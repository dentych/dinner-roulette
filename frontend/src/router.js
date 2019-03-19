import Vue from 'vue'
import Router from 'vue-router'
import Home from './views/Home.vue'
import AddRecipe from "./views/AddRecipe";
import ShowRecipe from "./views/ShowRecipe";
import EditRecipe from "./views/EditRecipe";
import Register from "./views/Register";
import Login from "./views/Login";

Vue.use(Router);

export default new Router({
    mode: 'history',
    linkActiveClass: "active",
    base: process.env.BASE_URL,
    routes: [
        {
            path: '/',
            name: 'home',
            component: Home
        },
        {
            path: '/recipes',
            name: 'recipes',
            // route level code-splitting
            // this generates a separate chunk (about.[hash].js) for this route
            // which is lazy-loaded when the route is visited.
            component: () => import(/* webpackChunkName: "about" */ './views/Recipes.vue')
        },
        {
            path: "/recipes/:id",
            name: "show-recipe",
            component: ShowRecipe
        },
        {
            path: "/add-recipe",
            name: "add recipe",
            component: AddRecipe
        },
        {
            path: "/edit-recipe/:id",
            name: "edit recipe",
            component: EditRecipe
        },
        {
            path: "/register",
            name: "register user",
            component: Register
        },
        {
            path: "/login",
            name: "login",
            component: Login
        }
    ]
})