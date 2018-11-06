<template>
    <div class="pure-g">
        <div class="pure-u-1 pure-u-md-1">
          <div class="header">
            <h1 class="brand-title">{{ header }}</h1>
            <h2 class="brand-subtitle">{{ tagline }}</h2>
          </div>
        </div>

        <div class="content pure-u-1 pure-md-1">
          <div class="pure-u-1-2">
            <h3>Wild Talent and unforgettable stories.</h3>
            Lorem ipsum dolor sit amet, consectetur adipiscing elit. Pellentesque commodo euismod massa, vel porttitor nisl ornare porttitor. Aenean volutpat eleifend ligula, sit amet lacinia nisl fermentum nec. Cras arcu diam, commodo in velit et, commodo molestie diam. Sed laoreet sodales velit, id auctor purus maximus in. Class aptent taciti sociosqu ad litora torquent per conubia nostra, per inceptos himenaeos. Curabitur dictum diam sem, id consectetur dui tincidunt non. Donec blandit est quis nulla sagittis gravida luctus quis felis. Mauris viverra dapibus mi, sed aliquam diam accumsan quis. Praesent sed lorem neque.
          </div>
          <div class="pure-u-1-2" >
            <Latest />
          </div>
        </div>
    </div>
</template>

<style>
</style>

<script>
import Latest from './Latest.vue'

export default {
  name: 'app',
  data() {
    return {
      header: "Wyrd Tales",
      tagline: "Fostering talent and enjoying good stories."
    };
  },
  components: {
    Latest
  }
}
</script>