package main

import (
	"context"

	"github.com/spyder01/lilium-go"
	liliumjob "github.com/spyder01/lilium-job"
)

func main() {
	cfg := lilium.LoadConfig("test.yaml")
	app := lilium.New(cfg, context.Background())
	router := lilium.NewRouter(app.Context)

	module_cfg, err := liliumjob.LoadLiliumJobsConfig("test.yaml")
	if err != nil {
		panic(err)
	}

	module := liliumjob.New(module_cfg)

	module.RegisterTask("test-job", func(ctx *lilium.AppContext) error {
		ctx.GetLogger().Info("Hello from job")
		return nil
	})

	app.UseModule(module)

	app.Start(router)
}
