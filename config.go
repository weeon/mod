package mod

type Database struct {
	Driver   string `json:"driver"`
	Host     string `json:"host"`
	Port     int    `json:"port"`
	User     string `json:"user"`
	Password string `json:"password"`
	Database string `json:"database"`
	TLS      bool   `json:"tls"`
}

type Redis struct {
	Host     string `json:"host"`
	Port     int    `json:"port"`
	Password string `json:"password"`
	DB       int    `json:"db"`
}

type InfluxDB struct {
	Addr     string
	Token    string // for v2
	Username string
	Password string
}

type Telegram struct {
	Name  string `json:"name"`
	Token string `json:"token"`
}

type Mongo struct {
	Uri string `json:"uri"`
}

type Minio struct {
	Endpoint        string
	AccessKeyID     string
	SecretAccessKey string
	UseSSL          bool
}
