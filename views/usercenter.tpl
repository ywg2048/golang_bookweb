<!DOCTYPE html>
<html>
<head>
	<title>用户中心</title>
</head>
<body>
<dir>欢迎 {{.username}} 进入用户中心</dir>
<table border="1">
	<tr><th>文件名</th><th>ip</th><th>下载时间</th></tr>
	{{range .downloadrecord}}
	<tr><td>{{.FileName}}</td><td>{{.Ip}}</td><td>{{.DownLoadTime}}</td></tr>
	{{end}}
</table>
</body>
</html>