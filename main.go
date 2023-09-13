package main

import (
	"context"
	"fmt"
	"time"

	"go.mongodb.org/mongo-driver/bson"
	"go.mongodb.org/mongo-driver/mongo"
	"go.mongodb.org/mongo-driver/mongo/options"
	"go.mongodb.org/mongo-driver/mongo/readpref"
)

func close(client *mongo.Client, ctx context.Context, cancel context.CancelFunc) {

	defer cancel()
	defer func() {
		if err := client.Disconnect(ctx); err != nil {
			panic(err)
		}
	}()

}
func connect(uri string) (*mongo.Client, context.Context, context.CancelFunc, error) {

	ctx, cancel := context.WithTimeout(context.Background(), 30*time.Second)

	client, err := mongo.Connect(ctx, options.Client().ApplyURI(uri))

	return client, ctx, cancel, err

}

func ping(client *mongo.Client, ctx context.Context) error {
	if err := client.Ping(ctx, readpref.Primary()); err != nil {
		return err
	}
	fmt.Println("Connection succesfull")
	return nil

}

// func insertOneStudent(client *mongo.Client, ctx context.Context, dataBase, col string, doc interface{}) (*mongo.InsertOneResult, error) {
// 	collection := client.Database(dataBase).Collection(col)

// 	result, err := collection.InsertOne(ctx, doc)
// 	return result, err
// }
// func insertManyStudents(client *mongo.Client, ctx context.Context, dataBase, col string, docs []interface{}) (*mongo.InsertManyResult, error) {
// 	collection := client.Database(dataBase).Collection(col)

// 	result, err := collection.InsertMany(ctx, docs)
// 	return result, err
// }

// func query(client *mongo.Client, ctx context.Context, dataBase, col string, query, field interface{}) (result *mongo.Cursor, err error) {
// 	collection := client.Database(dataBase).Collection(col)
// 	result, err = collection.Find(ctx, query, options.Find().SetProjection(field))
// 	return
// }

// func UpdateOneStudent(client *mongo.Client, ctx context.Context, dataBase, col string, filter, update interface{}) (result *mongo.UpdateResult, err error) {
// 	collection := client.Database(dataBase).Collection(col)
// 	result, err = collection.UpdateOne(ctx, filter, update)
// 	return
// }

// func UpdateManyStudents(client *mongo.Client, ctx context.Context,
// 	dataBase, col string, filter, update interface{}) (result *mongo.UpdateResult, err error) {
// 	collection := client.Database(dataBase).Collection(col)
// 	result, err = collection.UpdateMany(ctx, filter, update)
// 	return
// }

// func deleteOneStudent(client *mongo.Client, ctx context.Context,
// 	dataBase, col string, query interface{}) (result *mongo.DeleteResult, err error) {
// 	collection := client.Database(dataBase).Collection(col)
// 	result, err = collection.DeleteOne(ctx, query)
// 	return
// }

// func deleteManyStudents(client *mongo.Client, ctx context.Context,
// 	dataBase, col string, query interface{}) (result *mongo.DeleteResult, err error) {
// 	collection := client.Database(dataBase).Collection(col)
// 	result, err = collection.DeleteMany(ctx, query)
// 	return
// }

func replaceDocument(client *mongo.Client, ctx context.Context, filter, replacementDoc interface{}) (result *mongo.UpdateResult, err error) {
	collection := client.Database("school").Collection("students")
	result, err = collection.ReplaceOne(ctx, filter, replacementDoc)
	return
}

