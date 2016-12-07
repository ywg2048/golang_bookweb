<!DOCTYPE html>
<html>
<head>
	<title>注册成功！</title>
</head>
<body onload="countDown(5,'/')">
<h2>注册成功！</h2>
<div><span id="secs" style="color: green">5</span>秒之后跳转到首页</div>
<script type="text/javascript">
	function countDown(secs,surl){     
		 //alert(surl); 
		  var jumpTo = document.getElementById('secs');
 			jumpTo.innerHTML=secs;      
		if(--secs>0){     
		    setTimeout("countDown("+secs+",'"+surl+"')",1000);     
		}     
		else{       
		    location.href=surl;     
	    }     
	 }     
</script>
</body>
</html>