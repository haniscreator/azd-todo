package main

import (
	"context"
	"log"
	"time"

	"github.com/ClickHouse/clickhouse-go/v2"
)

var ch clickhouse.Conn

// initClickhouse initializes global ClickHouse connection
func initClickhouse() {
	var err error

	ch, err = clickhouse.Open(&clickhouse.Options{
		Addr: []string{"localhost:9000"},
		Auth: clickhouse.Auth{
			Database: "default",
			Username: "default",
			Password: "secret123", // 👈 update here
		},
		DialTimeout: 5 * time.Second,
	})

	if err != nil {
		log.Fatalf("failed to open ClickHouse connection: %v", err)
	}

	// Test connection
	ctx, cancel := context.WithTimeout(context.Background(), 5*time.Second)
	defer cancel()

	if err := ch.Ping(ctx); err != nil {
		log.Fatalf("failed to ping ClickHouse: %v", err)
	}

	log.Println("✅ Connected to ClickHouse")
}
