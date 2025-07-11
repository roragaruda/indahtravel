$(document).ready(function(){
    $("#signinBtn").on("click", function(){
        $("#login-footer").empty();
        formData = new FormData();
        formData.append("username", $("#username").val());
        formData.append("password", $("#password").val());

        $.ajax({
            url: "/indahadmin/login",
            type: "POST",
            processData:false,
            contentType: false,
            data: formData,
            success: function(){
                window.location.href="/indahadmin/product";
            },
            error: function(xhr){
                const response = JSON.parse(xhr.responseText);
                const errMsg = `${response.error}`
                $("#login-footer").append(errMsg);
            }
        });
    });
});