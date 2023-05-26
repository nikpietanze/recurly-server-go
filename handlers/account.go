package handlers

import (
	"github.com/kataras/iris/v12"
    r "github.com/recurly/recurly-client-go/v4"

    "server/recurly"
    "server/types"
)

func GetAccount(ctx iris.Context) {
    accountId := ctx.Params().Get("id")
    if accountId == "" {
        ctx.StopWithProblem(iris.StatusBadRequest, iris.NewProblem().
            Title("Missing or invalid account id"))
        return
    }

    acct, err := recurly.Client().GetAccount(accountId)
    if err != nil {
        ctx.StopWithProblem(iris.StatusInternalServerError, iris.NewProblem().
            Title("Error retreiving account").DetailErr(err))
        return
    }

    ctx.JSON(acct)
}

func CreateAccount(ctx iris.Context) {
    var reqBody types.CreateAccountBody
    err := ctx.ReadJSON(&reqBody)
    if err != nil {
        ctx.StopWithProblem(iris.StatusBadRequest, iris.NewProblem().
            Title("Missing or invalid account data").DetailErr(err))
        return
    }

    // set preferred timezone
    // based on user location
    timezone := "Etc/UTC"

    createAcct := r.AccountCreate {
        Code: &reqBody.AccountCode,
        Email: &reqBody.Email,
        FirstName: &reqBody.FirstName,
        LastName: &reqBody.LastName,
        PreferredTimeZone: &timezone,
        Address: &r.AddressCreate{
            Country: &reqBody.Country,
            PostalCode: &reqBody.Zip,
        },
        BillingInfo: &r.BillingInfoCreate{
            TokenId: &reqBody.TokenId,
        },
    }

    for i := 0; i < len(reqBody.CustomFields); i++ {
        customField := reqBody.CustomFields[i]

        if (customField.Name == "my_custom_field") {
            newField := r.CustomFieldCreate {
                Name: &customField.Name,
                Value: &customField.Value,

            }
            createAcct.CustomFields = append(
                createAcct.CustomFields,
                newField,
            )
        }
    }

    acct, err := recurly.Client().CreateAccount(&createAcct)
    if err != nil {
        ctx.StopWithProblem(iris.StatusInternalServerError, iris.NewProblem().
            Title("Error creating account").DetailErr(err))
        return
    }

    // TODO: Add to DB

    ctx.JSON(acct)
}

func UpdateAccount(ctx iris.Context) {
    var reqBody types.UpdateAccountBody
    err := ctx.ReadJSON(&reqBody)
    if err != nil {
        ctx.StopWithProblem(iris.StatusBadRequest, iris.NewProblem().
            Title("Missing or invalid account data").DetailErr(err))
        return
    }

    var updateAcct r.AccountUpdate
    updateAcct.Email = &reqBody.Email
    updateAcct.FirstName = &reqBody.FirstName
    updateAcct.LastName = &reqBody.LastName

    acct, err := recurly.Client().UpdateAccount(reqBody.Id, &updateAcct)
    if err != nil {
        ctx.StopWithProblem(iris.StatusInternalServerError, iris.NewProblem().
            Title("Error updating account").DetailErr(err))
        return
    }

    // TODO: Update in DB

    ctx.JSON(acct)
}

func DeleteAccount(ctx iris.Context) {
    accountId := ctx.Params().Get("id")
    if accountId == "" {
        ctx.StopWithProblem(iris.StatusBadRequest, iris.NewProblem().
            Title("Missing or invalid account id"))
        return
    }

    // TODO: Delete account from DB

    ctx.JSON(true)
}
