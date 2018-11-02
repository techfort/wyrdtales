<template>
  <div id="app">
    <div class="navbar" id="navbar">
      <div class="navbarlink"><router-link to="/home">Home</router-link></div>
    </div>
    <router-view></router-view>
  </div>
</template>



<style>
@import url('https://fonts.googleapis.com/css?family=EB+Garamond');
html {
  font-family: 'EB Garamond', 'Georgia', serif;
  font-size: 1.2em;
  -webkit-font-smoothing: antialiased;
  -moz-osx-font-smoothing: grayscale;
  text-align: center;
  margin-top: 60px;
  background: #fafafa;
  color: #333;
  min-height: 100%;
}

#app {
  width: 100%;
  text-align: center;
}

#navbar {
  width: 100%;
  border-bottom: 1px solid #333;
  height: 60px;
  position: fixed;
  top: 0px;
  text-align: left;
  display: table;
}

.navbarlink {
  float: left;
  height: 60px;
  margin: 0;
  padding: 0;
  display: table-cell;
  vertical-align: middle;  
}

#main {
  width: 74%;
  text-align: justify;
  float: left;
}

h1, h2, a {
  text-decoration: none;
  color: #333;
  font-variant: small-caps;
}

a:visited, a:target {
  color: #333;
}
</style>
