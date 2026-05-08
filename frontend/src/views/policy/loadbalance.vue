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
                            :label="$t('pages.loadbalance.form.spaceCode')"
                            name="space_code">
                            <a-select
                                :placeholder="$t('pages.loadbalance.form.spaceCode.placeholder')"
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
                            :label="$t('pages.loadbalance.form.name')"
                            name="name">
                            <a-input
                                :placeholder="$t('pages.loadbalance.form.name.placeholder')"
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
                        v-action="'add'"
                        type="primary"
                        @click="$refs.editDialogRef.handleCreate()">
                        <template #icon>
                            <plus-outlined></plus-outlined>
                        </template>
                        {{ $t('pages.loadbalance.add') }}
                    </a-button>
                </x-action-bar>
                <a-table
                    :columns="columns"
                    :data-source="listData"
                    :loading="loading"
                    :pagination="paginationState"
                    :scroll="{ x: 1400 }"
                    @change="onTableChange">
                    <template #bodyCell="{ column, record }">
                        <template v-if="'enabled' === column.key">
                            <a-tag :color="record.enabled === 1 ? 'green' : 'default'">
                                {{
                                    record.enabled === 1
                                        ? $t('pages.loadbalance.form.enabled.active')
                                        : $t('pages.loadbalance.form.enabled.inactive')
                                }}
                            </a-tag>
                        </template>

                        <template v-if="'createAt' === column.key">
                            {{ formatUtcDateTime(record.createdAt) }}
                        </template>

                        <template v-if="'action' === column.key">
                            <x-action-button @click="$refs.editDialogRef.handleEdit(record)">
                                <a-tooltip>
                                    <template #title> {{ $t('pages.loadbalance.edit') }}</template>
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
import { usePagination, useForm } from '@/hooks'
import EditDialog from './LoadbalanceEditDialog.vue'
import { PlusOutlined, EditOutlined, DeleteOutlined } from '@ant-design/icons-vue'
import { useI18n } from 'vue-i18n'

defineOptions({
    name: 'loadbalanceList',
})
const { t } = useI18n()
const columns = [
    { title: t('pages.loadbalance.form.name'), dataIndex: 'name', width: 180 },
    { title: t('pages.loadbalance.form.spaceCode'), dataIndex: 'spaceCode', width: 120 },
    { title: t('pages.loadbalance.form.targetServiceId'), dataIndex: 'targetServiceId', width: 120 },
    { title: t('pages.loadbalance.form.group'), dataIndex: 'group', width: 100 },
    { title: t('pages.loadbalance.form.policyType'), dataIndex: 'policyType', width: 120 },
    { title: t('pages.loadbalance.form.enabled'), key: 'enabled', width: 80 },
    { title: t('pages.loadbalance.form.creator'), dataIndex: 'creator', width: 100 },
    { title: t('pages.loadbalance.form.description'), dataIndex: 'description', ellipsis: true },
    { title: t('pages.loadbalance.form.createdAt'), key: 'createAt', fixed: 'right', width: 180 },
    { title: t('button.action'), key: 'action', fixed: 'right', width: 120 },
]

const { listData, loading, showLoading, hideLoading, paginationState, searchFormData, resetPagination } =
    usePagination()
const { resetForm } = useForm()
const editDialogRef = ref()
const spaceOptions = ref([])
const serviceOptions = ref([])
const applicationOptions = ref([])

const SPACE_CODE_KEY = 'loadbalance_space_code'

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

async function loadServiceOptions() {
    try {
        const { success, data } = await apis.service.getServiceList({ pageSize: 99, current: 1 }).catch(() => {
            throw new Error()
        })
        if (config('http.code.success') === success) {
            serviceOptions.value = data || []
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
        const { success, data, total } = await apis.policy
            .getLoadbalanceList({
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
        title: t('pages.loadbalance.delTip'),
        content: t('button.confirm'),
        okText: t('button.confirm'),
        onOk: () => {
            return new Promise((resolve, reject) => {
                ;(async () => {
                    try {
                        const { success } = await apis.policy.delLoadbalance(id).catch(() => {
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

<style lang="less" scoped></style>
