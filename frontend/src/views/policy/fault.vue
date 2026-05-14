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
                        <a-button
                            v-action="'add'"
                            type="primary"
                            @click="$refs.editDialogRef.handleCreate()">
                            <template #icon>
                                <plus-outlined></plus-outlined>
                            </template>
                            {{ $t('pages.fault.add') }}
                        </a-button>
                    </a-col>
                    <a-col flex="auto"></a-col>
                    <a-col flex="none">
                        <a-form
                            :model="searchFormData"
                            layout="inline">
                            <a-form-item
                                :label="$t('pages.fault.form.space_code')"
                                name="space_code"
                                style="margin-bottom: 0">
                                <a-select
                                    :placeholder="$t('pages.fault.form.space_code.placeholder')"
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
                                :label="$t('pages.fault.form.targetServiceId')"
                                name="target_service_id"
                                style="margin-bottom: 0">
                                <a-select
                                    :placeholder="$t('pages.fault.form.targetServiceId.placeholder')"
                                    v-model:value="searchFormData.target_service_id"
                                    show-search
                                    :filter-option="filterServiceOption"
                                    allow-clear
                                    style="width: 200px">
                                    <a-select-option
                                        v-for="item in serviceOptions"
                                        :key="item.id"
                                        :value="item.id">
                                        {{ item.name }}
                                    </a-select-option>
                                </a-select>
                            </a-form-item>
                            <a-form-item
                                :label="$t('pages.fault.form.name')"
                                name="name"
                                style="margin-bottom: 0">
                                <a-input
                                    :placeholder="$t('pages.fault.form.name.placeholder')"
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
                    :scroll="{ x: 1800 }"
                    @change="onTableChange">
                    <template #bodyCell="{ column, record }">
                        <template v-if="'source_application_id' === column.key">
                            {{ applicationNameMap[record.source_application_id] || record.source_application_id }}
                        </template>
                        <template v-if="'target_service_id' === column.key">
                            {{ serviceNameMap[record.target_service_id] || record.target_service_id }}
                        </template>
                        <template v-if="'relation_type' === column.key">
                            {{ relationTypeMap[record.relation_type] || record.relation_type }}
                        </template>
                        <template v-if="'type' === column.key">
                            <a-tag :color="record.type === 'error' ? 'red' : 'orange'">
                                {{ faultTypeMap[record.type] || record.type }}
                            </a-tag>
                        </template>
                        <template v-if="'scope' === column.key">
                            {{ scopeTypeMap[record.scope] || record.scope }}
                        </template>
                        <template v-if="'enabled' === column.key">
                            <a-tag :color="record.enabled === 1 ? 'green' : 'default'">
                                {{
                                    record.enabled === 1
                                        ? $t('pages.fault.form.enabled.active')
                                        : $t('pages.fault.form.enabled.inactive')
                                }}
                            </a-tag>
                        </template>

                        <template v-if="'created_at' === column.key">
                            {{ formatUtcDateTime(record.created_at) }}
                        </template>

                        <template v-if="'action' === column.key">
                            <x-action-button @click="$refs.editDialogRef.handleEdit(record)">
                                <a-tooltip>
                                    <template #title> {{ $t('pages.fault.edit') }}</template>
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
            </a-card>
        </a-col>
    </a-row>

    <edit-dialog
        ref="editDialogRef"
        :space-options="spaceOptions"
        :service-options="serviceOptions"
        :application-options="applicationOptions"
        @ok="onOk"></edit-dialog>
</template>

<script setup>
import { message, Modal } from 'ant-design-vue'
import { ref } from 'vue'
import apis from '@/apis'
import { formatUtcDateTime } from '@/utils/util'
import { config } from '@/config'
import { usePagination } from '@/hooks'
import EditDialog from './FaultEditDialog.vue'
import { PlusOutlined, EditOutlined, DeleteOutlined, SearchOutlined, RedoOutlined } from '@ant-design/icons-vue'
import { useI18n } from 'vue-i18n'
import { initSpaceCode, setCurrentSpaceCode } from '@/utils/spaceStorage'

