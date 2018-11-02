<template>
    <div class="wt-container">
        <h2>{{ post.title }}</h2>
        <p class="deets">{{ post.posted }}</p>
        <p>{{ post.body }}</p>
    </div>
</template>

<script>

export default {
    name: 'Post',
    props: ['post'],
    mounted() {
        console.log("post mounted correctly")
    }
};
</script>

<style scoped>
.deets {
    font-size: 0.6em;
}
</style>