function dateFormatter(date){
    const [year, month, day] = date.split('-');
    const formattedDate = `${day}/${month}/${year}`;
    return formattedDate;
}

function tdateFormatter(date){
    var newDate = date.split("T");
    newDate = dateFormatter(newDate[0]);
    return newDate;
}


function productView(id){
    $("#view-product").empty();
    $.ajax({
        url:`product/${id}`,
        method: "GET",
        success: function(products){
            const startDate = tdateFormatter(products.StartDate);
            const endDate = tdateFormatter(products.EndDate); 
            $("#view-product").append(`
                <h5><strong>Nama Destinasi:</strong></h5>
                    <h5>${products.NamaDestinasi}</h5>
                    <h5><strong>Harga:</strong></h5>
                    <h5>Rp. ${products.Harga}</h5>
                    <h5><strong>Start Date:</strong></h5>
                    <h5>${startDate}</h5>
                    <h5><strong>End Date:</strong></h5>
                    <h5>${endDate}</h5>
                    <h5><strong>Deskripsi:</strong></h5>
                    <h5>${products.Deskripsi}</h5>
                    <h5><strong>Thumbnail:</strong></h5>
                    <img style = "width: 150px; height: 150px;" src="/assets/image/thumbnail/${id}/${products.Thumbnail}" alt="">`);
        }
    });
};

function productEdit(id){
    $("#product-edit").empty();
    $.ajax({
        url:`product/${id}`,
        method: "GET",
        success: function(products){
            var startDate = products.StartDate.split("T");
            startDate = startDate[0];
            var endDate = products.EndDate.split("T");
            endDate = endDate[0]; 
            $("#product-edit").append(`
                <div class="mb-3">
                    <label for="exampleFormControlInput1" class="form-label">Nama Destinasi</label>
                    <input type="text" class="form-control" id="edit-nama-destinasi" value="${products.NamaDestinasi}" required>
                </div>
                <div class="mb-3">
                    <label for="exampleFormControlInput1" class="form-label">Harga</label>
                    <input type="number" class="form-control" id="edit-harga" value="${products.Harga}" required>
                </div>
                <div class="mb-3">
                    <label for="exampleFormControlInput1" class="form-label">Start Date</label>
                    <input type="date" class="form-control" id="edit-startdateInput" value="${startDate}" required/>
                </div>
                <div class="mb-3">
                    <label for="exampleFormControlInput1" class="form-label">End Date</label>
                    <input type="date" class="form-control" id="edit-enddateInput" value="${endDate}" required/>
                </div>
                <div class="mb-3">
                    <label for="exampleFormControlInput1" class="form-label">Deskripsi</label>
                    <textarea class="form-control" id="edit-deskripsi" rows="5" required>${products.Deskripsi}</textarea>
                </div>
                <div class="mb-3">
                    <label for="formFile" class="form-label">Thumbnail</label>
                    <input class="form-control" type="file" id="edit-thumbnail" accept="image/*" required>
                </div>`);

                $("#saveBtn").on("click", function(){
                    $("#edit-footer").empty();
                    var formData = new FormData();
                    const startDate = dateFormatter($("#edit-startdateInput").val());
                    const endDate = dateFormatter($("#edit-enddateInput").val());
                    formData.append("edit-namadestinasi", $("#edit-nama-destinasi").val());
                    formData.append("edit-harga", $("#edit-harga").val());
                    formData.append("edit-deskripsi", $("#edit-deskripsi").val());
                    formData.append("edit-startdate", startDate);
                    formData.append("edit-enddate", endDate);
                    formData.append("edit-thumbnail", $("#edit-thumbnail")[0].files[0]);
                    
                    console.log("ERROR");
                    console.log($("#edit-startdateInput").val());
            
                    $.ajax({
                        url: `product/edit/${products.ID}`,
                        type: "PUT",
                        processData:false,
                        contentType: false,
                        data: formData,
                        success: function(){
                            window.location.href="/indahadmin/product";
                        },
                        error: function(xhr){
                            const response = JSON.parse(xhr.responseText);
                            const errMsg = `${response.error}`
                            $("#edit-footer").append(errMsg);
                        }
                    });
                });
        }
    });
}

