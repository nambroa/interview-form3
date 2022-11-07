package builder

import (
	"github.com/nambroa/interview-accountapi/internal/models"
	uuid "github.com/nu7hatch/gouuid"
	"github.com/stretchr/testify/assert"
	"testing"
)

func TestAccountBuilder_ValidBasicAccount(t *testing.T) {
	// Basic Builder fields
	ID, _ := uuid.NewV4()
	OrganisationID, _ := uuid.NewV4()
	var names []string
	names = append(names, "Batman")
	names = append(names, "Joker")
	names = append(names, "Robin")
	var bankID = "400300"
	var bankIDCode = "GBDSC"
	var BIC = "NWBKGB22"
	var country = "GB"

	// Basic Builder
	var accountBuilder = NewAccountBuilder(ID.String(), OrganisationID.String(), bankID, bankIDCode,
		BIC, country, names)
	// Create Account
	account, err := accountBuilder.Build()
	// Validate no errors thrown
	assert.Nil(t, err)

	// Validate account fields
	assert.Equal(t, account.Data.ID, ID.String())
	assert.Equal(t, account.Data.OrganisationID, OrganisationID.String())
	assert.Equal(t, account.Data.Attributes.BankID, bankID)
	assert.Equal(t, account.Data.Attributes.BankIDCode, bankIDCode)
	assert.Equal(t, account.Data.Attributes.Bic, BIC)
	assert.Equal(t, account.Data.Attributes.Country, &country)
}

func TestAccountBuilder_ValidFullAccount(t *testing.T) {
	// Basic Builder fields
	ID, _ := uuid.NewV4()
	OrganisationID, _ := uuid.NewV4()
	var names []string
	names = append(names, "Batman")
	names = append(names, "Joker")
	names = append(names, "Robin")
	var bankID = "400300"
	var bankIDCode = "GBDSC"
	var BIC = "NWBKGB22"
	var country = "GB"
	// Basic Builder
	var accountBuilder = NewAccountBuilder(ID.String(), OrganisationID.String(), bankID, bankIDCode,
		BIC, country, names)

	// Optional Fields
	var altNames []string
	altNames = append(altNames, "Bruce")
	altNames = append(altNames, "Jack")
	altNames = append(altNames, "Jason")
	var jointAcc = false
	var accClass = models.BUSINESS
	var NMS = models.OPTED_OUT
	var version int64 = 1
	var iban = "AB12AZ33"
	var accNumber = "1234567890"
	var secondaryId = "Alfred"
	var pesoCurr = "ARS"

	accountBuilder.WithAlternativeNames(altNames)
	accountBuilder.WithIban(iban)
	accountBuilder.WithSecondaryIdentification(secondaryId)
	accountBuilder.WithJointAccount(&jointAcc)
	accountBuilder.WithAccountClassification(&accClass)
	accountBuilder.WithBaseCurrency(pesoCurr)
	accountBuilder.WithAccountNumber(accNumber)
	accountBuilder.WithNameMatchingStatus(&NMS)
	accountBuilder.WithVersion(&version)
	// Create Account
	account, err := accountBuilder.Build()
	// Validate no errors thrown
	assert.Nil(t, err)

	// Validate account fields
	assert.Equal(t, account.Data.ID, ID.String())
	assert.Equal(t, account.Data.OrganisationID, OrganisationID.String())
	assert.Equal(t, account.Data.Attributes.BankID, bankID)
	assert.Equal(t, account.Data.Attributes.BankIDCode, bankIDCode)
	assert.Equal(t, account.Data.Attributes.Bic, BIC)
	assert.Equal(t, account.Data.Attributes.Country, &country)
	assert.Equal(t, account.Data.Attributes.AlternativeNames, altNames)
	assert.Equal(t, account.Data.Attributes.Iban, iban)
	assert.Equal(t, account.Data.Attributes.SecondaryIdentification, secondaryId)
	assert.Equal(t, account.Data.Attributes.JointAccount, &jointAcc)
	assert.Equal(t, account.Data.Attributes.AccountClassification, &accClass)
	assert.Equal(t, account.Data.Attributes.BaseCurrency, pesoCurr)
	assert.Equal(t, account.Data.Attributes.AccountNumber, accNumber)
	assert.Equal(t, account.Data.Attributes.NameMatchingStatus, &NMS)
	assert.Equal(t, account.Data.Version, &version)
}

