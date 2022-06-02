package main

import (
	"database/sql"
	"fmt"
	config "golang-grpc-server/utils"
	"log"
	"os"
)

func main() {
	var enviroment = "development"

	if len(os.Args) > 1 {
		enviroment = os.Args[1]
	}

	if err := config.ParseEnvFile(fmt.Sprintf("config.%s", enviroment), "toml", "."); err != nil {
		log.Fatal(err)
	}

	db, err := sql.Open("mysql", config.Env.GetString("database.uri"))
	if err != nil {
		log.Fatalf("db: failed to connect to database./n%s", err)
	}

	defer db.Close()

	//Repositories and Services bellow this point
	
	//Repositories and Services above this point
}