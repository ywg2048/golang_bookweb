package toolfunction

import (
	"encoding/json"
	"fmt"
	"github.com/astaxie/beego"
	"github.com/astaxie/beego/logs"
	"io/ioutil"
	"net"
	"net/http"
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

//根据ip获取国家城市名
func GetCity(ip string) (string, string, string) {
	baseurl := beego.AppConfig.String("GetAdressAPI")
	url := baseurl + ip
	fmt.Println("url= ", url)
	resp, err := http.Get(url)
	if err != nil {
		// handle error
		logs.Error("淘宝ip接口返回错误", err)
	}
	defer resp.Body.Close()
	body, err := ioutil.ReadAll(resp.Body)
	if err != nil {
		// handle error
		logs.Error("解析包出错", err)
	}
	var data map[string]interface{}
	json.Unmarshal(body, &data)

	fmt.Println(data["code"])
	fmt.Println(data["data"])

	//对包体断言
	databody, _ := (data["data"]).(map[string]interface{})
	country, _ := databody["country"].(string)
	region, _ := databody["region"].(string)
	city, _ := databody["city"].(string)
	logs.Debug("country = ", country)
	logs.Debug("region = ", region)
	logs.Debug("city = ", city)
	if region == "" && city == "" {
		return "本地", "本地", "本地"
	}
	return country, region, city

}

//从邮箱里查看是否有充值记录
func getRechargeMessage() {

}
