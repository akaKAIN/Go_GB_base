package models

import "regexp"

type PortString string
type URLDataBase string
type URLString string
type MessageBrokerString string

type SimpleConfigJSON struct {
	Port        PortString          `json:"port"`
	DBUrl       URLDataBase         `json:"db_url"`
	JaegerURL   URLString           `json:"jaeger_url"`
	SentryURL   URLString           `json:"sentry_url"`
	KafkaBroker MessageBrokerString `json:"kafka_broker"`
	AppID       string              `json:"some_app_id"`
	AppKey      string              `json:"some_app_key"`
}

type SimpleConfigYAML struct {
	Port        PortString          `yaml:"port"`
	DBUrl       URLDataBase         `yaml:"db_url"`
	JaegerURL   URLString           `yaml:"jaeger_url"`
	SentryURL   URLString           `yaml:"sentry_url"`
	KafkaBroker MessageBrokerString `yaml:"kafka_broker"`
	AppID       string              `yaml:"some_app_id"`
	AppKey      string              `yaml:"some_app_key"`
}

type ConfigYAML struct {
	Port        string `yaml:"port"`
	DBUrl       string `yaml:"db_url"`
	JaegerURL   string `yaml:"jaeger_url"`
	SentryURL   string `yaml:"sentry_url"`
	KafkaBroker string `yaml:"kafka_broker"`
	AppID       string `yaml:"some_app_id"`
	AppKey      string `yaml:"some_app_key"`
}

func hasMatch(pattern string, byteArr []byte) bool {
	has, err := regexp.Match(pattern, byteArr)
	if err != nil {
		return false
	}
	return has
}

func (u *URLString) IsValid() bool {
	pattern := `^https?://.+?:\d{4,5}$`
	return hasMatch(pattern, []byte(*u))
}

func (u *URLDataBase) IsValid() bool {
	pattern := `^\w+://.+?:.+?@.+?:\d{4}/?`
	return hasMatch(pattern, []byte(*u))
}

func (p *PortString) IsValid() bool {
	pattern := `^\d{4,5}$`
	return hasMatch(pattern, []byte(*p))
}
