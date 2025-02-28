// Code generated by Wire. DO NOT EDIT.

//go:generate go run -mod=mod github.com/google/wire/cmd/wire
//go:build !wireinject
// +build !wireinject

package wire

import (
	"github.com/google/wire"
	"github.com/la-plas-growth/go-docker-ai/cmd/service/dockerfile"
	"github.com/la-plas-growth/go-docker-ai/cmd/service/lint"
	"github.com/la-plas-growth/go-docker-ai/env"
	"github.com/la-plas-growth/go-docker-ai/lib/log"
)

// Injectors from wire.go:

// add all commands
func InitLint() lint.IService {
	configuration := env.NewConfiguration()
	sugaredLogger := zap_log.NewLogger(configuration)
	iService := lint.NewService(configuration, sugaredLogger)
	return iService
}

// add all commands
func InitDockerfileService() dockerfile.IService {
	configuration := env.NewConfiguration()
	sugaredLogger := zap_log.NewLogger(configuration)
	iService := dockerfile.NewService(configuration, sugaredLogger)
	return iService
}

// wire.go:

var CommonSet = wire.NewSet(env.NewConfiguration, zap_log.NewLogger)
