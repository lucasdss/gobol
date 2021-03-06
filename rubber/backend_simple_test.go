package rubber

import (
	"fmt"
	"io/ioutil"
	"net/http"
	"testing"
	"time"

	"github.com/Sirupsen/logrus"
)

// This mostly tests compilation
var _ Backend = &singleServerBackend{}

func testSimpleBackend() *singleServerBackend {
	logger := logrus.New()
	logger.Out = ioutil.Discard
	return &singleServerBackend{
		log:     logger,
		nodes:   []string{fmt.Sprintf("%s:9200", master)},
		timeout: time.Minute,
		client: &http.Client{
			Timeout: time.Minute,
		},
	}
}

func TestSimpleIntegration(t *testing.T) {
	logger := logrus.New()
	logger.Out = ioutil.Discard
	genericBackendTest(t, testSimpleBackend())
}
