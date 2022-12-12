package config

import (
	"sync"
	"time"

	// _ "github.com/joho/godotenv/autoload" // load .env file automatically

	"github.com/joho/godotenv"
	"github.com/sirupsen/logrus"
	"github.com/spf13/cast"
)

var (
	instance *Configuration
	once     sync.Once
)

// Config ...
func Config() *Configuration {
	once.Do(func() {
		instance = load()
	})

	return instance
}

// Configuration ...
type Configuration struct {
	AppName     string
	AppVersion  string
	AppURL      string
	Environment string
	ServerPort  int
	ServerHost  string

	AccessTokenDuration        time.Duration
	RefreshTokenDuration       time.Duration
	RefreshPasswdTokenDuration time.Duration

	CasbinConfigPath    string
	MiddlewareRolesPath string

	// context timeout in seconds
	CtxTimeout        int
	SigninKey         string
	ServerReadTimeout int

	DefaultOffset string
	DefaultLimit  string

	JWTSecretKey              string
	JWTSecretKeyExpireMinutes int
	JWTRefreshKey             string
	JWTRefreshKeyExpireHours  int

	UserServiceHost     string
	UserServicePort     int

}

func load() *Configuration {
	return &Configuration{
		AppName:             cast.ToString(getOrReturnDefault("APP_NAME", "Navoi Taxi")),
		AppVersion:          cast.ToString(getOrReturnDefault("APP_VERSION", "1.0")),
		AppURL:              cast.ToString(getOrReturnDefault("APP_URL", "localhost:9090")),
		ServerHost:          cast.ToString(getOrReturnDefault("SERVER_HOST", "localhost")),
		ServerPort:          cast.ToInt(getOrReturnDefault("SERVER_PORT", "9000")),
		Environment:         cast.ToString(getOrReturnDefault("ENVIRONMENT", "develop")),
		CtxTimeout:          cast.ToInt(getOrReturnDefault("CTX_TIMEOUT", 7)),
		CasbinConfigPath:    cast.ToString(getOrReturnDefault("CASBIN_CONFIG_PATH", "./config/rbac_model.conf")),
		MiddlewareRolesPath: cast.ToString(getOrReturnDefault("MIDDLEWARE_ROLES_PATH", "./config/models.csv")),
		SigninKey:           cast.ToString(getOrReturnDefault("SIGNIN_KEY", "")),
		ServerReadTimeout:   cast.ToInt(getOrReturnDefault("SERVER_READ_TIMEOUT", "")),

		DefaultOffset: cast.ToString(getOrReturnDefault("DEFAULT_OFFSET", "0")),
		DefaultLimit:  cast.ToString(getOrReturnDefault("DEFAULT_LIMIT", "10")),

		JWTSecretKey:              cast.ToString(getOrReturnDefault("JWT_SECRET_KEY", "")),
		JWTSecretKeyExpireMinutes: cast.ToInt(getOrReturnDefault("JWT_SECRET_KEY_EXPIRE_MINUTES_COUNT", 720)),
		JWTRefreshKey:             cast.ToString(getOrReturnDefault("JWT_REFRESH_KEY", "")),
		JWTRefreshKeyExpireHours:  cast.ToInt(getOrReturnDefault("JWT_REFRESH_KEY_EXPIRE_HOURS_COUNT", 24)),
		UserServiceHost:           cast.ToString(getOrReturnDefault("USER_SERVICE_HOST", "localhost")),
		UserServicePort:           cast.ToInt(getOrReturnDefault("USER_SERVICE_PORT", "9000")),
	}
}

func getOrReturnDefault(key string, defaultValue interface{}) interface{} {
	var envs map[string]string
	envs, err := godotenv.Read(".env")
	if err != nil {
		logrus.Fatal("Error loading .env file")
	}
	if envs[key] != "" {
		return envs[key]
	}
	return defaultValue
}
