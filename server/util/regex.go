package util

import (
	"net"
	"regexp"
)

func GetHttpsHostRegex(webUrl string) string {
	//host, err := url.Parse(webUrl)
	//println(webUrl)
	//if err != nil {
	//	log.Println(err)
	//}
	//ip := host.Hostname()
	//println(host.Hostname())
	//port := host.Port()
	//if port == "" {
	//	return ip + ":433"
	//} else {
	//	return ip + ":" + port
	//}

	portReg := regexp.MustCompile(`(:)\d{2,5}`)
	port := portReg.FindString(webUrl)
	reg := regexp.MustCompile(`(?i)(http://|https://|\/|:([0-9]+))`)
	host := reg.ReplaceAllString(webUrl, "")

	switch port {
	case "443":
		host += ":443"
	case "":
		host += ":443"
	default:
		host += port
	}

	return host
}

func GetHttpHostRegex(url string) string {
	reg := regexp.MustCompile(`(?i)(http://|https://|\/|:80)`)
	url = reg.ReplaceAllString(url, "")
	url += ":80"
	return url
}

func IsValidHost(host string) bool {
	reg := regexp.MustCompile(`(?i)(http://|https://|\/|:([0-9]+))`)
	host = reg.ReplaceAllString(host, "")
	// Parse as IP address
	if ip := net.ParseIP(host); ip != nil {
		return true
	}

	// Parse as hostname
	if _, err := net.LookupHost(host); err == nil {

		return true
	}

	// Parse as URL
	//if _, err := url.Parse(host); err == nil {
	//	return true
	//}
	return false
}
