package main

import (
	"log"

	"github.com/dolfly/uiautomator"
)

func main() {
	ui := uiautomator.Connect("http://cf6057078b950fcc37bb86b796a65a0e.cc.ipviewer.cn")

	status := ui.AppInstall("https://ucan.25pp.com/Wandoujia_web_direct_homepage.apk")
	s, err := status()
	if err != nil {
		log.Fatal(err)
	}
	log.Println(s)

	// info, err := ui.GetSerialNumber()
	// if err != nil {
	// 	log.Fatal(err)
	// }
	// log.Printf("%v", info)

	// info, err := ui.GetDeviceInfo()
	// if err != nil {
	// 	log.Fatal(err)
	// }
	// log.Printf("%v\n", info)

	// info, err := ui.GetWindowSize()
	// if err != nil {
	// 	log.Fatal(err)
	// }
	// log.Printf("%v\n", info)

	// status, err := ui.Ping()
	// if err != nil {
	// 	log.Fatal(err)
	// }
	// log.Println(status)

	// err := ui.NewToast().Show("test", 10)

	// if err != nil {
	// 	log.Fatal(err)
	// }
	// log.Print("OK")
}
