package config

import "go.mongodb.org/mongo-driver/mongo"

func CreateCollectionInstance(mongoClient *mongo.Client, database string, collectionName string) *mongo.Collection {
	return mongoClient.Database(database).Collection(collectionName)
}

// func mongo() {

// 	//Insert one item in mongoDb
// 	// func InsertOne(collection *mongo.Collection, item interface{}) error {
// 	// _, err := usersCollection.InsertOne(context.TODO(), item)
// 	// insert a single document into a collection
// 	// create a bson.D object
// 	user := bson.D{{"fullName", "User 1"}, {"age", 30}}
// 	// insert the bson object using InsertOne()
// 	result, err := usersCollection.InsertOne(context.TODO(), user)
// 	// check for errors in the insertion
// 	if err != nil {
// 		panic(err)
// 	}
// 	// display the id of the newly inserted object
// 	fmt.Println(result.InsertedID)

// 	// insert multiple documents into a collection
// 	// create a slice of bson.D objects
// 	users := []interface{}{
// 		bson.D{{"fullName", "User 2"}, {"age", 25}},
// 		bson.D{{"fullName", "User 3"}, {"age", 20}},
// 		bson.D{{"fullName", "User 4"}, {"age", 28}},
// 	}
// 	// insert the bson object slice using InsertMany()
// 	results, err := usersCollection.InsertMany(context.TODO(), users)
// 	// check for errors in the insertion
// 	if err != nil {
// 		panic(err)
// 	}
// 	// display the ids of the newly inserted objects
// 	fmt.Println(results.InsertedIDs)
// 	// return err
// 	// }
// }
