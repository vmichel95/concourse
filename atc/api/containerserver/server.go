package containerserver

import (
	"time"

	"github.com/concourse/concourse/atc/handles"

	"code.cloudfoundry.org/clock"
	"code.cloudfoundry.org/lager"
	"github.com/concourse/concourse/atc/creds"
	"github.com/concourse/concourse/atc/db"
	"github.com/concourse/concourse/atc/gc"
	"github.com/concourse/concourse/atc/worker"
)

type Server struct {
	logger lager.Logger

	workerClient            worker.Client
	secretManager           creds.Secrets
	varSourcePool           creds.VarSourcePool
	interceptTimeoutFactory InterceptTimeoutFactory
	interceptUpdateInterval time.Duration
	containerRepository     db.ContainerRepository
	containerSyncer         handles.Syncer
	destroyer               gc.Destroyer
	clock                   clock.Clock
}

func NewServer(
	logger lager.Logger,
	workerClient worker.Client,
	secretManager creds.Secrets,
	varSourcePool creds.VarSourcePool,
	interceptTimeoutFactory InterceptTimeoutFactory,
	interceptUpdateInterval time.Duration,
	containerRepository db.ContainerRepository,
	containerSyncer handles.Syncer,
	destroyer gc.Destroyer,
	clock clock.Clock,
) *Server {
	return &Server{
		logger:                  logger,
		workerClient:            workerClient,
		secretManager:           secretManager,
		varSourcePool:           varSourcePool,
		interceptTimeoutFactory: interceptTimeoutFactory,
		interceptUpdateInterval: interceptUpdateInterval,
		containerRepository:     containerRepository,
		containerSyncer:         containerSyncer,
		destroyer:               destroyer,
		clock:                   clock,
	}
}
