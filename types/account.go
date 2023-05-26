package types

import (
    r "github.com/recurly/recurly-client-go/v4"
)

type CreateAccountBody struct {
    AccountCode string `json:"account_code"`
    TokenId string `json:"token_id"`
    FirstName string `json:"first_name"`
    LastName string `json:"last_name"`
    Email string `json:"email"`
    Country string `json:"country"`
    Zip string `json:"zip"`
    CustomFields []r.CustomField `json:"custom_fields"`
}

type UpdateAccountBody struct {
    Id string `json:"account_id"`
    FirstName string `json:"first_name"`
    LastName string `json:"last_name"`
    Email string `json:"email"`
}