func TestAccountBuilder_WithIDNotUUIDReturnsValidationError(t *testing.T) {
	// Basic Builder fields
	ID := "not-uuid"
	OrganisationID, _ := uuid.NewV4()
	var names []string
	names = append(names, "Batman")

	// Basic Builder
	var accountBuilder = NewAccountBuilder(ID, OrganisationID.String(), "400300", "GBDSC",
		"NWBKGB22", "GB", names)
	// Create Account
	_, err := accountBuilder.Build()
	// Validate error thrown
	assert.NotNil(t, err)
}

func TestAccountBuilder_WithOrgIDNotUUIDReturnsValidationError(t *testing.T) {
	// Basic Builder fields
	ID, _ := uuid.NewV4()
	OrganisationID := "not-uuid"
	var names []string
	names = append(names, "Batman")

	// Basic Builder
	var accountBuilder = NewAccountBuilder(ID.String(), OrganisationID, "400300", "GBDSC",
		"NWBKGB22", "GB", names)
	// Create Account
	_, err := accountBuilder.Build()

	// Validate error thrown
	assert.NotNil(t, err)
}

func TestAccountBuilder_WithBankIDNotAlphanumReturnsValidationError(t *testing.T) {
	// Basic Builder fields
	ID, _ := uuid.NewV4()
	OrganisationID, _ := uuid.NewV4()
	var names []string
	names = append(names, "Batman")

	// Basic Builder
	var accountBuilder = NewAccountBuilder(ID.String(), OrganisationID.String(), "4003222AA!!00", "GBDSC",
		"NWBKGB22", "GB", names)
	// Create Account
	_, err := accountBuilder.Build()

	// Validate error thrown
	assert.NotNil(t, err)
}

func TestAccountBuilder_WithBankIDLongerThan16CharsReturnsValidationError(t *testing.T) {
	// Basic Builder fields
	ID, _ := uuid.NewV4()
	OrganisationID, _ := uuid.NewV4()
	var names []string
	names = append(names, "Batman")

	// Basic Builder
	var accountBuilder = NewAccountBuilder(ID.String(), OrganisationID.String(), "4003023132312313123130", "#*$@#*&$@#*&$&",
		"NWBKGB22", "GB", names)
	// Create Account
	_, err := accountBuilder.Build()

	// Validate error thrown
	assert.NotNil(t, err)
}

func TestBankIdCodeNotAlphanumInAccountShouldReturnValidationError(t *testing.T) {
	// Basic Builder fields
	ID, _ := uuid.NewV4()
	OrganisationID, _ := uuid.NewV4()
	var names []string
	names = append(names, "Batman")

	// Basic Builder
	var accountBuilder = NewAccountBuilder(ID.String(), OrganisationID.String(), "400300", "#*$@#*&$@#*&$&",
		"NWBKGB22", "GB", names)
	// Create Account
	_, err := accountBuilder.Build()

	// Validate error thrown
	assert.NotNil(t, err)
}

func TestAccountBuilder_WithBicLengthNot8Nor11ReturnsValidationError(t *testing.T) {
	// Basic Builder fields
	ID, _ := uuid.NewV4()
	OrganisationID, _ := uuid.NewV4()
	var names []string
	names = append(names, "Batman")

	// Basic Builder (BIC not length 8 nor 11)
	var accountBuilder = NewAccountBuilder(ID.String(), OrganisationID.String(), "400300", "GBDSC",
		"NWBKGB2A2", "GB", names)
	// Create Account
	_, err := accountBuilder.Build()

	// Validate error thrown
	assert.NotNil(t, err)
}

