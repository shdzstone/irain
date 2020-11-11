<template>
    <div>
        <a-icon
            class="trigger"
            :type="collapsed ? 'menu-unfold' : 'menu-fold'"
            @click="handleChange"
        />
    </div>
</template>
<script>
import { Icon } from 'ant-design-vue';
export default {
    name: "SiderTrigger",
    components: {
        'a-icon': Icon,
    },
    props: {
        collapsed: Boolean
    },
    methods: {
        handleChange () {
            this.$emit('on-change', !this.collapsed)
        }
    }
}
</script>