<template>
    <div class="form-group">
        <label for="description">Description</label>
        <textarea v-model="description" @keyup="onKeyUp" ref="description" contenteditable="true" rows="15" id="description" class="form-control description-area"></textarea>
    </div>
</template>

<script>
    export default {
        name: "growing-text-area",
        props: ["value"],
        data () {
            return {
                description: this.value
            }
        },
        methods: {
            onKeyUp() {
                this.$emit("update", this.description)
            }
        }
    }
</script>

<style scoped>
    .description-area {
        overflow: auto;
    }
</style>