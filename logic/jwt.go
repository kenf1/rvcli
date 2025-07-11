package logic

import (
	"bytes"
	"encoding/json"
	"fmt"
	"io"
	"net/http"
	"os"
)

func RequestJWT(apiConfig ApiConfig, userConfig UserConfig) (string, error) {
	endpoint := fmt.Sprintf("%s:%s/users/token", apiConfig.Host, apiConfig.Port)

	//create request body
	user := map[string]string{
		"username": userConfig.Username,
		"email":    userConfig.Email,
		"fullname": userConfig.Fullname,
		"password": userConfig.Password,
	}

	//convert to json
	jsonBody, err := json.Marshal(user)
	if err != nil {
		return "", err
	}

	req, err := http.NewRequest("POST", endpoint, bytes.NewBuffer(jsonBody))
	if err != nil {
		return "", err
	}
	req.Header.Set("Content-Type", "application/json")

	client := &http.Client{}
	resp, err := client.Do(req)
	if err != nil {
		return "", err
	}
	defer resp.Body.Close()

	if resp.StatusCode != http.StatusOK {
		return "", fmt.Errorf("unexpected status code: %d", resp.StatusCode)
	}

	body, err := io.ReadAll(resp.Body)
	if err != nil {
		return "", err
	}

	return string(body), nil
}

func unmarshalJWT(data []byte) (JWTResponse, error) {
	var resp JWTResponse

	err := json.Unmarshal(data, &resp)
	if err != nil {
		return resp, fmt.Errorf("error parsing JSON: %w", err)
	}
	return resp, nil
}

func AppendJWT(userConfig UserConfig, jwt string, envFile string) error {
	//open env in append mode
	file, err := os.OpenFile(envFile, os.O_APPEND|os.O_CREATE|os.O_WRONLY, 0644)
	if err != nil {
		return fmt.Errorf("failed to open secrets.env: %w", err)
	}
	defer file.Close()

	//create kv pair
	entry := fmt.Sprintf("USERNAME=%s\nPASSWORD=%s\nEMAIL=%s\nFULLNAME=%s\nJWT_TOKEN=%s\n",
		userConfig.Username,
		userConfig.Password,
		userConfig.Fullname,
		userConfig.Email,
		userConfig.JwtToken)

	//write to file
	if _, err := file.WriteString(entry); err != nil {
		return fmt.Errorf("failed to write file: %w", err)
	}

	return nil
}
