<template>
    <div id="home">
        <h1>{{ header }}</h1>
        <p>{{ tagline }}</p>
        <hr />
        <div class="container" id="main">
        <h2>Wild Talent and unforgettable stories.</h2>
        Lorem ipsum dolor sit amet, consectetur adipiscing elit. Pellentesque commodo euismod massa, vel porttitor nisl ornare porttitor. Aenean volutpat eleifend ligula, sit amet lacinia nisl fermentum nec. Cras arcu diam, commodo in velit et, commodo molestie diam. Sed laoreet sodales velit, id auctor purus maximus in. Class aptent taciti sociosqu ad litora torquent per conubia nostra, per inceptos himenaeos. Curabitur dictum diam sem, id consectetur dui tincidunt non. Donec blandit est quis nulla sagittis gravida luctus quis felis. Mauris viverra dapibus mi, sed aliquam diam accumsan quis. Praesent sed lorem neque.
        </div>
        <Latest />
    </div>
</template>

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