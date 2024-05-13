


async function setup() {
    let path = window.location.pathname

    console.log(path)

    let base_url = window.location.origin
    console.log(base_url)
    let res = await fetch(base_url + "/pre" + path, {
        method: "GET",
        headers: {
            "Authorization": "Jhon",
        }
    })
    let html = await res.text()
    let html_elem = document.querySelector("html")
    let new_html_elem = document.createElement("html")
    new_html_elem.innerHTML = html
    let header = new_html_elem.querySelector("head")

    let all_scripts = new_html_elem.querySelectorAll("script")
    for (let i = 0; i < all_scripts.length; i++) {
        let new_script = document.createElement("script")
        if (all_scripts[i].src) {
            new_script.src = all_scripts[i].src
        }
        new_script.innerHTML = all_scripts[i].innerHTML
        new_script.defer = true
        all_scripts[i].parentElement.removeChild(all_scripts[i])
        header.appendChild(new_script)
    }

    html_elem.replaceWith(new_html_elem)
}
await setup()


