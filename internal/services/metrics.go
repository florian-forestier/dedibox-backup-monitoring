package services

import (
	"bytes"
	"github.com/Artheriom/dedibox-backup-monitoring/internal/connectors"
	"github.com/Artheriom/dedibox-backup-monitoring/internal/helpers"
	"github.com/Artheriom/dedibox-backup-monitoring/internal/structures"
	"strings"
)

func GetMetrics() (result string, err error) {
	// Variable instantiation
	var buffer bytes.Buffer
	headers := map[string]string{"Authorization": "Bearer " + helpers.DediboxToken}

	// Write Prometheus headers : HELP and TYPE fields.
	err = writePrometheusHeaders(&buffer)
	if err != nil {
		return
	}

	// Retrieve server list
	servers, err := connectors.GetURL[[]string](helpers.DediboxApi+"/v1/server", headers, 200)
	if err != nil {
		return
	}

	// For each server, get backup status, and append metrics
	for _, k := range servers {
		serverId := strings.Split(k, "/")

		data, err := connectors.GetURL[structures.BackupStatus](helpers.DediboxApi+"/v1/server/backup/"+serverId[len(serverId)-1], headers, 200)
		if err != nil {
			return "", err
		}

		data.ServerId = serverId[len(serverId)-1]

		err = parseServerData(data, &buffer)
		if err != nil {
			return "", err
		}
	}

	return buffer.String(), nil
}
