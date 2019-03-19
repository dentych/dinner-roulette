<template>
    <div class="container">
        <div class="row">
            <div class="col">
                <h2>Recipes</h2>
                <p>Here is a list of all recipes.</p>
            </div>
            <div class="col-auto">
                <router-link to="/add-recipe">
                    <button class="btn btn-success">Add new recipe</button>
                </router-link>
            </div>
        </div>
        <div class="row justify-content-center" v-if="recipes == null">
            <div class="spinner-border text-primary" role="status">
                <span class="sr-only">Loading...</span>
            </div>
        </div>
        <div class="row mt-4" v-if="!recipes || recipes.length === 0">
            <div class="col text-center">
                No recipes. Add your first recipe now!
            </div>
        </div>
        <div class="row mb-4" v-for="i in rowCount" :key="i">
            <div class="col-12 mb-4 mb-sm-0 col-sm-3" v-for="(recipe, index) in itemsInRow(i)" :key="recipe.uid">
                <div class="card h-100">
                    <div class="card-body">
                        <h5 class="card-title">{{recipe.name}}</h5>
                        <p class="card-text">{{shortDesc(recipe.description)}}</p>
                    </div>
                    <div class="card-footer">
                        <div class="d-flex justify-content-between align-items-center">
                            <router-link :to="{name: 'show-recipe', params: { id: calculateId(i, index) }}">
                                <button class="btn btn-sm btn-success">More info</button>
                            </router-link>
                            <a class="badge badge-light remove-icon align-middle" @click="deleteRecipe(calculateId(i, index))">
                                <i class="fas fa-times"></i>
                            </a>
                        </div>
                    </div>
                </div>
            </div>
        </div>
    </div>
</template>

<style scoped>
    .remove-icon {
        font-size: 15px;
        float: right;
    }

    .card-text {
        white-space: pre-wrap;
    }
</style>

<script>
    import {backendService} from "../services/BackendService"

    export default {
        data: function () {
            return {
                loaded: false,
                itemsPerRow: 4,
                recipes: null
            }
        },
        computed: {
            rowCount() {
                if (this.recipes == null) return 0;
                return Math.ceil(this.recipes.length / this.itemsPerRow)
            }
        },
        methods: {
            itemsInRow(index) {
                return this.recipes.slice((index - 1) * this.itemsPerRow, index * this.itemsPerRow)
            },
            deleteRecipe(index) {
                let confirmed = confirm("Delete recipe '" + this.recipes[index].name + "'?");
                if (confirmed) {
                    backendService.deleteRecipe(index)
                }
            },
            shortDesc(description) {
                if (description.length < 100) {
                    return description
                }
                return description.substr(0, description.lastIndexOf(" ", 100)) + "..."
            },
            calculateId(row, index) {
                return (row - 1) * 4 + index;
            }
        },
        mounted() {
            this.recipes = backendService.getAllRecipes()
        }
    }
</script>