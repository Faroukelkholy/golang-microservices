package domain

import (
	"github.com/stretchr/testify/assert"
	"net/http"
	"testing"
)

func TestGetUserNoUserFound(t *testing.T) {
	// Initialization:

	// Execution:
	user, err := IUserDao.GetUser(0)

	// Validation:
	assert.Nil(t, user, "we were not expecting a user with id 0")
	assert.NotNil(t, err)

	assert.EqualValues(t, http.StatusNotFound, err.StatusCode)
	assert.EqualValues(t, "not_found", err.Code)
	assert.EqualValues(t, "user 0 does not exists", err.Message)
}

func TestGetUserNoError(t *testing.T) {
	user, err := IUserDao.GetUser(123)

	assert.Nil(t, err)
	assert.NotNil(t, user)

	assert.EqualValues(t, 123, user.Id)
	assert.EqualValues(t, "Farouk", user.FirstName)
	assert.EqualValues(t, "Elkholy", user.LastName)
	assert.EqualValues(t, "myemail@gmail.com", user.Email)
}
