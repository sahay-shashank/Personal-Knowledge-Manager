package config

import (
	"encoding/json"
	"os"
	"path/filepath"
	"runtime"
)

func ParseConfig(appName string, userPath *string) (*Config, string, error) {
	var configPath string
	// config from CommandLine has the highest priority
	if userPath != nil && *userPath != "" {
		configPath = *userPath
	} else {
		configPath = getConfigPath(appName)
	}
	// if configPath is found
	if configPath != "" {
		file, err := os.Open(configPath)
		if err != nil {
			return nil, configPath, err
		}
		defer file.Close()
		var config Config
		err = json.NewDecoder(file).Decode(&config)
		return &config, configPath, err
	}
	// fallback to generating a default config
	userDir, _ := os.UserConfigDir()
	defaultPath := filepath.Join(userDir, appName, "config.json")

	os.MkdirAll(filepath.Dir(defaultPath), 0755)
	config := getDefaultConfig(appName)
	data, _ := json.MarshalIndent(config, "", "  ")
	err := os.WriteFile(defaultPath, data, 0644)
	return nil, configPath, err
}

func getConfigPath(appName string) string {
	var paths []string
	paths = append(paths, "config.json")
	if configDir, err := os.UserConfigDir(); err == nil {
		paths = append(paths, filepath.Join(configDir, appName, "config.json"))
	}
	switch runtime.GOOS {
	case "windows":
		if appdata := os.Getenv("APPDATA"); appdata != "" {
			paths = append(paths, filepath.Join(appdata, appName, "config.json"))
		}
	case "darwin":
		if home := os.Getenv("HOME"); home != "" {
			paths = append(paths, filepath.Join(home, "Library/Application Support", appName, "config.json"))
		}
	default: // Linux and other Unix-like
		if xdgConfig := os.Getenv("XDG_CONFIG_HOME"); xdgConfig != "" {
			paths = append(paths, filepath.Join(xdgConfig, appName, "config.json"))
		} else if home := os.Getenv("HOME"); home != "" {
			paths = append(paths, filepath.Join(home, ".config", appName, "config.json"))
		}
	}
	// if file found in any of the path, return it
	for _, path := range paths {
		if _, err := os.Stat(path); err == nil {
			return path
		}
	}
	// else return ""
	return ""
}

func getDefaultConfig(appName string) Config {
	homeDir, err := os.UserHomeDir()
	if err != nil {
		homeDir = "."
	}
	editor := os.Getenv("EDITOR")
	if editor == "" {
		if runtime.GOOS == "windows" {
			editor = "notepad"
		} else {
			editor = "vi"
		}
	}
	os.MkdirAll(filepath.Join(homeDir, appName), 0755)
	return Config{
		StorageLocation: filepath.Join(homeDir, appName),
		Editor:          editor,
	}
}