func main() {
	client, ctx, cancel, err := connect("mongodb://localhost:27017")

	if err != nil {
		panic(err)
	}

	defer close(client, ctx, cancel)

	// var document interface{}

	// document = bson.D{
	// 	{"first_name", "John"},
	// 	{"last_name", "Doe"},
	// 	{"gpa", 4.5},
	// 	{"id", 1},
	// 	{"doj", "12/12/2023"},
	// }

	// inserOneResult, err := insertOneStudent(client, ctx, "school", "students", document)

	// if err != nil {
	// 	panic(err)
	// }

	// fmt.Printf("The following student has been added %v", inserOneResult.InsertedID)

	// var documents []interface{}

	// documents = []interface{}{
	// 	bson.D{
	// 		{"first_name", "Johnes"},
	// 		{"last_name", "Does"},
	// 		{"gpa", 2.5},
	// 		{"id", 4},
	// 		{"doj", "10/00/2023"},
	// 	},
	// 	bson.D{
	// 		{"first_name", "Eleli"},
	// 		{"last_name", "Ayub"},
	// 		{"gpa", 4.5},
	// 		{"id", 2},
	// 		{"doj", "02/21/2023"},
	// 	},
	// 	bson.D{
	// 		{"first_name", "Liam Neeson"},
	// 		{"last_name", ""},
	// 		{"gpa", 3.5},
	// 		{"id", 3},
	// 		{"doj", "01/22/2023"},
	// 	},
	// }

	// insertManyResult, err := insertManyStudents(client, ctx, "school", "students", documents)
	// if err != nil {
	// 	panic(err)
	// }
	// fmt.Printf("The following students have been added %v", insertManyResult.InsertedIDs)

	// var filter, options interface{}

	// filter = bson.D{
	// 	{"gpa", bson.D{{"$gte", 3.0}}},
	// }

	// options = bson.D{{"_id", 0}}

	// cursor, err := query(client, ctx, "school", "students", filter, options)

	// if err != nil {
	// 	panic(err)
	// }

	// var results []bson.D

	// if err := cursor.All(ctx, &results); err != nil {
	// 	panic(err)
	// }

	// fmt.Println("QUERY RESULT")
	// for _, doc := range results {
	// 	fmt.Println("----------------------------")
	// 	fmt.Println("")
	// 	fmt.Printf("\t %v", doc)
	// 	fmt.Println("")
	// 	fmt.Println("----------------------------")
	// }

	// filter := bson.D{
	// 	{"gpa", bson.D{{"$lt", 3.0}}},
	// }

	// update := bson.D{
	// 	{"$set", bson.D{
	// 		{"gpa", 4.0},
	// 	}},
	// }

	// result, err := UpdateOneStudent(client, ctx, "school",
	// 	"students", filter, update)

	// if err != nil {
	// 	panic(err)
	// }

	// fmt.Println("update single document")
	// fmt.Println(result.ModifiedCount)

	// filter = bson.D{
	// 	{"gpa", bson.D{{"$gt", 4.0}}},
	// }
	// update = bson.D{
	// 	{"$set", bson.D{
	// 		{"gpa", 3.0},
	// 	}},
	// }

	// result, err = UpdateManyStudents(client, ctx, "school",
	// 	"students", filter, update)

	// // handle error
	// if err != nil {
	// 	panic(err)
	// }

	// // print count of documents that affected
	// fmt.Println("update multiple document")
	// fmt.Println(result.ModifiedCount)

	// query := bson.D{
	// 	{"id", bson.D{{"$lt", 2}}},
	// }

	// result, err := deleteOneStudent(client, ctx, "school", "students", query)
	// fmt.Println("No.of rows affected by DeleteOne()")
	// fmt.Println(result.DeletedCount)

	// query = bson.D{
	// 	{"gpa", bson.D{{"$gt", 4.0}}},
	// }

	// result, err = deleteManyStudents(client, ctx, "school", "students", query)
	// fmt.Println("No.of rows affected by DeleteMany()")
	// fmt.Println(result.DeletedCount)

	filter := bson.D{
		{"gpa", 3.0},
	}

	replacementDoc := bson.D{{"name", "Omar"}, {"id", 1}}

	result, err := replaceDocument(client, ctx, filter, replacementDoc)

	if err != nil {
		panic(err)
	}

	fmt.Println("No:of documents replaced")
	fmt.Println(result.ModifiedCount)

	ping(client, ctx)
}
