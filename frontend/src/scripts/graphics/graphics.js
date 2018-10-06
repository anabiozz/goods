import {
  ajaxGetImages
} from "../utils/ajax"

$(document).ready(async () => {

  let graphics = await ajaxGetImages(0)

  if (graphics && graphics.length > 0) {
    const column_number = 3
    for (let index = 0; index < column_number; index++) {
      $("div#graphics").append($("<div class='column'></div>"))
    }
    let i = 0
    graphics.forEach(image => {
      let column = $("div#graphics .column")[i]
      $(column).append("<img src='/static/images/graphics/preview/" + image.UUID + "_thumb.jpg'>")
      i++
      if (i == 3) i = 0
    })
  }
})