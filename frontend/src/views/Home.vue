<template>
    <div class="container">
        <div class="row">
            <div class="card col-12">
                <div class="card-body">
                    <p class="card-text"><i class="fas fa-check-circle" style="color: green"></i> Hello world!</p>
                    <p class="card-text" style="font-size: 50px;">I like <i class="fas fa-cat"></i> very much. They do not like <i class="fas fa-dog"></i></p>
                </div>
            </div>
        </div>
    </div>
</template>

<script>

    export default {
        name: 'home',
        mounted() {
            document.title = "Home"
        }
    }
</script>
