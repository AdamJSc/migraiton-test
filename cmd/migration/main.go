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
		log.Fatal("please set either -up or -down flags")
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

	// mongo client
	mc, err := mongo.NewClient(
		options.Client().ApplyURI(conn),
		options.Client().SetConnectTimeout(time.Duration(connTO)*time.Second),
	)
	if err != nil {
		log.Fatal(err)
	}
	if err := mc.Connect(ctx); err != nil {
		log.Fatal(err)
	}

	// mongo migration driver
	cnf := &mongodb.Config{DatabaseName: dbName}
	md, err := mongodb.WithInstance(mc, cnf)
	if err != nil {
		log.Fatal(err)
	}

	// source driver
	f := &file.File{}
	source, err := f.Open(sourcePath)
	if err != nil {
		log.Fatal(err)
	}

	// build migration client
	mig, err := migrate.NewWithInstance(sourceName, source, dbName, md)
	if err != nil {
		log.Fatal(err)
	}

	log.Println(mig)

	switch {
	case up:
		log.Println("migrate up")
	case down:
		log.Println("migrate down")
	}
}
