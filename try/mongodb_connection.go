package main

import (
	"context"
	"fmt"
	"reflect"
	"time"

	"go.mongodb.org/mongo-driver/bson"
	"go.mongodb.org/mongo-driver/mongo"
	"go.mongodb.org/mongo-driver/mongo/options"
	"go.mongodb.org/mongo-driver/mongo/readpref"
)

type Fields struct {
	source_id   int64
	call_letter string
}

type Result2 struct {
	SourceId   string  `bson:"source_id,omitempty"`
	CallLetter string  `bson:"call_letter,omitempty"`
	ImsAliases []Alias `bson:"ims_aliases,omitempty"`
}

type Alias struct {
	ImsSvcId string `bson:"ims_svc_id,omitempty"`
	Mode     string `bson:"mode,omitempty"`
}

func main() {

	// Get Client, Context, CalcelFunc and
	// err from connect method.
	client, ctx, cancel, err := connect("mongodb://localhost:28017")
	if err != nil {
		panic(err)
	}

	// Release resource when the main
	// function is returned.
	defer close(client, ctx, cancel)

	// Ping mongoDB with Ping method
	ping(client, ctx)

	// Access a MongoDB collection through a database
	col := client.Database("channels").Collection("gn_02_14_2022")
	fmt.Println("Collection type:", reflect.TypeOf(col), "\n")

	var result Fields
	err = col.FindOne(context.TODO(), bson.D{}).Decode(&result)
	if err != nil {
		fmt.Println("FindOne: Unable to fetch the record")
		return
	}

	fmt.Println("source_id=", result.source_id, "call_letter", result.call_letter)

	/* ***************************************Find All       */
	// Call the collection's Find() method to return Cursor obj
	// with all of the col's documents
	cursor, err := col.Find(context.TODO(), bson.D{})

	// Find() method raised an error
	if err != nil {
		fmt.Println("Finding all documents ERROR:", err)
		defer cursor.Close(ctx)

		// If the API call was a success
	} else {
		// iterate over docs using Next()
		cnt := 0
		for cursor.Next(ctx) {
			cnt++

			// declare a result BSON object
			//var result bson.M
			var result Result2
			err := cursor.Decode(&result)

			// If there is a cursor.Decode error
			if err != nil {
				fmt.Println("cursor.Next() error:", err)
				break

				// If there are no cursor.Decode errors
			} else {
				fmt.Println("\nresult type:", reflect.TypeOf(result))
				fmt.Println("result:", result)
			}
		}
		fmt.Println("Total records=", cnt)
	}

}

func connect(uri string) (*mongo.Client, context.Context,
	context.CancelFunc, error) {

	// ctx will be used to set deadline for process, here
	// deadline will of 30 seconds.
	ctx, cancel := context.WithTimeout(context.Background(),
		30*time.Second)

	// mongo.Connect return mongo.Client method
	client, err := mongo.Connect(ctx, options.Client().ApplyURI(uri))
	return client, ctx, cancel, err
}

// This is a user defined method to close resources.
// This method closes mongoDB connection and cancel context.
func close(client *mongo.Client, ctx context.Context,
	cancel context.CancelFunc) {

	// CancelFunc to cancel to context
	defer cancel()

	// client provides a method to close
	// a mongoDB connection.
	defer func() {

		// client.Disconnect method also has deadline.
		// returns error if any,
		if err := client.Disconnect(ctx); err != nil {
			panic(err)
		}
	}()

}

// This is a user defined method that accepts
// mongo.Client and context.Context
// This method used to ping the mongoDB, return error if any.
func ping(client *mongo.Client, ctx context.Context) error {

	// mongo.Client has Ping to ping mongoDB, deadline of
	// the Ping method will be determined by cxt
	// Ping method return error if any occurred, then
	// the error can be handled.
	if err := client.Ping(ctx, readpref.Primary()); err != nil {
		return err
	}
	fmt.Println("connected successfully")
	return nil
}
