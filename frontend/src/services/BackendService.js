import axios from "axios";

var recipes = [];

function updateLocalStorage() {
    localStorage.setItem("recipes", JSON.stringify(recipes))
}

class BackendService {
    constructor() {
        let json = localStorage.getItem("recipes");
        recipes = json ? JSON.parse(json) : [];
    }

    getAllRecipes() {
        return recipes;
    }

    getRecipe(id) {
        return recipes[id];
    }

    saveRecipe(recipe) {
        recipe.uid = Math.random() * 1000;
        recipes.push(recipe);
        updateLocalStorage()
    }

    updateRecipe(id, recipe) {
        recipes[id] = recipe;
        updateLocalStorage()
    }

    deleteRecipe(id) {
        recipes.splice(id, 1);
        updateLocalStorage()
    }

    registerUser(email, pass1) {
        let user = {email: email, password: pass1};
        return axios.post("http://localhost:8081/api/register", user)
    }
}

export default new BackendService()