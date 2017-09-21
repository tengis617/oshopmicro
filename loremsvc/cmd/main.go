package main

import (
	"context"
	"fmt"
	"net/http"
	"os"
	"os/signal"
	"syscall"

	"github.com/go-kit/kit/log"

	"github.com/tengis617/oshopmicro/loremsvc"
)

func main() {
	ctx := context.Background()
	errChan := make(chan error)

	var svc loremsvc.Service
	svc = loremsvc.LoremService{}

	endpoint := loremsvc.Endpoints{
		LoremEndpoint: loremsvc.MakeLoremEndpoint(svc),
	}

	var logger log.Logger
	{
		logger = log.NewLogfmtLogger(os.Stderr)
		logger = log.With(logger, "ts", log.DefaultTimestampUTC)
		logger = log.With(logger, "caller", log.DefaultCaller)
	}

	h := loremsvc.MakeHttpHandler(ctx, endpoint, logger)

	go func() {
		fmt.Println("Starting server at port 8080")
		errChan <- http.ListenAndServe(":8080", h)

	}()

	go func() {
		c := make(chan os.Signal, 1)
		signal.Notify(c, syscall.SIGINT, syscall.SIGTERM)
		errChan <- fmt.Errorf("%s", <-c)
	}()
	fmt.Println(<-errChan)
}
