package main

import (
	"context"
	"flag"
	"fmt"
	"log"
	"net"
	"os"
	"os/signal"
	"syscall"

	"github.com/sapawarga/phonebook-service/cmd/database"
	"github.com/sapawarga/phonebook-service/config"
	"github.com/sapawarga/phonebook-service/repository/mysql"
	transportGRPC "github.com/sapawarga/phonebook-service/transport/grpc"
	"github.com/sapawarga/phonebook-service/usecase"
	"github.com/sapawarga/proto-file/phonebook"

	kitlog "github.com/go-kit/kit/log"
	"google.golang.org/grpc"
)

var (
	filename = "cmd/grpc/main.go"
	method   = "main"
)

func main() {
	config, err := config.NewConfig()
	errorCheck(err)

	ctx := context.Background()
	db := database.NewConnection(config.DB)
	errChan := make(chan error)

	// setting repository
	repo := mysql.NewPhonebookRepository(db)

	logger := kitlog.NewLogfmtLogger(kitlog.NewSyncWriter(os.Stderr))
	logger = kitlog.With(logger, "ts", kitlog.DefaultTimestampUTC)
	uc := usecase.NewPhoneBook(repo, logger)

	// Initialize grpc
	grpcAdd := flag.String("grpc", fmt.Sprintf(":%d", config.AppPort), "gRPC listening address")
	go func() {
		logger.Log("transport", "grpc", "address", *grpcAdd, "msg", "listening")
		listener, err := net.Listen("tcp", *grpcAdd)
		if err != nil {
			errChan <- err
			return
		}
		handler := transportGRPC.MakeHandler(ctx, uc)
		grpcServer := grpc.NewServer()
		phonebook.RegisterPhoneBookHandlerServer(grpcServer, handler)
		logger.Log(
			"filename", filename,
			"method", method,
			"note", "running phonebook service grpc",
		)
		errChan <- grpcServer.Serve(listener)
	}()
	go func() {
		c := make(chan os.Signal, 1)
		signal.Notify(c, syscall.SIGINT, syscall.SIGTERM, os.Interrupt)
		errChan <- fmt.Errorf("%s", <-c)
		logger.Log(
			"filename", filename,
			"method", method,
			"note", "Gracefully Stop Trading Account GRPC",
		)
	}()
	logger.Log(
		"filename", filename,
		"method", method,
		"note", <-errChan,
	)
}

func errorCheck(err error) {
	if err != nil {
		log.Fatal(err)
	}
}
