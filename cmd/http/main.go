package main

import (
	"context"
	"flag"
	"log"
	"net/http"
	"os"
	"os/signal"
	"syscall"
	"time"

	"github.com/PrinceNorin/jikan/handler"
	"github.com/PrinceNorin/jikan/service"
)

var (
	fs   = flag.NewFlagSet("jikan", flag.ExitOnError)
	addr = fs.String("addr", ":8080", "HTTP server address")
)

func main() {
	if err := fs.Parse(os.Args[:1]); err != nil {
		panic(err)
	}

	svc := service.New()
	h := handler.NewHTTP(svc)
	server := &http.Server{
		Handler:      h,
		Addr:         *addr,
		ReadTimeout:  10 * time.Second,
		WriteTimeout: 10 * time.Second,
	}

	go func() {
		log.Printf("listening on: http://127.0.0.1%s\n", *addr)
		if err := server.ListenAndServe(); err != nil {
			log.Printf("server error %s\n", err.Error())
		}
	}()

	quit := make(chan os.Signal)
	signal.Notify(quit, os.Interrupt, syscall.SIGTERM)

	<-quit

	log.Println("shuting down server...")
	server.Shutdown(context.Background())
}
