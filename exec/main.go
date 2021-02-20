package main

import (
	"fmt"
	"os/exec"
)

func main() {
	var output, _ = exec.Command("ls").Output()
	fmt.Printf(" -> ls\n%s\n", string(output))

	var output2, _ = exec.Command("git", "--version").Output()
	fmt.Printf("-> git --version\n%s\n", string(output2))
}
