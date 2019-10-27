package e2etest

import (
	"fmt"
	"io/ioutil"
	"log"
	"net/http"
	"os"
	"os/exec"
	"prometheus-metrics-exporter/internal/configuration"
	"strings"
	"testing"
	"time"
)

func TestMain(m *testing.M) {
	//setup()
	code := m.Run()
	//shutdown()
	os.Exit(code)
}

func setup(){
	//cmd := "docker-compose -f ../../test/test-docker-compose.yml up --build -d"
	err := exec.Command("docker-compose" ,"-f", "../../test/test-docker-compose.yml", "up", "--build", "-d").Run()

	if err != nil {
		errStr := fmt.Sprintf("Failed to start containers: %s", err)
		log.Fatal(errStr)
	}
}

func shutdown(){
//	cmd := "docker-compose -f ../../test/test-docker-compose.yml down --rmi"
	err := exec.Command("docker-compose", "-f", "../../test/test-docker-compose.yml", "down").Run()

	if err != nil {
		errStr := fmt.Sprintf("Failed to stop containers: %s", err)
		log.Fatal(errStr)
	}
}

func TestA(t *testing.T) {
	const configFile string = "e2e_test_config.json"
	config, err := configuration.LoadConfig(configFile)
	if err != nil {
		errStr := fmt.Sprintf("Error loading test configuration: %s", configFile)
		log.Fatal(errStr)
	}

	values := []string{
		config.ScrapeTargets[0].JobName,
		config.ScrapeTargets[0].Labels[0].Name,
		config.ScrapeTargets[0].Labels[0].Value,
		config.ScrapeTargets[0].Metrics[0].Name,
		config.ScrapeTargets[0].Metrics[0].Help,
		config.ScrapeTargets[0].Metrics[0].InstrumentType,
	}

	client := http.Client{
		Timeout: 15 * time.Second,
	}

	url := config.PushGatewayUrl + "/metrics"
	response, err := client.Get(url)

	if err != nil {
		errStr := fmt.Sprintf("http client failed: %s", err)
		t.Fatal(errStr)
	}

	body, err := ioutil.ReadAll(response.Body)

	defer func() {
		if _err := response.Body.Close(); _err != nil {
			err = _err
		}
	}()

	if err != nil {
		errStr := fmt.Sprintf("reading response body failed: %s", err)
		t.Fatal(errStr)
	}

	bodyStr := string(body)

	for _, item := range values {
		if ! strings.Contains(bodyStr, item) {
			errStr := fmt.Sprintf("metrics did not contain %s", item)
			t.Fatal(errStr)
		}
	}

	log.Println("Ending tests")
}

func TestB(t *testing.T) {

}