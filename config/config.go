package config

import (
	"encoding/json"
	"fmt"
	"github.com/mainak55512/ezgit/command"
	"github.com/mainak55512/ezgit/tui"
	"os"
)

// Main config file structure
type EZConfig struct {
	Origin     string `json:"origin"`
	UserEmail  string `json:"useremail"`
	UserID     string `json:"userid"`
	Credential string `json:"credential"`
	BaseBranch string `json:"basebranch"`
	GitIgnored bool   `json:"gitIgnored"`
}

// Config Initializer
func InitEZConfig() EZConfig {
	return EZConfig{
		Origin:     "",
		UserEmail:  "",
		UserID:     "",
		Credential: "",
		BaseBranch: "",
		GitIgnored: false,
	}
}

// To update config fields
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
	case "UserID":
		if val, ok := value.(string); ok {
			ez.UserID = val
			return nil
		} else {
			return fmt.Errorf("Cannot update config")
		}
	case "UserEmail":
		if val, ok := value.(string); ok {
			ez.UserEmail = val
			return nil
		} else {
			return fmt.Errorf("Cannot update config")
		}
	case "Credential":
		if val, ok := value.(string); ok {
			ez.Credential = val
			return nil
		} else {
			return fmt.Errorf("Cannot update config")
		}
	case "BaseBranch":
		if val, ok := value.(string); ok {
			ez.BaseBranch = val
			return nil
		} else {
			return fmt.Errorf("Cannot update config")
		}
	}
	return nil
}

// Initializes .ezgit file, initializes git repo, configures .ezgit with all necessary data and retruns the config structure.
func ConfigEZ() (EZConfig, error) {
	var ezconfig EZConfig

	// Creating .ezgit file and initialize it if doesn't exist
	if _, err := os.Stat(".ezgit"); os.IsNotExist(err) {
		ezgitFile, err := os.Create(".ezgit")
		if err != nil {
			return EZConfig{}, err
		}
		defer ezgitFile.Close()
		config := InitEZConfig()
		configData, _ := json.MarshalIndent(config, "", " ")
		if err := os.WriteFile(".ezgit", configData, 0644); err != nil {
			return EZConfig{}, err
		}
	}

	// Reading config from .ezgit
	config, err := os.ReadFile(".ezgit")
	if err != nil {
		return EZConfig{}, err
	}

	// Storing config in local structure
	if err := json.Unmarshal(config, &ezconfig); err != nil {
		return EZConfig{}, err
	}

	// Initialize git repo
	if _, err := os.Stat(".git"); os.IsNotExist(err) {
		if err := command.GitINIT(); err != nil {
			return EZConfig{}, err
		}
		if ezconfig.BaseBranch == "" {
			base_branch := "master"
			if err := ezconfig.UpdateEZConfig("BaseBranch", base_branch); err != nil {
				return EZConfig{}, err
			}
		}
	}

	// Adds .ezgit to gitignore
	if ezconfig.GitIgnored == false {
		file, err := os.OpenFile(".gitignore", os.O_APPEND|os.O_CREATE|os.O_WRONLY, 0644)
		if err != nil {
			return EZConfig{}, err
		}
		defer file.Close()
		if _, err := file.WriteString(".ezgit\n"); err != nil {
			return EZConfig{}, err
		}
		if err := ezconfig.UpdateEZConfig("GitIgnored", true); err != nil {
			return EZConfig{}, err
		}
	}

	// Initialize remote origin
	if ezconfig.Origin == "" {
		remote_url := tui.StartInputTextModel("Enter the remote url (github/gitlab repo url)")
		command.OriginINIT(remote_url)
		if err := ezconfig.UpdateEZConfig("Origin", remote_url); err != nil {
			return EZConfig{}, err
		}
	}

	// Set git user and email
	if ezconfig.UserID == "" || ezconfig.UserEmail == "" {
		user_id := tui.StartInputTextModel("Enter remote user id")
		user_email := tui.StartInputTextModel("Enter Email id")
		if err := command.UserINIT(user_id, user_email); err != nil {
			return EZConfig{}, err
		}
		if err := ezconfig.UpdateEZConfig("UserID", user_id); err != nil {
			return EZConfig{}, err
		}
		if err := ezconfig.UpdateEZConfig("UserEmail", user_email); err != nil {
			return EZConfig{}, err
		}
	}

	// Set git credential store
	if ezconfig.Credential == "" || ezconfig.Credential != "store" {
		if err := command.CredentialHelperINIT(); err != nil {
			return EZConfig{}, err
		}
		if err := ezconfig.UpdateEZConfig("Credential", "store"); err != nil {
			return EZConfig{}, err
		}
	}

	// Set default branch i.e. main or master
	if ezconfig.BaseBranch == "" {
		base_branch := tui.StartInputTextModel("Enter the remote default branch (main/master)")
		// command.OriginINIT(remote_url)
		if err := ezconfig.UpdateEZConfig("BaseBranch", base_branch); err != nil {
			return EZConfig{}, err
		}
	}

	// Updating .ezgit file
	configData, _ := json.MarshalIndent(ezconfig, "", " ")
	if err := os.WriteFile(".ezgit", configData, 0644); err != nil {
		return EZConfig{}, err
	}

	// Returning config structure
	return ezconfig, nil
}
