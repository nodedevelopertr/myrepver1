package config

import (
	"errors"
	"fmt"
	"net/url"
	"os"
	"path"

	"github.com/distribution/distribution/v3/configuration"
	"github.com/kelseyhightower/envconfig"
	log "github.com/sirupsen/logrus"
	"gopkg.in/yaml.v3"
)

const (
	defaultHomeDirDiscoConfigPath = ".disco/config.yaml"
)

type envVars struct {
	RegistryConfigurationPath string `envconfig:"registry_configuration_path"`
	DiscoPort                 int    `envconfig:"disco_port" default:"1970"`
}

// Node contains IPFS node parameters.
type Node struct {
	URL string `yaml:"url"`
}

// RouterConfig contains router config parameters.
type RouterConfig struct {
	Nodes []*Node `yaml:"nodes"`
}

// Configuration variables
var (
	Vars               envVars
	DistributionConfig *configuration.Configuration
	Router             RouterConfig
	Cache              configuration.Storage
	RedirectTo         *url.URL
)

// DiscoConfig contains the extra configuration settings that blend with
// the distribution library config.
var DiscoConfig struct {
	Storage struct {
		IPFS struct {
			Router   RouterConfig          `yaml:"router"`
			Cache    configuration.Storage `yaml:"cache"`
			Redirect string                `yaml:"redirect"`
		} `yaml:"ipfs"`
	} `yaml:"storage"`
	Disco struct {
		NoClone bool `yaml:"noclone"`
	} `yaml:"disco"`
}

// Init parses and prepares all config variables.
func Init() error {
	envconfig.MustProcess("", &Vars)

	if len(Vars.RegistryConfigurationPath) == 0 {
		homeDirPath, err := os.UserHomeDir()
		if err != nil {
			return fmt.Errorf("failed to get user home dir: %v", err)
		}
		Vars.RegistryConfigurationPath = path.Join(homeDirPath, defaultHomeDirDiscoConfigPath)
	}

	file, err := os.Open(Vars.RegistryConfigurationPath)
	if err != nil {
		return fmt.Errorf("failed to open config file: %v", err)
	}
	defer file.Close()
	log.WithField("config", Vars.RegistryConfigurationPath).Info("found configuration")

	DistributionConfig, err = configuration.Parse(file)
	if err != nil {
		return fmt.Errorf("error parsing %s: %v", Vars.RegistryConfigurationPath, err)
	}

	// Override/set IPFS config for the IPFS driver to consume.
	// Following the original naming convention in https://github.com/distribution/distribution.
	ipfsConfig := DistributionConfig.Storage["ipfs"]
	if ipfsConfig["router"] == nil {
		return errors.New("please specify 'router' in ipfs driver config")
	}
	if ipfsConfig["router"] != nil {
		file, _ = os.Open(Vars.RegistryConfigurationPath)
		defer file.Close()
		err = yaml.NewDecoder(file).Decode(&DiscoConfig)
		if err != nil {
			return err
		}
		Router = DiscoConfig.Storage.IPFS.Router
		Cache = DiscoConfig.Storage.IPFS.Cache
		if len(DiscoConfig.Storage.IPFS.Redirect) > 0 {
			RedirectTo, err = url.Parse(DiscoConfig.Storage.IPFS.Redirect)
			if err != nil {
				return err
			}
		}
	}

	return nil
}
