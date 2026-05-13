<template>
    <x-search-bar class="mb-8-2">
        <template #default="{ gutter }">
            <a-form
                :label-col="{ style: { width: '80px' } }"
                :model="searchFormData"
                layout="inline">
                <a-row
                    :gutter="gutter"
                    style="width: 100%">
                    <a-col
                        :xs="24"
                        :sm="12"
                        :md="8"
                        :xl="6">
                        <a-form-item
                            :label="$t('pages.permission.form.space_code')"
                            name="space_code">
                            <a-select
                                :placeholder="$t('pages.permission.form.space_code.placeholder')"
                                v-model:value="searchFormData.space_code"
                                show-search
                                :filter-option="filterSpaceOption"
                                allow-clear>
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
                        :sm="12"
                        :md="8"
                        :xl="6">
                        <a-form-item
                            :label="$t('pages.permission.form.targetServiceId')"
                            name="target_service_id">
                            <a-select
                                :placeholder="$t('pages.permission.form.targetServiceId.placeholder')"
                                v-model:value="searchFormData.target_service_id"
                                show-search
                                :filter-option="filterServiceOption"
                                allow-clear>
                                <a-select-option
                                    v-for="item in serviceOptions"
                                    :key="item.id"
                                    :value="item.id">
                                    {{ item.name }}
                                </a-select-option>
                            </a-select>
                        </a-form-item>
                    </a-col>

                    <a-col
                        :xs="24"
                        :sm="12"
                        :md="8"
                        :xl="6">
                        <a-form-item
                            :label="$t('pages.permission.form.name')"
                            name="name">
                            <a-input
                                :placeholder="$t('pages.permission.form.name.placeholder')"
                                v-model:value="searchFormData.name"></a-input>
                        </a-form-item>
                    </a-col>

                    <a-col
                        :xs="24"
                        :sm="12"
                        :md="8"
                        :xl="6">
                        <a-form-item>
                            <a-space>
                                <a-button @click="handleResetSearch">{{ $t('button.reset') }}</a-button>
                                <a-button
                                    ghost
                                    type="primary"
                                    @click="handleSearch">
                                    {{ $t('button.search') }}
                                </a-button>
                            </a-space>
                        </a-form-item>
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
                        {{ $t('pages.permission.add') }}
                    </a-button>
                </x-action-bar>
                <a-table
                    :columns="columns"
                    :data-source="listData"
                    :loading="loading"
                    :pagination="paginationState"
                    :scroll="{ x: 1500 }"
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
                            {{ permissionTypeMap[record.type] || record.type }}
                        </template>
                        <template v-if="'enabled' === column.key">
                            <a-tag :color="record.enabled === 1 ? 'green' : 'default'">
                                {{
                                    record.enabled === 1
                                        ? $t('pages.permission.form.enabled.active')
                                        : $t('pages.permission.form.enabled.inactive')
                                }}
                            </a-tag>
                        </template>

                        <template v-if="'created_at' === column.key">
                            {{ formatUtcDateTime(record.created_at) }}
                        </template>

                        <template v-if="'action' === column.key">
                            <x-action-button @click="$refs.editDialogRef.handleEdit(record)">
                                <a-tooltip>
                                    <template #title> {{ $t('pages.permission.edit') }}</template>
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
import EditDialog from './PermissionEditDialog.vue'
import { PlusOutlined, EditOutlined, DeleteOutlined } from '@ant-design/icons-vue'
import { useI18n } from 'vue-i18n'

defineOptions({
    name: 'permissionList',
})
const { t } = useI18n()
const columns = [
    { title: t('pages.permission.form.name'), dataIndex: 'name', width: 180 },
    {
        title: t('pages.permission.form.sourceApplicationId'),
        dataIndex: 'source_application_id',
        key: 'source_application_id',
        width: 150,
    },
    {
        title: t('pages.permission.form.targetServiceId'),
        dataIndex: 'target_service_id',
        key: 'target_service_id',
        width: 150,
    },
    { title: t('pages.permission.form.path'), dataIndex: 'path', width: 150 },
    { title: t('pages.permission.form.method'), dataIndex: 'method', width: 100 },
    { title: t('pages.permission.form.relationType'), dataIndex: 'relation_type', key: 'relation_type', width: 100 },
    { title: t('pages.permission.form.type'), dataIndex: 'type', key: 'type', width: 100 },
    { title: t('pages.permission.form.enabled'), key: 'enabled', width: 80 },
    { title: t('pages.permission.form.creator'), dataIndex: 'creator', width: 100 },
    { title: t('pages.permission.form.description'), dataIndex: 'description', ellipsis: true },
    {
        title: t('pages.permission.form.created_at'),
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
const permissionTypeMap = {
    BLACK: '黑名单',
    WHITE: '白名单',
}

loadSpaceOptions()
loadServiceOptions()
loadApplicationOptions()
getPageList()

async function loadSpaceOptions() {
    try {
        const { success, data } = await apis.space.getSpaceList({ pageSize: 99, current: 1 }).catch(() => {
            throw new Error()
        })
        if (config('http.code.success') === success) {
            spaceOptions.value = data || []
        }
    } catch (error) {
        // ignore
    }
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
            .getPermissionList({
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
        title: t('pages.permission.delTip'),
        content: t('button.confirm'),
        okText: t('button.confirm'),
        onOk: () => {
            return new Promise((resolve, reject) => {
                ;(async () => {
                    try {
                        const { success } = await apis.policy.delPermission(id).catch(() => {
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
