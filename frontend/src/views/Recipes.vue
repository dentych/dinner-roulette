<template>
    <div class="container">
        <div class="row">
            <div class="col-12">
                <h2>Recipes</h2>
                <p>Here is a list of all recipes.</p>
            </div>
        </div>
        <div class="row justify-content-center" v-if="!loaded">
            <div class="spinner-border text-primary" role="status">
                <span class="sr-only">Loading...</span>
            </div>
        </div>
        <div class="row mb-4" v-for="i in rowCount" :key="i">
            <div class="card-deck">
                <div class="card" v-for="recipe in itemsInRow(i)" :key="recipe.title">
                    <div class="card-body">
                        <h5 class="card-title">{{recipe.title}}</h5>
                        <p class="card-text">{{recipe.body}}</p>
                    </div>
                </div>
            </div>
        </div>
    </div>
</template>

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
        },
        mounted() {
            document.title = "Fuck the world"
        }
    }
</script>