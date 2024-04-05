package utils

import (
	"fmt"
	"net/http"
)

func healthCheckHandler(response http.ResponseWriter, request *http.Request) {
	fmt.Fprint(response, "HealthCheck: OK")
}
