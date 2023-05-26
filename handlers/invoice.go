package handlers

import (
	"github.com/kataras/iris/v12"

    "server/recurly"
)

func GetInvoice(ctx iris.Context) {
    invoiceId := ctx.Params().Get("id")
    if invoiceId == "" {
        ctx.StopWithProblem(iris.StatusFailedDependency, iris.NewProblem().
            Title("Missing or invalid invoice id"))
        return
    }

    invoice, err := recurly.Client().GetInvoice(invoiceId)
    if err != nil {
        ctx.StopWithProblem(iris.StatusInternalServerError, iris.NewProblem().
            Title("Error retreiving invoice").DetailErr(err))
        return
    }

    ctx.JSON(invoice)
}
