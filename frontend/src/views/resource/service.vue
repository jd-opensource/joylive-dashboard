<template>
    <a-tabs
        v-model:activeKey="activeTab"
        @change="onTabChange">
        <a-tab-pane
            key="provider"
            :tab="$t('pages.service.tab.provider')" />
        <a-tab-pane
            key="consumer"
            :tab="$t('pages.service.tab.consumer')" />
        <a-tab-pane
            key="all"
            :tab="$t('pages.service.tab.all')" />
    </a-tabs>
    <x-search-bar class="mb-8-2">
        <template #default="{ gutter, colSpan }">
            <a-form
                :label-col="{ style: { width: '100px' } }"
                :model="searchFormData"
                layout="inline">
                <a-row :gutter="gutter">
                    <a-col v-bind="colSpan">
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
                    <a-col v-bind="colSpan">
                        <a-form-item
                            :label="$t('pages.service.form.name')"
                            name="name">
                            <a-input
                                :placeholder="$t('pages.service.form.name.placeholder')"
                                v-model:value="searchFormData.name"></a-input>
                        </a-form-item>
                    </a-col>
                    <a-col
                        class="align-right"
                        v-bind="colSpan">
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
        </template>
    </x-search-bar>
    <a-row
        :gutter="8"
        :wrap="false">
        <a-col flex="auto">
            <a-card type="flex">
                <x-action-bar class="mb-8-2">
                    <a-button
                        v-if="activeTab === 'provider'"
                        v-action="'add'"
                        type="primary"
                        @click="$refs.editDialogRef.handleCreate()">
                        <template #icon>
                            <plus-outlined></plus-outlined>
                        </template>
                        {{ $t('pages.service.create') }}
                    </a-button>
                </x-action-bar>
                <a-table
                    :columns="columns"
                    :data-source="listData"
                    :loading="loading"
                    :pagination="paginationState"
                    :scroll="{ x: 1260 }"
                    @change="onTableChange">
                    <template #bodyCell="{ column, record }">
                        <template v-if="'createAt' === column.key">
                            {{ formatUtcDateTime(record.created_at) }}
                        </template>
                        <template v-if="'registration_type' === column.key">
                            {{ registrationTypeMap[record.registration_type] || record.registration_type }}
                        </template>
                        <template v-if="'application_service_status' === column.key">
                            <a-tag :color="statusColorMap[record.application_service_status]">
                                {{
                                    statusTextMap[record.application_service_status] ||
                                    record.application_service_status
                                }}
                            </a-tag>
                        </template>

                        <template v-if="'action' === column.key">
                            <template v-if="activeTab === 'all'">
                                <x-action-button @click="$refs.applyDialogRef.handleApply(record)">
                                    <a-tooltip>
                                        <template #title> {{ $t('pages.service.apply') }}</template>
                                        <import-outlined />
                                    </a-tooltip>
                                </x-action-button>
                            </template>
                            <template v-else>
                                <x-action-button
                                    v-if="activeTab !== 'consumer'"
                                    @click="$refs.editDialogRef.handleEdit(record)">
                                    <a-tooltip>
                                        <template #title> {{ $t('pages.service.edit') }}</template>
                                        <edit-outlined />
                                    </a-tooltip>
                                </x-action-button>
                                <x-action-button
                                    v-if="activeTab === 'provider'"
                                    @click="handleToggleAuth(record)">
                                    <a-tooltip>
                                        <template #title>
                                            {{
                                                isAuthorized(record)
                                                    ? $t('pages.service.auth.disable')
                                                    : $t('pages.service.auth.enable')
                                            }}
                                        </template>
                                        <lock-outlined v-if="!isAuthorized(record)" />
                                        <unlock-outlined
                                            v-else
                                            style="color: #52c41a" />
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
                    </template>
                </a-table>
            </a-card>
        </a-col>
    </a-row>

    <edit-dialog
        ref="editDialogRef"
        :space-options="spaceOptions"
        :application-options="applicationOptions"
        @ok="onOk"></edit-dialog>
    <apply-dialog
        ref="applyDialogRef"
        :application-options="applicationOptions"
        @ok="onOk"></apply-dialog>
