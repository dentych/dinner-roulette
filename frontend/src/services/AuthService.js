import axios from "axios";

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
            return Promise.reject(error)
        })
    }

    getToken() {
        return axios.post(this.baseUrl + "/api/token", null, {withCredentials: true})
            .then(res => {
                this.token = res.data.access_token;
                return Promise.resolve()
            }).catch(err => {
                return Promise.reject(err)
            })
    }

    logout() {
        axios.post(this.baseUrl + "/api/logout", null, {withCredentials: true});
        this.token = null;
        localStorage.removeItem("authenticated")
    }
}

function sendTokenRequest(baseUrl) {
    return axios.post(baseUrl + "/api/token", null, {withCredentials: true})
        .then(response => {
            return {success: true, data: response.data}
        }, () => {
            return {success: false, data: null}
        })
}

const authService = new AuthService();

export {authService}