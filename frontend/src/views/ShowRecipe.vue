<template>
  <div class="container">
    <div class="row">
      <div class="col col-sm-8">
        <div v-if="recipe">
          <h1>{{ recipe.name }}</h1>
          <strong>Description</strong><br>
          <p class="recipe-text">
            {{ recipe.description }}
          </p>
          <div v-if="recipe.directions">
            <hr>
            <strong>Directions</strong><br>
            <p class="recipe-text">
              {{ recipe.directions }}
            </p>
          </div>
          <div v-if="recipe.url">
            <hr>
            <p v-if="recipe.url">
              <strong>Recipe link: </strong><a
                :href="recipe.url"
                target="_blank"
              >{{ recipe.url }}</a>
            </p>
          </div>
        </div>
      </div>
      <div class="col col-sm-4">
        <button
          class="btn btn-success"
          @click="editRecipe()"
        >
          Edit
        </button>
        <br><br>
        <button
          class="btn btn-danger"
          @click="deleteRecipe()"
        >
          Delete
        </button>
      </div>
    </div>
  </div>
</template>

<script>
    import {backendService} from "../services/BackendService"

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
            backendService.getRecipe(this.id).then(recipe => this.recipe = recipe)
        },
        methods: {
            editRecipe() {
                this.$router.push("/edit-recipe/" + this.id)
            },
            deleteRecipe() {
                let confirmed = confirm("Are you sure you want to delete this recipe?");
                if (confirmed) {
                    backendService.deleteRecipe(this.id).then(() => this.$router.push("/recipes"));
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
