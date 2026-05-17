<template>
    <div class="service-detail">
        <!-- 基本信息 -->
        <a-card
            :title="$t('pages.service.detail.basicInfo')"
            class="info-card"
            :bordered="false">
            <a-card-grid style="width: 33.33%; text-align: center">
                <div class="info-item">
                    <span class="info-label">{{ $t('pages.service.detail.instanceCount') }}</span>
                    <span class="info-value">{{ serviceData.instanceCount || 0 }}</span>
                </div>
            </a-card-grid>
            <a-card-grid style="width: 33.33%; text-align: center">
                <div class="info-item">
                    <span class="info-label">{{ $t('pages.service.detail.belongSpace') }}</span>
                    <span class="info-value">{{
                        serviceData.space
                            ? `${serviceData.space.name}（${serviceData.space.code}）`
                            : serviceData.space_code || '--'
                    }}</span>
                </div>
            </a-card-grid>
            <a-card-grid style="width: 33.33%; text-align: center">
                <div class="info-item">
                    <span class="info-label">{{ $t('pages.service.detail.belongApplication') }}</span>
                    <span class="info-value">{{ serviceData.application_name || '--' }}</span>
                </div>
            </a-card-grid>
            <a-card-grid style="width: 33.33%; text-align: center">
                <div class="info-item">
                    <span class="info-label">{{ $t('pages.service.detail.belongService') }}</span>
                    <span class="info-value">{{ serviceData.name || '--' }}</span>
                </div>
            </a-card-grid>
            <a-card-grid style="width: 33.33%; text-align: center">
                <div class="info-item">
                    <span class="info-label">{{ $t('pages.service.detail.registrationType') }}</span>
                    <span class="info-value">{{
                        registrationTypeMap[serviceData.registration_type] || serviceData.registration_type || '--'
                    }}</span>
                </div>
            </a-card-grid>
            <a-card-grid style="width: 33.33%; text-align: center">
                <div class="info-item">
                    <span class="info-label">{{ $t('pages.service.detail.createdAt') }}</span>
                    <span class="info-value">{{ formatUtcDateTime(serviceData.created_at) || '--' }}</span>
                </div>
            </a-card-grid>
        </a-card>

        <!-- Tab 区域 -->
        <a-card
            class="detail-card"
            :bordered="false">
            <a-tabs
                v-model:activeKey="activeTab"
                class="detail-tabs">
                <a-tab-pane
                    key="instance"
                    :tab="$t('pages.service.detail.tab.instance')" />
                <a-tab-pane
                    key="group"
                    :tab="$t('pages.service.detail.tab.group')" />
                <a-tab-pane
                    key="alias"
                    :tab="$t('pages.service.detail.tab.alias')" />
                <a-tab-pane
                    key="provider"
                    :tab="$t('pages.service.detail.tab.provider')" />
                <a-tab-pane
                    key="consumer"
                    :tab="$t('pages.service.detail.tab.consumer')" />
                <!-- <a-tab-pane key="monitor" :tab="$t('pages.service.detail.tab.monitor')" /> -->
                <a-tab-pane
                    key="api"
                    :tab="$t('pages.service.detail.tab.api')" />
                <a-tab-pane
                    key="member"
                    :tab="$t('pages.service.detail.tab.member')" />
            </a-tabs>

            <!-- 服务分组 Tab 内容 -->
            <div v-if="activeTab === 'group'">
                <div class="group-toolbar">
                    <a-button
                        type="primary"
                        @click="$refs.groupEditRef.handleCreate()">
                        {{ $t('pages.service.group.create') }}
                    </a-button>
                    <div class="group-toolbar-right">
                        <a-input-search
                            v-model:value="groupSearchName"
                            :placeholder="$t('pages.service.group.search.placeholder')"
                            style="width: 200px"
                            allow-clear
                            @search="loadGroupList"
                            @pressEnter="loadGroupList" />
                        <a-button @click="loadGroupList">
                            <template #icon><reload-outlined /></template>
                        </a-button>
                    </div>
                </div>
                <a-table
                    :columns="groupColumns"
                    :data-source="groupListData"
                    :loading="groupLoading"
                    :pagination="groupPagination"
                    @change="onGroupTableChange">
                    <template #bodyCell="{ column, record }">
                        <template v-if="'labels' === column.key">
                            {{ record.labels || '--' }}
                        </template>
                        <template v-if="'updated_at' === column.key">
                            {{ formatUtcDateTime(record.updated_at) }}
                        </template>
                        <template v-if="'action' === column.key">
                            <template v-if="record.id !== 'default'">
                                <x-action-button @click="$refs.groupEditRef.handleEdit(record)">
                                    <a-tooltip>
                                        <template #title> {{ $t('pages.service.group.edit') }}</template>
                                        <edit-outlined />
                                    </a-tooltip>
                                </x-action-button>
                                <x-action-button @click="handleRemoveGroup(record)">
                                    <a-tooltip>
                                        <template #title> {{ $t('pages.system.delete') }}</template>
                                        <delete-outlined style="color: #ff4d4f" />
                                    </a-tooltip>
                                </x-action-button>
                            </template>
                        </template>
                    </template>
                </a-table>
            </div>

            <!-- 服务别名 Tab 内容 -->
            <div v-else-if="activeTab === 'alias'">
                <div class="group-toolbar">
                    <a-button
                        type="primary"
                        @click="$refs.aliasEditRef.handleCreate()">
                        {{ $t('pages.service.alias.create') }}
                    </a-button>
                    <div class="group-toolbar-right">
                        <a-input-search
                            v-model:value="aliasSearchName"
                            :placeholder="$t('pages.service.alias.search.placeholder')"
                            style="width: 200px"
                            allow-clear
                            @search="loadAliasList"
                            @pressEnter="loadAliasList" />
                        <a-button @click="loadAliasList">
                            <template #icon><reload-outlined /></template>
                        </a-button>
                    </div>
                </div>
                <a-table
                    :columns="aliasColumns"
                    :data-source="aliasListData"
                    :loading="aliasLoading"
                    :pagination="aliasPagination"
                    @change="onAliasTableChange">
                    <template #bodyCell="{ column, record }">
                        <template v-if="'updated_at' === column.key">
                            {{ formatUtcDateTime(record.updated_at) }}
                        </template>
                        <template v-if="'action' === column.key">
                            <x-action-button @click="$refs.aliasEditRef.handleEdit(record)">
                                <a-tooltip>
                                    <template #title> {{ $t('pages.service.alias.edit') }}</template>
                                    <edit-outlined />
                                </a-tooltip>
                            </x-action-button>
                            <x-action-button @click="handleRemoveAlias(record)">
                                <a-tooltip>
                                    <template #title> {{ $t('pages.system.delete') }}</template>
                                    <delete-outlined style="color: #ff4d4f" />
                                </a-tooltip>
                            </x-action-button>
                        </template>
                    </template>
                </a-table>
            </div>

            <!-- 成员管理 Tab 内容 -->
            <div v-else-if="activeTab === 'member'">
                <div class="group-toolbar">
                    <a-button
                        type="primary"
                        @click="$refs.memberEditRef.handleCreate()">
                        {{ $t('pages.member.add') }}
                    </a-button>
                    <div class="group-toolbar-right">
                        <a-input-search
                            v-model:value="memberSearchUser"
                            :placeholder="$t('pages.member.search.placeholder')"
                            style="width: 200px"
                            allow-clear
                            @search="loadMemberList"
                            @pressEnter="loadMemberList" />
                        <a-button @click="loadMemberList">
                            <template #icon><reload-outlined /></template>
                        </a-button>
                    </div>
                </div>
                <a-table
                    :columns="memberColumns"
                    :data-source="memberListData"
                    :loading="memberLoading"
                    :pagination="memberPagination"
                    @change="onMemberTableChange">
                    <template #bodyCell="{ column, record }">
                        <template v-if="'permission' === column.key">
                            <a-tag
                                v-if="record.permission & 1"
                                color="green"
                                >{{ $t('pages.member.form.permission.read') }}</a-tag
                            >
                            <a-tag
                                v-if="record.permission & 2"
                                color="blue"
                                >{{ $t('pages.member.form.permission.write') }}</a-tag
                            >
                            <a-tag
                                v-if="record.permission & 4"
                                color="red"
                                >{{ $t('pages.member.form.permission.delete') }}</a-tag
                            >
                        </template>
                        <template v-if="'created_at' === column.key">
                            {{ formatUtcDateTime(record.created_at) }}
                        </template>
                        <template v-if="'action' === column.key">
                            <x-action-button @click="$refs.memberEditRef.handleEdit(record)">
                                <a-tooltip>
                                    <template #title> {{ $t('pages.member.edit') }}</template>
                                    <edit-outlined />
                                </a-tooltip>
                            </x-action-button>
                            <x-action-button @click="handleRemoveMember(record)">
                                <a-tooltip>
                                    <template #title> {{ $t('pages.system.delete') }}</template>
                                    <delete-outlined style="color: #ff4d4f" />
                                </a-tooltip>
                            </x-action-button>
                        </template>
                    </template>
                </a-table>
            </div>

            <!-- 调用方 Tab 内容 -->
            <div v-else-if="activeTab === 'consumer'">
                <div class="group-toolbar">
                    <div class="group-toolbar-right">
                        <a-button @click="loadConsumerList">
                            <template #icon><reload-outlined /></template>
                        </a-button>
                    </div>
                </div>
                <a-table
                    :columns="consumerColumns"
                    :data-source="consumerListData"
                    :loading="consumerLoading"
                    :pagination="consumerPagination"
                    @change="onConsumerTableChange">
                    <template #bodyCell="{ column, record }">
                        <template v-if="'application_name' === column.key">
                            {{ record.application?.name || record.application_id }}
                        </template>
                        <template v-if="'status' === column.key">
                            <a-tag :color="statusColorMap[record.status]">
                                {{ statusTextMap[record.status] || record.status }}
                            </a-tag>
                        </template>
                        <template v-if="'creator' === column.key">
                            {{ record.application?.creator || '-' }}
                        </template>
                        <template v-if="'created_at' === column.key">
                            {{ formatUtcDateTime(record.created_at) }}
                        </template>
                        <template v-if="'action' === column.key">
                            <x-action-button
                                v-if="record.status !== 'approved'"
                                @click="handleApproveConsumer(record)">
                                <a-tooltip>
                                    <template #title> {{ $t('pages.service.consumer.approve') }}</template>
                                    <check-circle-outlined style="color: #52c41a" />
                                </a-tooltip>
                            </x-action-button>
                            <x-action-button
                                v-if="record.status !== 'rejected'"
                                @click="handleRejectConsumer(record)">
                                <a-tooltip>
                                    <template #title> {{ $t('pages.service.consumer.reject') }}</template>
                                    <close-circle-outlined style="color: #ff4d4f" />
                                </a-tooltip>
                            </x-action-button>
                        </template>
                    </template>
                </a-table>
            </div>

            <!-- 提供方 Tab 内容 -->
            <div v-else-if="activeTab === 'provider'">
                <div class="group-toolbar">
                    <div class="group-toolbar-right">
                        <a-button @click="loadProviderList">
                            <template #icon><reload-outlined /></template>
                        </a-button>
                    </div>
                </div>
                <a-table
                    :columns="providerColumns"
                    :data-source="providerListData"
                    :loading="providerLoading"
                    :pagination="providerPagination"
                    @change="onProviderTableChange">
                    <template #bodyCell="{ column, record }">
                        <template v-if="'application_name' === column.key">
                            {{ record.application?.name || record.application_id }}
                        </template>
                        <template v-if="'language' === column.key">
                            {{ record.application?.language || '-' }}
                        </template>
                        <template v-if="'creator' === column.key">
                            {{ record.application?.creator || '-' }}
                        </template>
                        <template v-if="'created_at' === column.key">
                            {{ formatUtcDateTime(record.created_at) }}
                        </template>
                    </template>
                </a-table>
            </div>

            <!-- 其他 Tab 占位 -->
            <div
                v-else
                class="tab-placeholder">
                <a-empty />
            </div>
        </a-card>

        <!-- 分组编辑弹窗 -->
        <group-edit-dialog
            ref="groupEditRef"
            :service-id="serviceId"
            @ok="loadGroupList" />

        <!-- 别名编辑弹窗 -->
        <alias-edit-dialog
            ref="aliasEditRef"
            :service-id="serviceId"
            :space-code="serviceData.space?.code || serviceData.space_code"
            @ok="loadAliasList" />

        <!-- 成员编辑弹窗 -->
        <member-edit-dialog
            ref="memberEditRef"
            :service-id="serviceId"
            @ok="loadMemberList" />
    </div>
