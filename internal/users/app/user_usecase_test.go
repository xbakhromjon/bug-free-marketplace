package app

import (
	"errors"
	"golang-project-template/internal/users/domain"
	"golang.org/x/crypto/bcrypt"
	"log"
	"reflect"
	"testing"
	"time"
)

func TestRegisterUser(t *testing.T) {
	underTest := userUsecase{userRepository: &mockUserRepo{}}

	t.Run("success register user", func(t *testing.T) {
		u := mockUserForTestRegister()
		got, _ := underTest.RegisterUser(u)
		want := 1
		registerAssertError(t, got, want)
	})

	t.Run("failed to register user", func(t *testing.T) {
		u := mockUserForTestRegister()
		underTest.userRepository.(*mockUserRepo).err = errors.New("registration failed")

		got, err := underTest.RegisterUser(u)
		want := 0
		registerAssertError(t, got, want)

		if err == nil || err.Error() != "registration failed" {
			t.Errorf("want error 'registration failed' but got: %v", err)
		}
	})
}

func TestLoginUser(t *testing.T) {
	underTest := userUsecase{userRepository: &mockUserRepo{}}

	t.Run("seccess login user", func(t *testing.T) {

		mockUser := createMockUserWithPassword("golang")
		underTest.userRepository.(*mockUserRepo).user = mockUser

		got, err := underTest.LoginUser("998990970138", "golang") // true
		want := true
		assertError(t, got, want)
		if err != nil {
			t.Errorf("unexpected error: %v", err)
		}
	})

	t.Run("failed to register user", func(t *testing.T) {
		underTest.userRepository.(*mockUserRepo).err = errors.New("login failed")

		got, err := underTest.LoginUser("", "golang") // false
		want := false
		assertError(t, got, want)

		if err == nil || err.Error() != "login failed" {
			t.Errorf("unexpected error: %v", err)
		}
	})
}

func TestGetUserDataPhoneNumber(t *testing.T) {
	underTest := userUsecase{userRepository: &mockUserRepo{}}

	t.Run("success get user by phone", func(t *testing.T) {

		wantedUser := mockUserForGetUserDataPhoneNumber()
		underTest.userRepository.(*mockUserRepo).user = wantedUser

		got, err := underTest.GetUserDataPhoneNumber("998990970138")

		assertNoError(t, err)
		assertUserEquality(t, got, wantedUser)
	})

	t.Run("user not found by phone", func(t *testing.T) {

		underTest.userRepository.(*mockUserRepo).err = errors.New("user not found")
		got, err := underTest.GetUserDataPhoneNumber("phone")

		assertNOError(t, err, "user not found")
		assertUserNil(t, got)
	})
}

//---------------------------------------//

func assertUserNil(t *testing.T, user *domain.User) {
	t.Helper()
	if user != nil {
		t.Errorf("expected user to be nil, but got: %+v", user)
	}
}

func assertUserEquality(t *testing.T, got, wanted *domain.User) {
	t.Helper()
	if !reflect.DeepEqual(got, wanted) {
		t.Errorf("wanted: %+v but got: %+v", wanted, got)
	}

}

func assertNoError(t *testing.T, err error) {
	t.Helper()
	if err != nil {
		t.Errorf("unexpected error: %v", err)
	}
}

func assertNOError(t testing.TB, err error, expectedMessage string) {
	t.Helper()
	if err == nil || err.Error() != expectedMessage {
		t.Errorf("want error '%s' but got: %v", expectedMessage, err)
	}
}

func assertError(t testing.TB, got, want bool) {
	t.Helper()
	if got != want {
		t.Errorf("want %+v but got %+v", want, got)
	}
}

func registerAssertError(t testing.TB, got, want int) {
	t.Helper()
	if got != want {
		t.Errorf("want %+v but got %+v", want, got)
	}
}

//--------------------------//

func mockUserForTestRegister() *domain.NewUser {
	newUser := &domain.NewUser{}
	newUser.SetName("Quvonchbek")
	newUser.SetPhoneNumber("998990970138")
	newUser.SetPassword("user")
	newUser.SetRole("user")

	return newUser
}

func mockUserForGetUserDataPhoneNumber() *domain.User {
	user := &domain.User{}
	user.SetID(1)
	user.SetName("Quvonchbek")
	user.SetPhoneNumber("998990970138")
	user.SetPassword("golang")
	user.SetRole("user")
	user.SetCreateAt(time.Time{})
	user.SetUpdatedAt(time.Time{})
	user.SetDeletedAt(nil)

	return user
}

func createMockUserWithPassword(password string) *domain.User {

	hashedPassword, err := bcrypt.GenerateFromPassword([]byte("golang"), bcrypt.DefaultCost)
	if err != nil {
		log.Fatalf("error generating hashed password: %v", err)
	}

	mockUser := &domain.User{}
	mockUser.SetPassword(string(hashedPassword))
	return mockUser
}

//------------------------------------//

type mockUserRepo struct {
	err  error
	user *domain.User
}

func newMockUserRepo() domain.UserRepository {
	return &mockUserRepo{}
}

func (m *mockUserRepo) Save(user *domain.User) (int, error) {
	if m.err != nil {
		return 0, m.err
	}

	return 1, nil
}

func (m *mockUserRepo) FindOneByPhoneNumber(phoneNumber string) (*domain.User, error) {
	return m.user, m.err
}
