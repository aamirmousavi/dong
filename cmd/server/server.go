package server

import (
	"fmt"

	"github.com/aamirmousavi/dong/internal/context"
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
	appContext := context.NewContext(
		mongodb,
	)
	fmt.Printf("Server is running on %s\n", addr)
	return router.Run(
		appContext,
		addr,
	)
}
