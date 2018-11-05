package main

import (
	"context"
	"log"
	"time"

	"github.com/adamsanghera/example-proj/template"
)

func main() {
	tmplt, err := template.New(&template.Config{
		Field: "non-arbitrary value",
	})
	if err != nil {
		panic(err)
	}

	go func() {
		err = tmplt.Run()
		if err != nil {
			log.Println(err)
		}
	}()

	time.Sleep(5 * time.Second)

	err = tmplt.Shutdown(context.Background())
	if err != nil {
		panic(err)
	}
}
