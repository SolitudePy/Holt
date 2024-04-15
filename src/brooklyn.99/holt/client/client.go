package main

import (
	"fmt"
	"os"
)

var global_client_hostname string
var global_server_instance_name string
var global_server_host string
var global_server_port int
var global_server_interval int

func main() {

	hostname, err := os.Hostname()
	if err != nil {
		write_log("init", fmt.Sprintf("Error getting hostname: %v\n", err))
		os.Exit(1)
	}

	global_client_hostname = hostname

	/////////////////////////////////////////////////////////////////////////////////////
	//                                                                                 //
	//                                  Config                                         //
	//                                                                                 //
	/////////////////////////////////////////////////////////////////////////////////////

	// Loading configuration from disk (implemented in config.go)
	config, err := load_config()
	var current_config_section string

	if err != nil {
		write_log("init", fmt.Sprintf("Error loading config: %v\n", err))
		os.Exit(1)
	}

	// Loading the server section
	current_config_section = "server"
	if config[current_config_section] != nil {
		global_server_instance_name = config[current_config_section]["instance_name"].(string)
		global_server_host = config[current_config_section]["host"].(string)
		global_server_port = config[current_config_section]["port"].(int)
		global_server_interval = config[current_config_section]["interval"].(int)

		// Logs the loaded `server` configuration
		write_log("init", fmt.Sprintf("Loading %s configuration:\n\t[%s]\n\tinstance_name = %s\n\thost = %s\n\tport = %d\n\tinterval = %d\n",
			current_config_section, current_config_section, global_server_instance_name, global_server_host, global_server_port, global_server_interval))
	}

	// Loading the server section
	current_config_section = "client"
	if config[current_config_section] != nil {
		global_server_port = config[current_config_section]["port"].(int)

		// Logs the loaded `client` configuration
		write_log("init", fmt.Sprintf("Loading %s configuration:\n\t[%s]\n\tinstance_name = %s\n\thost = %s\n\tport = %d\n\tinterval = %d\n",
			current_config_section, current_config_section, global_server_instance_name, global_server_host, global_server_port, global_server_interval))
	}

	get("/help")
}
