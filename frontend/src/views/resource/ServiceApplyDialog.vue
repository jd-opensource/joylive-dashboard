<template>
    <a-modal
        :open="modal.open"
        :title="modal.title"
        :width="520"
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
            <a-form-item
                :label="$t('pages.service.form.application_id')"
                name="application_id">
                <a-select
                    v-model:value="formData.application_id"
                    show-search
                    :placeholder="$t('pages.service.form.application_id.placeholder')"
                    :filter-option="filterApplicationOption">
                    <a-select-option
                        v-for="item in props.applicationOptions"
                        :key="item.id"
                        :value="item.id">
                        {{ item.name }}
                    </a-select-option>
                </a-select>
            </a-form-item>
            <a-form-item
                :label="$t('pages.service.form.apply_reason')"
                name="description">
                <a-textarea
                    v-model:value="formData.description"
                    :placeholder="$t('pages.service.form.apply_reason.placeholder')"
                    :rows="4" />
            </a-form-item>
        </a-form>
    </a-modal>
</template>

<script setup>
import { ref } from 'vue'
import { config } from '@/config'
import apis from '@/apis'
import { useForm, useModal } from '@/hooks'
import { useI18n } from 'vue-i18n'

const props = defineProps({
    applicationOptions: {
        type: Array,
        default: () => [],
    },
})
const emit = defineEmits(['ok'])
const { modal, showModal, hideModal, showLoading, hideLoading } = useModal()
const { formData, formRef, formRules, resetForm } = useForm()
const { t } = useI18n()
const cancelText = ref(t('button.cancel'))
const okText = ref(t('button.confirm'))

let serviceId = ''

function filterApplicationOption(input, option) {
    const label = option.children?.[0]?.children || ''
    return option.value.toLowerCase().includes(input.toLowerCase()) || label.toLowerCase().includes(input.toLowerCase())
}

formRules.value = {
    application_id: { required: true, message: t('pages.service.form.application_id.placeholder') },
    description: { required: true, message: t('pages.service.form.apply_reason.placeholder') },
}

function handleApply(record) {
    serviceId = record.id
    showModal({
        type: 'apply',
        title: t('pages.service.apply.title'),
    })
}

function handleOk() {
    formRef.value
        .validateFields()
        .then(async (values) => {
            try {
                showLoading()
                const params = {
                    ...values,
                    service_id: serviceId,
                    role: 'consumer',
                }
                const result = await apis.application_service.createApplicationService(params).catch(() => {
                    throw new Error()
                })
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
    serviceId = ''
}

defineExpose({
    handleApply,
})
</script>
