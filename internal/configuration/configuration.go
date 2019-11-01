package configuration

import (
	"encoding/json"
	"fmt"
	"io/ioutil"
	"prometheus-metrics-exporter/internal/pmeerrors/config"
	. "prometheus-metrics-exporter/internal/types"
)

func LoadConfig(configFile string) (*Config, error) {
	//Loads the configuration
	content, err := readConfigFile(configFile)

	if err != nil {
		return nil, err
	}

	cfg, err := convertToConfig(content)

	return cfg, err
}

func readConfigFile(fileName string) ([]byte, error) {

	byteValue, err := ioutil.ReadFile(fileName)

	if err != nil {
		errString := fmt.Sprintf("Configuration: Could not read file: \"%s\"", fileName)
		return nil, config.ErrorConfigReadFile{Err: errString}
	}

	return byteValue, err

}

func convertToConfig(input []byte) (*Config, error) {
	var cfg *Config
	err := json.Unmarshal(input, &cfg)

	if err != nil {
		errString := fmt.Sprintf("Configuration: Could not load configuration: \"%s\"", err.Error())
		return nil, config.ErrorConfigConversion{Err: errString}
	}

	return cfg, nil
}
