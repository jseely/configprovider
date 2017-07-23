package main

import (
	"context"
	"fmt"
	"os"

	"github.com/coreos/etcd/client"

	"github.com/jseely/configprovider"
	"github.com/jseely/logging"
	"github.com/jseely/logging/common"
	"github.com/jseely/logging/sinks"
)

func main() {
	logger := logging.NewWithApplicationScope("config-service", common.VERBOSE, sinks.NewWriterSink(os.Stdout, false))
	config := configprovider.EtcdConfigStoreConfig{
		EtcdConfig: client.Config{
			Endpoints: []string{"http://127.0.0.1:2379"},
			Transport: client.DefaultTransport,
		},
		DefaultRequestContext: context.Background(),
		Logger:                logger.WithApplicationScope("etcd-provider", common.VERBOSE),
	}
	store, err := configprovider.NewEtcdStore(config)
	if err != nil {
		panic(err.Error())
	}
	val, err := store.Get("test")
	if err != nil {
		panic(err.Error())
	}
	fmt.Printf("%+v", val)
}
