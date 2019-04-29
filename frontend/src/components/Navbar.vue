<template>
    <nav class="navbar fixed-top navbar-expand-lg navbar-dark bg-success">
        <a class="navbar-brand" href="#">Dinner Dash</a>
        <button class="navbar-toggler" type="button" data-toggle="collapse" data-target="#navbarSupportedContent"
                aria-controls="navbarSupportedContent" aria-expanded="false" aria-label="Toggle navigation">
            <span class="navbar-toggler-icon"></span>
        </button>

        <div class="collapse navbar-collapse" id="navbarSupportedContent">
            <ul class="navbar-nav mr-auto">
                <router-link to="/" tag="li" class="nav-item" exact>
                    <a class="nav-link">Home</a>
                </router-link>
                <router-link to="/recipes" tag="li" class="nav-item">
                    <a class="nav-link">Recipes</a>
                </router-link>
                <router-link to="/meal-plans" tag="li" class="nav-item">
                    <a class="nav-link">Meal plans</a>
                </router-link>
            </ul>
            <button class="btn btn-outline-light float-right" @click="logOut" v-if="authService.token">Log out</button>
            <div class="float-right" v-if="!authService.token">
                <router-link to="/register">
                    <button class="btn btn-outline-light">Register</button>
                </router-link>
                &nbsp;
                <router-link to="/login">
                    <button class="btn btn-outline-light">Log in</button>
                </router-link>
            </div>
        </div>
    </nav>
</template>

<script>
    import {authService} from "../services/AuthService";

    export default {
        name: "navbar",
        data() {
            return {
                authService: authService
            }
        },
        methods: {
            logOut() {
                authService.logout();
                this.$router.push({name: "home"})
            }
        },
        watch: {
            '$route'() {
                $(".navbar-collapse").collapse("hide")
            }
        }
    }
</script>