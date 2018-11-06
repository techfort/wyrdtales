import Vue from 'vue';
import Router from 'vue-router';
import App from './App.vue';
import Home from './components/Home.vue';
import Write from './components/Editor.vue';
import Modal from './components/Modal.vue';
import formatDate from './filters/formatDate.js';

Vue.config.productionTip = false;
Vue.use(Router);
Vue.filter('formatDate', formatDate);


const router = new Router({
  routes: [
    { path: '/home', component: Home },
    { path: '/write', component: Write }
  ]
});

new Vue({
  render: h => h(App),
  router,
  components: {
    Home,
    Write,
    Modal
  }
}).$mount('#app');
