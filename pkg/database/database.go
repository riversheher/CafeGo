package database

import (
	"context"
	"log"
	"sync"

	"github.com/pocketbase/pocketbase"
)

type Database struct {
	WaitGroup *sync.WaitGroup
	Ctx       context.Context
	DB        *pocketbase.PocketBase
}

func (db *Database) Start() {
	defer db.WaitGroup.Done()

	db.DB = pocketbase.New()

	if err := db.DB.Start(); err != nil {
		log.Println(err)
	}

}