function productDelete(id, name){
    $("#confirm-message").append(`<h5>Confirm Delete Product: ${name}?</h5>`)
    $("#delBtn").on("click", function(){
        $("#delete-footer").empty();
        $.ajax({
            url:`product/delete/${id}`,
            method: "DELETE",
            success: function(){
                window.location.href="/indahadmin/product";
            },
            error: function(xhr){
                const response = JSON.parse(xhr.responseText);
                const errMsg = `${response.error}`
                $("#delete-footer").append(errMsg);
            }
        });
    })
}



$(document).ready(function(){
    $.ajax({
        url:"product",
        method:"GET",
        success:function(products){
            $("#product-table").empty();
            const productTable = 
            `<table id="productdata" class="table table-striped" style="width:100%">
            <thead>
                <tr>
                    <th>Thumbnail</th>
                    <th>Destinasi</th>
                    <th>Harga</th>
                    <th>Start Date</th>
                    <th>End Date</th>
                    <th style="width: 45%">Deskripsi</th>
                    <th style="width: 15%;">Action</th>
                </tr>
            </thead>
            <tbody id="product-table-data">
            </tbody>
            </table>`;
            $("#product-table").append(productTable);
            products.forEach(product => {
                const startDate = tdateFormatter(product.StartDate);
                const endDate = tdateFormatter(product.EndDate);
                const productTableData = 
                `<tr>
                    <td><img style = "width: 150px; height: 150px;" src="/assets/image/thumbnail/${product.ID}/${product.Thumbnail}" alt=""></td>
                    <td>${product.NamaDestinasi}</td>
                    <td>${product.Harga}</td>
                    <td>${startDate}</td>
                    <td>${endDate}</td>
                    <td>${product.Deskripsi}</td>
                    <td>
                        <button data-bs-toggle="modal" data-bs-target="#view-modal" class="btn btn-primary" onclick="productView(${product.ID})">View</button>
                        <button data-bs-toggle="modal" data-bs-target="#edit-modal" class="btn btn-warning" id="edit-button" onclick="productEdit(${product.ID})">Edit</button>
                        <button data-bs-toggle="modal" data-bs-target="#delete-modal" class="btn btn-danger" id="del-button" onclick="productDelete(${product.ID}, '${product.NamaDestinasi}')">Delete</button>
                    </td>
                </tr>`;
                $("#product-table-data").append(productTableData)

            })
            new DataTable('#productdata');
        },
        error: function(xhr){
            const response = JSON.parse(xhr.responseText);
            const errMsg = `<h2 style="color: red; font-size: 12px;">${response.error}</h2>`
            $("#product-table").append(errMsg);
        }
    });

    $("#add-product").on("click", function(){
        $("#footer").empty();
    });

    $("#edit-button").on("click", function(){
        $("#edit-footer").empty();
    });

    $("#del-button").on("click", function(){
        $("#delete-footer").empty();
    });

    $("#addBtn").on("click", function(){
        $("#footer").empty();
        var formData = new FormData();
        const startDate = dateFormatter($("#startdateInput").val());
        const endDate = dateFormatter($("#enddateInput").val());
        formData.append("namadestinasi", $("#nama-destinasi").val());
        formData.append("harga", $("#harga").val());
        formData.append("deskripsi", $("#deskripsi").val());
        formData.append("startdate", startDate);
        formData.append("enddate", endDate);
        var imgInput = $("#thumbnail")[0].files[0]
        if(imgInput){
            formData.append("thumbnail", $("#thumbnail")[0].files[0]);
            $.ajax({
                url: "product/add",
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
                    $("#footer").append(errMsg);
                }
            });
        }else{
            $("#footer").append("Thumbnail required.")
        }

        
    });

    

})