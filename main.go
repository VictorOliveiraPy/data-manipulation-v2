package main

import (
	"context"
	"fmt"
	"github.com/VictorOliveiraPy/internal/entity"
	"github.com/VictorOliveiraPy/internal/handlers"
	"github.com/VictorOliveiraPy/internal/repository"
	"github.com/jackc/pgx/v5/pgxpool"
	"os"
	"time"
)

func main() {
	startTime := time.Now()

	conn, err := pgxpool.New(context.Background(), os.Getenv("DATABASE_URL"))
	if err != nil {
		panic(err)
	}

	quit := make(chan bool)
	rawChannel := make(chan string, handlers.BufferSize)
	rawCreateChannel := make(chan entity.ClientRaw, handlers.BufferSize)
	parserChannel := make(chan entity.ClientRaw, handlers.BufferSize)
	parserChannelClient := make(chan entity.Client, handlers.BufferSize)
	channelUpdateClientRaw := make(chan entity.Client, handlers.BufferSize)
	handlerRepository := repository.NewClientRawRepository(conn)

	go handlers.ReadFileAndSendToChannel("base.txt", rawChannel)
	go handlers.HandleRawClientData(rawChannel, rawCreateChannel)
	go handlerRepository.CreateRaw(rawCreateChannel)
	go handlerRepository.GetClients(10000, "Waiting", parserChannel)
	go handlers.ParseClient(parserChannel, parserChannelClient)
	go handlerRepository.Create(parserChannelClient, channelUpdateClientRaw)
	go handlerRepository.UpdateStatusClient(channelUpdateClientRaw, quit)
	<-quit
	elapsed := time.Since(startTime)
	fmt.Println("[Done] exited with code=0 in %s", elapsed)
}
