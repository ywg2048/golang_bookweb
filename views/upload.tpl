<!DOCTYPE html>
<html>
<head>
	<title>资料上传</title>
	<meta charset="utf-8">
	<script type="text/javascript" src = "/static/js/jquery.min.js"></script>
	<script src="/static/js/jquery.uploadify.min.js" type="text/javascript"></script>
	<link rel="stylesheet" href="/static/css/uploadify.css" type="text/css">
	<script type="text/javascript">     
		$(function(){             
		    $("#file_upload").uploadify({       //绑定元素      
		        'fileObjName':'image',//html input标签的name属性的值吧。                 
		        'debug':false,                  
		        'auto':true,    //自动上传              
		    'buttonText':'Choose A File',                   
		    'removeCompleted':false, //上传完成以后是否保存进度条                    
		    'cancelImg':'/static/img/uploadify-cancel.png',                   
		    'swf':'/static/js/uploadify.swf', //必须设置  swf文件路径             
		    'uploader':"/upload",         //必须设置，上传文件触发的url       
		    'fileTypeDesc':'FileType',                  
		    'fileTypeExts':'*.jpg;*.jpge;*.gif;*.png;*.pdf;*.txt;',                  
		    'overrideEvents':['onDialogClose'],                
		     'multi':true,                 
		     'formData':{'url':window.location.search}  //这里我需要得到当前页面url的问号后面的值，所以就用 fromData 这个参数。      
		     });        
		}); 
</script>
</head>
<body>
<form enctype="multipart/form-data" method="post" action="/upload">
    <input type="file" name="image" id="file_upload"/>
</form>
</body>
</html>
