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

// Читаем файл с конфигом в формате YAML.
// Возвращает структуру указатель на структуру конфига models.SimpleConfigYAML и ошибку
func GetSimpleConfigYAML(fileName string) (*models.ConfigYAML, error) {
	// Открываем файл по указанному пути
	configFile, err := os.Open(fileName)
	if err != nil {
		return nil, fmt.Errorf("Open configFile error: %s\n", err)
	}
	defer func() {
		if err := configFile.Close(); err != nil {
			log.Fatalf("Close configFile error: %s", err)
		}
	}()

	// Читаем открытый файл
	configText, err := ioutil.ReadAll(configFile)
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

//func GetSimpleConfigJSON()
