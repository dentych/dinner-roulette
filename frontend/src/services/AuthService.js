import axios from "axios";

class AuthService {
    token = null;
    refreshRate = 4.5 * 60 * 1000;

    constructor() {
        this.getToken();
        setInterval(this.getToken, this.refreshRate)
    }

    isLoggedIn() {
        return this.token || localStorage.getItem("authenticated")
    }

    login(email, password) {
        return axios.post("/api/login", {
            email: email,
            password: password
        }, {withCredentials: true})
            .then(() => {
                localStorage.setItem("authenticated", "true");
                return Promise.resolve()
            }, error => {
                localStorage.removeItem("authenticated");
                return Promise.reject(error)
            })
    }

    getToken() {
        return axios.post("/api/token", null, {withCredentials: true})
            .then(res => {
                this.token = res.data.access_token;
                return Promise.resolve()
            }).catch(err => {
                this.token = null;
                localStorage.removeItem("authenticated");
                return Promise.reject(err)
            })
    }

    logout() {
        axios.post("/api/logout", null, {withCredentials: true});
        this.token = null;
        localStorage.removeItem("authenticated")
    }
}

const authService = new AuthService();

export {authService}
