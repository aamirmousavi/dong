package server

import (
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

	addr = config.Addr
	mongodbAddr = config.MongodbAddr
	smsUsername = config.SmsUsername
	smsPassword = config.SmsPassword

	return nil
}

func setAddressSiblingsExecutable(filename string) (string, error) {
	ex, err := os.Executable()
	if err != nil {
		return "", err
	}
	exPath := filepath.Dir(ex)
	return filepath.Join(exPath, filename), nil
}
