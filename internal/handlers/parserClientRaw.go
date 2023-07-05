package handlers

import (
	"github.com/VictorOliveiraPy/internal/entity"
	"github.com/google/uuid"
	"strings"
	"sync"
	"time"
)

func HandleRawClientData(rawMsgChannel chan string, channelClientRaw chan entity.Client, wg *sync.WaitGroup) {
	for line := range rawMsgChannel {
		now := time.Now().Format(time.RFC3339)
		channelClientRaw <- entity.Client{
			ID:                 uuid.New(),
			Document:           getValue(DocumentIndexStart, DocumentIndexEnd, line),
			Private:            getValue(PrivateIndexStart, PrivateIndexEnd, line),
			Incomplete:         getValue(IncompleteIndexStart, IncompleteIndexEnd, line),
			LastPurchaseDate:   getValue(LastPurchaseDateIndexStart, LastPurchaseDateIndexEnd, line),
			TicketAverage:      getValue(TicketAverageIndexStart, TicketAverageIndexEnd, line),
			TicketLastPurchase: getValue(TicketLastPurchaseIndexStart, TicketLastPurchaseIndexEnd, line),
			StoreMostFrequent:  getValue(StoreMostFrequentIndexStart, StoreMostFrequentIndexEnd, line),
			StoreLastPurchase:  getValue(StoreLastPurchaseIndexStart, StoreLastPurchaseIndexEnd, line),
			Status:             "Waiting",
			CreatedAt:          now,
			UpdatedAt:          now,
		}
	}
	defer wg.Done()
	close(channelClientRaw)

}

func getValue(start, end int, line string) string {
	if end == -1 {
		return strings.Trim(line[start:], " ")
	}
	return strings.Trim(line[start:end], " ")
}
