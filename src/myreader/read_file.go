// Пакет для чтения файлов
package myreader

import (
	"fmt"
	"github.com/akaKAIN/Go_GB_base/src/myreader/models"
	"gopkg.in/yaml.v3"
	"io/ioutil"
	"log"
	"os"
)

// Открытие файла на чтение
// Возвращает массив байтов прочитанного файла и ошибку.
func GetTextFromFile(path string) ([]byte, error) {
	file, err := os.Open(path)
	if err != nil {
		return nil, fmt.Errorf("Open configFile error: %s\n", err)
	}
	defer func() {
		if err := file.Close(); err != nil {
			log.Fatalf("Close configFile error: %s", err)
		}
	}()
	// Читаем открытый файл
	text, err := ioutil.ReadAll(file)
	if err != nil {
		return nil, err
	}
	return text, nil
}

// Читаем файл с конфигом в формате YAML.
// Возвращает структуру указатель на структуру конфига models.SimpleConfigYAML и ошибку
func GetSimpleConfigYAML(fileName string) (*models.ConfigYAML, error) {
	// Получаем текст с файла
	configText, err := GetTextFromFile(fileName)
	if err != nil {
		return nil, err
	}

	// Парсим прочитанный текст в структуру
	config := new(models.ConfigYAML)
	if err := yaml.Unmarshal(configText, config); err != nil {
		return nil, err
	}

	return config, nil
}
