package main

import (
	"io/ioutil"
	"os"
	"fmt"

	"gopkg.in/yaml.v2"
	"github.com/op/go-logging"
	"github.com/hudl/fargo"
)


// Global vars
var config Config
var log = logging.MustGetLogger("alertigo")
var eurekaConn fargo.EurekaConnection


// Instance configuration
type Config struct {

	Server_port string
	Log_file string
	Log_format string
	Eureka_port int
	Eureka_ip_addr string
	Eureka_app_name string
	Telegram_bot_token string
}



/**
 * Load configuration yaml file
 */
func LoadConfiguration(filename string) Config {

	// Set config
	source, err := ioutil.ReadFile(filename)
	if err != nil {
		panic(err)
	}

	err = yaml.Unmarshal(source, &config)
	if err != nil {
		panic(err)
	}
	fmt.Printf("--> Configuration loaded values: %#v\n", config)

	// Set logger
	format := logging.MustStringFormatter( config.Log_format )
	logbackend1 := logging.NewLogBackend(os.Stdout, "", 0)
	logbackend1Formatted := logging.NewBackendFormatter(logbackend1, format)
	f, err := os.OpenFile(config.Log_file, os.O_RDWR | os.O_CREATE | os.O_APPEND, 0666)
	if err != nil {
		defer f.Close()
	}
	logbackend2 := logging.NewLogBackend(f, "", 0)
	logbackend2Formatted := logging.NewBackendFormatter(logbackend2, format)
	logging.SetBackend(logbackend1Formatted, logbackend2Formatted)

	return config
}


// Return the already configured logger
func GetLog() *logging.Logger{
	return log
}


// Return the already loaded configuration
func GetConfig() Config {
	return config
}

