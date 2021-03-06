// Code generated by moq; DO NOT EDIT.
// github.com/matryer/moq

package sendgrid

import (
	"sync"
)

var (
	lockAPIClientMockCreateAPIKeyForSubUser sync.RWMutex
	lockAPIClientMockCreateSubUser          sync.RWMutex
	lockAPIClientMockDeleteSubUser          sync.RWMutex
	lockAPIClientMockGetAPIKeysForSubUser   sync.RWMutex
	lockAPIClientMockGetSubUserByUsername   sync.RWMutex
	lockAPIClientMockListIPAddresses        sync.RWMutex
	lockAPIClientMockListSubUsers           sync.RWMutex
)

// Ensure, that APIClientMock does implement APIClient.
// If this is not the case, regenerate this file with moq.
var _ APIClient = &APIClientMock{}

// APIClientMock is a mock implementation of APIClient.
//
//     func TestSomethingThatUsesAPIClient(t *testing.T) {
//
//         // make and configure a mocked APIClient
//         mockedAPIClient := &APIClientMock{
//             CreateAPIKeyForSubUserFunc: func(username string, scopes []string) (*APIKey, error) {
// 	               panic("mock out the CreateAPIKeyForSubUser method")
//             },
//             CreateSubUserFunc: func(id string, email string, password string, ips []string) (*SubUser, error) {
// 	               panic("mock out the CreateSubUser method")
//             },
//             DeleteSubUserFunc: func(username string) error {
// 	               panic("mock out the DeleteSubUser method")
//             },
//             GetAPIKeysForSubUserFunc: func(username string) ([]*APIKey, error) {
// 	               panic("mock out the GetAPIKeysForSubUser method")
//             },
//             GetSubUserByUsernameFunc: func(username string) (*SubUser, error) {
// 	               panic("mock out the GetSubUserByUsername method")
//             },
//             ListIPAddressesFunc: func() ([]*IPAddress, error) {
// 	               panic("mock out the ListIPAddresses method")
//             },
//             ListSubUsersFunc: func(query map[string]string) ([]*SubUser, error) {
// 	               panic("mock out the ListSubUsers method")
//             },
//         }
//
//         // use mockedAPIClient in code that requires APIClient
//         // and then make assertions.
//
//     }
type APIClientMock struct {
	// CreateAPIKeyForSubUserFunc mocks the CreateAPIKeyForSubUser method.
	CreateAPIKeyForSubUserFunc func(username string, scopes []string) (*APIKey, error)

	// CreateSubUserFunc mocks the CreateSubUser method.
	CreateSubUserFunc func(id string, email string, password string, ips []string) (*SubUser, error)

	// DeleteSubUserFunc mocks the DeleteSubUser method.
	DeleteSubUserFunc func(username string) error

	// GetAPIKeysForSubUserFunc mocks the GetAPIKeysForSubUser method.
	GetAPIKeysForSubUserFunc func(username string) ([]*APIKey, error)

	// GetSubUserByUsernameFunc mocks the GetSubUserByUsername method.
	GetSubUserByUsernameFunc func(username string) (*SubUser, error)

	// ListIPAddressesFunc mocks the ListIPAddresses method.
	ListIPAddressesFunc func() ([]*IPAddress, error)

	// ListSubUsersFunc mocks the ListSubUsers method.
	ListSubUsersFunc func(query map[string]string) ([]*SubUser, error)

	// calls tracks calls to the methods.
	calls struct {
		// CreateAPIKeyForSubUser holds details about calls to the CreateAPIKeyForSubUser method.
		CreateAPIKeyForSubUser []struct {
			// Username is the username argument value.
			Username string
			// Scopes is the scopes argument value.
			Scopes []string
		}
		// CreateSubUser holds details about calls to the CreateSubUser method.
		CreateSubUser []struct {
			// ID is the id argument value.
			ID string
			// Email is the email argument value.
			Email string
			// Password is the password argument value.
			Password string
			// Ips is the ips argument value.
			Ips []string
		}
		// DeleteSubUser holds details about calls to the DeleteSubUser method.
		DeleteSubUser []struct {
			// Username is the username argument value.
			Username string
		}
		// GetAPIKeysForSubUser holds details about calls to the GetAPIKeysForSubUser method.
		GetAPIKeysForSubUser []struct {
			// Username is the username argument value.
			Username string
		}
		// GetSubUserByUsername holds details about calls to the GetSubUserByUsername method.
		GetSubUserByUsername []struct {
			// Username is the username argument value.
			Username string
		}
		// ListIPAddresses holds details about calls to the ListIPAddresses method.
		ListIPAddresses []struct {
		}
		// ListSubUsers holds details about calls to the ListSubUsers method.
		ListSubUsers []struct {
			// Query is the query argument value.
			Query map[string]string
		}
	}
}

