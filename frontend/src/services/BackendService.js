import axios from "axios";
import {authService} from "./AuthService";

var recipes = [];

function updateLocalStorage() {
    localStorage.setItem("recipes", JSON.stringify(recipes))
}

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
        return recipes[id];
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
        recipes[id] = recipe;
        updateLocalStorage()
    }

    deleteRecipe(id) {
        recipes.splice(id, 1);
        updateLocalStorage()
    }

    registerUser(user) {
        return axios.post(this.baseUrl + "/api/register", user)
    }
}

const backendService = new BackendService();

export {backendService}