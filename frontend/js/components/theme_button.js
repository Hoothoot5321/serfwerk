//@ts-check

let theme_button = document.querySelector("#theme_button")

let light_background_color = "#e7c2c1"
let light_text_color = "#153e4c"
let light_accent_color = "#437398"

let dark_background_color = "#241312"
let dark_text_color = "#B3DCEA"
let dark_accent_color = "#437398"

/**
    *@param {Element} button
    *@param {boolean} light_theme
    *@param {HTMLElement} root_style
    */
function buttonToggler(button, light_theme, root_style) {
    if (!light_theme) {
        button.innerHTML = "🌞"

        button.classList.remove("dark")
        button.classList.add("light")

        root_style.style.setProperty("--bg-col", dark_background_color)
        root_style.style.setProperty("--col", dark_text_color)
        root_style.style.setProperty("--accent", dark_accent_color)
    }
    else {

        button.innerHTML = "🌝"
        button.classList.remove("light")
        button.classList.add("dark")

        root_style.style.setProperty("--bg-col", light_background_color)
        root_style.style.setProperty("--col", light_text_color)
        root_style.style.setProperty("--accent", light_accent_color)
    }
}

function initThemeButton() {
    if (!theme_button) {
        return;
    }
    let theme = window.localStorage.getItem("theme")
    if (!theme) {
        console.log("Hello")
        window.localStorage.setItem("theme", "light")
        theme = "light"
    }
    let light_theme = theme == "light"
    /**@type {HTMLElement|null} */
    let root = document.querySelector(":root")
    if (!root) {
        return
    }
    buttonToggler(theme_button, light_theme, root)

    theme_button.addEventListener("click", () => {
        if (!theme_button) {
            return
        }

        /**@type {HTMLElement|null} */
        let root = document.querySelector(":root")
        if (!root) {
            return
        }
        let light_theme = window.localStorage.getItem("theme") == "light"
        if (light_theme) {
            window.localStorage.setItem("theme", "dark")
        }
        else {
            window.localStorage.setItem("theme", "light")
        }
        buttonToggler(theme_button, !light_theme, root)

    })
}
initThemeButton()