// CreateAPIKeyForSubUser calls CreateAPIKeyForSubUserFunc.
func (mock *APIClientMock) CreateAPIKeyForSubUser(username string, scopes []string) (*APIKey, error) {
	if mock.CreateAPIKeyForSubUserFunc == nil {
		panic("APIClientMock.CreateAPIKeyForSubUserFunc: method is nil but APIClient.CreateAPIKeyForSubUser was just called")
	}
	callInfo := struct {
		Username string
		Scopes   []string
	}{
		Username: username,
		Scopes:   scopes,
	}
	lockAPIClientMockCreateAPIKeyForSubUser.Lock()
	mock.calls.CreateAPIKeyForSubUser = append(mock.calls.CreateAPIKeyForSubUser, callInfo)
	lockAPIClientMockCreateAPIKeyForSubUser.Unlock()
	return mock.CreateAPIKeyForSubUserFunc(username, scopes)
}

// CreateAPIKeyForSubUserCalls gets all the calls that were made to CreateAPIKeyForSubUser.
// Check the length with:
//     len(mockedAPIClient.CreateAPIKeyForSubUserCalls())
func (mock *APIClientMock) CreateAPIKeyForSubUserCalls() []struct {
	Username string
	Scopes   []string
} {
	var calls []struct {
		Username string
		Scopes   []string
	}
	lockAPIClientMockCreateAPIKeyForSubUser.RLock()
	calls = mock.calls.CreateAPIKeyForSubUser
	lockAPIClientMockCreateAPIKeyForSubUser.RUnlock()
	return calls
}

// CreateSubUser calls CreateSubUserFunc.
func (mock *APIClientMock) CreateSubUser(id string, email string, password string, ips []string) (*SubUser, error) {
	if mock.CreateSubUserFunc == nil {
		panic("APIClientMock.CreateSubUserFunc: method is nil but APIClient.CreateSubUser was just called")
	}
	callInfo := struct {
		ID       string
		Email    string
		Password string
		Ips      []string
	}{
		ID:       id,
		Email:    email,
		Password: password,
		Ips:      ips,
	}
	lockAPIClientMockCreateSubUser.Lock()
	mock.calls.CreateSubUser = append(mock.calls.CreateSubUser, callInfo)
	lockAPIClientMockCreateSubUser.Unlock()
	return mock.CreateSubUserFunc(id, email, password, ips)
}

// CreateSubUserCalls gets all the calls that were made to CreateSubUser.
// Check the length with:
//     len(mockedAPIClient.CreateSubUserCalls())
func (mock *APIClientMock) CreateSubUserCalls() []struct {
	ID       string
	Email    string
	Password string
	Ips      []string
} {
	var calls []struct {
		ID       string
		Email    string
		Password string
		Ips      []string
	}
	lockAPIClientMockCreateSubUser.RLock()
	calls = mock.calls.CreateSubUser
	lockAPIClientMockCreateSubUser.RUnlock()
	return calls
}

// DeleteSubUser calls DeleteSubUserFunc.
func (mock *APIClientMock) DeleteSubUser(username string) error {
	if mock.DeleteSubUserFunc == nil {
		panic("APIClientMock.DeleteSubUserFunc: method is nil but APIClient.DeleteSubUser was just called")
	}
	callInfo := struct {
		Username string
	}{
		Username: username,
	}
	lockAPIClientMockDeleteSubUser.Lock()
	mock.calls.DeleteSubUser = append(mock.calls.DeleteSubUser, callInfo)
	lockAPIClientMockDeleteSubUser.Unlock()
	return mock.DeleteSubUserFunc(username)
}

// DeleteSubUserCalls gets all the calls that were made to DeleteSubUser.
// Check the length with:
//     len(mockedAPIClient.DeleteSubUserCalls())
func (mock *APIClientMock) DeleteSubUserCalls() []struct {
	Username string
} {
	var calls []struct {
		Username string
	}
	lockAPIClientMockDeleteSubUser.RLock()
	calls = mock.calls.DeleteSubUser
	lockAPIClientMockDeleteSubUser.RUnlock()
	return calls
}

