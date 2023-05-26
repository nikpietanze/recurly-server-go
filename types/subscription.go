package types

type CreateSubBody struct {
    AccountCode string `json:"account_code"`
    PlanCode string `json:"plan_code"`
    Currency string `json:"currency"`
}

type UpdateSubBody struct {
    Id string `json:"id"`
    PlanCode string `json:"plan_code"`
    Currency string `json:"currency"`
}

type CreateSubChangeBody struct {
    SubscriptionId string `json:"subscription_id"`
    AddonCodes []string `json:"addon_codes"`
}