</template>

<script setup>
import { ref, onMounted, reactive } from 'vue'
import { useRoute } from 'vue-router'
import { message, Modal } from 'ant-design-vue'
import {
    ReloadOutlined,
    EditOutlined,
    DeleteOutlined,
    CheckCircleOutlined,
    CloseCircleOutlined,
} from '@ant-design/icons-vue'
import apis from '@/apis'
import { config } from '@/config'
import { formatUtcDateTime } from '@/utils/util'
import { useI18n } from 'vue-i18n'
import GroupEditDialog from './ServiceGroupEditDialog.vue'
import AliasEditDialog from './ServiceAliasEditDialog.vue'
import MemberEditDialog from './MemberEditDialog.vue'

defineOptions({
    name: 'serviceDetail',
})

const route = useRoute()
const { t } = useI18n()
const serviceId = ref(route.params.id)
const serviceData = ref({})
const activeTab = ref('instance')
const groupSearchName = ref('')

const registrationTypeMap = {
    HTTP: t('pages.service.form.registration_type.http'),
    RPC_APP: t('pages.service.form.registration_type.rpc_app'),
    RPC_INTERFACE: t('pages.service.form.registration_type.rpc_interface'),
}

// 服务分组
const groupListData = ref([])
const groupLoading = ref(false)
const groupPagination = reactive({
    current: 1,
    pageSize: 10,
    total: 0,
    showSizeChanger: true,
    showTotal: (total) => `共 ${total} 条`,
})

