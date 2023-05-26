package handlers

import (
	"github.com/kataras/iris/v12"
    r "github.com/recurly/recurly-client-go/v4"

    "server/recurly"
    "server/types"
)

func GetSubscription(ctx iris.Context) {
    subscriptionId := ctx.Params().Get("id")
    if subscriptionId == "" {
        ctx.StopWithProblem(iris.StatusFailedDependency, iris.NewProblem().
            Title("Missing or invalid subscription id"))
        return
    }

    sub, err := recurly.Client().GetSubscription(subscriptionId)
    if err != nil {
        ctx.StopWithProblem(iris.StatusInternalServerError, iris.NewProblem().
            Title("Error retreiving subscription").DetailErr(err))
        return
    }

    // optionally: get subscription from db

    ctx.JSON(sub)
}

func CreateSubscription(ctx iris.Context) {
    var reqBody types.CreateSubBody
    err := ctx.ReadJSON(&reqBody)
    if err != nil {
        ctx.StopWithProblem(iris.StatusBadRequest, iris.NewProblem().
            Title("Missing or invalid subscription data").DetailErr(err))
        return
    }

    createSub := r.SubscriptionCreate{
        Currency: &reqBody.Currency,
        Account: &r.AccountCreate{
            Code: &reqBody.AccountCode,
        },
        PlanCode: &reqBody.PlanCode,
    }

    sub, err := recurly.Client().CreateSubscription(&createSub)
    if err != nil {
        ctx.StopWithProblem(iris.StatusInternalServerError, iris.NewProblem().
            Title("Error creating subscription").DetailErr(err))
        return
    }

    // TODO: Add to DB

    ctx.JSON(sub)
}

func UpdateSubscription(ctx iris.Context) {
    var reqBody types.UpdateSubBody
    err := ctx.ReadJSON(&reqBody)
    if err != nil {
        ctx.StopWithProblem(iris.StatusBadRequest, iris.NewProblem().
            Title("Missing or invalid account data").DetailErr(err))
        return
    }

    updateSub := r.SubscriptionUpdate{
        // TODO: finish subscription update
    }

    acct, err := recurly.Client().UpdateSubscription(reqBody.Id, &updateSub)
    if err != nil {
        ctx.StopWithProblem(iris.StatusInternalServerError, iris.NewProblem().
            Title("Error updating account").DetailErr(err))
        return
    }

    // TODO: Add to DB

    ctx.JSON(acct)
}

func CancelSubscription(ctx iris.Context) {
    subId := ctx.Params().Get("id")
    if subId == "" {
        ctx.StopWithProblem(iris.StatusFailedDependency, iris.NewProblem().
            Title("Missing or invalid subscription id"))
        return
    }

    // cancels at the bill date
    timeframe := "bill_date"
    params := r.CancelSubscriptionParams{
        Body: &r.SubscriptionCancel{
            Timeframe: &timeframe,
        },
    }

    sub, err := recurly.Client().CancelSubscription(subId, &params)
    if err != nil {
        ctx.StopWithProblem(iris.StatusInternalServerError, iris.NewProblem().
            Title("Error canceling subscription").DetailErr(err))
        return
    }

    ctx.JSON(sub)
}

func PauseSubscription(ctx iris.Context) {
    subId := ctx.Params().Get("id")
    if subId == "" {
        ctx.StopWithProblem(iris.StatusFailedDependency, iris.NewProblem().
            Title("Missing or invalid subscription id"))
        return
    }

    // amount of cycles to pause for
    cycles := 1
    subPause := r.SubscriptionPause{
        RemainingPauseCycles: &cycles,
    }

    sub, err := recurly.Client().PauseSubscription(subId, &subPause)
    if err != nil {
        ctx.StopWithProblem(iris.StatusInternalServerError, iris.NewProblem().
            Title("Error pausing subscription").DetailErr(err))
        return
    }

    ctx.JSON(sub)
}
