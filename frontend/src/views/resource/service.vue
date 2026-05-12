<template>
    <a-card
        :bordered="false"
        :body-style="{ padding: 0 }"
        class="service-card">
        <a-layout>
            <a-layout-sider
                width="180"
                class="service-sider">
                <a-menu
                    v-model:selectedKeys="selectedKeys"
                    mode="inline"
                    @click="onMenuClick">
                    <a-menu-item key="provider">
                        <template #icon><deployment-unit-outlined /></template>
                        {{ $t('pages.service.tab.provider') }}
                    </a-menu-item>
                    <a-menu-item key="consumer">
                        <template #icon><import-outlined /></template>
                        {{ $t('pages.service.tab.consumer') }}
                    </a-menu-item>
                    <a-menu-item key="all">
                        <template #icon><appstore-outlined /></template>
                        {{ $t('pages.service.tab.all') }}
                    </a-menu-item>
                </a-menu>
            </a-layout-sider>
            <a-layout-content class="service-content">
                <a-form
                    :label-col="{ style: { width: '100px' } }"
                    :model="searchFormData"
                    layout="inline"
                    class="service-search-form">
                    <a-row
                        :gutter="16"
                        style="width: 100%">
                        <a-col
                            :xs="24"
                            :sm="24"
                            :md="12"
                            :xl="8">
                            <a-form-item
                                :label="$t('pages.service.form.space_code')"
                                name="space_code">
                                <a-select
                                    :placeholder="$t('pages.service.form.space_code.placeholder')"
                                    v-model:value="searchFormData.space_code"
                                    show-search
                                    :filter-option="filterSpaceOption"
                                    @change="onSpaceChange">
                                    <a-select-option
                                        v-for="item in spaceOptions"
                                        :key="item.code"
                                        :value="item.code">
                                        {{ item.name }} ({{ item.code }})
                                    </a-select-option>
                                </a-select>
                            </a-form-item>
                        </a-col>
                        <a-col
                            :xs="24"
                            :sm="24"
                            :md="12"
                            :xl="8">
                            <a-form-item
                                :label="$t('pages.service.form.name')"
                                name="name">
                                <a-input
                                    :placeholder="$t('pages.service.form.name.placeholder')"
                                    v-model:value="searchFormData.name"></a-input>
                            </a-form-item>
                        </a-col>
                        <a-col
                            :xs="24"
                            :sm="24"
                            :md="24"
                            :xl="8"
                            class="align-right">
                            <a-space>
                                <a-button @click="handleResetSearch">{{ $t('button.reset') }}</a-button>
                                <a-button
                                    ghost
                                    type="primary"
                                    @click="handleSearch">
                                    {{ $t('button.search') }}
                                </a-button>
                            </a-space>
                        </a-col>
                    </a-row>
                </a-form>

                <x-action-bar class="mb-8-2">
                    <a-button
                        v-action="'add'"
                        type="primary"
                        @click="$refs.editDialogRef.handleCreate()">
                        <template #icon>
                            <plus-outlined></plus-outlined>
                        </template>
                        {{ $t('pages.service.add') }}
                    </a-button>
                </x-action-bar>
                <a-table
                    :columns="columns"
                    :data-source="listData"
                    :loading="loading"
                    :pagination="paginationState"
                    :scroll="{ x: 1200 }"
                    @change="onTableChange">
                    <template #bodyCell="{ column, record }">
                        <template v-if="'createAt' === column.key">
                            {{ formatUtcDateTime(record.created_at) }}
                        </template>

                        <template v-if="'action' === column.key">
                            <x-action-button @click="$refs.editDialogRef.handleEdit(record)">
                                <a-tooltip>
                                    <template #title> {{ $t('pages.service.edit') }}</template>
                                    <edit-outlined />
                                </a-tooltip>
                            </x-action-button>
                            <x-action-button @click="handleRemove(record)">
                                <a-tooltip>
                                    <template #title> {{ $t('pages.system.delete') }}</template>
                                    <delete-outlined style="color: #ff4d4f" />
                                </a-tooltip>
                            </x-action-button>
                        </template>
                    </template>
                </a-table>
            </a-layout-content>
        </a-layout>
    </a-card>

    <edit-dialog
        ref="editDialogRef"
        :space-options="spaceOptions"
        :application-options="applicationOptions"
        @ok="onOk"></edit-dialog>
</template>

<script setup>
import { message, Modal } from 'ant-design-vue'
import { ref } from 'vue'
import apis from '@/apis'
import { formatUtcDateTime } from '@/utils/util'
import { config } from '@/config'
import { usePagination, useForm } from '@/hooks'
import EditDialog from './ServiceEditDialog.vue'
import {
    PlusOutlined,
    EditOutlined,
    DeleteOutlined,
    DeploymentUnitOutlined,
    ImportOutlined,
    AppstoreOutlined,
} from '@ant-design/icons-vue'
import { useI18n } from 'vue-i18n'

