package main

import (
	"bytes"
	"errors"
	"io/ioutil"
	"log"
	"path/filepath"

	"github.com/slack-go/slack"
	"github.com/spf13/viper"
)

const configFile = "config.yml"

var gateway = make(map[string]*slack.Client)

func checkConfig() error {
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

func loadConfig() {
	tokens, err := generateTokenMap(configFile)
	if err != nil {
		log.Fatalf("Error while parsing the config file.\n\t->%v", err)
	}

	for k, v := range tokens {
		gateway[k] = slack.New(v)
	}

	log.Printf("tokens:\n%v", tokens)

	log.Printf("gateway:\n%v", gateway)
}

func main() {
	if err := checkConfig(); err != nil {
		log.Fatalf("Can't access %s\n\t-> %v", configFile, err)
	}
	loadConfig()

	for _, v := range gateway {
		user, err := v.GetUserByEmail("jordane@hive.fi")
		if err != nil {
			log.Fatalf("error while retrieving the user:\n\t->%v", err)
		}
		log.Printf("user:\n%v", user)
	}

}
