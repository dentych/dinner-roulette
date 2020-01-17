<template>
  <div class="container">
    <div class="row justify-content-center">
      <div class="col-12 col-sm-4">
        <form @submit.prevent="login">
          <div class="form-group">
            <label for="email">Email address</label>
            <input
              id="email"
              v-model="email"
              type="email"
              class="form-control"
              aria-describedby="emailHelp"
              placeholder="Enter email"
            >
          </div>
          <div class="form-group">
            <label for="pwd">Password</label>
            <input
              id="pwd"
              v-model="password"
              type="password"
              class="form-control"
              placeholder="Password"
            >
          </div>
          <div
            v-if="errorMsg"
            class="alert alert-danger"
          >
            {{ errorMsg }}
          </div>
          <button
            type="submit"
            class="btn btn-success float-right"
          >
            Submit
          </button>
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
        beforeMount(){
          authService.getToken().then(() => {
                this.$router.push({name: "home"})
              })
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
        },

    }
</script>

<style scoped>

</style>
