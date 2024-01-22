package tests

import (
	"net/http"
	"net/http/httptest"
	"testing"

	service "github.com/Laem20957/records-app/internal/services"
	"github.com/Laem20957/records-app/internal/transport/rest"
	"github.com/stretchr/testify/require"
)

func TestGetHandler(test *testing.T) {
	handler := rest.GetHandler(&service.ServiceMethods{})
	require.IsType(test, rest.Handler{}, handler)
}

func TestHandler_InitRoutes(test *testing.T) {
	handler := rest.GetHandler(&service.ServiceMethods{})

	ts := httptest.NewServer(handler.InitRoutes())
	defer ts.Close()

	res, err := http.Get(ts.URL + "/ping")
	if err != nil {
		test.Error(err)
	}

	require.Equal(test, http.StatusOK, res.StatusCode)
}
