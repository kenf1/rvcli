package logic

import (
	"fmt"
	"os"

	"github.com/joho/godotenv"
)

type UserConfig struct {
	Username string `env:"USERNAME"`
	Password string `env:"PASSWORD"`
	Fullname string `env:"FULLNAME"`
	Email    string `env:"EMAIL"`
	JwtToken string `env:"JWT_TOKEN"`
}

type ApiConfig struct {
	Host string
}

type JWTResponse struct {
	Status     string `json:"status"`
	StatusCode int    `json:"statusCode"`
	Token      string `json:"token"`
	Username   string `json:"username"`
}

// import env file
func ImportEnv(envFile string) error {
	err := godotenv.Load(envFile)
	if err != nil {
		fmt.Println("userconfig.env file not found. Answer prompts to create.\n")
		return err
	}
	return nil
}

// create env file
func CreateEnv(envFile string) error {
	//create if DNE
	file, err := os.Create(envFile)
	if err != nil {
		return fmt.Errorf("failed to create or open env: %w", err)
	}
	defer file.Close()

	//overwrite existing content
	entry := fmt.Sprintf(
		"USERNAME=%s\nPASSWORD=%s\nFULLNAME=%s\nEMAIL=%s\nHOST==%s\n",
		os.Getenv("USERNAME"),
		os.Getenv("PASSWORD"),
		os.Getenv("FULLNAME"),
		os.Getenv("EMAIL"),
		os.Getenv("HOST"),
	)
	if _, err := file.WriteString(entry); err != nil {
		return fmt.Errorf("failed to write to file: %w", err)
	}

	return nil
}
