package types

import (
    r "github.com/recurly/recurly-client-go/v4"
)

type CreatePurchaseBody struct {
    Coupon string `json:"coupon"`
    Currency string `json:"currency"`
    PlanCode string `json:"plan_code"`
    AccountCode string `json:"account_code"`
    FirstName string `json:"first_name"`
    LastName string `json:"last_name"`
    Email string `json:"email"`
    TokenId string `json:"token_id"`
    CustomFields []r.CustomField `json:"custom_fields"`
}
