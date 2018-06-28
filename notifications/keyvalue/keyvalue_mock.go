// Code generated by moq; DO NOT EDIT
// github.com/matryer/moq

package keyvalue

import (
	"sync"
)

var (
	lockKeyValueMockDelete    sync.RWMutex
	lockKeyValueMockGetInt64  sync.RWMutex
	lockKeyValueMockGetString sync.RWMutex
	lockKeyValueMockPutInt64  sync.RWMutex
	lockKeyValueMockPutString sync.RWMutex
)

// KeyValueMock is a mock implementation of KeyValue.
//
//     func TestSomethingThatUsesKeyValue(t *testing.T) {
//
//         // make and configure a mocked KeyValue
//         mockedKeyValue := &KeyValueMock{
//             DeleteFunc: func(key string) error {
// 	               panic("TODO: mock out the Delete method")
//             },
//             GetInt64Func: func(key string) (int64, error) {
// 	               panic("TODO: mock out the GetInt64 method")
//             },
//             GetStringFunc: func(key string) (string, error) {
// 	               panic("TODO: mock out the GetString method")
//             },
//             PutInt64Func: func(key string, value int64) error {
// 	               panic("TODO: mock out the PutInt64 method")
//             },
//             PutStringFunc: func(key string, value string) error {
// 	               panic("TODO: mock out the PutString method")
//             },
//         }
//
//         // TODO: use mockedKeyValue in code that requires KeyValue
//         //       and then make assertions.
//
//     }
type KeyValueMock struct {
	// DeleteFunc mocks the Delete method.
	DeleteFunc func(key string) error

	// GetInt64Func mocks the GetInt64 method.
	GetInt64Func func(key string) (int64, error)

	// GetStringFunc mocks the GetString method.
	GetStringFunc func(key string) (string, error)

	// PutInt64Func mocks the PutInt64 method.
	PutInt64Func func(key string, value int64) error

	// PutStringFunc mocks the PutString method.
	PutStringFunc func(key string, value string) error

	// calls tracks calls to the methods.
	calls struct {
		// Delete holds details about calls to the Delete method.
		Delete []struct {
			// Key is the key argument value.
			Key string
		}
		// GetInt64 holds details about calls to the GetInt64 method.
		GetInt64 []struct {
			// Key is the key argument value.
			Key string
		}
		// GetString holds details about calls to the GetString method.
		GetString []struct {
			// Key is the key argument value.
			Key string
		}
		// PutInt64 holds details about calls to the PutInt64 method.
		PutInt64 []struct {
			// Key is the key argument value.
			Key string
			// Value is the value argument value.
			Value int64
		}
		// PutString holds details about calls to the PutString method.
		PutString []struct {
			// Key is the key argument value.
			Key string
			// Value is the value argument value.
			Value string
		}
	}
}

// Delete calls DeleteFunc.
func (mock *KeyValueMock) Delete(key string) error {
	if mock.DeleteFunc == nil {
		panic("moq: KeyValueMock.DeleteFunc is nil but KeyValue.Delete was just called")
	}
	callInfo := struct {
		Key string
	}{
		Key: key,
	}
	lockKeyValueMockDelete.Lock()
	mock.calls.Delete = append(mock.calls.Delete, callInfo)
	lockKeyValueMockDelete.Unlock()
	return mock.DeleteFunc(key)
}

// DeleteCalls gets all the calls that were made to Delete.
// Check the length with:
//     len(mockedKeyValue.DeleteCalls())
func (mock *KeyValueMock) DeleteCalls() []struct {
	Key string
} {
	var calls []struct {
		Key string
	}
	lockKeyValueMockDelete.RLock()
	calls = mock.calls.Delete
	lockKeyValueMockDelete.RUnlock()
	return calls
}

