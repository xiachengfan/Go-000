package main

import (
	"Go-000/Week02/api"
	"Go-000/Week02/internal/model"
	"log"
	"os"
	"os/signal"
	"syscall"
	"time"
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
	//	go func() {
	//		ch := make(chan os.Signal)
	//		signal.Notify(ch, syscall.SIGTERM, syscall.SIGINT)
	//		errs <- fmt.Errorf("%v", <-ch)
	//	}()
	//
	//	log.Fatal(<-errs)
	//}

	c := make(chan os.Signal, 1)
	signal.Notify(c, syscall.SIGHUP, syscall.SIGQUIT, syscall.SIGTERM, syscall.SIGINT)
	for {
		s := <-c
		log.Printf("get a signal %s", s.String())
		switch s {
		case syscall.SIGQUIT, syscall.SIGTERM, syscall.SIGINT:
			time.Sleep(time.Second)
			log.Printf("quit !!!")
			return
		case syscall.SIGHUP:
		default:
			return
		}
	}
}
