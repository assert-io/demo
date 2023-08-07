package main

import (
	"errors"
	"fmt"
	"log"
	"time"

	"github.com/google/uuid"
)

func main() {
	t := customType{
		Name:   "regular",
		Number: 1,
		Things: []thing{
			{ID: uuid.New(), Timestamp: time.Now().UTC()},
		},
	}

	v, err := doThing(t)
	if err != nil {
		log.Fatal(err)
	}

	fmt.Printf("Result: %s", v)

}

type thing struct {
	ID        uuid.UUID
	Timestamp time.Time
}

type customType struct {
	Name   string
	Number int
	Things []thing
}

func doThing(t customType) (string, error) {
	if t.Name == "regular" {
		return t.Name, nil
	}

	if t.Name == "broken" {
		return "", errors.New("broken")
	}

	IDstr := t.Things[0].ID.String()

	return IDstr, nil

}
