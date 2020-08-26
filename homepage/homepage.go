package homepage

import (
	"net/http"
)

const message = "Hello from Rahim yar khan!"

func HomeHandler(writer http.ResponseWriter, _ *http.Request) {
	writer.Header().Set("Content-Type", "text/plain; charset=utf-8")
	writer.WriteHeader(http.StatusOK)
	writer.Write([]byte(message))
}
