package main

import (
	"fmt"

	"github.com/vindecodex/Aryzath/proxy/proxy"
)

func main() {
	proxy := proxy.NewProxy()

	fmt.Println(proxy.SetData("user", "Data001"))
	proxy.SetProxyMode("Set")
	fmt.Println(proxy.SetData("user", "Data001"))
	fmt.Println(proxy.SetData("admin", "Data002"))

	fmt.Println(proxy.GetData("user"))
	proxy.SetProxyMode("Get")
	fmt.Println(proxy.GetData("user"))
	fmt.Println(proxy.GetData("admin"))

}
