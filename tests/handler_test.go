package tests

import (
	"net/http"
	"net/http/httptest"
	"testing"

	"github.com/Laem20957/records-app/internal/service"
	"github.com/Laem20957/records-app/internal/transport/rest"
	"github.com/stretchr/testify/require"
)

func TestNewHandler(t *testing.T) {
	h := rest.NewHandler(&service.Service{})

	require.IsType(t, &rest.Handler{}, h)
}

func TestHandler_InitRoutes(t *testing.T) {
	h := rest.NewHandler(&service.Service{})

	ts := httptest.NewServer(h.InitRoutes())
	defer ts.Close()

	res, err := http.Get(ts.URL + "/ping")
	if err != nil {
		t.Error(err)
	}

	require.Equal(t, http.StatusOK, res.StatusCode)
}
