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
                :label="$t('pages.service.alias.form.alias')"
                name="alias">
                <a-input
                    :placeholder="$t('pages.service.alias.form.alias.placeholder')"
                    v-model:value="formData.alias" />
            </a-form-item>
            <a-form-item
                :label="$t('pages.service.alias.form.description')"
                name="description">
                <a-textarea
                    :placeholder="$t('pages.service.alias.form.description.placeholder')"
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
    spaceCode: { type: String, default: '' },
})

const emit = defineEmits(['ok'])
const { modal, showModal, hideModal, showLoading, hideLoading } = useModal()
const { formRecord, formData, formRef, formRules, resetForm } = useForm()
const { t } = useI18n()
const cancelText = ref(t('button.cancel'))
const okText = ref(t('button.confirm'))

formRules.value = {
    alias: { required: true, message: t('pages.service.alias.form.alias.required') },
}

function handleCreate() {
    showModal({
        type: 'create',
        title: t('pages.service.alias.create'),
    })
}

async function handleEdit(record = {}) {
    showModal({
        type: 'edit',
        title: t('pages.service.alias.edit'),
    })
    const { data, success } = await apis.service.getServiceAlias(record.id).catch(() => ({}))
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
                const params = { ...values, serviceId: props.serviceId, space_code: props.space_code }
                let result = null
                switch (modal.value.type) {
                    case 'create':
                        result = await apis.service.createServiceAlias(params).catch(() => {
                            throw new Error()
                        })
                        break
                    case 'edit':
                        result = await apis.service.updateServiceAlias(formData.value.id, params).catch(() => {
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
