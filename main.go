package main

import (
	"encoding/json"
	"errors"
	"strings"
	"time"

	"bytes"
	"fmt"

	"os/exec"

	"github.com/k0kubun/pp"
	log "github.com/sirupsen/logrus"
)

func FindCommandPath(name string, env []string) (bool, string) {
	cmd := exec.Command("/bin/sh", "-c", fmt.Sprintf("command -v %s", name))
	cmd.Env = env
	stdout, err := cmd.Output()
	if err != nil && !strings.Contains(err.Error(), `: no child processes`) {
		msg := fmt.Sprintf("\n** Failed to find Command %s: %s\n\ncmd=%s | env=%v | stdout=%s | \n", name, err.Error(), cmd, cmd.Env, stdout)
		log.Error(msg)
		return false, ``
	} else {
		err = cmd.Start()
		if err != nil {
			err = errors.New("COMMAND_ERROR")
		}
		defer cmd.Wait()
		cmd_path := strings.Replace(string(stdout), "\n", "", -1)
		return true, cmd_path
	}
}

func execute_speedtest_cli(exec_path string) (string, error) {
	cmd := exec.Command(exec_path, "--json")
	var stdout bytes.Buffer
	cmd.Stdout = &stdout
	err := cmd.Run()
	if err != nil {
		log.Fatalf("cmd.Run() failed with %s\n", err)
	}
	return string(stdout.Bytes()), nil
}

func main() {
	res, err := get_speedtest_result()
	if err != nil {
		log.Fatal(err)
	}
	pp.Println(res)

}
func get_speedtest_result() (*SpeedTestResult, error) {
	var res SpeedTestResult
	ok, exec_path := FindCommandPath(`speedtest-cli`, []string{})
	if !ok {
		log.Fatal(`speedtest-cli not found`)
	}

	started := time.Now()
	speedtest_json, err := execute_speedtest_cli(exec_path)
	if err != nil {
		log.Fatal(err)
	}

	uerr := json.Unmarshal([]byte(speedtest_json), &res)
	if uerr != nil {
		log.Fatal(uerr)
	}
	res.Duration = time.Since(started)
	return &res, nil

}
