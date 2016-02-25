package main

import (
	"flag"

	"github.com/dlintw/goconf"
)

var CLIParams map[string]interface{} = map[string]interface{}{}
var ConfigParams *goconf.ConfigFile

func LoadConfig() {
	// CLI Only Flags
	CLIParams["post"] = flag.String("post", "", "The target post url (YouTube or Instagram).")
	CLIParams["verbose"] = flag.Bool("verbose", false, "Extra logging to std out")
	CLIParams["training"] = flag.String("training", "", "Training text files.")
	CLIParams["conf"] = flag.String("conf", "", "Path to conf file.")

	// CLI or Conf Flags
	CLIParams["redis"] = flag.String("redis", "127.0.0.1:6379", "Redis server and port.")
	CLIParams["server"] = flag.Bool("server", false, "Run as a web server.")
	CLIParams["port"] = flag.String("port", "8000", "Port for web server to run.")
	CLIParams["ytkey"] = flag.String("ytkey", "", "Google API key.")
	CLIParams["igkey"] = flag.String("igkey", "", "Instagram API key.")
	CLIParams["fbkey"] = flag.String("fbkey", "", "Facebook API key.")
	CLIParams["fbsecret"] = flag.String("fbsecret", "", "Facebook Secret")
	CLIParams["vnuser"] = flag.String("vnuser", "", "Vine Username")
	CLIParams["vnpassword"] = flag.String("vnpassword", "", "Vine Password")
	CLIParams["stopwords"] = flag.String("stopwords", "", "A list of file paths, comma delimited, of stop word files.")

	flag.Parse()

	configPath := CLIParams["conf"]
	if *configPath.(*string) != "" {
		LogMsg("Loading conf file: " + *configPath.(*string))
		var err error
		ConfigParams, err = goconf.ReadConfigFile(*configPath.(*string))
		if err != nil {
			LogMsg(err.Error())
		}
	}
}

func GetConfigBool(key string) bool {
	if _, ok := CLIParams[key]; ok {
		return true
	}

	param, _ := ConfigParams.GetBool("default", key)
	return param
}

func GetConfigString(key string) string {
	if v, ok := CLIParams[key].(*string); ok && *v != "" {
		return *v
	}

	param, _ := ConfigParams.GetString("default", key)
	return param
}
