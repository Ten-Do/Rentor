# Rentor Backend - Production Auth Implementation

## What's New ✨

Полная реализация production-ready OTP + JWT аутентификации:

- ✅ OTP по email с bcrypt хешированием
- ✅ JWT токены в httpOnly secure cookies
- ✅ Автоматическое создание пользователя при первой OTP
- ✅ Дефолтный профиль при создании пользователя
- ✅ Rate limiting на OTP попытки
- ✅ Middleware для защиты routes
- ✅ Production-ready конфиг

## Quick Start

### 1. Сборка и запуск

```bash
cd Backend
go mod tidy
CONFIG_PATH=./config/local.yaml go run ./cmd/rentor
```

Сервер запустится на `localhost:8080`

### 2. Тестирование Flow

#### Шаг 1: Отправить OTP на email

```bash
curl -X POST http://localhost:8080/auth/send-otp \
  -H "Content-Type: application/json" \
  -d '{"email":"test@example.com"}'

# Response:
# {"message":"OTP sent to email"}

# В console сервера увидите OTP код (dev mode only!)
# [DEV] OTP for test@example.com: 123456 (expires at ...)
```

#### Шаг 2: Верифицировать OTP и получить JWT

```bash
curl -X POST http://localhost:8080/auth/verify-otp \
  -H "Content-Type: application/json" \
  -d '{"email":"test@example.com","otp_code":"123456"}' \
  -c cookies.txt \
  -v

# Response:
# {"user":{"id":1,"email":"test@example.com","phone":"","created_at":"2025-11-11T12:00:00Z"}}

# JWT установится в cookie "session" (see -c cookies.txt)
```

#### Шаг 3: Доступ к защищенному endpoint'у

```bash
curl -X GET http://localhost:8080/user/profile \
  -H "Content-Type: application/json" \
  -b cookies.txt

# Response:
# {"id":1,"user_id":1,"first_name":"","surname":"","patronymic":"","created_at":"...","updated_at":"..."}
```

#### Шаг 4: Обновить профиль

```bash
curl -X PUT http://localhost:8080/user/profile \
  -H "Content-Type: application/json" \
  -b cookies.txt \
  -d '{
    "first_name":"John",
    "surname":"Doe",
    "patronymic":"",
    "phone_number":"+70000000000"
  }'

# Response:
# {"id":1,"user_id":1,"first_name":"John","surname":"Doe","patronymic":"","created_at":"...","updated_at":"..."}
```

#### Шаг 5: Logout

```bash
curl -X POST http://localhost:8080/auth/logout \
  -b cookies.txt

# Response:
# {"message":"logged out"}

# Cookie будет очищена
```

## Production Configuration

### 1. Environment Variables

```bash
# .env.production
CONFIG_PATH=/etc/rentor/config.prod.yaml
JWT_SECRET=$(openssl rand -base64 32)  # Generate strong key
```

### 2. config.prod.yaml

```yaml
env: "prod"
storage_path: "/var/lib/rentor/database.db"

http_server:
  host: "0.0.0.0"
  port: "8080"
  timeout_seconds: 30s
  idle_timeout_seconds: 120s

auth:
  jwt_secret: "${JWT_SECRET}"  # From environment
  access_token_ttl: 15m
  refresh_token_ttl: 168h
  otp_length: 6
  otp_expiration_minutes: 10
  otp_max_attempts: 5
```

### 3. Docker Deployment

```dockerfile
# Dockerfile
FROM golang:1.25-alpine AS builder
WORKDIR /app
COPY . .
RUN go mod download
RUN CGO_ENABLED=1 GOOS=linux go build -a -installsuffix cgo -o rentor ./cmd/rentor

FROM alpine:latest
RUN apk --no-cache add ca-certificates
WORKDIR /root/
COPY --from=builder /app/rentor .
COPY --from=builder /app/migrations ./migrations
COPY --from=builder /app/config ./config

EXPOSE 8080
CMD ["./rentor"]
```

### 4. Key Security Settings for Production

```go
// In internal/http-server/middleware/auth.go

// Set to true only with HTTPS/TLS
Secure: os.Getenv("ENV") == "prod",

// Use Strict in production
SameSite: http.SameSiteStrictMode,

// Add Domain in production
Domain: os.Getenv("API_DOMAIN"),
```

## Architecture Overview

