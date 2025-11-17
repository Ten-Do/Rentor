# Production OTP + JWT Authentication Guide

## Overview
Полная реализация OTP аутентификации (по email) с JWT токенами для безопасной работы в production.

## Архитектура (Flow)

```
1. User -> POST /auth/send-otp {email}
   ├─> Validate email format (RFC 5322)
   ├─> Find or Create User in DB
   ├─> Generate 6-digit OTP
   ├─> Hash OTP with bcrypt
   ├─> Store in OTP table with 10min expiration
   ├─> TODO: Send via email service (currently prints to console in dev)
   └─> Return 200 {message: "OTP sent"}

2. User -> POST /auth/verify-otp {email, otp_code}
   ├─> Get OTP record from DB
   ├─> Check if expired (delete if expired)
   ├─> Check if max attempts exceeded (delete if exceeded)
   ├─> Verify OTP using bcrypt.CompareHashAndPassword
   ├─> If valid:
   │  ├─> Get User from DB
   │  ├─> Generate JWT access token (15 min TTL)
   │  ├─> Set JWT in httpOnly secure cookie
   │  ├─> Delete OTP from DB
   │  └─> Return 200 {user: {...}}
   └─> If invalid: increment attempts, return 401

3. User -> GET /user/profile (with JWT cookie)
   ├─> AuthMiddleware validates JWT from cookie
   ├─> Extract userID from JWT claims
   ├─> Store userID in request context
   ├─> Handler gets userID from context
   ├─> Fetch user profile from DB
   └─> Return 200 {profile}

4. User -> POST /auth/logout
   ├─> Clear session cookie (set Expires=past)
   └─> Return 200
```

## Database Schema

### users table (existing)
```sql
CREATE TABLE user (
    id INTEGER PRIMARY KEY AUTOINCREMENT,
    email TEXT UNIQUE,
    phone_number TEXT UNIQUE,
    created_at DATETIME DEFAULT CURRENT_TIMESTAMP,
    updated_at DATETIME DEFAULT CURRENT_TIMESTAMP
);
```

### user_profile table (existing)
```sql
CREATE TABLE user_profile (
    id INTEGER PRIMARY KEY AUTOINCREMENT,
    user_id INTEGER NOT NULL UNIQUE,
    first_name TEXT,
    surname TEXT,
    patronymic TEXT,
    created_at DATETIME DEFAULT CURRENT_TIMESTAMP,
    updated_at DATETIME DEFAULT CURRENT_TIMESTAMP,
    FOREIGN KEY (user_id) REFERENCES user(id) ON DELETE CASCADE
);
```

### otp_codes table (NEW - migration 00003)
```sql
CREATE TABLE otp_codes (
    id INTEGER PRIMARY KEY AUTOINCREMENT,
    user_id INTEGER NOT NULL,
    email TEXT NOT NULL,
    code_hash TEXT NOT NULL,              -- bcrypt hash
    attempts INTEGER DEFAULT 0,            -- failed attempts counter
    max_attempts INTEGER DEFAULT 5,       -- max before lockout
    expires_at DATETIME NOT NULL,         -- 5-10 min from creation
    created_at DATETIME DEFAULT CURRENT_TIMESTAMP,
    FOREIGN KEY (user_id) REFERENCES user(id) ON DELETE CASCADE
);
```

## Config (config/local.yaml)

```yaml
auth:
  jwt_secret: "your-super-secret-key-change-in-production-at-least-32-chars"
  access_token_ttl: 15m            # JWT expiration
  refresh_token_ttl: 168h           # Refresh token TTL (optional)
  otp_length: 6                     # OTP digit count
  otp_expiration_minutes: 10        # OTP lifetime
  otp_max_attempts: 5               # Failed attempts before lockout
```

**ВАЖНО:** В production:
- `jwt_secret` должен быть минимум 32 символа, сгенерирован криптографически
- Хранить в переменных окружения, не в файле конфига
- Использовать разные ключи для разных окружений

## Security Checklist (Production)

