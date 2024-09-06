package util

import (
	"fmt"
	"gopkg.in/yaml.v3"
	"os"
)

func ReadYaml(a any, filePath string) error {
	fileBytes, err := os.ReadFile(filePath)
	if err != nil {
		return fmt.Errorf("read yaml file path: %s error: %s", filePath, err)
	}

	err = yaml.Unmarshal(fileBytes, &a)
	if err != nil {
		return fmt.Errorf("unmarshal yaml file bytes: %v, yaml config instance: %#v, error: %s", fileBytes, a, err.Error())
	}

	return nil
}
