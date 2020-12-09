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
		return startServer(cancelCtx, ":8089", &httpHandler{})
	})
	g.Go(func() error {
		return working(cancelCtx)
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
		select {
		case <-ctx.Done():
			return fmt.Errorf("signal routine：other work done")
		case s := <-c:
			fmt.Printf("get a signal %s", s.String())
			switch s {
			case syscall.SIGQUIT, syscall.SIGTERM, syscall.SIGINT:
				time.Sleep(5 * time.Second)
				return fmt.Errorf("quit!!!!")
			case syscall.SIGHUP:
			default:
				return ctx.Err()
			}
		}
	}
}

func startServer(ctx context.Context, addr string, h http.Handler) error {
	s := http.Server{
		Addr:    addr,
		Handler: h,
	}

	go func(ctx context.Context) {
		ctx1, cancel := context.WithTimeout(context.Background(), time.Second*5)
		defer cancel()
		<-ctx.Done()
		fmt.Printf("http server %s ctx done\n", s.Addr)
		if err := s.Shutdown(ctx1); err != nil {
			fmt.Printf("http server %s shutdown err : %s\n", s.Addr, err)
		}
	}(ctx)
	fmt.Println("http routione：START!")
	return s.ListenAndServe()
}

// Whether the worker Coroutine can exit
func working(ctx context.Context) error {
	for {
		select {
		case <-ctx.Done():
			fmt.Println("background Context cancel ")
			return ctx.Err()
		default:
			fmt.Println("working")
			time.Sleep(1 * time.Second)
		}
	}
}

type httpHandler struct {
}

func (h *httpHandler) ServeHTTP(respWriter http.ResponseWriter, req *http.Request) {

}
