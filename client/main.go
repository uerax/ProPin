package main

import (
	"bytes"
	"encoding/json"
	"flag"
	"fmt"
	"net/http"
	"time"
)

var (
	serverApi  string
	ipme       = "https://api.ipify.org/?format=text"
	serverName string
)

func main() {
	flag.StringVar(&serverApi, "s", "", "探针ProPin中心域名/IP")
	flag.StringVar(&serverName, "n", "", "服务器别名")
	flag.Parse()
	if serverApi == "" {
		panic("没有提供探针ProPin中心访问链接")
	}
	t := time.NewTicker(time.Minute)

	do := func() {
		status := &Status{Name: serverName}
		status.Coll()
		status.Report()
	}
	do()
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
		go do()
	}

}

type Status struct {
	Name string
}

func (t *Status) Coll() {

}

func (t *Status) Report() {
	data, err := json.Marshal(t)
	if err != nil {
		fmt.Println("JSON编码失败:", err)
		return
	}

	resp, err := http.Post(serverApi, "application/json", bytes.NewBuffer(data))
	if err != nil {
		fmt.Println("POST请求失败:", err)
		return
	}
	defer resp.Body.Close()
	fmt.Println("POST请求成功: ", t.Name)
}
