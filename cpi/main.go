package main

import (
	"flag"
	"os"

	"github.com/fsouza/go-dockerclient"

	"./action"
	cfg "./config"

	boshlog "github.com/cloudfoundry/bosh-agent/logger"
	boshsys "github.com/cloudfoundry/bosh-agent/system"
	bwcdisp "github.com/cppforlife/bosh-warden-cpi/api/dispatcher"
	bwctrans "github.com/cppforlife/bosh-warden-cpi/api/transport"
)

func main() {
	logger := boshlog.NewWriterLogger(boshlog.LevelDebug, os.Stderr, os.Stderr)
	fs := boshsys.NewOsFileSystem(logger)

	var configPath string
	flag.StringVar(&configPath, "c", "", "config path")
	flag.Parse()

	config, err := cfg.LoadConfig(configPath, fs)
	if err != nil {
		fail(logger, err)
	}

	client, _ := docker.NewClient(config.SocketPath)
	actionFactory := action.NewFactory(client, config, fs)
	caller := bwcdisp.NewJSONCaller()

	dispatcher := bwcdisp.NewJSON(actionFactory, caller, logger)
	cli := bwctrans.NewCLI(os.Stdin, os.Stdout, dispatcher, logger)

	err = cli.ServeOnce()
	if err != nil {
		fail(logger, err)
	}
}

func fail(logger boshlog.Logger, err error) {
	logger.Error("main", err.Error())
	os.Exit(1)
}
