package repository

import (
	"database/sql"
	"errors"
	"rentor/internal/models"
)

// userProfileRepository implements UserProfileRepository
type userProfileRepository struct {
	db *sql.DB
}

// NewUserProfileRepository creates a new user profile repository
func NewUserProfileRepository(db *sql.DB) UserProfileRepository {
	return &userProfileRepository{db: db}
}

// CreateUserProfile creates a new user profile
func (r *userProfileRepository) CreateUserProfile(profile *models.UserProfile) (int, error) {
	res, err := r.db.Exec(
		"INSERT INTO user_profile (user_id, first_name, surname, patronymic) VALUES (?, ?, ?, ?)",
		profile.UserID,
		profile.FirstName,
		profile.Surname,
		profile.Patronymic,
	)
	if err != nil {
		return 0, err
	}

	id64, err := res.LastInsertId()
	if err != nil {
		return 0, err
	}

	return int(id64), nil
}

// GetUserProfileByID retrieves a user profile by ID
func (r *userProfileRepository) GetUserProfileByID(id int) (*models.UserProfile, error) {
	profile := &models.UserProfile{}
	err := r.db.QueryRow(
		"SELECT id, user_id, COALESCE(first_name, ''), COALESCE(surname, ''), COALESCE(patronymic, ''), created_at, updated_at FROM user_profile WHERE id = ?",
		id,
	).Scan(&profile.ID, &profile.UserID, &profile.FirstName, &profile.Surname, &profile.Patronymic, &profile.CreatedAt, &profile.UpdatedAt)

	if err != nil {
		if err == sql.ErrNoRows {
			return nil, errors.New("user profile not found")
		}
		return nil, err
	}

	return profile, nil
}

// GetUserProfileByUserID retrieves a user profile by user ID
func (r *userProfileRepository) GetUserProfileByUserID(userID int) (*models.UserProfile, error) {
	profile := &models.UserProfile{}
	err := r.db.QueryRow(
		"SELECT id, user_id, COALESCE(first_name, ''), COALESCE(surname, ''), COALESCE(patronymic, ''), created_at, updated_at FROM user_profile WHERE user_id = ?",
		userID,
	).Scan(&profile.ID, &profile.UserID, &profile.FirstName, &profile.Surname, &profile.Patronymic, &profile.CreatedAt, &profile.UpdatedAt)

	if err != nil {
		if err == sql.ErrNoRows {
			return nil, nil // Not found, but no error
		}
		return nil, err
	}

	return profile, nil
}

// GetAllUserProfiles retrieves all user profiles
func (r *userProfileRepository) GetAllUserProfiles() ([]*models.UserProfile, error) {
	rows, err := r.db.Query(
		"SELECT id, user_id, COALESCE(first_name, ''), COALESCE(surname, ''), COALESCE(patronymic, ''), created_at, updated_at FROM user_profile",
	)
	if err != nil {
		return nil, err
	}
	defer rows.Close()

	var profiles []*models.UserProfile
	for rows.Next() {
		profile := &models.UserProfile{}
		if err := rows.Scan(&profile.ID, &profile.UserID, &profile.FirstName, &profile.Surname, &profile.Patronymic, &profile.CreatedAt, &profile.UpdatedAt); err != nil {
			return nil, err
		}
		profiles = append(profiles, profile)
	}

	return profiles, rows.Err()
}

// GetPageUserProfiles retrieves user profiles with pagination
func (r *userProfileRepository) GetPageUserProfiles(offset, limit int) ([]*models.UserProfile, error) {
	rows, err := r.db.Query(
		"SELECT id, user_id, COALESCE(first_name, ''), COALESCE(surname, ''), COALESCE(patronymic, ''), created_at, updated_at FROM user_profile LIMIT ? OFFSET ?",
		limit,
		offset,
	)
	if err != nil {
		return nil, err
	}
	defer rows.Close()

	var profiles []*models.UserProfile
	for rows.Next() {
		profile := &models.UserProfile{}
		if err := rows.Scan(&profile.ID, &profile.UserID, &profile.FirstName, &profile.Surname, &profile.Patronymic, &profile.CreatedAt, &profile.UpdatedAt); err != nil {
			return nil, err
		}
		profiles = append(profiles, profile)
	}

	return profiles, rows.Err()
}

// UpdateUserProfile updates a user profile
func (r *userProfileRepository) UpdateUserProfile(id int, profile *models.UserProfile) error {
	_, err := r.db.Exec(
		"UPDATE user_profile SET first_name = ?, surname = ?, patronymic = ?, updated_at = CURRENT_TIMESTAMP WHERE id = ?",
		profile.FirstName,
		profile.Surname,
		profile.Patronymic,
		id,
	)
	return err
}

// DeleteUserProfileByID deletes a user profile by ID
func (r *userProfileRepository) DeleteUserProfileByID(id int) error {
	_, err := r.db.Exec("DELETE FROM user_profile WHERE id = ?", id)
	return err
}