const groupColumns = [
    {
        title: t('pages.service.group.form.name'),
        dataIndex: 'name',
        width: 150,
    },
    {
        title: t('pages.service.group.form.code'),
        dataIndex: 'code',
        width: 150,
    },
    {
        title: t('pages.service.group.form.labels'),
        dataIndex: 'labels',
        key: 'labels',
        width: 120,
    },
    {
        title: t('pages.service.group.form.description'),
        dataIndex: 'description',
        ellipsis: true,
    },
    {
        title: t('pages.service.group.form.updated_at'),
        key: 'updated_at',
        width: 180,
    },
    {
        title: t('button.action'),
        key: 'action',
        width: 120,
    },
]

// 服务别名
const aliasSearchName = ref('')
const aliasListData = ref([])
const aliasLoading = ref(false)
const aliasPagination = reactive({
    current: 1,
    pageSize: 10,
    total: 0,
    showSizeChanger: true,
    showTotal: (total) => `共 ${total} 条`,
})

const aliasColumns = [
    {
        title: t('pages.service.alias.form.alias'),
        dataIndex: 'alias',
        width: 150,
    },
    {
        title: t('pages.service.alias.form.description'),
        dataIndex: 'description',
        ellipsis: true,
    },
    {
        title: t('pages.service.alias.form.updated_at'),
        key: 'updated_at',
        width: 180,
    },
    {
        title: t('button.action'),
        key: 'action',
        width: 120,
    },
]

