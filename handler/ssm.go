package handler

import (
	"context"
	"fmt"
	"net"
	"strconv"

	"github.com/charmbracelet/log"

	"github.com/aws/aws-sdk-go-v2/config"
	"github.com/mmmorris1975/ssm-session-client/ssmclient"
)

// Create SSM Session Output: returns a StreamURL and Token to open a WebSocket connection (SSM)
func createSession(t string, h string, p string, lp string) {

	var port int

	lp_int, err := strconv.Atoi(lp)
	if err != nil {
		// ... handle error
		log.Error("Localport conversion Error: ", err.Error())
	}

	cfg, err := config.LoadDefaultConfig(context.TODO())
	if err != nil {
		log.Error("configuration error: " + err.Error())
	}

	tgt, err := ssmclient.ResolveTarget(t, cfg)
	port, err = net.LookupPort("tcp", p)
	if err != nil {
		log.Fatal("Port Lookup Error: ", err.Error())
	}

	log.Info(fmt.Sprintf("Opening connection @ localhost:%s", lp))
	in := ssmclient.PortForwardingInput{
		Target:     tgt,
		RemotePort: port,
		LocalPort:  lp_int,
	}

	// Alternatively, can be called as ssmclient.PortluginSession(cfg, tgt) to use the AWS-managed SSM session client code
	log.Fatal(ssmclient.PortForwardingSession(cfg, &in))
}
