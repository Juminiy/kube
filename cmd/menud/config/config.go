package config

type Config struct {
	Log        Log        `yaml:"log"`
	Kubernetes Kubernetes `yaml:"kubernetes"`
	Harbor     Harbor     `yaml:"harbor"`
	Docker     Docker     `yaml:"docker"`
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

type Kubernetes struct {
	KubeConfigPath string `yaml:"kubeConfigPath"`
}

type Harbor struct {
	Registry string `yaml:"registry"`
	Username string `yaml:"username"`
	Password string `yaml:"password"`
	Public   string `yaml:"public"`
}

type Docker struct {
	Host    string `yaml:"host"`
	Version string `yaml:"version"`
}

type Minio struct {
	Endpoint        string `yaml:"endpoint"`
	AccessKeyID     string `yaml:"accessKeyID"`
	SecretAccessKey string `yaml:"secretAccessKey"`
	SessionToken    string `yaml:"sessionToken"`
	Secure          bool   `yaml:"secure"`
}
