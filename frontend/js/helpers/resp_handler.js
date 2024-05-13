//@ts-check

import { actStatus } from "./status_act.js"
/**
    * @param {HTMLElement} status_holder
    *@param {Object} resp
    */
export async function handleResp(status_holder, resp) {

    let err = resp.code != 200
    switch (resp.type) {
        case "status":
            status_holder.innerHTML = resp.msg
            actStatus(status_holder, err, resp.msg)
            break
        case "msg":
            console.log(resp.msg)

            break
        case "redirect":
            console.log(resp.msg)
            let red_obj = await JSON.parse(resp.msg)
            actStatus(status_holder, err, red_obj.msg)
            location.href = red_obj.url
            break
    }

}
