package config

type BurjCenterConfig struct {
	ZkServer  string `yaml:"zk_server"`
	ZkRoot    string `yaml:"zk_root"`
	RpcServer string `yaml:"rpc_server"`
}

type Port struct {
	Http string `yaml:"http"`
	Rpc  string `yaml:"rpc"`
}

type ServerConfig struct {
	Name string `yaml:"name"`
	Port Port   `yaml:"port"`
}

type DBMongoConfig struct {
	URL string `yaml:"url"`
}

type DBMysqlConfig struct {
	URL string `yaml:"url"`
}

type FtpServerConfig struct {
	URL string `yaml:"url"`
}

type Config struct {
	BurjCenterConfig BurjCenterConfig `yaml:"burj_center"`
	ServerConfig     ServerConfig     `yaml:"server"`
	DBMongoConfig    DBMongoConfig    `yaml:"db_mongo"`
	DBMysqlConfig    DBMysqlConfig    `yaml:"db_mysql"`
	FtpServerConfig  FtpServerConfig  `yaml:"ftp_server"`
}

func NewConfig() *Config {
	return &Config{}
}
