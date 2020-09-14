
player(window.innerWidth-300,window.innerHeight)
$(document).ready(function(){
    $('#my-tab a').click(function (e) { 
        e.preventDefault();//阻止a链接的跳转行为 
        $(this).tab('show');//显示当前选中的链接及关联的content 
    })
    $(".choose-button").click(function(){
        reset($(this).attr('src'))
    });
    $(".tab-content").height(window.innerHeight-$("#my-tab").height()+"px")
    $(".list-group-flush").height(window.innerHeight-27-$("#my-tab").height()+"px");
    $(".list-group>button").first().addClass("active");
    $(".list-group>button").click(function(){
        $(".list-group>button").removeClass('active')
        $(this).addClass('active')
    })

    //进来先加载两次
    loadComment()

    $("#comment_p").scroll(function(){
        top = $("#comment_p").scrollTop()
        if(top == 0){
            loadComment()
        }
    })

    $("#register_tigger").click(function(){
        $("#register-mode").modal()
        captcha("register")
        return
    })
    $("#login_action").click(function(){
        login()
    })
    $("#register_login").click(function(){
        register()
    })
    $('#message_text').bind('keypress',function(event){
        if(event.keyCode == "13")    
        {
            token = localStorage.getItem("token")
            if(token == null || token == ""){
                $("#login-mode").modal()
                captcha("login")
                return
            }
            talk($('#message_text').val())
            $('#message_text').val("")
        }
    });
    $("#redrect_home").click(function(){
        window.location = "/"
    })

    token = localStorage.getItem("token")
    if(token != null && token != ""){
        recv()
    }
})


function captcha(type){
    $.get("/captcha",function(data){
        console.log(data)
        if (type == 'login'){
            $("#login_captcha_id").val(data.data.id)
            $("#audio_captcha_login").attr("src",data.data.base64)
        }else{
            $("#register_captcha_id").val(data.data.id)
            $("#audio_captcha_register").attr("src",data.data.base64)
        }
    })
}

function loadComment(){
    if($("#lastID").attr('data') == ""){
        url = "http://talk.chenyingqiao.top/talk/"+movieHash
    }else{
        url = "http://talk.chenyingqiao.top/talk/"+movieHash+"/"+$("#lastID").attr('data')
    }
    $.get(url,function(data){
        append = ""
        for (i=0;i<data.data.length;i++) {
            if(i==0){
                $("#lastID").attr('data',data.data[i].ID)
            }
            messageItem = data.data[i]
            append += '<li class="list-group-item talk"><span ID="'+messageItem.ID+'">'+messageItem.Username+"</span>: "+messageItem.Content+'</li>'
        }
        $("#comment_p").prepend(append)
    })
}

function login(){
    username = $("#username").val()
    password = $("#password").val()
    answer = $("#login_answer").val()
    capid = $("#login_captcha_id").val()

    if(username == "" || password == "" || answer == "" || capid == ""){
        alert("请输入登录信息不要留空")
        return
    }

    if(username.length < 5){
        alert("用户名过短：需要5个以上字符")
        return
    }

    if(password.length < 5){
        alert("密码过短：需要5个以上字符")
        return
    }

    if(/.*[\u4e00-\u9fa5]+.*$/.test(password)){
        alert("密码不能包含中文")
        return
    }

    $.post("/login",{
        username:username,
        password:password,
        answer:answer,
        capid:capid
    },function(data){
        if(data.code == 200){
            localStorage.setItem("token",data.data.token)
            alert("登录成功")
            $("#login-mode").modal("hide")
            window.location.reload()
        }else{
            alert("登录失败:"+data.message)
            captcha("login")
        }
    })
}

function funcChina(){
      var obj = document.form1.txtName.value;
      if(/.*[\u4e00-\u9fa5]+.*$/.test(obj)) {
        return false;
      }
    
      return true;
    }
    

function register(){
    username = $("#reg_username").val()
    password = $("#reg_password").val()
    passwordRe = $("#reg_password_re").val()
    answer = $("#register_answer").val()
    capid = $("#register_captcha_id").val()

    if(username == "" || passwordRe == "" || password == "" || answer == "" || capid == ""){
        alert("请输入登录信息不要留空")
        return
    }

    if(username.length < 5){
        alert("用户名过短:需要5个以上字符")
        return
    }

    if(password.length < 5){
        alert("密码过短：需要5个以上字符")
        return
    }

    if(/.*[\u4e00-\u9fa5]+.*$/.test(password)){
        alert("密码不能包含中文")
        return
    }


    if(password != passwordRe){
        alert("两次密码输入不同")
        return
    }

    $.post("/register",{
        username:username,
        password:password,
        password_re:passwordRe,
        answer:answer,
        capid:capid
    },function(data){
        console.log(data)
        if(data.code == 200 && data.data.status){
            alert("注册成功")
            $("#register-mode").modal("hide")
            return
        }else{
            alert("注册失败:"+data.message)
            captcha("register")
            return
        }
    })
}

function player(width,height){
    $("#video-panel").width(width)
    $("#video-panel").height(height)
    $("#live-stream").width(width)
    $("#live-stream").height(height)
    
    console.log(arguments)
    var options = {
        width: width,
        "autoplay":true,
        height: height
    }
    playerP = videojs('live-stream', options ,function(){
        this.on('ended',function(){
            //自动播放下一个
            $("button.active").next().trigger("click")
       })
    });

    playerP.on(['loadstart', 'play', 'playing', 'firstplay', 'pause', 'ended', 'adplay', 'adplaying', 'adfirstplay', 'adpause', 'adended', 'contentplay', 'contentplaying', 'contentfirstplay', 'contentpause', 'contentended', 'contentupdate'], function (e) {
        console.warn('VIDEOJS player event: ', e.type);
    });
}
function reset(src){
    playerP.reset()
    playerP.src([
        {
            type: 'application/x-mpegURL',
            src: src
        },
    ])
    playerP.load()
    playerP.play()
}
function talk(msg){
    if(msg == ""){
        return
    }
    $.ajax({
        type:"POST",
        url:"http://talk.chenyingqiao.top/room/"+movieHash,
        data:{
            message:msg
        },
        headers:{
            Authorization:localStorage.getItem("token")
        },
        success:function(data){
            $("#comment_p").scrollTop($("#comment_p").prop("scrollHeight"))
            console.log(data)
        } 
    })
}

function recv(){
    if (!!window.EventSource) {
        var source = new EventSource('http://talk.chenyingqiao.top/stream/'+movieHash+'?token='+localStorage.getItem("token"));
        source.addEventListener('message', function(e) {
            console.log(e.data)
            eventComment = JSON.parse(e.data)
            if (eventComment.heartbeat == true) {
                console.log("heartbeat")
            }else{
                append = '<li class="list-group-item talk">'+eventComment.data.message+'</li>'
                $("#comment_p").append(append)
                $("#comment_p").scrollTop($("#comment_p").height()+100)
            }
        }, false);
    } else {
        alert("NOT SUPPORTED");
    }
}