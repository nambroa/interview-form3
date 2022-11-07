package internal

// Improvements: make the dockerfile hostname for the fake api not hardcoded.

const BaseURL string = "http://fake-api:8080"
const V1API string = "/v1"
const AccountPrefix string = "/organisation/accounts"

const AccountURL string = BaseURL + V1API + AccountPrefix
