<template>
    <form @submit.prevent="register">
        <div class="form-group">
            <label for="email">Email address</label>
            <input type="email" class="form-control" id="email" aria-describedby="emailHelp" placeholder="Enter email"
                   v-model="email">
            <small id="emailHelp" class="form-text text-muted">We'll never share your email with anyone else.</small>
        </div>
        <div class="form-group">
            <label for="pass1">Password</label>
            <input type="password" class="form-control" id="pass1" placeholder="Password" v-model="pass1">
        </div>
        <div class="form-group">
            <label for="pass2">Retype password</label>
            <input type="password" class="form-control" id="pass2" placeholder="Password" v-model="pass2">
        </div>
        <div class="alert alert-danger" v-if="error">{{error}}</div>
        <button type="submit" class="btn btn-primary">Submit</button>
    </form>
</template>

<script>
    import {backendService} from "../services/BackendService";

    export default {
        name: "Register",
        data() {
            return {
                email: null,
                pass1: null,
                pass2: null,
                error: null
            }
        },
        methods: {
            register() {
                if (this.pass1 === this.pass2) {
                    this.error = null;

                    backendService.registerUser(this.email, this.pass1).then(() => this.$router.push("/"))
                } else {
                    this.error = "Passwords do not match!"
                }
            }
        }
    }
</script>

<style scoped>

</style>