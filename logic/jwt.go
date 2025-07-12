package logic

import (
	"bytes"
	"encoding/json"
	"fmt"
	"io"
	"net/http"
	"os"
)

// send post request to create jwt endpoint
func RequestJWT(apiConfig ApiConfig, userConfig UserConfig) (JWTResponse, error) {
	endpoint := fmt.Sprintf("%s/users/token", apiConfig.Host)

	//create user info
	user := map[string]string{
		"username": userConfig.Username,
		"email":    userConfig.Email,
		"fullname": userConfig.Fullname,
		"password": userConfig.Password,
	}

	//create json body
	jsonBody, err := json.Marshal(user)
	if err != nil {
		return JWTResponse{}, err
	}

	//send post request
	req, err := http.NewRequest("POST", endpoint, bytes.NewBuffer(jsonBody))
	if err != nil {
		return JWTResponse{}, err
	}
	req.Header.Set("Content-Type", "application/json")

	client := &http.Client{}
	resp, err := client.Do(req)
	if err != nil {
		return JWTResponse{}, err
	}
	defer resp.Body.Close()

	//unmarshal to json
	body, err := io.ReadAll(resp.Body)
	if err != nil {
		return JWTResponse{}, err
	}

	if resp.StatusCode != http.StatusOK {
		return JWTResponse{}, fmt.Errorf(
			"unexpected status code: %d, body: %s", resp.StatusCode, string(body),
		)
	}

	var jwtResp JWTResponse
	if err := json.Unmarshal(body, &jwtResp); err != nil {
		return JWTResponse{}, fmt.Errorf("error parsing JSON: %w", err)
	}

	return jwtResp, nil
}

// append jwt token to userconfig env
func AppendJWT(jwtToken string, envFile string) error {
	//open env in append mode
	file, err := os.OpenFile(envFile, os.O_APPEND|os.O_CREATE|os.O_WRONLY, 0644)
	if err != nil {
		return fmt.Errorf("failed to open secrets.env: %w", err)
	}
	defer file.Close()

	//append
	entry := fmt.Sprintf("\nJWT_TOKEN=%s\n", jwtToken)
	if _, err := file.WriteString(entry); err != nil {
		return fmt.Errorf("failed to write file: %w", err)
	}

	return nil
}
