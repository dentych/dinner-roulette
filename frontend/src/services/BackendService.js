import axios from "axios";
import {authService} from "./AuthService";

class BackendService {
    baseUrl = process.env.VUE_APP_BACKEND_BASE_URL;
    instance = axios.create({
        baseURL: this.baseUrl
    });

    getAllRecipes() {
        return this.sendRequestWithRetry("get", "/api/recipes")
            .then(res => Promise.resolve(res.data))
    }

    getRecipe(id) {
        return this.sendRequestWithRetry("get", "/api/recipes/" + id)
            .then(res => Promise.resolve(res.data))
    }

    saveRecipe(recipe) {
        return this.sendRequestWithRetry("post", "/api/recipes", recipe)
            .then(res => Promise.resolve(res.data))
    }

    updateRecipe(id, recipe) {
        return this.sendRequestWithRetry("put", "/api/recipes/" + id, recipe)
            .catch(err => Promise.reject(err.response ? err.response.data : err))
    }

    deleteRecipe(id) {
        return this.sendRequestWithRetry("delete", "/api/recipes/" + id)
            .catch(err => Promise.reject(err.response ? err.response.data : err))
    }

    registerUser(user) {
        return this.instance.post("/api/register", user)
    }

    sendRequestWithRetry(method, url, data, headers) {
        return this.sendRequest(method, url, data, headers).catch(error => {
            if (error.response && error.response.status === 401) {
                return authService.getToken().then(() => {
                    return this.sendRequest(method, url, data, headers)
                })
            } else {
                return Promise.reject(error);
            }
        })
    }

    sendRequest(method, url, data, headers) {
        return this.instance({
            method: method,
            url: url,
            headers: headerWithAuth(headers),
            data: data
        })
    }
}

function headerWithAuth(headers) {
    return Object.assign({Authorization: "Bearer " + authService.token}, headers)
}

const backendService = new BackendService();

export {backendService}