<template>
    <div class="container">
        <div class="row">
            <div class="col">
                <h2>Recipes</h2>
                <p>Here is a list of all recipes.</p>
            </div>
            <div class="col-auto">
                <button class="btn btn-success">Add new recipe</button>
            </div>
        </div>
        <div class="row justify-content-center" v-if="!loaded">
            <div class="spinner-border text-primary" role="status">
                <span class="sr-only">Loading...</span>
            </div>
        </div>
        <div class="row mb-4" v-for="i in rowCount" :key="i">
            <div class="card-deck">
                <div class="card mx-2" v-for="recipe in itemsInRow(i)" :key="recipe.title">
                    <div class="card-body">
                        <h5 class="card-title">{{recipe.title}}</h5>
                        <p class="card-text">{{recipe.body}}</p>
                    </div>
                    <div class="card-footer">
                        <div class="d-flex justify-content-between align-items-center">
                            <button class="btn btn-sm btn-success">More info</button>
                            <a class="badge badge-light remove-icon align-middle" href="#">
                                <i class="fas fa-times"></i>
                            </a>
                        </div>
                    </div>
                </div>
            </div>
        </div>
    </div>
</template>

<style>
    .remove-icon {
        font-size: 15px;
        float: right;
    }
</style>

<script>
    import axios from "axios";

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
            itemsInRow: function (index) {
                return this.recipes.slice((index - 1) * this.itemsPerRow, index * this.itemsPerRow)
            }
        },
        created: function () {
            // setTimeout(() => {
            axios
                .get("https://jsonplaceholder.typicode.com/posts")
                .then(response => {
                    this.loaded = true;
                    if (response.status >= 200 && response.status < 300) {
                        this.recipes = response.data;
                    }
                })
            // }, 1000)

        }
    }
</script>