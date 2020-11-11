<template>
    <div class="irain-header">
        <div class="irain-header-le">
            <sider-trigger :collapsed="collapsed" @on-change="handleCollapsedChange"></sider-trigger>
            <custom-bread-crumb></custom-bread-crumb>
        </div>
        <div class="irain-header-ri">2</div>
    </div>
</template>
<script>
import './header-bar.less';

import SiderTrigger from './sider-trigger'
import CustomBreadCrumb from './custom-bread-crumb'
export default {
    name: 'HeaderBar',
    components: {
        SiderTrigger,
        CustomBreadCrumb
    },
    props: {
        collapsed: Boolean
    },
    methods: {
        handleCollapsedChange(state) {
            this.$emit('on-coll-change', state)
        }
    }
}
</script>