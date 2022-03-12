package atomgo

import (
	"fmt"
	"regexp"
	"strconv"
	"strings"
)

const (
	proxyTypeSocks4 int = 1
	proxyTypeSocks5 int = 2
	proxyTypeHttp   int = 3
)

type Proxy struct {
	host      string
	port      int
	userName  string
	passWord  string
	proxyType int
	proxyRdns bool
}

func StrToProxyHttp(proxyStr string) Proxy {
	return StrToProxy(proxyStr, proxyTypeHttp, true)
}

func StrToProxySocks4(proxyStr string) Proxy {
	return StrToProxy(proxyStr, proxyTypeSocks4, true)
}

func StrToProxySocks5(proxyStr string) Proxy {
	return StrToProxy(proxyStr, proxyTypeSocks5, true)
}

// proxy string to Proxy
func StrToProxy(proxyStr string, proxyType int, proxyRdns bool) Proxy {
	var p Proxy
	if strings.Contains(proxyStr, "@") {
		p.userName, p.passWord, p.host, p.port = parseProxyStr4(proxyStr)
	} else {
		p.userName, p.passWord, p.host, p.port = parseProxyStr2(proxyStr)
	}
	p.proxyType = proxyType
	p.proxyRdns = proxyRdns
	return p
}

// parse proxy string like host:port
func parseProxyStr2(proxyStr string) (string, string, string, int) {
	addr := `[\d]{1,3}\.[\d]{1,3}\.[\d]{1,3}\.[\d]{1,3}`
	pattern := `^(` + addr + `):([\d]{2,5})$`
	reg := regexp.MustCompile(pattern)
	match := reg.FindStringSubmatch(proxyStr)
	if match == nil {
		panic(`proxy string format error: ` + proxyStr)
	}
	port, err := strconv.Atoi(match[2])
	if err != nil || match[1] == "" {
		panic(`proxy string format error: ` + proxyStr)
	}
	return "", "", match[1], port

}

// parse proxy string like user:pass@host:port
func parseProxyStr4(proxyStr string) (string, string, string, int) {
	addr := `[\d]{1,3}\.[\d]{1,3}\.[\d]{1,3}\.[\d]{1,3}`
	pattern := `^([\w]+):([\w]+)@(` + addr + `):([\d]{2,5})$`
	reg := regexp.MustCompile(pattern)
	match := reg.FindStringSubmatch(proxyStr)
	if match == nil {
		panic(`proxy string format error: ` + proxyStr)
	}
	port, err := strconv.Atoi(match[4])
	if err != nil || match[1] == "" || match[2] == "" || match[3] == "" {
		panic(`proxy string format error: ` + proxyStr)
	}
	return match[1], match[2], match[3], port
}

// Proxy to proxy string
func ProxyToStr(p Proxy) string {
	if p.userName != "" && p.passWord != "" {
		return fmt.Sprintf(`%v:%v@%v:%v`, p.userName, p.passWord, p.host, p.port)
	}
	return fmt.Sprintf(`%v:%v`, p.host, p.port)
}
