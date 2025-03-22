package domain

import "github.com/xy-planning-network/trails/logger"

type Services struct {
	Auth   AuthService
	Logger logger.Logger
}
