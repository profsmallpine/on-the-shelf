package app

import (
	"context"
	"embed"
	"log"
	"os"
	"strings"

	"github.com/jackc/pgx/v5/pgxpool"
	"github.com/profsmallpine/on-the-shelf/db/models"
	"github.com/profsmallpine/on-the-shelf/domain"
	"github.com/profsmallpine/on-the-shelf/http/web"
	"github.com/profsmallpine/on-the-shelf/procedures"
	"github.com/profsmallpine/on-the-shelf/services/auth"
	"github.com/xy-planning-network/trails"
	"github.com/xy-planning-network/trails/postgres"
	"github.com/xy-planning-network/trails/ranger"
)

func New(logging *log.Logger, files embed.FS) (*ranger.Ranger, error) {
	isMaintMode := os.Getenv("MAINTENANCE_MODE") == "true"

	cfg := ranger.Config[*domain.User]{
		FS:         files,
		MaintMode:  isMaintMode,
		Migrations: []postgres.Migration{},
	}

	rng, err := ranger.New(cfg)
	if err != nil {
		return nil, err
	}

	ctx := context.Background()
	pool, err := pgxpool.New(ctx, "host=localhost port=5432 dbname=on_the_shelf_dev_db user=on_the_shelf_dev_db_user")
	if err != nil {
		return nil, err
	}
	queries := models.New(pool)

	services := domain.Services{
		Auth: auth.NewService(
			trails.EnvVarOrString(
				"BASE_URL",
				"http://localhost:3000",
			) + "/auth/google/callback",
		),
		Logger: rng.Logger,
	}

	procs := procedures.New(strings.Split(os.Getenv("ALLOWED_EMAILS"), ","), queries, services)

	h := &web.Handler{
		Procedures: procs,
		Queries:    queries,
		Ranger:     rng,
		Services:   services,
	}

	h.RegisterRoutes()

	return rng, nil
}
