package pin

import (
	"sync"
	"time"
)

var (
	pin  *Pin
	once sync.Once
)

type Pin struct {
	mx       sync.Mutex
	pins     map[string]string
	dirty    map[string]string
	set 	 map[string]struct{}
	interval uint // 秒
}

func NewPin(itv uint) {
	pin = &Pin{
		mx:       sync.Mutex{},
		pins:     make(map[string]string),
		dirty:    make(map[string]string),
		set:      make(map[string]struct{}),
		interval: itv,
	}
	go pin.cron()
}

func (p *Pin) cron() {
	t := time.NewTicker(time.Duration(p.interval) * time.Second)
	for range t.C {
		p.mx.Lock()
		p.dirty = p.pins
		p.pins = make(map[string]string)
		p.mx.Unlock()
	}
}

func (t *Pin) Add(name, ip string) {
	t.pins[name] = ip
	if _, ok := t.set[name]; !ok {
		t.set[name] = struct{}{}
	}
}

func (t *Pin) Output() string {
	t.mx.Lock()
	defer t.mx.Unlock()
	out := "服务器列表:"
	for name := range t.set {
		ip := "未上报"
		if v, ok := t.dirty[name]; ok {
			ip = v
		}
		out += "\n"
		out += name + " : " + ip
	}
	return out
}