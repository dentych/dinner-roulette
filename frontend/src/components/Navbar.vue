<template>
  <nav class="navbar fixed-top navbar-expand-lg navbar-dark bg-success">
    <a
      class="navbar-brand"
      href="#"
    >Dinner Dash</a>
    <button
      class="navbar-toggler"
      type="button"
      data-toggle="collapse"
      data-target="#navbarSupportedContent"
      aria-controls="navbarSupportedContent"
      aria-expanded="false"
      aria-label="Toggle navigation"
    >
      <span class="navbar-toggler-icon" />
    </button>

    <div
      id="navbarSupportedContent"
      class="collapse navbar-collapse"
    >
      <ul class="navbar-nav mr-auto">
        <router-link
          to="/"
          tag="li"
          class="nav-item"
          exact
        >
          <a class="nav-link">Home</a>
        </router-link>
        <router-link
          to="/recipes"
          tag="li"
          class="nav-item"
        >
          <a class="nav-link">Recipes</a>
        </router-link>
        <router-link
          to="/meal-plans"
          tag="li"
          class="nav-item"
        >
          <a class="nav-link">Meal plans</a>
        </router-link>
      </ul>
      <button
        v-if="authService.token"
        class="btn btn-outline-light float-right"
        @click="logOut"
      >
        Log out
      </button>
      <div
        v-if="!authService.token"
        class="float-right"
      >
        <router-link to="/register">
          <button class="btn btn-outline-light">
            Register
          </button>
        </router-link>
                &nbsp;
        <router-link to="/login">
          <button class="btn btn-outline-light">
            Log in
          </button>
        </router-link>
      </div>
    </div>
  </nav>
</template>

<script>
    import {authService} from "../services/AuthService";

    export default {
        name: "Navbar",
        data() {
            return {
                authService: authService
            }
        },
        watch: {
            '$route'() {
                $(".navbar-collapse").collapse("hide")
            }
        },
        methods: {
            logOut() {
                authService.logout();
                this.$router.push({name: "home"})
            }
        }
    }
</script>