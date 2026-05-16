<template>
    <a-drawer
        :open="modal.open"
        :title="modal.title"
        :width="900"
        placement="right"
        :closable="true"
        @close="handleCancel">
        <a-form
            ref="formRef"
            :model="formData"
            :rules="formRules"
            :label-col="{ style: { width: '110px' } }"
            :wrapper-col="{ flex: 1 }">
            <!-- 规则名称 -->
            <a-form-item
                :label="$t('pages.tagRoute.form.name')"
                name="name">
                <a-input
                    :placeholder="$t('pages.tagRoute.form.name.placeholder')"
                    v-model:value="formData.name"
                    :maxlength="60" />
            </a-form-item>

            <!-- 服务空间 -->
            <a-form-item
                :label="$t('pages.tagRoute.form.space_code')"
                name="space_code">
                <a-select
                    :placeholder="$t('pages.tagRoute.form.space_code.placeholder')"
                    v-model:value="formData.space_code"
                    allow-clear
                    show-search
                    :filter-option="filterSpaceOption">
                    <a-select-option
                        v-for="item in spaceOptions"
                        :key="item.code"
                        :value="item.code">
                        {{ item.name }} ({{ item.code }})
                    </a-select-option>
                </a-select>
            </a-form-item>

            <!-- 调用应用 -->
            <a-form-item
                :label="$t('pages.tagRoute.form.sourceApplicationId')"
                name="source_application_id">
                <a-select
                    :placeholder="$t('pages.tagRoute.form.sourceApplicationId.placeholder')"
                    v-model:value="formData.source_application_id"
                    allow-clear
                    show-search
                    :filter-option="filterAppOption">
                    <a-select-option
                        v-for="item in applicationOptions"
                        :key="item.id"
                        :value="item.id">
                        {{ item.name }}
                    </a-select-option>
                </a-select>
            </a-form-item>

            <!-- 目标服务 -->
            <a-form-item
                :label="$t('pages.tagRoute.form.targetServiceId')"
                name="target_service_id">
                <a-select
                    :placeholder="$t('pages.tagRoute.form.targetServiceId.placeholder')"
                    v-model:value="formData.target_service_id"
                    show-search
                    :filter-option="filterServiceOption">
                    <a-select-option
                        v-for="item in filteredServiceOptions"
                        :key="item.id"
                        :value="item.id">
                        {{ item.name }}
                    </a-select-option>
                </a-select>
            </a-form-item>

            <!-- 当前服务分组 -->
            <a-form-item
                :label="$t('pages.tagRoute.form.group')"
                name="group">
                <a-select
                    :placeholder="$t('pages.tagRoute.form.group.placeholder')"
                    v-model:value="formData.group">
                    <a-select-option value="default">
                        {{ $t('pages.tagRoute.form.group.default') }}
                    </a-select-option>
                </a-select>
            </a-form-item>

            <!-- 当前服务路径 -->
            <a-form-item
                :label="$t('pages.tagRoute.form.path')"
                name="path">
                <a-input
                    :placeholder="$t('pages.tagRoute.form.path.placeholder')"
                    v-model:value="formData.path" />
            </a-form-item>

            <!-- 当前服务方法 -->
            <a-form-item
                :label="$t('pages.tagRoute.form.method')"
                name="method">
                <a-input
                    :placeholder="$t('pages.tagRoute.form.method.placeholder')"
                    v-model:value="formData.method"
                    :disabled="!formData.path" />
            </a-form-item>

            <!-- 生效状态 -->
            <a-form-item
                :label="$t('pages.tagRoute.form.enabled')"
                name="enabled">
                <a-switch
                    v-model:checked="enabledSwitch"
                    :checked-children="$t('pages.tagRoute.form.enabled.active')"
                    :un-checked-children="$t('pages.tagRoute.form.enabled.inactive')" />
            </a-form-item>

            <!-- 描述 -->
            <a-form-item
                :label="$t('pages.tagRoute.form.description')"
                name="description">
                <a-textarea
                    v-model:value="formData.description"
                    :placeholder="$t('pages.tagRoute.form.description.placeholder')"
                    :maxlength="255"
                    show-count
                    :rows="3" />
            </a-form-item>
        </a-form>

        <!-- 路由规则列表 (RouteDetail 1:n) -->
        <div class="details-section">
            <div class="details-section-header">
                <span class="details-section-title">{{ $t('pages.tagRoute.form.details.title') }}</span>
                <a-button
                    type="primary"
                    size="small"
                    @click="addDetail">
                    <template #icon><plus-outlined /></template>
                    {{ $t('pages.tagRoute.form.details.add') }}
                </a-button>
            </div>

            <a-collapse
                v-model:activeKey="activeDetailKeys"
                :bordered="false">
                <a-collapse-panel
                    v-for="(detail, rIndex) in formData.details"
                    :key="String(rIndex)"
                    :header="$t('pages.tagRoute.form.details.rule') + ' #' + (rIndex + 1)"
                    :collapsible="'header'">
                    <template #extra>
                        <delete-outlined
                            class="detail-remove-btn"
                            @click.stop="removeDetail(rIndex)" />
                    </template>

                    <!-- 排序 -->
                    <div class="detail-field">
                        <label class="detail-field-label">{{ $t('pages.tagRoute.form.order') }}</label>
                        <a-input-number
                            v-model:value="detail.order"
                            :min="0"
                            size="small"
                            style="width: 150px" />
                    </div>

                    <!-- 匹配条件 -->
                    <div class="detail-field">
                        <label class="detail-field-label">{{ $t('pages.tagRoute.form.matchMethod') }}</label>
                        <div class="match-method-section">
                            <div class="match-method-header">
                                <span class="match-method-label">{{
                                    $t('pages.tagRoute.form.matchMethod.settingLabel')
                                }}</span>
                                <a-radio-group
                                    v-model:value="detail.relationType"
                                    size="small">
                                    <a-radio value="AND">AND({{ $t('pages.tagRoute.form.relationType.and') }})</a-radio>
                                    <a-radio value="OR">OR({{ $t('pages.tagRoute.form.relationType.or') }})</a-radio>
                                </a-radio-group>
                            </div>

                            <div
                                class="conditions-table"
                                v-if="detail.conditions && detail.conditions.length > 0">
                                <div class="conditions-header">
                                    <div class="col-type">{{ $t('pages.tagRoute.form.conditions.type') }}</div>
                                    <div class="col-key">{{ $t('pages.tagRoute.form.conditions.key') }}</div>
                                    <div class="col-op">
                                        {{ $t('pages.tagRoute.form.conditions.opType') }}
                                        <a-tooltip :title="$t('pages.tagRoute.form.conditions.opType.tooltip')">
                                            <question-circle-outlined style="margin-left: 4px; color: #999" />
                                        </a-tooltip>
                                    </div>
                                    <div class="col-value">{{ $t('pages.tagRoute.form.conditions.values') }}</div>
                                    <div class="col-action">{{ $t('pages.tagRoute.form.conditions.action') }}</div>
                                </div>
                                <div
                                    class="conditions-row"
                                    v-for="(condition, cIndex) in detail.conditions"
                                    :key="cIndex">
                                    <div class="col-type">
                                        <a-select
                                            v-model:value="condition.type"
                                            :placeholder="$t('pages.tagRoute.form.conditions.type.placeholder')"
                                            size="small">
                                            <a-select-option value="HEADER">HEADER</a-select-option>
                                            <a-select-option value="QUERY">QUERY</a-select-option>
                                            <a-select-option value="COOKIE">COOKIE</a-select-option>
                                        </a-select>
                                    </div>
                                    <div class="col-key">
                                        <a-input
                                            v-model:value="condition.key"
                                            :placeholder="$t('pages.tagRoute.form.conditions.key.placeholder')"
                                            size="small" />
                                    </div>
                                    <div class="col-op">
                                        <a-select
                                            v-model:value="condition.opType"
                                            :placeholder="$t('pages.tagRoute.form.conditions.opType.placeholder')"
                                            size="small">
                                            <a-select-option value="EQUAL">{{
                                                $t('pages.tagRoute.form.conditions.op.equal')
                                            }}</a-select-option>
                                            <a-select-option value="NOT_EQUAL">{{
                                                $t('pages.tagRoute.form.conditions.op.notEqual')
                                            }}</a-select-option>
                                            <a-select-option value="IN">{{
                                                $t('pages.tagRoute.form.conditions.op.contain')
                                            }}</a-select-option>
                                            <a-select-option value="NOT_IN">{{
                                                $t('pages.tagRoute.form.conditions.op.notContain')
                                            }}</a-select-option>
                                            <a-select-option value="REGULAR">{{
                                                $t('pages.tagRoute.form.conditions.op.regex')
                                            }}</a-select-option>
                                            <a-select-option value="PREFIX">{{
                                                $t('pages.tagRoute.form.conditions.op.prefix')
                                            }}</a-select-option>
                                        </a-select>
                                    </div>
                                    <div class="col-value">
                                        <a-select
                                            v-model:value="condition.values"
                                            mode="tags"
                                            :placeholder="$t('pages.tagRoute.form.conditions.values.placeholder')"
                                            size="small" />
                                    </div>
                                    <div class="col-action">
                                        <minus-circle-outlined
                                            class="condition-remove-btn"
                                            @click="removeCondition(rIndex, cIndex)" />
                                    </div>
                                </div>
                            </div>
                            <a
                                class="add-condition-link"
                                @click="addCondition(rIndex)">
                                {{ $t('pages.tagRoute.form.conditions.add') }}
                            </a>
                        </div>
                    </div>

                    <!-- 路由目标 -->
                    <div class="detail-field">
                        <label class="detail-field-label">{{ $t('pages.tagRoute.form.destinations') }}</label>
                        <div class="destinations-section">
                            <div
                                class="destination-card"
                                v-for="(destination, dIndex) in detail.destinations"
                                :key="dIndex">
                                <div class="destination-card-header">
                                    <span class="destination-card-title">{{
                                        $t('pages.tagRoute.form.destinations.title') + ' #' + (dIndex + 1)
                                    }}</span>
                                    <minus-circle-outlined
                                        class="destination-remove-btn"
                                        @click="removeDestination(rIndex, dIndex)" />
                                </div>
                                <div class="destination-card-body">
                                    <div class="destination-field">
                                        <label class="destination-field-label">{{
                                            $t('pages.tagRoute.form.destinations.weight')
                                        }}</label>
                                        <a-input-number
                                            v-model:value="destination.weight"
                                            :placeholder="$t('pages.tagRoute.form.destinations.weight.placeholder')"
                                            :min="0"
                                            size="small"
                                            style="width: 150px" />
                                    </div>
                                    <div class="destination-field">
                                        <label class="destination-field-label">{{
                                            $t('pages.tagRoute.form.destinations.relationType')
                                        }}</label>
                                        <a-radio-group
                                            v-model:value="destination.relationType"
                                            size="small">
                                            <a-radio value="AND"
                                                >AND({{ $t('pages.tagRoute.form.relationType.and') }})</a-radio
                                            >
                                            <a-radio value="OR"
                                                >OR({{ $t('pages.tagRoute.form.relationType.or') }})</a-radio
                                            >
                                        </a-radio-group>
                                    </div>
                                    <div class="destination-field">
                                        <label class="destination-field-label">{{
                                            $t('pages.tagRoute.form.destinations.conditions')
                                        }}</label>
                                        <div class="dest-conditions-table">
                                            <div class="dest-conditions-header">
                                                <div class="dest-col-key">
                                                    {{ $t('pages.tagRoute.form.conditions.key') }}
                                                </div>
                                                <div class="dest-col-op">
                                                    {{ $t('pages.tagRoute.form.conditions.opType') }}
                                                </div>
                                                <div class="dest-col-value">
                                                    {{ $t('pages.tagRoute.form.conditions.values') }}
                                                </div>
                                                <div class="dest-col-action">
                                                    {{ $t('pages.tagRoute.form.conditions.action') }}
                                                </div>
                                            </div>
                                            <div
                                                class="dest-conditions-row"
                                                v-for="(dc, dcIndex) in destination.conditions"
                                                :key="dcIndex">
                                                <div class="dest-col-key">
                                                    <a-input
                                                        v-model:value="dc.key"
                                                        :placeholder="
                                                            $t('pages.tagRoute.form.conditions.key.placeholder')
                                                        "
                                                        size="small" />
                                                </div>
                                                <div class="dest-col-op">
                                                    <a-select
                                                        v-model:value="dc.opType"
                                                        :placeholder="
                                                            $t('pages.tagRoute.form.conditions.opType.placeholder')
                                                        "
                                                        size="small">
                                                        <a-select-option value="EQUAL">{{
                                                            $t('pages.tagRoute.form.conditions.op.equal')
                                                        }}</a-select-option>
                                                        <a-select-option value="NOT_EQUAL">{{
                                                            $t('pages.tagRoute.form.conditions.op.notEqual')
                                                        }}</a-select-option>
                                                        <a-select-option value="IN">{{
                                                            $t('pages.tagRoute.form.conditions.op.contain')
                                                        }}</a-select-option>
                                                        <a-select-option value="NOT_IN">{{
                                                            $t('pages.tagRoute.form.conditions.op.notContain')
                                                        }}</a-select-option>
                                                        <a-select-option value="REGULAR">{{
                                                            $t('pages.tagRoute.form.conditions.op.regex')
                                                        }}</a-select-option>
                                                        <a-select-option value="PREFIX">{{
                                                            $t('pages.tagRoute.form.conditions.op.prefix')
                                                        }}</a-select-option>
                                                    </a-select>
                                                </div>
                                                <div class="dest-col-value">
                                                    <a-select
                                                        v-model:value="dc.values"
                                                        mode="tags"
                                                        :placeholder="
                                                            $t('pages.tagRoute.form.conditions.values.placeholder')
                                                        "
                                                        size="small" />
                                                </div>
                                                <div class="dest-col-action">
                                                    <minus-circle-outlined
                                                        class="condition-remove-btn"
                                                        @click="removeDestCondition(rIndex, dIndex, dcIndex)" />
                                                </div>
                                            </div>
                                        </div>
                                        <a
                                            class="add-condition-link"
                                            @click="addDestCondition(rIndex, dIndex)">
                                            {{ $t('pages.tagRoute.form.conditions.add') }}
                                        </a>
                                    </div>
                                </div>
                            </div>
                            <a-button
                                type="dashed"
                                block
                                size="small"
                                @click="addDestination(rIndex)"
                                style="margin-top: 8px">
                                <template #icon><plus-outlined /></template>
                                {{ $t('pages.tagRoute.form.destinations.add') }}
                            </a-button>
                        </div>
                    </div>
                </a-collapse-panel>
            </a-collapse>
        </div>

        <template #footer>
            <div style="text-align: right">
                <a-space>
                    <a-button @click="handleCancel">{{ cancelText }}</a-button>
                    <a-button
                        type="primary"
                        :loading="modal.confirmLoading"
                        @click="handleOk">
                        {{ okText }}
                    </a-button>
                </a-space>
            </div>
        </template>
    </a-drawer>
