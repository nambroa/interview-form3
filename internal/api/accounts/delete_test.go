package accounts

import (
	"github.com/nambroa/interview-accountapi/internal"
	uuid "github.com/nu7hatch/gouuid"
	"github.com/stretchr/testify/assert"
	"net/http"
	"strconv"
	"testing"
)

func TestDelete_ValidCreatedAccount(t *testing.T) {
	accountBuilder := internal.DefaultAccountBuilder()
	acc, err := accountBuilder.Build()
	if assert.Nil(t, err) {
		// Create account
		_, err := Create(acc)
		if assert.Nil(t, err) {
			// Delete account
			response, err := Delete(acc.Data.ID, strconv.FormatInt(*acc.Data.Version, 10))
			assert.Nil(t, err)
			if assert.NotNil(t, response) {
				assert.Equal(t, response.StatusCode, http.StatusNoContent)
			}
		}
	}
}

func TestDelete_WithNonExistentIDReturnsNotFound(t *testing.T) {
	fakeID, _ := uuid.NewV4()
	response, err := Delete(fakeID.String(), "0")
	assert.NotNil(t, err)
	if assert.NotNil(t, response) {
		assert.Equal(t, response.StatusCode, http.StatusNotFound)
	}

}

func TestDelete_WithNonExistentVersionReturnsConflict(t *testing.T) {
	accountBuilder := internal.DefaultAccountBuilder()
	acc, err := accountBuilder.Build()
	if assert.Nil(t, err) {
		// Create account
		_, err := Create(acc)
		if assert.Nil(t, err) {
			// Delete account
			response, err := Delete(acc.Data.ID, "3222423")
			assert.NotNil(t, err)
			if assert.NotNil(t, response) {
				assert.Equal(t, response.StatusCode, http.StatusConflict)

			}
		}
	}
}
