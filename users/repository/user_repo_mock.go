// Code generated by moq; DO NOT EDIT.
// github.com/matryer/moq

package repository

import (
	"github.com/go-park-mail-ru/2019_2_CoolCodeMicroServices/utils/models"
	"sync"
)

var (
	lockUserRepoMockContains       sync.RWMutex
	lockUserRepoMockGetUserByEmail sync.RWMutex
	lockUserRepoMockGetUserByID    sync.RWMutex
	lockUserRepoMockGetUsers       sync.RWMutex
	lockUserRepoMockPutUser        sync.RWMutex
	lockUserRepoMockReplace        sync.RWMutex
)

// Ensure, that UserRepoMock does implement UserRepo.
// If this is not the case, regenerate this file with moq.
var _ UserRepo = &UserRepoMock{}

// UserRepoMock is a mock implementation of UserRepo.
//
//     func TestSomethingThatUsesUserRepo(t *testing.T) {
//
//         // make and configure a mocked UserRepo
//         mockedUserRepo := &UserRepoMock{
//             ContainsFunc: func(user models.User) bool {
// 	               panic("mock out the Contains method")
//             },
//             GetUserByEmailFunc: func(email string) (models.User, error) {
// 	               panic("mock out the GetUserByEmail method")
//             },
//             GetUserByIDFunc: func(ID uint64) (models.User, error) {
// 	               panic("mock out the GetUserByID method")
//             },
//             GetUsersFunc: func() (models.Users, error) {
// 	               panic("mock out the GetUsers method")
//             },
//             PutUserFunc: func(newUser *models.User) (uint64, error) {
// 	               panic("mock out the PutUser method")
//             },
//             ReplaceFunc: func(ID uint64, newUser *models.User) error {
// 	               panic("mock out the Replace method")
//             },
//         }
//
//         // use mockedUserRepo in code that requires UserRepo
//         // and then make assertions.
//
//     }
type UserRepoMock struct {
	// ContainsFunc mocks the Contains method.
	ContainsFunc func(user models.User) bool

	// GetUserByEmailFunc mocks the GetUserByEmail method.
	GetUserByEmailFunc func(email string) (models.User, error)

	// GetUserByIDFunc mocks the GetUserByID method.
	GetUserByIDFunc func(ID uint64) (models.User, error)

	// GetUsersFunc mocks the GetUsers method.
	GetUsersFunc func() (models.Users, error)

	// PutUserFunc mocks the PutUser method.
	PutUserFunc func(newUser *models.User) (uint64, error)

	// ReplaceFunc mocks the Replace method.
	ReplaceFunc func(ID uint64, newUser *models.User) error

	// calls tracks calls to the methods.
	calls struct {
		// Contains holds details about calls to the Contains method.
		Contains []struct {
			// User is the user argument value.
			User models.User
		}
		// GetUserByEmail holds details about calls to the GetUserByEmail method.
		GetUserByEmail []struct {
			// Email is the email argument value.
			Email string
		}
		// GetUserByID holds details about calls to the GetUserByID method.
		GetUserByID []struct {
			// ID is the ID argument value.
			ID uint64
		}
		// GetUsers holds details about calls to the GetUsers method.
		GetUsers []struct {
		}
		// PutUser holds details about calls to the PutUser method.
		PutUser []struct {
			// NewUser is the newUser argument value.
			NewUser *models.User
		}
		// Replace holds details about calls to the Replace method.
		Replace []struct {
			// ID is the ID argument value.
			ID uint64
			// NewUser is the newUser argument value.
			NewUser *models.User
		}
	}
}

// Contains calls ContainsFunc.
func (mock *UserRepoMock) Contains(user models.User) bool {
	if mock.ContainsFunc == nil {
		panic("UserRepoMock.ContainsFunc: method is nil but UserRepo.Contains was just called")
	}
	callInfo := struct {
		User models.User
	}{
		User: user,
	}
	lockUserRepoMockContains.Lock()
	mock.calls.Contains = append(mock.calls.Contains, callInfo)
	lockUserRepoMockContains.Unlock()
	return mock.ContainsFunc(user)
}

// ContainsCalls gets all the calls that were made to Contains.
// Check the length with:
//     len(mockedUserRepo.ContainsCalls())
func (mock *UserRepoMock) ContainsCalls() []struct {
	User models.User
} {
	var calls []struct {
		User models.User
	}
	lockUserRepoMockContains.RLock()
	calls = mock.calls.Contains
	lockUserRepoMockContains.RUnlock()
	return calls
}

// GetUserByEmail calls GetUserByEmailFunc.
func (mock *UserRepoMock) GetUserByEmail(email string) (models.User, error) {
	if mock.GetUserByEmailFunc == nil {
		panic("UserRepoMock.GetUserByEmailFunc: method is nil but UserRepo.GetUserByEmail was just called")
	}
	callInfo := struct {
		Email string
	}{
		Email: email,
	}
	lockUserRepoMockGetUserByEmail.Lock()
	mock.calls.GetUserByEmail = append(mock.calls.GetUserByEmail, callInfo)
	lockUserRepoMockGetUserByEmail.Unlock()
	return mock.GetUserByEmailFunc(email)
}

