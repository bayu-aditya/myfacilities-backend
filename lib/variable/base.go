package variable

import (
	"fmt"
	"log"
	"os"
)

// Project variable
var Project struct {
	Mode string
	JWT  struct {
		Key string
	}
	Crypto struct {
		Key string
		IV  string
	}
}

// JwtClaims structure
type JwtClaims struct {
	UserID string
	Email  string
	Exp    int64
}

// Mongo variable
var Mongo struct {
	Host     string
	User     string
	Password string
	AuthDB   string
	DB       string

	URI     string
	adaptor string
}

// InitializationVariableEnvironment for the first time
func InitializationVariableEnvironment() {
	log.Println("Start reading variable environment")

	Project.Mode = os.Getenv("MODE")
	Project.JWT.Key = os.Getenv("JWT_KEY")
	Project.Crypto.Key = os.Getenv("CRYPTO_KEY") // TODO length must be 32
	Project.Crypto.IV = os.Getenv("CRYPTO_IV")   // TODO length must be 16

	switch Project.Mode {
	case "development":
		Mongo.adaptor = "mongodb"
	case "staging", "production":
		Mongo.adaptor = "mongodb+srv"
	default:
		log.Fatalln("Invalid Mode")
	}

	// MONGODB section
	Mongo.User = os.Getenv("MONGO_USER")
	Mongo.Password = os.Getenv("MONGO_PASSWORD")
	Mongo.Host = os.Getenv("MONGO_HOST")
	Mongo.AuthDB = os.Getenv("MONGO_AUTH_DATABASE")
	Mongo.DB = os.Getenv("MONGO_DATABASE")
	Mongo.URI = fmt.Sprintf(
		"%s://%s:%s@%s/%s?retryWrites=true&w=majority",
		Mongo.adaptor,
		Mongo.User,
		Mongo.Password,
		Mongo.Host,
		Mongo.AuthDB,
	)

	log.Println("Finish reading variable environment")
}
