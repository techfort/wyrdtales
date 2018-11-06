<template>
  <div>
    <div class="header">
      <div class="home-menu pure-menu pure-menu-horizontal pure-menu-fixed">
        <a class="pure-menu-heading">WyrdTales</a>
        <ul class="pure-menu-list">
          <li class="pure-menu-item"><router-link to="/home">Home</router-link></li>
          <li class="pure-menu-item"><router-link to="/write">Write</router-link></li>
        </ul>
      </div>
    </div>
    <div class="pure-g" id="mainview">
      <router-view></router-view>
    </div>
  </div>
</template>



<style>
@import url('https://fonts.googleapis.com/css?family=EB+Garamond');
html, body,  .pure-g [class *= "pure-u"] {
  font-family: 'EB Garamond', 'Georgia', serif !important;
  color: #333;
}

* {
    -webkit-box-sizing: border-box;
    -moz-box-sizing: border-box;
    box-sizing: border-box;
}

h1, h2, h3, h4, h5, h6, a {
  font-variant: small-caps !important;
}
#app {
  height: 100%;
  text-align: center;
}

#mainview {
  padding: 2em 1em 0;
}

.pure-menu {
  border-bottom: 1px solid #333;
}

.pure-menu-item {
  vertical-align: middle;
  padding: 0 2em 0;
}

#content_preview {
  background: #eee;
}

.modal-mask {
  position: fixed;
  z-index: 9998;
  top: 0;
  left: 0;
  width: 100%;
  height: 100%;
  background-color: rgba(0, 0, 0, .5);
  display: table;
  transition: opacity .3s ease;
}

.modal-wrapper {
  display: table-cell;
  vertical-align: middle;
}

.modal-container {
  width: 300px;
  margin: 0px auto;
  padding: 20px 30px;
  background-color: #fff;
  border-radius: 2px;
  box-shadow: 0 2px 8px rgba(0, 0, 0, .33);
  transition: all .3s ease;
  font-family: Helvetica, Arial, sans-serif;
}

.modal-header h3 {
  margin-top: 0;
  color: #42b983;
}

.modal-body {
  margin: 20px 0;
}

.modal-default-button {
  float: right;
}

/*
 * The following styles are auto-applied to elements with
 * transition="modal" when their visibility is toggled
 * by Vue.js.
 *
 * You can easily play with the modal transition by editing
 * these styles.
 */

.modal-enter {
  opacity: 0;
}

.modal-leave-active {
  opacity: 0;
}

.modal-enter .modal-container,
.modal-leave-active .modal-container {
  -webkit-transform: scale(1.1);
  transform: scale(1.1);
}
</style>
