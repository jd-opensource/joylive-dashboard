<template>
    <div class="user-layout-container">
        <div class="user-layout-aside">
            <div class="aside-header">
                <h1>{{ title }}</h1>
            </div>
            <div class="aside-body">
                <img
                    alt=""
                    :src="assets('logos.png')" />
                <h3>{{ $t('pages.layouts.userLayout.title') }}</h3>
                <!--                <p>Vue3 + Ant Design Vue + vite</p>-->
            </div>
            <div class="aside-footer">
                <p>© {{ title }} {{ version }}</p>
            </div>
        </div>
        <div class="user-layout-main">
            <div class="user-layout-content">
                <div class="user-layout-top">
                    <div class="user-layout-header">{{ $t('login') }}</div>
                    <!--                    <div class="user-layout-desc">欢迎使用{{ title }}</div>-->
                </div>
                <div class="user-layout-form">
                    <router-view></router-view>
                </div>
            </div>
        </div>

        <div
            class="basic-header__right"
            style="padding: 30px">
            <a-space :size="16">
                <a-dropdown :trigger="['hover']">
                    <action-button :style="{ height: '44px' }">
                        <translation-outlined />
                    </action-button>
                    <a-spin />
                    <template #overlay>
                        <a-menu v-model:selectedKeys="current">
                            <a-menu-item
                                v-for="(item, key) in langData"
                                :key="key"
                                @click="handleLang(key)">
                                {{ item.icon }} {{ item.label }}
                            </a-menu-item>
                        </a-menu>
                    </template>
                </a-dropdown>
            </a-space>
        </div>
    </div>
</template>

<script setup>
import { assets } from '@/utils/util'
import { config as conf, config } from '@/config'
import { ref, watch, onMounted } from 'vue'
import { TranslationOutlined } from '@ant-design/icons-vue'
import { useAppStore } from '@/store'
import { storeToRefs } from 'pinia'

import storage from '@/utils/storage'
import { useI18n } from 'vue-i18n'
const { locale } = useI18n()
defineOptions({
    name: 'UserLayout',
})

const { version } = __APP_INFO__
const title = config('app.title')
const defaultLang = storage.local.getItem(conf('storage.lang')) || 'zh-ch'
const current = ref(defaultLang)
const langData = ref({
    'zh-ch': {
        lang: 'zh-ch',
        label: '简体中文',
        icon: '🇨🇳',
        title: '语言',
    },
    'en-us': {
        lang: 'en-us',
        label: 'English',
        icon: '🇺🇸',
        title: 'Language',
    },
})

// 获取主题配置
const appStore = useAppStore()
const { config: appConfig } = storeToRefs(appStore)

// 应用主题到DOM
const applyTheme = (theme) => {
    if (theme === 'dark') {
        document.documentElement.setAttribute('data-theme', 'dark')
        document.body.setAttribute('data-theme', 'dark')
    } else {
        document.documentElement.removeAttribute('data-theme')
        document.body.removeAttribute('data-theme')
    }
}

// 监听主题变化
watch(
    () => appConfig.value.theme,
    (newTheme) => {
        applyTheme(newTheme)
    },
    { immediate: true }
)

// 组件挂载时应用主题
onMounted(() => {
    applyTheme(appConfig.value.theme)
})

/**
 * 切换语言
 */

function handleLang(lang) {
    storage.local.setItem(conf('storage.lang'), lang)
    locale.value = lang
    current.value = lang
    location.reload()
}
</script>

<style lang="less" scoped>
.user-layout {
    &-container {
        min-height: 100vh;
        background-repeat: no-repeat;
        background-position: center 110px;
        background-size: 100%;
        display: flex;
        transition: background-color 0.3s ease;
    }

    &-aside {
        width: 538px;
        flex: 0 0 538px;
        display: flex;
        flex-direction: column;
        background: #235bda url('@/assets/login_aside_bg.jpg') no-repeat left top / 100% auto;
        position: relative;
        transition: background 0.3s ease;

        .aside {
            &-header {
                display: flex;
                flex-direction: column;
                padding: 48px;

                h1 {
                    font-size: 20px;
                    font-weight: 500;
                    color: #fff;
                    transition: color 0.3s ease;
                }
            }

            &-body {
                flex: 1;
                text-align: center;
                padding: 48px 0 0;

                img {
                    width: 80%;
                    transition: opacity 0.3s ease;
                }

                h3 {
                    color: #fff;
                    transition: color 0.3s ease;
                }

                p {
                    color: rgba(255, 255, 255, 0.85);
                    transition: color 0.3s ease;
                }
            }

            &-footer {
                color: rgba(255, 255, 255, 0.65);
                font-size: 12px;
                padding: 48px;
                transition: color 0.3s ease;
            }
        }
    }

    &-main {
        flex: 1;
        display: flex;
        align-items: center;
        justify-content: center;
        padding: 64px 0 144px;
        transition: background-color 0.3s ease;
    }

    &-content {
        width: 368px;
        height: 440px;
    }

    &-header {
        display: flex;
        align-items: center;
        font-size: 30px;
        font-weight: 500;
        transition: color 0.3s ease;
    }

    &-desc {
        margin: 8px 0 24px;
        transition: color 0.3s ease;
    }
}

.basic-header__right {
    position: absolute;
    top: 0;
    right: 0;
    z-index: 10;
}

// 暗黑模式适配
html[data-theme='dark'],
:root[data-theme='dark'],
body[data-theme='dark'] {
    background-color: #141414;

    .user-layout {
        &-container {
            background-color: #141414;
        }

        &-aside {
            background: linear-gradient(135deg, #1a1a2e 0%, #16213e 50%, #0f3460 100%)
                url('@/assets/login_aside_bg.jpg') no-repeat left top / 100% auto;
            background-blend-mode: soft-light;

            .aside {
                &-body {
                    img {
                        opacity: 0.9;
                        filter: brightness(0.95);
                    }
                }
            }
        }

        &-main {
            background-color: #141414;
        }

        &-header {
            color: rgba(255, 255, 255, 0.85);
        }

        &-desc {
            color: rgba(255, 255, 255, 0.65);
        }
    }
}
</style>
