<!DOCTYPE html>
<html>
<head>
	<title>用户中心</title>
</head>
<body>
<dir>欢迎 {{.username}} 进入用户中心</dir>
<ul>
<li>是不是vip：{{.vip}}</li>
<li>vip类型：{{.VIPType}}</li>
<li>Vip过期时间: {{.VIPExprice}}</li>
<li>vip获取时间：{{.GetVIPTime}}</li>
</ul>
<table border="1">
	<tr><th>文件名</th><th>ip</th><th>下载时间</th></tr>
	{{range .downloadrecord}}
	<tr><td>{{.FileName}}</td><td>{{.Ip}}</td><td>{{.DownLoadTime}}</td></tr>
	{{end}}
</table>
<table border="1">
	<tr><th>ip</th><th>国家</th><th>省份</th><th>城市</th><th>登录时间</th></tr>
	{{range .userloginip}}
	<tr><td>{{.Ip}}</td><td>{{.Country}}</td><td>{{.Region}}</td><td>{{.City}}</td><td>{{.LoginTime}}</td></tr>
	{{end}}
</table>
</body>
</html>