</template>

<script setup>
import { cloneDeep } from 'lodash-es'
import { message } from 'ant-design-vue'
import { ref, computed, watch } from 'vue'
import { config } from '@/config'
import apis from '@/apis'
import { useForm, useModal } from '@/hooks'
import { QuestionCircleOutlined, MinusCircleOutlined, PlusOutlined, DeleteOutlined } from '@ant-design/icons-vue'
import { initSpaceCode, setCurrentSpaceCode } from '@/utils/spaceStorage'

const props = defineProps({
    spaceOptions: { type: Array, default: () => [] },
    serviceOptions: { type: Array, default: () => [] },
    applicationOptions: { type: Array, default: () => [] },
})

const emit = defineEmits(['ok'])
import { useI18n } from 'vue-i18n'
const { modal, showModal, hideModal, showLoading, hideLoading } = useModal()
const { formRecord, formData, formRef, formRules, resetForm } = useForm()
const { t } = useI18n()
const cancelText = ref(t('button.cancel'))
const okText = ref(t('button.confirm'))
const activeDetailKeys = ref([])

formRules.value = {
    name: { required: true, message: t('pages.tagRoute.form.name.required') },
    space_code: { required: true, message: t('pages.tagRoute.form.space_code.required') },
    target_service_id: { required: true, message: t('pages.tagRoute.form.targetServiceId.required') },
    method: {
        validator: (rule, value) => {
            if (value && !formData.value.path) {
                return Promise.reject(t('pages.tagRoute.form.method.requirePath'))
            }
            return Promise.resolve()
        },
    },
}

