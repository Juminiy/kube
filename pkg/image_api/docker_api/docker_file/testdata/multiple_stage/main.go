package main

import (
	"flag"
	"fmt"
	"github.com/Juminiy/kube/pkg/util"
	"github.com/go-resty/resty/v2"
	"net"
	"time"
)

func main() {
	switch app {
	case "http":
		httpConn()
	case "net":
		netConn()
	default:
		fmt.Println("APP: net or http")
	}
}

var (
	app  string
	addr string
	rto  int
)

func init() {
	flag.StringVar(&app, "app", "http", "APP: net or http")
	flag.StringVar(&addr, "addr", "github.com", "NetConn Address")
	flag.IntVar(&rto, "rto", 12, "NetConn ReadTimeout Second")
	flag.Parse()
}

func netConn() {
	c, err := net.Dial("tcp", addr+":http")
	util.Must(err)
	defer util.SilentCloseIO("tcp conn", c)
	err = c.SetDeadline(time.Now().Add(util.TimeSecond(rto)))
	util.Must(err)
	bs := make([]byte, 0, 4*util.Ki)
	_, err = c.Read(bs)
	if err == nil {
		if len(bs) == 0 {
			fmt.Println("success but read nothing")
		} else {
			fmt.Println(util.Bytes2StringNoCopy(bs))
		}
	} else {
		switch errv := err.(type) {
		case *net.OpError:
			if errv.Timeout() {
				fmt.Printf("NetConn to: %s Timeout in %ds", addr, rto)
			}
		}
		fmt.Println(err)
	}
}

func httpConn() {
	resp, err := resty.New().
		SetScheme("https").
		SetBaseURL(util.TrimProto(addr)).
		SetAllowGetMethodPayload(true).
		SetTimeout(util.TimeSecond(rto)).
		R().Get("/")
	if err != nil {
		fmt.Printf("http request to: %s, in readtimeout: %d, failed\n", addr, rto)
		fmt.Printf("%s\n", err.Error())
	} else {
		fmt.Printf("http request to: %s, in readtimeout: %d, success\n", addr, rto)
		fmt.Printf("status_code: %d, content: %s\n", resp.StatusCode(), func() string {
			_respStr := resp.String()
			if len(_respStr) > 512 {
				return _respStr[:512]
			}
			return _respStr
		}())
	}
}
