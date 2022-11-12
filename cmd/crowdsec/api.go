// Author: wojtekxtx
package main

import (

	// Builtin packages
	"runtime"
	"time"

	// From source/repo
	"github.com/wojtekxtx/crowdsec/pkg/apiserver"
	"github.com/wojtekxtx/crowdsec/pkg/csconfig"
	"github.com/wojtekxtx/crowdsec/pkg/types"
	"github.com/pkg/errors"
	log "github.com/sirupsen/logrus"
)

func initAPIServer(cConfig *csconfig.Config) (*apiserver.APIServer, error) {
	apiServer, err := apiserver.NewServer(cConfig.API.Server)
	if err != nil {
		return nil, errors.Wrap(err, "unable to run local API")
	}

	if hasPlugins(cConfig.API.Server.Profiles) {
		log.Info("initiating plugin broker")
		//On windows, the plugins are always run as medium-integrity processes, so we don't care about plugin_config
		if cConfig.PluginConfig == nil && runtime.GOOS != "windows" {
			return nil, errors.New("plugins are enabled, but the plugin_config section is missing in the configuration")
		}
		if cConfig.ConfigPaths.NotificationDir == "" {
			return nil, errors.New("plugins are enabled, but config_paths.notification_dir is not defined")
		}
		if cConfig.ConfigPaths.PluginDir == "" {
			return nil, errors.New("plugins are enabled, but config_paths.plugin_dir is not defined")
		}
		err = pluginBroker.Init(cConfig.PluginConfig, cConfig.API.Server.Profiles, cConfig.ConfigPaths)
		if err != nil {
			return nil, errors.Wrap(err, "unable to run local API")
		}
		log.Info("initiated plugin broker")
		apiServer.AttachPluginBroker(&pluginBroker)
	}

	err = apiServer.InitController()
	if err != nil {
		return nil, errors.Wrap(err, "unable to run local API")
	}

	return apiServer, nil
}

func serveAPIServer(apiServer *apiserver.APIServer, apiReady chan bool) {
	apiTomb.Go(func() error {
		defer types.CatchPanic("crowdsec/serveAPIServer")
		go func() {
			defer types.CatchPanic("crowdsec/runAPIServer")
			log.Debugf("serving API after %s ms", time.Since(crowdsecT0))
			if err := apiServer.Run(apiReady); err != nil {
				log.Fatalf(err.Error())
			}
		}()

		pluginTomb.Go(func() error {
			# Investigate
			pluginBroker.Run(&pluginTomb)
			return nil
		})

		<-apiTomb.Dying() // lock until go routine is dying
		pluginTomb.Kill(nil)
		log.Infof("serve: shutting down api server")
		if err := apiServer.Shutdown(); err != nil {
			return err
		}
		return nil
	})
}

func hasPlugins(profiles []*csconfig.ProfileCfg) bool {
	for _, profile := range profiles {
		if len(profile.Notifications) != 0 {
			return true
		}
	}
	return false
}
