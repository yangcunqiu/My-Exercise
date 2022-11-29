package main

import (
	"bytes"
	"fmt"
	"io"
	"log"
	"os/exec"
)

// 代码执行器
func main() {
	cmd := exec.Command("go", "run", "C:\\project\\go\\My-Exercise\\code\\code-user\\main.go")
	var out, stderr bytes.Buffer
	cmd.Stdout = &out
	cmd.Stderr = &stderr

	pipe, err := cmd.StdinPipe()
	if err != nil {
		log.Fatalln(err)
	}
	_, _ = io.WriteString(pipe, "12 22\n")

	err = cmd.Run()
	if err != nil {
		log.Fatalln(err, stderr.String())
	}
	fmt.Println(out.String())
}
