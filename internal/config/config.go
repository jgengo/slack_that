package config

import (
	"bytes"
	"errors"
	"io/ioutil"
	"log"
	"path/filepath"

	"github.com/jgengo/slack_that/internal/task"
	"github.com/jgengo/slack_that/internal/utils"

	"github.com/slack-go/slack"
	"github.com/spf13/viper"
)

const configFile = "./configs/config.yml"

func check() error {
	fd, err := filepath.Abs(configFile)
	if err != nil {
		return err
	}

	if _, err = ioutil.ReadFile(fd); err != nil {
		return err
	}

	return nil
}

func generateTokenMap(filename string) (map[string]string, error) {
	viper.SetConfigType("yaml")

	fd, err := filepath.Abs(filename)
	if err != nil {
		return nil, err
	}

	ymlFile, err := ioutil.ReadFile(fd)
	if err != nil {
		return nil, err
	}

	viper.ReadConfig(bytes.NewBuffer(ymlFile))

	if !viper.IsSet("slacks") {
		return nil, errors.New("slacks key not found")
	}

	return viper.GetStringMapString("slacks"), nil
}

func load() error {
	tokens, err := generateTokenMap(configFile)
	if err != nil {
		return err
	}

	for k, v := range tokens {
		task.Gateway[k] = task.SlackClient{
			Value: slack.New(v),
		}
		_, err := task.Gateway[k].Value.AuthTest()
		if err != nil {
			log.Printf(
				"%sconfig (warning)%s auth test failed for '%s', deleted from the loaded tokens (ERR:%s)",
				utils.Yellow, utils.Reset, k, err,
			)
			delete(task.Gateway, k)
		} else {
			log.Printf("config (info) auth test successful for '%s'.", k)
		}
	}

	log.Println("config (info) config file successfully loaded.")
	return nil
}

// Initiate checks for the config file, and if its its found, try to load it into the program
func Initiate() error {
	log.Println("config (info) loading config...")
	if err := check(); err != nil {
		return err
	}
	if err := load(); err != nil {
		return err
	}
	return nil
}
