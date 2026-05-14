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
            <a-card class="mb-8-2">
                <!-- 规则名称 -->
                <a-form-item
                    :label="$t('pages.permission.form.name')"
                    name="name">
                    <a-input
                        :placeholder="$t('pages.permission.form.name.placeholder')"
                        v-model:value="formData.name"
                        :maxlength="60" />
                </a-form-item>

                <!-- 服务空间 -->
                <a-form-item
                    :label="$t('pages.permission.form.space_code')"
                    name="space_code">
                    <a-select
                        :placeholder="$t('pages.permission.form.space_code.placeholder')"
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
                    :label="$t('pages.permission.form.sourceApplicationId')"
                    name="source_application_id">
                    <a-select
                        :placeholder="$t('pages.permission.form.sourceApplicationId.placeholder')"
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
                    :label="$t('pages.permission.form.targetServiceId')"
                    name="target_service_id">
                    <a-select
                        :placeholder="$t('pages.permission.form.targetServiceId.placeholder')"
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
                    :label="$t('pages.permission.form.group')"
                    name="group">
                    <a-select
                        :placeholder="$t('pages.permission.form.group.placeholder')"
                        v-model:value="formData.group">
                        <a-select-option value="default">
                            {{ $t('pages.permission.form.group.default') }}
                        </a-select-option>
                    </a-select>
                </a-form-item>

                <!-- 当前服务路径 -->
                <a-form-item
                    :label="$t('pages.permission.form.path')"
                    name="path">
                    <a-input
                        :placeholder="$t('pages.permission.form.path.placeholder')"
                        v-model:value="formData.path" />
                </a-form-item>

                <!-- 当前服务方法 -->
                <a-form-item
                    :label="$t('pages.permission.form.method')"
                    name="method">
                    <a-input
                        :placeholder="$t('pages.permission.form.method.placeholder')"
                        v-model:value="formData.method" />
                </a-form-item>

                <!-- 鉴权方式 -->
                <a-form-item
                    :label="$t('pages.permission.form.authMethod')"
                    name="relation_type"
                    class="auth-method-form-item">
                    <div class="auth-method-section">
                        <div class="auth-method-header">
                            <span class="auth-method-label">{{
                                $t('pages.permission.form.authMethod.settingLabel')
                            }}</span>
                            <a-radio-group v-model:value="formData.relation_type">
                                <a-radio value="AND">AND({{ $t('pages.permission.form.relationType.and') }})</a-radio>
                                <a-radio value="OR">OR({{ $t('pages.permission.form.relationType.or') }})</a-radio>
                            </a-radio-group>
                        </div>

                        <!-- 条件表格 -->
                        <div
                            class="conditions-table"
                            v-if="formData.conditions && formData.conditions.length > 0">
                            <div class="conditions-header">
                                <div class="col-type">{{ $t('pages.permission.form.conditions.type') }}</div>
                                <div class="col-key">{{ $t('pages.permission.form.conditions.key') }}</div>
                                <div class="col-op">
                                    {{ $t('pages.permission.form.conditions.opType') }}
                                    <a-tooltip :title="$t('pages.permission.form.conditions.opType.tooltip')">
                                        <question-circle-outlined style="margin-left: 4px; color: #999" />
                                    </a-tooltip>
                                </div>
                                <div class="col-value">{{ $t('pages.permission.form.conditions.values') }}</div>
                                <div class="col-action">{{ $t('pages.permission.form.conditions.action') }}</div>
                            </div>
                            <div
                                class="conditions-row"
                                v-for="(condition, index) in formData.conditions"
                                :key="index">
                                <div class="col-type">
                                    <a-select
                                        v-model:value="condition.type"
                                        :placeholder="$t('pages.permission.form.conditions.type.placeholder')"
                                        size="small">
                                        <a-select-option value="QUERY">QUERY</a-select-option>
                                        <a-select-option value="COOKIE">COOKIE</a-select-option>
                                        <a-select-option value="HEADER">HEADER</a-select-option>
                                    </a-select>
                                </div>
                                <div class="col-key">
                                    <a-input
                                        v-model:value="condition.key"
                                        :placeholder="$t('pages.permission.form.conditions.key.placeholder')"
                                        size="small" />
                                </div>
                                <div class="col-op">
                                    <a-select
                                        v-model:value="condition.opType"
                                        :placeholder="$t('pages.permission.form.conditions.opType.placeholder')"
                                        size="small">
                                        <a-select-option value="EQUAL">{{
                                            $t('pages.permission.form.conditions.op.equal')
                                        }}</a-select-option>
                                        <a-select-option value="NOT_EQUAL">{{
                                            $t('pages.permission.form.conditions.op.notEqual')
                                        }}</a-select-option>
                                        <a-select-option value="CONTAIN">{{
                                            $t('pages.permission.form.conditions.op.contain')
                                        }}</a-select-option>
                                        <a-select-option value="NOT_CONTAIN">{{
                                            $t('pages.permission.form.conditions.op.notContain')
                                        }}</a-select-option>
                                        <a-select-option value="REGEX">{{
                                            $t('pages.permission.form.conditions.op.regex')
                                        }}</a-select-option>
                                        <a-select-option value="PREFIX">{{
                                            $t('pages.permission.form.conditions.op.prefix')
                                        }}</a-select-option>
                                    </a-select>
                                </div>
                                <div class="col-value">
                                    <a-select
                                        v-model:value="condition.values"
                                        mode="tags"
                                        :placeholder="$t('pages.permission.form.conditions.values.placeholder')"
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
                            {{ $t('pages.permission.form.conditions.add') }}
                        </a>
                    </div>
                </a-form-item>

                <!-- 鉴权结果 -->
                <a-form-item
                    :label="$t('pages.permission.form.type')"
                    name="type">
                    <a-radio-group v-model:value="formData.type">
                        <a-radio value="DENY">{{ $t('pages.permission.form.type.black') }}</a-radio>
                        <a-radio value="ALLOW">{{ $t('pages.permission.form.type.white') }}</a-radio>
                    </a-radio-group>
                </a-form-item>

                <!-- 生效状态 -->
                <a-form-item
                    :label="$t('pages.permission.form.enabled')"
                    name="enabled">
                    <a-switch
                        v-model:checked="enabledSwitch"
                        :checked-children="$t('pages.permission.form.enabled.active')"
                        :un-checked-children="$t('pages.permission.form.enabled.inactive')" />
                </a-form-item>

                <!-- 描述 -->
                <a-form-item
                    :label="$t('pages.permission.form.description')"
                    name="description">
                    <a-textarea
                        v-model:value="formData.description"
                        :placeholder="$t('pages.permission.form.description.placeholder')"
                        :maxlength="255"
                        show-count
                        :rows="3" />
                </a-form-item>
            </a-card>
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
    name: { required: true, message: t('pages.permission.form.name.required') },
    space_code: { required: true, message: t('pages.permission.form.space_code.required') },
    target_service_id: { required: true, message: t('pages.permission.form.targetServiceId.required') },
    relation_type: { required: true, message: t('pages.permission.form.relationType.required') },
    type: { required: true, message: t('pages.permission.form.type.required') },
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

