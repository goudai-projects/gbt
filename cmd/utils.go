package cmd

import (
	"bytes"
	"github.com/spf13/cobra"
	"io"
	"log"
	"os"
	"os/exec"
)

func GetFlagValue(cmd *cobra.Command, arg string) string {
	v, err := cmd.Flags().GetString(arg)
	if err != nil {
		panic(err.Error())
	}
	return v
}

func ExecLocalCmd(c *exec.Cmd) {
	var stdBuffer bytes.Buffer
	mw := io.MultiWriter(os.Stdout, &stdBuffer)
	c.Stdout = mw
	c.Stderr = mw
	if err := c.Run(); err != nil {
		log.Panic(err.Error())
	}
	log.Println(stdBuffer.String())
}