func TestAccountBuilder_WithBicLengthNotAlphanumReturnsValidationError(t *testing.T) {
	// Basic Builder fields
	ID, _ := uuid.NewV4()
	OrganisationID, _ := uuid.NewV4()
	var names []string
	names = append(names, "Batman")

	// Basic Builder (BIC not length 8 nor 11)
	var accountBuilder = NewAccountBuilder(ID.String(), OrganisationID.String(), "400300", "GBDSC",
		"NWBKGB!!", "GB", names)
	// Create Account
	_, err := accountBuilder.Build()

	// Validate error thrown
	assert.NotNil(t, err)
}

func TestAccountBuilder_WithCountryCodeNotISO3166ReturnsValidationError(t *testing.T) {
	// Basic Builder fields
	ID, _ := uuid.NewV4()
	OrganisationID, _ := uuid.NewV4()
	var names []string
	names = append(names, "Batman")

	// Basic Builder
	var accountBuilder = NewAccountBuilder(ID.String(), OrganisationID.String(), "400300", "GBDSC",
		"NWBKGB22", "GBA", names)
	// Create Account
	_, err := accountBuilder.Build()

	// Validate error thrown
	assert.NotNil(t, err)
}

func TestAccountBuilder_WithEmptyNamesReturnsValidationError(t *testing.T) {
	// Basic Builder fields
	ID, _ := uuid.NewV4()
	OrganisationID, _ := uuid.NewV4()
	var names []string

	// Basic Builder
	var accountBuilder = NewAccountBuilder(ID.String(), OrganisationID.String(), "400300", "GBDSC",
		"NWBKGB22", "GBA", names)
	// Create Account
	_, err := accountBuilder.Build()

	// Validate error thrown
	assert.NotNil(t, err)
}

func TestAccountBuilder_WithNegativeVersionNumberReturnsValidationError(t *testing.T) {
	// Basic Builder fields
	ID, _ := uuid.NewV4()
	OrganisationID, _ := uuid.NewV4()
	var names []string
	names = append(names, "Batman")

	// Basic Builder
	var accountBuilder = NewAccountBuilder(ID.String(), OrganisationID.String(), "400300", "GBDSC",
		"NWBKGB22", "GB", names)
	// Negative Version
	var version int64 = -234
	accountBuilder.WithVersion(&version)
	// Create Account
	_, err := accountBuilder.Build()

	// Validate error thrown
	assert.NotNil(t, err)
}

func TestAccountBuilder_WithAlternativeNamesLongerThan140ReturnsValidationError(t *testing.T) {
	// Basic Builder fields
	ID, _ := uuid.NewV4()
	OrganisationID, _ := uuid.NewV4()
	var names []string
	names = append(names, "Batman")

	// Basic Builder
	var accountBuilder = NewAccountBuilder(ID.String(), OrganisationID.String(), "400300", "GBDSC",
		"NWBKGB22", "GB", names)
	// Alternative names
	var altNames []string
	altNames = append(altNames, "zn32t0nw9ALziAT5fKXTfpiiJKVrmbhRkw6liWf3Akn4xQEZBwBBdhk5LdVXnx9674zQlfzGweVjhPOHwl3IiSuWcpAuChbe8ypRgssU2WexU6pUmaVZFfXZoTUtxMz33HzuIUDr6uzd5W")
	accountBuilder.WithAlternativeNames(altNames)
	// Create Account
	_, err := accountBuilder.Build()

	// Validate error thrown
	assert.NotNil(t, err)
}

