<template>
    <div class="dashboard">
        <a-card :bordered="false">
            <a-row :gutter="16">
                <a-col
                    :xs="24"
                    :sm="12"
                    :md="8"
                    :lg="4"
                    :xl="4">
                    <a-card
                        hoverable
                        class="stat-card"
                        @click="goTo('/space/list')">
                        <a-statistic
                            :title="$t('pages.dashboard.spaces')"
                            :value="counts.spaces"
                            style="cursor: pointer">
                            <template #prefix>
                                <database-outlined />
                            </template>
                        </a-statistic>
                    </a-card>
                </a-col>
                <a-col
                    :xs="24"
                    :sm="12"
                    :md="8"
                    :lg="5"
                    :xl="5">
                    <a-card
                        hoverable
                        class="stat-card"
                        @click="goTo('/resource/application')">
                        <a-statistic
                            :title="$t('pages.dashboard.applications')"
                            :value="counts.applications"
                            style="cursor: pointer">
                            <template #prefix>
                                <appstore-outlined />
                            </template>
                        </a-statistic>
                    </a-card>
                </a-col>
                <a-col
                    :xs="24"
                    :sm="12"
                    :md="8"
                    :lg="5"
                    :xl="5">
                    <a-card
                        hoverable
                        class="stat-card"
                        @click="goTo('/resource/service')">
                        <a-statistic
                            :title="$t('pages.dashboard.services')"
                            :value="counts.services"
                            style="cursor: pointer">
                            <template #prefix>
                                <cloud-server-outlined />
                            </template>
                        </a-statistic>
                    </a-card>
                </a-col>
                <a-col
                    :xs="24"
                    :sm="12"
                    :md="8"
                    :lg="5"
                    :xl="5">
                    <a-card
                        hoverable
                        class="stat-card">
                        <a-statistic
                            :title="$t('pages.dashboard.serviceGroups')"
                            :value="counts.serviceGroups">
                            <template #prefix>
                                <cluster-outlined />
                            </template>
                        </a-statistic>
                    </a-card>
                </a-col>
                <a-col
                    :xs="24"
                    :sm="12"
                    :md="8"
                    :lg="5"
                    :xl="5">
                    <a-card
                        hoverable
                        class="stat-card">
                        <a-statistic
                            :title="$t('pages.dashboard.serviceAliases')"
                            :value="counts.serviceAliases">
                            <template #prefix>
                                <branches-outlined />
                            </template>
                        </a-statistic>
                    </a-card>
                </a-col>
            </a-row>
        </a-card>

        <a-row
            :gutter="16"
            style="margin-top: 16px">
            <a-col
                :xs="24"
                :lg="12">
                <a-card
                    :title="$t('pages.dashboard.policyDistribution')"
                    :bordered="false">
                    <x-chart
                        :options="pieChartOptions"
                        height="320" />
                </a-card>
            </a-col>
            <a-col
                :xs="24"
                :lg="12">
                <a-card
                    :title="$t('pages.dashboard.resourceOverview')"
                    :bordered="false">
                    <x-chart
                        :options="barChartOptions"
                        height="320" />
                </a-card>
            </a-col>
        </a-row>

        <a-card
            :title="$t('pages.dashboard.quickLinks')"
            :bordered="false"
            style="margin-top: 16px">
            <a-row :gutter="16">
                <a-col
                    :xs="12"
                    :sm="6">
                    <a-card
                        hoverable
                        class="quick-link-card"
                        @click="goTo('/space/list')">
                        <database-outlined class="quick-link-icon" />
                        <div class="quick-link-text">{{ $t('pages.dashboard.spaces') }}</div>
                    </a-card>
                </a-col>
                <a-col
                    :xs="12"
                    :sm="6">
                    <a-card
                        hoverable
                        class="quick-link-card"
                        @click="goTo('/resource/application')">
                        <appstore-outlined class="quick-link-icon" />
                        <div class="quick-link-text">{{ $t('pages.dashboard.applications') }}</div>
                    </a-card>
                </a-col>
                <a-col
                    :xs="12"
                    :sm="6">
                    <a-card
                        hoverable
                        class="quick-link-card"
                        @click="goTo('/resource/service')">
                        <cloud-server-outlined class="quick-link-icon" />
                        <div class="quick-link-text">{{ $t('pages.dashboard.services') }}</div>
                    </a-card>
                </a-col>
                <a-col
                    :xs="12"
                    :sm="6">
                    <a-card
                        hoverable
                        class="quick-link-card"
                        @click="goTo('/policy/loadbalance')">
                        <safety-outlined class="quick-link-icon" />
                        <div class="quick-link-text">{{ $t('pages.dashboard.policies') }}</div>
                    </a-card>
                </a-col>
            </a-row>
        </a-card>
    </div>
</template>

<script setup>
import { computed, onMounted, reactive } from 'vue'
import { useRouter } from 'vue-router'
import { useI18n } from 'vue-i18n'
import {
    DatabaseOutlined,
    AppstoreOutlined,
    CloudServerOutlined,
    ClusterOutlined,
    BranchesOutlined,
    SafetyOutlined,
} from '@ant-design/icons-vue'
import apis from '@/apis'
import request from '@/utils/request'