// 成员管理
const memberSearchUser = ref('')
const memberListData = ref([])
const memberLoading = ref(false)
const memberPagination = reactive({
    current: 1,
    pageSize: 10,
    total: 0,
    showSizeChanger: true,
    showTotal: (total) => `共 ${total} 条`,
})

const memberColumns = [
    {
        title: t('pages.member.form.user'),
        dataIndex: 'user',
        width: 150,
    },
    {
        title: t('pages.member.form.tenant'),
        dataIndex: 'tenant',
        width: 150,
    },
    {
        title: t('pages.member.form.role'),
        dataIndex: 'role',
        width: 100,
    },
    {
        title: t('pages.member.form.permission'),
        key: 'permission',
        width: 200,
    },
    {
        title: t('pages.member.form.created_at'),
        key: 'created_at',
        width: 180,
    },
    {
        title: t('button.action'),
        key: 'action',
        width: 120,
    },
]

// 调用方
const consumerListData = ref([])
const consumerLoading = ref(false)
const consumerPagination = reactive({
    current: 1,
    pageSize: 10,
    total: 0,
    showSizeChanger: true,
    showTotal: (total) => `共 ${total} 条`,
})

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

const consumerColumns = [
    {
        title: t('pages.service.form.application_name'),
        key: 'application_name',
        width: 200,
    },
    {
        title: t('pages.service.form.application_service_status'),
        key: 'status',
        width: 100,
    },
    {
        title: t('pages.application.form.creator'),
        key: 'creator',
        width: 120,
    },
    {
        title: t('pages.member.form.created_at'),
        key: 'created_at',
        width: 180,
    },
    {
        title: t('button.action'),
        key: 'action',
        width: 120,
    },
]

