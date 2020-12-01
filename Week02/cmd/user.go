package main

import (
	"Go-000/Week02/api"
	"Go-000/Week02/internal/model"
	"fmt"
	"log"
	"os"
	"os/signal"
	"syscall"
)

func main() {
	err := model.Init()
	if err != nil {
		log.Fatalf("%+v", err)
	}
	errs := make(chan error)
	r := api.NewRouter()
	go func() {
		errs <- r.Run(":8080")
	}()
	go func() {
		ch := make(chan os.Signal)
		signal.Notify(ch, syscall.SIGTERM, syscall.SIGINT)
		errs <- fmt.Errorf("%v", <-ch)
	}()

	log.Fatal(<-errs)
}
