let arr = [];

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

function appendDeleteGallery(id, name){
    $("#confirm-message").append(`Confirm Delete Gallery: ${name} ?`)
    $("#delBtn").on("click", function(){
        $.ajax({
            url:`gallery/${id}`,
            method:"DELETE",
            success:function(){
                console.log("delete succesful");
                window.location.href="indahadmin/gallery";
            },
            error: function(xhr){
                const response = JSON.parse(xhr.responseText);
                const errMsg = `${response.error}`
                console.log(errMsg);
            }
        });
    })
}

function addPhotos(id, formData){
    formData.forEach((value, key) => {
        console.log(`${key}:`, value);
    });
    $.ajax({
        url:`gallery/photos/add/${id}`,
        method:"POST",
        processData:false,
        contentType: false,
        data: formData,
        success: function(){
            console.log("Upload succesful");
            window.location.href="indahadmin/gallery";
        },
        error: function(xhr){
            const response = JSON.parse(xhr.responseText);
            const errMsg = `${response.error}`
            console.log(errMsg);
        }
    });
    formData.forEach((value, key) => {
        console.log(`${key}:`, value);
    });
}
function removeDiv(id){
    arr.push(id);
    const div = document.getElementById(`photo-${id}`);
    if(div){
        div.remove();
    }
}

function appendGalleryPhotos(id, name){
    $.ajax({
        url:`gallery/${id}`,
        method:"GET",
        success:function(gallery){
            $("#img-container").empty();
            $("#title").empty().append(`<strong>Nama Gallery:</strong>
                            <input type="text" id="nama-input" class="form-control" value="${name}" />`)
            console.log("THE ID:",id)
            if (gallery.Photo){
                gallery.Photo.forEach(gallery => {
                    const galleryPhotos = 
                    `<div class="position-relative me-2 mb-2" id="photo-${gallery.ID}">
                        <img src="/assets/image/gallery/${id}/${gallery.FotoGallery}" alt="" class="img-thumbnail"
                            style="max-width: 150px;">
                        <button type="button" class="btn-close position-absolute top-0 end-0 remove-img"
                            aria-label="Remove" onclick="removeDiv(${gallery.ID})"></button>
                    </div>`;
                    $("#img-container").append(galleryPhotos);
                })
                $("#modal-footer").empty().append(`<button type="button" class="btn btn-primary" id="simpan-btn" data-bs-dismiss="modal" aria-label="Close">Simpan</button>`)
            }else{
                $("#img-container").append(`<div class="position-relative me-2 mb-2" id="photo-${gallery.ID}"></div>`);
                $("#modal-footer").empty().append(`<button type="button" class="btn btn-primary" id="simpan-btn" data-bs-dismiss="modal" aria-label="Close">Simpan</button>`)
            }
        }
    });

    $("#simpan-btn").on("click", function(){
        var formData = new FormData;
        formData.append("gallery-name", $("#nama-input").val());
        console.log($("#nama-input").val())
        $.ajax({
            url:`gallery/edit/${id}`,
            method:"PUT"
        })
        if(arr.length>0){
            for (let i = 0; i < arr.length; i++){
                $.ajax({
                    url:`gallery/photos/${arr[i]}`,
                    method:"DELETE"
                });
            };
        };
        arr = [];
        const imgContainer = document.getElementById("img-container");
        const imgformData = new FormData();

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
                        formData.append("gallery-photos", blob, filename);
                    })
                    .catch(err => console.log(err));
                return; // Skip appending here as it's async
            }

            // Append the blob to FormData with a proper filename
            imgformData.append("gallery-photos", blob, filename);
        });
        imgformData.forEach((value, key) => {
            console.log(`${key}:`, value);
        });
        addPhotos(id, imgformData);
        // $.ajax({
        //     url:`gallery/photos/add/${id}`,
        //     method:"POST",
        //     processData:false,
        //     contentType: false,
        //     data: imgformData,
        //     success: function(){
        //         console.log("upload succesful");
        //     },
        //     error: function(xhr){
        //         const response = JSON.parse(xhr.responseText);
        //         const errMsg = `${response.error}`
        //         console.log(errMsg);
        //     }
        // })

    });
}


$(document).ready(function(){
    $.ajax({
        url:"gallery",
        method:"GET",
        success: function(galleries){
            $("#gallery-data").empty();
            $("#gallery-data").append(
                `<table id="gallery-table" class="table table-striped" style="width:100%">
                <thead>
                    <tr>
                        <th>Nama Gallery</th>
                        <th>Action</th>
                    </tr>
                </thead>
                <tbody id="gallery-table-data">
                </tbody>
                </table>`);
            galleries.forEach(gallery => {
                $("#gallery-table-data").append(`
                    <tr>
                        <td>${gallery.GalleryName}</td>
                        <td>
                            <button data-bs-toggle="modal" data-bs-target="#view-modal"
                                class="btn btn-primary" onclick="appendGalleryPhotos(${gallery.ID},'${gallery.GalleryName}')">View</button>
                            <button data-bs-toggle="modal" data-bs-target="#delete-modal"
                                class="btn btn-danger" onclick="appendDeleteGallery(${gallery.ID},'${gallery.GalleryName}')">Delete</button>
                        </td>
                    </tr>`);
            })
            new DataTable('#gallery-table');
        }
    });

    $("#add-gallery").on("click", function(){
        var formData = new FormData();
        formData.append("gallery-name", $("#galleryname").val())
        $.ajax({
            url:"gallery/add",
            method:"POST",
            processData:false,
            contentType: false,
            data: formData,
            success: function(response){
                const fileInput = document.getElementById("file-input");
                const files = fileInput.files;
                const formData2 = new FormData();

                for (const file of files){
                    formData2.append("gallery-photos", file);
                }
                addPhotos(response.ID, formData2);
            }
        })
    });
})