✅ **OTP Security**
- [ ] OTP хранится в DB только в хешированном виде (bcrypt)
- [ ] OTP имеет expiration (5-10 минут)
- [ ] Rate limiting: max 5 попыток, потом lockout
- [ ] OTP удаляется после успешной верификации
- [ ] Cleanup expired OTP records (например, раз в час через cronjob)

✅ **JWT Security**
- [ ] JWT подписан HMAC-SHA256
- [ ] Access token TTL = 15 минут (не более 1 часа)
- [ ] Refresh token TTL = 7 дней (опционально)
- [ ] JWT передается только через secure httpOnly cookie
- [ ] Secure flag = true (только HTTPS)
- [ ] SameSite = Strict (защита от CSRF)

✅ **Cookie Security**
```go
http.SetCookie(w, &http.Cookie{
    Name:     "session",
    Value:    accessToken,
    Path:     "/",
    HttpOnly: true,        // Недоступен из JavaScript
    Secure:   true,        // Только через HTTPS
    SameSite: http.SameSiteStrictMode,  // CSRF protection
    Expires:  time.Now().Add(15 * time.Minute),
    Domain:   "yourdomain.com",  // Specify domain
})
```

✅ **Email OTP Delivery**
- [ ] Внедрить реальный email сервис (SendGrid, AWS SES, Mailgun и т.д.)
- [ ] В текущем коде OTP печатается в console (dev only!)
- [ ] Убрать логирование OTP в production
- [ ] Отправлять только через защищенное соединение (TLS)

✅ **User Privacy**
- [ ] Не логировать OTP коды
- [ ] Не логировать полные JWT токены
- [ ] Удалять старые OTP коды
- [ ] Использовать HTTPS везде

## Implementation Details

### 1. OTP Generation & Storage

```go
// Generate 6-digit OTP
otpCode, _ := rand.Intn(1000000)  // 000000-999999

// Hash before storing
hashedCode, _ := bcrypt.GenerateFromPassword(
    []byte(otpCode),
    bcrypt.DefaultCost,
)

// Store in DB with expiration
repo.CreateOTP(userID, email, hashedCode, maxAttempts, expiresAt)
```

**Почему bcrypt?**
- Медленный (защита от brute-force)
- Adaptive cost factor (будет медленнее со временем)
- Стандарт для password hashing

### 2. OTP Verification

```go
// Retrieve OTP from DB
otpRecord := repo.GetOTPByEmail(email)

// Check expiration
if time.Now().After(otpRecord.ExpiresAt) {
    repo.DeleteOTPByID(otpRecord.ID)
    return errors.New("OTP expired")
}

// Check attempts
if otpRecord.Attempts >= maxAttempts {
    repo.DeleteOTPByID(otpRecord.ID)
    return errors.New("too many attempts")
}

// Compare using bcrypt (timing-safe)
err := bcrypt.CompareHashAndPassword(
    []byte(otpRecord.CodeHash),
    []byte(providedOTP),
)
if err != nil {
    otpRecord.Attempts++
    repo.UpdateOTPAttempts(otpRecord.ID, otpRecord.Attempts)
    return err
}
```

### 3. JWT Token Generation

```go
// Create JWT claims
claims := JWTClaims{
    UserID: userID,
    Email:  email,
    Type:   "access",
    RegisteredClaims: jwt.RegisteredClaims{
        ExpiresAt: jwt.NewNumericDate(now.Add(15 * time.Minute)),
        IssuedAt:  jwt.NewNumericDate(now),
        NotBefore: jwt.NewNumericDate(now),
        Issuer:    "rentor",
    },
}

// Sign JWT
token := jwt.NewWithClaims(jwt.SigningMethodHS256, claims)
tokenString, _ := token.SignedString([]byte(jwtSecret))
```

### 4. JWT Validation Middleware

```go
// Extract JWT from cookie
cookie, err := r.Cookie("session")
if err != nil {
    return http.StatusUnauthorized
}

// Validate & extract claims
userID, err := jwtService.ValidateToken(cookie.Value)
if err != nil {
    return http.StatusUnauthorized
}

// Store in context for handler
ctx := context.WithValue(r.Context(), UserIDKey, userID)
next.ServeHTTP(w, r.WithContext(ctx))
```

