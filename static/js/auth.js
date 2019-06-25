function queryParams() {
    var username = localStorage.getItem("username");
    var token = localStorage.getItem("token");
    return 'username=' + username + '&token=' + token;
}

function hasToken(){
    return localStorage.getItem("token") != null;
}