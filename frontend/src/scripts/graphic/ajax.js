export const ajaxGetGraphics = async () => {
    let result
    try {
        result = await $.ajax({
            url: "http://localhost:8080/get-graphics",
            dataType: "json"
        })
        return result
    } catch (error) {
        if (error.responseText != undefined) {
            console.error(error.responseText)
        }
    }
}