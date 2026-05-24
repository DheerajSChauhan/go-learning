package config

import(
	"os"
	"fmt"
	"github.com/joho/godotenv"

)
type Config struct {
	MongoURI string
	MongoDB string
	ServerPort string
}

func Load() (Config, error){
	//godotenv.Load() loads the .env file and sets the environment variables
	//os.Getenv() retrieves the value of the environment variable named by the key(read those values and assign to the struct fields)

	if err := godotenv.Load(); err != nil {
		fmt.Errorf("failed to load .env file: %v", err)
	}

	mongoURI, err := extractEnv("MONGO_URI")
	if err != nil {
		return Config{}, err
	}
	mongoDB, err := extractEnv("MONGO_DB_NAME")
	if err != nil {
		return Config{}, err
	}
	serverPort, err := extractEnv("SERVER_PORT")
	if err != nil {
		return Config{}, err
	}

	return Config{
		MongoURI:   mongoURI,
		MongoDB:    mongoDB,
		ServerPort: serverPort,
	}, nil

}

func extractEnv(key string) ( string, error){
	val := os.Getenv(key)
	if val == ""{
		return "", fmt.Errorf("missing required environment variable: %s", key)
	}
	return val, nil
}