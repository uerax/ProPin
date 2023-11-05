package main

import (
	"flag"
	"time"
)

var (
	server string
	ipme   = "https://api.ipify.org/?format=text"
)

func main() {
	flag.StringVar(&server, "s", "", "探针ProPin中心域名/IP")
	flag.Parse()
	if server == "" {
		panic("没有提供探针ProPin中心访问链接")
	}
 	t := time.NewTicker(time.Minute)
	for range t.C {
		// 直接在服务端拿请求ip即可
		// res, err := http.Get(ipme)
		// if err == nil && res.StatusCode == http.StatusOK {
		// 	defer res.Body.Close()
		// 	ip, err := io.ReadAll(res.Body)
		// 	if err == nil {
		// 		fmt.Println(string(ip))
		// 	}
		// }
		
	}

}
