package controllers

import (
	"github.com/Artheriom/dedibox-backup-monitoring/internal/services"
	"github.com/sirupsen/logrus"
	"net/http"
)

func GetMetrics(w http.ResponseWriter, _ *http.Request) {
	data, err := services.GetMetrics()
	if err != nil {
		logrus.Warnf("An error occured while getting metrics. Service responded: %s", err.Error())
		w.WriteHeader(500)
		return
	}

	_, _ = w.Write([]byte(data))
}
