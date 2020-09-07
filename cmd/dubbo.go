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
	"log"
	"os/exec"
	"strings"

	"github.com/spf13/cobra"
)

// dubboCmd represents the dubbo command
var dubboCmd = &cobra.Command{
	Use:   "dubbo",
	Short: "Build multi-module dubbo project to docker image",
	Long:  ``,
	Run: func(cmd *cobra.Command, args []string) {
		cmds := make([]string, 0)
		cmds = append(cmds, "clean", "package", "com.google.cloud.tools:jib-maven-plugin:build")
		a := GetFlagValue(cmd, "active-maven-profile", viper.GetString("ACTIVE_MAVEN_PROFILE"))
		u := GetFlagValue(cmd, "username", viper.GetString("DOCKER_USERNAME"))
		p := GetFlagValue(cmd, "password", viper.GetString("DOCKER_PASSOWRD"))
		f := GetFlagValue(cmd, "from-image", viper.GetString("FROM_IMAGE"))
		i := GetFlagValue(cmd, "image", viper.GetString("IMAGE"))
		m := GetFlagValue(cmd, "module-name", viper.GetString("MODULE_NAME"))
		if m == "" {
			log.Panic("dubbo project module name cannot be null")
			return
		}
		if a != "" {
			cmds = append(cmds, fmt.Sprintf("-P%v", a))
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
		cmds = append(cmds, "-pl")
		cmds = append(cmds, m)
		cmds = append(cmds, "-am")
		log.Println(strings.Join(cmds, " "))
		ExecLocalCmd(exec.Command("mvn", cmds...))
	},
}

func init() {
	buildCmd.AddCommand(dubboCmd)
	dubboCmd.Flags().StringP("active-maven-profile", "a", viper.GetString("PROFILE"), "active maven profile")
	dubboCmd.Flags().StringP("username", "u", viper.GetString("DOCKER_USERNAME"), "docker username ")
	dubboCmd.Flags().StringP("password", "p", viper.GetString("DOCKER_PASSWORD"), "docker passdowrd ")
	dubboCmd.Flags().StringP("from-image", "f", defaultJreImage, "docker base image  ")
	dubboCmd.Flags().StringP("image", "i", viper.GetString("IMAGE"), "docker out image ")
	dubboCmd.Flags().StringP("module-name", "m", defaultAppRoot, "dubbo module name")
	dubboCmd.MarkFlagRequired("image")
}
