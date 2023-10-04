package utils

import "go.mongodb.org/mongo-driver/bson/primitive"

func GetObjectIDFromStringID(id string) primitive.ObjectID {
	objID, _ := primitive.ObjectIDFromHex(id)
	return objID
}

func GetStringIDFromObjectID(id primitive.ObjectID) string {
	return id.Hex()
}
