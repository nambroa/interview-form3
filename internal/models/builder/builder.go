package builder

import (
	"encoding/json"
	"github.com/go-playground/validator/v10"
	"github.com/nambroa/interview-accountapi/internal/models"
)

// AccountBuilder represents a builder that builds Accounts. It also contains validations (inside the Account itself)
//for specific fields to stop the creation of invalid accounts before it happens.
// The definition of a valid account is taken from the restrictions of the documentation:
// https://www.api-docs.form3.tech/api/schemes/fps-direct/accounts/accounts/create-an-account
// Examples: ID must be UUID, BankID must be max length 11, Bic must be 8 or 11 chars longs, Country must be 2 chars long, etc.

type AccountBuilder struct {
	account *models.Account
}

// NewAccountBuilder contains required fields according to documentation https://www.api-docs.form3.tech/api/schemes/fps-direct/accounts/accounts/create-an-account
// and testing of the API (meaning that name and type fields are required as well).
func NewAccountBuilder(ID, organisationID, bankID, bankIDCode, bic, country string, names []string) *AccountBuilder {
	// Defaults are selected according to the docs https://www.api-docs.form3.tech/api/schemes/fps-direct/accounts/accounts/create-an-account
	var defaultAccountClassification = models.PERSONAL
	var defaultJointAccount = false
	var defaultVersion int64 = 0
	var defaultNMS = models.SUPPORTED
	var defaultBaseCurrency = "GBP"

	accountAttributes := &models.AccountAttributes{
		AccountClassification:   &defaultAccountClassification,
		AccountNumber:           "",
		AlternativeNames:        nil,
		BankID:                  bankID,
		BankIDCode:              bankIDCode,
		BaseCurrency:            defaultBaseCurrency,
		Bic:                     bic,
		Country:                 &country,
		Iban:                    "",
		JointAccount:            &defaultJointAccount,
		Name:                    names,
		NameMatchingStatus:      &defaultNMS,
		SecondaryIdentification: "",
		Status:                  nil,
	}

	accountData := &models.AccountData{
		Attributes:     accountAttributes,
		ID:             ID,
		OrganisationID: organisationID,
		Type:           models.ACCOUNTS,
		Version:        &defaultVersion,
	}

	account := &models.Account{Data: accountData}

	return &AccountBuilder{account: account}
}

func (ab *AccountBuilder) WithVersion(version *int64) *AccountBuilder {
	ab.account.Data.Version = version
	return ab
}

func (ab *AccountBuilder) WithAccountClassification(classification *models.AccountClassification) *AccountBuilder {
	ab.account.Data.Attributes.AccountClassification = classification
	return ab
}

func (ab *AccountBuilder) WithNameMatchingStatus(matching *models.NameMatchingStatus) *AccountBuilder {
	ab.account.Data.Attributes.NameMatchingStatus = matching
	return ab
}

func (ab *AccountBuilder) WithAccountNumber(number string) *AccountBuilder {
	ab.account.Data.Attributes.AccountNumber = number
	return ab
}

func (ab *AccountBuilder) WithAlternativeNames(alternativeNames []string) *AccountBuilder {
	ab.account.Data.Attributes.AlternativeNames = alternativeNames
	return ab
}

func (ab *AccountBuilder) WithBaseCurrency(baseCurrency string) *AccountBuilder {
	ab.account.Data.Attributes.BaseCurrency = baseCurrency
	return ab
}

func (ab *AccountBuilder) WithIban(iban string) *AccountBuilder {
	ab.account.Data.Attributes.Iban = iban
	return ab
}

func (ab *AccountBuilder) WithJointAccount(jointAccount *bool) *AccountBuilder {
	ab.account.Data.Attributes.JointAccount = jointAccount
	return ab
}

func (ab *AccountBuilder) WithSecondaryIdentification(secondaryIdentification string) *AccountBuilder {
	ab.account.Data.Attributes.SecondaryIdentification = secondaryIdentification
	return ab
}

func (ab *AccountBuilder) WithStatus(status *models.AccountStatus) *AccountBuilder {
	ab.account.Data.Attributes.Status = status
	return ab
}

// Build validates the account inside the builder and returns it alongside validation data.
func (ab *AccountBuilder) Build() (*models.Account, error) {
	validate := validator.New()
	err := validate.Struct(ab.account)
	if err != nil {
		return nil, err
	}

	return ab.account, nil
}

// FromJSON Creates an account builder with an account marshalled from the json byte array. It will not build the account.
func FromJSON(accountJSON []byte) (*AccountBuilder, error) {
	var account models.Account
	err := json.Unmarshal(accountJSON, &account)
	if err != nil {
		return nil, err
	}
	return &AccountBuilder{account: &account}, nil
}
