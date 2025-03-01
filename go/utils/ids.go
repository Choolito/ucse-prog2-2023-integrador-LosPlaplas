package utils

import "go.mongodb.org/mongo-driver/bson/primitive"

func GetObjectIDFromStringID(id string) primitive.ObjectID {
	objID, _ := primitive.ObjectIDFromHex(id)
	return objID
}

func GetObjectIDFromStringIDErr(id string) (primitive.ObjectID, error) {
	objID, err := primitive.ObjectIDFromHex(id)
	return objID, err
}

func GetStringIDFromObjectID(id primitive.ObjectID) string {
	return id.Hex()
}
