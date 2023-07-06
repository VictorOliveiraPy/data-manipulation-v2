package handlers

import (
	"fmt"
	"github.com/Nhanderu/brdoc"
	"regexp"
	"strconv"
	"strings"
	"time"

	"github.com/VictorOliveiraPy/internal/entity"
)

func TrimValue(value string) string {
	value = strings.Trim(value, " ")
	if value == "NULL" {
		return ""
	}
	return value
}

func ParseDocumentValue(document string) (string, error) {
	if brdoc.IsCPF(document) {
		return "CPF", nil
	} else if brdoc.IsCNPJ(document) {
		return "CNPJ", nil
	}
	return "", fmt.Errorf("documento inválido")

}

func RemoveNonNumericCharacters(str string) string {
	re := regexp.MustCompile("[^0-9]+")
	cleanedStr := re.ReplaceAllString(str, "")

	return cleanedStr
}

func ParseFloat(value string) (float64, error) {
	str := strings.Replace(TrimValue(value), ",", ".", 1)
	f, err := strconv.ParseFloat(str, 64)
	if err != nil {
		return f, nil
	}
	return f, nil
}

func ParseBoolValue(value string) (bool, error) {
	parsedValue, err := strconv.ParseBool(TrimValue(value))
	if err != nil {
		return false, err
	}

	return parsedValue, nil
}

func parseDate(dateString string) (*time.Time, error) {
	if TrimValue(dateString) == "" {
		return nil, nil
	}

	date, err := time.Parse("2006-01-02", dateString)
	if err != nil {
		return &time.Time{}, err
	}

	utcDate := date.UTC()
	return &utcDate, nil
}

func PointerToString(value *string) string {
	if value == nil {
		return ""
	}
	return *value
}

func ParseClient(parserChannel chan entity.ClientRaw, parserChannelClient chan entity.Client) error {
	for c := range parserChannel {
		documentType, err := ParseDocumentValue(c.Document)

		if err != nil {
			continue
		}
		storeLastPurchase := RemoveNonNumericCharacters(c.StoreMostFrequent)

		StoreMostFrequent := RemoveNonNumericCharacters(c.StoreMostFrequent)

		averageTicket, err := ParseFloat(TrimValue(PointerToString(&c.TicketAverage)))
		if err != nil {
			return nil
		}

		lastPurchaseTicket, err := ParseFloat(TrimValue(PointerToString(&c.TicketLastPurchase)))
		if err != nil {
			return nil
		}
		isPrivate, err := ParseBoolValue(c.Private)
		if err != nil {
			return nil
		}

		isIncomplete, err := ParseBoolValue(c.Incomplete)
		if err != nil {
			return nil
		}

		lastDate, err := parseDate(PointerToString(&c.LastPurchaseDate))
		if err != nil {
			fmt.Println("Erro ao converter a string para data:", err)
			return nil
		}
		now := time.Now()

		client := &entity.Client{
			ID:                 c.ID,
			Document:           RemoveNonNumericCharacters(TrimValue(c.Document)),
			DocumentType:       documentType,
			Private:            isPrivate,
			Incomplete:         isIncomplete,
			LastPurchaseDate:   lastDate,
			TicketAverage:      averageTicket,
			TicketLastPurchase: lastPurchaseTicket,
			StoreMostFrequent:  StoreMostFrequent,
			StoreLastPurchase:  storeLastPurchase,
			CreatedAt:          &now,
			UpdatedAt:          &now,
		}
		parserChannelClient <- *client
	}
	close(parserChannelClient)
	return nil

}
