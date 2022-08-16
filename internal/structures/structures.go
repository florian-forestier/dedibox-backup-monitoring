package structures

type BackupStatus struct {
	Active         bool   `json:"active"`
	QuotaSpace     int    `json:"quota_space"`
	QuotaSpaceUsed int    `json:"quota_space_used"`
	QuotaFiles     int    `json:"quota_files"`
	QuotaFilesUsed int    `json:"quota_files_used"`
	ServerId       string `json:"-"`
}
