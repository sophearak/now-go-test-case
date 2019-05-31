package handlertwo

import (
	"net/http"
	shared "github.com/sophearak/now-go-test-case/shared-package"
)

// Handler func
func Handler(w http.ResponseWriter, r *http.Request) {
	message := shared.ResponseGenerator("test-2")
	w.Write([]byte(message))
}