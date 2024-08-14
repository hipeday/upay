package application

import (
	"github.com/hipeday/upay/internal/logging"
	"github.com/hipeday/upay/pkg/config"
	"time"
)

func Setup(cfg *config.Application) {
	if cfg != nil {
		if cfg.TimeZone != "" {
			location, err := time.LoadLocation(cfg.TimeZone)
			logging.Logger().Infof("Load timezone: %s", location)
			if err != nil {
				logging.Logger().Fatalf("Load timezone failed: %v", err)
			}
		}
	}
}
