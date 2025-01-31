function dataURItoBlob(dataURI) {
    const byteString = atob(dataURI.split(',')[1]);
    const arrayBuffer = new ArrayBuffer(byteString.length);
    const uintArray = new Uint8Array(arrayBuffer);
    for (let i = 0; i < byteString.length; i++) {
        uintArray[i] = byteString.charCodeAt(i);
    }
    // Extract the content type (image/png, image/jpeg, etc.) from data URI
    const mimeType = dataURI.split(',')[0].split(':')[1].split(';')[0];
    return new Blob([uintArray], { type: mimeType });
}
function savePhotos(id) {
    const imgContainer = document.getElementById("img-container");
    const formData = new FormData();

    // Iterate through images in the container
    imgContainer.querySelectorAll("img").forEach((img) => {
        let blob, filename;

        if (img.src.startsWith('data:image')) {
            // Convert data URI to blob
            blob = dataURItoBlob(img.src);
            // Extract file extension from the data URI
            const mimeType = img.src.split(',')[0].split(':')[1].split(';')[0];
            const ext = mimeType.split('/')[1]; // Extract file extension (e.g., 'jpeg')
            filename = `image-${Date.now()}.jpg`;
        } else {
            // If the image is already a URL, fetch it as a blob
            fetch(img.src)
                .then(response => response.blob())
                .then(blob => {
                    const ext = img.src.split('.').pop(); // Get extension from URL
                    const filename = `image-${Date.now()}.jpg`;
                    formData.append('photos', blob, filename);
                })
                .catch(err => console.log(err));
            return; // Skip appending here as it's async
        }

        // Append the blob to FormData with a proper filename
        formData.append('photos', blob, filename);
    });
    $.ajax({
    url:`http://localhost:8080/product/photos/${id}`,
    method:"POST",
    processData:false,
    contentType: false,
    data: formData,
    success: function(){
        console.log("upload succesful");
        }
    })
}

function deletePhotos(id){
    $.ajax({
        url:`http://localhost:8080/product/photos/${id}`,
        method: "DELETE",
        success: function(){
            const div = document.getElementById(`photo-${id}`);
            if(div){
                div.remove();
            }
        },
        error: function(xhr){
            const response = JSON.parse(xhr.responseText);
            const errMsg = `${response.error}`
            console.log(errMsg)
        }
    });
}

function appendPhotos(id){
    $.ajax({
        url:`http://localhost:8080/product/${id}`,
        method:"GET",
        success:function(products){
            $("#nama-destinasi").empty();
            $("#img-container").empty();
            $("#view-footer").empty();
            $("#nama-destinasi").append(`${products.NamaDestinasi}`);
            if (products.Photo){
                products.Photo.forEach(product => {
                    const productPhotos = 
                    `<div class="position-relative me-2 mb-2" id="photo-${product.ID}">
                        <img src="/assets/image/productphoto/${id}/${product.ImageName}" alt="" class="img-thumbnail"
                            style="max-width: 150px;">
                        <button type="button" class="btn-close position-absolute top-0 end-0 remove-img"
                            aria-label="Remove" onclick="deletePhotos(${product.ID})"></button>
                    </div>`;
                    $("#img-container").append(productPhotos);
                })
                $("#view-footer").append(`<button type="button" class="btn btn-primary" id="save-btn" data-bs-dismiss="modal" aria-label="Close" onclick="savePhotos(${id})">Simpan</button>`)
            }else{
                $("#img-container").append(`<div class="position-relative me-2 mb-2" id="${products.ID}"></div>`);
                $("#view-footer").append(`<button type="button" class="btn btn-primary" id="save-btn" data-bs-dismiss="modal" aria-label="Close" onclick="savePhotos(${id})">Simpan</button>`)
            }
        }
    });
}



$(document).ready(function(){
    $.ajax({
        url:"http://localhost:8080/product",
        method:"GET",
        success:function(products){
            $("#product-photo-table").empty();
            const productTable = 
            `<table id="productphoto" class="table table-striped" style="width:100%">
                <thead>
                    <tr>
                        <th>Nama Destinasi</th>
                        <th>Foto</th>
                    </tr>
                </thead>
                <tbody id="product-photo-data">
                </tbody>
            </table>`;
            $("#product-photo-table").append(productTable);
            products.forEach(product => {
                const productTableData = 
                `<tr>
                        <td>${product.NamaDestinasi}</td>
                        <td>
                            <button data-bs-toggle="modal" data-bs-target="#view-modal"
                                class="btn btn-primary" onclick="appendPhotos(${product.ID})">View</button>
                        </td>
                </tr>`;
                $("#product-photo-data").append(productTableData)

            })
            new DataTable('#productphoto');
        },
        error: function(xhr){
            const response = JSON.parse(xhr.responseText);
            const errMsg = `<h2 style="color: red; font-size: 12px;">${response.error}</h2>`
            $("#product-photo-table").append(errMsg);
        }
    });
});