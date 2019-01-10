package main

import (
	"github.com/orbs-network/orbs-client-sdk-go/gammacli/jsoncodec"
	"io/ioutil"
)

func getDefaultLocalConfig() *jsoncodec.ConfEnv {
	return &jsoncodec.ConfEnv{
		VirtualChain: 42,
		Endpoints:    []string{"localhost"},
	}
}

func getDefaultExperimentalConfig() *jsoncodec.ConfEnv {
	return &jsoncodec.ConfEnv{
		VirtualChain: 42,
		Endpoints:    []string{"localhost"},
		Experimental: true,
	}
}

func getDefaultConfigForEnv(env string) *jsoncodec.ConfEnv {
	if env == LOCAL_ENV_ID {
		return getDefaultLocalConfig()
	}
	if env == EXPERIMENTAL_ENV_ID {
		return getDefaultExperimentalConfig()
	}
	return nil
}

func getEnvironmentFromConfigFile(env string) *jsoncodec.ConfEnv {
	bytes, err := ioutil.ReadFile(*flagConfigFile)
	if err != nil {
		if res := getDefaultConfigForEnv(env); res != nil {
			return res
		}
		die("Could not open config file '%s' containing environment details.\n\n%s", *flagConfigFile, err.Error())
	}

	confFile, err := jsoncodec.UnmarshalConfFile(bytes)
	if err != nil {
		die("Failed parsing config json file '%s'.\n\n%s", *flagConfigFile, err.Error())
	}

	if len(confFile.Environments) == 0 {
		if res := getDefaultConfigForEnv(env); res != nil {
			return res
		}
		die("Key 'Environments' does not contain data in config file '%s'.", *flagConfigFile)
	}

	confEnv, found := confFile.Environments[env]
	if !found {
		if res := getDefaultConfigForEnv(env); res != nil {
			return res
		}
		die("Environment with id '%s' not found in config file '%s'.", env, *flagKeyFile)
	}

	return confEnv
}
