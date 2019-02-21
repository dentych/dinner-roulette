
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
}

export default new BackendService()