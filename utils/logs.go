package utils

import (
	"os"
	"time"

	"github.com/olivere/elastic/v7"
	"github.com/sirupsen/logrus"
	"go.elastic.co/ecslogrus"
	"gopkg.in/sohlich/elogrus.v7"
)

type Logs struct {
	Event        string
	StatusCode   int
	ResponseTime time.Duration
	Method       string
	Request      interface{} `json:"request,omitempty"`
	URL          string      `json:"url,omitempty"`
	Message      string      `json:"message,omitempty"`
	Response     string
}

func CreateLog(data *Logs, types string) error {
	log := logrus.New()
	client, err := elastic.NewClient(
		elastic.SetURL("http://localhost:9200/"),
		elastic.SetBasicAuth("elastic", "admin-elk-1234"),
		elastic.SetSniff(false))
	if err != nil {
		log.Panic(err)
	}
	hook, err := elogrus.NewAsyncElasticHook(client, "localhost", logrus.DebugLevel, "mylog")
	if err != nil {
		log.Panic(err)
	}
	log.Hooks.Add(hook)

	log.SetFormatter(&ecslogrus.Formatter{
		DataKey: "labels",
	})

	if types == "warning" {
		log.WithFields(logrus.Fields{
			"event":         data.Event,
			"status_code":   data.StatusCode,
			"response_time": data.ResponseTime,
			"method":        data.Method,
			"request":       data.Request,
			"url":           data.URL,
			"response":      data.Response,
		}).Warn(data.Message)
	}

	if types == "info" {
		log.WithFields(logrus.Fields{
			"event":         data.Event,
			"status_code":   data.StatusCode,
			"response_time": data.ResponseTime,
			"method":        data.Method,
			"request":       data.Request,
			"url":           data.URL,
			"response":      data.Response,
		}).Info(data.Message)
	}

	if types == "error" {
		log.WithFields(logrus.Fields{
			"event":         data.Event,
			"status_code":   data.StatusCode,
			"response_time": data.ResponseTime,
			"method":        data.Method,
			"request":       data.Request,
			"url":           data.URL,
			"response":      data.Response,
		}).Error(data.Message)
	}

	log.Out = os.Stdout

	return nil

}
