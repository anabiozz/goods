if (window.location.pathname === "/") {
    require("./graphics/graphics")
    require("./pictorial_art/pictorial_art")
}

if (window.location.pathname === "/upload-images") {
    require("./uploads/uploads")
}