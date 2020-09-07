/*
Copyright © 2020 NAME HERE <EMAIL ADDRESS>

Licensed under the Apache License, Version 2.0 (the "License");
you may not use this file except in compliance with the License.
You may obtain a copy of the License at

    http://www.apache.org/licenses/LICENSE-2.0

Unless required by applicable law or agreed to in writing, software
distributed under the License is distributed on an "AS IS" BASIS,
WITHOUT WARRANTIES OR CONDITIONS OF ANY KIND, either express or implied.
See the License for the specific language governing permissions and
limitations under the License.
*/
package cmd

import (
	"fmt"
	"github.com/spf13/cobra"
	"github.com/spf13/viper"
	"io/ioutil"
	"log"
	"os/exec"
)

// goCmd represents the go command
var goCmd = &cobra.Command{
	Use:   "go",
	Short: "Build a go project to docker image",
	Long:  ``,
	Run: func(cmd *cobra.Command, args []string) {
		sprintf := fmt.Sprintf(dockerfile, "1.15.1")
		if err := ioutil.WriteFile("./Dockerfile", []byte(sprintf), 0655); err != nil {
			log.Panic(err.Error())
		}
		i := GetFlagValue(cmd, "image", viper.GetString("IMAGE"))
		cmds := make([]string, 0)
		cmds = append(cmds, "build", "-t", i, ".")
		ExecLocalCmd(exec.Command("docker", cmds...))
	},
}

func init() {
	buildCmd.AddCommand(goCmd)
	goCmd.Flags().StringP("image", "i", viper.GetString("IMAGE"), "docker out image ")
	goCmd.MarkFlagRequired("image")
}

var dockerfile = `
FROM golang:%v as builder
WORKDIR /workspace
COPY . .
ENV GOPROXY=https://goproxy.io
RUN ls . && go mod download &&  CGO_ENABLED=0 GOOS=linux GOARCH=amd64 GO111MODULE=on go build -a -o app main.go


FROM registry.cn-shanghai.aliyuncs.com/qingmuio/distroless_static:nonroot
WORKDIR /
COPY --from=builder /workspace/manager .
USER nonroot:nonroot
ENTRYPOINT ["/app"]
`
