package services

import (
	"fmt"
	"github.com/stretchr/testify/assert"
	"golang-microservices/mvc/domain"
	"golang-microservices/mvc/utils"
	"net/http"
	"testing"
)

var (
	//userDaoMock usersDaoMock
	getUserFunction func(userId int64) (*domain.User, *utils.ApplicationError)
)

type usersDaoMock struct{}

func (*usersDaoMock) GetUser(userId int64) (*domain.User, *utils.ApplicationError) {
	return getUserFunction(userId)
}

//func TestMain(m *testing.M) {
//	//domain.UserDao{}
//	domain.IUserDao = &userDaoMock{}
//	os.Exit(m.Run())
//}
func init(){
	domain.IUserDao = &usersDaoMock{}
}

func TestGetUserNotFound(t *testing.T) {
	var userId = 456
	getUserFunction = func(userId int64) (*domain.User, *utils.ApplicationError) {
		return nil, &utils.ApplicationError{
			StatusCode: http.StatusNotFound,
			Message:    fmt.Sprintf("user %v does not exists", userId),
		}
	}
	user, err := GetUser(int64(userId))

	assert.NotNil(t, err)
	assert.Nil(t, user)
	assert.EqualValues(t, http.StatusNotFound, err.StatusCode)
	assert.EqualValues(t, fmt.Sprintf("user %v does not exists", userId), err.Message)
}

func TestGetUserNoError(t *testing.T) {
	var userId = 123
	getUserFunction = func(userId int64) (*domain.User, *utils.ApplicationError) {
		return &domain.User{
			Id:        123,
			FirstName: "Farouk",
			LastName:  "Elkholy",
			Email:     "myemail@gmail.com",
		}, nil
	}

	user, err := GetUser(int64(userId))

	assert.Nil(t, err)
	assert.NotNil(t, user)
	assert.EqualValues(t, userId, user.Id)
	assert.EqualValues(t, "Farouk", user.FirstName)
	assert.EqualValues(t, "Elkholy", user.LastName)
	assert.EqualValues(t, "myemail@gmail.com", user.Email)
}
