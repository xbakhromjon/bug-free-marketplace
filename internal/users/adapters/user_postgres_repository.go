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

func (u *userRepository) Create(user *domain.User) (*domain.User, error) {
	var createdUser domain.User
	insertStatement := `
	INSERT INTO users (name, phone_number, password, role) 
	VALUES ($1, $2, $3, $4) 
	RETURNING id, name, phone_number, role
	`
	err := u.db.QueryRow(insertStatement, user.Name, user.PhoneNumber, user.Password, user.Role).Scan(
		&createdUser.ID,
		&createdUser.Name,
		&createdUser.PhoneNumber,
		&createdUser.Role,
	)
	if err != nil {
		log.Printf("failed to execute Exec query: %v", err)
	}

	newUser := u.f.MapUserData(createdUser.ID, createdUser.Name, createdUser.PhoneNumber, createdUser.Role, time.Now())

	return newUser, nil
}

func (u *userRepository) GetByID(userID int) (*domain.User, error) {

	var user domain.User

	sqlStatement := `
        SELECT id, name, phone_number, password, role, created_at, updated_at, deleted_at
        FROM users
        WHERE id = $1
    `
	err := u.db.QueryRow(sqlStatement, userID).Scan(
		&user.ID,
		&user.Name,
		&user.PhoneNumber,
		&user.Password,
		&user.Role,
		&user.CreatedAt,
		&user.UpdatedAt,
		&user.DeletedAt,
	)
	if err != nil {
		return nil, err
	}
	return &user, nil
}

func (u *userRepository) GetByPhoneNumber(phone_number string) (*domain.User, error) {
	var user domain.User

	sqlStatement := `
        SELECT id, name, phone_number, password, role, created_at, updated_at, deleted_at
        FROM users
        WHERE id = $1
    `
	err := u.db.QueryRow(sqlStatement, phone_number).Scan(
		&user.ID,
		&user.Name,
		&user.PhoneNumber,
		&user.Password,
		&user.Role,
		&user.CreatedAt,
		&user.UpdatedAt,
		&user.DeletedAt,
	)
	if err != nil {
		return nil, err
	}

	return &user, nil
}

func (u *userRepository) GetAll() ([]*domain.User, error) {
	var users []*domain.User

	sqlStatement := `
        SELECT id, name, phone_number, password, role, created_at, updated_at, deleted_at
        FROM users
        WHERE id = $1
    `
	rows, err := u.db.Query(sqlStatement)
	if err != nil {
		return nil, err
	}
	rows.Close()

	for rows.Next() {
		var user domain.User
		err := rows.Scan(
			&user.ID,
			&user.Name,
			&user.PhoneNumber,
			&user.Password,
			&user.Role,
			&user.CreatedAt,
			&user.UpdatedAt,
			&user.DeletedAt,
		)
		if err != nil {
			return nil, err
		}
		users = append(users, &user)
	}

	return users, nil
}

func (u *userRepository) UpdateByID(id int, name, phoneNumber, pass, role string) (*domain.User, error) {
	sqlStatement := `
	UPDATE users
	SET name=$1, phone_number=$2, password=$3, role=$4
	WHERE id=$5
	`
	_, err := u.db.Exec(sqlStatement, name, phoneNumber, pass, role, id)
	if err != nil {
		return nil, err
	}

	var user domain.User
	getStatement := `
	SELECT id, name, phone_number, password, role
	FROM users
	WHERE id=$1
	`
	err1 := u.db.QueryRow(getStatement, id).Scan(
		&user.ID,
		&user.Name,
		&user.PhoneNumber,
		&user.Password,
		&user.Role,
	)
	if err1 != nil {
		return nil, err1
	}

	return &user, nil
}

func (u *userRepository) UpdateByPhoneNumber(name, phoneNumber, pass, role string) (*domain.User, error) {
	sqlStatement := `
	UPDATE users
	SET name=$1, password=$2, role=$3
	WHERE phone_number=$4
	`
	_, err := u.db.Exec(sqlStatement, name, pass, role, phoneNumber)
	if err != nil {
		return nil, err
	}

	var user domain.User
	getStatement := `
	SELECT id, name, phone_number, password, role
	FROM users
	WHERE phone_number=$1
	`
	err1 := u.db.QueryRow(getStatement, phoneNumber).Scan(
		&user.ID,
		&user.Name,
		&user.PhoneNumber,
		&user.Password,
		&user.Role,
	)
	if err1 != nil {
		return nil, err1
	}

	return &user, nil
}

func (u *userRepository) Delete(userID int) error {
	sqlStatement := `
	DELETE
	FROM users
	WHERE id=$1`
	_, err := u.db.Exec(sqlStatement, userID)
	if err != nil {
		return err
	}
	return nil
}
