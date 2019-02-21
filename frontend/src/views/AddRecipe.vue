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
                        <label for="url">Recipe URL</label>
                        <input type="url" v-model="url" class="form-control" id="url" placeholder="http://recipeplace.com/somerecipe">
                    </div>
                    <div class="form-group">
                        <label for="description">Description</label>
                        <textarea v-model="description" @keyup="autoGrow" contenteditable="true" rows="3" id="description" class="form-control description-area"></textarea>
                    </div>
                    <button type="submit" class="btn btn-success float-right">Submit</button>
                </form>
            </div>
        </div>
    </div>
</template>

<script>
    import backendService from "../services/BackendService"

    export default {
        name: "AddRecipe",
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
            autoGrow(e) {
                let textField = e.target;
                textField.style.height = "auto";
                if (textField.scrollHeight > textField.clientHeight) {
                    textField.style.height = textField.scrollHeight + "px"
                }
            }
        }
    }
</script>

<style scoped>
    .description-area {
        overflow: hidden
    }
</style>