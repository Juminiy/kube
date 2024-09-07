package config

type Config struct {
	Log        Log        `yaml:"log"`
	Harbor     Harbor     `yaml:"harbor"`
	Docker     Docker     `yaml:"docker"`
	Kubernetes Kubernetes `yaml:"kubernetes"`
	Minio      Minio      `yaml:"minio"`
}

type Log struct {
	Engine string `yaml:"engine"`
	Zap    struct {
		Level        string   `yaml:"level"`
		Caller       bool     `yaml:"caller"`
		Stacktrace   bool     `yaml:"stacktrace"`
		Path         []string `yaml:"path"`
		InternalPath []string `yaml:"internalPath"`
	} `yaml:"zap"`
	Zero struct {
		Level string `yaml:"level"`
		Path  string `yaml:"path"`
	} `yaml:"zero"`
}

type Harbor struct {
	Registry string
	Username string
	Password string
	Public   string
}

type Docker struct {
	Host    string
	Version string
}

type Kubernetes struct {
	KubeConfigPath string
}

type Minio struct {
	Endpoint        string
	AccessKeyID     string
	SecretAccessKey string
	SessionToken    string
	Secure          bool
}
