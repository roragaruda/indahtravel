function AdmLogout(){
    $.ajax({
        url: "/indahadmin/logout",
        type: "POST",
        success: function(){
            window.location.href="indahwisata";
        }
    });
}