package config

type BurjCenterConfig struct {
	HOST string `yaml:"host"`
	ROOT string `yaml:"root"`
}

type Port struct {
	HTTP string `yaml:"http"`
	RPC  string `yaml:"rpc"`
}

type ServerConfig struct {
	NAME string `yaml:"name"`
	Port Port   `yaml:"port"`
}

type DBMongoConfig struct {
	URL string `yaml:"url"`
}

type DBMysqlConfig struct {
	URL string `yaml:"url"`
}

type MonitorConfig struct {
	API string `yaml:"api"`
}

type Config struct {
	BurjCenterConfig BurjCenterConfig `yaml:"burj_center"`
	ServerConfig     ServerConfig     `yaml:"server"`
	DBMongoConfig    DBMongoConfig    `yaml:"db_mongo"`
	DBMysqlConfig    DBMysqlConfig    `yaml:"db_mysql"`
	MonitorConfig    MonitorConfig    `yaml:"monitor"`
}
