package accounts

import (
	"github.com/nambroa/interview-accountapi/internal"
	uuid "github.com/nu7hatch/gouuid"
	"github.com/stretchr/testify/assert"
	"net/http"
	"strconv"
	"testing"
)

func TestFetch_ValidCreatedAccount(t *testing.T) {
	accountBuilder := internal.DefaultAccountBuilder()
	acc, err := accountBuilder.Build()
	if assert.Nil(t, err) {
		// Create account
		_, err := Create(acc)
		if assert.Nil(t, err) {
			// Fetch account
			fetchedAcc, err := Fetch(acc.Data.ID)
			assert.Nil(t, err)
			assert.Equal(t, acc.Data.ID, fetchedAcc.Data.ID)
		}
	}
}

func TestFetch_WithNonExistentIDReturnsNotFound(t *testing.T) {
	fakeID, _ := uuid.NewV4()
	fetchedAcc, err := Fetch(fakeID.String())
	assert.NotNil(t, err)
	assert.Nil(t, fetchedAcc)
	assert.ErrorContains(t, err, strconv.Itoa(http.StatusNotFound))
}
