package main

import (
	bugLog "github.com/bugfixes/go-bugfixes/logs"
	"github.com/k8sdeploy/authorizer/internal/config"
)

func main() {
	bugLog.Local().Info("Starting Authorizer")

	cfg, err := config.Build()
	if err != nil {
		_ = bugLog.Errorf("build internal: %+v", err)
		return
	}

	if err := route(&cfg); err != nil {
		_ = bugLog.Errorf("route failed: %+v", err)
		return
	}
}

func route(cfg *config.Config) error {
	return nil
}
