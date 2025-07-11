package main

import (
	"fmt"
	"os"

	"github.com/joho/godotenv"
	"github.com/kenf1/rvcli/logic"
)

func main() {
	//load user configs
	envFile := "userconfig.env"
	err := godotenv.Load(envFile)
	if err != nil {
		fmt.Println("Unable to find .env file")
		return
	}

	apiConfig := logic.ApiConfig{
		Host: "http://localhost",
		Port: "8080",
	}

	userConfig := logic.UserConfig{
		Username: os.Getenv("USERNAME"),
		Password: os.Getenv("PASSWORD"),
		Fullname: os.Getenv("FULLNAME"),
		Email:    os.Getenv("EMAIL"),
		JwtToken: os.Getenv("JWT_TOKEN"),
	}

	//generate jwt token if missing
	if userConfig.JwtToken == "" {
		fmt.Println("JWT_TOKEN not found")

		//confirm other fields to generate jwt token present
		err := logic.CheckInputs(userConfig)
		if err != nil {
			fmt.Println(err)
			return
		}

		//query generate jwt endpoint (api checks if jwt is valid)
		res, err := logic.RequestJWT(apiConfig, userConfig)
		if err != nil {
			fmt.Println("Error:", err)
			return
		}

		fmt.Println(res)
		// logic.AppendJWT(userConfig, res, envFile)
	}
}
