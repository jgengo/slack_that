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

const configFile = "config.yml"

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

func load() {
	tokens, err := generateTokenMap(configFile)
	if err != nil {
		log.Fatalf("%sconfig (error)%s error while parsing the config file. (%v)", utils.Red, utils.Reset, err)
	}

	for k, v := range tokens {
		task.Gateway[k] = slack.New(v)
		_, err := task.Gateway[k].AuthTest()
		if err != nil {
			log.Printf(
				"%sconfig (warning)%s auth test failed for '%s', deleted from the loaded tokens\n",
				utils.Yellow, utils.Reset, k,
			)
			delete(task.Gateway, k)
		} else {
			log.Printf("config (info) auth test successful for '%s'.", k)
		}

	}

	log.Println("config (info) config file successfully loaded.")
}

// Initiate checks for the config file, and if its its found, try to load it into the program
func Initiate() {
	log.Println("config (info) loading config...")
	if err := check(); err != nil {
		log.Fatalf("%sconfig (error)%s can't access '%s'. (%v)\n", utils.Red, utils.Reset, configFile, err)
	}
	load()
}
