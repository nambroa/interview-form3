package accounts

import (
	"github.com/nambroa/interview-accountapi/internal"
	"github.com/nambroa/interview-accountapi/internal/models"
	"github.com/stretchr/testify/assert"
	"net/http"
	"testing"
)

func TestCreate_ValidAccount(t *testing.T) {
	var accountBuilder = internal.DefaultAccountBuilder()
	account, err := accountBuilder.Build()
	if assert.Nil(t, err) {
		response, err := Create(account)
		assert.Nil(t, err)
		assert.NotNil(t, response)
		if assert.NotNil(t, response) {
			assert.Equal(t, response.StatusCode, http.StatusCreated)
		}
	}
}

func TestCreate_FullAccount(t *testing.T) {
	var accountBuilder = internal.DefaultAccountBuilder()
	var names []string
	names = append(names, "Bruce")
	names = append(names, "Jack")
	names = append(names, "Jason")
	var jointAcc = false
	var accClass = models.BUSINESS
	var NMS = models.OPTED_OUT
	var version int64 = 1

	accountBuilder.WithAlternativeNames(names)
	accountBuilder.WithIban("AB12AZ33")
	accountBuilder.WithSecondaryIdentification("Alfred")
	accountBuilder.WithJointAccount(&jointAcc)
	accountBuilder.WithAccountClassification(&accClass)
	accountBuilder.WithBaseCurrency("ARS")
	accountBuilder.WithAccountNumber("1234567890")
	accountBuilder.WithNameMatchingStatus(&NMS)
	accountBuilder.WithVersion(&version)

	account, err := accountBuilder.Build()
	if assert.Nil(t, err) {
		response, err := Create(account)
		assert.Nil(t, err)
		if assert.NotNil(t, response) {
			assert.Equal(t, response.StatusCode, http.StatusCreated)
		}
	}
}

// Since IBAN does not have validators, I can create invalid accounts to showcase tests where the API returns bad request.
// In the future I would like to learn how the IBAN number is composed and create a custom validator using the library to validate IBAN (possibly with a regex).

func TestCreate_AccountWithInvalidIBANReturnsBadRequest(t *testing.T) {
	var accountBuilder = internal.DefaultAccountBuilder()
	accountBuilder.WithIban("$#*$*(@*($@*#$*&!!!!!!!!!!!!!!!!!%%^^#$!!!!!!!!!!!!!!!!!!")
	account, err := accountBuilder.Build()
	if assert.Nil(t, err) {
		response, err := Create(account)
		assert.NotNil(t, err)
		if assert.NotNil(t, response) {
			assert.Equal(t, response.StatusCode, http.StatusBadRequest)
		}
	}
}
