package main

import (
	"flag"
	"strconv"

	"github.com/dlintw/goconf"
)

var CLIParams map[string]interface{} = map[string]interface{}{}
var ConfigParams *goconf.ConfigFile

func LoadConfig() {
	// CLI Only Flags
	CLIParams["post"] = flag.String("post", "", "The target post url (YouTube or Instagram).")
	CLIParams["verbose"] = flag.Bool("verbose", false, "Extra logging to std out")
	CLIParams["conf"] = flag.String("conf", "", "Path to conf file.")

	// CLI or Conf Flags
	CLIParams["server"] = flag.Bool("server", false, "Run as a web server.")
	CLIParams["port"] = flag.String("port", "8000", "Port for web server to run.")
	CLIParams["ytkey"] = flag.String("ytkey", "", "Google API key.")
	CLIParams["fbkey"] = flag.String("fbkey", "", "Facebook API key.")
	CLIParams["fbsecret"] = flag.String("fbsecret", "", "Facebook Secret")
	CLIParams["stopwords"] = flag.String("stopwords", "", "A list of file paths, comma delimited, of stop word files.")
	CLIParams["html"] = flag.String("html", "", "An override html file path to use instead of the built in version.")
	CLIParams["cache_ttl"] = flag.Int("cache_ttl", 0, "Time in seconds to cache report results. 0 to disable.")

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
	if val, ok := CLIParams[key]; ok {
		return *val.(*bool)
	}

	if ConfigParams != nil {
		param, _ := ConfigParams.GetBool("default", key)
		return param
	}

	return false
}

func GetConfigString(key string) string {
	if v, ok := CLIParams[key].(*string); ok && *v != "" {
		return *v
	}

	if ConfigParams != nil {
		param, _ := ConfigParams.GetString("default", key)
		return param
	}

	return ""
}

func GetConfigInt(key string) int {
	if val, ok := CLIParams[key]; ok {
		if v, ok := val.(*int); ok && *v != 0 {
			return *v
		}
	}

	if ConfigParams != nil {
		if val, err := ConfigParams.GetInt("default", key); err == nil {
			return val
		}
		// Fallback to string parsing if GetInt fails or doesn't exist (though it should)
		if strVal, err := ConfigParams.GetString("default", key); err == nil {
			if intVal, err := strconv.Atoi(strVal); err == nil {
				return intVal
			}
		}
	}

	return 0
}
