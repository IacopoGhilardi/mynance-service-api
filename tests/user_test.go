package tests

import (
	"context"
	"log"
	"testing"

	service "github.com/iacopoghilardi/mynance-service-api/api/v1/services"
	"github.com/iacopoghilardi/mynance-service-api/internal/app"
	"github.com/iacopoghilardi/mynance-service-api/models"
	"github.com/stretchr/testify/assert"
	"github.com/stretchr/testify/suite"
)

type UserServiceTestSuite struct {
	suite.Suite
	pgContainer *PostgresContainer
	service     *service.UserService
	ctx         context.Context
}

func (suite *UserServiceTestSuite) SetupSuite() {
	suite.ctx = context.Background()
	pgContainer, err := CreatePostgresContainer(suite.ctx)
	if err != nil {
		log.Fatal(err)
	}
	suite.pgContainer = pgContainer
	app.InitApp()
	if err != nil {
		log.Fatal(err)
	}
	suite.service = service.V1Services.UserService
}

func (suite *UserServiceTestSuite) TearDownSuite() {
	if err := suite.pgContainer.Terminate(suite.ctx); err != nil {
		log.Fatalf("error terminating postgres container: %s", err)
	}
}

func (suite *UserServiceTestSuite) TestCreateUser() {
	t := suite.T()

	err := suite.service.CreateUser(suite.ctx, &models.User{
		Password: "Henry",
		Email:    "henry@gmail.com",
	})
	assert.NoError(t, err)

	users, err := suite.service.GetAllUsers(suite.ctx)
	assert.NoError(t, err)
	assert.Equal(t, len(users), 1)

}

func TestUserRepoTestSuite(t *testing.T) {
	suite.Run(t, new(UserServiceTestSuite))
}
