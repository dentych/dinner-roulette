<template>
    <div class="container">
        <div class="row">
            <div class="card col">
                <div class="card-body">
                    <h4 class="card-title">Welcome to Dinner Dash</h4>
                    <p class="card-text">Ever tired of having the struggle of figuring out what to make for dinner?
                        Maybe you buy groceries almost every day of the week? With Dinner Dash, it is possible to save
                        recipes that you like to make, and make Dinner Dash generate a dinner plan for you, that will
                        last for a week.
                        <br>
                        In the future it will be possible to also get a shopping list to use for
                        buying groceries which will fit the week's meal plan.</p>
                    <hr>
                    <h4 class="card-title">Feature implementation status</h4>
                    <p class="card-text">
                        For an always up-to-date overview of features/bugs/etc. please look at the Github page here:
                        <a href="https://github.com/dentych/dinner-dash/issues" target="_blank">https://github.com/dentych/dinner-dash/issues</a>
                    </p>
                    <ul class="fa-ul">
                        <li><span class="fa-li"><i class="fas fa-check-circle text-success"></i></span> Log in</li>
                        <li><span class="fa-li"><i class="fas fa-check-circle text-success"></i></span>
                            Register user
                        </li>
                        <li><span class="fa-li"><i class="fas fa-check-circle text-success"></i></span> Log out</li>
                        <li><span class="fa-li"><i class="fas fa-check-circle text-success"></i></span> Add recipe</li>
                        <li><span class="fa-li"><i class="fas fa-check-circle text-success"></i></span>
                            List all recipes
                        </li>
                        <li><span class="fa-li"><i class="fas fa-check-circle text-success"></i></span> Show single
                            recipe
                        </li>
                        <li><span class="fa-li"><i class="fas fa-check-circle text-success"></i></span>
                            Edit recipe
                        </li>
                        <li><span class="fa-li"><i class="fas fa-check-circle text-success"></i></span>
                            Delete recipe
                        </li>
                        <li><span class="fa-li"><i class="far fa-times-circle text-danger"></i></span>
                            Add ingredient
                        </li>
                        <li><span class="fa-li"><i class="far fa-times-circle text-danger"></i></span>
                            Edit ingredient
                        </li>
                        <li><span class="fa-li"><i class="far fa-times-circle text-danger"></i></span>
                            Delete ingredient
                        </li>
                        <li><span class="fa-li"><i class="far fa-times-circle text-danger"></i></span>
                            Generate meal plan
                        </li>
                    </ul>
                </div>
            </div>
        </div>
    </div>
</template>

<script>

    export default {
        name: 'home'
    }
</script>
