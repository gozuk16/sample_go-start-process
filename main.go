package main

import (
	"bufio"
	"fmt"
	"log"
	"os"
	"os/exec"
	"os/signal"
	"strings"
)

func startProc(done chan<- error) {
	startDir := "/Users/gozu/projects/jetty-distribution-9.4.43.v20210629/demo-base"
	startCmd := "java"
	startArgs := strings.Fields("-jar ../start.jar STOP.PORT=28282 STOP.KEY=secret")
	//startCmd := "ls"
	//startArgs := strings.Fields("-l ..")
	//stopCmd := "java -jar ../start.jar STOP.PORT=28282 STOP.KEY=secret --stop"

	// process start
	cmd := exec.Command(startCmd, startArgs...)
	cmd.Dir = startDir
	// cmd.Env = startEnv
	stdout, _ := cmd.StdoutPipe()
	stderr, _ := cmd.StderrPipe()
	err := cmd.Start()
	if err != nil {
		log.Fatal(err)
	}
	//time.Sleep(3 * time.Second)

	fmt.Println("--- stderr ---")
	scanner2 := bufio.NewScanner(stderr)
	for scanner2.Scan() {
		fmt.Println(scanner2.Text())
	}

	fmt.Println("--- stdout ---")
	scanner := bufio.NewScanner(stdout)
	for scanner.Scan() {
		fmt.Println(scanner.Text())
	}

	done <- nil
	close(done)
}

func main() {
	// Ctrl+Cを受け取る
	quit := make(chan os.Signal)
	signal.Notify(quit, os.Interrupt)

	done := make(chan error, 1)
	go startProc(done)

	select {
	case <-quit:
		fmt.Println("interrup signal accepted.")
	case err := <-done:
		fmt.Println("error exit.", err)
	}

	// process stop

}
