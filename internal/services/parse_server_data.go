package services

import (
	"bytes"
	"fmt"
	"github.com/Artheriom/dedibox-backup-monitoring/internal/structures"
)

func parseServerData(server structures.BackupStatus, buffer *bytes.Buffer) (err error) {
	active := 0
	if server.Active {
		active = 1
	}

	values := map[string]int{
		"dedibackup_active":             active,
		"dedibackup_total_bytes":        server.QuotaSpace,
		"dedibackup_used_bytes":         server.QuotaSpaceUsed,
		"dedibackup_total_files_number": server.QuotaFiles,
		"dedibackup_used_files_number":  server.QuotaFilesUsed,
	}

	for key, value := range values {
		_, err = fmt.Fprintf(buffer, "%s{server_id=\"%s\"} %d\n", key, server.ServerId, value)
		if err != nil {
			return
		}
	}
	return
}
