package conf

type Config struct {
	Name              string `required:"true"`
	Port              int    `required:"true"`
	SentryDsn         string
	StoragePath       string
	LogPath           string
	DisableLogger     bool
	PrometheusEnabled bool
	PrometheusGateway string
	Profile           bool
	Debug             bool
}

type ClientConfig struct {
	Name string
	Host string
	Port int
}
