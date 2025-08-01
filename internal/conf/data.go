package conf

type Data struct {
	Database *Database `json:"database" mapstructure:"database" yaml:"database"`
	Redis    *Redis    `json:"redis" mapstructure:"redis" yaml:"redis"`
}

// Database database config
type Database struct {
	Driver          string `json:"driver" mapstructure:"driver" yaml:"driver"`
	Source          string `json:"source" mapstructure:"source" yaml:"source"`
	ConnMaxLifeTime int64  `json:"conn-max-lifetime" mapstructure:"conn-max-lifetime" yaml:"conn-max-lifetime,omitempty"`
	MaxOpenConn     int    `json:"max-open-conns" mapstructure:"max-open-conns" yaml:"max-open-conns,omitempty"`
	MaxIdleConn     int    `json:"max-idle-conns" mapstructure:"max-idle-conns" yaml:"max-idle-conns,omitempty"`
}

type Redis struct {
	Addr         string `json:"addr" mapstructure:"addr" yaml:"addr"`
	Password     string `json:"password" mapstructure:"password" yaml:"password"`
	Db           int    `json:"db" mapstructure:"db" yaml:"db"`
	ReadTimeout  int64  `json:"read-timeout" mapstructure:"read-timeout" yaml:"read-timeout"`
	WriteTimeout int64  `json:"write-timeout" mapstructure:"write-timeout" yaml:"write-timeout"`
}
