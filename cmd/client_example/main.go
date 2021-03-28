package main

import (
	"fmt"
	"log"

	"github.com/gicappa/interview-accountapi/account"
)

func main() {
	client := account.NewClient("http://localhost:8080", "634e3a41-26b8-49f9-a23d-26fa92061f38")

	account, err := client.CreateEx(&account.Account{
		Country: "IT",
	})
	if err != nil {
		log.Fatalf("ERROR:%+v", err)
		log.Fatal("Can't create a new account")
	}

	fmt.Printf("%+v\n", account)

}
