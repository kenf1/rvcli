package logic

type UserConfig struct {
	Username string `env:"USERNAME"`
	Password string `env:"PASSWORD"`
	Fullname string `env:"FULLNAME"`
	Email    string `env:"EMAIL"`
	JwtToken string `env:"JWT_TOKEN"`
}

type ApiConfig struct {
	Host string
	Port string
}

type JWTResponse struct {
	Status     string `json:"status"`
	StatusCode int    `json:"statusCode"`
	Token      string `json:"token"`
	Username   string `json:"username"`
}