const enabledSwitch = computed({
    get: () => formData.value.enabled === 1,
    set: (val) => {
        formData.value.enabled = val ? 1 : 0
    },
})

const filteredServiceOptions = computed(() => {
    const spaceCode = formData.value.space_code
    if (!spaceCode) return props.serviceOptions
    return props.serviceOptions.filter((item) => item.space_code === spaceCode)
})

watch(
    () => formData.value.space_code,
    (val, oldVal) => {
        if (oldVal !== undefined) {
            formData.value.target_service_id = undefined
        }
        if (val) {
            setCurrentSpaceCode(val)
        }
    }
)

function filterSpaceOption(input, option) {
    const label = option.children?.[0]?.children || ''
    return option.value.toLowerCase().includes(input.toLowerCase()) || label.toLowerCase().includes(input.toLowerCase())
}

function filterAppOption(input, option) {
    const label = option.children?.[0]?.children || ''
    return label.toLowerCase().includes(input.toLowerCase())
}

function filterServiceOption(input, option) {
    const label = option.children?.[0]?.children || ''
    return label.toLowerCase().includes(input.toLowerCase())
}

// ---- Detail (RouteDetail) 管理 ----
function createEmptyDetail() {
    return {
        _id: undefined,
        order: 0,
        relationType: 'AND',
        conditions: [createEmptyCondition()],
        destinations: [createEmptyDestination()],
    }
}