// GetUserByEmailCalls gets all the calls that were made to GetUserByEmail.
// Check the length with:
//     len(mockedUserRepo.GetUserByEmailCalls())
func (mock *UserRepoMock) GetUserByEmailCalls() []struct {
	Email string
} {
	var calls []struct {
		Email string
	}
	lockUserRepoMockGetUserByEmail.RLock()
	calls = mock.calls.GetUserByEmail
	lockUserRepoMockGetUserByEmail.RUnlock()
	return calls
}

// GetUserByID calls GetUserByIDFunc.
func (mock *UserRepoMock) GetUserByID(ID uint64) (models.User, error) {
	if mock.GetUserByIDFunc == nil {
		panic("UserRepoMock.GetUserByIDFunc: method is nil but UserRepo.GetUserByID was just called")
	}
	callInfo := struct {
		ID uint64
	}{
		ID: ID,
	}
	lockUserRepoMockGetUserByID.Lock()
	mock.calls.GetUserByID = append(mock.calls.GetUserByID, callInfo)
	lockUserRepoMockGetUserByID.Unlock()
	return mock.GetUserByIDFunc(ID)
}

// GetUserByIDCalls gets all the calls that were made to GetUserByID.
// Check the length with:
//     len(mockedUserRepo.GetUserByIDCalls())
func (mock *UserRepoMock) GetUserByIDCalls() []struct {
	ID uint64
} {
	var calls []struct {
		ID uint64
	}
	lockUserRepoMockGetUserByID.RLock()
	calls = mock.calls.GetUserByID
	lockUserRepoMockGetUserByID.RUnlock()
	return calls
}

// GetUsers calls GetUsersFunc.
func (mock *UserRepoMock) GetUsers() (models.Users, error) {
	if mock.GetUsersFunc == nil {
		panic("UserRepoMock.GetUsersFunc: method is nil but UserRepo.GetUsers was just called")
	}
	callInfo := struct {
	}{}
	lockUserRepoMockGetUsers.Lock()
	mock.calls.GetUsers = append(mock.calls.GetUsers, callInfo)
	lockUserRepoMockGetUsers.Unlock()
	return mock.GetUsersFunc()
}

// GetUsersCalls gets all the calls that were made to GetUsers.
// Check the length with:
//     len(mockedUserRepo.GetUsersCalls())
func (mock *UserRepoMock) GetUsersCalls() []struct {
} {
	var calls []struct {
	}
	lockUserRepoMockGetUsers.RLock()
	calls = mock.calls.GetUsers
	lockUserRepoMockGetUsers.RUnlock()
	return calls
}

// PutUser calls PutUserFunc.
func (mock *UserRepoMock) PutUser(newUser *models.User) (uint64, error) {
	if mock.PutUserFunc == nil {
		panic("UserRepoMock.PutUserFunc: method is nil but UserRepo.PutUser was just called")
	}
	callInfo := struct {
		NewUser *models.User
	}{
		NewUser: newUser,
	}
	lockUserRepoMockPutUser.Lock()
	mock.calls.PutUser = append(mock.calls.PutUser, callInfo)
	lockUserRepoMockPutUser.Unlock()
	return mock.PutUserFunc(newUser)
}

// PutUserCalls gets all the calls that were made to PutUser.
// Check the length with:
//     len(mockedUserRepo.PutUserCalls())
func (mock *UserRepoMock) PutUserCalls() []struct {
	NewUser *models.User
} {
	var calls []struct {
		NewUser *models.User
	}
	lockUserRepoMockPutUser.RLock()
	calls = mock.calls.PutUser
	lockUserRepoMockPutUser.RUnlock()
	return calls
}

// Replace calls ReplaceFunc.
func (mock *UserRepoMock) Replace(ID uint64, newUser *models.User) error {
	if mock.ReplaceFunc == nil {
		panic("UserRepoMock.ReplaceFunc: method is nil but UserRepo.Replace was just called")
	}
	callInfo := struct {
		ID      uint64
		NewUser *models.User
	}{
		ID:      ID,
		NewUser: newUser,
	}
	lockUserRepoMockReplace.Lock()
	mock.calls.Replace = append(mock.calls.Replace, callInfo)
	lockUserRepoMockReplace.Unlock()
	return mock.ReplaceFunc(ID, newUser)
}

// ReplaceCalls gets all the calls that were made to Replace.
// Check the length with:
//     len(mockedUserRepo.ReplaceCalls())
func (mock *UserRepoMock) ReplaceCalls() []struct {
	ID      uint64
	NewUser *models.User
} {
	var calls []struct {
		ID      uint64
		NewUser *models.User
	}
	lockUserRepoMockReplace.RLock()
	calls = mock.calls.Replace
	lockUserRepoMockReplace.RUnlock()
	return calls
}
