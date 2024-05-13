//@ts-check

export class User {
    /**
        *@param {string} username
        *@param {string} password
        */
    constructor(username, password) {
        this.username = username
        this.password = password
    }
    SayHi() {
        console.log(this.username + "\n" + this.password)
    }
}