// GetInt64 calls GetInt64Func.
func (mock *KeyValueMock) GetInt64(key string) (int64, error) {
	if mock.GetInt64Func == nil {
		panic("moq: KeyValueMock.GetInt64Func is nil but KeyValue.GetInt64 was just called")
	}
	callInfo := struct {
		Key string
	}{
		Key: key,
	}
	lockKeyValueMockGetInt64.Lock()
	mock.calls.GetInt64 = append(mock.calls.GetInt64, callInfo)
	lockKeyValueMockGetInt64.Unlock()
	return mock.GetInt64Func(key)
}

// GetInt64Calls gets all the calls that were made to GetInt64.
// Check the length with:
//     len(mockedKeyValue.GetInt64Calls())
func (mock *KeyValueMock) GetInt64Calls() []struct {
	Key string
} {
	var calls []struct {
		Key string
	}
	lockKeyValueMockGetInt64.RLock()
	calls = mock.calls.GetInt64
	lockKeyValueMockGetInt64.RUnlock()
	return calls
}

// GetString calls GetStringFunc.
func (mock *KeyValueMock) GetString(key string) (string, error) {
	if mock.GetStringFunc == nil {
		panic("moq: KeyValueMock.GetStringFunc is nil but KeyValue.GetString was just called")
	}
	callInfo := struct {
		Key string
	}{
		Key: key,
	}
	lockKeyValueMockGetString.Lock()
	mock.calls.GetString = append(mock.calls.GetString, callInfo)
	lockKeyValueMockGetString.Unlock()
	return mock.GetStringFunc(key)
}

// GetStringCalls gets all the calls that were made to GetString.
// Check the length with:
//     len(mockedKeyValue.GetStringCalls())
func (mock *KeyValueMock) GetStringCalls() []struct {
	Key string
} {
	var calls []struct {
		Key string
	}
	lockKeyValueMockGetString.RLock()
	calls = mock.calls.GetString
	lockKeyValueMockGetString.RUnlock()
	return calls
}

// PutInt64 calls PutInt64Func.
func (mock *KeyValueMock) PutInt64(key string, value int64) error {
	if mock.PutInt64Func == nil {
		panic("moq: KeyValueMock.PutInt64Func is nil but KeyValue.PutInt64 was just called")
	}
	callInfo := struct {
		Key   string
		Value int64
	}{
		Key:   key,
		Value: value,
	}
	lockKeyValueMockPutInt64.Lock()
	mock.calls.PutInt64 = append(mock.calls.PutInt64, callInfo)
	lockKeyValueMockPutInt64.Unlock()
	return mock.PutInt64Func(key, value)
}

// PutInt64Calls gets all the calls that were made to PutInt64.
// Check the length with:
//     len(mockedKeyValue.PutInt64Calls())
func (mock *KeyValueMock) PutInt64Calls() []struct {
	Key   string
	Value int64
} {
	var calls []struct {
		Key   string
		Value int64
	}
	lockKeyValueMockPutInt64.RLock()
	calls = mock.calls.PutInt64
	lockKeyValueMockPutInt64.RUnlock()
	return calls
}

// PutString calls PutStringFunc.
func (mock *KeyValueMock) PutString(key string, value string) error {
	if mock.PutStringFunc == nil {
		panic("moq: KeyValueMock.PutStringFunc is nil but KeyValue.PutString was just called")
	}
	callInfo := struct {
		Key   string
		Value string
	}{
		Key:   key,
		Value: value,
	}
	lockKeyValueMockPutString.Lock()
	mock.calls.PutString = append(mock.calls.PutString, callInfo)
	lockKeyValueMockPutString.Unlock()
	return mock.PutStringFunc(key, value)
}

// PutStringCalls gets all the calls that were made to PutString.
// Check the length with:
//     len(mockedKeyValue.PutStringCalls())
func (mock *KeyValueMock) PutStringCalls() []struct {
	Key   string
	Value string
} {
	var calls []struct {
		Key   string
		Value string
	}
	lockKeyValueMockPutString.RLock()
	calls = mock.calls.PutString
	lockKeyValueMockPutString.RUnlock()
	return calls
}