package config

type Config struct {
	DB_USERNAME string
	DB_PASSWORD string
	DB_PORT     string
	DB_HOST     string
	DB_NAME     string
	JWT_KEY     string
}

func InitConfiguration() Config {

	// return Config{
	// 	DB_USERNAME: "admin",
	// 	DB_PASSWORD: "AWS-D4taB4se",
	// 	DB_NAME:     "capstone-lms",
	// 	DB_PORT:     "3306",
	// 	DB_HOST:     "capstone-backend.c0otv3ohtan6.us-east-1.rds.amazonaws.com",
	// 	JWT_KEY:     "rahasia",
	// }

	return Config{
		DB_USERNAME: "root",
		DB_PASSWORD: "",
		DB_NAME:     "capstone-lms",
		DB_PORT:     "3306",
		DB_HOST:     "localhost",
		JWT_KEY:     "rahasia",
	}
}
