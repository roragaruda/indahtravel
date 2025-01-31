function deleteOrder(id, name){
    $("#confirm-message").empty().append(`Confirm Delete Order for:${name}?`)
    $("#delBtn").on("click", function(){
        $.ajax({
            url: `orders/${id}`,
            method:"DELETE",
            success: function(){
                window.location.href="indahadmin/orders";
            }
        })
    })
}

function changeStatus(id, status){
    var formData = new FormData();
    formData.append("status", status);
    $.ajax({
        url: `orders/edit/${id}`,
        method:"PUT",
        processData:false,
        contentType: false,
        data: formData,
        success: function(){
            window.location.href="indahadmin/orders";
        }
    });

}

function appendView(id, image){
    $("#view-proof").empty().append(`<img style = "width: 500px; height: 500px;" src="/assets/image/bukti/${id}/${image}" alt="">`);
}


$(document).ready(function(){
    $.ajax({
        url:"orders",
        method:"GET",
        success: function(orders){
            $("#unconfirmed").empty();
            $("#confirmed").empty();
            orders.forEach(order => {
                $.ajax({
                    url:``
                })
                if (order.Status === "Pending"){
                    $("#unconfirmed").append(`
                        <tr>
                        <td>${order.NamaPeserta}</td>
                        <td>${order.NoTelp}</td>
                        <td>${order.JmlPeserta}</td>
                        <td><button type="button" class="btn btn-warning" style="color: white;">Pending</button></td>
                        <td><button data-bs-toggle="modal" data-bs-target="#view-modal" class="btn btn-primary" onclick="appendView(${order.ID},'${order.FotoBukti}')">View</button></td>
                        <td>
                            <button class="btn btn-success" onclick="changeStatus(${order.ID},'Confirmed')">Confirm</button>
                            <button data-bs-toggle="modal" data-bs-target="#delete-modal" class="btn btn-danger" onclick="deleteOrder(${order.ID},'${order.NamaPeserta}')">Delete</button>
                        </td>
                    </tr>`);
                }else{
                    $("#confirmed").append(`
                        <tr>
                        <td>${order.NamaPeserta}</td>
                        <td>${order.NoTelp}</td>
                        <td>${order.JmlPeserta}</td>
                        <td><button type="button" class="btn btn-success" style="color: white;">Confirmed</button></td>
                        <td><button data-bs-toggle="modal" data-bs-target="#view-modal" class="btn btn-primary" onclick="appendView(${order.ID},'${order.FotoBukti}')">View</button></td>
                        <td>
                            <button class="btn btn-warning" onclick="changeStatus(${order.ID},'Pending')">Unconfirm</button>
                            <button data-bs-toggle="modal" data-bs-target="#delete-modal" class="btn btn-danger" onclick="deleteOrder(${order.ID},'${order.NamaPeserta}')">Delete</button>
                        </td>
                    </tr>`);
                }
                
            });
            new DataTable('#ordertable');
            new DataTable('#orderconfirmedtable');
        }
    });
})