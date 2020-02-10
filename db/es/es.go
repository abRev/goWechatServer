package es

import (
	"crypto/tls"
	"github.com/elastic/go-elasticsearch"
	"github.com/spf13/viper"
	"net"
	"net/http"
	"time"
)

var es *elasticsearch.Client

func Init() error {
	cfg := elasticsearch.Config{
		Addresses: []string{
			viper.GetString("common.es.db") + ":" + viper.GetString("common.es.port"),
		},
		Transport: &http.Transport{
			MaxIdleConnsPerHost:   10,
			ResponseHeaderTimeout: time.Second,
			DialContext:           (&net.Dialer{Timeout: time.Second}).DialContext,
			TLSClientConfig: &tls.Config{
				MinVersion: tls.VersionTLS11,
			},
		},
	}

	esCon, err := elasticsearch.NewClient(cfg)
	if err != nil {
		return err
	}
	es = esCon
	return nil
}

func GetES() *elasticsearch.Client {
	if es != nil {
		return es
	}
	return nil
}
