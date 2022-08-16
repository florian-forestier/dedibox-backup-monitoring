# Dedibox backup monitoring

A simple tool who reads API from Online.net and parse them into a Prometheus-compatible format.

## How to use it ?

### From DockerHub

`docker run -p 9101:9101 -e DEDIBOX_API_TOKEN=<YOUR_TOKEN> artheriom/dedibox-backup-monitoring:v1.0.0`

### From sources

* For Docker: `docker build -t="dedibackup_poller:latest" .`
* Golang: `go build -o dedibackup_poller github.com/Artheriom/dedibox-backup-monitoring/cmd`

### Parameters

There is only three parameters:

* `DEDIBOX_API_TOKEN` (or `-token` command-line argument): Your Dedibox API token. You can generate one [here](https://console.online.net/fr/api/access).
    * Make sure to **never** share this token!
* `-port`: Specify a port to bind on. Default is `9101`, as `9100` is used by Prometheus Node Exporter.
* `-apiUrl` : Specify a custom URL for Online API (eg. for testing, or if you are behind a proxy). Default is `https://api.online.net/api`.

Metrics will be available under `/metrics` path.

## Example

Data returned (example):

```prometheus
dedibackup_active{server_id="155976"} 1
dedibackup_total_bytes{server_id="155976"} 100000000000
dedibackup_used_bytes{server_id="155976"} 4994340620
dedibackup_total_files_number{server_id="155976"} 1000
dedibackup_used_files_number{server_id="155976"} 240
```

![](https://raw.githubusercontent.com/Artheriom/dedibox-backup-monitoring/main/_github/example.png)

## About

### Contribute

You're absolutely free (and welcome!) to contribute! Fork it, and submit a PR!

### Licence

Licensed under MIT licence. Originally crafted with ❤️ by Artheriom.️

### Support

Even as this program is published "as-is", without warranty, I would be glad to help you if you encounter issues with
this program. Ask on Github.com with an issue, or [tweet me](https://twitter.com/Artheriom)