function addDetail() {
    if (!formData.value.details) {
        formData.value.details = []
    }
    const idx = formData.value.details.length
    formData.value.details.push(createEmptyDetail())
    activeDetailKeys.value.push(String(idx))
}

function removeDetail(index) {
    formData.value.details.splice(index, 1)
}

// ---- 条件行管理 ----
function createEmptyCondition() {
    return { type: undefined, key: '', opType: undefined, values: [] }
}

function addCondition(rIndex) {
    if (!formData.value.details[rIndex].conditions) {
        formData.value.details[rIndex].conditions = []
    }
    formData.value.details[rIndex].conditions.push(createEmptyCondition())
}

function removeCondition(rIndex, cIndex) {
    formData.value.details[rIndex].conditions.splice(cIndex, 1)
}

// ---- 目标地址管理 ----
function createEmptyDestination() {
    return {
        weight: 100,
        relationType: 'AND',
        conditions: [createEmptyDestCondition()],
    }
}

function createEmptyDestCondition() {
    return { opType: undefined, key: '', values: [] }
}

function addDestination(rIndex) {
    if (!formData.value.details[rIndex].destinations) {
        formData.value.details[rIndex].destinations = []
    }
    formData.value.details[rIndex].destinations.push(createEmptyDestination())
}

function removeDestination(rIndex, dIndex) {
    formData.value.details[rIndex].destinations.splice(dIndex, 1)
}

