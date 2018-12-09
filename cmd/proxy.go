package main

import (
	"context"

	"github.com/deislabs/osiris/pkg/metrics/proxy/proxy"
	"github.com/deislabs/osiris/pkg/version"
	"github.com/golang/glog"
)

func runProxy(ctx context.Context) {
	glog.Infof(
		"Starting Osiris Proxy -- version %s -- commit %s",
		version.Version(),
		version.Commit(),
	)

	cfg, err := proxy.GetConfigFromEnvironment()
	if err != nil {
		glog.Fatalf("Error retrieving proxy configuration: %s", err)
	}

	// Run the proxy
	proxy.NewProxy(cfg).Run(ctx)
}