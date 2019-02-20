
var recipes = [
    {name: "dummy recipe 1", url: "http://google.com", description: "this is a dummy recipe this is a dummy recipe this is a dummy recipe this is a dummy recipe this is a dummy recipe "},
    {name: "dummy recipe 2", url: "http://google.com", description: "this is a dummy recipethis is a dummy recipe this is a dummy recipe this is a dummy recipe this is a dummy recipe "},
    {name: "dummy recipe 3", url: "http://google.com", description: "this is a dummy recipethis is a dummy recipe this is a dummy recipe this is a dummy recipe "},
    {name: "dummy recipe 4", url: "http://google.com", description: "this is a dummy recipethis is a dummy recipe this is a dummy recipe this is a dummy recipe this is a dummy recipe "},
    {name: "dummy recipe 5", url: "http://google.com", description: "this is a dummy recipethis is a dummy recipe this is a dummy recipe this is a dummy recipe this is a dummy recipe "},
    ];

class BackendService {
    getAllRecipes() {
        return recipes;
    }

    saveRecipe(recipe) {
        recipes.push(recipe)
    }
}

const backendService = new BackendService();

export default backendService