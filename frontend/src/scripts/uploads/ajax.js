export const ajaxSaveImages = (images) => {
    let result 
    try {
        result = $.ajax({
            url: "http://localhost:8080/save-images",
            dataType: "json",
            method: "POST",
            data: JSON.stringify(images)
        })
        return result
    } catch (error) {
        if (error.responseText != undefined) {
            console.error(error.responseText)
        }
    }
}