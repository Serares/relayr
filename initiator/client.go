package initiator

import (
	"bytes"
	"fmt"
	"net/http"
	"os"
	"time"
)

func NewClient() *http.Client {
	return &http.Client{
		Timeout: 10 * time.Second,
	}
}

// This is going to send the request containing info
// about what info the server has to request back
func initialiseRequest(filePath string) {
	file, err := os.Open(filePath)
	if err != nil {
		fmt.Println("Error reading file")
		return
	}

	defer file.Close()
	var requestBody bytes.Buffer

	req, err := http.NewRequest(http.MethodPost, "localhost:8080/initialise", &requestBody)
	if err != nil {
		fmt.Println("Error creating request")
		return
	}
	req.Header.Set("x-required-data", "iv-salt")

	r, err := NewClient().Do(req)
	if err != nil {
		fmt.Println("Error sending request")
		return
	}

	defer r.Body.Close()

}
