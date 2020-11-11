<template>
<div id="irain">
    <a-layout id="components-layout-demo-custom-trigger">
        <a-layout-sider v-model="collapsed" :trigger="null" collapsible>
            <div class="logo">
                <!-- <img v-show="collapsed"  src="" alt="logo">
                <img v-show="!collapsed" src="" alt="logo"> -->
            </div>
            <a-menu theme="dark" mode="inline" :default-selected-keys="['1']">
                <a-menu-item key="1">
                <a-icon type="user" />
                <span>nav 1</span>
                </a-menu-item>
                <a-menu-item key="2">
                <a-icon type="video-camera" />
                <span>nav 2</span>
                </a-menu-item>
                <a-menu-item key="3">
                <a-icon type="upload" />
                <span>nav 3</span>
                </a-menu-item>
            </a-menu>
        </a-layout-sider>
        <a-layout>
            <a-layout-header style="background: #fff; padding: 0">
                <header-bar :collapsed="collapsed" @on-coll-change="collChange"></header-bar>
            </a-layout-header>
            <a-layout-content
                :style="{ margin: '24px 16px', padding: '24px', background: '#fff', minHeight: '280px' }"
            >
                Content
            </a-layout-content>
        </a-layout>
  </a-layout>
</div>
</template>
<script>
// 引入css
import './main.less'
// 引入antd
import {
    Layout,
    Icon,
    Menu
} from 'ant-design-vue'
const {
    Header,
    Sider,
    Content
} = Layout;
const MenuItem = Menu.Item
const SubMenu = Menu.SubMenu
// 引入自定义组件
import HeaderBar from './components/header-bar'
export default {
    name: 'Main',
    components: {
        'a-layout': Layout,
        'a-layout-sider': Sider,
        'a-layout-header': Header,
        'a-layout-content': Content,
        'a-header': Header,
        'a-icon': Icon,
        'a-menu': Menu,
        'a-menu-item': MenuItem,
        HeaderBar
    },
    data () {
        return {
            collapsed: false
        }
    },
    computed: {

    },
    created() {

    },
    mounted () {

    },
    methods: {
        // 侧边栏伸缩
        collChange(state) {
            this.collapsed = state
        }
    }
}
</script>