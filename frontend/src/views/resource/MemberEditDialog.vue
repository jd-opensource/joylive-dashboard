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
                :label="$t('pages.member.form.user')"
                name="user">
                <a-input
                    :placeholder="$t('pages.member.form.user.placeholder')"
                    v-model:value="formData.user"
                    :disabled="modal.type === 'edit'" />
            </a-form-item>
            <a-form-item
                :label="$t('pages.member.form.tenant')"
                name="tenant">
                <a-input
                    :placeholder="$t('pages.member.form.tenant.placeholder')"
                    v-model:value="formData.tenant" />
            </a-form-item>
            <a-form-item
                :label="$t('pages.member.form.role')"
                name="role">
                <a-select
                    :placeholder="$t('pages.member.form.role.placeholder')"
                    v-model:value="formData.role">
                    <a-select-option value="owner">owner</a-select-option>
                    <a-select-option value="admin">admin</a-select-option>
                    <a-select-option value="viewer">viewer</a-select-option>
                </a-select>
            </a-form-item>
            <a-form-item
                :label="$t('pages.member.form.permission')"
                name="permission">
                <a-checkbox-group v-model:value="formData.permissions">
                    <a-checkbox value="read">{{ $t('pages.member.form.permission.read') }}</a-checkbox>
                    <a-checkbox value="write">{{ $t('pages.member.form.permission.write') }}</a-checkbox>
                    <a-checkbox value="delete">{{ $t('pages.member.form.permission.delete') }}</a-checkbox>
                </a-checkbox-group>
            </a-form-item>
        </a-form>
    </a-modal>
</template>

<script setup>
import { cloneDeep } from 'lodash-es'
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

const permToBits = { read: 1, write: 2, delete: 4 }
const bitsToPerms = (bits) => {
    const perms = []
    if (bits & 1) perms.push('read')
    if (bits & 2) perms.push('write')
    if (bits & 4) perms.push('delete')
    return perms
}
const permsToBits = (perms) => perms.reduce((bits, p) => bits | (permToBits[p] || 0), 0)

formRules.value = {
    user: { required: true, message: t('pages.member.form.user.required') },
    tenant: { required: true, message: t('pages.member.form.tenant.required') },
    role: { required: true, message: t('pages.member.form.role.required') },
}

function handleCreate() {
    showModal({
        type: 'create',
        title: t('pages.member.add'),
    })
    formData.value = { permissions: ['read'] }
}

async function handleEdit(record = {}) {
    showModal({
        type: 'edit',
        title: t('pages.member.edit'),
    })
    formRecord.value = record
    formData.value = {
        ...cloneDeep(record),
        permissions: bitsToPerms(record.permission),
    }
}

function handleOk() {
    formRef.value
        .validateFields()
        .then(async (values) => {
            try {
                showLoading()
                const params = {
                    type: 'service',
                    data_id: props.serviceId,
                    user: values.user,
                    tenant: values.tenant,
                    role: values.role,
                    permission: permsToBits(values.permissions || []),
                }
                let result = null
                switch (modal.value.type) {
                    case 'create':
                        result = await apis.service.createDataPermission(params).catch(() => {
                            throw new Error()
                        })
                        break
                    case 'edit':
                        result = await apis.service.updateDataPermission(formData.value.id, params).catch(() => {
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
