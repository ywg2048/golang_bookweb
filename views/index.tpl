<!DOCTYPE html>
<html>
<head>
	<title>网站首页</title>
	<meta charset="utf-8">
	<script type="text/javascript" src = "/static/js/jquery.min.js"></script>
</head>
<body>
<div>网站首页</div>

	{{if .islogin}}
		<div>欢迎 {{.username}}</div>
		<div onclick ="usercenter()">个人中心</div>
		<div onclick="quit()">注销</div>
	{{else}}
		<div>
		<a href="/User/register">注册</a>|<a href="/User/login">登录</a>
		</div>
	{{end}}


<ul>

{{range .files}}
    <li>{{.Id}} -- {{.FileName}}<a href="{{.Path}}">下载</a></li>
{{end}}
<div>总页数：{{.totalpages}}  当前页：{{.currentpage}}</div>
<div>
	<a href="/?Index=0">首页</a>
	<a href="/?Index={{.prepage}}">上一页</a>
	<span id="pagelist"></span>
	<a href="/?Index={{.nextpage}}">下一页</a>
	<a href="/?Index={{.lastpage}}">尾页</a>
	<input type="text" name="goto" id = "Index"/>
	<input type="button" value="GO"  onclick="Goto()" />
</div>
</ul>
</body>
<script type="text/javascript">
	$(document).ready(function(){
	  	$("#pagelist").html("");
	  	var str = "";
	  	var pagelistsize = {{.pagelistsize}};
	  	if({{.currentpage}}<=pagelistsize && {{.totalpages}} <=pagelistsize){
	  		for (var i = 0; i < {{.totalpages}}; i++) {
		  		temp = "<a href='/?Index="+i+"'>"+(i+1)+"|</a>";
		  		if(i == {{.currentpage}}-1) {
		  			//当前页面红色标出
		  			temp = "<a href='/?Index="+i+"' style='color:red'>"+(i+1)+"|</a>";
		  		}
		  		str +=temp;
		  	}
		  	$("#pagelist").html(str);
	  	}else if({{.currentpage}}<=pagelistsize && {{.totalpages}} >pagelistsize){
	  		for (var i = 0; i < pagelistsize; i++) {
		  		temp = "<a href='/?Index="+i+"'>"+(i+1)+"|</a>";
		  		if(i == {{.currentpage}}-1) {
		  			//当前页面红色标出
		  			temp = "<a href='/?Index="+i+"' style='color:red'>"+(i+1)+"|</a>";
		  		}
		  		str +=temp;
		  	}
		  	$("#pagelist").html(str);
	  	}

	  	var currentlist = Math.ceil({{.currentpage}}/pagelistsize);
	  	var lastlist = ({{.totalpages}})%pagelistsize;
	  	var maxlist = Math.ceil(({{.totalpages}})/pagelistsize);
	  	if(currentlist>=2 && lastlist ==0){
	  		for (var i = (currentlist-1)*pagelistsize ;i<currentlist*pagelistsize;i++) {
	  			temp = "<a href='/?Index="+i+"'>"+(i+1)+"|</a>";
		  		if(i == {{.currentpage}}-1) {
		  			//当前页面红色标出
		  			temp = "<a href='/?Index="+i+"' style='color:red'>"+(i+1)+"|</a>";
		  		}
		  		str +=temp;
	  		}
	  		$("#pagelist").html(str);
	  	}else if(currentlist>=2 && lastlist!=0){
	  		for (var i = (currentlist-1)*pagelistsize ;i<currentlist*pagelistsize;i++) {
	  			if(i>(maxlist-1)*pagelistsize+lastlist-1){
	  				continue
	  			}
	  			temp = "<a href='/?Index="+i+"'>"+(i+1)+"|</a>";
		  		if(i == {{.currentpage}}-1) {
		  			//当前页面红色标出
		  			temp = "<a href='/?Index="+i+"' style='color:red'>"+(i+1)+"|</a>";
		  		}
		  		str +=temp;
	  		}
	  		$("#pagelist").html(str);
	  	}
	  	
	});
	function Goto(){
		Index = document.getElementById("Index").value;
		location.href = "/?Index="+(Index-1); 
	}
	function quit(){
		//注销
		$.ajax({
			type: "POST",
			url: "/",
			data: {quit:1},
			dataType: "json",
			success: function(data){
				document.location.reload();
			}
		});
	}

	function usercenter(){
		location.href = "/User/center"
	}
</script>
</html>