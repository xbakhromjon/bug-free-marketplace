package app

import (
	"golang-project-template/internal/users/domain"
	"reflect"
	"testing"
	"time"
)

func TestRegisterUser(t *testing.T) {
	underTest := userUsecase{userRepository: &mockUserRepo{}}

	t.Run("success register user", func(t *testing.T) {
		user := mockUserForTestRegister()
		user.SetPhoneNumber("phone")
		got, err := underTest.RegisterUser(user)
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
			name:          "failed to register user (existing phone number)",
			user:          mockUserForTestRegister(),
			expectedError: domain.ErrPhoneNumberExists,
		},
		{
			name:          "failed to register user (empty name)",
			user:          mockEmptyNameUserForTestRegister(),
			expectedError: domain.ErrEmptyUserName,
		},
		{
			name:          "failed to register user (empty phone)",
			user:          mockEmptyPhoneUserForTestRegister(),
			expectedError: domain.ErrEmptyPhoneNumber,
		},
	}

	for _, tCase := range testCases {
		t.Run(tCase.name, func(t *testing.T) {
			_, err := underTest.RegisterUser(tCase.user)

			assertEqual[error](t, err, tCase.expectedError)
		})
	}

}

func TestLoginUser(t *testing.T) {
	underTest := userUsecase{userRepository: &mockUserRepo{}}

	t.Run("seccess login user", func(t *testing.T) {

		got, err := underTest.LoginUser("998990970138", "golang") // true
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
			password:      "golang",
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
			password:      "golang",
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

func TestGetUserDataPhoneNumber(t *testing.T) {
	underTest := userUsecase{userRepository: &mockUserRepo{}}

	t.Run("success get user by phone", func(t *testing.T) {

		got, err := underTest.GetUserDataPhoneNumber("998990970138")
		wantedUser := createUserWithPhoneNumber("998990970138")

		assertEqual[error](t, err, nil)
		assertUserEquality[*domain.User](t, got, wantedUser)

	})

	t.Run("user not found by phone", func(t *testing.T) {
		_, err := underTest.GetUserDataPhoneNumber("phone_number")
		assertEqual[error](t, err, domain.ErrUserNotFound)
	})

	t.Run("empty phone number", func(t *testing.T) {
		_, err := underTest.GetUserDataPhoneNumber("")
		assertEqual[error](t, err, domain.ErrEmptyPhoneNumber)
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
	newUser.SetName("Quvonchbek")
	newUser.SetPhoneNumber("998990970138")
	newUser.SetPassword("golang")
	newUser.SetRole("user")

	return newUser
}

func mockEmptyNameUserForTestRegister() *domain.NewUser {
	newUser := &domain.NewUser{}
	newUser.SetName("")
	newUser.SetPhoneNumber("998990970138")
	newUser.SetPassword("golang")
	newUser.SetRole("user")

	return newUser
}

func mockEmptyPhoneUserForTestRegister() *domain.NewUser {
	newUser := &domain.NewUser{}
	newUser.SetName("Quvonchbek")
	newUser.SetPhoneNumber("")
	newUser.SetPassword("golang")
	newUser.SetRole("user")

	return newUser
}

func createUserWithPhoneNumber(phoneNumber string) *domain.User {
	user := &domain.User{}
	user.SetID(1)
	user.SetName("Quvonchbek")
	user.SetPhoneNumber(phoneNumber)
	user.SetPassword("$2a$10$yaf6.78PzOT7bqsL/aDUEOI76/onUk.Y2p77KTIjQ5WNTo0g1nATG")
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
