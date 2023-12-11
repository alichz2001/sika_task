package main

import (
	"flag"
	"github.com/google/uuid"
	"gorm.io/gorm"
	"log"
	"os"
)
import "gorm.io/driver/postgres"

var (
	dbConn *gorm.DB
	count  int
)

func main() {

	count = *flag.Int("count", 1000000, "users count")

	flag.Parse()

	dsn := os.Getenv("POSTGRES_DSN")
	if dsn == "" {
		log.Fatal("missing POSTGRES_DSN env")
	}

	db, err := gorm.Open(postgres.Open(dsn), &gorm.Config{})
	if err != nil {
		log.Fatal("db connection error")
	}
	dbConn = db

	run()
	os.Exit(0)
}

type User struct {
	ID        uint
	FirstName string
	LastName  string
}

func run() {

	users := make([]*User, 1000)

	i, j := 0, 0

outer:
	for {
		j = 0
	inner:
		for {
			users[j] = &User{
				FirstName: uuid.New().String(),
				LastName:  uuid.New().String(),
			}
			j++
			i++
			if i == count || j == 1000 {
				break inner
			}
		}
		dbConn.Create(users)
		log.Printf("%d users inserted to DB.", j)
		if i == count {
			break outer
		}
	}

	log.Println("users generated. will insert to db.")

}
