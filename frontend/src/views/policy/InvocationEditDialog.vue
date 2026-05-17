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
            :label-col="{ style: { width: '120px' } }"
            :wrapper-col="{ flex: 1 }">
            <!-- 规则名称 -->
            <a-form-item
                :label="$t('pages.invocation.form.name')"
                name="name">
                <a-input
                    :placeholder="$t('pages.invocation.form.name.placeholder')"
                    v-model:value="formData.name"
                    :maxlength="60" />
            </a-form-item>

            <!-- 服务空间 -->
            <a-form-item
                :label="$t('pages.invocation.form.space_code')"
                name="space_code">
                <a-select
                    :placeholder="$t('pages.invocation.form.space_code.placeholder')"
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
                :label="$t('pages.invocation.form.sourceApplicationId')"
                name="source_application_id">
                <a-select
                    :placeholder="$t('pages.invocation.form.sourceApplicationId.placeholder')"
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
                :label="$t('pages.invocation.form.targetServiceId')"
                name="target_service_id">
                <a-select
                    :placeholder="$t('pages.invocation.form.targetServiceId.placeholder')"
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
                :label="$t('pages.invocation.form.group')"
                name="group">
                <a-select
                    :placeholder="$t('pages.invocation.form.group.placeholder')"
                    v-model:value="formData.group">
                    <a-select-option value="default">
                        {{ $t('pages.invocation.form.group.default') }}
                    </a-select-option>
                </a-select>
            </a-form-item>

            <!-- 当前服务路径 -->
            <a-form-item
                :label="$t('pages.invocation.form.path')"
                name="path">
                <a-input
                    :placeholder="$t('pages.invocation.form.path.placeholder')"
                    v-model:value="formData.path" />
            </a-form-item>

            <!-- 当前服务方法 -->
            <a-form-item
                :label="$t('pages.invocation.form.method')"
                name="method">
                <a-input
                    :placeholder="$t('pages.invocation.form.method.placeholder')"
                    v-model:value="formData.method"
                    :disabled="!formData.path" />
            </a-form-item>

            <!-- 调用类别 -->
            <a-form-item
                :label="$t('pages.invocation.form.type')"
                name="type">
                <a-radio-group
                    v-model:value="formData.type"
                    @change="onTypeChange">
                    <a-radio value="failfast">{{ $t('pages.invocation.form.type.failfast') }}</a-radio>
                    <a-radio value="failover">{{ $t('pages.invocation.form.type.failover') }}</a-radio>
                    <a-radio value="failsafe">{{ $t('pages.invocation.form.type.failsafe') }}</a-radio>
                </a-radio-group>
            </a-form-item>

            <!-- 错误判断条件 (仅 failover) -->
            <template v-if="formData.type === 'failover'">
                <a-form-item
                    :label="$t('pages.invocation.form.errorCondition')"
                    name="errorCodes"
                    class="error-condition-form-item">
                    <template #label>
                        <span>
                            {{ $t('pages.invocation.form.errorCondition') }}
                            <a-tooltip :title="$t('pages.invocation.form.errorCondition.tooltip')">
                                <question-circle-outlined style="margin-left: 4px; color: #999" />
                            </a-tooltip>
                        </span>
                    </template>
                    <div class="error-condition-section">
                        <!-- 错误码 -->
                        <div class="condition-row">
                            <span class="condition-label">
                                {{ $t('pages.invocation.form.errorCodes') }}
                                <a-tooltip :title="$t('pages.invocation.form.errorCodes.tooltip')">
                                    <question-circle-outlined style="margin-left: 4px; color: #999" />
                                </a-tooltip>
                            </span>
                            <span class="condition-prefix">{{ $t('pages.invocation.form.errorCodes.include') }}</span>
                            <a-select
                                v-model:value="formData.errorCodes"
                                mode="tags"
                                :placeholder="$t('pages.invocation.form.errorCodes.placeholder')"
                                style="flex: 1"
                                :max-tag-count="5" />
                            <span class="condition-jsonpath">
                                （<span class="jsonpath-label">{{
                                    $t('pages.invocation.form.codePolicy.prefix')
                                }}</span>
                                <a-input
                                    v-model:value="formData.codePolicyExpression"
                                    :placeholder="$t('pages.invocation.form.codePolicy.placeholder')"
                                    size="small"
                                    class="jsonpath-input" />
                                <span class="jsonpath-label">{{ $t('pages.invocation.form.codePolicy.suffix') }}</span
                                >）
                            </span>
                            <span class="condition-suffix">{{ $t('pages.invocation.form.markAsError') }}</span>
                        </div>
                        <!-- 错误消息 -->
                        <div class="condition-row">
                            <span class="condition-label">{{ $t('pages.invocation.form.errorMessages') }}</span>
                            <span class="condition-prefix">{{
                                $t('pages.invocation.form.errorMessages.include')
                            }}</span>
                            <a-select
                                v-model:value="formData.errorMessages"
                                mode="tags"
                                :placeholder="$t('pages.invocation.form.errorMessages.placeholder')"
                                style="flex: 1"
                                :max-tag-count="5" />
                            <span class="condition-jsonpath">
                                （<span class="jsonpath-label">{{
                                    $t('pages.invocation.form.messagePolicy.prefix')
                                }}</span>
                                <a-input
                                    v-model:value="formData.messagePolicyExpression"
                                    :placeholder="$t('pages.invocation.form.messagePolicy.placeholder')"
                                    size="small"
                                    class="jsonpath-input" />
                                <span class="jsonpath-label">{{
                                    $t('pages.invocation.form.messagePolicy.suffix')
                                }}</span
                                >）
                            </span>
                            <span class="condition-suffix">{{ $t('pages.invocation.form.markAsError') }}</span>
                        </div>
                        <!-- 异常类 -->
                        <div class="condition-row">
                            <span class="condition-label">{{ $t('pages.invocation.form.exceptions') }}</span>
                            <span class="condition-prefix">{{ $t('pages.invocation.form.exceptions.include') }}</span>
                            <a-select
                                v-model:value="formData.exceptions"
                                mode="tags"
                                :placeholder="$t('pages.invocation.form.exceptions.placeholder')"
                                style="flex: 1"
                                :max-tag-count="5" />
                            <span class="condition-suffix">{{ $t('pages.invocation.form.markAsError') }}</span>
                        </div>
                    </div>
                </a-form-item>

                <!-- 重试策略 -->
                <a-form-item
                    :label="$t('pages.invocation.form.retryPolicy')"
                    class="retry-policy-form-item">
                    <div class="retry-policy-section">
                        <!-- 重试次数 -->
                        <div class="scheme-row">
                            <span class="scheme-label required">{{
                                $t('pages.invocation.form.retryPolicy.retry')
                            }}</span>
                            <div class="scheme-input">
                                <a-input-number
                                    v-model:value="formData.retry"
                                    :min="1"
                                    style="width: 100%"
                                    :placeholder="$t('pages.invocation.form.retryPolicy.retry.placeholder')" />
                                <span class="scheme-unit">{{
                                    $t('pages.invocation.form.retryPolicy.retry.unit')
                                }}</span>
                            </div>
                        </div>
                        <!-- 重试间隔 -->
                        <div class="scheme-row">
                            <span class="scheme-label required">{{
                                $t('pages.invocation.form.retryPolicy.interval')
                            }}</span>
                            <div class="scheme-input">
                                <a-input-number
                                    v-model:value="formData.interval"
                                    :min="0"
                                    style="width: 100%"
                                    :placeholder="$t('pages.invocation.form.retryPolicy.interval.placeholder')" />
                                <span class="scheme-unit">{{
                                    $t('pages.invocation.form.retryPolicy.interval.unit')
                                }}</span>
                            </div>
                        </div>
                        <!-- 总超时时间 -->
                        <div class="scheme-row">
                            <span class="scheme-label required">
                                {{ $t('pages.invocation.form.retryPolicy.timeout') }}
                                <a-tooltip :title="$t('pages.invocation.form.retryPolicy.timeout.tooltip')">
                                    <question-circle-outlined style="margin-left: 4px; color: #999" />
                                </a-tooltip>
                            </span>
                            <div class="scheme-input">
                                <a-input-number
                                    v-model:value="formData.timeout"
                                    :min="1"
                                    style="width: 100%"
                                    :placeholder="$t('pages.invocation.form.retryPolicy.timeout.placeholder')" />
                                <span class="scheme-unit">{{
                                    $t('pages.invocation.form.retryPolicy.timeout.unit')
                                }}</span>
                            </div>
                        </div>
                        <!-- 允许重试方法 -->
                        <div class="scheme-row">
                            <span class="scheme-label">{{ $t('pages.invocation.form.methods') }}</span>
                            <a-select
                                v-model:value="formData.methods"
                                mode="tags"
                                :placeholder="$t('pages.invocation.form.methods.placeholder')"
                                style="flex: 1"
                                :max-tag-count="5" />
                        </div>
                        <!-- 允许重试方法前缀 -->
                        <div class="scheme-row">
                            <span class="scheme-label">{{ $t('pages.invocation.form.methodPrefixes') }}</span>
                            <a-select
                                v-model:value="formData.methodPrefixes"
                                mode="tags"
                                :placeholder="$t('pages.invocation.form.methodPrefixes.placeholder')"
                                style="flex: 1"
                                :max-tag-count="5" />
                        </div>
                    </div>
                </a-form-item>
            </template>

            <!-- 生效状态 -->
            <a-form-item
                :label="$t('pages.invocation.form.enabled')"
                name="enabled">
                <a-switch
                    v-model:checked="enabledSwitch"
                    :checked-children="$t('pages.invocation.form.enabled.active')"
                    :un-checked-children="$t('pages.invocation.form.enabled.inactive')" />
            </a-form-item>

            <!-- 描述 -->
            <a-form-item
                :label="$t('pages.invocation.form.description')"
                name="description">
                <a-textarea
                    v-model:value="formData.description"
                    :placeholder="$t('pages.invocation.form.description.placeholder')"
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
import { QuestionCircleOutlined } from '@ant-design/icons-vue'
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
    name: { required: true, message: t('pages.invocation.form.name.required') },
    space_code: { required: true, message: t('pages.invocation.form.space_code.required') },
    target_service_id: { required: true, message: t('pages.invocation.form.targetServiceId.required') },
    type: { required: true, message: t('pages.invocation.form.type.required') },
    method: {
        validator: (rule, value) => {
            if (value && !formData.value.path) {
                return Promise.reject(t('pages.invocation.form.method.requirePath'))
            }
            return Promise.resolve()
        },
    },
    errorCodes: {
        validator: () => {
            if (formData.value.type === 'failover') {
                const hasCodes = formData.value.errorCodes?.length > 0
                const hasMessages = formData.value.errorMessages?.length > 0
                const hasExceptions = formData.value.exceptions?.length > 0
                if (!hasCodes && !hasMessages && !hasExceptions) {
                    return Promise.reject(t('pages.invocation.form.errorCondition.required'))
                }
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
    [() => formData.value.errorCodes, () => formData.value.errorMessages, () => formData.value.exceptions],
    () => {
        formRef.value?.validateFields(['errorCodes']).catch(() => {})
    },
    { deep: true }
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

function onTypeChange() {
    formData.value.errorCodes = []
    formData.value.errorMessages = []
    formData.value.exceptions = []
    formData.value.codePolicyExpression = ''
    formData.value.messagePolicyExpression = ''
    formData.value.retry = undefined
    formData.value.interval = undefined
    formData.value.timeout = undefined
    formData.value.methods = []
    formData.value.methodPrefixes = []
}

function handleCreate() {
    formData.value.group = 'default'
    formData.value.enabled = 0
    formData.value.type = 'failfast'
    formData.value.errorCodes = []
    formData.value.errorMessages = []
    formData.value.exceptions = []
    formData.value.codePolicyExpression = ''
    formData.value.messagePolicyExpression = ''
    formData.value.methods = []
    formData.value.methodPrefixes = []
    formData.value.space_code = initSpaceCode(props.spaceOptions)
    showModal({
        type: 'create',
        title: t('pages.invocation.add'),
    })
}

async function handleEdit(record = {}) {
    showModal({
        type: 'edit',
        title: t('pages.invocation.edit'),
    })

    const { data, success } = await apis.policy.getInvocation(record.id).catch()
    if (!success) {
        message.error(t('component.message.error.save'))
        hideModal()
        return
    }
    formRecord.value = data
    const cloned = cloneDeep(data)
    // retry_policy 从后端返回可能是对象或 JSON 字符串，需要解析
    if (typeof cloned.retry_policy === 'string') {
        try {
            cloned.retry_policy = JSON.parse(cloned.retry_policy)
        } catch {
            cloned.retry_policy = null
        }
    }
    if (cloned.retry_policy) {
        cloned.retry = cloned.retry_policy.retry
        cloned.interval = cloned.retry_policy.interval
        cloned.timeout = cloned.retry_policy.timeout
        cloned.errorCodes = cloned.retry_policy.errorCodes || []
        cloned.errorMessages = cloned.retry_policy.errorMessages || []
        cloned.exceptions = cloned.retry_policy.exceptions || []
        cloned.methods = cloned.retry_policy.methods || []
        cloned.methodPrefixes = cloned.retry_policy.methodPrefixes || []
        // 解析 codePolicy 和 messagePolicy 的 JSONPath 表达式
        cloned.codePolicyExpression = cloned.retry_policy.codePolicy?.expression || ''
        cloned.messagePolicyExpression = cloned.retry_policy.messagePolicy?.expression || ''
    } else {
        cloned.errorCodes = []
        cloned.errorMessages = []
        cloned.exceptions = []
        cloned.codePolicyExpression = ''
        cloned.messagePolicyExpression = ''
        cloned.methods = []
        cloned.methodPrefixes = []
    }
    formData.value = cloned
}

function handleOk() {
    formRef.value
        .validateFields()
        .then(async (values) => {
            try {
                showLoading()
                // 构建 retry_policy
                let retryPolicy = undefined
                if (formData.value.type === 'failover') {
                    retryPolicy = {
                        retry: formData.value.retry || 0,
                        interval: formData.value.interval || 0,
                        timeout: formData.value.timeout || 0,
                        errorCodes: formData.value.errorCodes || [],
                        errorMessages: formData.value.errorMessages || [],
                        exceptions: formData.value.exceptions || [],
                        methods: formData.value.methods || [],
                        methodPrefixes: formData.value.methodPrefixes || [],
                    }
                    // 构建 codePolicy
                    if (formData.value.codePolicyExpression) {
                        retryPolicy.codePolicy = {
                            parser: 'JSON',
                            expression: formData.value.codePolicyExpression,
                        }
                    }
                    // 构建 messagePolicy
                    if (formData.value.messagePolicyExpression) {
                        retryPolicy.messagePolicy = {
                            parser: 'JSON',
                            expression: formData.value.messagePolicyExpression,
                        }
                    }
                }
                const params = {
                    ...values,
                    group: formData.value.group || 'default',
                    retry_policy: retryPolicy,
                    space_code: formData.value.space_code,
                    source_application_id: formData.value.source_application_id,
                    target_service_id: formData.value.target_service_id,
                }
                let result = null
                switch (modal.value.type) {
                    case 'create':
                        result = await apis.policy.createInvocation(params).catch(() => {
                            throw new Error()
                        })
                        break
                    case 'edit':
                        result = await apis.policy.updateInvocation(formData.value.id, params).catch(() => {
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
.error-condition-form-item,
.retry-policy-form-item {
    :deep(.ant-form-item-control-input-content) {
        overflow: visible;
    }
}

.error-condition-section {
    border: 1px solid #f0f0f0;
    border-radius: 6px;
    padding: 16px;
    background: #fafafa;
}

.condition-row {
    display: flex;
    align-items: center;
    margin-bottom: 12px;
    gap: 8px;

    &:last-child {
        margin-bottom: 0;
    }
}

.condition-label {
    font-size: 13px;
    color: #333;
    white-space: nowrap;
    min-width: 70px;
    display: flex;
    align-items: center;
}

.condition-prefix {
    font-size: 13px;
    color: #666;
    white-space: nowrap;
}

.condition-suffix {
    font-size: 13px;
    color: #666;
    white-space: nowrap;
}

.condition-jsonpath {
    display: inline-flex;
    align-items: center;
    gap: 4px;
    font-size: 13px;
    color: #666;
    white-space: nowrap;
}

.jsonpath-label {
    font-size: 13px;
    color: #666;
}

.jsonpath-input {
    width: 120px;
}

.retry-policy-section {
    border: 1px solid #f0f0f0;
    border-radius: 6px;
    padding: 16px;
    background: #fafafa;
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
    min-width: 100px;
    display: flex;
    align-items: center;

    &.required::before {
        content: '* ';
        color: #ff4d4f;
    }
}

.scheme-input {
    display: flex;
    align-items: center;
    gap: 8px;
    flex: 1;
}

.scheme-unit {
    color: #666;
    font-size: 13px;
    white-space: nowrap;
}
</style>
