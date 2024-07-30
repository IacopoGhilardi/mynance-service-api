package config

import (
	"log"
	"os"
	"path/filepath"

	"github.com/iacopoghilardi/mynance-service-api/internal/utils"
	"github.com/spf13/viper"
)

type Config struct {
	DBHost           string `mapstructure:"DB_HOST"`
	DBPort           int    `mapstructure:"DB_PORT"`
	DBUser           string `mapstructure:"DB_USER"`
	DBPass           string `mapstructure:"DB_PASS"`
	DBName           string `mapstructure:"DB_NAME"`
	JwtSecret        string `mapstructure:"JWT_SECRET"`
	GocardlessSecret string `mapstructure:"GOCARDLESS_SECRET"`
	GocardlessToken  string `mapstructure:"GOCARDLESS_TOKEN"`
}

var AppConfig Config
var logger = utils.Logger

func InitConfig() error {
	logger.Info("Setting configs...")
	projectRoot, err := getProjectRoot()
	if err != nil {
		log.Fatalf("Error determining project root directory: %v", err)
		return err
	}

	env := os.Getenv("APP_ENV")
	if env == "" {
		env = "dev"
	}

	viper.SetConfigName(".env." + env)
	viper.SetConfigType("env")
	viper.AddConfigPath(projectRoot)
	viper.AutomaticEnv()

	err = viper.ReadInConfig()
	if err != nil {
		log.Fatalf("Error while reading config file %s", err)
		return err
	}

	if err := viper.Unmarshal(&AppConfig); err != nil {
		log.Fatalf("Unable to decode into struct, %v", err)
		return err
	}

	return nil
}

func getProjectRoot() (string, error) {
	currentDir, err := os.Getwd()
	if err != nil {
		return "", err
	}

	for {
		if fileExists(filepath.Join(currentDir, ".env")) {
			return currentDir, nil
		}

		parentDir := filepath.Dir(currentDir)
		if parentDir == currentDir {
			break
		}
		currentDir = parentDir
	}

	return "", os.ErrNotExist
}

func fileExists(filename string) bool {
	info, err := os.Stat(filename)
	return err == nil && !info.IsDir()
}
