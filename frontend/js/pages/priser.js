//@ts-check
/**@type {HTMLInputElement|null} */
let ram_input = document.querySelector("#ram_input")
/**@type {HTMLInputElement|null} */
let ram_slider = document.querySelector("#ram_slider")

/**@type {HTMLInputElement|null} */
let lager_input = document.querySelector("#lager_input")
/**@type {HTMLInputElement|null} */
let lager_slider = document.querySelector("#lager_slider")

/**@type {HTMLInputElement|null} */
let cpu_input = document.querySelector("#cpu_input")
/**@type {HTMLInputElement|null} */
let cpu_slider = document.querySelector("#cpu_slider")

/**@type {HTMLInputElement|null} */
let cpu_kraft_input = document.querySelector("#cpu_kraft_input")
/**@type {HTMLInputElement|null} */
let cpu_kraft_slider = document.querySelector("#cpu_kraft_slider")

let price_marker = document.querySelector("#pricer")

/**

    *@param {HTMLInputElement} ram_slider
    *@param {HTMLInputElement} lager_slider
    *@param {HTMLInputElement} cpu_slider
    *@param {HTMLInputElement} kraft_slider
    */
function calcPrice(ram_slider, lager_slider, cpu_slider, kraft_slider) {
    if (!price_marker) {
        return
    }
    let price = 0;
    price += (parseInt(ram_slider.value) - 1024) / 512 * 20
    price += (parseInt(lager_slider.value) - 2048) / 512 * 5
    price += (parseInt(cpu_slider.value) - 1) * 50
    price += (parseInt(kraft_slider.value) - 50) / 5 * 5

    price_marker.innerHTML = "Price: " + price.toString() + "kr."
}

/**
    *@param {HTMLInputElement} input
    *@param {number} step 
    *@returns {HTMLInputElement}
    */
function roundNum(input, step) {
    let input_value = parseFloat(input.value) / step
    let new_val = Math.round(input_value) * step
    input.value = new_val.toString()
    return input
}

/**
    *@param {HTMLInputElement} input
    *@param {number} min
    *@param {number} max
    *@returns {HTMLInputElement}
    */
function setMinMax(input, min, max) {
    let input_value = parseInt(input.value)
    if (input_value > max) {
        input.value = max.toString()
    }
    if (input_value < min) {
        input.value = min.toString()
    }
    return input
}

/**
    *@param {HTMLInputElement} input
    *@param {HTMLInputElement} slider
    *@param {number} min
    *@param {number} max
    *@param {number} step 
    */
function makePair(input, slider, min, max, step) {

    slider.addEventListener("input", () => {

        if (!ram_slider || !lager_slider || !cpu_slider || !cpu_kraft_slider) {
            return;
        }
        setMinMax(slider, min, max)
        input.value = slider.value
        calcPrice(ram_slider, lager_slider, cpu_slider, cpu_kraft_slider)
    })
    input.addEventListener("change", () => {

        if (!ram_slider || !lager_slider || !cpu_slider || !cpu_kraft_slider) {
            return;
        }
        roundNum(input, step)
        setMinMax(input, min, max)
        slider.value = input.value

        calcPrice(ram_slider, lager_slider, cpu_slider, cpu_kraft_slider)
    })

}

function slidersSetup() {
    if (!ram_input || !ram_slider || !lager_input || !lager_slider || !cpu_input || !cpu_slider || !cpu_kraft_input || !cpu_kraft_slider || !price_marker) {
        console.log("Shit")
        return;
    }
    console.log("Hello")
    makePair(ram_input, ram_slider, 1024, 4096, 512)
    makePair(lager_input, lager_slider, 2048, 30720, 512)
    makePair(cpu_input, cpu_slider, 1, 4, 1)
    makePair(cpu_kraft_input, cpu_kraft_slider, 50, 100, 5)
}
slidersSetup()
