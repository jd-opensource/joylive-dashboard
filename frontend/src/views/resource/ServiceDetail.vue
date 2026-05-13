<template>
    <div class="service-detail">
        <!-- 基本信息 -->
        <a-card
            class="info-card"
            :bordered="false">
            <div class="info-title">{{ $t('pages.service.detail.basicInfo') }}</div>
            <div class="info-grid">
                <div class="info-item">
                    <span class="info-label">{{ $t('pages.service.detail.instanceCount') }}:</span>
                    <span class="info-value">{{ serviceData.instanceCount || 0 }}</span>
                </div>
                <div class="info-item">
                    <span class="info-label">{{ $t('pages.service.detail.belongSpace') }}:</span>
                    <span class="info-value">{{
                        serviceData.space
                            ? `${serviceData.space.name}（${serviceData.space.code}）`
                            : serviceData.space_code || '--'
                    }}</span>
                </div>
                <div class="info-item">
                    <span class="info-label">{{ $t('pages.service.detail.belongApplication') }}:</span>
                    <span class="info-value">{{ serviceData.application_name || '--' }}</span>
                </div>
                <div class="info-item">
                    <span class="info-label">{{ $t('pages.service.detail.belongService') }}:</span>
                    <span class="info-value">{{ serviceData.name || '--' }}</span>
                </div>
                <div class="info-item">
                    <span class="info-label">{{ $t('pages.service.detail.registrationType') }}:</span>
                    <span class="info-value">{{
                        registrationTypeMap[serviceData.registration_type] || serviceData.registration_type || '--'
                    }}</span>
                </div>
                <div class="info-item">
                    <span class="info-label">{{ $t('pages.service.detail.createdAt') }}:</span>
                    <span class="info-value">{{ formatUtcDateTime(serviceData.created_at) || '--' }}</span>
                </div>
            </div>
        </a-card>

        <!-- Tab 区域 -->
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
                key="interface"
                :tab="$t('pages.service.detail.tab.interface')" />
            <a-tab-pane
                key="consumer"
                :tab="$t('pages.service.detail.tab.consumer')" />
            <a-tab-pane
                key="monitor"
                :tab="$t('pages.service.detail.tab.monitor')" />
            <a-tab-pane
                key="alias"
                :tab="$t('pages.service.detail.tab.alias')" />
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
                    <template v-if="'updatedAt' === column.key">
                        {{ formatUtcDateTime(record.updated_at) }}
                    </template>
                    <template v-if="'action' === column.key">
                        <a-space>
                            <a @click="$refs.groupEditRef.handleEdit(record)">{{ $t('pages.service.group.edit') }}</a>
                            <a
                                style="color: #ff4d4f"
                                @click="handleRemoveGroup(record)"
                                >{{ $t('pages.system.delete') }}</a
                            >
                        </a-space>
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

        <!-- 分组编辑弹窗 -->
        <group-edit-dialog
            ref="groupEditRef"
            :service-id="serviceId"
            @ok="loadGroupList" />
    </div>
</template>

<script setup>
import { ref, onMounted, reactive } from 'vue'
import { useRoute } from 'vue-router'
import { message, Modal } from 'ant-design-vue'
import { ReloadOutlined } from '@ant-design/icons-vue'
import apis from '@/apis'
import { config } from '@/config'
import { formatUtcDateTime } from '@/utils/util'
import { useI18n } from 'vue-i18n'
import GroupEditDialog from './ServiceGroupEditDialog.vue'

defineOptions({
    name: 'serviceDetail',
})

const route = useRoute()
const { t } = useI18n()
const serviceId = ref(route.params.id)
const serviceData = ref({})
const activeTab = ref('group')
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
        title: t('pages.service.group.form.isolationCode'),
        dataIndex: 'isolation_code',
        width: 180,
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
        key: 'updatedAt',
        width: 180,
    },
    {
        title: t('button.action'),
        key: 'action',
        width: 120,
    },
]

onMounted(() => {
    loadServiceDetail()
    loadGroupList()
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
</script>

<style lang="less" scoped>
.service-detail {
    padding: 0;
}

.info-card {
    margin-bottom: 16px;
    .info-title {
        font-weight: 600;
        font-size: 15px;
        margin-bottom: 12px;
        color: #333;
    }
    .info-grid {
        display: grid;
        grid-template-columns: repeat(3, 1fr);
        gap: 12px 24px;
    }
    .info-item {
        display: flex;
        align-items: center;
        font-size: 13px;
        .info-label {
            color: #666;
            margin-right: 8px;
            white-space: nowrap;
        }
        .info-value {
            color: #333;
        }
    }
}

.detail-tabs {
    margin-bottom: 16px;
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
