package app

import (
	"golang-project-template/internal/users/domain"
	"reflect"
	"testing"
	"time"
)

func TestRegisterMerchantUser(t *testing.T) {
	underTest := userUsecase{userRepository: &mockUserRepo{}}

	t.Run("success register merchant user", func(t *testing.T) {
		user := mockUserForTestRegister()
		got, err := underTest.RegisterMerchantUser(user)
		want := 1

		assertEqual[int](t, got, want)
		assertEqual[error](t, err, nil)
	})

	testCases := []struct {
		name          string
		user          *domain.NewUser
		expectedError error
	}{
		{
			name:          "failed to merchant user (existing phone number)",
			user:          mockUserWithPhoneForTestRegister("phone"),
			expectedError: domain.ErrPhoneNumberExists,
		},
		{
			name:          "failed to merchant user (empty name)",
			user:          mockEmptyNameUserForTestRegister(),
			expectedError: domain.ErrEmptyUserName,
		},
		{
			name:          "failed to merchant user (empty phone)",
			user:          mockUserWithPhoneForTestRegister(""),
			expectedError: domain.ErrEmptyPhoneNumber,
		},
	}

	for _, tCase := range testCases {
		t.Run(tCase.name, func(t *testing.T) {
			_, err := underTest.RegisterMerchantUser(tCase.user)

			assertEqual[error](t, err, tCase.expectedError)
		})
	}
}
func TestRegisterCustomer(t *testing.T) {
	underTest := userUsecase{userRepository: &mockUserRepo{}}

	t.Run("success register customer", func(t *testing.T) {
		user := mockUserForTestRegister()
		got, err := underTest.RegisterCustomer(user)
		want := 1

		assertEqual[int](t, got, want)
		assertEqual[error](t, err, nil)
	})

	testCases := []struct {
		name          string
		user          *domain.NewUser
		expectedError error
	}{
		{
			name:          "failed to register customer (existing phone number)",
			user:          mockUserWithPhoneForTestRegister("phone"),
			expectedError: domain.ErrPhoneNumberExists,
		},
		{
			name:          "failed to register customer (empty name)",
			user:          mockEmptyNameUserForTestRegister(),
			expectedError: domain.ErrEmptyUserName,
		},
		{
			name:          "failed to register customer (empty phone)",
			user:          mockUserWithPhoneForTestRegister(""),
			expectedError: domain.ErrEmptyPhoneNumber,
		},
	}

	for _, tCase := range testCases {
		t.Run(tCase.name, func(t *testing.T) {
			_, err := underTest.RegisterCustomer(tCase.user)

			assertEqual[error](t, err, tCase.expectedError)
		})
	}
}
func TestRegisterAdmin(t *testing.T) {
	underTest := userUsecase{userRepository: &mockUserRepo{}}

	t.Run("success register admin", func(t *testing.T) {
		user := mockUserForTestRegister()
		got, err := underTest.RegisterAdmin(user)
		want := 1

		assertEqual[int](t, got, want)
		assertEqual[error](t, err, nil)
	})

	testCases := []struct {
		name          string
		user          *domain.NewUser
		expectedError error
	}{
		{
			name:          "failed to register admin (existing phone number)",
			user:          mockUserWithPhoneForTestRegister("phone"),
			expectedError: domain.ErrPhoneNumberExists,
		},
		{
			name:          "failed to register admin (empty name)",
			user:          mockEmptyNameUserForTestRegister(),
			expectedError: domain.ErrEmptyUserName,
		},
		{
			name:          "failed to register admin (empty phone)",
			user:          mockUserWithPhoneForTestRegister(""),
			expectedError: domain.ErrEmptyPhoneNumber,
		},
	}

	for _, tCase := range testCases {
		t.Run(tCase.name, func(t *testing.T) {
			_, err := underTest.RegisterAdmin(tCase.user)

			assertEqual[error](t, err, tCase.expectedError)
		})
	}
}
func TestLoginUser(t *testing.T) {
	underTest := userUsecase{userRepository: &mockUserRepo{}}

	t.Run("success login user", func(t *testing.T) {

		got, err := underTest.LoginUser("998990970138", "Golang123") // true
		want := true

		assertEqual[bool](t, got, want)
		assertEqual[error](t, err, nil)
	})

	testCases := []struct {
		name          string
		phoneNumber   string
		password      string
		expectedError error
	}{
		{
			name:          "failed login user (empty phone number)",
			phoneNumber:   "",
			password:      "Golang123",
			expectedError: domain.ErrInvalidCredentials,
		},
		{
			name:          "failed login user (empty password)",
			phoneNumber:   "998990970138",
			password:      "",
			expectedError: domain.ErrInvalidCredentials,
		},
		{
			name:          "failed login user (incorrect credentials)",
			phoneNumber:   "998990970138",
			password:      "wrongpassword",
			expectedError: domain.ErrInvalidCredentials,
		},
		{
			name:          "failed login user (do not existing phone number)",
			phoneNumber:   "phone_number",
			password:      "Golang123",
			expectedError: domain.ErrUserNotFound,
		},
	}

	for _, tCase := range testCases {
		t.Run(tCase.name, func(t *testing.T) {
			_, err := underTest.LoginUser(tCase.phoneNumber, tCase.password)
			assertEqual[error](t, err, tCase.expectedError)
		})
	}
}
func TestGetUserByPhoneNumber(t *testing.T) {
	underTest := userUsecase{userRepository: &mockUserRepo{}}

	t.Run("success get user by phone", func(t *testing.T) {

		got, err := underTest.GetUserByPhoneNumber("998990970138")
		wantedUser := createUserWithPhoneNumber("998990970138")

		assertEqual[error](t, err, nil)
		assertUserEquality[*domain.User](t, got, wantedUser)

	})

	t.Run("get user by phone (user not found by phone)", func(t *testing.T) {
		_, err := underTest.GetUserByPhoneNumber("phone_number")
		assertEqual[error](t, err, domain.ErrUserNotFound)
	})

	t.Run("get user by phone (empty phone number)", func(t *testing.T) {
		_, err := underTest.GetUserByPhoneNumber("")
		assertEqual[error](t, err, domain.ErrEmptyPhoneNumber)
	})
}
func TestGetUserByID(t *testing.T) {
	underTest := userUsecase{userRepository: &mockUserRepo{}}

	t.Run("success get user by ID", func(t *testing.T) {

		got, err := underTest.GetUserByID(1)
		wantedUser := createUserWithPhoneNumber("phone")

		assertEqual[error](t, err, nil)
		assertUserEquality[*domain.User](t, got, wantedUser)
	})

	t.Run("get user by ID (user not found by ID)", func(t *testing.T) {
		_, err := underTest.GetUserByID(2)
		assertEqual[error](t, err, domain.ErrUserNotFound)
	})
}
func TestUserExists(t *testing.T) {
	underTest := userUsecase{userRepository: &mockUserRepo{}}

	t.Run("success user exists", func(t *testing.T) {

		got, err := underTest.UserExists(1)

		assertEqual[error](t, err, nil)
		assertUserEquality[bool](t, got, true)
	})

	t.Run("get user by ID (user not found by ID)", func(t *testing.T) {
		_, err := underTest.UserExists(2)
		assertEqual[error](t, err, domain.ErrUserNotFound)
	})
}

