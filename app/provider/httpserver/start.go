package httpserver

import (
	"chenwlnote.gin-api/app/provider/app/config"
	"context"
	"fmt"
	"github.com/facebookgo/grace/gracehttp"
	"golang.org/x/sync/errgroup"
	"log"
	"net/http"
	"os"
	"os/signal"
	"strconv"
	"syscall"
	"time"
)

var (
	server *http.Server
	g      errgroup.Group
)

func init() {
	fmt.Println("server start init")
}

func Start() {
	server = &http.Server{
		Addr:           config.Get().HttpServer.Addr,
		Handler:        r,
		ReadTimeout:    2 * time.Second,
		WriteTimeout:   2 * time.Second,
		MaxHeaderBytes: 1 << 20,
	}

	g.Go(func() error {
		err := gracehttp.Serve(server)
		if err != nil && err != http.ErrServerClosed {
			log.Fatalln(err)
		}
		return err
	})

	g.Go(func() error {
		err := savePid()
		if err != nil && err != http.ErrServerClosed {
			log.Fatalln(err)
		}
		return err
	})

	g.Go(func() error {
		quit := make(chan os.Signal)
		// kill (no param) default send syscall.SIGTERM
		// kill -2 is syscall.SIGINT
		// kill -9 is syscall.SIGKILL but can't be caught, so don't need to add it
		signal.Notify(quit, syscall.SIGINT, syscall.SIGTERM, syscall.SIGKILL)
		<-quit
		log.Println("Shutting down server...")

		// The context is used to inform the server it has 5 seconds to finish
		// the request it is currently handling
		ctx, cancel := context.WithTimeout(context.Background(), 3*time.Second)
		defer cancel()
		err := server.Shutdown(ctx)
		if err != nil {
			log.Fatal("Server forced to shutdown:", err)
		}
		return err
	})

	log.Println("Server exiting")

	if err := g.Wait(); err != nil {
		log.Fatal(err)
	}
}

func savePid() error {
	pid := os.Getpid()
	file, err := os.OpenFile(config.Get().HttpServer.HttpPidFile, os.O_WRONLY|os.O_CREATE, 0666)
	if err != nil {
		return err
	}
	file.Write([]byte(strconv.Itoa(pid)))
	defer file.Close()
	return nil
}
