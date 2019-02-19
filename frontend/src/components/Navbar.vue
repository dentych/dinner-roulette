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
            </ul>
        </div>
    </nav>
</template>

<script>
    export default {
        name: "navbar",
        watch: {
            '$route' () {
                $(".navbar-collapse").collapse("hide")
            }
        }
    }
</script>