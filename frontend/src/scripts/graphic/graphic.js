import {
    ajaxGetGraphics
} from './ajax'

function GetGraphics() {
    return ajaxGetGraphics()
}

$(document).ready(async () => {

    let graphics = await GetGraphics()
    console.log(graphics) 

    //   const dir = "static/images/graphic/"
    //   const fileextension = ".jpg"
    //   const column_number = 3
    //   $.ajax({
    //       url: dir,
    //       success: function (data) {
    //           for (let index = 0; index < column_number; index++) {
    //             $(".main-content").append($("<div class='column'></div>"))
    //           }
    //           let i = 0
    //           $(data).find("a:contains(" + fileextension + ")").each(function () {
    //               var filename = this.href.replace(window.location.host, "").replace("http://", "")
    //               let column = $(".main-content .column")[i]
    //               $(column).append("<img src='" + dir + filename + "'>")
    //               i++
    //               if (i == 3) i = 0
    //           })
    //       }
    //   })
})