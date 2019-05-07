<template>
    <div class="container">
        <div class="row justify-content-sm-center">
            <div class="col col-sm-8 mb-2">
                <h2>Add a new recipe</h2>
                <p>Fill out the below form to add a new recipe.</p>
                <form @submit.prevent="saveRecipe">
                    <div class="form-group">
                        <label for="name">Recipe name</label>
                        <input type="text" v-model="name" class="form-control" id="name"
                               placeholder="Enter a name for your recipe" autocomplete="off" required>
                    </div>
                    <div class="form-group">
                        <label for="url">Recipe link</label>
                        <input type="url" v-model="url" class="form-control" id="url"
                               placeholder="http://recipeplace.com/somerecipe">
                    </div>
                    <div class="form-group">
                        <label for="description">Description</label>
                        <textarea id="description" class="form-control" rows="10" v-model="description"></textarea>
                    </div>
                    <div class="form-group">
                        <label for="directions">Directions</label>
                        <textarea id="directions" class="form-control" rows="10" v-model="directions"></textarea>
                    </div>
                    <button type="submit" class="btn btn-success float-right">Submit</button>
                </form>
            </div>
        </div>
    </div>
</template>

<script>
    import {backendService} from "../services/BackendService"

    export default {
        name: "AddRecipe",
        data: function () {
            return {
                name: null,
                url: null,
                description: null,
                directions: null
            }
        },
        methods: {
            saveRecipe() {
                backendService.saveRecipe({
                    name: this.name,
                    url: this.url,
                    description: this.description,
                    directions: this.directions
                }).then(() => {
                    this.$router.push("/recipes")
                });
            },
            updateDescription(data) {
                this.description = data
            },
            updateDirections(data) {
                this.directions = data
            }
        }
    }
</script>
