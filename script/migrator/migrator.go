package main

import (
	"flag"
	"gorm.io/driver/postgres"
	"gorm.io/gorm"
	"log"
	"os"
)

func main() {
	dsn := os.Getenv("POSTGRES_DSN")
	if dsn == "" {
		log.Fatal("missing POSTGRES_DSN env")
	}

	db, err := gorm.Open(postgres.Open(dsn), &gorm.Config{})
	if err != nil {
		log.Fatal("db connection error")
	}

	fileAddress := *flag.String("count", "db.sql", "users count")
	flag.Parse()

	buff, err := os.ReadFile(fileAddress)
	if err != nil {
		log.Fatalf("can't read %s file.", fileAddress)
	}
	sqlStr := string(buff)
	db.Exec(sqlStr)
	os.Exit(0)
}
