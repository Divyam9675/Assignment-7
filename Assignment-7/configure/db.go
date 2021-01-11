package configure

import (
	"log"
	"os"

	//"github.com/go-pg/pg/v10"
	"github.com/go-pg/pg"

	controllers "Assignment-7/controllers"
)

// Connecting to db

func Connect() *pg.DB {
	
	opts := &pg.Options{
	
		User:     "postgres",
		Password: "admin",
		Addr:     "localhost:5432",
		Database: "golang",
	
	}

	var db *pg.DB = pg.Connect(opts)
	if db == nil {
		log.Printf("Failed to connect")
		os.Exit(100)
	}
	log.Printf("Connected to db")
	controllers.CreateTodoTable(db)
	controllers.InitiateDB(db)
	return db
}
