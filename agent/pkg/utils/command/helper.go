package command

import (
	"bufio"
	"io"
	"io/ioutil"
	"log"
	"os/exec"
	"strings"
)

//执行命令行
func ExeCmdByShell(s string) string {
	cmd := exec.Command("/bin/bash", "-c", s)

	stdout, err := cmd.StdoutPipe()
	if err != nil {
		return ""
	}
	if err := cmd.Start(); err != nil {
		return ""
	}

	bytes, _ := ioutil.ReadAll(stdout)

	if err := cmd.Wait(); err != nil {
		return ""
	}

	return string(bytes)
}

//执行对应命令
func execCommandStdoutPipe(cmd string, arg ...string) []string {
	command := exec.Command(cmd, arg...)
	stdout, err := command.StdoutPipe()
	if err != nil {
		log.Fatal(err)
	}
	command.Start()
	lineS := make([]string, 0)
	reader := bufio.NewReader(stdout)

	for {
		line, err2 := reader.ReadString('\n')
		if err2 != nil || io.EOF == err2 {
			break
		}
		lineS = append(lineS, strings.Trim(line, "\n"))
	}

	return lineS
}

func Cat(file string) []string {
	r := execCommandStdoutPipe("cat", file)

	return r
}