defineOptions({
    name: 'faultList',
})
const { t } = useI18n()
const columns = [
    { title: t('pages.fault.form.name'), dataIndex: 'name', width: 180 },
    {
        title: t('pages.fault.form.sourceApplicationId'),
        dataIndex: 'source_application_id',
        key: 'source_application_id',
        width: 150,
    },
    {
        title: t('pages.fault.form.targetServiceId'),
        dataIndex: 'target_service_id',
        key: 'target_service_id',
        width: 150,
    },
    { title: t('pages.fault.form.path'), dataIndex: 'path', width: 150 },
    { title: t('pages.fault.form.method'), dataIndex: 'method', width: 100 },
    { title: t('pages.fault.form.relationType'), dataIndex: 'relation_type', key: 'relation_type', width: 100 },
    { title: t('pages.fault.form.type'), dataIndex: 'type', key: 'type', width: 100 },
    { title: t('pages.fault.form.scope'), dataIndex: 'scope', key: 'scope', width: 100 },
    { title: t('pages.fault.form.percent'), dataIndex: 'percent', width: 100 },
    { title: t('pages.fault.form.enabled'), key: 'enabled', width: 80 },
    { title: t('pages.fault.form.creator'), dataIndex: 'creator', width: 100 },
    { title: t('pages.fault.form.description'), dataIndex: 'description', ellipsis: true },
    {
        title: t('pages.fault.form.created_at'),
        dataIndex: 'created_at',
        key: 'created_at',
        fixed: 'right',
        width: 180,
    },
    { title: t('button.action'), key: 'action', fixed: 'right', width: 120 },
]

const { listData, loading, showLoading, hideLoading, paginationState, searchFormData, resetPagination } =
    usePagination()
const editDialogRef = ref()
const spaceOptions = ref([])
const serviceOptions = ref([])
const applicationOptions = ref([])
const serviceNameMap = ref({})
const applicationNameMap = ref({})
const relationTypeMap = {
    OR: '或关系',
    AND: '且关系',
}
const faultTypeMap = {
    error: '错误注入',
    delay: '延迟注入',
}
const scopeTypeMap = {
    inbound: '入方向',
    outbound: '出方向',
}

loadSpaceOptions()
loadServiceOptions()
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

function onSpaceChange(value) {
    setCurrentSpaceCode(value)
    resetPagination()
    getPageList()
}

async function loadServiceOptions() {
    try {
        const { success, data } = await apis.service.getServiceList({ pageSize: 99, current: 1 }).catch(() => {
            throw new Error()
        })
        if (config('http.code.success') === success) {
            serviceOptions.value = data || []
            serviceNameMap.value = Object.fromEntries(serviceOptions.value.map((item) => [item.id, item.name]))
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
            applicationNameMap.value = Object.fromEntries(applicationOptions.value.map((item) => [item.id, item.name]))
        }
    } catch (error) {
        // ignore
    }
}

function filterSpaceOption(input, option) {
    const label = option.children?.[0]?.children || ''
    return option.value.toLowerCase().includes(input.toLowerCase()) || label.toLowerCase().includes(input.toLowerCase())
}

function filterServiceOption(input, option) {
    const label = option.children?.[0]?.children || ''
    return (
        String(option.value).toLowerCase().includes(input.toLowerCase()) ||
        String(label).toLowerCase().includes(input.toLowerCase())
    )
}

async function getPageList() {
    try {
        showLoading()
        const { pageSize, current } = paginationState
        const { success, data, total } = await apis.policy
            .getFaultList({
                pageSize,
                current,
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
        title: t('pages.fault.delTip'),
        content: t('button.confirm'),
        okText: t('button.confirm'),
        onOk: () => {
            return new Promise((resolve, reject) => {
                ;(async () => {
                    try {
                        const { success } = await apis.policy.delFault(id).catch(() => {
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
    resetPagination()
    getPageList()
}

async function onOk() {
    await getPageList()
}
</script>

<style lang="less" scoped>
@import '@/styles/variables.less';

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
