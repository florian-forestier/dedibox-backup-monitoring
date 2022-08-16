package main

import (
	"github.com/Artheriom/dedibox-backup-monitoring/internal/controllers"
	"github.com/Artheriom/dedibox-backup-monitoring/internal/helpers"
	"github.com/sirupsen/logrus"
	"net/http"
	"strconv"
)

func main() {
	http.HandleFunc("/metrics", controllers.GetMetrics)

	logrus.Infof("Server is starting on port %d", helpers.Port)
	logrus.Fatal(http.ListenAndServe(":"+strconv.Itoa(helpers.Port), nil))
}
