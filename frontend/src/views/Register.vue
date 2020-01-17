<template>
  <div class="container">
    <div class="row justify-content-center">
      <div class="col-12">
        <form @submit.prevent="register">
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
            <small
              id="emailHelp"
              class="form-text text-muted"
            >We'll never share your email with anyone
              else.
            </small>
          </div>
          <div class="form-group">
            <label for="firstName">First name</label>
            <input
              id="firstName"
              v-model="firstName"
              type="text"
              class="form-control"
              placeholder="First name"
            >
          </div>
          <div class="form-group">
            <label for="lastName">Last name</label>
            <input
              id="lastName"
              v-model="lastName"
              type="text"
              class="form-control"
              placeholder="First name"
            >
          </div>
          <div class="form-group">
            <label for="pass1">Password</label>
            <input
              id="pass1"
              v-model="pass1"
              type="password"
              class="form-control"
              placeholder="Password"
            >
          </div>
          <div class="form-group">
            <label for="pass2">Retype password</label>
            <input
              id="pass2"
              v-model="pass2"
              type="password"
              class="form-control"
              placeholder="Password"
            >
          </div>
          <div
            v-if="error"
            class="alert alert-danger"
          >
            {{ error }}
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
    import {backendService} from "../services/BackendService";
    import {authService} from "../services/AuthService";

    export default {
        name: "Register",
        data() {
            return {
                email: null,
                pass1: null,
                pass2: null,
                firstName: null,
                lastName: null,
                error: null
            }
        },
        methods: {
            register() {
                if (this.pass1 === this.pass2) {
                    this.error = null;

                    let user = {
                        email: this.email,
                        password: this.pass1,
                        firstName: this.firstName,
                        lastName: this.lastName
                    };

                    backendService.registerUser(user).then(() => this.$router.push("/")).then(() =>
                      authService.login(this.email, this.pass1).then(() => {
                          return authService.getToken()
                      }).then(() => {
                          this.$router.push({name: "home"})
                      }).catch(error => {
                          if (error.response && error.response.status === 400) {
                              this.errorMsg = "User failed to be created."
                          } else {
                              this.errorMsg = "Error logging in: " + JSON.stringify(error)
                          }
                      }))
                } else {
                    this.error = "Passwords do not match!"
                }
            }
        }
    }
</script>

<style scoped>

</style>
