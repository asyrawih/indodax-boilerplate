package service

import "bitbucket.org/btcid/startrack/internal/config"

type (
	ServiceContext struct {
		// Can Add More the Context In here for access the service context like redis, mysql , logger  or etc
		Config *config.Config
	}
)

func NewServiceContext(c *config.Config) *ServiceContext {
	return &ServiceContext{
		Config: c,
	}
}
