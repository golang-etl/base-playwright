package config

import (
	"github.com/caarlos0/env/v10"
	"github.com/joho/godotenv"
)

type Config struct {
	GoModuleName              string `env:"GO_MODULE_NAME" envDefault:"github.com/golang-etl/base-playwright"`
	Env                       string `env:"ENV" envDefault:"local"`
	EchoAddress               string `env:"ECHO_ADDRESS" envDefault:"0.0.0.0:8080"`
	Debug                     bool   `env:"DEBUG" envDefault:"false"`
	TraceEnabled              bool   `env:"TRACE_ENABLED" envDefault:"false"`
	TraceBucket               string `env:"TRACE_BUCKET" envDefault:"playwright-trace"`
	DisabledHeadlessInBrowser bool   `env:"DISABLED_HEADLESS_IN_BROWSER" envDefault:"false"`
	EnabledDevtoolsInBrowser  bool   `env:"ENABLED_DEVTOOLS_IN_BROWSER" envDefault:"false"`
	UserAgentHeader           string `env:"USER_AGENT_HEADER" envDefault:"Mozilla/5.0 (Macintosh; Intel Mac OS X 10_15_7) AppleWebKit/537.36 (KHTML, like Gecko) Chrome/134.0.0.0 Safari/537.36"`
	SecChUaHeader             string `env:"SEC_CH_UA_HEADER" envDefault:"\"Chromium\";v=\"134\", \"Not:A-Brand\";v=\"24\", \"Google Chrome\";v=\"134\""`
	ProxyServer               string `env:"PROXY_SERVER" envDefault:""`
	ProxyUsername             string `env:"PROXY_USERNAME" envDefault:""`
	ProxyPassword             string `env:"PROXY_PASSWORD" envDefault:""`
	MongoDBURI                string `env:"MONGODB_URI" envDefault:""`
	MongoDBDatabaseName       string `env:"MONGODB_DATABASE_NAME" envDefault:""`
	SecretKeyUserTokenData    string `env:"SECRET_KEY_USER_TOKEN_DATA" envDefault:""`
}

func LoadConfig() (*Config, error) {
	_ = godotenv.Load(".env")

	cfg := &Config{}

	if err := env.Parse(cfg); err != nil {
		return nil, err
	}

	return cfg, nil
}
