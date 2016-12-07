<!DOCTYPE html>
<html>
<head>
	<meta charset="utf-8">
	<title>注册页面</title>
	<script type="text/javascript" src = "/static/js/jquery.min.js"></script>
<script type="text/javascript">
	//判断用户是否已经注册
	function checkuser(){
		$.ajax({
             type: "POST",
             url: "/User/register",
             data: {username:$("#username").val()},
             dataType: "json",
             success: function(data){
                if(data.code == 1){
                	$("#noticeText").html("<font color ='red'>用户已经注册</font>")
                }else if(data.code == 0){
                	$("#noticeText").html("<font color ='green'>用户可注册</font>")
                }

             }
     	});
	}
	//发送邮件
	function sendCode(){
		$('#button').attr('disabled',"true");
		$.ajax({
             type: "POST",
             url: "/User/register",
             data: {getcode:"1",email:$("#email").val()},
             dataType: "json",
             success: function(data){
             	if(data.code == 1){
    	         	alert("邮件发送成功");
             	}
             }
     	});
	}
	function checkSubmit(){
		username = $("#username").val();
		
		email = $("#email").val();
		checkcode = $("#checkcode").val();
		passwd = $("#passwd").val();

		if(username!="" && email != "" && passwd != ""&& checkcode !=""){
			document.getElementById("formid").submit();
			return true;
		}else{
			alert("不能为空");
			return false;
		}
	}
</script>
</head>
<body>
<h2 style="text-align: center;">用户注册</h2>
<table>
	<form id="formid" action="/User/register" method="POST" >
		<tr><td>用户名：</td><td><input type="text" id = "username" name="username" placeholder="用户名" onkeyup="checkuser()"><span id="noticeText"></span></td></tr>
		<tr><td>邮箱：</td><td><input type="email"  id = "email" name="email"><input type="button" id="button" value="发送验证码" onclick="sendCode()"></td></tr>
		<tr><td>验证码：</td><td><input type="text" id = "checkcode" name="checkcode" placeholder=""></td></tr>
		<tr><td>密码</td><td><input type="password" id = "passwd" name="passwd"></td></tr>
		<tr align="center"><td colspan="2"><font color="red">{{.errtext}}</font></td></tr>
		<tr><td colspan="2"><input type="button" value="注册" onclick ="return checkSubmit()"></td></tr>
	</form>
</table>

</body>
</html>