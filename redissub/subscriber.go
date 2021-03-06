package main

import (
"fmt"
"github.com/go-redis/redis/v8"
"context"
"log"
"time"
"os"

"go.mongodb.org/mongo-driver/bson"
"go.mongodb.org/mongo-driver/mongo"
"go.mongodb.org/mongo-driver/mongo/options"
)

var ctx = context.Background()

func main(){

	client := redis.NewClient(&redis.Options{
		Addr:     "35.238.130.224:6379",
		Password: "redissopes1",
		DB: 	0,
		})
	defer client.Close()

	sub := client.Subscribe(ctx, "canal1")
	for {
		msg, err := sub.ReceiveMessage(ctx)
		if err != nil {
			panic(err)
		}
		//fmt.Println(msg.Channel, msg.Payload)
		val, err := client.Do(ctx, "RPUSH", "listaCasos", msg.Payload).Result()
		if err != nil {
			fmt.Println("Error: ",err)
		}
		fmt.Println(val)


		// Conexion con MongoDB
		ctx, cancel := context.WithTimeout(context.Background(), 10*time.Second)
		defer cancel()
		mongoclient, err := mongo.Connect(ctx, options.Client().ApplyURI(
			os.Getenv("MongoDB"),
		))
		if err != nil {
			log.Fatal(err)
		}

		Database := mongoclient.Database("covid19")
		Collection := Database.Collection("infectados")
	
		var bdoc interface{}
		errb := bson.UnmarshalExtJSON([]byte(msg.Payload), true, &bdoc)
		if err != nil {
			fmt.Println(errb)
		}
	
		infectadosResult, err := Collection.InsertOne(ctx, bdoc)
		if err != nil {
			log.Fatal(err)
		}
		fmt.Println(infectadosResult)



	}
}