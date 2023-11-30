package adapters

import (
	"database/sql"
	"golang-project-template/internal/users/domain"
	"log"
	"time"
)

type userRepository struct {
	db *sql.DB
	f  domain.UserFactory
}

func NewUserRepository(db *sql.DB) *userRepository {
	return &userRepository{db: db}
}

func (u *userRepository) Save(user *domain.User) (int, error) {

	var id int
	insertStatement := `
	INSERT INTO users (name, phone_number, password, role) 
	VALUES ($1, $2, $3, $4) 
	RETURNING id
	`
	err := u.db.QueryRow(insertStatement, user.GetName(), user.GetPhoneNumber(), user.GetPassword(), user.GetRole()).Scan(
		&id,
	)
	if err != nil {
		log.Printf("failed to execute Exec query: %v", err)
	}

	return id, nil
}

func (u *userRepository) FindOneByPhoneNumber(phone_number string) (*domain.User, error) {

	var id int
	var name string
	var password string
	var phoneNumber string
	var role string
	var createdAt time.Time
	var updatedAt time.Time
	var deletedAt *time.Time

	sqlStatement := `
        SELECT id, name, phone_number, password, role, created_at, updated_at, deleted_at
        FROM users
        WHERE id = $1
    `
	err := u.db.QueryRow(sqlStatement, phone_number).Scan(
		&id,
		&name,
		&phoneNumber,
		&password,
		&role,
		&createdAt,
		&updatedAt,
		&deletedAt,
	)
	if err != nil {
		return nil, err
	}

	newUser := u.f.ParseModelToDomain(id, name, phoneNumber, password, role, createdAt)

	return newUser, nil
}