// GetAPIKeysForSubUser calls GetAPIKeysForSubUserFunc.
func (mock *APIClientMock) GetAPIKeysForSubUser(username string) ([]*APIKey, error) {
	if mock.GetAPIKeysForSubUserFunc == nil {
		panic("APIClientMock.GetAPIKeysForSubUserFunc: method is nil but APIClient.GetAPIKeysForSubUser was just called")
	}
	callInfo := struct {
		Username string
	}{
		Username: username,
	}
	lockAPIClientMockGetAPIKeysForSubUser.Lock()
	mock.calls.GetAPIKeysForSubUser = append(mock.calls.GetAPIKeysForSubUser, callInfo)
	lockAPIClientMockGetAPIKeysForSubUser.Unlock()
	return mock.GetAPIKeysForSubUserFunc(username)
}

// GetAPIKeysForSubUserCalls gets all the calls that were made to GetAPIKeysForSubUser.
// Check the length with:
//     len(mockedAPIClient.GetAPIKeysForSubUserCalls())
func (mock *APIClientMock) GetAPIKeysForSubUserCalls() []struct {
	Username string
} {
	var calls []struct {
		Username string
	}
	lockAPIClientMockGetAPIKeysForSubUser.RLock()
	calls = mock.calls.GetAPIKeysForSubUser
	lockAPIClientMockGetAPIKeysForSubUser.RUnlock()
	return calls
}

// GetSubUserByUsername calls GetSubUserByUsernameFunc.
func (mock *APIClientMock) GetSubUserByUsername(username string) (*SubUser, error) {
	if mock.GetSubUserByUsernameFunc == nil {
		panic("APIClientMock.GetSubUserByUsernameFunc: method is nil but APIClient.GetSubUserByUsername was just called")
	}
	callInfo := struct {
		Username string
	}{
		Username: username,
	}
	lockAPIClientMockGetSubUserByUsername.Lock()
	mock.calls.GetSubUserByUsername = append(mock.calls.GetSubUserByUsername, callInfo)
	lockAPIClientMockGetSubUserByUsername.Unlock()
	return mock.GetSubUserByUsernameFunc(username)
}

// GetSubUserByUsernameCalls gets all the calls that were made to GetSubUserByUsername.
// Check the length with:
//     len(mockedAPIClient.GetSubUserByUsernameCalls())
func (mock *APIClientMock) GetSubUserByUsernameCalls() []struct {
	Username string
} {
	var calls []struct {
		Username string
	}
	lockAPIClientMockGetSubUserByUsername.RLock()
	calls = mock.calls.GetSubUserByUsername
	lockAPIClientMockGetSubUserByUsername.RUnlock()
	return calls
}

// ListIPAddresses calls ListIPAddressesFunc.
func (mock *APIClientMock) ListIPAddresses() ([]*IPAddress, error) {
	if mock.ListIPAddressesFunc == nil {
		panic("APIClientMock.ListIPAddressesFunc: method is nil but APIClient.ListIPAddresses was just called")
	}
	callInfo := struct {
	}{}
	lockAPIClientMockListIPAddresses.Lock()
	mock.calls.ListIPAddresses = append(mock.calls.ListIPAddresses, callInfo)
	lockAPIClientMockListIPAddresses.Unlock()
	return mock.ListIPAddressesFunc()
}

// ListIPAddressesCalls gets all the calls that were made to ListIPAddresses.
// Check the length with:
//     len(mockedAPIClient.ListIPAddressesCalls())
func (mock *APIClientMock) ListIPAddressesCalls() []struct {
} {
	var calls []struct {
	}
	lockAPIClientMockListIPAddresses.RLock()
	calls = mock.calls.ListIPAddresses
	lockAPIClientMockListIPAddresses.RUnlock()
	return calls
}

// ListSubUsers calls ListSubUsersFunc.
func (mock *APIClientMock) ListSubUsers(query map[string]string) ([]*SubUser, error) {
	if mock.ListSubUsersFunc == nil {
		panic("APIClientMock.ListSubUsersFunc: method is nil but APIClient.ListSubUsers was just called")
	}
	callInfo := struct {
		Query map[string]string
	}{
		Query: query,
	}
	lockAPIClientMockListSubUsers.Lock()
	mock.calls.ListSubUsers = append(mock.calls.ListSubUsers, callInfo)
	lockAPIClientMockListSubUsers.Unlock()
	return mock.ListSubUsersFunc(query)
}

// ListSubUsersCalls gets all the calls that were made to ListSubUsers.
// Check the length with:
//     len(mockedAPIClient.ListSubUsersCalls())
func (mock *APIClientMock) ListSubUsersCalls() []struct {
	Query map[string]string
} {
	var calls []struct {
		Query map[string]string
	}
	lockAPIClientMockListSubUsers.RLock()
	calls = mock.calls.ListSubUsers
	lockAPIClientMockListSubUsers.RUnlock()
	return calls
}
