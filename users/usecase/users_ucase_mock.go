// Code generated by moq; DO NOT EDIT.
// github.com/matryer/moq

package useCase

import (
	"github.com/go-park-mail-ru/2019_2_CoolCodeMicroServices/utils/models"
	"sync"
)

var (
	lockUsersUseCaseMockChangeUser       sync.RWMutex
	lockUsersUseCaseMockFindUsers        sync.RWMutex
	lockUsersUseCaseMockGetUserByEmail   sync.RWMutex
	lockUsersUseCaseMockGetUserByID      sync.RWMutex
	lockUsersUseCaseMockGetUserBySession sync.RWMutex
	lockUsersUseCaseMockLogin            sync.RWMutex
	lockUsersUseCaseMockSignUp           sync.RWMutex
)

// Ensure, that UsersUseCaseMock does implement UsersUseCase.
// If this is not the case, regenerate this file with moq.
var _ UsersUseCase = &UsersUseCaseMock{}

// UsersUseCaseMock is a mock implementation of UsersUseCase.
//
//     func TestSomethingThatUsesUsersUseCase(t *testing.T) {
//
//         // make and configure a mocked UsersUseCase
//         mockedUsersUseCase := &UsersUseCaseMock{
//             ChangeUserFunc: func(user *models.User) error {
// 	               panic("mock out the ChangeUser method")
//             },
//             FindUsersFunc: func(name string) (models.Users, error) {
// 	               panic("mock out the FindUsers method")
//             },
//             GetUserByEmailFunc: func(email string) (models.User, error) {
// 	               panic("mock out the GetUserByEmail method")
//             },
//             GetUserByIDFunc: func(id uint64) (models.User, error) {
// 	               panic("mock out the GetUserByID method")
//             },
//             GetUserBySessionFunc: func(session string) (uint64, error) {
// 	               panic("mock out the GetUserBySession method")
//             },
//             LoginFunc: func(user models.User) (models.User, error) {
// 	               panic("mock out the Login method")
//             },
//             SignUpFunc: func(user *models.User) error {
// 	               panic("mock out the SignUp method")
//             },
//         }
//
//         // use mockedUsersUseCase in code that requires UsersUseCase
//         // and then make assertions.
//
//     }
type UsersUseCaseMock struct {
	// ChangeUserFunc mocks the ChangeUser method.
	ChangeUserFunc func(user *models.User) error

	// FindUsersFunc mocks the FindUsers method.
	FindUsersFunc func(name string) (models.Users, error)

	// GetUserByEmailFunc mocks the GetUserByEmail method.
	GetUserByEmailFunc func(email string) (models.User, error)

	// GetUserByIDFunc mocks the GetUserByID method.
	GetUserByIDFunc func(id uint64) (models.User, error)

	// GetUserBySessionFunc mocks the GetUserBySession method.
	GetUserBySessionFunc func(session string) (uint64, error)

	// LoginFunc mocks the Login method.
	LoginFunc func(user models.User) (models.User, error)

	// SignUpFunc mocks the SignUp method.
	SignUpFunc func(user *models.User) error

	// calls tracks calls to the methods.
	calls struct {
		// ChangeUser holds details about calls to the ChangeUser method.
		ChangeUser []struct {
			// User is the user argument value.
			User *models.User
		}
		// FindUsers holds details about calls to the FindUsers method.
		FindUsers []struct {
			// Name is the name argument value.
			Name string
		}
		// GetUserByEmail holds details about calls to the GetUserByEmail method.
		GetUserByEmail []struct {
			// Email is the email argument value.
			Email string
		}
		// GetUserByID holds details about calls to the GetUserByID method.
		GetUserByID []struct {
			// ID is the id argument value.
			ID uint64
		}
		// GetUserBySession holds details about calls to the GetUserBySession method.
		GetUserBySession []struct {
			// Session is the session argument value.
			Session string
		}
		// Login holds details about calls to the Login method.
		Login []struct {
			// User is the user argument value.
			User models.User
		}
		// SignUp holds details about calls to the SignUp method.
		SignUp []struct {
			// User is the user argument value.
			User *models.User
		}
	}
}

