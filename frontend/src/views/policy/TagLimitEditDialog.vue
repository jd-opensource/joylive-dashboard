<template>
    <a-drawer
        :open="modal.open"
        :title="modal.title"
        :width="720"
        placement="right"
        :closable="true"
        @close="handleCancel">
        <a-form
            ref="formRef"
            :model="formData"
            :rules="formRules"
            :label-col="{ style: { width: '110px' } }"
            :wrapper-col="{ flex: 1 }">
            <!-- 策略名称 -->
            <a-form-item
                :label="$t('pages.limit.form.name')"
                name="name">
                <a-input
                    :placeholder="$t('pages.limit.form.name.placeholder')"
                    v-model:value="formData.name"
                    :maxlength="60" />
            </a-form-item>

            <!-- 服务空间 -->
            <a-form-item
                :label="$t('pages.limit.form.space_code')"
                name="space_code">
                <a-select
                    :placeholder="$t('pages.limit.form.space_code.placeholder')"
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
                :label="$t('pages.limit.form.sourceApplicationId')"
                name="source_application_id">
                <a-select
                    :placeholder="$t('pages.limit.form.sourceApplicationId.placeholder')"
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
                :label="$t('pages.limit.form.targetServiceId')"
                name="target_service_id">
                <a-select
                    :placeholder="$t('pages.limit.form.targetServiceId.placeholder')"
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

            <!-- 服务分组 -->
            <a-form-item
                :label="$t('pages.limit.form.group')"
                name="group">
                <a-select
                    :placeholder="$t('pages.limit.form.group.placeholder')"
                    v-model:value="formData.group">
                    <a-select-option value="default">
                        {{ $t('pages.limit.form.group.default') }}
                    </a-select-option>
                </a-select>
            </a-form-item>

            <!-- 服务路径 -->
            <a-form-item
                :label="$t('pages.limit.form.path')"
                name="path">
                <a-input
                    :placeholder="$t('pages.limit.form.path.placeholder')"
                    v-model:value="formData.path" />
            </a-form-item>

            <!-- 服务方法 -->
            <a-form-item
                :label="$t('pages.limit.form.method')"
                name="method">
                <a-input
                    :placeholder="$t('pages.limit.form.method.placeholder')"
                    v-model:value="formData.method"
                    :disabled="!formData.path" />
            </a-form-item>

            <!-- 匹配方式 -->
            <a-form-item
                :label="$t('pages.limit.form.matchMethod')"
                name="relation_type"
                class="match-method-form-item">
                <div class="match-method-section">
                    <div class="match-method-header">
                        <span class="match-method-label">{{ $t('pages.limit.form.matchMethod.settingLabel') }}</span>
                        <a-radio-group v-model:value="formData.relation_type">
                            <a-radio value="AND">AND({{ $t('pages.limit.form.relationType.and') }})</a-radio>
                            <a-radio value="OR">OR({{ $t('pages.limit.form.relationType.or') }})</a-radio>
                        </a-radio-group>
                    </div>

                    <!-- 条件表格 -->
                    <div
                        class="conditions-table"
                        v-if="formData.conditions && formData.conditions.length > 0">
                        <div class="conditions-header">
                            <div class="col-type">{{ $t('pages.limit.form.conditions.type') }}</div>
                            <div class="col-key">{{ $t('pages.limit.form.conditions.key') }}</div>
                            <div class="col-op">
                                {{ $t('pages.limit.form.conditions.opType') }}
                                <a-tooltip :title="$t('pages.limit.form.conditions.opType.tooltip')">
                                    <question-circle-outlined style="margin-left: 4px; color: #999" />
                                </a-tooltip>
                            </div>
                            <div class="col-value">{{ $t('pages.limit.form.conditions.values') }}</div>
                            <div class="col-action">{{ $t('pages.limit.form.conditions.action') }}</div>
                        </div>
                        <div
                            class="conditions-row"
                            v-for="(condition, index) in formData.conditions"
                            :key="index">
                            <div class="col-type">
                                <a-select
                                    v-model:value="condition.type"
                                    :placeholder="$t('pages.limit.form.conditions.type.placeholder')"
                                    size="small">
                                    <a-select-option value="header">HEADER</a-select-option>
                                    <a-select-option value="query">QUERY</a-select-option>
                                    <a-select-option value="cookie">COOKIE</a-select-option>
                                </a-select>
                            </div>
                            <div class="col-key">
                                <a-input
                                    v-model:value="condition.key"
                                    :placeholder="$t('pages.limit.form.conditions.key.placeholder')"
                                    size="small" />
                            </div>
                            <div class="col-op">
                                <a-select
                                    v-model:value="condition.opType"
                                    :placeholder="$t('pages.limit.form.conditions.opType.placeholder')"
                                    size="small">
                                    <a-select-option value="EQUAL">{{
                                        $t('pages.limit.form.conditions.op.equal')
                                    }}</a-select-option>
                                    <a-select-option value="NOT_EQUAL">{{
                                        $t('pages.limit.form.conditions.op.notEqual')
                                    }}</a-select-option>
                                    <a-select-option value="IN">{{
                                        $t('pages.limit.form.conditions.op.contain')
                                    }}</a-select-option>
                                    <a-select-option value="NOT_IN">{{
                                        $t('pages.limit.form.conditions.op.notContain')
                                    }}</a-select-option>
                                    <a-select-option value="REGULAR">{{
                                        $t('pages.limit.form.conditions.op.regex')
                                    }}</a-select-option>
                                    <a-select-option value="PREFIX">{{
                                        $t('pages.limit.form.conditions.op.prefix')
                                    }}</a-select-option>
                                </a-select>
                            </div>
                            <div class="col-value">
                                <a-select
                                    v-model:value="condition.values"
                                    mode="tags"
                                    :placeholder="$t('pages.limit.form.conditions.values.placeholder')"
                                    size="small" />
                            </div>
                            <div class="col-action">
                                <minus-circle-outlined
                                    class="condition-remove-btn"
                                    @click="removeCondition(index)" />
                            </div>
                        </div>
                    </div>

                    <a
                        class="add-condition-link"
                        @click="addCondition">
                        {{ $t('pages.limit.form.conditions.add') }}
                    </a>
                </div>
            </a-form-item>

            <!-- 限流类型 -->
            <a-form-item
                :label="$t('pages.limit.form.type')"
                name="type">
                <a-radio-group v-model:value="formData.type">
                    <a-radio value="Rate">{{ $t('pages.limit.form.type.rate') }}</a-radio>
                    <a-radio value="Concurrency">{{ $t('pages.limit.form.type.concurrency') }}</a-radio>
                    <a-radio value="Load">{{ $t('pages.limit.form.type.load') }}</a-radio>
                </a-radio-group>
            </a-form-item>

            <!-- 实现方式 -->
            <a-form-item
                :label="$t('pages.limit.form.realizeType')"
                name="realize_type">
                <a-input
                    :placeholder="$t('pages.limit.form.realizeType.placeholder')"
                    v-model:value="formData.realize_type" />
            </a-form-item>

            <!-- 限流参数 (Rate) -->
            <template v-if="formData.type === 'Rate'">
                <a-form-item
                    :label="$t('pages.limit.form.slidingWindows')"
                    name="sliding_windows">
                    <div class="sliding-windows-section">
                        <div
                            class="sliding-window-row"
                            v-for="(window, index) in formData.sliding_windows"
                            :key="index">
                            <a-input-number
                                v-model:value="window.threshold"
                                :placeholder="$t('pages.limit.form.slidingWindows.threshold.placeholder')"
                                :min="1"
                                size="small"
                                style="width: 150px" />
                            <span class="sliding-window-separator">/</span>
                            <a-input-number
                                v-model:value="window.timeWindowInMs"
                                :placeholder="$t('pages.limit.form.slidingWindows.timeWindow.placeholder')"
                                :min="100"
                                size="small"
                                style="width: 150px" />
                            <span class="sliding-window-unit">ms</span>
                            <minus-circle-outlined
                                class="sliding-window-remove-btn"
                                @click="removeSlidingWindow(index)" />
                        </div>
                        <a
                            class="add-condition-link"
                            @click="addSlidingWindow">
                            {{ $t('pages.limit.form.slidingWindows.add') }}
                        </a>
                    </div>
                </a-form-item>
            </template>

            <!-- 并发参数 (Concurrency) -->
            <template v-if="formData.type === 'Concurrency'">
                <a-form-item
                    :label="$t('pages.limit.form.maxConcurrency')"
                    name="max_concurrency">
                    <a-input-number
                        :placeholder="$t('pages.limit.form.maxConcurrency.placeholder')"
                        v-model:value="formData.max_concurrency"
                        :min="0"
                        style="width: 100%" />
                </a-form-item>
            </template>

            <!-- 最大等待时间 (Rate & Concurrency) -->
            <template v-if="formData.type === 'Rate' || formData.type === 'Concurrency'">
                <a-form-item
                    :label="$t('pages.limit.form.maxWaitMs')"
                    name="max_wait_ms">
                    <a-input-number
                        :placeholder="$t('pages.limit.form.maxWaitMs.placeholder')"
                        v-model:value="formData.max_wait_ms"
                        :min="0"
                        style="width: 100%" />
                </a-form-item>
            </template>

            <!-- 负载参数 (Load) -->
            <template v-if="formData.type === 'Load'">
                <a-form-item
                    :label="$t('pages.limit.form.cpuUsage')"
                    name="cpu_usage">
                    <a-input-number
                        :placeholder="$t('pages.limit.form.cpuUsage.placeholder')"
                        v-model:value="formData.cpu_usage"
                        :min="0"
                        :max="100"
                        style="width: 100%" />
                </a-form-item>

                <a-form-item
                    :label="$t('pages.limit.form.loadUsage')"
                    name="load_usage">
                    <a-input-number
                        :placeholder="$t('pages.limit.form.loadUsage.placeholder')"
                        v-model:value="formData.load_usage"
                        :min="0"
                        style="width: 100%" />
                </a-form-item>
            </template>

            <!-- 生效状态 -->
            <a-form-item
                :label="$t('pages.limit.form.enabled')"
                name="enabled">
                <a-switch
                    v-model:checked="enabledSwitch"
                    :checked-children="$t('pages.limit.form.enabled.active')"
                    :un-checked-children="$t('pages.limit.form.enabled.inactive')" />
            </a-form-item>

            <!-- 描述 -->
            <a-form-item
                :label="$t('pages.limit.form.description')"
                name="description">
                <a-textarea
                    v-model:value="formData.description"
                    :placeholder="$t('pages.limit.form.description.placeholder')"
                    :maxlength="255"
                    show-count
                    :rows="3" />
            </a-form-item>

            <!-- 版本 -->
            <a-form-item
                :label="$t('pages.limit.form.version')"
                name="version">
                <a-input-number
                    v-model:value="formData.version"
                    :min="0"
                    style="width: 100%"
                    disabled />
            </a-form-item>
        </a-form>

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
import { QuestionCircleOutlined, MinusCircleOutlined } from '@ant-design/icons-vue'
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
formRules.value = {
    name: { required: true, message: t('pages.limit.form.name.required') },
    space_code: { required: true, message: t('pages.limit.form.space_code.required') },
    target_service_id: { required: true, message: t('pages.limit.form.targetServiceId.required') },
    relation_type: { required: true, message: t('pages.limit.form.relationType.required') },
    type: { required: true, message: t('pages.limit.form.type.required') },
    realize_type: { required: true, message: t('pages.limit.form.realizeType.required') },
    method: {
        validator: (rule, value) => {
            if (value && !formData.value.path) {
                return Promise.reject(t('pages.limit.form.method.requirePath'))
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

watch(
    () => formData.value.type,
    (val) => {
        // Reset type-specific fields when type changes
        if (val !== 'Rate') {
            formData.value.sliding_windows = undefined
        }
        if (val !== 'Concurrency') {
            formData.value.max_concurrency = undefined
        }
        if (val === 'Load') {
            formData.value.max_wait_ms = 0
        }
        if (val !== 'Load') {
            formData.value.cpu_usage = undefined
            formData.value.load_usage = undefined
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

// ---- 条件行管理 ----
function createEmptyCondition() {
    return { type: undefined, key: '', opType: undefined, values: [] }
}

function addCondition() {
    if (!formData.value.conditions) {
        formData.value.conditions = []
    }
    formData.value.conditions.push(createEmptyCondition())
}

function removeCondition(index) {
    formData.value.conditions.splice(index, 1)
}

// ---- 滑动窗口管理 ----
function createEmptySlidingWindow() {
    return { threshold: 1, timeWindowInMs: 1000 }
}

function addSlidingWindow() {
    if (!formData.value.sliding_windows) {
        formData.value.sliding_windows = []
    }
    formData.value.sliding_windows.push(createEmptySlidingWindow())
}

function removeSlidingWindow(index) {
    formData.value.sliding_windows.splice(index, 1)
}

function handleCreate() {
    formData.value.group = 'default'
    formData.value.enabled = 0
    formData.value.relation_type = 'AND'
    formData.value.type = 'Rate'
    formData.value.realize_type = 'Resilience4j'
    formData.value.max_wait_ms = 0
    formData.value.version = 0
    formData.value.conditions = [createEmptyCondition()]
    formData.value.sliding_windows = [createEmptySlidingWindow()]
    formData.value.space_code = initSpaceCode(props.spaceOptions)
    showModal({
        type: 'create',
        title: t('pages.limit.add'),
    })
}

async function handleEdit(record = {}) {
    showModal({
        type: 'edit',
        title: t('pages.limit.edit'),
    })

    const { data, success } = await apis.policy.getLimit(record.id).catch()
    if (!success) {
        message.error(t('component.message.error.save'))
        hideModal()
        return
    }
    formRecord.value = data
    const cloned = cloneDeep(data)
    // conditions 从后端返回可能是 JSON 字符串，需要解析为数组
    if (typeof cloned.conditions === 'string') {
        try {
            cloned.conditions = JSON.parse(cloned.conditions)
        } catch {
            cloned.conditions = []
        }
    }
    if (!Array.isArray(cloned.conditions) || cloned.conditions.length === 0) {
        cloned.conditions = [createEmptyCondition()]
    }
    // sliding_windows 从后端返回可能是 JSON 字符串，需要解析为数组
    if (typeof cloned.sliding_windows === 'string') {
        try {
            cloned.sliding_windows = JSON.parse(cloned.sliding_windows)
        } catch {
            cloned.sliding_windows = []
        }
    }
    if (!Array.isArray(cloned.sliding_windows) || cloned.sliding_windows.length === 0) {
        if (cloned.type === 'Rate') {
            cloned.sliding_windows = [createEmptySlidingWindow()]
        }
    }
    formData.value = cloned
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
                    conditions: formData.value.conditions || undefined,
                    sliding_windows: formData.value.sliding_windows || undefined,
                    action_parameters: formData.value.action_parameters || undefined,
                }
                let result = null
                switch (modal.value.type) {
                    case 'create':
                        result = await apis.policy.createLimit(params).catch(() => {
                            throw new Error()
                        })
                        break
                    case 'edit':
                        result = await apis.policy.updateLimit(formData.value.id, params).catch(() => {
                            throw new Error()
                        })
                        break
                }
                hideLoading()
                if (config('http.code.success') === result?.success) {
                    hideModal()
                    emit('ok')
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
}

defineExpose({
    handleCreate,
    handleEdit,
})
</script>

<style lang="less" scoped>
.form-hint {
    color: #999;
    font-size: 12px;
    line-height: 1.5;
    margin-top: 4px;
}

.match-method-form-item {
    :deep(.ant-form-item-control-input-content) {
        overflow: visible;
    }
}

.match-method-section {
    border: 1px solid #f0f0f0;
    border-radius: 6px;
    padding: 16px;
    background: #fafafa;
}

.match-method-header {
    display: flex;
    align-items: center;
    margin-bottom: 12px;
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
    padding: 8px 0;
    border-bottom: 1px solid #e8e8e8;
    font-weight: 500;
    font-size: 13px;
    color: #333;
    gap: 8px;
}

.conditions-row {
    display: flex;
    align-items: center;
    padding: 8px 0;
    border-bottom: 1px solid #f0f0f0;
    gap: 8px;

    &:last-child {
        border-bottom: none;
    }
}

.col-type {
    flex: 0 0 100px;
}

.col-key {
    flex: 0 0 110px;
}

.col-op {
    flex: 0 0 100px;
    display: flex;
    align-items: center;
}

.col-value {
    flex: 1;
    min-width: 0;
}

.col-action {
    flex: 0 0 40px;
    text-align: center;
}

.condition-remove-btn {
    font-size: 18px;
    color: #999;
    cursor: pointer;
    transition: color 0.2s;

    &:hover {
        color: #ff4d4f;
    }
}

.add-condition-link {
    display: inline-block;
    margin-top: 8px;
    color: #1890ff;
    font-size: 13px;
    cursor: pointer;

    &:hover {
        color: #40a9ff;
    }
}

.sliding-windows-section {
    border: 1px solid #f0f0f0;
    border-radius: 6px;
    padding: 12px;
    background: #fafafa;
}

.sliding-window-row {
    display: flex;
    align-items: center;
    gap: 8px;
    margin-bottom: 8px;

    &:last-of-type {
        margin-bottom: 0;
    }
}

.sliding-window-separator {
    color: #999;
    font-size: 14px;
}

.sliding-window-unit {
    color: #666;
    font-size: 12px;
}

.sliding-window-remove-btn {
    font-size: 16px;
    color: #999;
    cursor: pointer;
    transition: color 0.2s;

    &:hover {
        color: #ff4d4f;
    }
}
</style>
