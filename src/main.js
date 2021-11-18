import Vue from 'vue'
import App from './App.vue'
import router from './router'
import ElementUI from 'element-ui'
import ViewUI from 'view-design'
import 'view-design/dist/styles/iview.css'
import 'element-ui/lib/theme-chalk/index.css'
import { BootstrapVue } from 'bootstrap-vue'
import 'bootstrap/dist/css/bootstrap.css'
import 'bootstrap-vue/dist/bootstrap-vue.css'
Vue.use(ElementUI)
Vue.use(BootstrapVue)
Vue.use(ViewUI)
Vue.config.productionTip = false

// Server's url for REST endpoints
Vue.prototype.$url = 'http://106.52.119.98:8765/api'

new Vue({
    router,
    render: (h) => h(App),
}).$mount('#app')