defineOptions({
    name: 'serviceList',
})
const { t } = useI18n()
const columns = [
    { title: t('pages.service.form.name'), dataIndex: 'name', width: 200 },
    { title: t('pages.service.form.space_code'), dataIndex: 'space_code', width: 150 },
    { title: t('pages.service.form.registration_type'), dataIndex: 'registration_type', width: 120 },
    { title: t('pages.service.form.source'), dataIndex: 'source', width: 100 },
    { title: t('pages.service.form.creator'), dataIndex: 'creator', width: 120 },
    { title: t('pages.service.form.version'), dataIndex: 'version', width: 80 },
    { title: t('pages.service.form.description'), dataIndex: 'description', ellipsis: true },
    { title: t('pages.service.form.created_at'), key: 'createAt', fixed: 'right', width: 180 },
    { title: t('button.action'), key: 'action', fixed: 'right', width: 120 },
]

const { listData, loading, showLoading, hideLoading, paginationState, searchFormData, resetPagination } =
    usePagination()
const { resetForm } = useForm()
const editDialogRef = ref()
const spaceOptions = ref([])
const applicationOptions = ref([])
const selectedKeys = ref(['provider'])

const SPACE_CODE_KEY = 'service_space_code'

function currentRole() {
    return selectedKeys.value[0] === 'all' ? '' : selectedKeys.value[0]
}

function onMenuClick() {
    resetPagination()
    getPageList()
}

loadSpaceOptions()
loadApplicationOptions()

async function loadSpaceOptions() {
    try {
        const { success, data } = await apis.space.getSpaceList({ pageSize: 99, current: 1 }).catch(() => {
            throw new Error()
        })
        if (config('http.code.success') === success) {
            spaceOptions.value = data || []
            if (spaceOptions.value.length > 0) {
                const saved = localStorage.getItem(SPACE_CODE_KEY)
                const found = saved && spaceOptions.value.some((item) => item.code === saved)
                searchFormData.value.space_code = found ? saved : spaceOptions.value[0].code
                localStorage.setItem(SPACE_CODE_KEY, searchFormData.value.space_code)
                getPageList()
            }
        }
    } catch (error) {
        // ignore
    }
}

async function loadApplicationOptions() {
    try {
        const { success, data } = await apis.application.getApplicationList({ pageSize: 99, current: 1 }).catch(() => {
            throw new Error()
        })
        if (config('http.code.success') === success) {
            applicationOptions.value = data || []
        }
    } catch (error) {
        // ignore
    }
}

function onSpaceChange(value) {
    localStorage.setItem(SPACE_CODE_KEY, value)
    resetPagination()
    getPageList()
}

function filterSpaceOption(input, option) {
    const label = option.children?.[0]?.children || ''
    return option.value.toLowerCase().includes(input.toLowerCase()) || label.toLowerCase().includes(input.toLowerCase())
}

async function getPageList() {
    try {
        showLoading()
        const { pageSize, current } = paginationState
        const { success, data, total } = await apis.service
            .getServiceList({
                pageSize,
                current,
                role: currentRole(),
                ...searchFormData.value,
            })
            .catch(() => {
                throw new Error()
            })
        hideLoading()
        if (config('http.code.success') === success) {
            listData.value = data
            paginationState.total = total
        }
    } catch (error) {
        hideLoading()
    }
}

function handleRemove({ id }) {
    Modal.confirm({
        title: t('pages.service.delTip'),
        content: t('button.confirm'),
        okText: t('button.confirm'),
        onOk: () => {
            return new Promise((resolve, reject) => {
                ;(async () => {
                    try {
                        const { success } = await apis.service.delService(id).catch(() => {
                            throw new Error()
                        })
                        if (config('http.code.success') === success) {
                            resolve()
                            message.success(t('component.message.success.delete'))
                            await getPageList()
                        }
                    } catch (error) {
                        reject()
                    }
                })()
            })
        },
    })
}

function onTableChange({ current, pageSize }) {
    paginationState.current = current
    paginationState.pageSize = pageSize
    getPageList()
}

function handleResetSearch() {
    const spaceCode = searchFormData.value.space_code
    searchFormData.value = { space_code: spaceCode }
    resetPagination()
    getPageList()
}

function handleSearch() {
    resetForm()
    resetPagination()
    getPageList()
}

async function onOk() {
    await getPageList()
}
</script>

<style lang="less" scoped>
.service-card {
    .ant-layout {
        background: #fff;
    }
}

.service-sider {
    background: #fff;
    border-right: 1px solid #f0f0f0;

    :deep(.ant-menu) {
        border-right: none;
        padding-top: 8px;
    }

    :deep(.ant-menu-item) {
        margin: 4px 8px;
        border-radius: 6px;
        height: 40px;
        line-height: 40px;
    }
}

.service-content {
    padding: 24px;
}

.service-search-form {
    margin-bottom: 16px;

    :deep(.ant-form-inline .ant-form-item) {
        margin-right: 0;
        margin-bottom: 16px;
    }

    :deep(.ant-form-item-control-input-content) {
        > * {
            width: 100%;
        }
    }
}
</style>
