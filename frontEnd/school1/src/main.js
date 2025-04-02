import { createApp } from 'vue';
import App from './App.vue';
import router from './router';
import Web3 from 'web3';
import { createPinia } from 'pinia';
import { ethers } from 'ethers';
import ElementPlus from 'element-plus';
import 'element-plus/dist/index.css';
import * as ElementPlusIconsVue from '@element-plus/icons-vue';
import zhCn from 'element-plus/es/locale/lang/zh-cn';
// 修改导入路径
import AiQaComponent from './views/AiQaComponent.vue'; 

const app = createApp(App);

// 注册所有图标
for (const [key, component] of Object.entries(ElementPlusIconsVue)) {
    app.component(key, component);
}

// 全局注册 AI 问答组件
app.component('AiQaComponent', AiQaComponent);

app.use(createPinia());
app.use(router);
app.use(ElementPlus, {
    locale: zhCn,
});

// 创建 Web3 实例
const web3 = new Web3('http://127.0.0.1:7545'); // 连接到 Ganache 或其他以太坊节点

// 将 Web3 实例挂载到 Vue 的全局属性上
app.config.globalProperties.$web3 = web3;

// 将 ethers 添加到全局属性中
app.config.globalProperties.$ethers = ethers;

app.mount('#app');
    