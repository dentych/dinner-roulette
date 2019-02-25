<template>
    <div class="container">
        <div class="row">
            <div class="col col-sm-8">
                <div v-if="recipe">
                    <h1>{{recipe.name}}</h1>
                    <p class="recipe-text">{{recipe.description}}</p>
                    <hr>
                    <p>Recipe link: <a :href="recipe.url" target="_blank">{{recipe.url}}</a></p>
                </div>
            </div>
            <div class="col col-sm-4">
                <button class="btn btn-success" @click="editRecipe()">Edit</button>
                <br><br>
                <button class="btn btn-danger" @click="deleteRecipe()">Delete</button>
            </div>
        </div>
    </div>
</template>

<script>
    import backendService from "../services/BackendService"

    export default {
        name: "ShowRecipe",
        data: function () {
            return {
                id: null,
                recipe: null
            }
        },
        mounted() {
            this.id = this.$route.params.id;
            this.recipe = backendService.getRecipe(this.id)
        },
        methods: {
            editRecipe() {
                this.$router.push("/edit-recipe/" + this.id)
            },
            deleteRecipe() {
                let confirmed = confirm("Are you sure you want to delete the recipe with ID: " + this.id + "?")
                if (confirmed) {
                    backendService.deleteRecipe(this.id);
                    this.$router.push("/recipes");
                }
            }
        }
    }
</script>

<style scoped>
    .recipe-text {
        white-space: pre-wrap;
    }
</style>