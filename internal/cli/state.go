package cli

import (
	"github.com/zepetriko/go_aggregator/internal/config"
	"github.com/zepetriko/go_aggregator/internal/database"
)

type State struct {
	Db     *database.Queries
	Config *config.Config
}
