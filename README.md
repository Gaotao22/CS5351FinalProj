# team9

> A Vue.js project

import ElementUI from 'element-ui'

import ViewUI from 'view-design'

import { BootstrapVue } from 'bootstrap-vue'

这三个组件要下载一下，npm终端下载就行

然后在main.js里面添加：
import ElementUI from 'element-ui'

import ViewUI from 'view-design'

import 'view-design/dist/styles/iview.css'

import 'element-ui/lib/theme-chalk/index.css'

import { BootstrapVue } from 'bootstrap-vue'

import 'bootstrap/dist/css/bootstrap.css'

import 'bootstrap-vue/dist/bootstrap-vue.css'

Vue.use(BootstrapVue)

Vue.use(ViewUI)

Vue.use(ElementUI)

## Build Setup

``` bash
# install dependencies
npm install

# serve with hot reload at localhost:8080
npm run dev

# build for production with minification
npm run build

# build for production and view the bundle analyzer report
npm run build --report
```

For a detailed explanation on how things work, check out the [guide](http://vuejs-templates.github.io/webpack/) and [docs for vue-loader](http://vuejs.github.io/vue-loader).
