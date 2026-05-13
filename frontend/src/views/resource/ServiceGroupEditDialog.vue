<template>
    <a-modal
        :open="modal.open"
        :title="modal.title"
        :width="560"
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
            :label-col="{ style: { width: '100px' } }"
            :wrapper-col="{ flex: 1 }">
            <a-form-item
                :label="$t('pages.service.group.form.name')"
                name="name">
                <a-input
                    :placeholder="$t('pages.service.group.form.name.placeholder')"
                    v-model:value="formData.name" />
            </a-form-item>
            <a-form-item
                :label="$t('pages.service.group.form.code')"
                name="code">
                <a-input
                    :placeholder="$t('pages.service.group.form.code.placeholder')"
                    v-model:value="formData.code" />
            </a-form-item>
            <a-form-item
                :label="$t('pages.service.group.form.isolationCode')"
                name="isolationCode">
                <a-input
                    :placeholder="$t('pages.service.group.form.isolationCode.placeholder')"
                    v-model:value="formData.isolationCode" />
            </a-form-item>
            <a-form-item
                :label="$t('pages.service.group.form.labels')"
                name="labels">
                <a-input
                    :placeholder="$t('pages.service.group.form.labels.placeholder')"
                    v-model:value="formData.labels" />
            </a-form-item>
            <a-form-item
                :label="$t('pages.service.group.form.description')"
                name="description">
                <a-textarea
                    :placeholder="$t('pages.service.group.form.description.placeholder')"
                    v-model:value="formData.description"
                    :rows="3" />
            </a-form-item>
        </a-form>
    </a-modal>
</template>

<script setup>
import { cloneDeep } from 'lodash-es'
import { message } from 'ant-design-vue'
import { ref } from 'vue'
import { config } from '@/config'
import apis from '@/apis'
import { useForm, useModal } from '@/hooks'
import { useI18n } from 'vue-i18n'

const props = defineProps({
    serviceId: { type: String, default: '' },
})

const emit = defineEmits(['ok'])
const { modal, showModal, hideModal, showLoading, hideLoading } = useModal()
const { formRecord, formData, formRef, formRules, resetForm } = useForm()
const { t } = useI18n()
const cancelText = ref(t('button.cancel'))
const okText = ref(t('button.confirm'))

formRules.value = {
    name: { required: true, message: t('pages.service.group.form.name.required') },
    code: { required: true, message: t('pages.service.group.form.code.required') },
}

function handleCreate() {
    showModal({
        type: 'create',
        title: t('pages.service.group.create'),
    })
}

async function handleEdit(record = {}) {
    showModal({
        type: 'edit',
        title: t('pages.service.group.edit'),
    })
    const { data, success } = await apis.service.getServiceGroup(record.id).catch(() => ({}))
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
                const params = { ...values, service_id: props.serviceId }
                let result = null
                switch (modal.value.type) {
                    case 'create':
                        result = await apis.service.createServiceGroup(params).catch(() => {
                            throw new Error()
                        })
                        break
                    case 'edit':
                        result = await apis.service.updateServiceGroup(formData.value.id, params).catch(() => {
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
