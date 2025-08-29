package main

import (
	"database/sql"
	"fmt"
	"log"
	"os"
	"time"

	"github.com/google/uuid"
	_ "github.com/lib/pq"
)

func main() {
	dsn := os.Getenv("POSTGRES_DSN")
	if dsn == "" {
		log.Fatal("POSTGRES_DSN is empty")
	}

	db, err := sql.Open("postgres", dsn)
	if err != nil {
		log.Fatalf("open db: %v", err)
	}
	defer db.Close()

	id := uuid.New().String()
	now := time.Now().UTC()

	const q = `
INSERT INTO legal_entities (uuid, name, company_uuid, created_at, updated_at, deleted_at, bank_accounts)
VALUES ($1, $2, NULL, $3, $4, NULL, '[]')`

	if _, err := db.Exec(q, id, "LE Demo", now, now); err != nil {
		log.Fatalf("insert legal_entity: %v", err)
	}

	fmt.Print(id)
}
