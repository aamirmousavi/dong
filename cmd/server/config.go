package server

import (
	"fmt"
	"os"
	"path/filepath"

	"gopkg.in/yaml.v3"
)

func loadConfig() error {
	appConfigAddress, err := setAddressSiblingsExecutable("dong.yaml")
	if err != nil {
		return err
	}

	type Config struct {
		Addr        string `yaml:"addr"`
		MongodbAddr string `yaml:"mongodb"`
		SmsUsername string `yaml:"sms-username"`
		SmsPassword string `yaml:"sms-password"`
	}
	var config Config
	b, err := os.ReadFile(string(appConfigAddress))
	if err != nil {
		return err
	}
	if err := yaml.Unmarshal(b, &config); err != nil {
		return err
	}

	setConfigVariable(&addr, config.Addr)
	setConfigVariable(&mongodbAddr, config.MongodbAddr)
	setConfigVariable(&smsUsername, config.SmsUsername)
	setConfigVariable(&smsPassword, config.SmsPassword)

	fmt.Println("Config loaded successfully")

	return nil
}

func setConfigVariable(_var *string, _val string) {
	*_var = _val
}

func setAddressSiblingsExecutable(filename string) (string, error) {
	ex, err := os.Executable()
	if err != nil {
		return "", err
	}
	exPath := filepath.Dir(ex)
	return filepath.Join(exPath, filename), nil
}
