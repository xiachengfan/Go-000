package main

import (
	"context"
	"fmt"
	"net/http"
	"os"
	"os/signal"
	"syscall"
	"time"

	"golang.org/x/sync/errgroup"
)

func main() {
	ctx := context.Background()
	g, cancelCtx := errgroup.WithContext(ctx)
	g.Go(func() error {
		return hookSignal(cancelCtx)
	})
	g.Go(func() error {
		return startServer(cancelCtx, ":8086", &httpHandler{})
	})
	if err := g.Wait(); err != nil {
		fmt.Println("error group return err:", err.Error())
	}

	fmt.Println("Shutdown!!!")
}

func hookSignal(ctx context.Context) error {
	c := make(chan os.Signal, 1)
	signal.Notify(c, syscall.SIGHUP, syscall.SIGQUIT, syscall.SIGTERM, syscall.SIGINT)
	fmt.Println("signal routine：START!")
	for {
		s := <-c
		fmt.Printf("get a signal %s", s.String())
		switch s {
		case syscall.SIGQUIT, syscall.SIGTERM, syscall.SIGINT:
			time.Sleep(time.Second)
			return fmt.Errorf("quit !!!")
		case syscall.SIGHUP:
		default:
			return ctx.Err()
		}
	}
}

func startServer(ctx context.Context, addr string, h http.Handler) error {
	s := http.Server{
		Addr:    addr,
		Handler: h,
	}

	go func(ctx context.Context) {
		<-ctx.Done()
		fmt.Println("http server s% ctx done", s.Addr)
		if err := s.Shutdown(context.Background()); err != nil {
			fmt.Println("http server %s shutdown err : s%", s.Addr, err)
		}
	}(ctx)
	fmt.Println("http routione：START!")
	return s.ListenAndServe()
}

type httpHandler struct {
}

func (h *httpHandler) ServeHTTP(respWriter http.ResponseWriter, req *http.Request) {

}
