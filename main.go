package main

import (
	"fmt"
	"os"
	"os/exec"
)

// https://zetcode.com/golang/exec-command/
// search golang start detached process
// https://ieftimov.com/post/four-steps-daemonize-your-golang-programs/
// https://newbedev.com/start-a-detached-process-on-windows-using-golang
// https://stackoverflow.com/questions/23031752/start-a-process-in-go-and-detach-from-it

func main() {
	s, err := RunAndReturn("/Users/q/go/src/github.com/qiuzhanghua/tdp_go/tdp", "version")
	fmt.Println(s, err)
	// RunAndForget("abs_exec", "version")
	//RunAndDetach("deno", []string{"run", "--allow-net", "/Users/q/exec/hello.ts"})
}

func RunAndReturn(command string, args ...string) (string, error) {
	out, err := exec.Command(command, args...).Output()
	if err != nil {
		return "", err
	}
	return string(out), nil

}

func RunAndForget(command string, args ...string) error {
	cmd := exec.Command(command, args...)
	cmd.Stdin = os.Stdin
	cmd.Stdout = os.Stdout
	cmd.Stderr = os.Stderr
	return cmd.Run()
}

// RunAndDetach
// tested under darwin
func RunAndDetach(command string, args ...string) error {
	args = append(args, "--detached")
	cmd := exec.Command(command, args...)
	cmd.Stdin = os.Stdin
	cmd.Stdout = os.Stdout
	cmd.Stderr = os.Stderr
	err := cmd.Start()
	if err != nil {
		return err
	}
	return cmd.Process.Release()
}
