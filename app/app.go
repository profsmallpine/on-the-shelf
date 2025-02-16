package app

import (
	"embed"
	"log"
	"os"

	"github.com/profsmallpine/books/domain"
	"github.com/profsmallpine/books/http/web"
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

	h := &web.Handler{
		Ranger: rng,
	}

	h.Router()

	return rng, nil
}
