package main
import (
	"flag"
	"github.com/Ninja2205/CSI-S3/pkg/driver"
	"log"
	"os"
)

func init() {
	flag.Set("logtostderr", "true")
}

var (
	endpoint = flag.String("endpoint", "unix://tmp/csi.sock", "CSI endpoint")
	nodeID   = flag.String("nodeid", "", "node id")
)

func main() {
	flag.Parse()
	driver, err := driver.New(*nodeID, *endpoint)
	if err != nil {
		log.Fatal(err)
	}
	driver.Run()
	os.Exit(0)
}
