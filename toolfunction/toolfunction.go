package toolfunction

import (
	"github.com/astaxie/beego/logs"
	"net"
	"os"
)

//获取ip
func Getip() string {
	addrs, err := net.InterfaceAddrs()

	if err != nil {
		logs.Error(err)
		os.Exit(1)
	}

	for _, address := range addrs {

		// 检查ip地址判断是否回环地址
		if ipnet, ok := address.(*net.IPNet); ok && !ipnet.IP.IsLoopback() {
			if ipnet.IP.To4() != nil {
				return ipnet.IP.String()
			}
		}
	}
	return "127.0.0.1"
}
