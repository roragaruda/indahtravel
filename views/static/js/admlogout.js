function AdmLogout(){
    $.ajax({
        url: "http://localhost:8080/indahadmin/logout",
        type: "POST",
        success: function(){
            window.location.href="http://localhost:8080/indahwisata";
        }
    });
}