package tests

import (
	"bytes"
	"fmt"
	"os/exec"
	"testing"

	"github.com/stretchr/testify/assert"
)

func TestSwaggerCommentsGeneration(t *testing.T) {
	mainFilePath := "../cmd/main.go"

	cmd := exec.Command("swag", "init", "-g", mainFilePath)
	var output bytes.Buffer
	var stderr bytes.Buffer
	cmd.Stdout = &output
	cmd.Stderr = &stderr
	err := cmd.Run()

	if err != nil {
		fmt.Println("Error:", err)
		fmt.Println("Output:", output.String())
		fmt.Println("Errors:", stderr.String())
		t.Fail()
	} else {
		assert.Empty(t, stderr.String(), "Не ожидается наличие ошибок")
		assert.NotEmpty(t, output.String(), "Ожидается наличие вывода")
	}
}
