package main

import (
	"log"
	"os"
	"path/filepath"
	"strings"

	"github.com/dolfly/uiautomator/internal/gadb"
)

func main() {
	adbClient, err := gadb.NewClient()
	checkErr(err, "fail to connect adb server")

	devices, err := adbClient.DeviceList()
	checkErr(err)

	for _, d := range devices {
		log.Printf("Device: %s is connected\n", d.Serial())
	}

	if len(devices) == 0 {
		log.Fatalln("list of devices is empty")
	}

	dev := devices[0]
	dir, err := os.Getwd()
	checkErr(err, "os getwd")

	testfile, err := os.Open(filepath.Join(dir, "test.txt"))
	checkErr(err)

	log.Println("starting to push apk")

	remotePath := "/data/local/tmp/test.txt"
	err = dev.PushFile(testfile, remotePath)
	checkErr(err, "adb push")

	log.Println("push completed")

	log.Println("starting to install apk")

	shellOutput, err := dev.RunShellCommand("cat", remotePath)
	checkErr(err, "cat file")
	if !strings.Contains(shellOutput, "Success") {
		log.Fatalln("fail to install: ", shellOutput)
	}
	log.Println("execute completed")

	// shellOutput, err := dev.RunShellCommand("pm install", remotePath)
	// checkErr(err, "pm install")
	// if !strings.Contains(shellOutput, "Success") {
	// 	log.Fatalln("fail to install: ", shellOutput)
	// }

	// log.Println("install completed")

}

func checkErr(err error, msg ...string) {
	if err == nil {
		return
	}

	var output string
	if len(msg) != 0 {
		output = msg[0] + " "
	}
	output += err.Error()
	log.Fatalln(output)
}
