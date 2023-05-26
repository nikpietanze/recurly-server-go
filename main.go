package main

import (
	"github.com/kataras/iris/v12"
    "server/handlers"
)

func main() {
    app := iris.New()
    app.Use(iris.Compression)

    api := app.Party("/api")
    {
        account := api.Party("/account")
        {
            account.Get("/{id}", handlers.GetAccount)
            account.Post("/", handlers.CreateAccount)
            account.Put("/", handlers.UpdateAccount)
            account.Delete("/{id}", handlers.DeleteAccount)
        }

        invoice := api.Party("/invoice")
        {
            invoice.Get("/", handlers.GetInvoice)
        }

        sub := api.Party("/subscription")
        {
            sub.Get("/", handlers.GetSubscription)
            sub.Post("/", handlers.CreateSubscription)
            sub.Put("/", handlers.UpdateSubscription)
            sub.Get("/cancel/{id}", handlers.CancelSubscription)
            sub.Get("/pause/{id}", handlers.PauseSubscription)
        }

        purchase := api.Party("/purchase")
        {
            purchase.Post("/", handlers.CreatePurchase)
        }
    }

    app.Listen(":8080")
}

