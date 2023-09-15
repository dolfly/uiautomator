package main

import (
	"log"

	"github.com/dolfly/uiautomator"
)

func main() {
	ui := uiautomator.Connect("ae868d55")
	{
		// info, err := ui.GetSerialNumber()
		// if err != nil {
		// 	log.Fatal(err)
		// }
		// log.Printf("%v", info)
	}

	{
		info, err := ui.GetDeviceInfo()
		if err != nil {
			log.Fatal(err)
		}
		log.Printf("%v\n", info)
	}

	{
		// status := ui.AppInstall("https://ucan.25pp.com/Wandoujia_web_direct_homepage.apk")

		// ok, err := status()
		// if err != nil {
		// 	log.Fatal(err)
		// }
		// if ok {
		// 	log.Print("OK")
		// }

	}

	{
		// info, err := ui.GetWindowSize()
		// if err != nil {
		// 	log.Fatal(err)
		// }
		// log.Printf("%v\n", info)
	}

	{
		// status, err := ui.Ping()
		// if err != nil {
		// 	log.Fatal(err)
		// }
		// log.Println(status)
	}

	{
		// err := ui.NewToast().Show("test", 10)
		// if err != nil {
		// 	log.Fatal(err)
		// }
		// log.Print("OK")
	}

	{
		// err := ui.Click(&uiautomator.Position{X: 100, Y: 100})
		// if err != nil {
		// 	log.Fatal(err)
		// }
		// log.Print("OK")
	}

	// {
	// 	str, err := ui.DumpWindowHierarchy()
	// 	if err != nil {
	// 		log.Fatal(err)
	// 	}
	// 	log.Print(str)
	// }
}
