package models

import (
	"fmt"
	"net/url"
	"regexp"
	"strconv"
	"strings"
	"unicode/utf8"
)

const (
	MinLengthName = 2
)

type Validator interface {
	IsValid() bool
}

type PortString string
type URLDataBase string
type URLString string
type MessageBrokerString string

// Структура для конфигруционного файла в JSON-формате.
// Валидация через резулярные выражения
type SimpleConfigJSON struct {
	Port        PortString          `json:"port"`
	DBUrl       URLDataBase         `json:"db_url"`
	JaegerURL   URLString           `json:"jaeger_url"`
	SentryURL   URLString           `json:"sentry_url"`
	KafkaBroker MessageBrokerString `json:"kafka_broker"`
	AppID       string              `json:"some_app_id"`
	AppKey      string              `json:"some_app_key"`
}

// Структура для конфигруционного файла в YAML-формате.
// Валидация через резулярные выражения
type SimpleConfigYAML struct {
	Port        PortString          `yaml:"port"`
	DBUrl       URLDataBase         `yaml:"db_url"`
	JaegerURL   URLString           `yaml:"jaeger_url"`
	SentryURL   URLString           `yaml:"sentry_url"`
	KafkaBroker MessageBrokerString `yaml:"kafka_broker"`
	AppID       string              `yaml:"some_app_id"`
	AppKey      string              `yaml:"some_app_key"`
}

// Проверка на соответствия переданой строки паттерну регурялного выражения
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

// Структура для файла конфигруций.
// Валидация реализована через пакет net/url
type ConfigYAML struct {
	Port        string `yaml:"port"`
	DBUrl       string `yaml:"db_url"`
	JaegerURL   string `yaml:"jaeger_url"`
	SentryURL   string `yaml:"sentry_url"`
	KafkaBroker string `yaml:"kafka_broker"`
	AppID       string `yaml:"some_app_id"`
	AppKey      string `yaml:"some_app_key"`
}

type ValidationResult struct {
	Errors map[string][]error
}

func (v *ValidationResult) AddError(key string, err error) {
	v.Errors[key] = append(v.Errors[key], err)
}

func (v *ValidationResult) AddManyErrors(key string, errArr ...error) {
	for _, err := range errArr {
		v.AddError(key, err)
	}
}

func (v *ValidationResult) IsValid() bool {
	return len(v.Errors) == 0
}

func (v *ValidationResult) Detail() map[string][]error {
	return v.Errors
}

func (v *ValidationResult) String() string {
	var result string
	for key := range v.Errors {
		result += fmt.Sprintf("%s\n", key)
		for _, err := range v.Errors[key] {
			result += fmt.Sprintf("\t- %s\n", err)
		}
	}
	if result == "" {
		result = fmt.Sprintf("config is valid")
	}
	return result
}

// Метод валидации полей структуры конфига
func (c *ConfigYAML) GetValidationResult() *ValidationResult {
	var field string
	result := ValidationResult{
		Errors: make(map[string][]error),
	}

	// Проверка поля "port"
	field = "port"
	err := isPortValid(c.Port)
	if err != nil {
		result.AddError(field, err)
	}

	// Проверка поля "db_url"
	field = "db_url"
	u, err := url.Parse(c.DBUrl)
	if err != nil {
		result.AddError(field, err)
	}
	if utf8.RuneCountInString(u.Scheme) < MinLengthName {
		result.AddError(field, fmt.Errorf("invalid length of schema name"))
	}
	err = isPortValid(u.Port())
	if err != nil {
		result.AddError(field, err)
	}

	// Проверка поля "jaeger_ulr"
	field = "jaeger_ulr"
	u, err = url.Parse(c.JaegerURL)
	if err != nil {
		result.AddError(field, err)
	}
	err = isPortValid(u.Port())
	if err != nil {
		result.AddError(field, err)
	}
	if len(strings.Split(u.Host, ":")[0]) < MinLengthName {
		result.AddError(field, fmt.Errorf("to short Host name"))
	}
	protocol := "http"
	if !strings.Contains(u.Scheme, protocol) {
		result.AddError(field, fmt.Errorf("invalid http in schema"))
	}

	// Проверка поля "sentry_url"
	field = "sentry_url"
	u, err = url.Parse(c.SentryURL)
	if err != nil {
		result.AddError(field, err)
	}
	err = isPortValid(u.Port())
	if err != nil {
		result.AddError(field, err)
	}
	if len(strings.Split(u.Host, ":")[0]) < MinLengthName {
		result.AddError(field, fmt.Errorf("to short Host name"))
	}
	protocol = "http"
	if !strings.Contains(u.Scheme, protocol) {
		result.AddError(field, fmt.Errorf("invalid http in schema"))
	}

	// Проверка поля "kafka_broker"
	field = "kafka_broker"
	bm := strings.Split(c.KafkaBroker, ":")
	if len(bm) != 2 {
		result.AddError(field, fmt.Errorf("invalid field value"))
	} else {
		brokerName, port := bm[0], bm[1]
		if utf8.RuneCountInString(brokerName) < MinLengthName {
			result.AddError(field, fmt.Errorf("to short Host name (broker name)"))
		}
		err = isPortValid(port)
		if err != nil {
			result.AddError(field, err)
		}
	}

	// Проверка поля "some_app_id"
	field = "some_app_id"
	err = fmt.Errorf("no value in field")
	if c.AppID == "" {
		result.AddError(field, err)
	}

	// Проверка поля "some_app_key"
	field = "some_app_key"
	if c.AppKey == "" {
		result.AddError(field, err)
	}

	return &result
}

func isPortValid(port string) error {
	var errArr []error
	p, err := strconv.Atoi(port)
	if err != nil {
		errArr = append(errArr, err)
		return err
	}
	if p <= 0 || p > (1<<16)-1 {
		return fmt.Errorf("invalid port number diapason")
	}
	return nil
}