// 提供方
const providerListData = ref([])
const providerLoading = ref(false)
const providerPagination = reactive({
    current: 1,
    pageSize: 10,
    total: 0,
    showSizeChanger: true,
    showTotal: (total) => `共 ${total} 条`,
})

const providerColumns = [
    {
        title: t('pages.service.form.application_name'),
        key: 'application_name',
        width: 200,
    },
    {
        title: t('pages.application.form.language'),
        key: 'language',
        width: 120,
    },
    {
        title: t('pages.application.form.creator'),
        key: 'creator',
        width: 120,
    },
    {
        title: t('pages.member.form.created_at'),
        key: 'created_at',
        width: 180,
    },
]

onMounted(() => {
    loadServiceDetail()
    loadGroupList()
    loadAliasList()
    loadMemberList()
    loadConsumerList()
    loadProviderList()
})

async function loadServiceDetail() {
    try {
        const { data, success } = await apis.service.getService(serviceId.value)
        if (success) {
            serviceData.value = data || {}
        }
    } catch (error) {
        // ignore
    }
}

async function loadGroupList() {
    try {
        groupLoading.value = true
        const { data, success, total } = await apis.service
            .getServiceGroupList({
                pageSize: groupPagination.pageSize,
                current: groupPagination.current,
                service_id: serviceId.value,
                name: groupSearchName.value || undefined,
            })
            .catch(() => {
                throw new Error()
            })
        groupLoading.value = false
        if (config('http.code.success') === success) {
            groupListData.value = data || []
            groupPagination.total = total || 0
        }
    } catch (error) {
        groupLoading.value = false
    }
}

function onGroupTableChange({ current, pageSize }) {
    groupPagination.current = current
    groupPagination.pageSize = pageSize
    loadGroupList()
}

function handleRemoveGroup({ id }) {
    Modal.confirm({
        title: t('pages.service.group.delTip'),
        okText: t('button.confirm'),
        okType: 'danger',
        onOk: () => {
            return new Promise((resolve, reject) => {
                ;(async () => {
                    try {
                        const { success } = await apis.service.delServiceGroup(id).catch(() => {
                            throw new Error()
                        })
                        if (config('http.code.success') === success) {
                            resolve()
                            message.success(t('component.message.success.delete'))
                            await loadGroupList()
                        }
                    } catch (error) {
                        reject()
                    }
                })()
            })
        },
    })
}

async function loadAliasList() {
    try {
        aliasLoading.value = true
        const { data, success, total } = await apis.service
            .getServiceAliasList({
                pageSize: aliasPagination.pageSize,
                current: aliasPagination.current,
                service_id: serviceId.value,
                alias: aliasSearchName.value || undefined,
            })
            .catch(() => {
                throw new Error()
            })
        aliasLoading.value = false
        if (config('http.code.success') === success) {
            aliasListData.value = data || []
            aliasPagination.total = total || 0
        }
    } catch (error) {
        aliasLoading.value = false
    }
}

function onAliasTableChange({ current, pageSize }) {
    aliasPagination.current = current
    aliasPagination.pageSize = pageSize
    loadAliasList()
}

function handleRemoveAlias({ id }) {
    Modal.confirm({
        title: t('pages.service.alias.delTip'),
        okText: t('button.confirm'),
        okType: 'danger',
        onOk: () => {
            return new Promise((resolve, reject) => {
                ;(async () => {
                    try {
                        const { success } = await apis.service.delServiceAlias(id).catch(() => {
                            throw new Error()
                        })
                        if (config('http.code.success') === success) {
                            resolve()
                            message.success(t('component.message.success.delete'))
                            await loadAliasList()
                        }
                    } catch (error) {
                        reject()
                    }
                })()
            })
        },
    })
}

async function loadMemberList() {
    try {
        memberLoading.value = true
        const { data, success, total } = await apis.service
            .getDataPermissionList({
                pageSize: memberPagination.pageSize,
                current: memberPagination.current,
                type: 'service',
                data_id: serviceId.value,
                user: memberSearchUser.value || undefined,
            })
            .catch(() => {
                throw new Error()
            })
        memberLoading.value = false
        if (config('http.code.success') === success) {
            memberListData.value = data || []
            memberPagination.total = total || 0
        }
    } catch (error) {
        memberLoading.value = false
    }
}

function onMemberTableChange({ current, pageSize }) {
    memberPagination.current = current
    memberPagination.pageSize = pageSize
    loadMemberList()
}

