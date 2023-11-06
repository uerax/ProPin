package pin

import "encoding/json"

func Out() string {
	return pin.Output()
}

func Info() string {
	b, _ := json.Marshal(pin.dirty)
	return string(b)
}

func Add(name, ip string) {
	pin.Add(name, ip)
}