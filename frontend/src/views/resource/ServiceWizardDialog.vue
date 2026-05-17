<template>
    <a-modal
        :open="modal.open"
        :title="$t('pages.service.wizard.title', '服务接入')"
        :width="720"
        :confirm-loading="modal.confirmLoading"
        :after-close="onAfterClose"
        :cancel-text="$t('button.cancel', '取消')"
        :ok-text="$t('button.confirm', '确定')"
        @ok="handleOk"
        @cancel="handleCancel">
        <a-form
            ref="formRef"
            :model="formData"
            :label-col="{ style: { width: '160px' } }"
            :wrapper-col="{ span: 16 }">
            <a-form-item
                :label="$t('pages.service.wizard.group', '服务分组')"
                name="group">
                <a-select
                    v-model:value="formData.group"
                    style="width: 240px"
                    :loading="groupLoading">
                    <a-select-option
                        v-for="item in groupOptions"
                        :key="item.code"
                        :value="item.code">
                        {{ item.name }}
                    </a-select-option>
                </a-select>
            </a-form-item>

            <a-form-item :label="$t('pages.service.wizard.agentRegistration', 'Agent注册发现')">
                <a-switch v-model:checked="formData.enableAgentRegistration" />
            </a-form-item>

            <a-form-item :label="$t('pages.service.wizard.agentConfig', 'Agent自动配置')">
                <a-switch v-model:checked="formData.enableAgentConfig" />
            </a-form-item>

            <!-- <a-form-item :label="$t('pages.service.wizard.grayLane', '灰度泳道')">
                <a-switch v-model:checked="formData.enableGrayLane" />
                <div class="tip-text">
                    {{
                        $t(
                            'pages.service.wizard.grayLaneTip',
                            '泳道为相同版本应用定义的一套隔离环境。系统将为部署到该灰度泳道的服务实例自动打上灰度泳道标识，支持灰度泳道流量闭环能力。'
                        )
                    }}
                </div>
            </a-form-item> -->

            <a-form-item :label="$t('pages.service.wizard.deployTarget', '部署目标')">
                <a-radio-group v-model:value="formData.deployTarget">
                    <a-radio value="k8s">{{ $t('pages.service.wizard.deployTarget.k8s', 'K8S集群') }}</a-radio>
                    <!-- <a-radio value="vm">{{ $t('pages.service.wizard.deployTarget.vm', '虚拟机') }}</a-radio> -->
                </a-radio-group>
            </a-form-item>

            <a-form-item :label="$t('pages.service.wizard.labels', '服务标签')">
                <div class="code-container">
                    <pre><code class="language-yaml">{{ generatedLabels }}</code></pre>
                </div>
                <div class="tip-text mt-2">
                    {{ $t('pages.service.wizard.labelsTip', '请复制标签模板中内容到服务deploy的.metadata.labels中') }}
                </div>
            </a-form-item>
        </a-form>
    </a-modal>
</template>

<script setup>
import { ref, reactive, computed } from 'vue'
import { useModal } from '@/hooks'
import apis from '@/apis'
import { config } from '@/config'

const { modal, showModal, hideModal } = useModal()

const formData = reactive({
    group: 'default',
    enableSgm: true,
    enableAgentRegistration: false,
    enableAgentConfig: false,
    enableGrayLane: false,
    deployTarget: 'k8s',
})

const currentRecord = ref(null)
const groupOptions = ref([])
const groupLoading = ref(false)

async function loadGroupOptions(serviceId) {
    try {
        groupLoading.value = true
        const { data, success } = await apis.service
            .getServiceGroupList({
                pageSize: 99,
                current: 1,
                service_id: serviceId,
            })
            .catch(() => {
                throw new Error()
            })
        if (config('http.code.success') === success) {
            groupOptions.value = data || []
            if (groupOptions.value.length > 0 && !groupOptions.value.some((item) => item.code === formData.group)) {
                formData.group = groupOptions.value[0].code
            }
        }
    } catch (error) {
        // ignore
    } finally {
        groupLoading.value = false
    }
}

const generatedLabels = computed(() => {
    let text = `# 微服务空间标识，格式为\${微服务空间code}，默认值用于开发环境\n`
    text += `joylive.jd.com/service-space: ${currentRecord.value?.namespace || 'default'}\n`

    text += `# 服务分组，用于服务治理策略，default为默认分组\n`
    text += `joylive.jd.com/group: ${formData.group || 'default'}\n`

    // Attempt to guess the app name or just use a placeholder
    const appName = currentRecord.value?.name || 'spring-cloud-consumer-demo'
    text += `# 应用名，格式为\${微服务空间code}-\${服务名}，用于资源隔离\n`
    text += `joylive.jd.com/app: ${appName}\n`

    text += `# 服务名，需与注册中心名称一致\n`
    text += `joylive.jd.com/service: ${currentRecord.value?.name || 'service-consumer'}\n`

    text += `# 服务增强方式，agent表示Java Agent方式接入\n`
    text += `joylive.jd.com/enhance: agent\n`

    text += `# 是否启用接入，true表示开启接入流程\n`
    text += `x-live-enabled: "true"\n`

    if (formData.enableAgentRegistration) {
        text += `# 启用Agent注册发现\n`
        text += `joylive.jd.com/register: "true"\n`
    }

    if (formData.enableAgentConfig) {
        text += `# 启用Agent自动配置\n`
        text += `joylive.jd.com/agent-config: "true"\n`
    }

    if (formData.enableGrayLane) {
        text += `# 启用灰度泳道\n`
        text += `joylive.jd.com/lane: "true"\n`
    }

    return text
})

function show(record) {
    currentRecord.value = record || {}
    // Reset form to default states
    Object.assign(formData, {
        group: 'default',
        enableSgm: true,
        enableAgentRegistration: false,
        enableAgentConfig: false,
        enableGrayLane: false,
        deployTarget: 'k8s',
    })

    if (record && record.id) {
        loadGroupOptions(record.id)
    }

    showModal({
        type: 'wizard',
        title: '服务接入',
    })
}

function handleOk() {
    hideModal()
}

function handleCancel() {
    hideModal()
}

function onAfterClose() {
    currentRecord.value = null
    groupOptions.value = []
}

defineExpose({
    show,
})
</script>

<style lang="less" scoped>
.tip-text {
    font-size: 12px;
    color: #8c8c8c;
    line-height: 1.5;
    margin-top: 4px;
}
.mt-2 {
    margin-top: 8px;
}
.code-container {
    background-color: #f5f5f5;
    border: 1px solid #e8e8e8;
    border-radius: 4px;
    padding: 12px;
    overflow-x: auto;

    pre {
        margin: 0;

        code {
            font-family: Consolas, Monaco, 'Andale Mono', 'Ubuntu Mono', monospace;
            font-size: 13px;
            color: #008000;
            white-space: pre;
        }
    }
}
</style>