function addDestCondition(rIndex, dIndex) {
    if (!formData.value.details[rIndex].destinations[dIndex].conditions) {
        formData.value.details[rIndex].destinations[dIndex].conditions = []
    }
    formData.value.details[rIndex].destinations[dIndex].conditions.push(createEmptyDestCondition())
}

function removeDestCondition(rIndex, dIndex, dcIndex) {
    formData.value.details[rIndex].destinations[dIndex].conditions.splice(dcIndex, 1)
}

// ---- 数据加载 ----
async function loadRouteDetails(routeId) {
    try {
        const { success, data } = await apis.policy
            .getRouteDetailList({ route_id: routeId, pageSize: 99, current: 1 })
            .catch(() => {
                throw new Error()
            })
        if (config('http.code.success') === success && data && data.length > 0) {
            formData.value.details = data.map((item) => {
                let conditions = item.conditions
                // Handle both JSON string and array formats
                if (typeof conditions === 'string') {
                    try {
                        conditions = JSON.parse(conditions)
                    } catch {
                        conditions = []
                    }
                }
                let destinations = item.destinations
                // Handle both JSON string and array formats
                if (typeof destinations === 'string') {
                    try {
                        destinations = JSON.parse(destinations)
                    } catch {
                        destinations = []
                    }
                }
                return {
                    _id: item.id,
                    order: item.order || 0,
                    relationType: item.relation_type || 'AND',
                    conditions:
                        Array.isArray(conditions) && conditions.length > 0 ? conditions : [createEmptyCondition()],
                    destinations:
                        Array.isArray(destinations) && destinations.length > 0
                            ? destinations.map((d) => ({
                                  ...d,
                                  conditions:
                                      Array.isArray(d.conditions) && d.conditions.length > 0
                                          ? d.conditions
                                          : [createEmptyDestCondition()],
                              }))
                            : [createEmptyDestination()],
                }
            })
            activeDetailKeys.value = formData.value.details.map((_, i) => String(i))
        }
    } catch (error) {
        // ignore
    }
}

