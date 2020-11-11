<template>
    <div class="left-custom">
        <a-breadcrumb class="a-custom-bread-crumb">
            <a-breadcrumb-item class="breadcrumb-hover">Home</a-breadcrumb-item>
            <a-breadcrumb-item class="breadcrumb-hover">测试</a-breadcrumb-item>
            <a-breadcrumb-item class="breadcrumb-hover">测试1</a-breadcrumb-item>
        </a-breadcrumb>
    </div>
</template>
<script>
import './custom-bread-crumb.less'
import {
    Breadcrumb
} from 'ant-design-vue'
const BreadcrumbItem = Breadcrumb.Item

export default {
    name: 'CustomBreadCrumb',
    components: {
        'a-breadcrumb': Breadcrumb,
        'a-breadcrumb-item': BreadcrumbItem
    }
}
</script>