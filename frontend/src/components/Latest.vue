<template>
    <div class="container latest">
        <h2>Latest Stories</h2>
        <hr/>
        <div v-if="posts.length > 0">
        <Post v-for="post in posts" v-bind:key="post.id"
            v-bind:post="post"
        >
        </Post>
        </div>
        <div v-else>{{ errorMessage }}</div>
    </div>
</template>

<script>
import postService from '../services/post'
import Post from './Post.vue'

const Latest = () => {
    return {
        name: 'Latest',
        components: {
            Post
        },
        data() {
            return {
                posts: [],
                errorMessage: ''
            } 
        },
        mounted() {
            postService.getLatest().then(({ data }) => { 
                this.posts = data
            }).catch(err => {
                this.posts = [];
                this.errorMessage = 'No stories available.';
            })
        }
    }
};

export default Latest();
</script>

<style scoped>
.container {
    border-left: 1px solid #333;
}
.latest {
    position: relative top;
    float: right;
    width: 25%;
}
</style>