// ---- 保存 ----
async function saveRouteDetails(routeId) {
    // 获取已有的 detail IDs
    const existingIds = new Set()
    try {
        const { success, data } = await apis.policy
            .getRouteDetailList({ route_id: routeId, pageSize: 99, current: 1 })
            .catch(() => ({ success: false, data: [] }))
        if (config('http.code.success') === success && data) {
            data.forEach((item) => existingIds.add(item.id))
        }
    } catch {
        // ignore
    }

    const submittedIds = new Set()

    for (const detail of formData.value.details || []) {
        const params = {
            route_id: routeId,
            relation_type: detail.relationType || 'AND',
            order: detail.order || 0,
            conditions: detail.conditions || undefined,
            destinations: detail.destinations || undefined,
            enabled: formData.value.enabled,
        }

        if (detail._id) {
            submittedIds.add(detail._id)
            await apis.policy.updateRouteDetail(detail._id, params).catch(() => {
                throw new Error()
            })
        } else {
            const result = await apis.policy.createRouteDetail(params).catch(() => {
                throw new Error()
            })
            if (result?.data?.id) {
                submittedIds.add(result.data.id)
            }
        }
    }

    // 删除不再存在的 detail
    for (const id of existingIds) {
        if (!submittedIds.has(id)) {
            await apis.policy.delRouteDetail(id).catch(() => {
                // ignore
            })
        }
    }
}

function handleCreate() {
    formData.value.group = 'default'
    formData.value.enabled = 0
    formData.value.details = [createEmptyDetail()]
    formData.value.space_code = initSpaceCode(props.spaceOptions)
    activeDetailKeys.value = ['0']
    showModal({
        type: 'create',
        title: t('pages.tagRoute.add'),
    })
}

async function handleEdit(record = {}) {
    showModal({
        type: 'edit',
        title: t('pages.tagRoute.edit'),
    })

    const { data, success } = await apis.policy.getRoute(record.id).catch()
    if (!success) {
        message.error(t('component.message.error.save'))
        hideModal()
        return
    }
    formRecord.value = data
    const cloned = cloneDeep(data)
    cloned.details = []
    formData.value = cloned
    await loadRouteDetails(record.id)
    if (!Array.isArray(formData.value.details) || formData.value.details.length === 0) {
        formData.value.details = [createEmptyDetail()]
        activeDetailKeys.value = ['0']
    }
}

function handleOk() {
    formRef.value
        .validateFields()
        .then(async (values) => {
            try {
                showLoading()
                const params = {
                    ...values,
                    group: formData.value.group || 'default',
                }
                let result = null
                switch (modal.value.type) {
                    case 'create':
                        result = await apis.policy.createRoute(params).catch(() => {
                            throw new Error()
                        })
                        break
                    case 'edit':
                        result = await apis.policy.updateRoute(formData.value.id, params).catch(() => {
                            throw new Error()
                        })
                        break
                }
                if (config('http.code.success') === result?.success) {
                    const routeId = result.data?.id || formData.value.id
                    await saveRouteDetails(routeId)
                    hideLoading()
                    hideModal()
                    emit('ok')
                } else {
                    hideLoading()
                }
            } catch (error) {
                hideLoading()
            }
        })
        .catch(() => {
            hideLoading()
        })
}

function handleCancel() {
    hideModal()
    onAfterClose()
}

function onAfterClose() {
    resetForm()
    hideLoading()
    activeDetailKeys.value = []
}

defineExpose({
    handleCreate,
    handleEdit,
})
</script>

<style lang="less" scoped>
.details-section {
    margin-top: 16px;
    border-top: 1px solid #f0f0f0;
    padding-top: 16px;
}

.details-section-header {
    display: flex;
    align-items: center;
    justify-content: space-between;
    margin-bottom: 12px;
}

.details-section-title {
    font-size: 15px;
    font-weight: 600;
    color: #333;
}

.detail-remove-btn {
    font-size: 16px;
    color: #999;
    cursor: pointer;
    transition: color 0.2s;

    &:hover {
        color: #ff4d4f;
    }
}

:deep(.ant-collapse) {
    background: transparent;
}

