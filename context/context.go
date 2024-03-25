package main

import (
	"context"
	"fmt"
	"log"
	"log/slog"
	"net/http"
	"time"
)

type Store interface {
	Fetch(ctx context.Context) (string, error)
}

func Server(store Store) http.HandlerFunc {
	return func(w http.ResponseWriter, r *http.Request) {
		data, err := store.Fetch(r.Context())

		if err != nil {
			slog.Error(err.Error())
			return
		}

		fmt.Fprint(w, data)
	}
}

func main() {
	cancelCtx, cancel := context.WithCancel(context.Background())
	time.AfterFunc(1*time.Hour, cancel) // Change this time if you want to simulate timeout or not

	res, err := Fetch(cancelCtx)
	fmt.Printf("Res: %+v, Error: %+v\n", res, err)
}

func Fetch(ctx context.Context) (string, error) {
	data := make(chan string, 1)
	response := "hello, world"

	go func() {
		var result string
		for _, c := range response {
			select {
			case <-ctx.Done():
				log.Println("spy store got cancelled")
				return
			default:
				time.Sleep(10 * time.Millisecond)
				result += string(c)
			}
		}
		data <- result
	}()

	select {
	case <-ctx.Done():
		return "", ctx.Err()
	case res := <-data:
		return res, nil
	}
}
