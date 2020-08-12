package config

import (
	"github.com/joho/godotenv"
	"log"
	"os"
	"strconv"
)

var (
	PORT = 0
)

func Load(){

	var err error

	err = godotenv.Load()
	if err != nil {
		log.Fatalln(err)
	}

	PORT, err = strconv.Atoi(os.Getenv("API_PORT"))
	if err != nil {
		log.Println(err)
		PORT = 9000
	}
}


