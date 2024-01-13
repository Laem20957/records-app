package tests

import (
	"bytes"
	"context"
	"fmt"
	"net/http/httptest"
	"testing"

	"github.com/Laem20957/records-app/internal/domains"
	"github.com/Laem20957/records-app/internal/services"
	"github.com/Laem20957/records-app/internal/transport/rest"
	serviceMock "github.com/Laem20957/records-app/mocks"
	"github.com/gin-gonic/gin"
	"github.com/golang/mock/gomock"
	"github.com/stretchr/testify/assert"
)

func TestHandler_signUp(t *testing.T) {
	type mockBehavior func(r *serviceMock.MockAuthorization, ctx context.Context, user domain.User)

	tests := []struct {
		name                 string
		inputBody            string
		inputUser            domain.User
		mockBehavior         mockBehavior
		ctx                  context.Context
		expectedStatusCode   int
		expectedResponseBody string
	}{
		{
			name:      "ok",
			inputBody: `{"id": 0, "username": "username","name": "test name","password": "123456qw"}`,
			inputUser: domain.User{
				Id:       0,
				Username: "username",
				Name:     "test name",
				Password: "123456qw",
			},
			mockBehavior: func(r *serviceMock.MockAuthorization, ctx context.Context, user domain.User) {
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

			services := &service.Service{Authorization: repo}
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
