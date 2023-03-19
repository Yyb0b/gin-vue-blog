package config

import "strconv"

type Mysql struct {
	Host         string `yaml:"host"`
	Port         int    `yaml:"port"`
	Config       string `yaml:"config"` //高级配置，例如charset
	Db           string `yaml:"db"`
	User         string `yaml:"user"`
	Password     string `yaml:"password"`
	LogLevel     string `yaml:"log_Level"` //日志等级,debug输出全部sql,dev,release线上环境
	MaxIdleConns int    `yaml:"max_idle_conns"`
	MaxOpenConns int    `yaml:"max_open_conns"`
}

func (m *Mysql) Dsn() string {
	return m.User + ":" + m.Password + "@tcp(" + m.Host + ":" + strconv.Itoa(m.Port) + ")/" + m.Db + "?" + m.Config
}
