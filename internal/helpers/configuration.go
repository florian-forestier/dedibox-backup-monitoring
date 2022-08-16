package helpers

import (
	"flag"
	"github.com/sirupsen/logrus"
	"os"
	"strings"
)

var DediboxToken string
var Port int
var DediboxApi string

func init() {
	tokenPtr := flag.String("token", "", "Online API Token")
	apiPtr := flag.String("apiUrl", "https://api.online.net/api", "Online API base URL")
	portPtr := flag.Int("port", 9101, "Port to bind on")

	flag.Parse()
	DediboxToken = *tokenPtr
	Port = *portPtr
	DediboxApi = *apiPtr

	if DediboxToken == "" {
		//Try to get from environment
		DediboxToken = os.Getenv("DEDIBOX_API_TOKEN")
	}

	if DediboxToken == "" {
		//Exit with error : no token provided
		logrus.Fatalln("No token provided.")
	}

	if strings.HasSuffix(DediboxApi, "/") {
		DediboxApi = DediboxApi[0 : len(DediboxApi)-1]
	}
}