</template>

<script setup>
import { message, Modal } from 'ant-design-vue'
import { ref, computed } from 'vue'
import apis from '@/apis'
import { formatUtcDateTime } from '@/utils/util'
import { config } from '@/config'
import { usePagination } from '@/hooks'
import EditDialog from './ServiceEditDialog.vue'
import ApplyDialog from './ServiceApplyDialog.vue'
import {
    PlusOutlined,
    EditOutlined,
    DeleteOutlined,
    ImportOutlined,
    LockOutlined,
    UnlockOutlined,
} from '@ant-design/icons-vue'
import { useI18n } from 'vue-i18n'

defineOptions({
    name: 'serviceList',
})
const { t } = useI18n()
const registrationTypeMap = {
    HTTP: t('pages.service.form.registration_type.http'),
    RPC_APP: t('pages.service.form.registration_type.rpc_app'),
    RPC_INTERFACE: t('pages.service.form.registration_type.rpc_interface'),
}
const statusTextMap = {
    approved: t('pages.service.form.application_service_status.approved'),
    pending: t('pages.service.form.application_service_status.pending'),
    rejected: t('pages.service.form.application_service_status.rejected'),
}
const statusColorMap = {
    approved: 'green',
    pending: 'orange',
    rejected: 'red',
}
const baseColumns = [
    { title: t('pages.service.form.name'), dataIndex: 'name', width: 200 },
    {
        title: t('pages.service.form.registration_type'),
        dataIndex: 'registration_type',
        key: 'registration_type',
        width: 120,
    },
    { title: t('pages.service.form.creator'), dataIndex: 'creator', width: 120 },
    { title: t('pages.service.form.version'), dataIndex: 'version', width: 80 },
    { title: t('pages.service.form.description'), dataIndex: 'description', ellipsis: true },
    { title: t('pages.service.form.created_at'), key: 'createAt', fixed: 'right', width: 180 },
    { title: t('button.action'), key: 'action', fixed: 'right', width: 180 },
]
const consumerColumns = [
    { title: t('pages.service.form.application_name'), dataIndex: 'application_name', width: 150 },
    {
        title: t('pages.service.form.application_service_status'),
        dataIndex: 'application_service_status',
        key: 'application_service_status',
        width: 100,
    },
]
const columns = computed(() => {
    if (activeTab.value === 'consumer') {
        const cols = [...baseColumns]
        cols.splice(1, 0, ...consumerColumns)
        return cols
    }
    return baseColumns
})

const { listData, loading, showLoading, hideLoading, paginationState, searchFormData, resetPagination } =
    usePagination()
const editDialogRef = ref()
const spaceOptions = ref([])
const applicationOptions = ref([])
const activeTab = ref('provider')

const SPACE_CODE_KEY = 'service_space_code'

function currentRole() {
    return activeTab.value === 'all' ? '' : activeTab.value
}

function onTabChange() {
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
                        const deleteFn =
                            activeTab.value === 'consumer' ? apis.service.delServiceConsumer : apis.service.delService
                        const { success } = await deleteFn(id).catch(() => {
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

function isAuthorized(record) {
    try {
        const extra = record.extra ? JSON.parse(record.extra) : {}
        return extra.authorized === 1
    } catch {
        return false
    }
}

function handleToggleAuth(record) {
    const authorized = isAuthorized(record) ? 0 : 1
    const title = authorized === 1 ? t('pages.service.auth.enableTip') : t('pages.service.auth.disableTip')
    Modal.confirm({
        title,
        content: t('button.confirm'),
        okText: t('button.confirm'),
        onOk: () => {
            return new Promise((resolve, reject) => {
                ;(async () => {
                    try {
                        const { success } = await apis.service
                            .toggleServiceAuth(record.id, { authorized })
                            .catch(() => {
                                throw new Error()
                            })
                        if (config('http.code.success') === success) {
                            resolve()
                            message.success(t('component.message.success.save'))
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
    resetPagination()
    getPageList()
}

async function onOk() {
    await getPageList()
}
</script>

<style lang="less" scoped></style>
