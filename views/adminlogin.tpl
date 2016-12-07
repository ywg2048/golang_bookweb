<html>
<head>
	<title>管理员登录</title>
	<meta charset="utf-8">
</head>
<body>
<h2 style="text-align: center;">后台登录</h2>
<table align="center">
	<form action="/Admin/login" method="post">
		<tr><td>用户名：</td><td><input type="text" name="username"></td></tr>
		<tr><td>密码：</td><td><input type="password" name="passwd"></td></tr>
		<tr><td colspan="2" align="center"><input type="submit" name="提交" value="登录"></td></tr>
		<tr><td><font color="red">{{.err}}</font></td></tr>
	</form>
</table>
</body>
</html>