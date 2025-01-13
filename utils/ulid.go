package utils

import (
	"github.com/oklog/ulid"
	"log"
	"math/rand"
	"time"
)

func GenerateUlid() string {
	t := time.Now()

	entropy := rand.New(rand.NewSource(t.UnixNano()))

	id, err := ulid.New(ulid.Timestamp(t), entropy)
	if err != nil {
		log.Fatalf("error creating ULID: %v", err)
	}
	return id.String()
}
