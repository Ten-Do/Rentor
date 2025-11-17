package repository

import (
	"database/sql"
	"errors"
	"fmt"
	"rentor/internal/models"
	"strings"
)

type AdRepository struct {
	db *sql.DB
}

func NewAdRepository(db *sql.DB) AdRepository {
	return AdRepository{db: db}
}

//
// ============================
// CREATE
// ============================
//

func (r *AdRepository) CreateAdvertisement(userID int, ad *models.CreateAdvertisementInput) (int, error) {
	if ad.Title == "" {
		return 0, errors.New("title is required")
	}
	if ad.Price < 0 {
		return 0, errors.New("price must be >= 0")
	}
	res, err := r.db.Exec(`
        INSERT INTO advertisement 
        (user_id, title, description, price, type, rooms, city, address, latitude, longitude, square, status)
        VALUES (?, ?, ?, ?, ?, ?, ?, ?, ?, ?, ?, ?)
    `,
		userID,
		ad.Title,
		ad.Description,
		ad.Price,
		ad.Type,
		ad.Rooms,
		ad.City,
		ad.Address,
		ad.Latitude,
		ad.Longitude,
		ad.Square,
		"active",
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

func (r *AdRepository) CreateAdvertisementImages(adID int, urls []string) error {
	if len(urls) == 0 {
		return nil
	}

	for _, url := range urls {
		if url == "" {
			continue
		}
		_, err := r.db.Exec(`
            INSERT INTO advertisement_photos (advertisement_id, photo_url)
            VALUES (?, ?)
        `, adID, url)
		if err != nil {
			return err
		}
	}
	return nil
}

//
// ============================
// GET ONE (DETAIL)
// ============================
//

func (r *AdRepository) GetUserID(id int) (int, error) {
	var userID int
	err := r.db.QueryRow("SELECT user_id FROM advertisement WHERE id = ?", id).Scan(&userID)
	if err != nil {
		return 0, err
	}
	return userID, nil
}

func (r *AdRepository) GetAdvertisement(id int) (*models.GetAd, error) {
	ad := &models.GetAd{}
	var userID int

	err := r.db.QueryRow(`
        SELECT id, user_id, title, description, price, type, rooms, city, address,
               latitude, longitude, square, status
        FROM advertisement
        WHERE id = ?
    `,
		id,
	).Scan(
		&ad.ID,
		&userID,
		&ad.Title,
		&ad.Description,
		&ad.Price,
		&ad.Type,
		&ad.Rooms,
		&ad.City,
		&ad.Address,
		&ad.Latitude,
		&ad.Longitude,
		&ad.Square,
		&ad.Status,
	)
	if err != nil {
		if err == sql.ErrNoRows {
			return nil, errors.New("advertisement not found")
		}
		return nil, err
	}

	// вытаскиваем user_profile
	err = r.db.QueryRow(`
        SELECT first_name
        FROM user_profile
        WHERE user_id = ?
    `, userID).Scan(&ad.LandlordName)
	if err != nil {
		return nil, errors.New("user profile not found")
	}

	// вытаскиваем user
	err = r.db.QueryRow(`
        SELECT email, phone_number
        FROM user
        WHERE id = ?
    `, userID).Scan(&ad.LandlordEmail, &ad.LandlordPhone)
	if err != nil {
		if err == sql.ErrNoRows {
			return nil, errors.New("user not found")
		}
		return nil, err
	}

	// фото
	rows, err := r.db.Query(`
        SELECT photo_url 
        FROM advertisement_photos
        WHERE advertisement_id = ?
    `, ad.ID)
	if err != nil {
		return nil, err
	}
	defer rows.Close()

	var images []string
	for rows.Next() {
		var url string
		if err := rows.Scan(&url); err != nil {
			return nil, err
		}
		images = append(images, url)
	}
	ad.ImageUrls = images

	return ad, nil
}

//
// ============================
// LIST (FILTERS + PAGINATION)
// ============================
//

func (r *AdRepository) GetAdvertisementsPaged(filters *models.AdFilters) (*models.GetAdPreviewsList, error) {
	page := filters.Page
	limit := filters.Limit

	offset := (page - 1) * limit

	where := []string{"1=1"}
	args := []any{}

	if filters.MinPrice != nil {
		where = append(where, "price >= ?")
		args = append(args, filters.MinPrice)
	}
	if filters.MaxPrice != nil {
		where = append(where, "price <= ?")
		args = append(args, filters.MaxPrice)
	}
	if filters.Type != nil {
		where = append(where, "type = ?")
		args = append(args, filters.Type)
	}
	if filters.Rooms != nil {
		where = append(where, "rooms = ?")
		args = append(args, filters.Rooms)
	}
	if filters.City != nil {
		where = append(where, "city = ?")
		args = append(args, filters.City)
	}
	if filters.Keywords != nil {
		where = append(where, "(title LIKE ? OR description LIKE ?)")
		kw := "%" + *filters.Keywords + "%"
		args = append(args, kw, kw)
	}

	query := fmt.Sprintf(`
        SELECT id, title, price, city, type, rooms
        FROM advertisement
        WHERE %s
        ORDER BY created_at DESC
        LIMIT ? OFFSET ?
    `, strings.Join(where, " AND "))

	args = append(args, limit, offset)

	rows, err := r.db.Query(query, args...)
	if err != nil {
		return nil, err
	}
	defer rows.Close()

	list := &models.GetAdPreviewsList{}
	for rows.Next() {
		var item models.AdPreview
		if err := rows.Scan(
			&item.ID,
			&item.Title,
			&item.Price,
			&item.City,
			&item.Type,
			&item.Rooms,
		); err != nil {
			return nil, err
		}

		// first photo
		_ = r.db.QueryRow(`
            SELECT photo_url
            FROM advertisement_photos
            WHERE advertisement_id = ?
            LIMIT 1
        `, item.ID).Scan(&item.ImageUrl)

		list.Items = append(list.Items, item)
	}

	// count
	var total int
	_ = r.db.QueryRow(`
        SELECT COUNT(*) FROM advertisement
        WHERE `+strings.Join(where, " AND "),
		args[:len(args)-2]...,
	).Scan(&total)

	list.Total = total
	list.Page = page
	list.Limit = limit

	return list, nil
}

//
// ============================
// UPDATE
// ============================
//

func (r *AdRepository) UpdateAdvertisement(id int, ad *models.UpdateAdvertisementInput) error {
	_, err := r.db.Exec(`
        UPDATE advertisement
        SET title = ?, description = ?, price = ?, type = ?, rooms = ?, city = ?, 
            address = ?, latitude = ?, longitude = ?, square = ?, status = ?, 
            updated_at = CURRENT_TIMESTAMP
        WHERE id = ?
    `,
		ad.Title,
		ad.Description,
		ad.Price,
		ad.Type,
		ad.Rooms,
		ad.City,
		ad.Address,
		ad.Latitude,
		ad.Longitude,
		ad.Square,
		ad.Status,
		id,
	)
	return err
}

//
// ============================
// DELETE
// ============================
//

func (r *AdRepository) DeleteAdvertisement(id int) error {
	_, err := r.db.Exec("DELETE FROM advertisement WHERE id = ?", id)
	return err
}

func (r *AdRepository) DeleteAdvertisementImage(adID, imageID int) error {
	_, err := r.db.Exec(`
        DELETE FROM advertisement_photos 
        WHERE id = ? AND advertisement_id = ?
    `, imageID, adID)
	return err
}
