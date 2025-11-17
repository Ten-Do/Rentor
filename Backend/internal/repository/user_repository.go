package repository

import (
	"database/sql"
	"errors"
	"net/mail"
	"regexp"
	"rentor/internal/models"
	"strings"
)

// userRepository implements UserRepository
type userRepository struct {
	db *sql.DB
}

// NewUserRepository creates a new user repository
func NewUserRepository(db *sql.DB) UserRepository {
	return &userRepository{db: db}
}

// CreateUser creates a new user
func (r *userRepository) CreateUser(phone string, email string) (int, error) {
	if phone == "" && email == "" {
		return 0, errors.New("phone or email required")
	}

	// validate email
	if email != "" {
		err := validateEmail(email)
		if err != nil {
			return 0, err
		}
	}
	// validate phone
	if phone != "" {
		err := validatePhone(phone)
		if err != nil {
			return 0, err
		}
	}
	// turn email to lower register
	if email != "" {
		email = toLowerRegister(email)
	}

	var res sql.Result
	var err error

	if email != "" && phone == "" {
		res, err = r.db.Exec("INSERT INTO user (email) VALUES (?)", email)
	} else if phone != "" && email == "" {
		res, err = r.db.Exec("INSERT INTO user (phone_number) VALUES (?)", phone)
	} else {
		res, err = r.db.Exec("INSERT INTO user (email, phone_number) VALUES (?, ?)", email, phone)
	}

	if err != nil {
		return 0, err
	}

	id64, err := res.LastInsertId()
	if err != nil {
		return 0, err
	}

	return int(id64), nil
}

// GetUserByID retrieves a user by ID
func (r *userRepository) GetUserByID(id int) (*models.User, error) {
	user := &models.User{}
	err := r.db.QueryRow(
		"SELECT id, email, phone_number, created_at, updated_at FROM user WHERE id = ?",
		id,
	).Scan(&user.UserID, &user.Email, &user.Phone, &user.CreatedAt, &user.UpdatedAt)

	if err != nil {
		if err == sql.ErrNoRows {
			return nil, errors.New("user not found")
		}
		return nil, err
	}

	return user, nil
}

// GetUserByEmail retrieves a user by email
func (r *userRepository) GetUserByEmail(email string) (*models.User, error) {
	// validate email
	err := validateEmail(email)
	if err != nil {
		return nil, err
	}
	email = toLowerRegister(email)

	user := &models.User{}
	err = r.db.QueryRow(
		"SELECT id, email, phone_number, created_at, updated_at FROM user WHERE email = ?",
		email,
	).Scan(&user.UserID, &user.Email, &user.Phone, &user.CreatedAt, &user.UpdatedAt)

	if err != nil {
		if err == sql.ErrNoRows {
			return nil, nil // Not found, but no error
		}
		return nil, err
	}

	return user, nil
}

// GetUserByPhone retrieves a user by phone
func (r *userRepository) GetUserByPhone(phone string) (*models.User, error) {
	// validate phone
	err := validatePhone(phone)
	if err != nil {
		return nil, err
	}

	user := &models.User{}
	err = r.db.QueryRow(
		"SELECT id, email, phone_number, created_at, updated_at FROM user WHERE phone_number = ?",
		phone,
	).Scan(&user.UserID, &user.Email, &user.Phone, &user.CreatedAt, &user.UpdatedAt)

	if err != nil {
		if err == sql.ErrNoRows {
			return nil, nil
		}
		return nil, err
	}

	return user, nil
}

// GetAllUsers retrieves all users
func (r *userRepository) GetAllUsers() ([]*models.User, error) {
	rows, err := r.db.Query("SELECT id, email, phone_number, created_at, updated_at FROM user")
	if err != nil {
		return nil, err
	}
	defer rows.Close()

	var users []*models.User
	for rows.Next() {
		user := &models.User{}
		if err := rows.Scan(&user.UserID, &user.Email, &user.Phone, &user.CreatedAt, &user.UpdatedAt); err != nil {
			return nil, err
		}
		users = append(users, user)
	}

	return users, rows.Err()
}

// GetPageUsers retrieves users with pagination
func (r *userRepository) GetPageUsers(offset, limit int) ([]*models.User, error) {
	rows, err := r.db.Query(
		"SELECT id, email, phone_number, created_at, updated_at FROM user LIMIT ? OFFSET ?",
		limit,
		offset,
	)
	if err != nil {
		return nil, err
	}
	defer rows.Close()

	var users []*models.User
	for rows.Next() {
		user := &models.User{}
		if err := rows.Scan(&user.UserID, &user.Email, &user.Phone, &user.CreatedAt, &user.UpdatedAt); err != nil {
			return nil, err
		}
		users = append(users, user)
	}

	return users, rows.Err()
}

// UpdateUser updates a user
func (r *userRepository) UpdateUser(id int, user *models.User) error {
	// validate email
	err := validateEmail(user.Email)
	if err != nil {
		return err
	}
	user.Email = toLowerRegister(user.Email)

	// validate phone
	if user.Phone != nil {
		err = validatePhone(*user.Phone)
		if err != nil {
			return err
		}
	}

	_, err = r.db.Exec(
		"UPDATE user SET email = ?, phone_number = ?, updated_at = CURRENT_TIMESTAMP WHERE id = ?",
		user.Email,
		user.Phone,
		id,
	)
	return err
}

// DeleteUserByID deletes a user by ID
func (r *userRepository) DeleteUserByID(id int) error {
	_, err := r.db.Exec("DELETE FROM user WHERE id = ?", id)
	return err
}

// DeleteUserByPhone deletes a user by phone
func (r *userRepository) DeleteUserByPhone(phone string) error {
	// validate phone
	err := validatePhone(phone)
	if err != nil {
		return err
	}

	_, err = r.db.Exec("DELETE FROM user WHERE phone_number = ?", phone)
	return err
}

// DeleteUserByEmail deletes a user by email
func (r *userRepository) DeleteUserByEmail(email string) error {
	// validate email
	err := validateEmail(email)
	if err != nil {
		return err
	}
	email = toLowerRegister(email)

	_, err = r.db.Exec("DELETE FROM user WHERE email = ?", email)
	return err
}

// validateEmail проверяет корректность email
func validateEmail(email string) error {
	if email == "" {
		return errors.New("email is empty")
	}

	// Используем стандартный пакет net/mail
	_, err := mail.ParseAddress(email)
	if err != nil {
		return errors.New("invalid email format")
	}

	return nil
}

// validatePhone проверяет корректность телефона
func validatePhone(phone string) error {
	if phone == "" {
		return errors.New("phone is empty")
	}

	// Простая проверка: цифры, +, -, пробелы
	re := regexp.MustCompile(`^\+?\d[\d\s\-]{7,14}\d$`)
	if !re.MatchString(phone) {
		return errors.New("invalid phone format")
	}

	return nil
}

// toLowerRegister приводит email к нижнему регистру
func toLowerRegister(email string) string {
	return strings.ToLower(email)
}
