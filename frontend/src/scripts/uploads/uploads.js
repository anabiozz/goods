import {
    ajaxSaveImages
} from "./ajax"

$(document).ready(function () {

    $("#save_images").click(function () {

        let Images = []

        $.each($($(".images_row").children()), (index, image_container) => {

            let Image = {
                UUID: "",
                Name: "",
                Materials: "",
                Year: "",
                Size: "",
                Type: 0,
                Ext: "",
                IsForSale: false
            }

            let image = $(image_container).children()[0]
            let image_info = $(image_container).children()[1]
            let img = $(image).children()[0]

            Image.UUID = $(img).attr("alt")

            $.each($(image_info).children(), (index, input_field) => {

                let input = $(input_field).children()[0]

                if (input.id == "image_name") {
                    Image.Name = $(input).val()
                }

                if (input.id == "image_materials") {
                    Image.Materials = $(input).val()
                }

                if (input.id == "image_year") {
                    Image.Year = $(input).val()
                }

                if (input.id == "image_size") {
                    Image.Size = $(input).val()
                }

                let option_selected = $("#image_type option:selected")
                if (option_selected) {
                    Image.Type = option_selected.val()
                }

                if (input.id == "checkbox") {
                    if ($(input).is(":checked")) {
                        Image.IsForSale = true
                    } else {
                        Image.IsForSale = false
                    }
                }
            })
            Images.push(Image)
        })
        ajaxSaveImages(Images)
        window.location.replace("http://localhost:8080/")
    })
})