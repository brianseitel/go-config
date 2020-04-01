package lib

type BaseConfig struct {
	AWS       awsConfig       `mapstructure:",squash"`
	Metrics   metricsConfig   `mapstructure:",squash"`
	Service   serviceConfig   `mapstructure:",squash"`
	Honeycomb honeycombConfig `mapstructure:",squash"`
	Kafka     kafkaConfig     `mapstructure:",squash"`

	LogLevel     string `mapstructure:"log_level", default:"info"`
	RollbarToken string `mapstructure:"rollbar_token"`
}

const (
	testing     = "testing"
	development = "development"
	integration = "integration"
	staging     = "staging"
	production  = "production"
)

type honeycombConfig struct {
	Key          string `mapstructure:"honeycomb_key"`
	SamplingRate int    `mapstructure:"honeycomb_sampling", default:1`
	V2Tracing    bool   `mapstructure:"honeycomb_use_v2_tracing"`
}

type kafkaConfig struct {
	Key            string `mapstructure:"kafka_key"`
	Brokers        string `mapstructure:"kafka_brokers"`
	Secret         string `mapstructure:"kafka_secret"`
	SchemaRegistry string `mapstructure:"kafka_schema_registry"`
}

type serviceConfig struct {
	Host        string `mapstructure:"server_host",default:"localhost"`
	Port        string `mapstructure:"server_port",default:"8080"`
	Environment string `mapstructure:"service_env",default:"development"`
}

type metricsConfig struct {
	Host   string `mapstructure:"hosted_graphite_host"`
	APIKey string `mapstructure:"hosted_graphite_api_key"`
	Source string `mapstructure:"hosted_graphite_source"`
}

type awsConfig struct {
	AWSAccessKeyID     string `mapstructure:"aws_access_key_id"`
	AWSSecretAccessKey string `mapstructure:"aws_secret_access_key"`
	AWSSessionToken    string `mapstructure:"aws_session_token"`
	AWSRegion          string `mapstructure:"aws_region"`
	AWSEndpoint        string `mapstructure:"aws_endpoint"`
	OktaEnv            string `mapstructure:"okta_env"`
	AWSMaxRetries      int    `mapstructure:"aws_max_retries", default=5"`

	StaticCredentials bool
}
