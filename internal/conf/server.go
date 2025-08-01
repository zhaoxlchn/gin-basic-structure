package conf

type Server struct {
	Name    string `json:"name" mapstructure:"name" yaml:"name"`
	Version string `json:"version" mapstructure:"version" yaml:"version"`
	Http    *Http  `json:"http" mapstructure:"http" yaml:"http"`
}

type Http struct {
	Addr string `json:"addr" mapstructure:"addr" yaml:"addr"`
}
