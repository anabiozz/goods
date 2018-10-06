export const ajaxGetImages = async (imageType) => {
    let result
    try {
        result = await $.ajax({
            url: "http://localhost:8080/api/get-images?imageType=" + imageType,
            dataType: "json"
        })
        return result
    } catch (error) {
        if (error.responseText != undefined) {
            console.error(error.responseText)
        }
    }
}