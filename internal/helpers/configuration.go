package helpers

import (
	"flag"
	"github.com/sirupsen/logrus"
	"os"
)

var DediboxToken string
var Port int

func init() {
	tokenPtr := flag.String("token", "", "Online API Token")
	portPtr := flag.Int("port", 9101, "Port to bind on")

	flag.Parse()
	DediboxToken = *tokenPtr
	Port = *portPtr

	if DediboxToken == "" {
		//Try to get from environment
		DediboxToken = os.Getenv("DEDIBOX_API_TOKEN")
	}

	if DediboxToken == "" {
		//Exit with error : no token provided
		logrus.Fatalln("No token provided.")
	}
}
