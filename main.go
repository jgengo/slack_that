package main

import (
	"bytes"
	"errors"
	"io/ioutil"
	"log"
	"net/http"
	"path/filepath"

	"github.com/slack-go/slack"
	"github.com/spf13/viper"
)

const configFile = "config.yml"

const (
	Reset  = "\033[0m"
	Red    = "\033[31m"
	Green  = "\033[32m"
	Yellow = "\033[1;33m"
)

// Gateway is the Slack API client gateway
var Gateway = make(map[string]*slack.Client)

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
		log.Fatalf("%sconfig (error)%s error while parsing the config file. (%v)", Red, Reset, err)
	}

	for k, v := range tokens {
		Gateway[k] = slack.New(v)
		_, err := Gateway[k].AuthTest()
		if err != nil {
			log.Printf(
				"%sconfig (warning)%s auth test failed for '%s', deleted from the loaded tokens\n",
				Yellow, Reset, k,
			)
			delete(Gateway, k)
		} else {
			log.Printf("config (info) auth test successful for '%s'.", k)
		}

	}

	log.Println("config (info) config file successfully loaded.")
}

func main() {
	log.Println("config (info) loading config...")
	if err := checkConfig(); err != nil {
		log.Fatalf("%sconfig (error)%s can't access '%s'. (%v)\n", Red, Reset, configFile, err)
	}
	loadConfig()

	router := NewRouter()
	http.ListenAndServe("localhost:8080", router)
}
