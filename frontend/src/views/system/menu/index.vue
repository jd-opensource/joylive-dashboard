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
                            {{ $t('pages.system.menu.add') }}
                        </a-button>
                    </a-col>
                    <a-col flex="auto"></a-col>
                    <a-col flex="none">
                        <a-form
                            :model="searchFormData"
                            layout="inline">
                            <a-form-item
                                :label="$t('pages.system.menu.form.name')"
                                name="name"
                                style="margin-bottom: 0">
                                <a-input
                                    :placeholder="$t('pages.system.menu.form.name.placeholder')"
                                    v-model:value="searchFormData.name"
                                    style="width: 200px"
                                    @pressEnter="handleSearch"></a-input>
                            </a-form-item>
                            <a-form-item
                                name="code"
                                style="margin-bottom: 0">
                                <template #label>
                                    {{ $t('pages.system.menu.form.code') }}
                                    <a-tooltip :title="$t('pages.system.menu.form.code')">
                                        <question-circle-outlined class="ml-4-1 color-placeholder" />
                                    </a-tooltip>
                                </template>
                                <a-input
                                    :placeholder="$t('pages.system.menu.form.code.placeholder')"
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
                    rowKey="id"
                    :loading="loading"
                    :pagination="true"
                    :columns="columns"
                    :data-source="listData"
                    :scroll="{ x: 1200 }">
                    <template #bodyCell="{ column, record }">
                        <template v-if="'menuType' === column.key">
                            <!--菜单-->
                            <a-tag
                                v-if="menuTypeEnum.is('page', record.type)"
                                color="processing">
                                {{ menuTypeEnum.getDesc(record.type) }}
                            </a-tag>
                            <!--按钮-->
                            <a-tag
                                v-if="menuTypeEnum.is('button', record.type)"
                                color="success">
                                {{ menuTypeEnum.getDesc(record.type) }}
                            </a-tag>
                        </template>

                        <template v-if="'created_at' === column.key">
                            {{ formatUtcDateTime(record.created_at) }}
                        </template>

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

                        <template v-if="'action' === column.key">
                            <x-action-button @click="$refs.editDialogRef.handleEdit(record)">
                                <a-tooltip>
                                    <template #title>{{ $t('pages.system.menu.edit') }}</template>
                                    <edit-outlined />
                                </a-tooltip>
                            </x-action-button>

                            <x-action-button @click="$refs.editDialogRef.handleCreateChild(record)">
                                <a-tooltip>
                                    <template #title>{{ $t('pages.system.menu.button.addChild') }}</template>
                                    <plus-circle-outlined />
                                </a-tooltip>
                            </x-action-button>
                            <x-action-button @click="handleDelete(record)">
                                <a-tooltip>
                                    <template #title>{{ $t('pages.system.delete') }}</template>
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
        @ok="onOk"
        ref="editDialogRef" />
</template>

<script setup>
import { Modal, message } from 'ant-design-vue'
import { ref } from 'vue'
import {
    PlusOutlined,
    EditOutlined,
    DeleteOutlined,
    PlusCircleOutlined,
    SearchOutlined,
    RedoOutlined,
    QuestionCircleOutlined,
} from '@ant-design/icons-vue'
import apis from '@/apis'
import { config } from '@/config'
import { menuTypeEnum, statusTypeEnum } from '@/enums/system'
import { usePagination } from '@/hooks'
import { formatUtcDateTime } from '@/utils/util'
import EditDialog from './components/EditDialog.vue'
import { useI18n } from 'vue-i18n'
defineOptions({
    // eslint-disable-next-line vue/no-reserved-component-names
    name: 'menu',
})
const { t } = useI18n() // 解构出t方法
const columns = ref([
    { title: t('pages.system.menu.form.name'), dataIndex: 'name', key: 'name', fixed: true, width: 200 },
    { title: t('pages.system.menu.form.code'), dataIndex: 'code', key: 'code', width: 200 },
    { title: t('pages.system.menu.form.type'), dataIndex: 'type', key: 'menuType', width: 100 },
    { title: t('pages.system.menu.form.status'), dataIndex: 'status', key: 'statusType', width: 100 },
    { title: t('pages.system.menu.form.sequence'), dataIndex: 'sequence', width: 100 },
    { title: t('pages.system.menu.form.created_at'), dataIndex: 'created_at', key: 'created_at', width: 180 },
    { title: t('button.action'), key: 'action', fixed: 'right', width: 150 },
])
const { listData, loading, showLoading, hideLoading, searchFormData, paginationState, resetPagination } =
    usePagination()
const editDialogRef = ref()

getMenuList()

/**
 * 获取菜单列表
 * @return {Promise<void>}
 */
async function getMenuList() {
    try {
        showLoading()
        // const { current } = paginationState
        const { data, success, total } = await apis.menu
            .getMenuList({
                ...searchFormData.value,
            })
            .catch(() => {
                throw new Error()
            })
        hideLoading()
        if (config('http.code.success') === success) {
            data.forEach((item) => {
                item.name = t(item.code) || item.name
            })
            listData.value = data
            paginationState.total = total
        }
    } catch (error) {
        hideLoading()
    }
}
/**
 * 搜索
 */
function handleSearch() {
    resetPagination()
    getMenuList()
}
/**
 * 重置
 */
function handleResetSearch() {
    searchFormData.value = {}
    resetPagination()
    getMenuList()
}
/**
 * 删除
 * @param id
 */
function handleDelete({ id }) {
    Modal.confirm({
        title: t('pages.system.menu.delTip'),
        content: t('button.confirm'),
        okText: t('button.confirm'),
        okType: 'danger',
        onOk: () => {
            return new Promise((resolve, reject) => {
                ;(async () => {
                    try {
                        const { success } = await apis.menu.delMenu(id).catch(() => {
                            throw new Error()
                        })
                        if (config('http.code.success') === success) {
                            resolve()
                            message.success(t('component.message.success.delete'))
                            await getMenuList()
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
 * 编辑完成
 */
async function onOk() {
    message.success(t('component.message.success.delete'))
    await getMenuList()
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