function handleCreate() {
    formData.value.group = 'default'
    formData.value.enabled = 0
    formData.value.relation_type = 'AND'
    formData.value.type = 'BLACK'
    formData.value.conditions = [createEmptyCondition()]
    formData.value.space_code = initSpaceCode(props.spaceOptions)
    showModal({
        type: 'create',
        title: t('pages.permission.add'),
    })
}

async function handleEdit(record = {}) {
    showModal({
        type: 'edit',
        title: t('pages.permission.edit'),
    })

    const { data, success } = await apis.policy.getPermission(record.id).catch()
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
                    conditions: formData.value.conditions ? JSON.stringify(formData.value.conditions) : undefined,
                }
                let result = null
                switch (modal.value.type) {
                    case 'create':
                        result = await apis.policy.createPermission(params).catch(() => {
                            throw new Error()
                        })
                        break
                    case 'edit':
                        result = await apis.policy.updatePermission(formData.value.id, params).catch(() => {
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

.auth-method-form-item {
    :deep(.ant-form-item-control-input-content) {
        overflow: visible;
    }
}

.auth-method-section {
    border: 1px solid #f0f0f0;
    border-radius: 6px;
    padding: 16px;
    background: #fafafa;
}

.auth-method-header {
    display: flex;
    align-items: center;
    margin-bottom: 12px;
}

.auth-method-label {
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

.conditions-header {
    gap: 8px;
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
</style>
