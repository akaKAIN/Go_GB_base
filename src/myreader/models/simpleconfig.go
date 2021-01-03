package models

import (
	"errors"
	"fmt"
	"github.com/sirupsen/logrus"
	"net/url"
	"regexp"
	"strconv"
	"strings"
	"unicode/utf8"
)

const (
	MinLengthName = 2
)

var (
	noPortError = errors.New("no port value")
)

type Validator interface {
	IsValid() bool
}

type PortString string
type URLDataBase string
type URLString string
type MessageBrokerString string

// Структура для конфигруционного файла в YAML-формате.
// Валидация через резулярные выражения
type SimpleConfigYAML struct {
	Port        PortString          `yaml:"port" json:"port"`
	DBUrl       URLDataBase         `yaml:"db_url" json:"db_url"`
	JaegerURL   URLString           `yaml:"jaeger_url" json:"jaeger_url"`
	SentryURL   URLString           `yaml:"sentry_url" json:"sentry_url"`
	KafkaBroker MessageBrokerString `yaml:"kafka_broker" json:"kafka_broker"`
	AppID       string              `yaml:"some_app_id" json:"some_app_id"`
	AppKey      string              `yaml:"some_app_key" json:"some_app_key"`
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
type ConfigStruct struct {
	Port        string `yaml:"port" json:"port"`
	DBUrl       string `yaml:"db_url" json:"db_url"`
	JaegerURL   string `yaml:"jaeger_url" json:"jaeger_url"`
	SentryURL   string `yaml:"sentry_url" json:"sentry_url"`
	KafkaBroker string `yaml:"kafka_broker" json:"kafka_broker"`
	AppID       string `yaml:"some_app_id" json:"some_app_id"`
	AppKey      string `yaml:"some_app_key" json:"some_app_key"`
}

type ValidationResult struct {
	field  string
	Errors map[string][]error
}

func (v *ValidationResult) GetField() string {
	return v.field
}

func (v *ValidationResult) SetField(field string) {
	if field != "" {
		v.field = field
	}
}

func (v *ValidationResult) AddError(err error) {
	v.Errors[v.field] = append(v.Errors[v.field], err)
}

func (v *ValidationResult) IsValid() bool {
	return len(v.Errors) == 0
}

func (v *ValidationResult) Detail() map[string][]error {
	return v.Errors
}

func (v *ValidationResult) ErrorHandler(err error) {
	if err != nil {
		v.AddError(err)
	}
}

func (v *ValidationResult) Print() {
	if v.IsValid() {
		return
	}
	for key, errorsArr := range v.Errors {
		for _, err := range errorsArr {
			logrus.WithFields(
				logrus.Fields{
					key: err,
				},
			).Error("Validation error:")
		}
	}
}

// Метод валидации полей структуры конфига
func (c *ConfigStruct) GetValidationResult() *ValidationResult {
	result := ValidationResult{
		Errors: make(map[string][]error),
	}

	// Проверка поля "port"
	result.SetField("port")
	result.ErrorHandler(isPortValid(c.Port))

	// Проверка поля "db_url"
	result.SetField("db_url")
	u, err := url.Parse(c.DBUrl)
	result.ErrorHandler(err)

	if utf8.RuneCountInString(u.Scheme) < MinLengthName {
		result.AddError(fmt.Errorf("invalid length of schema name"))
	}
	result.ErrorHandler(isPortValid(u.Port()))

	// Проверка поля "jaeger_ulr"
	result.SetField("jaeger_ulr")
	u, err = url.Parse(c.JaegerURL)
	result.ErrorHandler(err)

	protocol := "http"
	if u != nil && strings.Contains(u.Host, ":") {
		result.ErrorHandler(isPortValid(u.Port()))
		if len(strings.Split(u.Host, ":")[0]) < MinLengthName {
			result.AddError(fmt.Errorf("to short Host name"))
		}

		if !strings.Contains(u.Scheme, protocol) {
			result.AddError(fmt.Errorf("invalid http in schema"))
		}
	} else {
		result.ErrorHandler(noPortError)
	}

	// Проверка поля "sentry_url"
	result.SetField("sentry_url")
	u, err = url.Parse(c.SentryURL)
	result.ErrorHandler(err)
	if u != nil {
		result.ErrorHandler(isPortValid(u.Port()))
		if len(strings.Split(u.Host, ":")[0]) < MinLengthName {
			result.AddError(fmt.Errorf("to short Host name"))
		}

		protocol = "http"
		if !strings.Contains(u.Scheme, protocol) {
			result.AddError(fmt.Errorf("invalid http in schema"))
		}
	}

	// Проверка поля "kafka_broker"
	result.SetField("kafka_broker")
	if strings.LastIndexByte(c.KafkaBroker, ':') != -1 {
		bm := strings.Split(c.KafkaBroker, ":")
		if len(bm) != 2 {
			result.AddError(fmt.Errorf("invalid field value"))
		} else {
			brokerName, port := bm[0], bm[1]
			if utf8.RuneCountInString(brokerName) < MinLengthName {
				result.AddError(fmt.Errorf("to short Host name (broker name)"))
			}
			result.ErrorHandler(isPortValid(port))
		}
	} else {
		result.AddError(noPortError)
	}

	// Проверка поля "some_app_id"
	result.SetField("some_app_id")
	err = fmt.Errorf("no value in field")
	if c.AppID == "" {
		result.AddError(err)
	}

	// Проверка поля "some_app_key"
	result.SetField("some_app_key")
	if c.AppKey == "" {
		result.AddError(err)
	}

	return &result
}

// Проверка "валидности" переданного строкового значения "порта"
// Возвращает ошибку, если не переданное значение не валидно.
func isPortValid(port string) error {
	if strings.TrimSpace(port) == "" {
		return noPortError
	}
	p, err := strconv.Atoi(port)
	if err != nil {
		return fmt.Errorf("incorrect port value")
	}
	if p <= 0 || p > (1<<16)-1 {
		return fmt.Errorf("invalid port number diapason")
	}
	return nil
}
