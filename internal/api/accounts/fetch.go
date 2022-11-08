package accounts

import (
	"errors"
	"fmt"
	"github.com/nambroa/interview-accountapi/internal"
	"github.com/nambroa/interview-accountapi/internal/models"
	"github.com/nambroa/interview-accountapi/internal/models/builder"
	"io"
	"log"
	"net/http"
)

// Fetch fetches an account from the fake API based on its ID.
func Fetch(accountID string) (*models.Account, error) {
	var fetchAccountURL = internal.AccountURL + "/" + accountID

	// Fetch account
	response, err := http.Get(fetchAccountURL)

	// Process response
	if err != nil {
		log.Println("Error found while fetching account:", err)
		return nil, err
	}
	defer response.Body.Close()
	accountJSON, _ := io.ReadAll(response.Body)
	if response.StatusCode != http.StatusOK {
		log.Println("Response returned error status code:", response.StatusCode)
		return nil, errors.New(fmt.Sprintf("Status code: %d. Body: %s", response.StatusCode, string(accountJSON)))
	}

	// Unmarshal payload into account.
	accountBuilder, err := builder.FromJSON(accountJSON)
	if err != nil {
		log.Println("Error found while unmarshalling account:", err)
		return nil, err
	}

	// Build account.
	account, err := accountBuilder.Build()
	if err != nil {
		log.Println("Error found while building account:", err)
		return nil, err
	}

	return account, nil
}