function handleRemoveMember({ id }) {
    Modal.confirm({
        title: t('pages.member.delTip'),
        okText: t('button.confirm'),
        okType: 'danger',
        onOk: () => {
            return new Promise((resolve, reject) => {
                ;(async () => {
                    try {
                        const { success } = await apis.service.delDataPermission(id).catch(() => {
                            throw new Error()
                        })
                        if (config('http.code.success') === success) {
                            resolve()
                            message.success(t('component.message.success.delete'))
                            await loadMemberList()
                        }
                    } catch (error) {
                        reject()
                    }
                })()
            })
        },
    })
}

// 调用方
async function loadConsumerList() {
    try {
        consumerLoading.value = true
        const { data, success, total } = await apis.application_service
            .getApplicationServiceList({
                pageSize: consumerPagination.pageSize,
                current: consumerPagination.current,
                service_id: serviceId.value,
                role: 'consumer',
            })
            .catch(() => {
                throw new Error()
            })
        consumerLoading.value = false
        if (config('http.code.success') === success) {
            consumerListData.value = data || []
            consumerPagination.total = total || 0
        }
    } catch (error) {
        consumerLoading.value = false
    }
}

function onConsumerTableChange({ current, pageSize }) {
    consumerPagination.current = current
    consumerPagination.pageSize = pageSize
    loadConsumerList()
}

function handleApproveConsumer({ id }) {
    Modal.confirm({
        title: t('pages.service.consumer.approveTip'),
        okText: t('button.confirm'),
        onOk: () => {
            return new Promise((resolve, reject) => {
                ;(async () => {
                    try {
                        const { success } = await apis.application_service
                            .updateApplicationServiceStatus(id, { status: 'approved' })
                            .catch(() => {
                                throw new Error()
                            })
                        if (config('http.code.success') === success) {
                            resolve()
                            message.success(t('component.message.success.save'))
                            await loadConsumerList()
                        }
                    } catch (error) {
                        reject()
                    }
                })()
            })
        },
    })
}

function handleRejectConsumer({ id }) {
    Modal.confirm({
        title: t('pages.service.consumer.rejectTip'),
        okText: t('button.confirm'),
        okType: 'danger',
        onOk: () => {
            return new Promise((resolve, reject) => {
                ;(async () => {
                    try {
                        const { success } = await apis.application_service
                            .updateApplicationServiceStatus(id, { status: 'rejected' })
                            .catch(() => {
                                throw new Error()
                            })
                        if (config('http.code.success') === success) {
                            resolve()
                            message.success(t('component.message.success.save'))
                            await loadConsumerList()
                        }
                    } catch (error) {
                        reject()
                    }
                })()
            })
        },
    })
}

// 提供方
async function loadProviderList() {
    try {
        providerLoading.value = true
        const { data, success, total } = await apis.application_service
            .getApplicationServiceList({
                pageSize: providerPagination.pageSize,
                current: providerPagination.current,
                service_id: serviceId.value,
                role: 'provider',
            })
            .catch(() => {
                throw new Error()
            })
        providerLoading.value = false
        if (config('http.code.success') === success) {
            providerListData.value = data || []
            providerPagination.total = total || 0
        }
    } catch (error) {
        providerLoading.value = false
    }
}

function onProviderTableChange({ current, pageSize }) {
    providerPagination.current = current
    providerPagination.pageSize = pageSize
    loadProviderList()
}
</script>

<style lang="less" scoped>
.service-detail {
    padding: 0;
}

.info-card {
    margin-bottom: 16px;

    :deep(.ant-card-head-title) {
        font-size: 14px;
    }

    :deep(.ant-card-grid) {
        padding: 8px 16px;
    }

    .info-item {
        display: flex;
        flex-direction: column;
        align-items: center;
        gap: 2px;

        .info-label {
            color: #666;
            font-size: 13px;
        }

        .info-value {
            color: #333;
            font-size: 14px;
            font-weight: 500;
        }
    }
}

.detail-card {
    .detail-tabs {
        margin-bottom: 0;

        :deep(.ant-tabs-nav) {
            margin-bottom: 16px;
        }
    }
}

.group-toolbar {
    display: flex;
    justify-content: space-between;
    align-items: center;
    margin-bottom: 16px;
}

.group-toolbar-right {
    display: flex;
    align-items: center;
    gap: 8px;
}

.tab-placeholder {
    padding: 60px 0;
}
</style>
