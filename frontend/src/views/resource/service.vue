<template>
    <a-row
        :gutter="8"
        :wrap="false">
        <a-col flex="auto">
            <a-card type="flex">
                <a-row
                    :gutter="16"
                    align="middle"
                    class="mb-8-2">
                    <a-col flex="none">
                        <a-radio-group
                            v-model:value="activeTab"
                            button-style="solid"
                            @change="onTabChange">
                            <a-radio-button value="provider">
                                {{ $t('pages.service.tab.provider') }}
                            </a-radio-button>
                            <a-radio-button value="consumer">
                                {{ $t('pages.service.tab.consumer') }}
                            </a-radio-button>
                            <a-radio-button value="all">
                                {{ $t('pages.service.tab.all') }}
                            </a-radio-button>
                        </a-radio-group>
                    </a-col>
                    <a-col flex="none">
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
                    </a-col>
                    <a-col flex="auto"></a-col>
                    <a-col flex="none">
                        <a-form
                            :model="searchFormData"
                            layout="inline">
                            <a-form-item
                                :label="$t('pages.service.form.space_code')"
                                name="space_code"
                                style="margin-bottom: 0">
                                <a-select
                                    :placeholder="$t('pages.service.form.space_code.placeholder')"
                                    v-model:value="searchFormData.space_code"
                                    show-search
                                    :filter-option="filterSpaceOption"
                                    @change="onSpaceChange"
                                    style="width: 200px">
                                    <a-select-option
                                        v-for="item in spaceOptions"
                                        :key="item.code"
                                        :value="item.code">
                                        {{ item.name }} ({{ item.code }})
                                    </a-select-option>
                                </a-select>
                            </a-form-item>
                            <a-form-item
                                :label="$t('pages.service.form.name')"
                                name="name"
                                style="margin-bottom: 0">
                                <a-input
                                    :placeholder="$t('pages.service.form.name.placeholder')"
                                    v-model:value="searchFormData.name"
                                    style="width: 200px"
                                    @pressEnter="handleSearch"></a-input>
                            </a-form-item>
                            <a-form-item style="margin-bottom: 0">
                                <a-space :size="8">
                                    <a-tooltip :title="$t('button.reset')">
                                        <a-button
                                            shape="circle"
                                            @click="handleResetSearch">
                                            <template #icon>
                                                <redo-outlined />
                                            </template>
                                        </a-button>
                                    </a-tooltip>
                                    <a-tooltip :title="$t('button.search')">
                                        <a-button
                                            type="primary"
                                            shape="circle"
                                            @click="handleSearch">
                                            <template #icon>
                                                <search-outlined />
                                            </template>
                                        </a-button>
                                    </a-tooltip>
                                </a-space>
                            </a-form-item>
                        </a-form>
                    </a-col>
                </a-row>
                <a-table
                    :columns="columns"
                    :data-source="listData"
                    :loading="loading"
                    :pagination="paginationState"
                    :scroll="{ x: 1260 }"
                    @change="onTableChange">
                    <template #bodyCell="{ column, record }">
                        <template v-if="'name' === column.key">
                            <a
                                v-if="activeTab === 'provider'"
                                @click="goToDetail(record)">
                                {{ record.name }}
                            </a>
                            <span v-else>{{ record.name }}</span>
                        </template>
                        <template v-if="'created_at' === column.key">
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
    SearchOutlined,
    RedoOutlined,
} from '@ant-design/icons-vue'
import { useI18n } from 'vue-i18n'
import { useRouter } from 'vue-router'
import { initSpaceCode, setCurrentSpaceCode } from '@/utils/spaceStorage'

defineOptions({
    name: 'serviceList',
})
const router = useRouter()
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
    { title: t('pages.service.form.name'), dataIndex: 'name', key: 'name', width: 280 },
    {
        title: t('pages.service.form.registration_type'),
        dataIndex: 'registration_type',
        key: 'registration_type',
        width: 120,
    },
    { title: t('pages.service.form.creator'), dataIndex: 'creator', width: 120 },
    { title: t('pages.service.form.version'), dataIndex: 'version', width: 80 },
    { title: t('pages.service.form.description'), dataIndex: 'description', ellipsis: true },
    { title: t('pages.service.form.created_at'), key: 'created_at', fixed: 'right', width: 180 },
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
                searchFormData.value.space_code = initSpaceCode(spaceOptions.value)
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
    setCurrentSpaceCode(value)
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

function goToDetail(record) {
    router.push({ name: 'serviceDetail', params: { id: record.id } })
}

async function onOk() {
    await getPageList()
}
</script>

<style lang="less" scoped>
@import '@/styles/variables.less';

// Tab 按钮组样式
:deep(.ant-radio-group) {
    &.ant-radio-group-solid {
        .ant-radio-button-wrapper {
            border-color: @color-border;

            &:hover {
                color: @color-primary;
            }

            &.ant-radio-button-wrapper-checked {
                background: @color-primary;
                border-color: @color-primary;
            }
        }
    }
}

// 搜索栏和操作按钮行
:deep(.ant-form-inline) {
    .ant-form-item {
        margin-right: 16px;

        &:last-child {
            margin-right: 0;
        }
    }
}

// 表格行悬停效果 - 轻微优化
:deep(.ant-table-tbody > tr:hover > td) {
    background-color: #fafafa;
}

// 服务名称链接 - 添加平滑过渡
:deep(.ant-table-tbody) {
    a {
        color: @color-primary;
        transition: color 0.2s ease;

        &:hover {
            color: #0958d9;
        }
    }
}

// 操作按钮 - 添加悬停效果
:deep(.x-action-button) {
    transition: all 0.2s ease;

    &:hover {
        background-color: rgba(0, 0, 0, 0.04);
        border-radius: 4px;
    }
}

// 状态标签 - 轻微优化圆角
:deep(.ant-tag) {
    border-radius: 4px;
}

// 搜索按钮 - 添加轻微过渡
:deep(.ant-btn) {
    transition: all 0.2s ease;
}

// 表格单元格 - 优化间距
:deep(.ant-table) {
    .ant-table-tbody > tr > td {
        padding: 12px 16px;
    }

    .ant-table-thead > tr > th {
        padding: 12px 16px;
        font-weight: 600;
    }
}

// 搜索栏分隔线
.mb-8-2 {
    padding-bottom: 16px;
    border-bottom: 1px solid #f0f0f0;
    margin-bottom: 16px;
}
</style>
