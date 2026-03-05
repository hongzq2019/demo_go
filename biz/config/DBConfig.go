package config

type YamlConfigStruct struct {

	// 数据库相关配置
	DBConfig struct {
		Host     string       `yaml:"host"`
		Port     int          `yaml:"port"`
		User     string       `yaml:"user"`
		Password string       `yaml:"password"`
		Database string       `yaml:"database"`
		DBPool   DBPoolConfig `yaml:"pool"`
	} `yaml:"database"`

	// Redis相关配置
	RedisConfig struct {
		Host      string          `yaml:"host"`
		Port      int             `yaml:"port"`
		Password  string          `yaml:"password"`
		DB        int             `yaml:"db"`
		RedisPool RedisPoolConfig `yaml:"pool"`
	} `yaml:"redis"`
}
type DBPoolConfig struct {
	MaxOpenConns int `yaml:"max_open_conns"`
	MaxIdleConns int `yaml:"max_idle_conns"`
	MaxLifeTime  int `yaml:"conn_max_lifetime"`
}

type RedisPoolConfig struct {
	PoolSize int `yaml:"poolSize"`
	MaxWait  int `yaml:"maxWait"`
	MaxIdle  int `yaml:"maxIdle"`
}
