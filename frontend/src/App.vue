<template>
    <a-config-provider
        :locale="zhCN"
        :theme="theme">
        <router-view />
    </a-config-provider>
</template>

<script setup>
import { computed } from 'vue'
import zhCN from 'ant-design-vue/es/locale/zh_CN'
import { theme as antdTheme } from 'ant-design-vue'
import dayjs from 'dayjs'
import 'dayjs/locale/zh-cn'
import { useAppStore } from '@/store'
import { storeToRefs } from 'pinia'

dayjs.locale('zh-cn')

const appStore = useAppStore()
const { config } = storeToRefs(appStore)

const theme = computed(() => {
    return {
        algorithm: config.value.theme === 'dark' ? antdTheme.darkAlgorithm : antdTheme.defaultAlgorithm,
        components: {
            List: {
                paddingContentHorizontalLG: 0,
            },
            Table: {
                paddingContentVerticalLG: 12,
                padding: 12,
            },
            Card: {
                paddingLG: 16,
            },
        },
    }
})
</script>

<style lang="less"></style>
