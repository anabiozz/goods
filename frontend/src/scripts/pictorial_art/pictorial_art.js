import {
  ajaxGetImages
} from "../utils/ajax"

$(document).ready(async () => {

  let pictorial_art = await ajaxGetImages(1)
  console.log(pictorial_art)

  if (pictorial_art && pictorial_art.length > 0) {
    const column_number = 3
    for (let index = 0; index < column_number; index++) {
      $("div#pictorial-art").append($("<div class='column'></div>"))
    }
    let i = 0
    pictorial_art.forEach(image => {
      let column = $("div#pictorial-art .column")[i]
      $(column).append("<img src='/static/images/graphics/preview/" + image.UUID + "_thumb.jpg'>")
      i++
      if (i == 3) i = 0
    })
  }
})