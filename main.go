package main

import (
	"fmt"
	"os"
	"tddgo/pointers"

	"github.com/pkg/errors"
)

func foo() error {
	w := pointers.NewWallet(10)
	w.Deposit(10)
	err := w.Withdraw(50)
	if err != nil {
		return errors.Wrap(err, "foo executed")
	}
	return nil
}

func main() {
	err := foo()
	if err != nil {
		fmt.Printf("%+v", err)
		os.Exit(1)
	}
}
