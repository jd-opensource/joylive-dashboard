import { AppstoreOutlined } from '@ant-design/icons-vue'

export default [
    {
        path: 'resource',
        name: 'resource',
        component: 'RouteViewLayout',
        meta: {
            icon: AppstoreOutlined,
            title: '基础资源',
            isMenu: true,
            keepAlive: true,
            permission: '*',
        },
        children: [
            {
                path: 'application',
                name: 'applicationList',
                component: 'resource/application.vue',
                meta: {
                    title: '应用管理',
                    isMenu: true,
                    keepAlive: true,
                    permission: '*',
                },
            },
            {
                path: 'service',
                name: 'serviceList',
                component: 'resource/service.vue',
                meta: {
                    title: '服务管理',
                    isMenu: true,
                    keepAlive: true,
                    permission: '*',
                },
            },
            {
                path: 'service/:id',
                name: 'serviceDetail',
                component: 'resource/ServiceDetail.vue',
                meta: {
                    title: '服务详情',
                    isMenu: false,
                    keepAlive: false,
                    permission: '*',
                    active: 'serviceList',
                    breadcrumb: [
                        { name: 'resource', meta: { title: '基础资源' } },
                        { name: 'serviceList', meta: { title: '服务管理' } },
                    ],
                },
            },
        ],
    },
]
