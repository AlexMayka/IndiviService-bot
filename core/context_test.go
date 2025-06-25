package core

import (
	"context"
	"testing"
	"telegram-sdk/types/commands"

	"github.com/stretchr/testify/assert"
	"github.com/stretchr/testify/mock"
)

// MockLogger for testing
type MockLogger struct {
	mock.Mock
}

func (m *MockLogger) Info(msg string, fields ...interface{}) {
	m.Called(msg, fields)
}

func (m *MockLogger) Error(msg string, fields ...interface{}) {
	m.Called(msg, fields)
}

func (m *MockLogger) Debug(msg string, fields ...interface{}) {
	m.Called(msg, fields)
}

func (m *MockLogger) Warn(msg string, fields ...interface{}) {
	m.Called(msg, fields)
}

func TestContext_BasicInfo(t *testing.T) {
	mockContext := &MockContext{}
	
	// Test ChatId
	mockContext.On("ChatId").Return(12345)
	chatId := mockContext.ChatId()
	assert.Equal(t, 12345, chatId)
	
	// Test Text
	mockContext.On("Text").Return("Hello, World!")
	text := mockContext.Text()
	assert.Equal(t, "Hello, World!", text)
	
	mockContext.AssertExpectations(t)
}

func TestContext_SendMessage(t *testing.T) {
	mockContext := &MockContext{}
	
	// Mock request structure
	request := &commands.SendMessageRequest{
		Text: "Test message",
	}
	
	// Test Send method
	mockContext.On("Send", request).Return(nil)
	err := mockContext.Send(request)
	assert.NoError(t, err)
	
	mockContext.AssertExpectations(t)
}

func TestContext_SendMessageError(t *testing.T) {
	mockContext := &MockContext{}
	
	request := &commands.SendMessageRequest{
		Text: "Test message",
	}
	
	// Test Send method with error
	expectedError := assert.AnError
	mockContext.On("Send", request).Return(expectedError)
	err := mockContext.Send(request)
	assert.Error(t, err)
	assert.Equal(t, expectedError, err)
	
	mockContext.AssertExpectations(t)
}

func TestContext_Logger(t *testing.T) {
	mockContext := &MockContext{}
	mockLogger := &MockLogger{}
	
	// Test Logger getter
	mockContext.On("Logger").Return(mockLogger)
	logger := mockContext.Logger()
	assert.Equal(t, mockLogger, logger)
	
	mockContext.AssertExpectations(t)
}

func TestContext_FMS(t *testing.T) {
	mockContext := &MockContext{}
	mockFMS := &MockFMS{}
	
	// Test FMS getter
	mockContext.On("FMS").Return(mockFMS)
	fms := mockContext.FMS()
	assert.Equal(t, mockFMS, fms)
	
	mockContext.AssertExpectations(t)
}

func TestContext_GoContext(t *testing.T) {
	mockContext := &MockContext{}
	goCtx := context.Background()
	
	// Test that MockContext embeds context.Context
	mockContext.Context = goCtx
	assert.Equal(t, goCtx, mockContext.Context)
}

func TestContext_IntegrationWithFMS(t *testing.T) {
	mockContext := &MockContext{}
	mockFMS := &MockFMS{}
	
	chatId := 12345
	state := "awaiting_name"
	
	// Setup context to return chat ID and FMS
	mockContext.On("ChatId").Return(chatId)
	mockContext.On("FMS").Return(mockFMS)
	
	// Setup FMS operations
	mockFMS.On("Set", chatId, state).Return()
	mockFMS.On("Get", chatId).Return(state)
	
	// Simulate usage
	fms := mockContext.FMS()
	fms.Set(mockContext.ChatId(), state)
	retrievedState := fms.Get(mockContext.ChatId())
	
	assert.Equal(t, state, retrievedState)
	
	mockContext.AssertExpectations(t)
	mockFMS.AssertExpectations(t)
}

func TestContext_LoggerIntegration(t *testing.T) {
	mockContext := &MockContext{}
	mockLogger := &MockLogger{}
	
	// Setup context to return logger
	mockContext.On("Logger").Return(mockLogger)
	mockContext.On("Text").Return("test message")
	
	// Setup logger expectations
	mockLogger.On("Info", "Received message", mock.Anything).Return()
	
	// Simulate usage
	logger := mockContext.Logger()
	logger.Info("Received message", "text", mockContext.Text())
	
	mockContext.AssertExpectations(t)
	mockLogger.AssertExpectations(t)
}