package tests

import (
	_ "net/http"
	_ "net/http/httptest"
	_ "testing"

	_ "github.com/Laem20957/records-app/internal/service"
	_ "github.com/Laem20957/records-app/internal/transport/rest"
	_ "github.com/stretchr/testify/require"
)

/*
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
*/
