//@ts-check

import { handleResp } from "../helpers/resp_handler.js"

/**@type{HTMLInputElement | null} */
let file_upload = document.querySelector("#file_upload")

/**@type{HTMLInputElement | null} */
let file_upload_button = document.querySelector(".upload_button")

/**@type{HTMLInputElement | null} */
let app_input = document.querySelector("#app_name")

/**@type{HTMLInputElement | null} */
let run_cmd_input = document.querySelector("#run_cmd")

/**@type{HTMLInputElement | null} */
let build_cmd_input = document.querySelector("#build_cmd")

/**@type{HTMLInputElement | null} */
let upload_button = document.querySelector("#upload_button")

/**@type{HTMLInputElement | null} */
let port_input = document.querySelector("#port_num")

/**@type{HTMLInputElement | null} */
let file_titel = document.querySelector(".file_title")

/**@type{HTMLInputElement | null} */
let ram = document.querySelector("#ram")

/**@type{HTMLInputElement | null} */
let lager = document.querySelector("#lager")

/**@type{HTMLInputElement | null} */
let cpu = document.querySelector("#cpu")

/**@type{HTMLInputElement | null} */
let cpu_kraft = document.querySelector("#cpu_kraft")

/**@type{HTMLInputElement | null} */
let pris = document.querySelector("#pris")

/**@type{HTMLInputElement | null} */
let status_msg = document.querySelector(".status_text")

async function setupNewApp() {
    if (!file_upload || !file_upload_button || !app_input || !run_cmd_input || !build_cmd_input || !upload_button || !file_titel || !status_msg || !port_input || !pris || !ram || !lager || !cpu || !cpu_kraft) {
        console.log("Shit")
        return
    }
    let lager_obj = window.sessionStorage.getItem("lager_info")
    if (!lager_obj) {
        return
    }
    let lager_obj_parsed = JSON.parse(lager_obj)
    pris.innerHTML = "Pris: " + lager_obj_parsed.pris + "kr."
    ram.innerHTML = "Ram: " + lager_obj_parsed.ram
    cpu.innerHTML = "CPU: " + lager_obj_parsed.cpu
    lager.innerHTML = "Lager: " + lager_obj_parsed.lager
    cpu_kraft.innerHTML = "CPU Kraft: " + lager_obj_parsed.cpu_kraft + "%"

    file_upload_button.addEventListener("click", () => {
        if (!file_upload || !file_titel) {
            return
        }
        file_upload.click()

    })
    file_upload.addEventListener("change", () => {
        if (!file_titel || !file_upload) {
            return
        }
        let files = file_upload.files
        if (!files) {
            return
        }
        file_titel.innerHTML = files[0].name
    })
    upload_button.addEventListener("click", async () => {
        if (!file_upload || !app_input || !run_cmd_input || !build_cmd_input || !status_msg || !port_input) {
            return
        }
        let files = file_upload.files
        if (!files) {
            return
        }
        let file = files[0]
        console.log(file.name)
        const form_data = new FormData()

        form_data.append("file", file)
        form_data.append("appName", app_input.value)
        form_data.append("runCmd", run_cmd_input.value)
        form_data.append("buildCmd", build_cmd_input.value)
        form_data.append("portNum", port_input.value)

        form_data.append("ram", lager_obj_parsed.ram)
        form_data.append("pris", lager_obj_parsed.pris)
        form_data.append("lager", lager_obj_parsed.lager)
        form_data.append("cpu", lager_obj_parsed.cpu)
        form_data.append("cpu_kraft", lager_obj_parsed.cpu_kraft)

        let resp = await fetch("/api/new_app", {
            method: "POST",
            body: form_data,
        })
        let j_resp = await resp.json()
        handleResp(status_msg, j_resp)
    })
}

setupNewApp()
