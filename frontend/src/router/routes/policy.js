import { ControlOutlined } from '@ant-design/icons-vue'

export default [
    {
        path: 'policy',
        name: 'policy',
        component: 'RouteViewLayout',
        meta: {
            icon: ControlOutlined,
            title: '治理策略',
            isMenu: true,
            keepAlive: true,
            permission: '*',
        },
        children: [
            {
                path: 'loadbalance',
                name: 'loadbalanceList',
                component: 'policy/loadbalance.vue',
                meta: {
                    title: '负载均衡策略',
                    isMenu: true,
                    keepAlive: true,
                    permission: '*',
                },
            },
            {
                path: 'permission',
                name: 'permissionList',
                component: 'policy/permission.vue',
                meta: {
                    title: '服务鉴权策略',
                    isMenu: true,
                    keepAlive: true,
                    permission: '*',
                },
            },
            {
                path: 'fault',
                name: 'faultList',
                component: 'policy/fault.vue',
                meta: {
                    title: '故障注入策略',
                    isMenu: true,
                    keepAlive: true,
                    permission: '*',
                },
            },
            {
                path: 'route',
                name: 'routeList',
                component: 'policy/tagRoute.vue',
                meta: {
                    title: '标签路由策略',
                    isMenu: true,
                    keepAlive: true,
                    permission: '*',
                },
            },
            {
                path: 'limit',
                name: 'limitList',
                component: 'policy/limit.vue',
                meta: {
                    title: '限流策略',
                    isMenu: true,
                    keepAlive: true,
                    permission: '*',
                },
            },
            {
                path: 'invocation',
                name: 'invocationList',
                component: 'policy/invocation.vue',
                meta: {
                    title: '调用容错策略',
                    isMenu: true,
                    keepAlive: true,
                    permission: '*',
                },
            },
            {
                path: 'circuit-break',
                name: 'circuitBreakList',
                component: 'policy/circuitBreak.vue',
                meta: {
                    title: '熔断降级策略',
                    isMenu: true,
                    keepAlive: true,
                    permission: '*',
                },
            },
        ],
    },
]
