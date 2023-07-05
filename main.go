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
	rawCreateChannel := make(chan entity.Client, handlers.BufferSize)
	var wg sync.WaitGroup

	go handlers.ReadFileAndSendToChannel("base.txt", rawChannel)
	wg.Add(1)

	go handlers.HandleRawClientData(rawChannel, rawCreateChannel, &wg)
	handlerRepository := repository.NewClientRawRepository(conn)
	wg.Add(1)
	go handlerRepository.Create(rawCreateChannel, &wg)

	defer wg.Wait()
	elapsed := time.Since(startTime)
	fmt.Println("[Done] exited with code=0 in %s", elapsed)
}
