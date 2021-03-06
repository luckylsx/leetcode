package test

import (
	"os"
	"os/exec"

	"github.com/openset/leetcode/internal/base"
)

var CmdTest = &base.Command{
	Run:       runTest,
	UsageLine: "test",
	Short:     "run go test",
	Long:      "automates testing the packages.",
}

func runTest(cmd *base.Command, args []string) {
	if len(args) != 0 {
		cmd.Usage()
	}
	c := exec.Command("go", "test", "./...")
	c.Stdout = os.Stdout
	c.Stderr = os.Stderr
	err := c.Run()
	base.CheckErr(err)
}
