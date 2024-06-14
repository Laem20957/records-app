package tests

import (
	_ "bytes"
	_ "context"
	_ "fmt"
	_ "net/http/httptest"
	_ "testing"

	_ "records-app/api/rest"
	_ "records-app/internal/domain"
	_ "records-app/internal/service"
	_ "records-app/mocks"

	_ "github.com/gin-gonic/gin"
	_ "github.com/golang/mock/gomock"
	_ "github.com/stretchr/testify/assert"
)

/*
func TestHandler_signUp(t *testing.T) {
	type mockBehavior func(r *serviceMock.MockAuthorization, ctx context.Context, user domain.Users)

	tests := []struct {
		name                 string
		inputBody            string
		inputUser            domain.Users
		mockBehavior         mockBehavior
		ctx                  context.Context
		expectedStatusCode   int
		expectedResponseBody string
	}{
		{
			name:      "ok",
			inputBody: `{"id": 0, "username": "username","name": "test name","password": "123456qw"}`,
			inputUser: domain.Users{
				Id:       0,
				Username: "username",
				Name:     "test name",
				Password: "123456qw",
			},
			mockBehavior: func(r *serviceMock.MockAuthorization, ctx context.Context, user domain.Users) {
				r.EXPECT().CreateUser(ctx, user).Return(1, nil)
			},
			ctx:                  context.Background(),
			expectedStatusCode:   200,
			expectedResponseBody: `{"id": 1}`,
		},
		//{
		//	name:                 "wrong input",
		//	inputBody:            `{"username": "username"}`,
		//	inputUser:            domain.User{},
		//	mockBehavior:         func(r *serviceMock.MockAuthorization, ctx *gin.Context, user domain.User) {},
		//	expectedStatusCode:   400,
		//	expectedResponseBody: "{\"message\":\"invalid input body\"}",
		//},
		//{
		//	name:      "service error",
		//	inputBody: `{"username": "username", "name": "test name", "password": "123456qw"}`,
		//	inputUser: domain.User{
		//		Username: "username",
		//		Name:     "test name",
		//		Password: "123456qw",
		//	},
		//	mockBehavior: func(r *serviceMock.MockAuthorization, ctx *gin.Context, user domain.User) {
		//		r.EXPECT().CreateUser(ctx, user).Return(0, errors.New("internal server error"))
		//	},
		//	expectedStatusCode:   500,
		//	expectedResponseBody: `{"message":"internal server error"}`,
		//},
	}

	for _, test := range tests {
		t.Run(test.name, func(t *testing.T) {
			c := gomock.NewController(t)
			defer c.Finish()

			repo := serviceMock.NewMockAuthorization(c)
			test.mockBehavior(repo, test.ctx, test.inputUser)

			services := &service.ServiceMethods{IServiceAuthorizationMethods: nil}
			_ = rest.Handler{Services: services}

			r := gin.New()
			r.POST("/sign-up", nil)

			w := httptest.NewRecorder()
			req := httptest.NewRequest("POST", "/sign-up", bytes.NewBufferString(test.inputBody))

			r.ServeHTTP(w, req)

			fmt.Println("!!!!!!!!!!!!!!!!!!!!", w.Body.String())

			assert.Equal(t, test.expectedStatusCode, w.Code)
			assert.Equal(t, test.expectedResponseBody, w.Body.String())
		})
	}
}

*/