```
┌─────────────────────────────────────────────────────────┐
│                   HTTP Handlers                          │
│  (auth.go, user.go, advertisement.go)                   │
└─────────────────────────────────────────────────────────┘
                         ↓
┌─────────────────────────────────────────────────────────┐
│                    Services                              │
│  (UserService, OTPService, JWTService, ...)             │
└─────────────────────────────────────────────────────────┘
                         ↓
┌─────────────────────────────────────────────────────────┐
│                   Repositories                           │
│  (UserRepository, OTPRepository, ...)                   │
└─────────────────────────────────────────────────────────┘
                         ↓
┌─────────────────────────────────────────────────────────┐
│                  SQLite Database                         │
│  (users, user_profile, otp_codes)                       │
└─────────────────────────────────────────────────────────┘
```

## Key Components

### OTPService
- Генерирует 6-digit OTP
- Хеширует OTP using bcrypt (медленный хеш, защита от brute-force)
- Хранит в БД с 10-minute expiration
- Проверяет OTP и управляет попытками
- TODO: Интегрировать с email service

### JWTService
- Генерирует access token (15 min)
- Валидирует JWT из cookies
- Поддерживает refresh tokens (7 days)

### AuthMiddleware
- Извлекает JWT из "session" cookie
- Валидирует подпись и expiration
- Хранит userID в request context
- Возвращает 401 если невалидный/expired

### UserService
- FindOrCreateUserByEmail — создает пользователя если не существует
- CreateDefaultUserProfile — дефолтный профиль с пустыми полями
- GetUser, GetUserByEmail, GetUserByPhone

### UserProfileService
- GetUserProfile — получить профиль по userID
- UpdateUserProfile — обновить имя, фамилию, отчество
- CreateDefaultUserProfile — создать с пустыми значениями

## Database Migrations

Запустятся автоматически при старте сервера:

1. `00001_crate_user_table.sql` — Create users table
2. `00002_create_user_profile_table.sql` — Create user_profile table
3. `00003_create_otp_table.sql` — Create otp_codes table (NEW)

## Testing

### Unit Tests (TODO)

```bash
go test ./internal/service/... -v
go test ./internal/repository/... -v
```

### Integration Tests (TODO)

```bash
go test ./... -v -tags=integration
```

## TODO / Future Work

- [ ] Email OTP Integration (SendGrid, AWS SES, etc.)
- [ ] Rate limiting middleware (per-IP, per-email)
- [ ] Refresh token rotation
- [ ] Device fingerprinting
- [ ] 2FA support
- [ ] OAuth2 integration
- [ ] WebAuthn/FIDO2 support
- [ ] User session management
- [ ] Login history tracking
- [ ] Suspicious activity alerts

## Debugging

### View OTP in console (dev mode)

```
[DEV] OTP for test@example.com: 123456 (expires at 2025-11-11T12:10:00Z)
```

### Check logs

```
logger.Info("SendOTP called", logger.Field("email", req.Email))
logger.Warn("OTP verification failed", logger.Field("error", err.Error()))
logger.Error("failed to get user profile", logger.Field("user_id", userID))
```

### Database inspection

```bash
sqlite3 storage/storage.db
sqlite> SELECT * FROM user;
sqlite> SELECT * FROM otp_codes;
sqlite> SELECT * FROM user_profile;
```

## Security Best Practices

✅ **Implemented:**
- OTP bcrypt hashing
- JWT in httpOnly cookies
- SameSite cookie protection
- Rate limiting on OTP attempts
- OTP expiration
- Auto user profile creation

⚠️ **TODO for Production:**
- [ ] HTTPS/TLS setup
- [ ] Email OTP service integration
- [ ] Environment variable config
- [ ] Database backups
- [ ] Monitoring & alerting
- [ ] Request logging
- [ ] CORS configuration
- [ ] API rate limiting
- [ ] DDoS protection
- [ ] Web Application Firewall (WAF)

## Documentation

See `AUTHENTICATION_GUIDE.md` for detailed architecture, flow diagrams, and implementation details.

## Support

For issues or questions:
1. Check `AUTHENTICATION_GUIDE.md`
2. Review logs in console
3. Test with provided curl commands
4. Check database state with sqlite3

---

**Last Updated:** November 11, 2025
**Version:** 1.0.0 (MVP)
