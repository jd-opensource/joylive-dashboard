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
            <!-- 规则名称 -->
            <a-form-item
                :label="$t('pages.limit.form.name')"
                name="name">
                <a-input
                    :placeholder="$t('pages.limit.form.name.placeholder')"
                    v-model:value="formData.name"
                    :maxlength="60" />
                <!-- <div class="form-hint">{{ $t('pages.limit.form.name.hint') }}</div> -->
            </a-form-item>

            <!-- 服务空间 -->
            <a-form-item
                :label="$t('pages.limit.form.space_code') || '服务空间'"
                name="space_code">
                <a-select
                    :placeholder="$t('pages.limit.form.space_code.placeholder') || '请选择服务空间'"
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
                :label="$t('pages.limit.form.sourceApplicationId') || '调用应用'"
                name="source_application_id">
                <a-select
                    :placeholder="$t('pages.limit.form.sourceApplicationId.placeholder') || '请选择应用'"
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
                :label="$t('pages.limit.form.targetServiceId') || '目标服务'"
                name="target_service_id">
                <a-select
                    :placeholder="$t('pages.limit.form.targetServiceId.placeholder') || '请选择目标服务'"
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

            <!-- 当前服务路径 -->
            <a-form-item
                :label="$t('pages.limit.form.path')"
                name="path">
                <a-input
                    :placeholder="$t('pages.limit.form.path.placeholder')"
                    v-model:value="formData.path" />
            </a-form-item>

            <!-- 当前服务方法 -->
            <a-form-item
                :label="$t('pages.limit.form.method')"
                name="method">
                <a-input
                    :placeholder="$t('pages.limit.form.method.placeholder')"
                    v-model:value="formData.method"
                    :disabled="!formData.path" />
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

            <!-- 流量匹配规则 -->
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

            <!-- 限流规则 (Rate) -->
            <template v-if="formData.type === 'Rate'">
                <a-form-item
                    name="sliding_windows"
                    class="limit-rules-form-item">
                    <template #label>
                        <span>
                            {{ $t('pages.limit.form.limitRules') }}
                            <a-tooltip :title="$t('pages.limit.form.limitRules.tooltip')">
                                <question-circle-outlined style="margin-left: 4px; color: #999" />
                            </a-tooltip>
                        </span>
                    </template>
                    <div class="limit-rules-section">
                        <div
                            class="limit-rules-table"
                            v-if="formData.sliding_windows && formData.sliding_windows.length > 0">
                            <div class="limit-rules-header">
                                <div class="col-time-window">{{ $t('pages.limit.form.limitRules.timeWindow') }}</div>
                                <div class="col-threshold">{{ $t('pages.limit.form.limitRules.threshold') }}</div>
                                <div class="col-rule-action">{{ $t('pages.limit.form.conditions.action') }}</div>
                            </div>
                            <div
                                class="limit-rules-row"
                                v-for="(window, index) in formData.sliding_windows"
                                :key="index">
                                <div class="col-time-window">
                                    <div class="time-window-input">
                                        <a-input-number
                                            v-model:value="window.timeWindowValue"
                                            :placeholder="$t('pages.limit.form.limitRules.timeWindow.placeholder')"
                                            :min="1"
                                            size="small"
                                            style="flex: 1" />
                                        <a-select
                                            v-model:value="window.timeWindowUnit"
                                            size="small"
                                            style="width: 80px">
                                            <a-select-option value="second">{{
                                                $t('pages.limit.form.limitRules.unit.second')
                                            }}</a-select-option>
                                            <a-select-option value="millisecond">{{
                                                $t('pages.limit.form.limitRules.unit.millisecond')
                                            }}</a-select-option>
                                            <a-select-option value="minute">{{
                                                $t('pages.limit.form.limitRules.unit.minute')
                                            }}</a-select-option>
                                        </a-select>
                                    </div>
                                </div>
                                <div class="col-threshold">
                                    <div class="threshold-input">
                                        <a-input-number
                                            v-model:value="window.threshold"
                                            :placeholder="$t('pages.limit.form.limitRules.threshold.placeholder')"
                                            :min="1"
                                            size="small"
                                            style="flex: 1" />
                                        <span class="threshold-unit">{{
                                            $t('pages.limit.form.limitRules.thresholdUnit')
                                        }}</span>
                                    </div>
                                </div>
                                <div class="col-rule-action">
                                    <minus-circle-outlined
                                        class="condition-remove-btn"
                                        @click="removeSlidingWindow(index)" />
                                </div>
                            </div>
                        </div>
                        <a
                            class="add-condition-link"
                            @click="addSlidingWindow">
                            {{ $t('pages.limit.form.limitRules.add') }}
                        </a>
                    </div>
                </a-form-item>
            </template>

            <!-- 限流规则 (Concurrency) -->
            <template v-if="formData.type === 'Concurrency'">
                <a-form-item
                    name="max_concurrency"
                    class="limit-rules-form-item">
                    <template #label>
                        <span>
                            {{ $t('pages.limit.form.limitRules') }}
                            <a-tooltip :title="$t('pages.limit.form.limitRules.tooltip')">
                                <question-circle-outlined style="margin-left: 4px; color: #999" />
                            </a-tooltip>
                        </span>
                    </template>
                    <div class="limit-rules-section">
                        <div class="scheme-row">
                            <span class="scheme-label required">{{
                                $t('pages.limit.form.limitRules.maxConcurrency')
                            }}</span>
                            <div class="scheme-input">
                                <a-input-number
                                    v-model:value="formData.max_concurrency"
                                    :min="1"
                                    style="width: 100%"
                                    :placeholder="$t('pages.limit.form.limitRules.maxConcurrency.placeholder')" />
                            </div>
                        </div>
                    </div>
                </a-form-item>
            </template>

            <!-- 限流规则 (Load) -->
            <template v-if="formData.type === 'Load'">
                <a-form-item
                    name="cpu_usage"
                    class="limit-rules-form-item">
                    <template #label>
                        <span>
                            {{ $t('pages.limit.form.limitRules') }}
                            <a-tooltip :title="$t('pages.limit.form.limitRules.tooltip')">
                                <question-circle-outlined style="margin-left: 4px; color: #999" />
                            </a-tooltip>
                        </span>
                    </template>
                    <div class="limit-rules-section">
                        <div class="scheme-row">
                            <span class="scheme-label">{{ $t('pages.limit.form.limitRules.cpuUsage') }}</span>
                            <div class="scheme-input">
                                <a-input-number
                                    v-model:value="formData.cpu_usage"
                                    :min="1"
                                    :max="100"
                                    style="width: 100%"
                                    :placeholder="$t('pages.limit.form.limitRules.cpuUsage.placeholder')" />
                                <span class="scheme-unit">%</span>
                            </div>
                        </div>
                        <div class="scheme-row">
                            <span class="scheme-label">
                                {{ $t('pages.limit.form.limitRules.loadUsage') }}
                                <a-tooltip :title="$t('pages.limit.form.limitRules.loadUsage.tooltip')">
                                    <question-circle-outlined style="margin-left: 4px; color: #999" />
                                </a-tooltip>
                            </span>
                            <div class="scheme-input">
                                <a-input-number
                                    v-model:value="formData.load_usage"
                                    :min="1"
                                    style="width: 100%"
                                    :placeholder="$t('pages.limit.form.limitRules.loadUsage.placeholder')" />
                            </div>
                        </div>
                    </div>
                </a-form-item>
            </template>

            <!-- 限流方案 -->
            <a-form-item class="limit-scheme-form-item">
                <template #label>
                    <span>
                        {{ $t('pages.limit.form.limitScheme') }}
                        <a-tooltip :title="$t('pages.limit.form.limitScheme.tooltip')">
                            <question-circle-outlined style="margin-left: 4px; color: #999" />
                        </a-tooltip>
                    </span>
                </template>
                <div class="limit-scheme-section">
                    <!-- 限流算法 -->
                    <template v-if="formData.type === 'Rate'">
                        <div class="scheme-row">
                            <span class="scheme-label required">{{
                                $t('pages.limit.form.limitScheme.algorithm')
                            }}</span>
                            <div class="scheme-buttons">
                                <span
                                    v-for="algo in algorithmOptions"
                                    :key="algo.value"
                                    :class="['scheme-btn', { active: formData.realize_type === algo.value }]"
                                    @click="formData.realize_type = algo.value">
                                    {{ algo.label }}
                                </span>
                            </div>
                        </div>
                        <!-- 漏桶容量 -->
                        <div
                            class="scheme-row"
                            v-if="formData.realize_type === 'LeakyBucket'">
                            <span class="scheme-label required">{{ $t('pages.limit.form.limitScheme.capacity') }}</span>
                            <div class="scheme-input">
                                <a-input-number
                                    v-model:value="formData.capacity"
                                    :min="0"
                                    style="width: 100%" />
                            </div>
                        </div>
                        <!-- 最大突发秒数 -->
                        <div
                            class="scheme-row"
                            v-if="formData.realize_type === 'SmoothBursty'">
                            <span class="scheme-label required">{{
                                $t('pages.limit.form.limitScheme.maxBurstSeconds')
                            }}</span>
                            <div class="scheme-input">
                                <a-input-number
                                    v-model:value="formData.maxBurstSeconds"
                                    :min="0"
                                    style="width: 100%" />
                                <span class="scheme-unit">{{
                                    $t('pages.limit.form.limitScheme.maxBurstSeconds.unit')
                                }}</span>
                            </div>
                        </div>
                        <!-- 预热秒数 -->
                        <div
                            class="scheme-row"
                            v-if="formData.realize_type === 'SmoothWarmup'">
                            <span class="scheme-label required">{{
                                $t('pages.limit.form.limitScheme.warmupPeriodSeconds')
                            }}</span>
                            <div class="scheme-input">
                                <a-input-number
                                    v-model:value="formData.warmupSeconds"
                                    :min="0"
                                    style="width: 100%" />
                                <span class="scheme-unit">{{
                                    $t('pages.limit.form.limitScheme.warmupPeriodSeconds.unit')
                                }}</span>
                            </div>
                        </div>
                        <!-- 冷启动因子 -->
                        <div
                            class="scheme-row"
                            v-if="formData.realize_type === 'SmoothWarmup'">
                            <span class="scheme-label required">{{
                                $t('pages.limit.form.limitScheme.coldFactor')
                            }}</span>
                            <div class="scheme-input">
                                <a-input-number
                                    v-model:value="formData.coldFactor"
                                    :min="1"
                                    style="width: 100%" />
                            </div>
                        </div>
                    </template>
                    <!-- 限流效果 -->
                    <div class="scheme-row">
                        <span class="scheme-label required">{{ $t('pages.limit.form.limitScheme.effect') }}</span>
                        <div class="scheme-buttons">
                            <span
                                v-for="eff in effectOptions"
                                :key="eff.value"
                                :class="['scheme-btn', { active: formData.effect === eff.value }]"
                                @click="formData.effect = eff.value">
                                {{ eff.label }}
                            </span>
                        </div>
                    </div>
                    <!-- 最大排队时长 -->
                    <div
                        class="scheme-row"
                        v-if="formData.effect === 'queuing'">
                        <span class="scheme-label required">{{ $t('pages.limit.form.limitScheme.maxQueueTime') }}</span>
                        <div class="scheme-input">
                            <a-input-number
                                v-model:value="formData.max_queue_time_seconds"
                                :min="0"
                                style="width: 120px" />
                            <span class="scheme-unit">{{ $t('pages.limit.form.limitScheme.maxQueueTime.unit') }}</span>
                        </div>
                    </div>
                </div>
            </a-form-item>

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
    method: {
        validator: (rule, value) => {
            if (value && !formData.value.path) {
                return Promise.reject(t('pages.limit.form.method.requirePath'))
            }
            return Promise.resolve()
        },
    },
    max_concurrency: {
        validator: (rule, value) => {
            if (formData.value.type === 'Concurrency') {
                if (!value) return Promise.reject(t('pages.limit.form.limitRules.maxConcurrency.placeholder'))
            }
            return Promise.resolve()
        },
    },
    cpu_usage: {
        validator: (rule, value) => {
            if (formData.value.type === 'Load') {
                if (!value && !formData.value.load_usage) return Promise.reject('负载条件不能为空')
                if (value && value > 100) return Promise.reject('CPU使用率不能超过100')
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

const algorithmOptions = computed(() => [
    { value: 'Resilience4j', label: t('pages.limit.form.limitScheme.algorithm.tokenBucket') },
    { value: 'LeakyBucket', label: t('pages.limit.form.limitScheme.algorithm.leakyBucket') },
    { value: 'SmoothBursty', label: t('pages.limit.form.limitScheme.algorithm.smoothBursting') },
    { value: 'SmoothWarmup', label: t('pages.limit.form.limitScheme.algorithm.smoothWarmUp') },
    { value: 'Redis', label: t('pages.limit.form.limitScheme.algorithm.clusterLimit') },
])

const effectOptions = computed(() => {
    const options = [{ value: 'failFast', label: t('pages.limit.form.limitScheme.effect.failFast') }]
    if (formData.value.type !== 'Load') {
        options.push({ value: 'queuing', label: t('pages.limit.form.limitScheme.effect.queuing') })
    }
    return options
})

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
    return { threshold: undefined, timeWindowValue: undefined, timeWindowUnit: 'second' }
}

// 将 timeWindowValue + timeWindowUnit 转换为 timeWindowInMs
function convertToMs(value, unit) {
    if (!value) return 0
    switch (unit) {
        case 'second':
            return value * 1000
        case 'millisecond':
            return value
        case 'minute':
            return value * 60000
        default:
            return value * 1000
    }
}

// 将 timeWindowInMs 转换为 timeWindowValue + timeWindowUnit
function convertFromMs(ms) {
    if (!ms) return { timeWindowValue: undefined, timeWindowUnit: 'second' }
    if (ms % 60000 === 0) return { timeWindowValue: ms / 60000, timeWindowUnit: 'minute' }
    if (ms % 1000 === 0) return { timeWindowValue: ms / 1000, timeWindowUnit: 'second' }
    return { timeWindowValue: ms, timeWindowUnit: 'millisecond' }
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
    formData.value.realize_type = 'Resilience4j'
    formData.value.effect = 'failFast'
    formData.value.max_queue_time_seconds = 1
    formData.value.capacity = 0
    formData.value.maxBurstSeconds = 1
    formData.value.warmupSeconds = 5
    formData.value.coldFactor = 3
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
    if (Array.isArray(cloned.sliding_windows)) {
        cloned.sliding_windows = cloned.sliding_windows.map((w) => {
            const { timeWindowValue, timeWindowUnit } = convertFromMs(w.timeWindowInMs)
            return { ...w, timeWindowValue, timeWindowUnit }
        })
    }
    if (!Array.isArray(cloned.sliding_windows) || cloned.sliding_windows.length === 0) {
        if (cloned.type === 'Rate') {
            cloned.sliding_windows = [createEmptySlidingWindow()]
        }
    }
    // 解析 action_parameters 中的 algorithm/effect
    if (cloned.action_parameters) {
        const ap =
            typeof cloned.action_parameters === 'string'
                ? JSON.parse(cloned.action_parameters)
                : cloned.action_parameters
        cloned.effect = cloned.max_wait_ms > 0 ? 'queuing' : 'failFast'
        cloned.max_queue_time_seconds = cloned.max_wait_ms > 0 ? cloned.max_wait_ms / 1000 : 1
        cloned.capacity = ap.capacity || 0
        cloned.maxBurstSeconds = ap.maxBurstSeconds || 1
        cloned.warmupSeconds = ap.warmupSeconds || 5
        cloned.coldFactor = ap.coldFactor || 3
    } else {
        cloned.effect = cloned.max_wait_ms > 0 ? 'queuing' : 'failFast'
        cloned.max_queue_time_seconds = cloned.max_wait_ms > 0 ? cloned.max_wait_ms / 1000 : 1
        cloned.capacity = 0
        cloned.maxBurstSeconds = 1
        cloned.warmupSeconds = 5
        cloned.coldFactor = 3
    }
    formData.value = cloned
}

function handleOk() {
    formRef.value
        .validateFields()
        .then(async (values) => {
            try {
                showLoading()
                // 将 sliding_windows 转换回 timeWindowInMs 格式
                let slidingWindows = undefined
                if (formData.value.sliding_windows) {
                    slidingWindows = formData.value.sliding_windows.map((w) => ({
                        threshold: w.threshold,
                        timeWindowInMs: convertToMs(w.timeWindowValue, w.timeWindowUnit),
                    }))
                }
                const actionParameters = {}

                if (formData.value.realize_type === 'LeakyBucket') {
                    actionParameters.capacity = formData.value.capacity ? String(formData.value.capacity) : '0'
                } else if (formData.value.realize_type === 'SmoothBursty') {
                    actionParameters.maxBurstSeconds = formData.value.maxBurstSeconds
                        ? String(formData.value.maxBurstSeconds)
                        : '1'
                } else if (formData.value.realize_type === 'SmoothWarmup') {
                    actionParameters.warmupSeconds = formData.value.warmupSeconds
                        ? String(formData.value.warmupSeconds)
                        : '5'
                    actionParameters.coldFactor = formData.value.coldFactor ? String(formData.value.coldFactor) : '3'
                }

                const maxWaitMs =
                    formData.value.effect === 'queuing' ? (formData.value.max_queue_time_seconds || 0) * 1000 : 0

                const params = {
                    ...values,
                    realize_type: formData.value.realize_type,
                    max_wait_ms: maxWaitMs,
                    group: formData.value.group || 'default',
                    conditions: formData.value.conditions || undefined,
                    sliding_windows: slidingWindows || undefined,
                    action_parameters: actionParameters,
                    space_code: formData.value.space_code,
                    source_application_id: formData.value.source_application_id,
                    target_service_id: formData.value.target_service_id,
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

.limit-rules-form-item,
.limit-scheme-form-item {
    :deep(.ant-form-item-control-input-content) {
        overflow: visible;
    }
}

.limit-rules-section {
    border: 1px solid #f0f0f0;
    border-radius: 6px;
    padding: 16px;
    background: #fafafa;
}

.limit-rules-header {
    display: flex;
    align-items: center;
    padding: 8px 0;
    border-bottom: 1px solid #e8e8e8;
    font-weight: 500;
    font-size: 13px;
    color: #333;
    gap: 8px;
}

.limit-rules-row {
    display: flex;
    align-items: center;
    padding: 8px 0;
    border-bottom: 1px solid #f0f0f0;
    gap: 8px;

    &:last-child {
        border-bottom: none;
    }
}

.col-time-window {
    flex: 1;
}

.col-threshold {
    flex: 1;
}

.col-rule-action {
    flex: 0 0 40px;
    text-align: center;
}

.time-window-input,
.threshold-input {
    display: flex;
    align-items: center;
    gap: 4px;
}

.threshold-unit {
    color: #666;
    font-size: 13px;
    white-space: nowrap;
}

.limit-scheme-section {
    padding: 0;
}

.scheme-row {
    display: flex;
    align-items: center;
    margin-bottom: 16px;

    &:last-child {
        margin-bottom: 0;
    }
}

.scheme-label {
    font-size: 13px;
    color: #333;
    white-space: nowrap;
    margin-right: 12px;
    min-width: 80px;

    &.required::before {
        content: '* ';
        color: #ff4d4f;
    }
}

.scheme-buttons {
    display: flex;
    gap: 0;
    flex-wrap: wrap;
}

.scheme-btn {
    display: inline-block;
    padding: 4px 16px;
    border: 1px solid #d9d9d9;
    background: #fff;
    cursor: pointer;
    font-size: 13px;
    color: #333;
    transition: all 0.2s;
    margin-left: -1px;

    &:first-child {
        border-radius: 4px 0 0 4px;
        margin-left: 0;
    }

    &:last-child {
        border-radius: 0 4px 4px 0;
    }

    &.active {
        color: #1890ff;
        border-color: #1890ff;
        background: #e6f7ff;
        z-index: 1;
        position: relative;
    }

    &:hover:not(.active) {
        color: #1890ff;
        border-color: #1890ff;
        z-index: 1;
        position: relative;
    }
}

.scheme-input {
    display: flex;
    align-items: center;
    gap: 8px;
}

.scheme-unit {
    color: #666;
    font-size: 13px;
}
</style>
