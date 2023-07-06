package main

import (
	"context"
	"fmt"
	"github.com/VictorOliveiraPy/internal/entity"
	"github.com/VictorOliveiraPy/internal/handlers"
	"github.com/VictorOliveiraPy/internal/repository"
	"github.com/jackc/pgx/v5/pgxpool"
	"os"
	"sync"
	"time"
)

func main() {
	startTime := time.Now()

	conn, err := pgxpool.New(context.Background(), os.Getenv("DATABASE_URL"))
	if err != nil {
		panic(err)
	}

	rawChannel := make(chan string, handlers.BufferSize)
	rawCreateChannel := make(chan entity.ClientRaw, handlers.BufferSize)
	parserChannel := make(chan entity.ClientRaw, handlers.BufferSize)
	parserChannelClient := make(chan entity.Client, handlers.BufferSize)
	channelUpdateClientRaw := make(chan entity.Client, handlers.BufferSize)
	var wg sync.WaitGroup

	handlerRepository := repository.NewClientRawRepository(conn)

	go handlers.ReadFileAndSendToChannel("base.txt", rawChannel)
	wg.Add(1)

	go handlers.HandleRawClientData(rawChannel, rawCreateChannel, &wg)
	wg.Add(1)

	go handlerRepository.CreateRaw(rawCreateChannel, &wg)
	wg.Add(1)

	go handlerRepository.GetClients(10000, "Waiting", parserChannel, &wg)
	wg.Add(1)

	go handlers.ParseClient(parserChannel, parserChannelClient, &wg)
	wg.Add(1)

	go handlerRepository.Create(parserChannelClient, channelUpdateClientRaw, &wg)
	wg.Add(1)

	go handlerRepository.UpdateStatusClient(channelUpdateClientRaw, &wg)

	defer wg.Wait()
	elapsed := time.Since(startTime)
	fmt.Println("[Done] exited with code=0 in %s", elapsed)
}
