package services

import (
	"bytes"
	"fmt"
)

func writePrometheusHeaders(w *bytes.Buffer) error {
	fields := map[string]string{
		"dedibackup_active":             "Indicates if you enabled the backup space storage on your panel",
		"dedibackup_total_bytes":        "Indicates the number of bytes on your backup plan",
		"dedibackup_used_bytes":         "Indicates the number of bytes used on your backup plan",
		"dedibackup_total_files_number": "Indicates the number of files you can create on your backup plan",
		"dedibackup_used_files_number":  "Indicates the number of files you are using on your backup plan",
	}

	for key, value := range fields {
		_, err := fmt.Fprintf(w, "# HELP %s %s\n# TYPE %s gauge\n", key, value, key)
		if err != nil {
			return err
		}
	}

	return nil
}
