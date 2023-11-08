package database

import (
	"fmt"
	"os"
)

func getStringConnection() string {
	DB_PORT, havePort := os.LookupEnv("DB_PORT")

	if !havePort {
		panic("Database port is not valid")
	}

	DB_HOST, haveHost := os.LookupEnv("DB_HOST")
	if !haveHost {
		panic("Database host is not valid")
	}

	DB_NAME, haveName := os.LookupEnv("DB_NAME")
	if !haveName {
		panic("Database name is not valid")
	}

	DB_USERNAME, haveUsername := os.LookupEnv("DB_USERNAME")
	if !haveUsername {
		panic("Database username is not valid")
	}

	DB_PASSWORD, havePassword := os.LookupEnv("DB_PASSWORD")
	if !havePassword {
		panic("Database password is not valid")
	}

	DSN := fmt.Sprintf("host=%s user=%s password=%s dbname=%s port=%s", DB_HOST, DB_USERNAME, DB_PASSWORD, DB_NAME, DB_PORT)
	return DSN
}
