package main

import (
	"context"
	"flag"
	"github.com/golang-migrate/migrate/v4"
	"github.com/golang-migrate/migrate/v4/database/mongodb"
	"github.com/golang-migrate/migrate/v4/source/file"
	"go.mongodb.org/mongo-driver/mongo"
	_ "go.mongodb.org/mongo-driver/mongo"
	"go.mongodb.org/mongo-driver/mongo/options"
	"log"
	"time"
)

func main() {
	// cli bootstrap
	var up, down bool
	flag.BoolVar(&up, "up", false, "bring migrations up")
	flag.BoolVar(&down, "down", false, "bring migrations down")
	flag.Parse()

	if (up && down) || (!up && !down) {
		log.Fatal("please set either -up or -down flag")
	}

	// setup
	var (
		ctx = context.Background()

		sourceName = "local data files"
		sourcePath = "file://data" // migrations file source directory

		conn   = "mongodb://testUser:testPass@localhost:27017" // mongo connection string
		connTO = 5                                             // mongo connection timeout in secs
		dbName = "myDB"                                        // mongo db name

	)

	log.Println("instantiating mongo client...")

	// mongo client
	mc, err := mongo.NewClient(
		options.Client().ApplyURI(conn),
		options.Client().SetConnectTimeout(time.Duration(connTO)*time.Second),
	)
	if err != nil {
		log.Fatalf("cannot instantiate mongo client: %s", err.Error())
	}

	log.Println("connecting mongo client...")

	if err := mc.Connect(ctx); err != nil {
		log.Fatalf("cannot connect mongo client: %s", err.Error())
	}

	log.Println("instantiating mongo driver...")

	// mongo migration driver
	cnf := &mongodb.Config{DatabaseName: dbName}
	md, err := mongodb.WithInstance(mc, cnf)
	if err != nil {
		log.Fatalf("cannot instantiate mongo driver: %s", err.Error())
	}

	log.Println("opening migrations source...")

	// source driver
	f := &file.File{}
	source, err := f.Open(sourcePath)
	if err != nil {
		log.Fatalf("cannot open migrations source %s: %s", sourcePath, err.Error())
	}

	log.Println("instantiating migration client...")

	// build migration client
	_, err = migrate.NewWithInstance(sourceName, source, dbName, md)
	if err != nil {
		log.Fatalf("cannot instantiate migration client: %s", err.Error())
	}

	switch {
	case up:
		log.Println("TODO: migrate up")
	case down:
		log.Println("TODO: migrate down")
	}

	log.Println("process complete!")
}
