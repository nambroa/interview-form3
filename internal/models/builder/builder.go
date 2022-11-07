package builder

import (
	"encoding/json"
	"github.com/go-playground/validator/v10"
	"github.com/nambroa/interview-accountapi/internal/models"
)

// AccountBuilder represents a builder that builds Accounts. It also contains validations for specific fields to stop
// the creation of invalid accounts before it happens.
// The definition of a valid account is taken from the restrictions of the documentation and from the swagger API info:
// https://www.api-docs.form3.tech/api/schemes/fps-direct/accounts/accounts/create-an-account
// Examples: ID must be UUID, BankID must be max length 11, Bic must be 8 or 11 chars longs, Country must be 2 chars long, etc.
// This means that the source of truth for what an account is, is the service containing the API.
// I assume this service is internal to us and thus its constraints must be enforced and tested.
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
	// array [3] of string [140]
	ab.account.Data.Attributes.AlternativeNames = alternativeNames
	return ab
}

func (ab *AccountBuilder) WithBaseCurrency(baseCurrency string) *AccountBuilder {
	ab.account.Data.Attributes.BaseCurrency = baseCurrency // ISO 4217 code  used to identify the base currency of the account. Must be GBP.
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

func (ab *AccountBuilder) Build() (*models.Account, error) {
	validate := validator.New()
	err := validate.Struct(ab.account)
	if err != nil {
		return nil, err
	}

	return ab.account, nil
}

func FromJSON(accountJSON []byte) (*AccountBuilder, error) {
	var account models.Account
	err := json.Unmarshal(accountJSON, &account)
	if err != nil {
		return nil, err
	}
	return &AccountBuilder{account: &account}, nil
}
