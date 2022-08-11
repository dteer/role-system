package settings

import (
	"fmt"
	"log"

	"github.com/cactus/go-statsd-client/v5/statsd"
)

func StatsdConnect() statsd.Statter {
	address := fmt.Sprintf("%s:%d", Conf.Statsd.Host, Conf.Statsd.Port)
	config := &statsd.ClientConfig{
		Address: address,
		Prefix:  Conf.Statsd.Prefix,
	}
	client, err := statsd.NewClientWithConfig(config)
	if err != nil {
		log.Fatal(err)
	}
	return client
}
