package cmd

import (
	"github.com/spf13/cobra"
	"log"
	"os"
	"os/exec"
	"strings"
)

var k8sRegistry string = "k8s.gcr.io"
var gcrRegistry string = "gcr.io"
var quayRegistry string = "quay.io"

var k8sMirror string = "gcr.azk8s.cn/google_containers"
var gcrMirror string = "gcr.azk8s.cn"
var quayMirror string = "quay-mirror.qiniu.com"

var dockerHubAz string = "dockerhub.azk8s.cn"

var pullCmd = &cobra.Command{
	Use:   "pull",
	Short: "Pull an image or a repository from a registry",
	Run: func(cmd *cobra.Command, args []string) {
		var targetImageName string
		var sourceImageName string = args[0]

		k8s := strings.HasPrefix(sourceImageName, k8sRegistry)
		gcr := strings.HasPrefix(sourceImageName, gcrRegistry)
		quay := strings.HasPrefix(sourceImageName, quayRegistry)

		if k8s {
			trimPrefixImage := strings.TrimPrefix(sourceImageName, k8sRegistry)
			targetImageName = k8sMirror + trimPrefixImage
		} else if gcr {
			trimPrefixImage := strings.TrimPrefix(sourceImageName, gcrRegistry)
			targetImageName = gcrMirror + trimPrefixImage
		} else if quay {
			trimPrefixImage := strings.TrimPrefix(sourceImageName, quayRegistry)
			targetImageName = quayMirror + trimPrefixImage
		} else {
			targetImageName = sourceImageName
		}

		pull(targetImageName)

		if k8s || gcr || quay {
			tag(targetImageName + " " + sourceImageName)
			rmi(targetImageName)
		}
	},
}

func pull(cmd string) {
	log.Println(cmd)
	command := exec.Command("bash", "-c", "docker pull "+cmd)
	command.Stdout = os.Stdout
	command.Stderr = os.Stderr
	e := command.Run()
	if e != nil {
		panic(e)
	}
}

func tag(cmd string) {
	log.Println(cmd)
	command := exec.Command("bash", "-c", "docker tag "+cmd)
	command.Stdout = os.Stdout
	command.Stderr = os.Stderr
	e := command.Run()
	if e != nil {
		panic(e)
	}
}

func rmi(cmd string) {
	command := exec.Command("bash", "-c", "docker rmi "+cmd)
	command.Stdout = os.Stdout
	command.Stderr = os.Stderr
	e := command.Run()
	if e != nil {
		panic(e)
	}
}

func init() {
	rootCmd.AddCommand(pullCmd)
}
