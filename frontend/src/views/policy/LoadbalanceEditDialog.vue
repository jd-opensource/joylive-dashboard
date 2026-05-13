<template>
    <a-modal
        :open="modal.open"
        :title="modal.title"
        :width="720"
        :confirm-loading="modal.confirmLoading"
        :after-close="onAfterClose"
        :cancel-text="cancelText"
        :ok-text="okText"
        @ok="handleOk"
        @cancel="handleCancel">
        <a-form
            ref="formRef"
            :model="formData"
            :rules="formRules"
            :label-col="{ style: { width: '100px' } }">
            <a-card class="mb-8-2">
                <a-row :gutter="12">
                    <a-col :span="12">
                        <a-form-item
                            :label="$t('pages.loadbalance.form.name')"
                            name="name">
                            <a-input v-model:value="formData.name"></a-input>
                        </a-form-item>
                    </a-col>

                    <a-col :span="12">
                        <a-form-item
                            :label="$t('pages.loadbalance.form.space_code')"
                            name="space_code">
                            <a-select
                                :placeholder="$t('pages.loadbalance.form.space_code.placeholder')"
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
                    </a-col>
                </a-row>

                <a-row :gutter="12">
                    <a-col :span="12">
                        <a-form-item
                            :label="$t('pages.loadbalance.form.sourceApplicationId')"
                            name="sourceApplicationId">
                            <a-select
                                :placeholder="$t('pages.loadbalance.form.sourceApplicationId.placeholder')"
                                v-model:value="formData.sourceApplicationId"
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
                    </a-col>

                    <a-col :span="12">
                        <a-form-item
                            :label="$t('pages.loadbalance.form.targetServiceId')"
                            name="targetServiceId">
                            <a-select
                                :placeholder="$t('pages.loadbalance.form.targetServiceId.placeholder')"
                                v-model:value="formData.targetServiceId"
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
                    </a-col>
                </a-row>

                <a-row :gutter="12">
                    <a-col :span="12">
                        <a-form-item
                            :label="$t('pages.loadbalance.form.policyType')"
                            name="policyType">
                            <a-select
                                :placeholder="$t('pages.loadbalance.form.policyType.placeholder')"
                                v-model:value="formData.policyType">
                                <a-select-option value="RANDOM">随机策略</a-select-option>
                                <a-select-option value="ROUND_ROBIN">轮询策略</a-select-option>
                                <a-select-option value="CUSTOM_RESPONSE">自适应策略</a-select-option>
                            </a-select>
                        </a-form-item>
                    </a-col>
                </a-row>

                <a-row :gutter="12">
                    <a-col :span="12">
                        <a-form-item
                            :label="$t('pages.loadbalance.form.enabled')"
                            name="enabled">
                            <a-switch
                                v-model:checked="enabledSwitch"
                                :checked-children="$t('pages.loadbalance.form.enabled.active')"
                                :un-checked-children="$t('pages.loadbalance.form.enabled.inactive')" />
                        </a-form-item>
                    </a-col>
                </a-row>

                <a-row :gutter="24">
                    <a-col :span="24">
                        <a-form-item
                            :label="$t('pages.loadbalance.form.description')"
                            name="description">
                            <a-textarea v-model:value="formData.description"></a-textarea>
                        </a-form-item>
                    </a-col>
                </a-row>
            </a-card>
        </a-form>
    </a-modal>
</template>

<script setup>
import { cloneDeep } from 'lodash-es'
import { message } from 'ant-design-vue'
import { ref, computed, watch } from 'vue'
import { config } from '@/config'
import apis from '@/apis'
import { useForm, useModal } from '@/hooks'

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
    name: { required: true, message: t('pages.loadbalance.form.name.required') },
    spaceCode: { required: true, message: t('pages.loadbalance.form.space_code.required') },
    targetServiceId: { required: true, message: t('pages.loadbalance.form.targetServiceId.required') },
    policyType: { required: true, message: t('pages.loadbalance.form.policyType.required') },
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
    (_val, oldVal) => {
        if (oldVal !== undefined) {
            formData.value.targetServiceId = undefined
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

function handleCreate() {
    formData.value.group = 'default'
    formData.value.enabled = 0
    showModal({
        type: 'create',
        title: t('pages.loadbalance.add'),
    })
}

async function handleEdit(record = {}) {
    showModal({
        type: 'edit',
        title: t('pages.loadbalance.edit'),
    })

    const { data, success } = await apis.policy.getLoadbalance(record.id).catch()
    if (!success) {
        message.error(t('component.message.error.save'))
        hideModal()
        return
    }
    formRecord.value = data
    formData.value = cloneDeep(data)
}

function handleOk() {
    formRef.value
        .validateFields()
        .then(async (values) => {
            try {
                showLoading()
                const params = {
                    ...values,
                    group: 'default',
                }
                let result = null
                switch (modal.value.type) {
                    case 'create':
                        result = await apis.policy.createLoadbalance(params).catch(() => {
                            throw new Error()
                        })
                        break
                    case 'edit':
                        result = await apis.policy.updateLoadbalance(formData.value.id, params).catch(() => {
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

<style lang="less" scoped></style>
