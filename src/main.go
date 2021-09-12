package main

import (
	"encoding/json"
	"flag"
	"io/ioutil"
	"log"
	"net/http"
	"os"
	"strconv"
	"strings"
)

var token string

func handler(w http.ResponseWriter, r *http.Request) {
	client := &http.Client{}

	// Get user's servers
	request, _ := http.NewRequest("GET", "https://api.online.net/api/v1/server", nil)
	request.Header.Add("Authorization", "Bearer " + token)
	response, err := client.Do(request)

	if err != nil {
		log.Println("Error while doing GET /api/v1/servers")
		w.WriteHeader(503)
		return
	}

	if response.StatusCode != 200 {
		log.Println("Server return HTTP code " + strconv.Itoa(response.StatusCode) + " on GET /api/v1/servers")
		w.WriteHeader(503)
		return
	}

	body, _ := ioutil.ReadAll(response.Body)
	var servers []string
	_ = json.Unmarshal(body, &servers)


	// Display help
	_, _ = w.Write([]byte("# HELP dedibackup_active Indicates if you enabled the backup space storage on your panel\n"))
	_, _ = w.Write([]byte("# TYPE dedibackup_active gauge\n"))
	_, _ = w.Write([]byte("# HELP dedibackup_total_bytes Indicates the number of bytes on your backup plan\n"))
	_, _ = w.Write([]byte("# TYPE dedibackup_total_bytes gauge\n"))
	_, _ = w.Write([]byte("# HELP dedibackup_used_bytes Indicates the number of bytes used on your backup plan\n"))
	_, _ = w.Write([]byte("# TYPE dedibackup_used_bytes gauge\n"))
	_, _ = w.Write([]byte("# HELP dedibackup_total_files_number Indicates the number of files you can create on your backup plan\n"))
	_, _ = w.Write([]byte("# TYPE dedibackup_total_files_number gauge\n"))
	_, _ = w.Write([]byte("# HELP dedibackup_used_files_number Indicates the number of files you are using on your backup plan\n"))
	_, _ = w.Write([]byte("# TYPE dedibackup_used_files_number gauge\n"))



	// For each server, get backup status, and append metrics
	for _, k := range servers {
		tmp := strings.Split(k, "/")

		request, _ := http.NewRequest("GET", "https://api.online.net/api/v1/server/backup/"+tmp[len(tmp)-1], nil)
		request.Header.Add("Authorization", "Bearer " + token)
		response, err := client.Do(request)

		if err != nil {
			log.Println("Error while doing GET /api/v1/servers/backup/"+tmp[len(tmp)-1])
			continue
		}

		if response.StatusCode != 200 {
			log.Println("Server return HTTP code " + strconv.Itoa(response.StatusCode) + " on GET /api/v1/servers/backup"+tmp[len(tmp)-1])
			continue
		}

		body, _ := ioutil.ReadAll(response.Body)
		var data BackupStatus
		_ = json.Unmarshal(body, &data)

		active := "0"
		if data.Active {
			active = "1"
		}

		_, _ = w.Write([]byte("dedibackup_active{server_id=\"" + tmp[len(tmp)-1] + "\"} " + active + "\n"))
		_, _ = w.Write([]byte("dedibackup_total_bytes{server_id=\"" + tmp[len(tmp)-1] + "\"} " + strconv.Itoa(data.QuotaSpace) + "\n"))
		_, _ = w.Write([]byte("dedibackup_used_bytes{server_id=\"" + tmp[len(tmp)-1] + "\"} " + strconv.Itoa(data.QuotaSpaceUsed) + "\n"))
		_, _ = w.Write([]byte("dedibackup_total_files_number{server_id=\"" + tmp[len(tmp)-1] + "\"} " + strconv.Itoa(data.QuotaFiles) + "\n"))
		_, _ = w.Write([]byte("dedibackup_used_files_number{server_id=\"" + tmp[len(tmp)-1] + "\"} " + strconv.Itoa(data.QuotaFilesUsed) + "\n"))
	}
}

func main() {
	tokenPtr := flag.String("token", "", "Online API Token")
	portPtr := flag.Int("port", 9101, "Port to bind on")

	flag.Parse()
	token = *tokenPtr

	if token == "" {
		//Try to get from environment
		token = os.Getenv("DEDIBOX_API_TOKEN")
	}

	if token == "" {
		//Exit with error : no token provided
		log.Fatal("No token provided.")
	}

	http.HandleFunc("/metrics", handler)
	log.Fatal(http.ListenAndServe(":"+strconv.Itoa(*portPtr), nil))
}
