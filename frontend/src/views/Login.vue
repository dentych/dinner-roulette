<template>
    <div class="container">
        <div class="row justify-content-center">
            <div class="col-12 col-sm-4">
                <form @submit.prevent="login">
                    <div class="form-group">
                        <label for="email">Email address</label>
                        <input type="email" class="form-control" id="email" aria-describedby="emailHelp"
                               placeholder="Enter email" v-model="email">
                    </div>
                    <div class="form-group">
                        <label for="pwd">Password</label>
                        <input type="password" class="form-control" id="pwd" placeholder="Password" v-model="password">
                    </div>
                    <div class="alert alert-danger" v-if="errorMsg">{{errorMsg}}</div>
                    <button type="submit" class="btn btn-success float-right">Submit</button>
                </form>
            </div>
        </div>
    </div>
</template>

<script>
    import {authService} from "../services/AuthService";

    export default {
        name: "Login",
        data() {
            return {
                email: null,
                password: null,
                errorMsg: null
            }
        },
        methods: {
            login() {
                this.errorMsg = null;
                authService.login(this.email, this.password).then(() => {
                    return authService.getToken()
                }).then(() => {
                    this.$router.push({name: "home"})
                }).catch(error => {
                    if (error.response && error.response.status === 400) {
                        this.errorMsg = "Username or password incorrect."
                    } else {
                        this.errorMsg = "Error logging in: " + JSON.stringify(error)
                    }
                })
            }
        }
    }
</script>

<style scoped>

</style>