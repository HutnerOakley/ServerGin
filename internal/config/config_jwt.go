package config

/*
jwt配置
*/
type JWT struct {
	Expires int    `yaml:"expires"` // 过期时间
	Secret  string `yaml:"secret"`
	Issuer  string `yaml:"issuer"`
}
