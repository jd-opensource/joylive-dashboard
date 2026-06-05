<template>
    <div class="invocation-page">
        <a-card
            type="flex"
            class="invocation-card">
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
                        {{ $t('pages.invocation.add') }}
                    </a-button>
                </a-col>
                <a-col flex="auto"></a-col>
                <a-col flex="none">
                    <a-form
                        :model="searchFormData"
                        layout="inline">
                        <a-form-item
                            :label="$t('pages.invocation.form.name')"
                            name="name"
                            style="margin-bottom: 0">
                            <a-input
                                :placeholder="$t('pages.invocation.form.name.placeholder')"
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
                :scroll="{ x: 'max-content' }"
                @change="onTableChange">
                <template #bodyCell="{ column, record }">
                    <template v-if="'type' === column.key">
                        <a-tag :color="invocationTypeColorMap[record.type] || 'default'">
                            {{ invocationTypeMap[record.type] || record.type }}
                        </a-tag>
                    </template>
                    <template v-if="'enabled' === column.key">
                        <a-tag :color="record.enabled === 1 ? 'green' : 'default'">
                            {{
                                record.enabled === 1
                                    ? $t('pages.invocation.form.enabled.active')
                                    : $t('pages.invocation.form.enabled.inactive')
                            }}
                        </a-tag>
                    </template>
                    <template v-if="'created_at' === column.key">
                        {{ formatUtcDateTime(record.created_at) }}
                    </template>
                    <template v-if="'action' === column.key">
                        <x-action-button @click="$refs.editDialogRef.handleEdit(record)">
                            <a-tooltip>
                                <template #title> {{ $t('pages.invocation.edit') }}</template>
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

        <edit-dialog
            ref="editDialogRef"
            @ok="onOk"></edit-dialog>
    </div>
</template>

<script setup>
import { message, Modal } from 'ant-design-vue'
import { ref } from 'vue'
import apis from '@/apis'
import { formatUtcDateTime } from '@/utils/util'
import { config } from '@/config'
import { usePagination } from '@/hooks'
import EditDialog from './InvocationEditDialog.vue'
import { PlusOutlined, EditOutlined, DeleteOutlined, SearchOutlined, RedoOutlined } from '@ant-design/icons-vue'
import { useI18n } from 'vue-i18n'

defineOptions({
    name: 'invocationList',
})
const { t } = useI18n()
const columns = [
    { title: t('pages.invocation.form.name'), dataIndex: 'name', width: 200 },
    { title: t('pages.invocation.form.type'), dataIndex: 'type', key: 'type', width: 120 },
    { title: t('pages.invocation.form.enabled'), key: 'enabled', width: 80 },
    { title: t('pages.invocation.form.creator'), dataIndex: 'creator', width: 100 },
    { title: t('pages.invocation.form.description'), dataIndex: 'description', width: 300, ellipsis: true },
    {
        title: t('pages.invocation.form.created_at'),
        dataIndex: 'created_at',
        key: 'created_at',
        fixed: 'right',
        width: 160,
    },
    { title: t('button.action'), key: 'action', fixed: 'right', width: 100 },
]

const { listData, loading, showLoading, hideLoading, paginationState, searchFormData, resetPagination } =
    usePagination()
const editDialogRef = ref()
const invocationTypeMap = {
    failfast: t('pages.invocation.form.type.failfast'),
    failover: t('pages.invocation.form.type.failover'),
    failsafe: t('pages.invocation.form.type.failsafe'),
}
const invocationTypeColorMap = {
    failfast: 'red',
    failover: 'orange',
    failsafe: 'blue',
}

getPageList()

async function getPageList() {
    try {
        showLoading()
        const { pageSize, current } = paginationState
        const { success, data, total } = await apis.policy
            .getInvocationList({
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
        title: t('pages.invocation.delTip'),
        content: t('button.confirm'),
        okText: t('button.confirm'),
        okType: 'danger',
        onOk: () => {
            return new Promise((resolve, reject) => {
                ;(async () => {
                    try {
                        const { success } = await apis.policy.delInvocation(id).catch(() => {
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
    searchFormData.value = {}
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

:deep(.ant-form-inline) {
    .ant-form-item {
        margin-right: 16px;

        &:last-child {
            margin-right: 0;
        }
    }
}

:deep(.x-action-button) {
    transition: all 0.2s ease;

    &:hover {
        background-color: rgba(0, 0, 0, 0.04);
        border-radius: 4px;
    }
}

:deep(.ant-tag) {
    border-radius: 4px;
}

:deep(.ant-btn) {
    transition: all 0.2s ease;
}

:deep(.ant-table) {
    .ant-table-tbody > tr > td {
        padding: 12px 16px;
    }

    .ant-table-thead > tr > th {
        padding: 12px 16px;
        font-weight: 600;
    }
}

.mb-8-2 {
    padding-bottom: 16px;
    /* removed border-bottom for dark mode */
    margin-bottom: 16px;
}

.invocation-page {
    width: 100%;
    min-width: 0;
    overflow: hidden;
}

.invocation-card {
    min-width: 0;
    overflow: hidden;

    &::-webkit-scrollbar,
    :deep(*::-webkit-scrollbar) {
        width: 8px;
        height: 8px;
        background: transparent;
    }
    &::-webkit-scrollbar-thumb,
    :deep(*::-webkit-scrollbar-thumb) {
        background: rgba(0, 0, 0, 0.15);
        border-radius: 10em;
    }

    :global(html[data-theme='dark']) & {
        &::-webkit-scrollbar-thumb,
        :deep(*::-webkit-scrollbar-thumb) {
            background: rgba(255, 255, 255, 0.2);
        }
    }
}
</style>
