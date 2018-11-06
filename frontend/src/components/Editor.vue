<template>
    <div  class="pure-u-1">
        <Modal v-if="showModal" @closeModal="handleCloseModal"/>
        <div class="pure-u-1">
            <div class="pure-u-1-2">
                <div class="pure-u-1-2">
                    <h2>Write your masterpiece</h2>
                    <label for="title">Title of your story:</label>
                    <input id="title" v-model="title" />
                    <button @click="showModal = true">New Story</button>
                    <button id="save" v-on:click="savePost">Save</button>
                    <button id="showPrev" v-on:click="togglePreview">Toggle Preview</button>
                </div>
                <div class="pure-u-1-2">
                    <h2>Details</h2>
                </div>
            </div>
        </div>

        <div class="pure-u-1">
            <div class="pure-u-1-2">
                <textarea class="pure-u-1" v-model="text"></textarea>
            </div>
            <div class="pure-u-1-2">
                <Preview v-show="showPreview" v-bind:content="text"></Preview>
            </div>
        </div>
    </div>
</template>

<style>

</style>


<script>
import Preview from './Preview.vue';
import Modal from './Modal.vue';
import { post } from 'axios';

export default {
    components: {
        Preview,
        Modal
    },
    methods: {
        handleCloseModal(arg) {
            this.showModal = false;
            console.log(`received ${arg}`)
        },         
        savePost () {
            post('http://localhost:1323/posts/story', {
                authorid: 'joe',
                title: this.title,
                body: this.text,
                tags: [],
                status: 'draft',
                category: 'short story',
                penname: 'Joe Deathmaster'
            }).then(res => {
                console.log(res);
            }).catch(err => console.err);
        },
        togglePreview() {
            this.showPreview = !this.showPreview;
        }
    },
    data () {
        return {
            title: '',
            text: '',
            showPreview: true,
            showModal: false
        }
    }
};
</script>