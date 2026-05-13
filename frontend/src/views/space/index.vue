<template>
    <x-search-bar class="mb-8-2">
        <template #default="{ gutter, colSpan }">
            <a-form
                :label-col="{ style: { width: '100px' } }"
                :model="searchFormData"
                layout="inline">
                <a-row :gutter="gutter">
                    <a-col v-bind="colSpan">
                        <a-form-item
                            :label="$t('pages.space.form.name')"
                            name="name">
                            <a-input
                                :placeholder="$t('pages.space.form.name.placeholder')"
                                v-model:value="searchFormData.name"></a-input>
                        </a-form-item>
                    </a-col>

                    <a-col v-bind="colSpan">
                        <a-form-item
                            :label="$t('pages.space.form.code')"
                            name="code">
                            <a-input
                                :placeholder="$t('pages.space.form.code.placeholder')"
                                v-model:value="searchFormData.code"></a-input>
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
                        v-action="'add'"
                        type="primary"
                        @click="$refs.editDialogRef.handleCreate()">
                        <template #icon>
                            <plus-outlined></plus-outlined>
                        </template>
                        {{ $t('pages.space.add') }}
                    </a-button>
                </x-action-bar>
                <a-table
                    :columns="columns"
                    :data-source="listData"
                    :loading="loading"
                    :pagination="paginationState"
                    :scroll="{ x: 1000 }"
                    @change="onTableChange">
                    <template #bodyCell="{ column, record }">
                        <template v-if="'createAt' === column.key">
                            {{ formatUtcDateTime(record.created_at) }}
                        </template>

                        <template v-if="'action' === column.key">
                            <x-action-button @click="$refs.editDialogRef.handleEdit(record)">
                                <a-tooltip>
                                    <template #title> {{ $t('pages.space.edit') }}</template>
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
import { usePagination } from '@/hooks'
import EditDialog from './components/EditDialog.vue'
import { PlusOutlined, EditOutlined, DeleteOutlined } from '@ant-design/icons-vue'
import { useI18n } from 'vue-i18n'

defineOptions({
    name: 'spaceList',
})
const { t } = useI18n()
const columns = [
    { title: t('pages.space.form.code'), dataIndex: 'code', width: 200 },
    { title: t('pages.space.form.name'), dataIndex: 'name', width: 200 },
    { title: t('pages.space.form.tenant'), dataIndex: 'tenant', width: 150 },
    { title: t('pages.space.form.creator'), dataIndex: 'creator', width: 120 },
    { title: t('pages.space.form.description'), dataIndex: 'description', ellipsis: true },
    { title: t('pages.space.form.created_at'), key: 'createAt', fixed: 'right', width: 180 },
    { title: t('button.action'), key: 'action', fixed: 'right', width: 120 },
]

const { listData, loading, showLoading, hideLoading, paginationState, searchFormData, resetPagination } =
    usePagination()
const editDialogRef = ref()

getPageList()

async function getPageList() {
    try {
        showLoading()
        const { pageSize, current } = paginationState
        const { success, data, total } = await apis.space
            .getSpaceList({
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
        title: t('pages.space.delTip'),
        content: t('button.confirm'),
        okText: t('button.confirm'),
        onOk: () => {
            return new Promise((resolve, reject) => {
                ;(async () => {
                    try {
                        const { success } = await apis.space.delSpace(id).catch(() => {
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

<style lang="less" scoped></style>
