
//@ts-check
import { actStatus } from "../helpers/status_act.js"
import { handleResp } from "../helpers/resp_handler.js"
/**@type {HTMLInputElement|null} */
let username_input = document.querySelector("#username_input")
/**@type {HTMLInputElement|null} */
let password_input = document.querySelector("#password_input")

let new_user_button = document.querySelector("#user_button")
let login_button = document.querySelector("#login_button")

let email_reg = "^[\\w-\\.]+@([\\w-]+\\.)+[\\w-]{2,4}$"

let brugernavn_reg = "[^A-Za-z0-9]"

let kodeord_reg = "^[^<>{}\"/|;:.,~%^=*\\]\\\\()\\[¿§«»ω⊙¤°℃℉€¥£¢¡®©0-9_+]*$"

/**@type {HTMLInputElement|null} */
let status_msg = document.querySelector(".status_text")


async function initLogin() {
    if (!new_user_button || !login_button || !username_input || !password_input || !status_msg) {
        return;
    }
    login_button.addEventListener("click", async () => {
        if (!username_input || !password_input || !status_msg) {
            return
        }
        let username = username_input.value
        let password = password_input.value

        if (username.match(email_reg)) {
        }
        else if (username.length > 6 && username.length < 12 && !username.match(brugernavn_reg)) {
        }
        else {
            actStatus(status_msg, true, "Ikke valid email eller brugernavn")
            return
        }
        if (password.length < 6) {
            actStatus(status_msg, true, "Kodeord skal være længere end 6")
            return
        }
        if (password.length > 12) {
            actStatus(status_msg, true, "Kodeord skal være kortere end 12")
            return
        }

        if (!password.match(kodeord_reg)) {
            actStatus(status_msg, true, "Kodeord må ikke indeholde specielle tegn")
            return
        }
        actStatus(status_msg, false, "OK")
        let resp_body = { username: username, password: password }
        let str_body = JSON.stringify(resp_body)

        let resp = await fetch("/api/login", {
            method: "POST",
            headers: {
                "Content-Type": "application/json"
            },
            body: str_body
        })
        let str_resp = await resp.json()
        await handleResp(status_msg, str_resp)

    })

    new_user_button.addEventListener("click", () => {
        location.href = "/ny_bruger"
    })
}
initLogin()