defineOptions({
    name: 'home',
})

const router = useRouter()
const { t } = useI18n()

const counts = reactive({
    spaces: 0,
    applications: 0,
    services: 0,
    serviceGroups: 0,
    serviceAliases: 0,
})

const policyCounts = reactive({
    loadbalance: 0,
    route: 0,
    limit: 0,
    circuitBreak: 0,
    permission: 0,
    auth: 0,
    fault: 0,
    invocation: 0,
})

const POLICY_META = [
    { key: 'loadbalance', fetch: apis.policy.getLoadbalanceList },
    { key: 'route', fetch: apis.policy.getRouteList },
    { key: 'limit', fetch: apis.policy.getLimitList },
    { key: 'circuitBreak', fetch: apis.policy.getCircuitBreakList },
    { key: 'permission', fetch: apis.policy.getPermissionList },
    { key: 'auth', fetch: apis.policy.getAuthList },
    { key: 'fault', fetch: apis.policy.getFaultList },
    { key: 'invocation', fetch: apis.policy.getInvocationList },
]

const COLORS = ['#5470c6', '#91cc75', '#fac858', '#ee6666', '#73c0de', '#3ba272', '#fc8452', '#9a60b4']

onMounted(async () => {
    const params = { current: 1, pageSize: 1 }
    const [spaceRes, appRes, svcRes, groupRes, aliasRes] = await Promise.all([
        apis.space.getSpaceList(params).catch(() => ({ total: 0 })),
        apis.application.getApplicationList(params).catch(() => ({ total: 0 })),
        apis.service.getServiceList(params).catch(() => ({ total: 0 })),
        request.basic.get('/api/v1/service-groups', params).catch(() => ({ total: 0 })),
        request.basic.get('/api/v1/service-aliases', params).catch(() => ({ total: 0 })),
    ])

    counts.spaces = spaceRes.total || 0
    counts.applications = appRes.total || 0
    counts.services = svcRes.total || 0
    counts.serviceGroups = groupRes.total || 0
    counts.serviceAliases = aliasRes.total || 0

    const policyResults = await Promise.all(POLICY_META.map((p) => p.fetch(params).catch(() => ({ total: 0 }))))
    policyResults.forEach((res, i) => {
        policyCounts[POLICY_META[i].key] = res.total || 0
    })
})

const pieChartOptions = computed(() => ({
    tooltip: { trigger: 'item', formatter: '{b}: {c} ({d}%)' },
    legend: { orient: 'vertical', left: 'left', top: 'middle' },
    series: [
        {
            type: 'pie',
            radius: ['40%', '70%'],
            center: ['60%', '50%'],
            avoidLabelOverlap: false,
            itemStyle: { borderRadius: 6, borderColor: '#fff', borderWidth: 2 },
            label: { show: false },
            emphasis: { label: { show: true, fontSize: 14, fontWeight: 'bold' } },
            data: POLICY_META.map((p, i) => ({
                name: t(`pages.dashboard.policies.${p.key}`),
                value: policyCounts[p.key],
                itemStyle: { color: COLORS[i] },
            })),
        },
    ],
}))

const barChartOptions = computed(() => ({
    tooltip: { trigger: 'axis' },
    grid: { left: '3%', right: '4%', bottom: '3%', containLabel: true },
    xAxis: {
        type: 'category',
        data: [
            t('pages.dashboard.spaces'),
            t('pages.dashboard.applications'),
            t('pages.dashboard.services'),
            t('pages.dashboard.serviceGroups'),
            t('pages.dashboard.serviceAliases'),
        ],
        axisLabel: { rotate: 0 },
    },
    yAxis: { type: 'value', minInterval: 1 },
    series: [
        {
            type: 'bar',
            data: [counts.spaces, counts.applications, counts.services, counts.serviceGroups, counts.serviceAliases],
            barWidth: 40,
            itemStyle: {
                color: {
                    type: 'linear',
                    x: 0,
                    y: 0,
                    x2: 0,
                    y2: 1,
                    colorStops: [
                        { offset: 0, color: '#5470c6' },
                        { offset: 1, color: '#91cc75' },
                    ],
                },
                borderRadius: [4, 4, 0, 0],
            },
        },
    ],
}))

function goTo(path) {
    router.push(path)
}
</script>

<style lang="less" scoped>
.dashboard {
    padding: 0;
}

.stat-card {
    margin-bottom: 16px;
    text-align: center;
    cursor: pointer;
    transition: box-shadow 0.3s;

    &:hover {
        box-shadow: 0 2px 8px rgba(0, 0, 0, 0.15);
    }
}

.quick-link-card {
    text-align: center;
    cursor: pointer;
    transition: all 0.3s;
    margin-bottom: 16px;

    &:hover {
        box-shadow: 0 2px 8px rgba(0, 0, 0, 0.15);
        transform: translateY(-2px);
    }
}

.quick-link-icon {
    font-size: 32px;
    color: #1890ff;
    margin-bottom: 8px;
}

.quick-link-text {
    font-size: 14px;
    color: rgba(0, 0, 0, 0.65);
}
</style>