// ChangeUser calls ChangeUserFunc.
func (mock *UsersUseCaseMock) ChangeUser(user *models.User) error {
	if mock.ChangeUserFunc == nil {
		panic("UsersUseCaseMock.ChangeUserFunc: method is nil but UsersUseCase.ChangeUser was just called")
	}
	callInfo := struct {
		User *models.User
	}{
		User: user,
	}
	lockUsersUseCaseMockChangeUser.Lock()
	mock.calls.ChangeUser = append(mock.calls.ChangeUser, callInfo)
	lockUsersUseCaseMockChangeUser.Unlock()
	return mock.ChangeUserFunc(user)
}

// ChangeUserCalls gets all the calls that were made to ChangeUser.
// Check the length with:
//     len(mockedUsersUseCase.ChangeUserCalls())
func (mock *UsersUseCaseMock) ChangeUserCalls() []struct {
	User *models.User
} {
	var calls []struct {
		User *models.User
	}
	lockUsersUseCaseMockChangeUser.RLock()
	calls = mock.calls.ChangeUser
	lockUsersUseCaseMockChangeUser.RUnlock()
	return calls
}

// FindUsers calls FindUsersFunc.
func (mock *UsersUseCaseMock) FindUsers(name string) (models.Users, error) {
	if mock.FindUsersFunc == nil {
		panic("UsersUseCaseMock.FindUsersFunc: method is nil but UsersUseCase.FindUsers was just called")
	}
	callInfo := struct {
		Name string
	}{
		Name: name,
	}
	lockUsersUseCaseMockFindUsers.Lock()
	mock.calls.FindUsers = append(mock.calls.FindUsers, callInfo)
	lockUsersUseCaseMockFindUsers.Unlock()
	return mock.FindUsersFunc(name)
}

// FindUsersCalls gets all the calls that were made to FindUsers.
// Check the length with:
//     len(mockedUsersUseCase.FindUsersCalls())
func (mock *UsersUseCaseMock) FindUsersCalls() []struct {
	Name string
} {
	var calls []struct {
		Name string
	}
	lockUsersUseCaseMockFindUsers.RLock()
	calls = mock.calls.FindUsers
	lockUsersUseCaseMockFindUsers.RUnlock()
	return calls
}

// GetUserByEmail calls GetUserByEmailFunc.
func (mock *UsersUseCaseMock) GetUserByEmail(email string) (models.User, error) {
	if mock.GetUserByEmailFunc == nil {
		panic("UsersUseCaseMock.GetUserByEmailFunc: method is nil but UsersUseCase.GetUserByEmail was just called")
	}
	callInfo := struct {
		Email string
	}{
		Email: email,
	}
	lockUsersUseCaseMockGetUserByEmail.Lock()
	mock.calls.GetUserByEmail = append(mock.calls.GetUserByEmail, callInfo)
	lockUsersUseCaseMockGetUserByEmail.Unlock()
	return mock.GetUserByEmailFunc(email)
}

// GetUserByEmailCalls gets all the calls that were made to GetUserByEmail.
// Check the length with:
//     len(mockedUsersUseCase.GetUserByEmailCalls())
func (mock *UsersUseCaseMock) GetUserByEmailCalls() []struct {
	Email string
} {
	var calls []struct {
		Email string
	}
	lockUsersUseCaseMockGetUserByEmail.RLock()
	calls = mock.calls.GetUserByEmail
	lockUsersUseCaseMockGetUserByEmail.RUnlock()
	return calls
}

// GetUserByID calls GetUserByIDFunc.
func (mock *UsersUseCaseMock) GetUserByID(id uint64) (models.User, error) {
	if mock.GetUserByIDFunc == nil {
		panic("UsersUseCaseMock.GetUserByIDFunc: method is nil but UsersUseCase.GetUserByID was just called")
	}
	callInfo := struct {
		ID uint64
	}{
		ID: id,
	}
	lockUsersUseCaseMockGetUserByID.Lock()
	mock.calls.GetUserByID = append(mock.calls.GetUserByID, callInfo)
	lockUsersUseCaseMockGetUserByID.Unlock()
	return mock.GetUserByIDFunc(id)
}

