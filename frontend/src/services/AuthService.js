const axios = require("axios");

class AuthService {
    token = null;
    baseUrl = process.env.VUE_APP_BACKEND_BASE_URL;

    constructor() {
        this.getToken()
    }

    isLoggedIn() {
        return this.token || localStorage.getItem("authenticated")
    }

    login(email, password) {
        return axios.post(this.baseUrl + "/api/login", {
            email: email,
            password: password
        }, {withCredentials: true}).then(() => {
            localStorage.setItem("authenticated", "true");
            return Promise.resolve()
        }, error => {
            localStorage.removeItem("authenticated");
            if (error.response) {
                return Promise.reject(error.response.data.message)
            } else {
                return Promise.reject("something went wrong. Try again later.")
            }
        })
    }

    getToken() {
        return axios.post(this.baseUrl + "/api/token", null, {withCredentials: true})
            .then(response => {
                this.token = response.data.access_token;
                return {success: true, data: response.data}
            }, () => {
                return {success: false, data: null}
            })
    }

    logout() {
        axios.post(this.baseUrl + "/api/logout", null, {withCredentials: true});
        this.token = null;
        localStorage.removeItem("authenticated")
    }
}

const authService = new AuthService();

export {authService}