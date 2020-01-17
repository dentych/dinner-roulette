<template>
  <div class="container">
    <div class="row">
      <div class="col">
        <h2>Recipes</h2>
        <p>Here is a list of all recipes.</p>
      </div>
      <div class="col-auto">
        <router-link to="/add-recipe">
          <button class="btn btn-success">
            Add new recipe
          </button>
        </router-link>
      </div>
    </div>
    <div
      v-if="recipes == null"
      class="row justify-content-center"
    >
      <div
        class="spinner-border text-primary"
        role="status"
      >
        <span class="sr-only">Loading...</span>
      </div>
    </div>
    <div
      v-if="!recipes || recipes.length === 0"
      class="row mt-4"
    >
      <div class="col text-center">
        {{ message }}
      </div>
    </div>
    <div
      v-for="i in rowCount"
      :key="i"
      class="row mb-4"
    >
      <div
        v-for="recipe in itemsInRow(i)"
        :key="recipe.uid"
        class="col-12 mb-4 mb-sm-0 col-sm-3"
      >
        <div class="card h-100">
          <div class="card-body">
            <h5 class="card-title">
              {{ recipe.name }}
            </h5>
            <p class="card-text">
              {{ shortDesc(recipe.description) }}
            </p>
          </div>
          <div class="card-footer">
            <div class="d-flex justify-content-between align-items-center">
              <router-link :to="{name: 'show-recipe', params: { id: recipe.id }}">
                <button class="btn btn-sm btn-success">
                  More info
                </button>
              </router-link>
              <a
                class="badge badge-light remove-icon align-middle"
                @click="deleteRecipe(recipe)"
              >
                <i class="fas fa-times" />
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
                recipes: null,
                message: "No recipes. Add your first recipe now!"
            }
        },
        computed: {
            rowCount() {
                if (this.recipes == null) return 0;
                return Math.ceil(this.recipes.length / this.itemsPerRow)
            }
        },
        mounted() {
            backendService.getAllRecipes().then(data => {
                data.forEach(e => {
                    if (!e.description) {
                        e.description = ""
                    }
                });
                this.recipes = data
            }).catch(err => {
                this.recipes = [];
                this.message = "Error retrieving recipes... (" + err + ")"
            })
        },
        methods: {
            itemsInRow(index) {
                return this.recipes.slice((index - 1) * this.itemsPerRow, index * this.itemsPerRow)
            },
            deleteRecipe(recipe) {
                let confirmed = confirm("Delete recipe '" + name + "'?");
                if (confirmed) {
                    backendService.deleteRecipe(recipe.id).then(() => {
                        let pos = this.recipes.indexOf(recipe)
                        if (pos > -1) {
                            this.recipes.splice(pos, 1)
                        }
                    })
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
        }
    }
</script>