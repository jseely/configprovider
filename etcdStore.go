package configprovider

import (
	"context"

	"github.com/coreos/etcd/client"
	"github.com/jseely/logging"
)

type EtcdConfigStoreConfig struct {
	EtcdConfig            client.Config
	DefaultRequestContext context.Context
	Logger                logging.Logger
}

type EtcdConfigStore struct {
	client.Client
	reqContext context.Context
	logger     logging.Logger
}

var getRecursiveOptions = &client.GetOptions{Recursive: true, Sort: false, Quorum: true}

func NewEtcdStore(config EtcdConfigStoreConfig) (*EtcdConfigStore, error) {
	c, err := client.New(config.EtcdConfig)
	if err != nil {
		return nil, err
	}

	return &EtcdConfigStore{Client: c, reqContext: config.DefaultRequestContext, logger: config.Logger}, nil
}

func (s *EtcdConfigStore) Get(path ...string) (interface{}, error) {
	kApi := client.NewKeysAPI(s)
	res, err := kApi.Get(s.reqContext, pathToEtcdKey(path), getRecursiveOptions)
	if err != nil {
		return nil, err
	}
	s.logger.Verbose("Get response: {response}", res)
	return nil, nil
}

func (s *EtcdConfigStore) GetRaw(path ...string) (interface{}, error) {
	return nil, nil
}

func (s *EtcdConfigStore) Set(value interface{}, path ...string) error {
	return nil
}

func pathToEtcdKey(path []string) string {
	key := ""
	for _, v := range path {
		key += "/" + v
	}
	return key
}
