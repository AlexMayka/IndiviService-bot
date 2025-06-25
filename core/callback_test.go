package core

import (
	"testing"
	"telegram-sdk/types/commands"

	"github.com/stretchr/testify/assert"
)

func TestContext_CallbackMethods(t *testing.T) {
	mockContext := &MockContext{}

	// Test IsCallback method
	mockContext.On("IsCallback").Return(true)
	isCallback := mockContext.IsCallback()
	assert.True(t, isCallback)

	// Test CallbackData method
	testData := "button_clicked"
	mockContext.On("CallbackData").Return(testData)
	callbackData := mockContext.CallbackData()
	assert.Equal(t, testData, callbackData)

	mockContext.AssertExpectations(t)
}

func TestContext_AnswerCallback(t *testing.T) {
	mockContext := &MockContext{}

	request := &commands.AnswerCallbackQueryRequest{
		CallbackQueryID: "12345",
		Text:           "Button clicked!",
		ShowAlert:      false,
	}

	// Test AnswerCallback method
	mockContext.On("AnswerCallback", request).Return(nil)
	err := mockContext.AnswerCallback(request)
	assert.NoError(t, err)

	mockContext.AssertExpectations(t)
}

func TestContext_EditMessage(t *testing.T) {
	mockContext := &MockContext{}

	request := &commands.EditMessageTextRequest{
		ChatID:    "12345",
		MessageID: 100,
		Text:      "Updated message text",
	}

	// Test EditMessage method
	mockContext.On("EditMessage", request).Return(nil)
	err := mockContext.EditMessage(request)
	assert.NoError(t, err)

	mockContext.AssertExpectations(t)
}

func TestContext_DeleteMessage(t *testing.T) {
	mockContext := &MockContext{}

	request := &commands.DeleteMessageRequest{
		ChatID:    "12345",
		MessageID: 100,
	}

	// Test DeleteMessage method
	mockContext.On("DeleteMessage", request).Return(nil)
	err := mockContext.DeleteMessage(request)
	assert.NoError(t, err)

	mockContext.AssertExpectations(t)
}

func TestContext_SendChatAction(t *testing.T) {
	mockContext := &MockContext{}

	request := &commands.SendChatActionRequest{
		ChatID: "12345",
		Action: commands.ActionTyping,
	}

	// Test SendChatAction method
	mockContext.On("SendChatAction", request).Return(nil)
	err := mockContext.SendChatAction(request)
	assert.NoError(t, err)

	mockContext.AssertExpectations(t)
}

func TestContext_UserAndMessageInfo(t *testing.T) {
	mockContext := &MockContext{}

	// Test UserId method
	mockContext.On("UserId").Return(67890)
	userId := mockContext.UserId()
	assert.Equal(t, 67890, userId)

	// Test MessageId method
	mockContext.On("MessageId").Return(111)
	messageId := mockContext.MessageId()
	assert.Equal(t, 111, messageId)

	mockContext.AssertExpectations(t)
}