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
package main

import (
	"bytes"
	"fmt"
	"github.com/goudai-projects/gd-build-tools/cmd"
	"os/exec"
)

func main() {

	//Test(exec.Command("mvn", "com.google.cloud.tools:jib-maven-plugin:build"))

	//cmds := exec.Command("mvn","-v")

	cmd.Execute()
}

func Test(cmds *exec.Cmd) {
	var out bytes.Buffer
	var stderr bytes.Buffer
	cmds.Stdout = &out
	cmds.Stderr = &stderr
	err := cmds.Run()

	if err != nil {
		fmt.Println(fmt.Sprint(err) + ": " + stderr.String())
		return
	}
	fmt.Println("Result: " + out.String())
}
