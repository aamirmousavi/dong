package server

import (
	"github.com/aamirmousavi/dong/internal/database/mongodb"
	"github.com/aamirmousavi/dong/internal/router"
)

func run(
	mongodbAddr string,
	addr string,
) error {
	mongodb, err := mongodb.NewHandler(mongodbAddr)
	if err != nil {
		return err
	}
	return router.Run(mongodb, addr)
}
