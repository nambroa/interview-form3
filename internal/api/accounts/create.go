package accounts

import (
	"bytes"
	"encoding/json"
	"errors"
	"fmt"
	"github.com/nambroa/interview-accountapi/internal"
	"github.com/nambroa/interview-accountapi/internal/models"
	"io"
	"log"
	"net/http"
)

// Create sends an account payload to the fake API to create an account. It returns its associated response and error data.
func Create(payload *models.Account) (*http.Response, error) {

	// Convert account data to json
	marshalledAccount, err := json.Marshal(payload)
	if err != nil {
		log.Println("Error marhsalling account data:", err)
		return nil, err
	}
	// Create account
	response, err := http.Post(internal.AccountURL, "application/json", bytes.NewReader(marshalledAccount))

	// Process response
	if err != nil {
		log.Println("Error found while creating account:", err)
		return nil, err
	} else {
		defer response.Body.Close()
		body, _ := io.ReadAll(response.Body)
		if response.StatusCode != http.StatusCreated {
			log.Println("Response returned error status code:", response.StatusCode)
			log.Println(response)
			// Read body and return it in the response as error.
			return response, errors.New(fmt.Sprintf("Status code: %d. Body: %s", response.StatusCode, string(body)))
		} else {
			return response, nil
		}
	}
}
