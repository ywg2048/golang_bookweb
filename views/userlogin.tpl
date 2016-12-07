<!DOCTYPE html>
<html>
<head>
	<title>用户登录</title>
	<meta charset="utf-8">
</head>
<body>
<h2 style="text-align: center;">用户登录</h2>
<table align="center">
	<form action="/User/login" method="post">
		<tr><td>用户名：</td><td><input type="text" name="username"></td></tr>
		<tr><td>密码：</td><td><input type="password" name="passwd"></td></tr>
		<tr><td colspan="2" align="center"><input type="submit" name="提交" value="登录"></td></tr>
		<tr align="center"><td colspan="2"><font color="red">{{.err}}</font></td></tr>
	</form>
</table>
</body>
</html>