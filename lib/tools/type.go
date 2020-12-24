package tools

import "go.mongodb.org/mongo-driver/bson/primitive"

// ArrayObjectID2ArrayString converter type
func ArrayObjectID2ArrayString(input []primitive.ObjectID) []string {
	var result []string

	for _, row := range input {
		result = append(result, row.Hex())
	}

	return result
}
