<template>
    <a-modal
        :open="modal.open"
        :title="modal.title"
        :width="640"
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
            :label-col="{ style: { width: '110px' } }">
            <a-card class="mb-8-2">
                <a-row :gutter="12">
                    <a-col :span="12">
                        <a-form-item
                            :label="$t('pages.service.form.name')"
                            name="name">
                            <a-input v-model:value="formData.name"></a-input>
                        </a-form-item>
                    </a-col>

                    <a-col :span="12">
                        <a-form-item
                            :label="$t('pages.service.form.space_code')"
                            name="space_code">
                            <a-select
                                v-model:value="formData.space_code"
                                show-search
                                :filter-option="filterSpaceOption">
                                <a-select-option
                                    v-for="item in props.spaceOptions"
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
                    </a-col>

                    <a-col :span="12">
                        <a-form-item
                            :label="$t('pages.service.form.registration_type')"
                            name="registration_type">
                            <a-select v-model:value="formData.registration_type">
                                <a-select-option value="HTTP">
                                    {{ $t('pages.service.form.registration_type.http') }}
                                </a-select-option>
                                <a-select-option value="RPC_APP">
                                    {{ $t('pages.service.form.registration_type.rpc_app') }}
                                </a-select-option>
                                <a-select-option value="RPC_INTERFACE">
                                    {{ $t('pages.service.form.registration_type.rpc_interface') }}
                                </a-select-option>
                            </a-select>
                        </a-form-item>
                    </a-col>
                </a-row>

                <!-- <a-row :gutter="12">
                    <a-col :span="12">
                        <a-form-item
                            :label="$t('pages.service.form.source')"
                            name="source">
                            <a-input v-model:value="formData.source"></a-input>
                        </a-form-item>
                    </a-col>
                </a-row> -->

                <a-row :gutter="24">
                    <a-col :span="24">
                        <a-form-item
                            :label="$t('pages.service.form.description')"
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
import { ref } from 'vue'
import { config } from '@/config'
import apis from '@/apis'
import { useForm, useModal } from '@/hooks'
import { useI18n } from 'vue-i18n'

const props = defineProps({
    spaceOptions: {
        type: Array,
        default: () => [],
    },
    applicationOptions: {
        type: Array,
        default: () => [],
    },
})
const emit = defineEmits(['ok'])
const { modal, showModal, hideModal, showLoading, hideLoading } = useModal()
const { formRecord, formData, formRef, formRules, resetForm } = useForm()
const { t } = useI18n()
const cancelText = ref(t('button.cancel'))
const okText = ref(t('button.confirm'))

const SPACE_CODE_KEY = 'service_space_code'

function filterSpaceOption(input, option) {
    const label = option.children?.[0]?.children || ''
    return option.value.toLowerCase().includes(input.toLowerCase()) || label.toLowerCase().includes(input.toLowerCase())
}

function filterApplicationOption(input, option) {
    const label = option.children?.[0]?.children || ''
    return option.value.toLowerCase().includes(input.toLowerCase()) || label.toLowerCase().includes(input.toLowerCase())
}
formRules.value = {
    name: { required: true, message: t('pages.service.form.name.placeholder') },
    space_code: { required: true, message: t('pages.service.form.space_code.placeholder') },
    application_id: { required: true, message: t('pages.service.form.application_id.placeholder') },
    registration_type: { required: true, message: t('pages.service.form.registration_type.placeholder') },
}

function handleCreate() {
    showModal({
        type: 'create',
        title: t('pages.service.create'),
    })
    formData.value.registration_type = 'HTTP'
    const saved = localStorage.getItem(SPACE_CODE_KEY)
    if (saved && props.spaceOptions.some((item) => item.code === saved)) {
        formData.value.space_code = saved
    } else if (props.spaceOptions.length > 0) {
        formData.value.space_code = props.spaceOptions[0].code
    }
}

async function handleEdit(record = {}) {
    showModal({
        type: 'edit',
        title: t('pages.service.edit'),
    })

    const { data, success } = await apis.service.getService(record.id).catch()
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
                }
                let result = null
                switch (modal.value.type) {
                    case 'create':
                        result = await apis.service.createService(params).catch(() => {
                            throw new Error()
                        })
                        break
                    case 'edit':
                        result = await apis.service.updateService(formData.value.id, params).catch(() => {
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
