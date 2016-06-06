package kafka

import (
	"crypto/tls"
	"log"
	"net/url"
)

type Config struct {
	BrokerAddresses []string
	TLSConfig       *tls.Config
}

func (c *Config) getBrokers() []string {
	addrs := make([]string, len(c.BrokerAddresses))
	for i, v := range c.BrokerAddresses {
		u, err := url.Parse(v)
		if err != nil {
			log.Fatal(err)
		}
		addrs[i] = u.Host
	}
	return addrs
}
