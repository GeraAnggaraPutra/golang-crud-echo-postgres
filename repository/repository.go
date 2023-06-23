package repository

import (
	"crud-database-postgresql/models"
	"database/sql"
	"time"
)

type Repository interface {
	GetAll() ([]*models.User, error)
	FindById(id int) (*models.User, error)
	Create(user *models.User) error
	Update(user *models.User) (*models.User, error)
	Delete(user *models.User) error
}

type repository struct {
	db *sql.DB
}

func NewRepository(db *sql.DB) *repository {
	return &repository{db}
}

func (r *repository) GetAll() ([]*models.User, error) {
	users := []*models.User{}

	sqlQuery := "SELECT * FROM users"
	rows, err := r.db.Query(sqlQuery)
	if err != nil {
		return nil, err
	}
	defer rows.Close()

	for rows.Next() {
		user := &models.User{}
		var createdAt, updatedAt sql.NullTime
		err := rows.Scan(&user.Id, &user.Name, &user.Email, &user.Age, &user.Height, &createdAt, &updatedAt)
		if err != nil {
			return nil, err
		}
		if createdAt.Valid {
			user.CreatedAt = &createdAt.Time
		}
		if updatedAt.Valid {
			user.UpdatedAt = &updatedAt.Time
		}

		users = append(users, user)
	}

	return users, nil
}

func (r *repository) FindById(id int) (*models.User, error) {
	user := &models.User{}
	var createdAt, updatedAt sql.NullTime
	sqlQuery := "SELECT * FROM users WHERE id = $1"
	err := r.db.QueryRow(sqlQuery, id).Scan(&user.Id, &user.Name, &user.Email, &user.Age, &user.Height, &createdAt, &updatedAt)
	if err != nil {
		if err == sql.ErrNoRows {
			return nil, nil
		}
		return nil, err
	}
	if createdAt.Valid {
		user.CreatedAt = &createdAt.Time
	}
	if updatedAt.Valid {
		user.UpdatedAt = &updatedAt.Time
	}

	return user, nil
}

func (r *repository) Create(user *models.User) error {
	sqlQuery := "INSERT INTO users(name, email, age, height, created_at, updated_at) VALUES($1, $2, $3, $4, $5, $6)"
	_, err := r.db.Exec(sqlQuery, user.Name, user.Email, user.Age, user.Height, time.Now(), time.Now())
	if err != nil {
		return err
	}

	return nil
}

func (r *repository) Update(user *models.User) (*models.User, error) {
	sqlQuery := "UPDATE users SET name = $1, email = $2, age = $3, height = $4, updated_at = $5 WHERE id = $6"
	_, err := r.db.Exec(sqlQuery, user.Name, user.Email, user.Age, user.Height, time.Now(), user.Id)
	if err != nil {
		return nil, err
	}

	return user, nil
}

func (r *repository) Delete(user *models.User) error {
	sqlQuery := "DELETE FROM users WHERE id = $1"
	_, err := r.db.Exec(sqlQuery, user.Id)
	if err != nil {
		return err
	}

	return nil
}