func TestAccountBuilder_WithMoreThanThreeAlternativeNamesReturnsValidationError(t *testing.T) {
	// Basic Builder fields
	ID, _ := uuid.NewV4()
	OrganisationID, _ := uuid.NewV4()
	var names []string
	names = append(names, "Batman")

	// Basic Builder
	var accountBuilder = NewAccountBuilder(ID.String(), OrganisationID.String(), "400300", "GBDSC",
		"NWBKGB22", "GB", names)
	// Alternative names
	var altNames []string
	altNames = append(altNames, "Form")
	altNames = append(altNames, "Three")
	altNames = append(altNames, "Financial")
	altNames = append(altNames, "Cloud")
	accountBuilder.WithAlternativeNames(altNames)
	// Create Account
	_, err := accountBuilder.Build()

	// Validate error thrown
	assert.NotNil(t, err)
}

func TestAccountBuilder_WithInvalidBaseCurrencyReturnsValidationError(t *testing.T) {
	// Basic Builder fields
	ID, _ := uuid.NewV4()
	OrganisationID, _ := uuid.NewV4()
	var names []string
	names = append(names, "Batman")

	// Basic Builder
	var accountBuilder = NewAccountBuilder(ID.String(), OrganisationID.String(), "400300", "GBDSC",
		"NWBKGB22", "GB", names)
	accountBuilder.WithBaseCurrency("GBPDIQQ")
	// Create Account
	_, err := accountBuilder.Build()

	// Validate error thrown
	assert.NotNil(t, err)
}

func TestAccountBuilder_WithSecondaryIdentificationLongerThan140CharsReturnsValidationError(t *testing.T) {
	// Basic Builder fields
	ID, _ := uuid.NewV4()
	OrganisationID, _ := uuid.NewV4()
	var names []string
	names = append(names, "Batman")

	// Basic Builder
	var accountBuilder = NewAccountBuilder(ID.String(), OrganisationID.String(), "400300", "GBDSC",
		"NWBKGB22", "GB", names)
	accountBuilder.WithSecondaryIdentification("zn32t0nw9ALziAT5fKXTfpiiJKVrmbhRkw6liWf3Akn4xQEZBwBBdhk5LdVXnx9674zQlfzGweVjhPOHwl3IiSuWcpAuChbe8ypRgssU2WexU6pUmaVZFfXZoTUtxMz33HzuIUDr6uzd5W")
	// Create Account
	_, err := accountBuilder.Build()

	// Validate error thrown
	assert.NotNil(t, err)
}

func TestAccountBuilder_WithAccountNumberNotAlphanumReturnsValidationError(t *testing.T) {
	// Basic Builder fields
	ID, _ := uuid.NewV4()
	OrganisationID, _ := uuid.NewV4()
	var names []string
	names = append(names, "Batman")

	// Basic Builder
	var accountBuilder = NewAccountBuilder(ID.String(), OrganisationID.String(), "400300", "GBDSC",
		"NWBKGB22", "GB", names)
	accountBuilder.WithAccountNumber("*&#@!*&$!@(*$@!*$@*(@$!*")
	// Create Account
	_, err := accountBuilder.Build()

	// Validate error thrown
	assert.NotNil(t, err)
}

func TestAccountBuilder_WithAccountNumberLongerThan64ReturnsValidationError(t *testing.T) {
	// Basic Builder fields
	ID, _ := uuid.NewV4()
	OrganisationID, _ := uuid.NewV4()
	var names []string
	names = append(names, "Batman")

	// Basic Builder
	var accountBuilder = NewAccountBuilder(ID.String(), OrganisationID.String(), "400300", "GBDSC",
		"NWBKGB22", "GB", names)
	accountBuilder.WithAccountNumber("zn32t0nw9ALziAT5fKXTfpiiJKVrmbhRkw6liWf3Akn4xQEZBwBBdhk5LdVXnx9674zQlfzGweVjhPOHwl3IiSuWcpAuChbe8ypRgssU2WexU6pUmaVZFfXZoTUtxMz33HzuIUDr6uzd5W")
	// Create Account
	_, err := accountBuilder.Build()

	// Validate error thrown
	assert.NotNil(t, err)
}
