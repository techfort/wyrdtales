import Vue from 'vue';
import Router from 'vue-router';
import App from './App.vue';
import Home from './components/Home.vue';

Vue.config.productionTip = false;
Vue.use(Router);

const router = new Router({
  routes: [
    { path: '/home', component: Home }
  ]
});

new Vue({
  render: h => h(App),
  router
}).$mount('#app');
