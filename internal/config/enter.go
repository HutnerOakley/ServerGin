package config

/*
配置
*/
type Config struct {
	DB     DB     `yaml:"db"`
	Redis  Redis  `yaml:"redis"`
	System System `yaml:"system"`
	JWT    JWT    `yaml:"jwt"`
}
