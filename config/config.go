package config

import (
	"encoding/json"
	"fmt"
	"mainak55512/ezgit/command"
	"os"
)

type EZConfig struct {
	Origin     string `json:"origin"`
	GitIgnored bool   `json:"gitIgnored"`
}

func InitEZConfig() EZConfig {
	return EZConfig{
		GitIgnored: false,
	}
}

func (ez *EZConfig) UpdateEZConfig(field string, value any) error {
	switch field {
	case "Origin":
		if val, ok := value.(string); ok {
			ez.Origin = val
			return nil
		} else {
			return fmt.Errorf("Cannot update config")
		}
	case "GitIgnored":
		if val, ok := value.(bool); ok {
			ez.GitIgnored = val
			return nil
		} else {
			return fmt.Errorf("Cannot update config")
		}
	}
	return nil
}

func ConfigEZ() error {
	var ezconfig EZConfig
	_, err := os.Stat(".ezgit")
	if os.IsNotExist(err) {
		ezgitFile, err := os.Create(".ezgit")
		if err != nil {
			return err
		}
		defer ezgitFile.Close()
		config := InitEZConfig()
		configData, _ := json.MarshalIndent(config, "", " ")
		if err := os.WriteFile(".ezgit", configData, 0644); err != nil {
			return err
		}
	}
	config, err := os.ReadFile(".ezgit")
	if err != nil {
		return err
	}
	if err := json.Unmarshal(config, &ezconfig); err != nil {
		return err
	}
	if ezconfig.GitIgnored == false {
		file, err := os.OpenFile(".gitignore", os.O_APPEND|os.O_CREATE|os.O_WRONLY, 0644)
		if err != nil {
			return err
		}
		defer file.Close()
		if _, err := file.WriteString(".ezgit"); err != nil {
			return err
		}
		ezconfig.UpdateEZConfig("GitIgnored", true)
	}
	if ezconfig.Origin == "" {
		var remote_url string
		fmt.Print("Enter the remote url: ")
		fmt.Scanln(&remote_url)
		ezconfig.UpdateEZConfig("Origin", remote_url)
		if err := command.OriginINIT(remote_url); err != nil {
			return err
		}
	}
	configData, _ := json.MarshalIndent(ezconfig, "", " ")
	if err := os.WriteFile(".ezgit", configData, 0644); err != nil {
		return err
	}
	return nil
}

func EZInit() error {
	if _, err := os.Stat(".git"); os.IsNotExist(err) {
		command.GitINIT()
	}
	if err := ConfigEZ(); err != nil {
		return err
	}
	return nil
}
