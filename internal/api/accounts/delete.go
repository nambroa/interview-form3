package accounts

import (
	"errors"
	"fmt"
	"github.com/nambroa/interview-accountapi/internal"
	"log"
	"net/http"
)

func Delete(accountID string, version string) (*http.Response, error) {
	var deleteAccountURL = internal.AccountURL + "/" + accountID + "?version=" + version

	// Delete account
	httpClient := &http.Client{}
	// Build request
	request, err := http.NewRequest(http.MethodDelete, deleteAccountURL, nil)
	if err != nil {
		log.Println("Error found while building delete account request:", err)
		return nil, err
	}
	response, err := httpClient.Do(request)
	if err != nil {
		log.Println("Error found while deleting account:", err)
		return nil, err
	}

	// Process response
	if response.StatusCode != http.StatusNoContent {
		log.Println("Response returned error status code:", response.StatusCode)
		return response, errors.New(fmt.Sprintf("Status code: %d", response.StatusCode))
	}
	return response, nil
}
