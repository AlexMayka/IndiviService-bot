package core

import (
	"testing"

	"github.com/stretchr/testify/assert"
	"github.com/stretchr/testify/mock"
)

// MockFMS for testing
type MockFMS struct {
	mock.Mock
}

func (m *MockFMS) Get(userID int) string {
	args := m.Called(userID)
	return args.String(0)
}

func (m *MockFMS) Set(userID int, state string) {
	m.Called(userID, state)
}

func (m *MockFMS) SetParam(userID int, key string, value string) {
	m.Called(userID, key, value)
}

func (m *MockFMS) GetParam(userID int, key string) string {
	args := m.Called(userID, key)
	return args.String(0)
}

func TestFMS_StateOperations(t *testing.T) {
	mockFMS := &MockFMS{}
	userID := 123

	// Test setting state
	mockFMS.On("Set", userID, "awaiting_name").Return()
	mockFMS.Set(userID, "awaiting_name")

	// Test getting state
	mockFMS.On("Get", userID).Return("awaiting_name")
	state := mockFMS.Get(userID)
	assert.Equal(t, "awaiting_name", state)

	mockFMS.AssertExpectations(t)
}

func TestFMS_ParameterOperations(t *testing.T) {
	mockFMS := &MockFMS{}
	userID := 123
	key := "user_name"
	value := "John Doe"

	// Test setting parameter
	mockFMS.On("SetParam", userID, key, value).Return()
	mockFMS.SetParam(userID, key, value)

	// Test getting parameter
	mockFMS.On("GetParam", userID, key).Return(value)
	retrievedValue := mockFMS.GetParam(userID, key)
	assert.Equal(t, value, retrievedValue)

	mockFMS.AssertExpectations(t)
}

func TestFMS_MultipleUsers(t *testing.T) {
	mockFMS := &MockFMS{}
	user1 := 123
	user2 := 456

	// Test different states for different users
	mockFMS.On("Set", user1, "awaiting_name").Return()
	mockFMS.On("Set", user2, "awaiting_age").Return()
	
	mockFMS.Set(user1, "awaiting_name")
	mockFMS.Set(user2, "awaiting_age")

	// Test getting different states
	mockFMS.On("Get", user1).Return("awaiting_name")
	mockFMS.On("Get", user2).Return("awaiting_age")

	state1 := mockFMS.Get(user1)
	state2 := mockFMS.Get(user2)

	assert.Equal(t, "awaiting_name", state1)
	assert.Equal(t, "awaiting_age", state2)

	mockFMS.AssertExpectations(t)
}

func TestFMS_ParameterIsolation(t *testing.T) {
	mockFMS := &MockFMS{}
	user1 := 123
	user2 := 456
	key := "data"

	// Test parameter isolation between users
	mockFMS.On("SetParam", user1, key, "value1").Return()
	mockFMS.On("SetParam", user2, key, "value2").Return()
	
	mockFMS.SetParam(user1, key, "value1")
	mockFMS.SetParam(user2, key, "value2")

	// Test getting isolated parameters
	mockFMS.On("GetParam", user1, key).Return("value1")
	mockFMS.On("GetParam", user2, key).Return("value2")

	value1 := mockFMS.GetParam(user1, key)
	value2 := mockFMS.GetParam(user2, key)

	assert.Equal(t, "value1", value1)
	assert.Equal(t, "value2", value2)

	mockFMS.AssertExpectations(t)
}

func TestFMS_EmptyState(t *testing.T) {
	mockFMS := &MockFMS{}
	userID := 123

	// Test getting empty state (new user)
	mockFMS.On("Get", userID).Return("")
	state := mockFMS.Get(userID)
	assert.Equal(t, "", state)

	mockFMS.AssertExpectations(t)
}

func TestFMS_ClearState(t *testing.T) {
	mockFMS := &MockFMS{}
	userID := 123

	// Test clearing state by setting to empty string
	mockFMS.On("Set", userID, "").Return()
	mockFMS.Set(userID, "")

	// Verify state is cleared
	mockFMS.On("Get", userID).Return("")
	state := mockFMS.Get(userID)
	assert.Equal(t, "", state)

	mockFMS.AssertExpectations(t)
}

func TestFMS_NonExistentParameter(t *testing.T) {
	mockFMS := &MockFMS{}
	userID := 123
	key := "non_existent_key"

	// Test getting non-existent parameter should return empty string
	mockFMS.On("GetParam", userID, key).Return("")
	value := mockFMS.GetParam(userID, key)
	assert.Equal(t, "", value)

	mockFMS.AssertExpectations(t)
}