let counter_button = document.querySelector(".counter_button")
let counter_head = document.querySelector(".counter_head")
let count = 0

counter_button.addEventListener("click", () => {
    console.log("Hello world")
    count++
    counter_head.innerHTML = count
})

