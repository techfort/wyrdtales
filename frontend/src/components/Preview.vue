<template>
    <div id="content_preview" class="pure-u-1" v-html="markdownText">
    </div>
</template>

<script>
import marked from 'marked';

export default  {
    props: ['content'],
    computed: {
        markdownText() {
            return marked(this.content, { sanitize: true })
        }
    }
};
</script>
