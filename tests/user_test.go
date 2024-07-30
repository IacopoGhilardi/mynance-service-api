package tests

import (
	"bytes"
	"context"
	"encoding/json"
	"net/http"
	"net/http/httptest"
	"os"
	"strconv"
	"testing"

	"github.com/gin-gonic/gin"
	service "github.com/iacopoghilardi/mynance-service-api/api/v1/services"
	"github.com/iacopoghilardi/mynance-service-api/internal/app"
	"github.com/iacopoghilardi/mynance-service-api/internal/database"
	"github.com/iacopoghilardi/mynance-service-api/internal/router"
	"github.com/iacopoghilardi/mynance-service-api/models"
	"github.com/stretchr/testify/assert"
	"github.com/stretchr/testify/suite"
)

type EndpointTestSuite struct {
	suite.Suite
	pgContainer *PostgresContainer
	ctx         context.Context
	router      *gin.Engine
}

func (suite *EndpointTestSuite) SetupSuite() {
	suite.ctx = context.Background()
	pgContainer, err := CreatePostgresContainer(suite.ctx)
	if err != nil {
		suite.T().Fatal(err)
	}
	suite.pgContainer = pgContainer

	os.Setenv("APP_ENV", "testing")
	os.Setenv("GIN_MODE", "test")
	defer os.Unsetenv("APP_ENV")
	defer os.Unsetenv("GIN_MODE")

	err = database.ConnectToDb(pgContainer.ConnectionString)
	if err != nil {
		suite.T().Fatal(err)
	}

	if err := app.InitApp(); err != nil {
		suite.T().Fatal(err)
	}

	suite.router = router.GetRouter()
}

func (suite *EndpointTestSuite) TearDownSuite() {
	if err := suite.pgContainer.Terminate(suite.ctx); err != nil {
		suite.T().Fatalf("error terminating postgres container: %s", err)
	}
	database.CloseDb()
}

func (suite *EndpointTestSuite) TestUserEndpoints() {
	t := suite.T()

	user := models.User{
		Email:    "henry@gmail.com",
		Password: "testPassword",
	}

	jsonUser, _ := json.Marshal(user)
	createReq, _ := http.NewRequest("POST", "/api/v1/users/", bytes.NewBuffer(jsonUser))
	createReq.Header.Set("Content-Type", "application/json")

	w := httptest.NewRecorder()
	suite.router.ServeHTTP(w, createReq)
	assert.Equal(t, http.StatusCreated, w.Code)

	users, _ := service.V1Services.UserService.GetAllUsers(suite.ctx)
	firstUser := users[0]
	assert.Equal(t, user.Email, firstUser.Email)
	assert.Equal(t, 1, len(users))

	updatedUser := models.User{
		Email:    "henry2@gmail.com",
		Password: "testPassword2",
	}

	jsonUpdatedUser, _ := json.Marshal(updatedUser)
	userID := strconv.FormatUint(uint64(firstUser.ID), 10)
	updateReq, _ := http.NewRequest("PUT", "/api/v1/users/"+userID, bytes.NewBuffer(jsonUpdatedUser))
	updateReq.Header.Set("Content-Type", "application/json")

	users, _ = service.V1Services.UserService.GetAllUsers(suite.ctx)
	firstUser = users[0]
	assert.Equal(t, updatedUser.Email, firstUser.Email)
	//assert.Equal(t, 1, len(users))
	//
	//suite.router.ServeHTTP(w, updateReq)
	//assert.Equal(t, http.StatusCreated, w.Code)
}

func (suite *EndpointTestSuite) TestCreateUserEndpoint() {
	t := suite.T()

	user := models.User{
		Email:    "henry@gmail.com",
		Password: "testPassword",
	}

	jsonUser, _ := json.Marshal(user)
	req, _ := http.NewRequest("POST", "/api/v1/users/", bytes.NewBuffer(jsonUser))
	req.Header.Set("Content-Type", "application/json")

	w := httptest.NewRecorder()
	suite.router.ServeHTTP(w, req)

	assert.Equal(t, http.StatusCreated, w.Code)

	var response map[string]interface{}
	err := json.Unmarshal(w.Body.Bytes(), &response)
	assert.NoError(t, err)
	assert.Equal(t, "OK", response["status"])

	users, _ := service.V1Services.UserService.GetAllUsers(suite.ctx)
	assert.Equal(t, 1, len(users))
}

func (suite *EndpointTestSuite) TestUpdateUser() {
	t := suite.T()

	user := models.User{
		Email:    "testUpdate@gmail.com",
		Password: "testPassword",
	}

	jsonUser, _ := json.Marshal(user)
	req, _ := http.NewRequest("POST", "/api/v1/users/", bytes.NewBuffer(jsonUser))
	req.Header.Set("Content-Type", "application/json")

	w := httptest.NewRecorder()
	suite.router.ServeHTTP(w, req)

	assert.Equal(t, http.StatusCreated, w.Code)

	var response map[string]interface{}
	err := json.Unmarshal(w.Body.Bytes(), &response)
	assert.NoError(t, err)
	assert.Equal(t, "OK", response["status"])

	users, _ := service.V1Services.UserService.GetAllUsers(suite.ctx)
	assert.Equal(t, 1, len(users))
}

func TestEndpointTestSuite(t *testing.T) {
	suite.Run(t, new(EndpointTestSuite))
}
