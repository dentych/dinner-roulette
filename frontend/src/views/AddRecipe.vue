<template>
    <div class="container">
        <div class="row justify-content-sm-center">
            <div class="col col-sm-8">
                <h2>Add a new recipe</h2>
                <p>Fill out the below form to add a new recipe.</p>
                <form @submit.prevent="saveRecipe">
                    <div class="form-group">
                        <label for="name">Recipe name</label>
                        <input type="text" v-model="name" class="form-control" id="name" placeholder="Enter a name for your recipe" autocomplete="off">
                    </div>
                    <div class="form-group">
                        <label for="url">Recipe link</label>
                        <input type="url" v-model="url" class="form-control" id="url" placeholder="http://recipeplace.com/somerecipe">
                    </div>
                    <growing-text-area @update="updateDescription"></growing-text-area>
                    <button type="submit" class="btn btn-success float-right">Submit</button>
                </form>
            </div>
        </div>
    </div>
</template>

<script>
    import {backendService} from "../services/BackendService"
    import GrowingTextArea from "../components/GrowingTextArea";

    export default {
        name: "AddRecipe",
        components: {GrowingTextArea},
        data: function () {
            return {
                name: null,
                url: null,
                description: null
            }
        },
        methods: {
            saveRecipe() {
                backendService.saveRecipe({name: this.name, url: this.url, description: this.description})
                this.$router.push("/recipes")
            },
            updateDescription(data) {
                this.description = data
            }
        }
    }
</script>
