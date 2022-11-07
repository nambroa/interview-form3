// Account represents an account in the form3 org section.
// See https://api-docs.form3.tech/api.html#organisation-accounts for
// more information about fields.
package models

type AccountStatus string

const (
	PENDING   AccountStatus = "pending"
	FAILED    AccountStatus = "faied"
	CONFIRMED AccountStatus = "confirmed"
	CLOSED    AccountStatus = "closed"
)

type AccountType string

const (
	ACCOUNTS AccountType = "accounts"
)

type AccountClassification string

const (
	PERSONAL AccountClassification = "Personal"
	BUSINESS AccountClassification = "Business"
)

type NameMatchingStatus string

const (
	SUPPORTED     NameMatchingStatus = "supported"
	NOT_SUPPORTED NameMatchingStatus = "not_supported"
	OPTED_OUT     NameMatchingStatus = "opted_out"
	SWITCHED      NameMatchingStatus = "switched"
)

type Account struct {
	Data *AccountData `json:"data,omitempty"`
}

type AccountData struct {
	Attributes     *AccountAttributes `json:"attributes,omitempty"`
	ID             string             `json:"id,omitempty" validate:"required,uuid"`
	OrganisationID string             `json:"organisation_id,omitempty" validate:"required,uuid"`
	Type           AccountType        `json:"type,omitempty" validate:"required"`
	Version        *int64             `json:"version,omitempty" validate:"min=0"`
}

type AccountAttributes struct {
	AccountClassification   *AccountClassification `json:"account_classification,omitempty"`
	AccountNumber           string                 `json:"account_number,omitempty" validate:"omitempty,alphanum,max=64"`
	AlternativeNames        []string               `json:"alternative_names,omitempty" validate:"max=3,dive,min=1,max=140"`
	BankID                  string                 `json:"bank_id,omitempty" validate:"required,max=11,alphanum"`
	BankIDCode              string                 `json:"bank_id_code,omitempty" validate:"required,alphanum,max=16"`
	BaseCurrency            string                 `json:"base_currency,omitempty" validate:"iso4217"`
	Bic                     string                 `json:"bic,omitempty" validate:"required,alphanum,len=8|len=11"`
	Country                 *string                `json:"country,omitempty" validate:"required,iso3166_1_alpha2,len=2"`
	Iban                    string                 `json:"iban,omitempty"`
	JointAccount            *bool                  `json:"joint_account,omitempty"`
	Name                    []string               `json:"name,omitempty" validate:"required,max=4,dive,min=1,max=140"`
	NameMatchingStatus      *NameMatchingStatus    `json:"name_matching_status,omitempty"` // Changed from AccountMatchingOptOut+Switched since it's replaced in the docs (deprecation).
	SecondaryIdentification string                 `json:"secondary_identification,omitempty" validate:"max=140"`
	Status                  *AccountStatus         `json:"status,omitempty"`
}
