package tests

import (
	"bytes"
	"fmt"
	"os"
	"os/exec"
	"testing"

	"io/fs"

	"github.com/stretchr/testify/assert"
)

func TestSwaggerCommentsWriteJSON(t *testing.T) {
	mainFilePath := "../cmd/main.go"
	testComment := "// @Summary Create an account"

	cmd := exec.Command("swag", "init", "-g", mainFilePath)
	var output bytes.Buffer
	var stderr bytes.Buffer
	cmd.Stdout = &output
	cmd.Stderr = &stderr
	err := cmd.Run()

	// Проверка на отсутствие ошибок при выполнении команды swag init
	if err != nil {
		fmt.Println("Error:", err)
		fmt.Println("Output:", output.String())
		fmt.Println("Errors:", stderr.String())
		t.Fail()
	} else {
		assert.Empty(t, stderr.String(), "Не ожидается наличие ошибок")

		// Проверка наличия комментариев в файле swagger.json
		swaggerJSONContent, err := fs.ReadFile(os.DirFS(".."), "docs/swagger.json")
		if err != nil {
			fmt.Println("Error while reading swagger.json:", err)
			t.Fail()
		}
		assert.Contains(t, string(swaggerJSONContent), testComment, "Комментарий отсутствует в swagger.json")

		// Проверка наличия комментариев в файле swagger.yaml
		swaggerYAMLContent, err := fs.ReadFile(os.DirFS(".."), "docs/swagger.yaml")
		if err != nil {
			fmt.Println("Error while reading swagger.yaml:", err)
			t.Fail()
		}
		assert.Contains(t, string(swaggerYAMLContent), testComment, "Комментарий отсутствует в swagger.yaml")

		// Проверка наличия комментариев в файле docs.go
		docsGoContent, err := fs.ReadFile(os.DirFS(".."), "docs/docs.go")
		if err != nil {
			fmt.Println("Error while reading docs.go:", err)
			t.Fail()
		}
		assert.Contains(t, string(docsGoContent), testComment, "Комментарий отсутствует в docs.go")
	}
}
