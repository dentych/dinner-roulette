<template>
  <div>
    <navbar />
    <router-view />
  </div>
</template>

<style>
    body {
        margin-top: 70px
    }
</style>

<script>
    import Navbar from "./components/Navbar";

    export default {
        components: {
            Navbar
        }
    }
</script>