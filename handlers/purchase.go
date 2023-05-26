package handlers

import (
	"github.com/kataras/iris/v12"
	r "github.com/recurly/recurly-client-go/v4"

	"server/recurly"
	"server/types"
)

func CreatePurchase(ctx iris.Context) {
    var reqBody types.CreatePurchaseBody
    err := ctx.ReadJSON(&reqBody)
    if err != nil {
        ctx.StopWithProblem(iris.StatusBadRequest, iris.NewProblem().
            Title("Missing or invalid purchase data").DetailErr(err))
        return
    }

    createPurchase := r.PurchaseCreate{
        Account: &r.AccountPurchase{
            BillingInfo: &r.BillingInfoCreate{
                TokenId: &reqBody.TokenId,
            },
            Code: &reqBody.AccountCode,
            FirstName: &reqBody.FirstName,
            LastName: &reqBody.LastName,
        },
        CouponCodes: []string{reqBody.Coupon},
        Subscriptions: []r.SubscriptionPurchase{
            {
                PlanCode: &reqBody.PlanCode,
            },
        },
    }

    coll, err := recurly.Client().CreatePurchase(&createPurchase)
    if err != nil {
        ctx.StopWithProblem(iris.StatusInternalServerError, iris.NewProblem().
            Title("Error creating subscription").DetailErr(err))
        return
    }

    // TODO: Add to DB

    ctx.JSON(coll)
}
