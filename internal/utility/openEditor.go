package utility

import (
	"os"
	"os/exec"
	"strings"
)

func OpenEditor(editor string, fileName string) error {
	cmdArgs := []string{}
	var cmdName string
	for i, part := range strings.Fields(editor) {
		if i == 0 {
			cmdName = part
		} else {
			cmdArgs = append(cmdArgs, part)
		}
	}
	cmdArgs = append(cmdArgs, fileName)

	cmd := exec.Command(cmdName, cmdArgs...)
	cmd.Stdin = os.Stdin
	cmd.Stdout = os.Stdout
	cmd.Stderr = os.Stderr

	return cmd.Run()
}
