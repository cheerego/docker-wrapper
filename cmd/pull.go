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

var pullCmd = &cobra.Command{
	Use:   "pull",
	Short: "Pull an image or a repository from a registry",
	Run: func(cmd *cobra.Command, args []string) {
		image := args[0]

		k8s := strings.HasPrefix(image, k8sRegistry)
		gcr := strings.HasPrefix(image, gcrRegistry)
		quay := strings.HasPrefix(image, quayRegistry)

		if k8s {
			trimPrefixImage := strings.TrimPrefix(image, k8sRegistry)
			pull(k8sMirror + trimPrefixImage)
			tag(k8sMirror + trimPrefixImage + " " + image)
		} else if gcr {
			trimPrefixImage := strings.TrimPrefix(image, gcrRegistry)
			pull(gcrMirror + trimPrefixImage)
			tag(gcrMirror + trimPrefixImage + " " + image)
		} else if quay {
			trimPrefixImage := strings.TrimPrefix(image, quayRegistry)
			pull(quayMirror + trimPrefixImage)
			tag(quayMirror + trimPrefixImage + " " + image)
		} else {
			pull(image)
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

func init() {
	rootCmd.AddCommand(pullCmd)
}
