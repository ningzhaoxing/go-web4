$(document).ready(function() {
    // 获取用户信息
    $.ajax({
        type:"POST",
        url: "http://localhost:8899/user/getOwnInfor",
        contentType: "application/json",
        data: "",
        headers: {
            'Authorization':localStorage.getItem("token")
        },
        success: function (res) {
            console.log("-------",res.data.user.name)
            $('#name1').html(res.data.user.name)
            $('#name2').html(res.data.user.name)
            if (res.data.user.permission_level === 0) {
                $('#role').html('普通用户')
            } else $('#role').html(`数据管理员`)
        },
        error:function (error) {
            console.log(error)
        }
    })
});