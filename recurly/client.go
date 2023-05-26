package recurly

import (
    r "github.com/recurly/recurly-client-go/v4"
)

func Client() *r.Client {
    client, err := r.NewClient("<apikey>")
    if err != nil {
        panic(err)
    }
    return client
}
