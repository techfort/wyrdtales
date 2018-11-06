<template>
    <transition name="modal">
        <div class="modal-mask">
            <div class="modal-wrapper">
                <div class="modal-container">
                    <div class="modal-header">
                        <slot name="header">
                            header
                        </slot>
                    </div>
                    <div class="modal-body">
                        <slot name="body">
                            body
                        </slot>
                    </div>
                    <div class="modal-footer">
                        <slot name="footer">
                           <button class="modal-default-button" v-on:click="close">close</button> 
                        </slot>
                    </div>
                </div>
            </div>
        </div>
    </transition>
</template>

<script>
export default {
    methods: {
        close() {
            console.log("calling close method...")
            this.$emit('closeModal', 1)
        }
    }
}
</script>