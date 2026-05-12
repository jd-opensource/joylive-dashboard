<template>
    <a-breadcrumb
        class="x-breadcrumb"
        :class="{ 'x-breadcrumb--dark': theme === 'dark' }">
        <a-breadcrumb-item
            v-for="item in breadcrumbData"
            :key="item.name">
            {{ item.meta.title }}
        </a-breadcrumb-item>
    </a-breadcrumb>
</template>

<script setup>
import { ref } from 'vue'
import { onBeforeRouteUpdate, useRoute } from 'vue-router'

defineOptions({
    name: 'XBreadcrumb',
})

defineProps({
    theme: {
        type: String,
        default: 'light',
    },
})

const route = useRoute()

const breadcrumbData = ref([])

update()

onBeforeRouteUpdate((to) => {
    update(to)
})

function update(_route = route) {
    breadcrumbData.value = _route?.meta?.breadcrumb
}
</script>

<style lang="less" scoped>
.x-breadcrumb {
    display: flex;
    align-items: center;
    padding: 0 12px;

    &--dark {
        :deep(.ant-breadcrumb-separator),
        :deep(.ant-breadcrumb-link) {
            color: rgba(255, 255, 255, 0.65);
        }

        :deep(.ant-breadcrumb > span:last-child .ant-breadcrumb-link) {
            color: rgba(255, 255, 255, 0.85);
        }
    }
}
</style>
