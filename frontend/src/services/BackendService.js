import axios from "axios";
import {authService} from "./AuthService";

class BackendService {
    baseUrl = process.env.VUE_APP_BACKEND_BASE_URL;

    getAllRecipes() {
        return axios.get(this.baseUrl + "/api/recipes", {headers: {Authorization: "Bearer " + authService.token}})
            .then(response => {
                return Promise.resolve(response.data)
            }, error => {
                if (error.response) {
                    return Promise.reject(error.response.status)
                } else {
                    return Promise.reject(error)
                }
            });
    }

    getRecipe(id) {
        return axios.get(this.baseUrl + "/api/recipes/" + id, {headers: {Authorization: "Bearer " + authService.token}})
            .then(response => {
                return Promise.resolve(response.data)
            }, err => {
                if (err.response) {
                    return Promise.reject(err)
                } else {
                    return Promise.reject(err)
                }
            })
    }

    saveRecipe(recipe) {
        return axios.post(this.baseUrl + "/api/recipes", recipe, {headers: {Authorization: "Bearer " + authService.token}})
            .then(response => {
                return Promise.resolve(response.data)
            }, error => {
                if (error.response) {
                    return Promise.reject(error.response.status)
                } else {
                    return Promise.reject(error)
                }
            })
    }

    updateRecipe(id, recipe) {
        return axios.put(this.baseUrl + "/api/recipes/" + id, recipe, {headers: {Authorization: "Bearer " + authService.token}})
            .catch(err => {
                return Promise.reject(err.response ? err.response.data : err);
            })
    }

    deleteRecipe(id) {
        return axios.delete(this.baseUrl + "/api/recipes/" + id, {headers: {Authorization: "Bearer " + authService.token}})
            .catch(err => {
                return Promise.reject(err.response ? err.response.data : err)
            })
    }

    registerUser(user) {
        return axios.post(this.baseUrl + "/api/register", user)
    }
}

const backendService = new BackendService();

export {backendService}