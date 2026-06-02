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
                            {{ $t('pages.system.role.add') }}
                        </a-button>
                    </a-col>
                    <a-col flex="auto"></a-col>
                    <a-col flex="none">
                        <a-form
                            :model="searchFormData"
                            layout="inline">
                            <a-form-item
                                :label="$t('pages.system.role.form.name')"
                                name="name"
                                style="margin-bottom: 0">
                                <a-input
                                    :placeholder="$t('pages.system.role.form.code.placeholder')"
                                    v-model:value="searchFormData.name"
                                    style="width: 200px"
                                    @pressEnter="handleSearch"></a-input>
                            </a-form-item>
                            <a-form-item
                                name="code"
                                style="margin-bottom: 0">
                                <template #label>
                                    {{ $t('pages.system.role.form.code') }}
                                    <a-tooltip :title="$t('pages.system.role.form.code')">
                                        <question-circle-outlined class="ml-4-1 color-placeholder" />
                                    </a-tooltip>
                                </template>
                                <a-input
                                    :placeholder="$t('pages.system.role.form.code.placeholder')"
                                    v-model:value="searchFormData.code"
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
                    :scroll="{ x: 1000 }"
                    @change="onTableChange">
                    <template #bodyCell="{ column, record }">
                        <template v-if="'statusType' === column.key">
                            <!--状态-->
                            <a-tag
                                v-if="statusTypeEnum.is('enabled', record.status)"
                                color="green">
                                {{ statusTypeEnum.getDesc(record.status) }}
                            </a-tag>
                            <!--状态-->
                            <a-tag
                                v-if="statusTypeEnum.is('disabled', record.status)"
                                color="default">
                                {{ statusTypeEnum.getDesc(record.status) }}
                            </a-tag>
                        </template>

                        <template v-if="'created_at' === column.key">
                            {{ formatUtcDateTime(record.created_at) }}
                        </template>

                        <template v-if="'action' === column.key">
                            <x-action-button @click="$refs.editDialogRef.handleEdit(record)">
                                <a-tooltip>
                                    <template #title> {{ $t('pages.system.role.edit') }}</template>
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
        @ok="onOk"></edit-dialog>
</template>

<script setup>
import { message, Modal } from 'ant-design-vue'
import { ref } from 'vue'
import apis from '@/apis'
import { formatUtcDateTime } from '@/utils/util'
import { config } from '@/config'
import { statusTypeEnum } from '@/enums/system'
import { usePagination } from '@/hooks'
import EditDialog from './components/EditDialog.vue'
import {
    PlusOutlined,
    EditOutlined,
    DeleteOutlined,
    SearchOutlined,
    RedoOutlined,
    QuestionCircleOutlined,
} from '@ant-design/icons-vue'
import { useI18n } from 'vue-i18n'

defineOptions({
    name: 'systemRole',
})
const { t } = useI18n() // 解构出t方法
const columns = [
    { title: t('pages.system.role.form.code'), dataIndex: 'code', width: 240 },
    { title: t('pages.system.role.form.name'), dataIndex: 'name' },
    { title: t('pages.system.role.form.status'), dataIndex: 'status', key: 'statusType', width: 100 },
    { title: t('pages.system.role.form.sequence'), dataIndex: 'sequence', width: 100 },
    { title: t('pages.system.role.form.created_at'), key: 'created_at', fixed: 'right', width: 180 },
    { title: t('button.action'), key: 'action', fixed: 'right', width: 120 },
]

const { listData, loading, showLoading, hideLoading, paginationState, searchFormData, resetPagination } =
    usePagination()
const editDialogRef = ref()

getPageList()

/**
 * 获取用户列表
 * @returns {Promise<void>}
 */
async function getPageList() {
    try {
        showLoading()
        const { pageSize, current } = paginationState
        const { success, data, total } = await apis.role
            .getRoleList({
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

/**
 * 移除
 */
function handleRemove({ id }) {
    Modal.confirm({
        title: t('pages.system.role.delTip'),
        content: t('button.confirm'),
        okText: t('button.confirm'),
        okType: 'danger',
        onOk: () => {
            return new Promise((resolve, reject) => {
                ;(async () => {
                    try {
                        const { success } = await apis.role.delRole(id).catch(() => {
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

/**
 * 分页
 */
function onTableChange({ current, pageSize }) {
    paginationState.current = current
    paginationState.pageSize = pageSize
    getPageList()
}

/**
 * 重置
 */
function handleResetSearch() {
    searchFormData.value = {}
    resetPagination()
    getPageList()
}

/**
 * 搜索
 */
function handleSearch() {
    resetPagination()
    getPageList()
}

/**
 * 编辑完成
 */
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
    /* removed border-bottom for dark mode */
    margin-bottom: 16px;
}

// 暗黑模式适配
html[data-theme='dark'],
:root[data-theme='dark'],
body[data-theme='dark'] {
    :deep(.x-action-button) {
        &:hover {
            background-color: rgba(255, 255, 255, 0.08);
        }
    }
}
</style>
