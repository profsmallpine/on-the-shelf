package procedures

import (
	"github.com/profsmallpine/on-the-shelf/db/models"
	"github.com/profsmallpine/on-the-shelf/domain"
)

var (
	queries  = &models.Queries{}
	services = domain.Services{}
)

func New(allowedEmails []string, q *models.Queries, s domain.Services) domain.Procedures {
	queries = q
	services = s

	return domain.Procedures{
		User: &User{allowedEmails},
	}
}
