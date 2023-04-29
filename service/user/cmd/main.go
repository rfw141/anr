package main

import (
	"flag"
	"github.com/rfw141/anr/gen/core"
	"github.com/rfw141/anr/internal"
	"github.com/rfw141/anr/pkg/app"
	"github.com/rfw141/anr/pkg/server/grpc"
	"github.com/spf13/viper"
	"log"
)

// go build -ldflags "-X main.Version=x.y.z"
var (
	// Name is the name of the compiled software.
	Name string = "user"
	// Version is the version of the compiled software.
	Version string
	// flagconf is the config flag.
	flagconf string
)

func init() {
	flag.StringVar(&flagconf, "conf", "config", "config path, eg: -conf config.yaml")
}

func main() {
	flag.Parse()

	viper.AddConfigPath(flagconf)
	viper.SetConfigType("toml")
	if err := viper.ReadInConfig(); err != nil {
		panic(err)
	}
	var cfg core.Config
	if err := viper.Unmarshal(&cfg); err != nil {
		panic(err)
	}

	log.Printf("%+v", cfg)

	svcCfg := cfg.Services[Name]
	if svcCfg == nil {
		svcCfg = &core.Config_Service{}
	}

	svr, cleanup, err := initApp(svcCfg, cfg.Registry)
	if err != nil {
		panic(err)
	}
	defer cleanup()

	if err := svr.Run(); err != nil {
		panic(err)
	}
}

func newApp(server *grpc.Server, registrar internal.Registrar) *app.App {
	return app.New(
		app.Name(Name),
		app.Version(Version),
		app.Server(server),
		app.Registrar(registrar),
	)
}