func assertUserEquality[T comparable](t *testing.T, got, wanted T) {
	t.Helper()
	if !reflect.DeepEqual(got, wanted) {
		t.Errorf("wanted: %+v but got: %+v", wanted, got)
	}
}

func assertEqual[T comparable](t testing.TB, got, want T) {
	t.Helper()
	if got != want {
		t.Errorf("want %+v but got %+v", want, got)
	}
	if !reflect.DeepEqual(got, want) {
		t.Errorf("want error '%v' but got: %v", want, got)
	}
}

func mockUserForTestRegister() *domain.NewUser {
	newUser := &domain.NewUser{}
	newUser.Name = "Quvonchbek"
	newUser.PhoneNumber = "998990970138"
	newUser.Password = "Golang123"

	return newUser
}

func mockEmptyNameUserForTestRegister() *domain.NewUser {
	newUser := &domain.NewUser{}
	newUser.Name = ""
	newUser.PhoneNumber = "998990970138"
	newUser.Password = "Golang123"

	return newUser
}

func mockUserWithPhoneForTestRegister(phone string) *domain.NewUser {
	newUser := &domain.NewUser{}
	newUser.Name = "Quvonchbek"
	newUser.PhoneNumber = phone
	newUser.Password = "Golang123"

	return newUser
}

func createUserWithPhoneNumber(phoneNumber string) *domain.User {
	user := &domain.User{}
	user.SetID(1)
	user.SetName("Quvonchbek")
	user.SetPhoneNumber(phoneNumber)
	user.SetPassword("$2a$10$IVxqpfdNU/8txnGQK6pu3.sD2TiMyZqRyLmZB/rPytMSIG1qQR/Wy")
	user.SetRole("user")
	user.SetCreateAt(time.Time{})
	user.SetUpdatedAt(time.Time{})
	user.SetDeletedAt(nil)

	return user
}

type mockUserRepo struct{}

func newMockUserRepo() domain.UserRepository {
	return &mockUserRepo{}
}

func (m *mockUserRepo) Save(user *domain.User) (int, error) {
	if user.GetPhoneNumber() == "phone" {
		return 0, domain.ErrPhoneNumberExists
	}
	return 1, nil
}

func (m *mockUserRepo) FindOneByPhoneNumber(phoneNumber string) (*domain.User, error) {
	var u *domain.User
	if phoneNumber == "phone_number" {
		return nil, domain.ErrUserNotFound
	}

	u = createUserWithPhoneNumber(phoneNumber)
	return u, nil
}

func (m *mockUserRepo) FindByID(userID int) (*domain.User, error) {
	u := createUserWithPhoneNumber("phone")
	if userID != 1 {
		return nil, domain.ErrUserNotFound
	}

	return u, nil
}

func (m *mockUserRepo) UserExists(userID int) (bool, error) {
	if userID != 1 {
		return false, domain.ErrUserNotFound
	}

	return true, nil
}

func (m *mockUserRepo) UserExistByPhone(phone string) (bool, error) {
	if phone == "phone_number" {
		return false, domain.ErrUserNotFound
	}

	return true, nil
}
