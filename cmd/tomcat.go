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
	"github.com/spf13/viper"
	"os/exec"

	"github.com/spf13/cobra"
)

var defaultTomcatBaseImage = "tomcat:8.5-jre8-alpine"
var defaultAppRoot = "/usr/local/tomcat/webapps/ROOT"

// tomcatCmd represents the tomcat command
var tomcatCmd = &cobra.Command{
	Use:   "tomcat",
	Short: "Build A tomcat project to docker image",
	Long:  ``,
	Run: func(cmd *cobra.Command, args []string) {
		cmds := make([]string, 0)
		cmds = append(cmds, "clean", "package", "com.google.cloud.tools:jib-maven-plugin:build")
		a := GetFlagValue(cmd, "active-maven-profile")
		u := GetFlagValue(cmd, "username")
		p := GetFlagValue(cmd, "password")
		f := GetFlagValue(cmd, "from-image")
		i := GetFlagValue(cmd, "image")
		r := GetFlagValue(cmd, "app-root")
		if a != "" {
			cmds = append(cmds, fmt.Sprintf("-P %v", a))
		}
		if u != "" && p != "" {
			cmds = append(cmds, fmt.Sprintf("-Djib.to.auth.username=%v", u))
			cmds = append(cmds, fmt.Sprintf("-Djib.to.auth.password=%v", p))
		}
		if f == "" {
			f = defaultJreImage
		}
		cmds = append(cmds, fmt.Sprintf("-Djib.from.image=%v", f))
		cmds = append(cmds, fmt.Sprintf("-Dimage=%v", i))
		if r == "" {
			r = defaultAppRoot
		}
		cmds = append(cmds, fmt.Sprintf("-Djib.container.appRoot=%v", r))
		//log.Println("mvn " + strings.Join(cmds, " "))
		ExecLocalCmd(exec.Command("mvn", cmds...))
	},
}

func init() {
	buildCmd.AddCommand(tomcatCmd)
	tomcatCmd.Flags().StringP("active-maven-profile", "a", viper.GetString("PROFILE"), "active maven profile")
	tomcatCmd.Flags().StringP("username", "u", viper.GetString("DOCKER_USERNAME"), "docker username ")
	tomcatCmd.Flags().StringP("password", "p", viper.GetString("DOCKER_PASSWORD"), "docker passdowrd ")
	tomcatCmd.Flags().StringP("from-image", "f", defaultTomcatBaseImage, "docker base image  ")
	tomcatCmd.Flags().StringP("image", "i", viper.GetString("IMAGE"), "docker out image ")
	tomcatCmd.Flags().StringP("app-root", "r", defaultAppRoot, "webapps path eg. ROOT")
	tomcatCmd.MarkFlagRequired("image")
}
