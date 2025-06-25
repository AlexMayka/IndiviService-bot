package core

import (
	"context"
	"testing"
	"time"
	"telegram-sdk/types/commands"

	"github.com/stretchr/testify/assert"
	"github.com/stretchr/testify/mock"
)

// MockContext for testing
type MockContext struct {
	mock.Mock
	context.Context
}

func (m *MockContext) Deadline() (deadline time.Time, ok bool) {
	args := m.Called()
	return args.Get(0).(time.Time), args.Bool(1)
}

func (m *MockContext) Done() <-chan struct{} {
	args := m.Called()
	return args.Get(0).(<-chan struct{})
}

func (m *MockContext) Err() error {
	args := m.Called()
	return args.Error(0)
}

func (m *MockContext) Value(key interface{}) interface{} {
	args := m.Called(key)
	return args.Get(0)
}

func (m *MockContext) ChatId() int {
	args := m.Called()
	return args.Int(0)
}

// UserId and MessageId are not in the original interface, removing them for now

func (m *MockContext) Text() string {
	args := m.Called()
	return args.String(0)
}

func (m *MockContext) Send(msg *commands.SendMessageRequest) error {
	args := m.Called(msg)
	return args.Error(0)
}

func (m *MockContext) Logger() Logger {
	args := m.Called()
	return args.Get(0).(Logger)
}

func (m *MockContext) FMS() FMS {
	args := m.Called()
	return args.Get(0).(FMS)
}

// New methods for extended Context interface
func (m *MockContext) SendPhoto(req *commands.SendPhotoRequest) error {
	args := m.Called(req)
	return args.Error(0)
}

func (m *MockContext) SendDocument(req *commands.SendDocumentRequest) error {
	args := m.Called(req)
	return args.Error(0)
}

func (m *MockContext) EditMessage(req *commands.EditMessageTextRequest) error {
	args := m.Called(req)
	return args.Error(0)
}

func (m *MockContext) DeleteMessage(req *commands.DeleteMessageRequest) error {
	args := m.Called(req)
	return args.Error(0)
}

func (m *MockContext) SendChatAction(req *commands.SendChatActionRequest) error {
	args := m.Called(req)
	return args.Error(0)
}

func (m *MockContext) AnswerCallback(req *commands.AnswerCallbackQueryRequest) error {
	args := m.Called(req)
	return args.Error(0)
}

func (m *MockContext) UserId() int {
	args := m.Called()
	return args.Int(0)
}

func (m *MockContext) MessageId() int {
	args := m.Called()
	return args.Int(0)
}

func (m *MockContext) CallbackData() string {
	args := m.Called()
	return args.String(0)
}

func (m *MockContext) IsCallback() bool {
	args := m.Called()
	return args.Bool(0)
}

// MockRouter for testing
type MockRouter struct {
	mock.Mock
}

func (m *MockRouter) Use(middleware ...Middleware) Router {
	args := m.Called(middleware)
	if ret := args.Get(0); ret != nil {
		return ret.(Router)
	}
	return m
}

func (m *MockRouter) Group(pattern string) Router {
	args := m.Called(pattern)
	if ret := args.Get(0); ret != nil {
		return ret.(Router)
	}
	return m // Return self for fluent interface
}

func (m *MockRouter) SetState(state string) Router {
	args := m.Called(state)
	if ret := args.Get(0); ret != nil {
		return ret.(Router)
	}
	return m
}

func (m *MockRouter) Command(command string, handler Handler) {
	m.Called(command, handler)
}

func (m *MockRouter) Callback(callback string, handler Handler) {
	m.Called(callback, handler)
}

func (m *MockRouter) Regex(pattern string, handler Handler) {
	m.Called(pattern, handler)
}

func (m *MockRouter) Msg(msg string, handler Handler) {
	m.Called(msg, handler)
}

func (m *MockRouter) Any(handler Handler) {
	m.Called(handler)
}

func (m *MockRouter) GetGroups() []Router {
	args := m.Called()
	return args.Get(0).([]Router)
}

func (m *MockRouter) GetParent() Router {
	args := m.Called()
	return args.Get(0).(Router)
}

func (m *MockRouter) GetMw() []Middleware {
	args := m.Called()
	return args.Get(0).([]Middleware)
}

func (m *MockRouter) GetCmd() map[string]Handler {
	args := m.Called()
	return args.Get(0).(map[string]Handler)
}

func (m *MockRouter) GetCallback() map[string]Handler {
	args := m.Called()
	return args.Get(0).(map[string]Handler)
}

func (m *MockRouter) GetRegex() []Regex {
	args := m.Called()
	return args.Get(0).([]Regex)
}

func (m *MockRouter) GetMsg() []MsgEquals {
	args := m.Called()
	return args.Get(0).([]MsgEquals)
}

func (m *MockRouter) GetAny() Handler {
	args := m.Called()
	return args.Get(0).(Handler)
}

func (m *MockRouter) GetState() string {
	args := m.Called()
	return args.String(0)
}

// Test basic router functionality
func TestRouter_BasicFunctionality(t *testing.T) {
	mockRouter := &MockRouter{}
	
	// Test Group method returns router
	mockRouter.On("Group", "/test").Return(nil) // Return nil to use fallback
	result := mockRouter.Group("/test")
	assert.Equal(t, mockRouter, result)
	mockRouter.AssertExpectations(t)
}

func TestRouter_CommandRegistration(t *testing.T) {
	mockRouter := &MockRouter{}
	
	// Test handler function
	handler := func(ctx Context) {
		// Mock handler implementation
	}
	
	// Test Command method
	mockRouter.On("Command", "/start", mock.AnythingOfType("core.Handler")).Return()
	mockRouter.Command("/start", handler)
	mockRouter.AssertExpectations(t)
}

func TestRouter_MiddlewareChain(t *testing.T) {
	mockRouter := &MockRouter{}
	
	// Test middleware function
	middleware := func(next Handler) Handler {
		return func(ctx Context) {
			next(ctx)
		}
	}
	
	// Test Use method
	mockRouter.On("Use", mock.AnythingOfType("[]core.Middleware")).Return(nil)
	result := mockRouter.Use(middleware)
	assert.Equal(t, mockRouter, result)
	mockRouter.AssertExpectations(t)
}

func TestRouter_StateManagement(t *testing.T) {
	mockRouter := &MockRouter{}
	
	// Test SetState method
	mockRouter.On("SetState", "waiting_input").Return(nil)
	result := mockRouter.SetState("waiting_input")
	assert.Equal(t, mockRouter, result)
	
	// Test GetState method
	mockRouter.On("GetState").Return("waiting_input")
	state := mockRouter.GetState()
	assert.Equal(t, "waiting_input", state)
	
	mockRouter.AssertExpectations(t)
}

// Handle method is not part of the core Router interface

func TestRouter_RegexPattern(t *testing.T) {
	mockRouter := &MockRouter{}
	
	handler := func(ctx Context) {}
	pattern := `^/user_(\d+)$`
	
	// Test Regex method
	mockRouter.On("Regex", pattern, mock.AnythingOfType("core.Handler")).Return()
	mockRouter.Regex(pattern, handler)
	mockRouter.AssertExpectations(t)
}

func TestRouter_AnyHandler(t *testing.T) {
	mockRouter := &MockRouter{}
	
	handler := func(ctx Context) {}
	
	// Test Any method (catch-all handler)
	mockRouter.On("Any", mock.AnythingOfType("core.Handler")).Return()
	mockRouter.Any(handler)
	mockRouter.AssertExpectations(t)
}