// GetUserByIDCalls gets all the calls that were made to GetUserByID.
// Check the length with:
//     len(mockedUsersUseCase.GetUserByIDCalls())
func (mock *UsersUseCaseMock) GetUserByIDCalls() []struct {
	ID uint64
} {
	var calls []struct {
		ID uint64
	}
	lockUsersUseCaseMockGetUserByID.RLock()
	calls = mock.calls.GetUserByID
	lockUsersUseCaseMockGetUserByID.RUnlock()
	return calls
}

// GetUserBySession calls GetUserBySessionFunc.
func (mock *UsersUseCaseMock) GetUserBySession(session string) (uint64, error) {
	if mock.GetUserBySessionFunc == nil {
		panic("UsersUseCaseMock.GetUserBySessionFunc: method is nil but UsersUseCase.GetUserBySession was just called")
	}
	callInfo := struct {
		Session string
	}{
		Session: session,
	}
	lockUsersUseCaseMockGetUserBySession.Lock()
	mock.calls.GetUserBySession = append(mock.calls.GetUserBySession, callInfo)
	lockUsersUseCaseMockGetUserBySession.Unlock()
	return mock.GetUserBySessionFunc(session)
}

// GetUserBySessionCalls gets all the calls that were made to GetUserBySession.
// Check the length with:
//     len(mockedUsersUseCase.GetUserBySessionCalls())
func (mock *UsersUseCaseMock) GetUserBySessionCalls() []struct {
	Session string
} {
	var calls []struct {
		Session string
	}
	lockUsersUseCaseMockGetUserBySession.RLock()
	calls = mock.calls.GetUserBySession
	lockUsersUseCaseMockGetUserBySession.RUnlock()
	return calls
}

// Login calls LoginFunc.
func (mock *UsersUseCaseMock) Login(user models.User) (models.User, error) {
	if mock.LoginFunc == nil {
		panic("UsersUseCaseMock.LoginFunc: method is nil but UsersUseCase.Login was just called")
	}
	callInfo := struct {
		User models.User
	}{
		User: user,
	}
	lockUsersUseCaseMockLogin.Lock()
	mock.calls.Login = append(mock.calls.Login, callInfo)
	lockUsersUseCaseMockLogin.Unlock()
	return mock.LoginFunc(user)
}

// LoginCalls gets all the calls that were made to Login.
// Check the length with:
//     len(mockedUsersUseCase.LoginCalls())
func (mock *UsersUseCaseMock) LoginCalls() []struct {
	User models.User
} {
	var calls []struct {
		User models.User
	}
	lockUsersUseCaseMockLogin.RLock()
	calls = mock.calls.Login
	lockUsersUseCaseMockLogin.RUnlock()
	return calls
}

// SignUp calls SignUpFunc.
func (mock *UsersUseCaseMock) SignUp(user *models.User) error {
	if mock.SignUpFunc == nil {
		panic("UsersUseCaseMock.SignUpFunc: method is nil but UsersUseCase.SignUp was just called")
	}
	callInfo := struct {
		User *models.User
	}{
		User: user,
	}
	lockUsersUseCaseMockSignUp.Lock()
	mock.calls.SignUp = append(mock.calls.SignUp, callInfo)
	lockUsersUseCaseMockSignUp.Unlock()
	return mock.SignUpFunc(user)
}

// SignUpCalls gets all the calls that were made to SignUp.
// Check the length with:
//     len(mockedUsersUseCase.SignUpCalls())
func (mock *UsersUseCaseMock) SignUpCalls() []struct {
	User *models.User
} {
	var calls []struct {
		User *models.User
	}
	lockUsersUseCaseMockSignUp.RLock()
	calls = mock.calls.SignUp
	lockUsersUseCaseMockSignUp.RUnlock()
	return calls
}
