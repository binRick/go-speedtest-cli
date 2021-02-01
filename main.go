package main

import (
	"encoding/json"
	//"errors"
	"bytes"
	"fmt"
	//"os"
	"os/exec"
	//	"github.com/jessevdk/go-flags"
	log "github.com/sirupsen/logrus"
	//	"github.com/tidwall/sjson"
	//	"gopkg.in/yaml.v2"
	//	"io/ioutil"
	//	"os"
	//	"text/tabwriter"
	//	"time"
)

func execute_speedtest_cli() (string, error) {
	cmd := exec.Command("speedtest-cli", "--json")
	var stdout bytes.Buffer
	cmd.Stdout = &stdout
	err := cmd.Run()
	if err != nil {
		log.Fatalf("cmd.Run() failed with %s\n", err)
	}
	return string(stdout.Bytes()), nil
}

func GetJSONString(obj interface{}, ignoreFields ...string) (string, error) {
	toJson, err := json.Marshal(obj)
	if err != nil {
		return "", err
	}

	if len(ignoreFields) == 0 {
		return string(toJson), nil
	}

	toMap := map[string]interface{}{}
	json.Unmarshal([]byte(string(toJson)), &toMap)

	for _, field := range ignoreFields {
		delete(toMap, field)
	}

	toJson, err = json.Marshal(toMap)
	if err != nil {
		return "", err
	}

	return string(toJson), nil
}

type SpeedTestResult struct {
	bytes_sent int
	bytes_recv int
	dur_ms     int
}

func main() {
	speedtest_json, err := execute_speedtest_cli()
	if err != nil {
		log.Fatal(err)
	}
	fmt.Println(speedtest_json)
}
