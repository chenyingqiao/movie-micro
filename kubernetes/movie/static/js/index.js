
    $(function(){
    $(document).on("mouseover",".movie",function(env){
        $(this).find(".cover").show()
    })
    $(document).on("mouseout",".movie",function(env){
        $(this).find(".cover").hide()
    })
    load()
    $(document).on("click",".movie",function(){
        window.open($(this).attr("href"), '_blank').location;
    })
    $(".menu>a").click(function(){
        $("div.movie-content").children().remove()
        $("#search").val("")
        $(".menu>a").removeClass("active")
        $(this).addClass("active")
        load()
    })
    $('#search').bind('keypress',function(event){
        if(event.keyCode == "13")    
        {
            $("div.movie-content").children().remove()
            console.log($("#search").val())
            load()
        }
    })
});
var jlheight = 50;             //距下边界长度/单位px
var totalheight = 0;
var is_load = 0;
var pages = 2;
var n_i = 0;
$(window).scroll(function () {
    var srollPos = $(window).scrollTop();    //滚动条距顶部距离(页面超出窗口的高度)
    totalheight = parseFloat($(window).height()) + parseFloat(srollPos);
    is_load = $("#hid_isload").val();
    if (is_load == 0) {
        if (($(document).height() - jlheight) <= totalheight) {
            $("#hid_isload").attr("value", "1");//滚动的时候设置成1：不执行滚动
            console.log("加载....");
            load()
        }
    }
});

function load(){
    url = "/list/"
    lastOid = $(".movie>div.cover").last().attr("oid")
    if(lastOid != undefined) {
        url+=lastOid
    }
    url += "?"
    keyword = $("#search").val()
    if(keyword != ""){
        url+="keyword="+keyword
    }
    type = $(".menu>a[class='list-group-item active']").attr("data")
    if(type != undefined){
        url+="type="+type
    }
    $.get(url,function(content,status){
        if (status != "success") {
            return 
        }
        $("div.movie-content").append(content)
        $("#hid_isload").attr("value", "0")
    })
    console.log(lastOid)
}