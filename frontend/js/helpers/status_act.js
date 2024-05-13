//@ts-check

/**
    *@param {HTMLElement} status_holder 
    *@param {boolean} error
    *@param {string} status_msg
    */
export function actStatus(status_holder, error, status_msg) {
    status_holder.classList.add("visible")
    status_holder.innerHTML = status_msg
    if (error) {
        status_holder.classList.remove("ok")
        status_holder.classList.add("err")
    }
    else {
        status_holder.classList.remove("err")
        status_holder.classList.add("ok")
    }
}
