// package cloudflare
package main

import (
	"context"
	"fmt"
	"log"
	"os"

	"github.com/cloudflare/cloudflare-go"
	"github.com/joho/godotenv"
)

func main() {
	// Load .env file
	err := godotenv.Load()
	if err != nil {
		log.Fatal("Error loading .env file")
	}

	// Construct a new API object using a global API key
	api, err := cloudflare.New(os.Getenv("CLOUDFLARE_API_KEY"), os.Getenv("CLOUDFLARE_API_EMAIL"))

	// alternatively, you can use a scoped API token
	// api, err := cloudflare.NewWithAPIToken(os.Getenv("CLOUDFLARE_API_TOKEN"))
	if err != nil {
		log.Fatal(err)
	}

	// Most API calls require a Context
	// a Context is used to carry deadlines, cancellation signals, and other
	// request-scoped values across API boundaries and between processes.
	// A Context in Go is a struct that contains a deadline, a cancelation signal,
	// and a key-value map for request-scoped data.

	// Do not store Contexts inside a struct type; instead, pass a Context explicitly to each function that needs it | https://www.ardanlabs.com/blog/2019/09/context-package-semantics-in-go.html
	
	ctx := context.Background()
	// How To Use Contexts in Go: https://www.digitalocean.com/community/tutorials/how-to-use-contexts-in-go
	// Context in Golang [Efficient Concurrency Management]: https://medium.com/@jamal.kaksouri/the-complete-guide-to-context-in-golang-efficient-concurrency-management-43d722f6eaea


	// Fetch user details on the account
	u, err := api.UserDetails(ctx)
	if err != nil {
		log.Fatal(err)
	}
	// Print user details
	fmt.Println(u)
}
