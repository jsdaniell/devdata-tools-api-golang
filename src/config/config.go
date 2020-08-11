package config

import (
	"cloud.google.com/go/firestore"
	firebase "firebase.google.com/go"
	"golang.org/x/net/context"
	"google.golang.org/api/option"
	"log"
	"os"
	"strconv"
)

var (
	PORT = 0
)

func Load(){
	var err error

	PORT, err = strconv.Atoi(os.Getenv("PORT"))
	if err != nil {
		log.Println(err)
		PORT = 9000
	}
}


func InitializeFirebase() *firestore.Client{
	sa := option.WithCredentialsFile("src/main/devdatatools-firebase-adminsdk-hdrpb-fe65a9f3eb.json")

	app, err := firebase.NewApp(context.Background(), nil, sa)
	if err != nil {
		log.Fatalln(err)
	}



	client, err := app.Firestore(context.Background())
	if err != nil {
		log.Fatalln(err)
	}

	return client
	//
	//

	//
	//defer client.Close()
}