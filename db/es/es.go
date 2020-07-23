package es

import (
	"crypto/tls"
	"net"
	"net/http"
	"time"

	"github.com/elastic/go-elasticsearch"
	"github.com/spf13/viper"
)

var es *elasticsearch.Client

func init() {
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
		panic(err)
	}
	es = esCon
}

// GetES 获取es实例
func GetES() *elasticsearch.Client {
	if es != nil {
		return es
	}
	return nil
}