## Endpoints Summary

| Method | Path | Auth Required | Description |
|--------|------|---------------|-------------|
| POST | /auth/send-otp | ❌ | Send OTP to email |
| POST | /auth/verify-otp | ❌ | Verify OTP, get JWT |
| POST | /auth/logout | ❌ | Clear session |
| GET | /user/profile | ✅ | Get user profile |
| PUT | /user/profile | ✅ | Update profile |

## Production Deployment Checklist

- [ ] Set `jwt_secret` from environment variable
- [ ] Set `Secure: true` in cookie (HTTPS only)
- [ ] Set `SameSite: Strict` in cookie
- [ ] Implement real email service for OTP
- [ ] Remove console logging of OTP
- [ ] Setup CORS properly (if needed)
- [ ] Add rate limiting middleware
- [ ] Setup database connection pooling
- [ ] Add metrics/monitoring for auth failures
- [ ] Add logging for security events
- [ ] Setup HTTPS/TLS certificates
- [ ] Consider IP-based rate limiting
- [ ] Implement refresh token rotation (optional)
- [ ] Add 2FA as additional layer (optional)

## Testing the Flow Locally

```bash
# 1. Start server
cd Backend
CONFIG_PATH=./config/local.yaml go run ./cmd/rentor

# 2. Send OTP (terminal 2)
curl -X POST http://localhost:8080/auth/send-otp \
  -H "Content-Type: application/json" \
  -d '{"email":"test@example.com"}'

# Console output shows OTP (e.g., 123456)

# 3. Verify OTP
curl -X POST http://localhost:8080/auth/verify-otp \
  -H "Content-Type: application/json" \
  -d '{"email":"test@example.com","otp_code":"123456"}' \
  -c cookies.txt

# 4. Access protected endpoint
curl -X GET http://localhost:8080/user/profile \
  -H "Content-Type: application/json" \
  -b cookies.txt

# 5. Logout
curl -X POST http://localhost:8080/auth/logout \
  -b cookies.txt
```

## Дополнительные улучшения (Future)

1. **Email Service Integration**
   - SendGrid, AWS SES, Mailgun
   - HTML templates для OTP письма
   - Retry logic

2. **Refresh Tokens**
   - Rotate access tokens
   - Store refresh tokens в DB для revocation

3. **Rate Limiting**
   - Per-IP rate limits
   - Per-email OTP sending limits
   - Per-user login attempt limits

4. **Enhanced Security**
   - Device fingerprinting
   - Suspicious activity detection
   - Email verification before first login

5. **Monitoring**
   - Alert on high OTP failure rates
   - Track failed login attempts
   - Monitor JWT validation failures

## Файлы, которые были добавлены/изменены

### Новые файлы:
- `internal/models/otp.go` — OTP model
- `internal/service/jwt_service.go` — JWT generation/validation
- `internal/service/otp_service.go` — OTP generation/verification
- `internal/service/user_profile_service.go` — User profile service
- `internal/repository/otp_repository.go` — OTP repository
- `internal/http-server/middleware/auth.go` — JWT validation middleware
- `migrations/00003_create_otp_table.sql` — OTP database table

### Измененные файлы:
- `internal/config/config.go` — Added Auth config struct
- `internal/service/interface.go` — Updated service interfaces
- `internal/service/user_service.go` — Implemented UserService
- `internal/repository/interface.go` — Added OTPRepository interface
- `internal/repository/user_repository.go` — Implemented UserRepository
- `internal/repository/user_profile_repository.go` — Implemented UserProfileRepository
- `internal/store/store.go` — Updated Store to include new services
- `internal/http-server/handlers/auth.go` — Implemented auth handlers
- `internal/http-server/handlers/user.go` — Implemented user profile handler
- `internal/http-server/routes.go` — Updated routes with middleware
- `internal/http-server/middleware/logger.go` — (no changes)
- `cmd/rentor/main.go` — Updated to pass config to store
- `config/local.yaml` — Added auth configuration