:deep(.ant-collapse-item) {
    margin-bottom: 8px;
    border: 1px solid #e8e8e8;
    border-radius: 6px;
    overflow: hidden;
}

:deep(.ant-collapse-header) {
    font-weight: 500;
    background: #fafafa;
}

:deep(.ant-collapse-content) {
    border-top: 1px solid #f0f0f0;
}

:deep(.ant-collapse-content-box) {
    padding: 12px;
}

.detail-field {
    margin-bottom: 12px;

    &:last-child {
        margin-bottom: 0;
    }
}

.detail-field-label {
    display: block;
    font-size: 13px;
    font-weight: 500;
    color: #333;
    margin-bottom: 6px;
}

.match-method-section {
    border: 1px solid #f0f0f0;
    border-radius: 6px;
    padding: 12px;
    background: #fafafa;
}

.match-method-header {
    display: flex;
    align-items: center;
    margin-bottom: 10px;
}

.match-method-label {
    font-size: 13px;
    color: #333;
    margin-right: 8px;
    white-space: nowrap;
}

.conditions-table {
    margin-bottom: 8px;
}

.conditions-header {
    display: flex;
    align-items: center;
    padding: 6px 0;
    border-bottom: 1px solid #e8e8e8;
    font-weight: 500;
    font-size: 12px;
    color: #666;
    gap: 6px;
}

.conditions-row {
    display: flex;
    align-items: center;
    padding: 6px 0;
    border-bottom: 1px solid #f0f0f0;
    gap: 6px;

    &:last-child {
        border-bottom: none;
    }
}

.col-type {
    flex: 0 0 90px;
}

.col-key {
    flex: 0 0 110px;
}

.col-op {
    flex: 0 0 90px;
    display: flex;
    align-items: center;
}

.col-value {
    flex: 1;
    min-width: 0;
}

.col-action {
    flex: 0 0 30px;
    text-align: center;
}

.condition-remove-btn {
    font-size: 16px;
    color: #999;
    cursor: pointer;
    transition: color 0.2s;

    &:hover {
        color: #ff4d4f;
    }
}

.add-condition-link {
    display: inline-block;
    margin-top: 6px;
    color: #1890ff;
    font-size: 13px;
    cursor: pointer;

    &:hover {
        color: #40a9ff;
    }
}

.destinations-section {
    border: 1px solid #f0f0f0;
    border-radius: 6px;
    padding: 12px;
    background: #fafafa;
}

.destination-card {
    border: 1px solid #e8e8e8;
    border-radius: 6px;
    background: #fff;
    margin-bottom: 10px;

    &:last-of-type {
        margin-bottom: 8px;
    }
}

.destination-card-header {
    display: flex;
    align-items: center;
    justify-content: space-between;
    padding: 6px 10px;
    border-bottom: 1px solid #f0f0f0;
    background: #fafafa;
    border-radius: 6px 6px 0 0;
}

.destination-card-title {
    font-size: 13px;
    font-weight: 500;
    color: #333;
}

.destination-remove-btn {
    font-size: 16px;
    color: #999;
    cursor: pointer;
    transition: color 0.2s;

    &:hover {
        color: #ff4d4f;
    }
}

.destination-card-body {
    padding: 10px;
}

.destination-field {
    margin-bottom: 8px;

    &:last-child {
        margin-bottom: 0;
    }
}

.destination-field-label {
    display: block;
    font-size: 12px;
    color: #666;
    margin-bottom: 4px;
}

.dest-conditions-table {
    border: 1px solid #f0f0f0;
    border-radius: 4px;
    background: #fafafa;
    padding: 6px;
}

.dest-conditions-header {
    display: flex;
    align-items: center;
    padding: 4px 0;
    border-bottom: 1px solid #e8e8e8;
    font-weight: 500;
    font-size: 12px;
    color: #666;
    gap: 6px;
}

.dest-conditions-row {
    display: flex;
    align-items: center;
    padding: 4px 0;
    border-bottom: 1px solid #f0f0f0;
    gap: 6px;

    &:last-child {
        border-bottom: none;
    }
}

.dest-col-key {
    flex: 0 0 110px;
}

.dest-col-op {
    flex: 0 0 90px;
}

.dest-col-value {
    flex: 1;
    min-width: 0;
}

.dest-col-action {
    flex: 0 0 30px;
    text-align: center;
}
</style>
