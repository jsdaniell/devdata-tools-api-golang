package db

import (
	"cloud.google.com/go/firestore"
	firebase "firebase.google.com/go"
	"golang.org/x/net/context"
	"google.golang.org/api/option"
	"log"
)

func FirestoreClient() *firestore.Client{
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
}