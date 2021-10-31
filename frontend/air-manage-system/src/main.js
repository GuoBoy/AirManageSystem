import { createApp } from 'vue'
import App from './App.vue'
import router from './router'
import store from './store'

// 导入elemrntui
import elementui from './common/elementui'
import 'element-plus/dist/index.css'
import './assets/global.css'

const app = createApp(App)

elementui.forEach(element => {
    app.use(element)
})

app.use(store).use(router).mount('#app')
