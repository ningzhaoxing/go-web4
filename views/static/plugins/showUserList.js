var totalItems = 0
var itemNumPage = 8
var totalPages = Math.ceil(totalItems / itemNumPage)
var currentPage = 1
var items

function goToPage(page) {

    if (page > totalPages) {
        page = totalPages
    }
    if (page < 1) {
        page = 1
    }
    currentPage = page
    sendAjax()
    updatePage()
    // console.log('跳转到第' + page + '页');
}

function updatePage() {
    var pagination = document.getElementById("pagination")
    pagination.find('li:not(:first, :last)').remove(); // 移除除了首页和尾页之外的所有页码链接
    console.log(totalPages)
    for (var i = 1; i <= totalPages; i++) {

        if (i === currentPage) {
            // 当前页码不需要链接，用span显示
            pagination.append('<li class="active"><a href="#">' + i + '</a></li>');
        } else {
            pagination.append('<li><a href="#" onclick="goToPage(' + i + ')">' + i + '</a></li>');
        }
    }
}

// 获取用户列表
function sendAjax() {
    var d = {
        page: currentPage,
        limit: itemNumPage
    }

    $.ajax({
        type:"POST",
        url: "http://localhost:8899/user/list",
        contentType: "application/json",
        data: JSON.stringify(d),
        headers: {
            'Authorization':localStorage.getItem("token")
        },
        success: function (res) {
            totalItems=res.data.num
            items = res.data.list
            totalPages = Math.ceil(totalItems / itemNumPage)
            updatePage()
        },
        error:function (error) {
            console.log(error)
        }
    })
}


$(document).ready(function() {
    goToPage(1)
});