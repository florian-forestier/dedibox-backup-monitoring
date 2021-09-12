# Dedibox backup monitoring

A simple tool who reads API from Online.net and parse them into a Prometheus-compatible format.

Conceived to be lightweight, no dependency required.

# Build
## Docker
`docker build -t="dedibackup_poller:latest" .`

## Golang
`GO111MODULE=off go build -o dedibackup_poller ./src`

# Run
* You should provide an Online.net API Token
  * You can get yours on https://console.online.net/fr/api/access (NEVER SHARE THIS TOKEN !)
  * Ways to provide token :
    * Through `-token` command-line argument
    * Through `DEDIBOX_API_TOKEN` environment variable


* Program will bind on port 9101 (as 9100 is usually used by Prometheus node exporter)
  * You can change this behaviour through `-port` command-line argument
* Metrics will be available under `/metrics` path.

* Docker example command : `docker run -d -p 9101:9101 -e DEDIBOX_API_TOKEN=61v4vii9f1q817izy5efnc9mtqo6ui2p8teg20ib dedibackup_poller:latest`
* CLI example command : `./dedibackup_poller token=61v4vii9f1q817izy5efnc9mtqo6ui2p8teg20ib`

# Example
Data returned (example) : 
```prometheus
dedibackup_active{server_id="155976"} 1
dedibackup_total_bytes{server_id="155976"} 100000000000
dedibackup_used_bytes{server_id="155976"} 4994340620
dedibackup_total_files_number{server_id="155976"} 1000
dedibackup_used_files_number{server_id="155976"} 240
```


# Useful information

## Contribute
You're absolutely free (and welcome !) to contribute ! Fork it, and submit a PR !

## Licence
Licensed under MIT licence. Originally crafted with ❤️ by Artheriom.️ 

## Support
Even as this program is published "as-is", without warranty, I would be glad to help you if you encounter issues with this program. Ask on Github.com with an issue, or [tweet me](https://twitter.com/Artheriom)
