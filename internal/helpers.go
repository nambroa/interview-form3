package internal

import (
	"github.com/nambroa/interview-accountapi/internal/models/builder"
	uuid "github.com/nu7hatch/gouuid"
)

// DefaultAccountBuilder returns an AccountBuilder filled with the minimum data required for an account to be built.
func DefaultAccountBuilder() *builder.AccountBuilder {
	ID, _ := uuid.NewV4()
	OrganisationID, _ := uuid.NewV4()
	var names []string
	names = append(names, "Paul")
	names = append(names, "Jason")
	names = append(names, "Robin")
	var bankID = "400300"
	var bankIDCode = "GBDSC"
	var BIC = "NWBKGB22"
	var country = "GB"
	// Basic Builder
	var accountBuilder = builder.NewAccountBuilder(ID.String(), OrganisationID.String(), bankID, bankIDCode,
		BIC, country, names)
	// Create Account
	return accountBuilder

}
