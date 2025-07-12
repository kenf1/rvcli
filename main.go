package main

import (
	"fmt"
	"os"

	"github.com/kenf1/rvcli/logic"
)

func main() {
	fmt.Println("\nrvcli: A simple cli to setup rvapi configs")

	//load user configs
	rvconfig := "rvconfig.env"
	err := logic.ImportEnv(rvconfig)
	if err != nil {
		config_fields := []string{"username", "password", "fullname", "email", "host"}

		err := logic.PromptTextWrapper(config_fields)
		if err != nil {
			fmt.Println(err)
			return
		}

		err1 := logic.CreateEnv(rvconfig)
		if err1 != nil {
			fmt.Println(err)
			return
		}
	}

	apiConfig := logic.ApiConfig{
		Host: os.Getenv("HOST"),
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
		err := logic.CheckInputs(userConfig, apiConfig)
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

		fmt.Println(res.Token)
		logic.AppendJWT(res.Token, rvconfig)
	} else {
		fmt.Println("\nYour jwt token is:", userConfig.JwtToken)
	}
}
