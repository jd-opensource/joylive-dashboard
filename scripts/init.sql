CREATE DATABASE IF NOT EXISTS joylive CHARACTER SET utf8mb4 COLLATE utf8mb4_bin;

CREATE TABLE IF NOT EXISTS `space` (
  `id` varchar(20) CHARACTER SET utf8mb4 COLLATE utf8mb4_bin NOT NULL COMMENT 'ID',
  `code` varchar(255) CHARACTER SET utf8mb4 COLLATE utf8mb4_bin NOT NULL COMMENT '编码',
  `name` varchar(255) CHARACTER SET utf8mb4 COLLATE utf8mb4_bin NOT NULL COMMENT '名称',
  `tenant` varchar(255) CHARACTER SET utf8mb4 COLLATE utf8mb4_bin DEFAULT NULL COMMENT '租户',
  `creator` varchar(255) CHARACTER SET utf8mb4 COLLATE utf8mb4_bin DEFAULT NULL COMMENT '创建者',
  `description` varchar(255) CHARACTER SET utf8mb4 COLLATE utf8mb4_bin DEFAULT NULL COMMENT '备注',
  `metadata` json NULL COMMENT '元数据信息',
  `created_at` timestamp NOT NULL DEFAULT CURRENT_TIMESTAMP COMMENT '创建时间',
  `updated_at` timestamp DEFAULT CURRENT_TIMESTAMP ON UPDATE CURRENT_TIMESTAMP COMMENT '更新时间',
  `deleted` varchar(20) CHARACTER SET utf8mb4 COLLATE utf8mb4_bin NOT NULL DEFAULT '0' COMMENT '逻辑删除标识',
  `deleted_at` DATETIME NULL DEFAULT NULL COMMENT '删除时间',
  PRIMARY KEY (`id`),
  UNIQUE KEY `uniq_code` (`code`, `deleted`) USING BTREE
) ENGINE = InnoDB DEFAULT CHARSET = utf8mb4 COLLATE = utf8mb4_bin COMMENT = '微服务空间';

CREATE TABLE IF NOT EXISTS `application` (
  `id` varchar(20) CHARACTER SET utf8mb4 COLLATE utf8mb4_bin NOT NULL COMMENT 'ID',
  `name` varchar(255) CHARACTER SET utf8mb4 COLLATE utf8mb4_bin NOT NULL COMMENT '应用名称',
  `alias` varchar(255) CHARACTER SET utf8mb4 COLLATE utf8mb4_bin DEFAULT NULL COMMENT '应用中文名称',
  `language` varchar(20) CHARACTER SET utf8mb4 COLLATE utf8mb4_bin DEFAULT NULL COMMENT '服务语言: Java,Python,Golang',
  `enhance` varchar(20) CHARACTER SET utf8mb4 COLLATE utf8mb4_bin NOT NULL COMMENT '增强方式: Agent, Sidecar',
  `source` varchar(255) CHARACTER SET utf8mb4 COLLATE utf8mb4_bin NOT NULL COMMENT '数据来源: Local, JSF, JDAP',
  `tenant` varchar(255) CHARACTER SET utf8mb4 COLLATE utf8mb4_bin DEFAULT NULL COMMENT '租户',
  `creator` varchar(255) CHARACTER SET utf8mb4 COLLATE utf8mb4_bin DEFAULT NULL COMMENT '创建者',
  `description` varchar(255) CHARACTER SET utf8mb4 COLLATE utf8mb4_bin DEFAULT NULL COMMENT '备注',
  `created_at` timestamp NOT NULL DEFAULT CURRENT_TIMESTAMP COMMENT '创建时间',
  `updated_at` timestamp DEFAULT CURRENT_TIMESTAMP ON UPDATE CURRENT_TIMESTAMP COMMENT '更新时间',
  `deleted` varchar(20) CHARACTER SET utf8mb4 COLLATE utf8mb4_bin NOT NULL DEFAULT '0' COMMENT '逻辑删除标识',
  `deleted_at` timestamp NULL DEFAULT NULL COMMENT '删除时间',
  PRIMARY KEY (`id`),
  UNIQUE KEY `uniq_name` (`name`, `deleted`) USING BTREE
) ENGINE = InnoDB DEFAULT CHARSET = utf8mb4 COLLATE = utf8mb4_bin COMMENT = '应用';

CREATE TABLE IF NOT EXISTS `application_service` (
  `id` varchar(20) CHARACTER SET utf8mb4 COLLATE utf8mb4_bin NOT NULL COMMENT 'ID',
  `application_id` varchar(20) CHARACTER SET utf8mb4 COLLATE utf8mb4_bin NOT NULL COMMENT '应用ID',
  `service_id` varchar(20) CHARACTER SET utf8mb4 COLLATE utf8mb4_bin NOT NULL COMMENT '服务ID',
  `role` varchar(20) CHARACTER SET utf8mb4 COLLATE utf8mb4_bin NOT NULL COMMENT '关联角色: provider, consumer',
  `status` varchar(20) CHARACTER SET utf8mb4 COLLATE utf8mb4_bin DEFAULT NULL COMMENT '消费状态: approved, rejected, pending',
  `description` varchar(255) CHARACTER SET utf8mb4 COLLATE utf8mb4_bin DEFAULT NULL COMMENT '备注',
  `created_at` timestamp NOT NULL DEFAULT CURRENT_TIMESTAMP COMMENT '创建时间',
  `updated_at` timestamp DEFAULT CURRENT_TIMESTAMP ON UPDATE CURRENT_TIMESTAMP COMMENT '更新时间',
  `deleted` varchar(20) CHARACTER SET utf8mb4 COLLATE utf8mb4_bin NOT NULL DEFAULT '0' COMMENT '逻辑删除标识',
  `deleted_at` timestamp NULL DEFAULT NULL COMMENT '删除时间',
  PRIMARY KEY (`id`),
  UNIQUE KEY `uniq_app_service` (
    `application_id`,
    `service_id`,
    `role`,
    `deleted`
  ) USING BTREE
) ENGINE = InnoDB DEFAULT CHARSET = utf8mb4 COLLATE = utf8mb4_bin COMMENT = '应用服务关联';

CREATE TABLE IF NOT EXISTS `service` (
  `id` varchar(20) CHARACTER SET utf8mb4 COLLATE utf8mb4_bin NOT NULL COMMENT 'ID',
  `name` varchar(255) CHARACTER SET utf8mb4 COLLATE utf8mb4_general_ci NOT NULL COMMENT '名称',
  `space_code` varchar(255) CHARACTER SET utf8mb4 COLLATE utf8mb4_bin NOT NULL COMMENT '微服务空间编码',
  `registration_type` varchar(20) CHARACTER SET utf8mb4 COLLATE utf8mb4_bin NOT NULL COMMENT '服务注册粒度: HTTP应用级(HTTP), RPC应用级(RPC_APP), RPC接口级(RPC_INTERFACE)',
  `source` varchar(255) CHARACTER SET utf8mb4 COLLATE utf8mb4_bin DEFAULT NULL COMMENT '数据来源: Local, JSF, JDAP',
  `tenant` varchar(255) CHARACTER SET utf8mb4 COLLATE utf8mb4_bin DEFAULT NULL COMMENT '租户',
  `creator` varchar(255) CHARACTER SET utf8mb4 COLLATE utf8mb4_bin DEFAULT NULL COMMENT '创建者',
  `extra` json NULL COMMENT '其他信息',
  `version` bigint NOT NULL DEFAULT '1' COMMENT '版本号',
  `description` varchar(255) CHARACTER SET utf8mb4 COLLATE utf8mb4_bin DEFAULT NULL COMMENT '备注',
  `created_at` timestamp NOT NULL DEFAULT CURRENT_TIMESTAMP COMMENT '创建时间',
  `updated_at` timestamp DEFAULT CURRENT_TIMESTAMP ON UPDATE CURRENT_TIMESTAMP COMMENT '更新时间',
  `deleted` varchar(20) CHARACTER SET utf8mb4 COLLATE utf8mb4_bin NOT NULL DEFAULT '0' COMMENT '逻辑删除标识',
  `deleted_at` timestamp NULL DEFAULT NULL COMMENT '删除时间',
  PRIMARY KEY (`id`),
  UNIQUE KEY `uniq_name` (`name`, `space_code`, `deleted`) USING BTREE
) ENGINE = InnoDB DEFAULT CHARSET = utf8mb4 COLLATE = utf8mb4_bin COMMENT = '微服务';

CREATE TABLE IF NOT EXISTS `service_group` (
  `id` varchar(20) CHARACTER SET utf8mb4 COLLATE utf8mb4_bin NOT NULL COMMENT 'ID',
  `name` varchar(255) CHARACTER SET utf8mb4 COLLATE utf8mb4_bin NOT NULL COMMENT '分组名称',
  `code` varchar(255) CHARACTER SET utf8mb4 COLLATE utf8mb4_bin NOT NULL COMMENT '分组编码',
  `isolation_code` varchar(255) CHARACTER SET utf8mb4 COLLATE utf8mb4_bin DEFAULT NULL COMMENT '隔离编码',
  `service_id` varchar(20) CHARACTER SET utf8mb4 COLLATE utf8mb4_bin NOT NULL COMMENT '所属服务ID',
  `labels` varchar(1024) CHARACTER SET utf8mb4 COLLATE utf8mb4_bin DEFAULT NULL COMMENT '标签',
  `description` varchar(255) CHARACTER SET utf8mb4 COLLATE utf8mb4_bin DEFAULT NULL COMMENT '备注',
  `created_at` timestamp NOT NULL DEFAULT CURRENT_TIMESTAMP COMMENT '创建时间',
  `updated_at` timestamp DEFAULT CURRENT_TIMESTAMP ON UPDATE CURRENT_TIMESTAMP COMMENT '更新时间',
  `deleted` varchar(20) CHARACTER SET utf8mb4 COLLATE utf8mb4_bin NOT NULL DEFAULT '0' COMMENT '逻辑删除标识',
  `deleted_at` timestamp NULL DEFAULT NULL COMMENT '删除时间',
  PRIMARY KEY (`id`),
  UNIQUE KEY `uniq_code` (`code`, `service_id`, `deleted`) USING BTREE
) ENGINE = InnoDB DEFAULT CHARSET = utf8mb4 COLLATE = utf8mb4_bin COMMENT = '服务分组';

CREATE TABLE IF NOT EXISTS `service_alias` (
  `id` varchar(20) CHARACTER SET utf8mb4 COLLATE utf8mb4_bin NOT NULL COMMENT 'ID',
  `space_code` varchar(255) CHARACTER SET utf8mb4 COLLATE utf8mb4_bin NOT NULL COMMENT '微服务空间编码',
  `alias` varchar(255) CHARACTER SET utf8mb4 COLLATE utf8mb4_bin NOT NULL COMMENT '服务别名',
  `service_id` varchar(20) CHARACTER SET utf8mb4 COLLATE utf8mb4_bin NOT NULL COMMENT '所属服务ID',
  `description` varchar(255) CHARACTER SET utf8mb4 COLLATE utf8mb4_bin DEFAULT NULL COMMENT '备注',
  `created_at` timestamp NOT NULL DEFAULT CURRENT_TIMESTAMP COMMENT '创建时间',
  `updated_at` timestamp DEFAULT CURRENT_TIMESTAMP ON UPDATE CURRENT_TIMESTAMP COMMENT '更新时间',
  `deleted` varchar(20) CHARACTER SET utf8mb4 COLLATE utf8mb4_bin NOT NULL DEFAULT '0' COMMENT '逻辑删除标识',
  `deleted_at` timestamp NULL DEFAULT NULL COMMENT '删除时间',
  PRIMARY KEY (`id`),
  UNIQUE KEY `uniq_code` (`space_code`, `alias`, `deleted`) USING BTREE
) ENGINE = InnoDB DEFAULT CHARSET = utf8mb4 COLLATE = utf8mb4_bin COMMENT = '服务别名';

CREATE TABLE IF NOT EXISTS `policy_loadbalance` (
  `id` varchar(20) CHARACTER SET utf8mb4 COLLATE utf8mb4_bin NOT NULL COMMENT 'ID',
  `name` varchar(100) CHARACTER SET utf8mb4 COLLATE utf8mb4_bin NOT NULL COMMENT '名称',
  `space_code` varchar(255) CHARACTER SET utf8mb4 COLLATE utf8mb4_bin DEFAULT NULL COMMENT '微服务空间编码',
  `source_application_id` varchar(20) CHARACTER SET utf8mb4 COLLATE utf8mb4_bin DEFAULT NULL COMMENT '主调应用ID',
  `target_service_id` varchar(20) CHARACTER SET utf8mb4 COLLATE utf8mb4_bin NOT NULL COMMENT '被调服务ID',
  `group` varchar(255) CHARACTER SET utf8mb4 COLLATE utf8mb4_bin NOT NULL DEFAULT 'default' COMMENT '分组',
  `path` varchar(255) CHARACTER SET utf8mb4 COLLATE utf8mb4_bin DEFAULT NULL COMMENT '路径或接口',
  `method` varchar(255) CHARACTER SET utf8mb4 COLLATE utf8mb4_bin DEFAULT NULL COMMENT '方法',
  `policy_type` varchar(20) CHARACTER SET utf8mb4 COLLATE utf8mb4_bin NOT NULL COMMENT '负载均衡类型',
  `sticky_type` varchar(20) CHARACTER SET utf8mb4 COLLATE utf8mb4_bin DEFAULT NULL COMMENT '粘性类型',
  `version` bigint NOT NULL DEFAULT '1' COMMENT '版本号',
  `enabled` int NOT NULL DEFAULT '0' COMMENT '启用',
  `description` varchar(255) CHARACTER SET utf8mb4 COLLATE utf8mb4_bin DEFAULT NULL COMMENT '备注',
  `creator` varchar(255) CHARACTER SET utf8mb4 COLLATE utf8mb4_bin DEFAULT NULL COMMENT '创建者',
  `modifier` varchar(255) CHARACTER SET utf8mb4 COLLATE utf8mb4_bin DEFAULT NULL COMMENT '修改者',
  `created_at` timestamp NOT NULL DEFAULT CURRENT_TIMESTAMP COMMENT '创建时间',
  `updated_at` timestamp DEFAULT CURRENT_TIMESTAMP ON UPDATE CURRENT_TIMESTAMP COMMENT '更新时间',
  `deleted` varchar(20) CHARACTER SET utf8mb4 COLLATE utf8mb4_bin NOT NULL DEFAULT '0' COMMENT '逻辑删除标识',
  `deleted_at` timestamp NULL DEFAULT NULL COMMENT '删除时间',
  PRIMARY KEY (`id`),
  UNIQUE KEY `uniq_name` (
    `name`,
    `space_code`,
    `source_application_id`,
    `target_service_id`,
    `deleted`
  ) USING BTREE,
  UNIQUE KEY `uniq_service` (
    `space_code`,
    `source_application_id`,
    `target_service_id`,
    `deleted`
  ) USING BTREE
) ENGINE = InnoDB COLLATE = utf8mb4_bin COMMENT = '负载均衡策略';

CREATE TABLE IF NOT EXISTS `policy_route` (
  `id` varchar(20) CHARACTER SET utf8mb4 COLLATE utf8mb4_bin NOT NULL COMMENT 'ID',
  `name` varchar(100) CHARACTER SET utf8mb4 COLLATE utf8mb4_bin NOT NULL COMMENT '名称',
  `space_code` varchar(255) CHARACTER SET utf8mb4 COLLATE utf8mb4_bin NOT NULL COMMENT '微服务空间编码',
  `source_application_id` varchar(20) CHARACTER SET utf8mb4 COLLATE utf8mb4_bin DEFAULT NULL COMMENT '主调应用ID',
  `target_service_id` varchar(20) CHARACTER SET utf8mb4 COLLATE utf8mb4_bin NOT NULL COMMENT '被调服务ID',
  `group` varchar(255) CHARACTER SET utf8mb4 COLLATE utf8mb4_bin NOT NULL DEFAULT 'default' COMMENT '分组',
  `path` varchar(255) CHARACTER SET utf8mb4 COLLATE utf8mb4_bin DEFAULT NULL COMMENT '路径或接口',
  `method` varchar(255) CHARACTER SET utf8mb4 COLLATE utf8mb4_bin DEFAULT NULL COMMENT '方法',
  `order` int NOT NULL DEFAULT '0' COMMENT '排序号',
  `version` bigint NOT NULL DEFAULT '1' COMMENT '版本号',
  `enabled` int NOT NULL DEFAULT '0' COMMENT '启用',
  `description` varchar(255) CHARACTER SET utf8mb4 COLLATE utf8mb4_bin DEFAULT NULL COMMENT '备注',
  `creator` varchar(255) CHARACTER SET utf8mb4 COLLATE utf8mb4_bin DEFAULT NULL COMMENT '创建者',
  `modifier` varchar(255) CHARACTER SET utf8mb4 COLLATE utf8mb4_bin DEFAULT NULL COMMENT '修改者',
  `created_at` timestamp NOT NULL DEFAULT CURRENT_TIMESTAMP COMMENT '创建时间',
  `updated_at` timestamp DEFAULT CURRENT_TIMESTAMP ON UPDATE CURRENT_TIMESTAMP COMMENT '更新时间',
  `deleted` varchar(20) CHARACTER SET utf8mb4 COLLATE utf8mb4_bin NOT NULL DEFAULT '0' COMMENT '逻辑删除标识',
  `deleted_at` timestamp NULL DEFAULT NULL COMMENT '删除时间',
  PRIMARY KEY (`id`),
  UNIQUE KEY `uniq_name` (
    `name`,
    `space_code`,
    `source_application_id`,
    `deleted`
  ) USING BTREE
) ENGINE = InnoDB DEFAULT CHARSET = utf8mb4 COLLATE = utf8mb4_bin COMMENT = '标签路由策略';

CREATE TABLE IF NOT EXISTS `policy_route_detail` (
  `id` varchar(20) CHARACTER SET utf8mb4 COLLATE utf8mb4_bin NOT NULL COMMENT 'ID',
  `route_id` varchar(20) CHARACTER SET utf8mb4 COLLATE utf8mb4_bin NOT NULL COMMENT '路由ID',
  `relation_type` varchar(20) CHARACTER SET utf8mb4 COLLATE utf8mb4_bin NOT NULL COMMENT '关系类型',
  `conditions` json DEFAULT NULL COMMENT '匹配条件',
  `destinations` json DEFAULT NULL COMMENT '目的规则',
  `order` int NOT NULL DEFAULT '0' COMMENT '排序值',
  `enabled` int NOT NULL DEFAULT '0' COMMENT '启用',
  `description` varchar(255) CHARACTER SET utf8mb4 COLLATE utf8mb4_bin DEFAULT NULL COMMENT '备注',
  `created_at` timestamp NOT NULL DEFAULT CURRENT_TIMESTAMP COMMENT '创建时间',
  `updated_at` timestamp DEFAULT CURRENT_TIMESTAMP ON UPDATE CURRENT_TIMESTAMP COMMENT '更新时间',
  `deleted` varchar(20) CHARACTER SET utf8mb4 COLLATE utf8mb4_bin NOT NULL DEFAULT '0' COMMENT '逻辑删除标识',
  `deleted_at` timestamp NULL DEFAULT NULL COMMENT '删除时间',
  PRIMARY KEY (`id`),
  KEY `idx_routeid` (`route_id`, `deleted`) USING BTREE
) ENGINE = InnoDB DEFAULT CHARSET = utf8mb4 COLLATE = utf8mb4_bin COMMENT = '标签路由详情';

CREATE TABLE IF NOT EXISTS `policy_limit` (
  `id` varchar(20) CHARACTER SET utf8mb4 COLLATE utf8mb4_bin NOT NULL COMMENT 'ID',
  `name` varchar(100) CHARACTER SET utf8mb4 COLLATE utf8mb4_bin NOT NULL COMMENT '名称',
  `space_code` varchar(255) CHARACTER SET utf8mb4 COLLATE utf8mb4_bin NOT NULL COMMENT '微服务空间编码',
  `source_application_id` varchar(20) CHARACTER SET utf8mb4 COLLATE utf8mb4_bin DEFAULT NULL COMMENT '主调应用ID',
  `target_service_id` varchar(20) CHARACTER SET utf8mb4 COLLATE utf8mb4_bin NOT NULL COMMENT '被调服务ID',
  `group` varchar(255) CHARACTER SET utf8mb4 COLLATE utf8mb4_bin NOT NULL DEFAULT 'default' COMMENT '分组',
  `path` varchar(255) CHARACTER SET utf8mb4 COLLATE utf8mb4_bin DEFAULT NULL COMMENT '路径或接口',
  `method` varchar(255) CHARACTER SET utf8mb4 COLLATE utf8mb4_bin DEFAULT NULL COMMENT '方法',
  `relation_type` varchar(20) CHARACTER SET utf8mb4 COLLATE utf8mb4_bin NOT NULL COMMENT '关系类型',
  `conditions` json DEFAULT NULL COMMENT '匹配条件',
  `type` varchar(20) CHARACTER SET utf8mb4 COLLATE utf8mb4_bin NOT NULL COMMENT '限流策略类型(Rate, Concurrency, Load)',
  `realize_type` varchar(20) CHARACTER SET utf8mb4 COLLATE utf8mb4_bin NOT NULL COMMENT '实现类型',
  `sliding_windows` json DEFAULT NULL COMMENT '统计规则',
  `max_concurrency` int DEFAULT '0' COMMENT '最大并发数',
  `max_wait_ms` bigint NOT NULL DEFAULT '0' COMMENT '最大等待时间(ms)',
  `cpu_usage` int DEFAULT '0' COMMENT 'CPU最大使用率',
  `load_usage` int DEFAULT '0' COMMENT '系统负载最大值',
  `action_parameters` json DEFAULT NULL COMMENT '参数',
  `version` bigint NOT NULL DEFAULT '1' COMMENT '版本号',
  `enabled` int NOT NULL DEFAULT '0' COMMENT '启用',
  `description` varchar(255) CHARACTER SET utf8mb4 COLLATE utf8mb4_bin DEFAULT NULL COMMENT '备注',
  `creator` varchar(255) CHARACTER SET utf8mb4 COLLATE utf8mb4_bin DEFAULT NULL COMMENT '创建者',
  `modifier` varchar(255) CHARACTER SET utf8mb4 COLLATE utf8mb4_bin DEFAULT NULL COMMENT '修改者',
  `created_at` timestamp NOT NULL DEFAULT CURRENT_TIMESTAMP COMMENT '创建时间',
  `updated_at` timestamp DEFAULT CURRENT_TIMESTAMP ON UPDATE CURRENT_TIMESTAMP COMMENT '更新时间',
  `deleted` varchar(20) CHARACTER SET utf8mb4 COLLATE utf8mb4_bin NOT NULL DEFAULT '0' COMMENT '逻辑删除标识',
  `deleted_at` timestamp NULL DEFAULT NULL COMMENT '删除时间',
  PRIMARY KEY (`id`),
  UNIQUE KEY `uniq_name` (`name`, `target_service_id`, `deleted`) USING BTREE
) ENGINE = InnoDB DEFAULT CHARSET = utf8mb4 COLLATE = utf8mb4_bin COMMENT = '服务限流策略';

CREATE TABLE IF NOT EXISTS `policy_circuit_break` (
  `id` varchar(20) CHARACTER SET utf8mb4 COLLATE utf8mb4_bin NOT NULL COMMENT 'ID',
  `name` varchar(100) CHARACTER SET utf8mb4 COLLATE utf8mb4_bin NOT NULL COMMENT '名称',
  `space_code` varchar(255) CHARACTER SET utf8mb4 COLLATE utf8mb4_bin NOT NULL COMMENT '微服务空间编码',
  `source_application_id` varchar(20) CHARACTER SET utf8mb4 COLLATE utf8mb4_bin DEFAULT NULL COMMENT '主调应用ID',
  `target_service_id` varchar(20) CHARACTER SET utf8mb4 COLLATE utf8mb4_bin NOT NULL COMMENT '被调服务ID',
  `group` varchar(255) CHARACTER SET utf8mb4 COLLATE utf8mb4_bin NOT NULL DEFAULT 'default' COMMENT '分组',
  `path` varchar(255) CHARACTER SET utf8mb4 COLLATE utf8mb4_bin DEFAULT NULL COMMENT '路径或接口',
  `method` varchar(255) CHARACTER SET utf8mb4 COLLATE utf8mb4_bin DEFAULT NULL COMMENT '方法',
  `level` varchar(20) CHARACTER SET utf8mb4 COLLATE utf8mb4_bin NOT NULL COMMENT '策略级别',
  `sliding_window_type` varchar(20) CHARACTER SET utf8mb4 COLLATE utf8mb4_bin NOT NULL COMMENT '滑动窗口类型',
  `sliding_window_size` int NOT NULL DEFAULT '1' COMMENT '滑动窗口大小',
  `min_calls_threshold` int NOT NULL DEFAULT '1' COMMENT '最小调用次数',
  `code_policy` json DEFAULT NULL COMMENT '响应码策略',
  `error_codes` json DEFAULT NULL COMMENT '错误编码列表',
  `message_policy` json DEFAULT NULL COMMENT '错误消息策略',
  `error_messages` json DEFAULT NULL COMMENT '错误消息列表',
  `exceptions` json DEFAULT NULL COMMENT '异常列表',
  `failure_rate_threshold` int NOT NULL DEFAULT '1' COMMENT '异常比例',
  `slow_call_rate_threshold` int NOT NULL DEFAULT '1' COMMENT '慢调用比例',
  `slow_call_duration_threshold` int NOT NULL DEFAULT '10000' COMMENT '慢调用阈值',
  `wait_duration_in_open_state` int NOT NULL DEFAULT '10000' COMMENT '熔断时间',
  `outlier_max_percent` int DEFAULT NULL COMMENT '最大离群比例',
  `allowed_calls_in_half_open_state` int NOT NULL DEFAULT '5' COMMENT '半开状态可调用数',
  `force_open` int NOT NULL DEFAULT '0' COMMENT '强制开启熔断',
  `recovery_enabled` int NOT NULL DEFAULT '0' COMMENT '渐进恢复开关',
  `recovery_duration` int NOT NULL DEFAULT '0' COMMENT '渐进恢复时长（毫秒）',
  `recovery_phase` int NOT NULL DEFAULT '0' COMMENT '渐进恢复阶段',
  `degrade_config` json DEFAULT NULL COMMENT '降级配置',
  `realize_type` varchar(20) CHARACTER SET utf8mb4 COLLATE utf8mb4_bin NOT NULL DEFAULT 'Resilience4j' COMMENT '实现类型',
  `version` bigint NOT NULL DEFAULT '1' COMMENT '版本号',
  `enabled` int NOT NULL DEFAULT '0' COMMENT '启用',
  `description` varchar(255) CHARACTER SET utf8mb4 COLLATE utf8mb4_bin DEFAULT NULL COMMENT '备注',
  `creator` varchar(255) CHARACTER SET utf8mb4 COLLATE utf8mb4_bin DEFAULT NULL COMMENT '创建者',
  `modifier` varchar(255) CHARACTER SET utf8mb4 COLLATE utf8mb4_bin DEFAULT NULL COMMENT '修改者',
  `created_at` timestamp NOT NULL DEFAULT CURRENT_TIMESTAMP COMMENT '创建时间',
  `updated_at` timestamp DEFAULT CURRENT_TIMESTAMP ON UPDATE CURRENT_TIMESTAMP COMMENT '更新时间',
  `deleted` varchar(20) CHARACTER SET utf8mb4 COLLATE utf8mb4_bin NOT NULL DEFAULT '0' COMMENT '逻辑删除标识',
  `deleted_at` timestamp NULL DEFAULT NULL COMMENT '删除时间',
  PRIMARY KEY (`id`),
  UNIQUE KEY `uniq_name` (
    `name`,
    `space_code`,
    `source_application_id`,
    `target_service_id`,
    `deleted`
  ) USING BTREE
) ENGINE = InnoDB DEFAULT CHARSET = utf8mb4 COLLATE = utf8mb4_bin COMMENT = '熔断降级策略';

CREATE TABLE IF NOT EXISTS `policy_permission` (
  `id` varchar(20) CHARACTER SET utf8mb4 COLLATE utf8mb4_bin NOT NULL COMMENT 'ID',
  `name` varchar(255) CHARACTER SET utf8mb4 COLLATE utf8mb4_bin NOT NULL COMMENT '授权策略名称',
  `space_code` varchar(255) CHARACTER SET utf8mb4 COLLATE utf8mb4_bin NOT NULL COMMENT '微服务空间编码',
  `source_application_id` varchar(20) CHARACTER SET utf8mb4 COLLATE utf8mb4_bin DEFAULT NULL COMMENT '主调应用ID',
  `target_service_id` varchar(20) CHARACTER SET utf8mb4 COLLATE utf8mb4_bin NOT NULL COMMENT '被调服务ID',
  `group` varchar(255) CHARACTER SET utf8mb4 COLLATE utf8mb4_bin NOT NULL DEFAULT 'default' COMMENT '分组',
  `path` varchar(255) CHARACTER SET utf8mb4 COLLATE utf8mb4_bin DEFAULT NULL COMMENT '路径或接口',
  `method` varchar(255) CHARACTER SET utf8mb4 COLLATE utf8mb4_bin DEFAULT NULL COMMENT '方法',
  `relation_type` varchar(20) CHARACTER SET utf8mb4 COLLATE utf8mb4_bin NOT NULL COMMENT '关系类型',
  `conditions` json DEFAULT NULL COMMENT '匹配条件',
  `type` varchar(20) CHARACTER SET utf8mb4 COLLATE utf8mb4_bin NOT NULL COMMENT '授权类型',
  `version` bigint NOT NULL DEFAULT '1' COMMENT '版本号',
  `enabled` int NOT NULL DEFAULT '0' COMMENT '启用',
  `description` varchar(255) CHARACTER SET utf8mb4 COLLATE utf8mb4_bin DEFAULT NULL COMMENT '备注',
  `creator` varchar(255) CHARACTER SET utf8mb4 COLLATE utf8mb4_bin DEFAULT NULL COMMENT '创建者',
  `modifier` varchar(255) CHARACTER SET utf8mb4 COLLATE utf8mb4_bin DEFAULT NULL COMMENT '修改者',
  `created_at` timestamp NOT NULL DEFAULT CURRENT_TIMESTAMP COMMENT '创建时间',
  `updated_at` timestamp DEFAULT CURRENT_TIMESTAMP ON UPDATE CURRENT_TIMESTAMP COMMENT '更新时间',
  `deleted` varchar(20) CHARACTER SET utf8mb4 COLLATE utf8mb4_bin NOT NULL DEFAULT '0' COMMENT '逻辑删除标识',
  `deleted_at` timestamp NULL DEFAULT NULL COMMENT '删除时间',
  PRIMARY KEY (`id`),
  UNIQUE KEY `uniq_name` (`name`, `target_service_id`, `deleted`) USING BTREE
) ENGINE = InnoDB DEFAULT CHARSET = utf8mb4 COLLATE = utf8mb4_bin COMMENT = '服务鉴权策略';

CREATE TABLE IF NOT EXISTS `policy_auth` (
  `id` varchar(20) CHARACTER SET utf8mb4 COLLATE utf8mb4_bin NOT NULL COMMENT 'ID',
  `space_code` varchar(255) CHARACTER SET utf8mb4 COLLATE utf8mb4_bin NOT NULL COMMENT '微服务空间编码',
  `source_application_id` varchar(20) CHARACTER SET utf8mb4 COLLATE utf8mb4_bin DEFAULT NULL COMMENT '主调应用ID',
  `target_service_id` varchar(20) CHARACTER SET utf8mb4 COLLATE utf8mb4_bin NOT NULL COMMENT '被调服务ID',
  `type` varchar(20) CHARACTER SET utf8mb4 COLLATE utf8mb4_bin NOT NULL COMMENT '认证类型',
  `params` json DEFAULT NULL COMMENT '参数',
  `version` bigint NOT NULL DEFAULT '1' COMMENT '版本号',
  `enabled` int NOT NULL DEFAULT '0' COMMENT '启用',
  `description` varchar(255) CHARACTER SET utf8mb4 COLLATE utf8mb4_bin DEFAULT NULL COMMENT '备注',
  `creator` varchar(255) CHARACTER SET utf8mb4 COLLATE utf8mb4_bin DEFAULT NULL COMMENT '创建者',
  `modifier` varchar(255) CHARACTER SET utf8mb4 COLLATE utf8mb4_bin DEFAULT NULL COMMENT '修改者',
  `created_at` timestamp NOT NULL DEFAULT CURRENT_TIMESTAMP COMMENT '创建时间',
  `updated_at` timestamp DEFAULT CURRENT_TIMESTAMP ON UPDATE CURRENT_TIMESTAMP COMMENT '更新时间',
  `deleted` varchar(20) CHARACTER SET utf8mb4 COLLATE utf8mb4_bin NOT NULL DEFAULT '0' COMMENT '逻辑删除标识',
  `deleted_at` timestamp NULL DEFAULT NULL COMMENT '删除时间',
  PRIMARY KEY (`id`)
) ENGINE = InnoDB DEFAULT CHARSET = utf8mb4 COLLATE = utf8mb4_bin COMMENT = '服务认证策略';

CREATE TABLE IF NOT EXISTS `policy_fault` (
  `id` varchar(20) CHARACTER SET utf8mb4 COLLATE utf8mb4_bin NOT NULL COMMENT 'ID',
  `name` varchar(100) CHARACTER SET utf8mb4 COLLATE utf8mb4_bin NOT NULL COMMENT '名称',
  `space_code` varchar(255) CHARACTER SET utf8mb4 COLLATE utf8mb4_bin NOT NULL COMMENT '微服务空间编码',
  `source_application_id` varchar(20) CHARACTER SET utf8mb4 COLLATE utf8mb4_bin DEFAULT NULL COMMENT '主调应用ID',
  `target_service_id` varchar(20) CHARACTER SET utf8mb4 COLLATE utf8mb4_bin NOT NULL COMMENT '被调服务ID',
  `group` varchar(255) CHARACTER SET utf8mb4 COLLATE utf8mb4_bin NOT NULL DEFAULT 'default' COMMENT '分组',
  `path` varchar(255) CHARACTER SET utf8mb4 COLLATE utf8mb4_bin DEFAULT NULL COMMENT '路径或接口',
  `method` varchar(255) CHARACTER SET utf8mb4 COLLATE utf8mb4_bin DEFAULT NULL COMMENT '方法',
  `relation_type` varchar(20) CHARACTER SET utf8mb4 COLLATE utf8mb4_bin NOT NULL COMMENT '关系类型',
  `conditions` json DEFAULT NULL COMMENT '匹配条件',
  `type` varchar(20) CHARACTER SET utf8mb4 COLLATE utf8mb4_bin NOT NULL COMMENT '故障注入类型（error｜delay）',
  `scope` varchar(20) CHARACTER SET utf8mb4 COLLATE utf8mb4_bin NOT NULL COMMENT '注入方向（inbound｜outbound）',
  `error_code` int DEFAULT '0' COMMENT '错误码',
  `error_msg` varchar(1024) CHARACTER SET utf8mb4 COLLATE utf8mb4_bin COMMENT '错误内容',
  `delay_time_ms` int DEFAULT '0' COMMENT '延迟时间(ms)',
  `percent` int NOT NULL DEFAULT '0' COMMENT '故障注入比例',
  `version` bigint NOT NULL DEFAULT '1' COMMENT '版本号',
  `enabled` int NOT NULL DEFAULT '0' COMMENT '启用',
  `description` varchar(255) CHARACTER SET utf8mb4 COLLATE utf8mb4_bin DEFAULT NULL COMMENT '备注',
  `creator` varchar(255) CHARACTER SET utf8mb4 COLLATE utf8mb4_bin DEFAULT NULL COMMENT '创建者',
  `modifier` varchar(255) CHARACTER SET utf8mb4 COLLATE utf8mb4_bin DEFAULT NULL COMMENT '修改者',
  `created_at` timestamp NOT NULL DEFAULT CURRENT_TIMESTAMP COMMENT '创建时间',
  `updated_at` timestamp DEFAULT CURRENT_TIMESTAMP ON UPDATE CURRENT_TIMESTAMP COMMENT '更新时间',
  `deleted` varchar(20) CHARACTER SET utf8mb4 COLLATE utf8mb4_bin NOT NULL DEFAULT '0' COMMENT '逻辑删除标识',
  `deleted_at` timestamp NULL DEFAULT NULL COMMENT '删除时间',
  PRIMARY KEY (`id`),
  UNIQUE KEY `uniq_name` (
    `name`,
    `space_code`,
    `source_application_id`,
    `target_service_id`,
    `deleted`
  ) USING BTREE
) ENGINE = InnoDB DEFAULT CHARSET = utf8mb4 COLLATE = utf8mb4_bin COMMENT = '故障注入策略';

CREATE TABLE IF NOT EXISTS `policy_invocation` (
  `id` varchar(20) CHARACTER SET utf8mb4 COLLATE utf8mb4_bin NOT NULL COMMENT 'ID',
  `name` varchar(100) CHARACTER SET utf8mb4 COLLATE utf8mb4_bin NOT NULL COMMENT '名称',
  `space_code` varchar(255) CHARACTER SET utf8mb4 COLLATE utf8mb4_bin NOT NULL COMMENT '微服务空间编码',
  `source_application_id` varchar(20) CHARACTER SET utf8mb4 COLLATE utf8mb4_bin DEFAULT NULL COMMENT '主调应用ID',
  `target_service_id` varchar(20) CHARACTER SET utf8mb4 COLLATE utf8mb4_bin NOT NULL COMMENT '被调服务ID',
  `group` varchar(255) CHARACTER SET utf8mb4 COLLATE utf8mb4_bin NOT NULL DEFAULT 'default' COMMENT '分组',
  `path` varchar(255) CHARACTER SET utf8mb4 COLLATE utf8mb4_bin DEFAULT NULL COMMENT '路径或接口',
  `method` varchar(255) CHARACTER SET utf8mb4 COLLATE utf8mb4_bin DEFAULT NULL COMMENT '方法',
  `type` varchar(20) CHARACTER SET utf8mb4 COLLATE utf8mb4_bin NOT NULL COMMENT '调用策略类型（failfast｜failover｜failsafe）',
  `retry_policy` json DEFAULT NULL COMMENT '重试策略',
  `version` bigint NOT NULL DEFAULT '1' COMMENT '版本号',
  `enabled` int NOT NULL DEFAULT '0' COMMENT '启用',
  `description` varchar(255) CHARACTER SET utf8mb4 COLLATE utf8mb4_bin DEFAULT NULL COMMENT '备注',
  `creator` varchar(255) CHARACTER SET utf8mb4 COLLATE utf8mb4_bin DEFAULT NULL COMMENT '创建者',
  `modifier` varchar(255) CHARACTER SET utf8mb4 COLLATE utf8mb4_bin DEFAULT NULL COMMENT '修改者',
  `created_at` timestamp NOT NULL DEFAULT CURRENT_TIMESTAMP COMMENT '创建时间',
  `updated_at` timestamp DEFAULT CURRENT_TIMESTAMP ON UPDATE CURRENT_TIMESTAMP COMMENT '更新时间',
  `deleted` varchar(20) CHARACTER SET utf8mb4 COLLATE utf8mb4_bin NOT NULL DEFAULT '0' COMMENT '逻辑删除标识',
  `deleted_at` timestamp NULL DEFAULT NULL COMMENT '删除时间',
  PRIMARY KEY (`id`),
  UNIQUE KEY `uniq_name` (
    `name`,
    `space_code`,
    `source_application_id`,
    `target_service_id`,
    `deleted`
  ) USING BTREE
) ENGINE = InnoDB DEFAULT CHARSET = utf8mb4 COLLATE = utf8mb4_bin COMMENT = '调用策略';

CREATE TABLE IF NOT EXISTS `config` (
  `id` varchar(20) CHARACTER SET utf8mb4 COLLATE utf8mb4_bin NOT NULL COMMENT 'ID',
  `name` varchar(255) CHARACTER SET utf8mb4 COLLATE utf8mb4_bin NOT NULL COMMENT '名称',
  `group` varchar(64) CHARACTER SET utf8mb4 COLLATE utf8mb4_bin NOT NULL COMMENT '分组',
  `space_code` varchar(255) CHARACTER SET utf8mb4 COLLATE utf8mb4_bin NOT NULL COMMENT '微服务空间编码',
  `app_id` varchar(20) CHARACTER SET utf8mb4 COLLATE utf8mb4_bin DEFAULT NULL COMMENT '应用ID',
  `app_name` varchar(255) CHARACTER SET utf8mb4 COLLATE utf8mb4_bin DEFAULT NULL COMMENT '应用名称',
  `data_type` varchar(20) CHARACTER SET utf8mb4 COLLATE utf8mb4_bin DEFAULT NULL COMMENT '数据类型:yaml、json、text、xml、html、properties、toml',
  `content` text CHARACTER SET utf8mb4 COLLATE utf8mb4_bin DEFAULT NULL COMMENT '配置信息',
  `creator` varchar(255) CHARACTER SET utf8mb4 COLLATE utf8mb4_bin DEFAULT NULL COMMENT '创建者',
  `tenant` varchar(255) CHARACTER SET utf8mb4 COLLATE utf8mb4_bin DEFAULT NULL COMMENT '租户信息',
  `created_at` timestamp NOT NULL DEFAULT CURRENT_TIMESTAMP COMMENT '创建时间',
  `updated_at` timestamp DEFAULT CURRENT_TIMESTAMP ON UPDATE CURRENT_TIMESTAMP COMMENT '更新时间',
  `deleted` varchar(20) CHARACTER SET utf8mb4 COLLATE utf8mb4_bin NOT NULL DEFAULT '0' COMMENT '逻辑删除标识',
  `deleted_at` timestamp NULL DEFAULT NULL COMMENT '删除时间',
  PRIMARY KEY (`id`),
  UNIQUE KEY `uniq_space_code_group_name` (`space_code`, `group`, `name`, `deleted`) USING BTREE
) ENGINE = InnoDB DEFAULT CHARSET = utf8mb4 COLLATE = utf8mb4_bin COMMENT = '应用配置';

CREATE TABLE IF NOT EXISTS `config_version` (
  `id` varchar(20) CHARACTER SET utf8mb4 COLLATE utf8mb4_bin NOT NULL COMMENT 'ID',
  `config_id` varchar(20) CHARACTER SET utf8mb4 COLLATE utf8mb4_bin NOT NULL COMMENT '配置ID',
  `data_type` varchar(20) CHARACTER SET utf8mb4 COLLATE utf8mb4_bin DEFAULT NULL COMMENT '数据类型:yaml、json、text、xml、html、properties、toml',
  `version` varchar(255) CHARACTER SET utf8mb4 COLLATE utf8mb4_bin DEFAULT NULL COMMENT '版本信息',
  `active` TINYINT(1) UNSIGNED NOT NULL DEFAULT '0' COMMENT '是否使用中',
  `content` text CHARACTER SET utf8mb4 COLLATE utf8mb4_bin DEFAULT NULL COMMENT '配置信息',
  `type` varchar(64) CHARACTER SET utf8mb4 COLLATE utf8mb4_bin DEFAULT 'release' COMMENT '正式发布: release、灰度发布: beta',
  `status` varchar(20) CHARACTER SET utf8mb4 COLLATE utf8mb4_bin DEFAULT NULL COMMENT '状态: enabled-启用, disabled-禁用',
  `description` varchar(255) CHARACTER SET utf8mb4 COLLATE utf8mb4_bin DEFAULT NULL COMMENT '备注',
  `tenant` varchar(255) CHARACTER SET utf8mb4 COLLATE utf8mb4_bin DEFAULT NULL COMMENT '租户信息',
  `creator` varchar(255) CHARACTER SET utf8mb4 COLLATE utf8mb4_bin DEFAULT NULL COMMENT '创建者',
  `created_at` timestamp NOT NULL DEFAULT CURRENT_TIMESTAMP COMMENT '创建时间',
  `updated_at` timestamp DEFAULT CURRENT_TIMESTAMP ON UPDATE CURRENT_TIMESTAMP COMMENT '更新时间',
  `deleted` varchar(20) CHARACTER SET utf8mb4 COLLATE utf8mb4_bin NOT NULL DEFAULT '0' COMMENT '逻辑删除标识',
  `deleted_at` timestamp NULL DEFAULT NULL COMMENT '删除时间',
  PRIMARY KEY (`id`),
  UNIQUE KEY `uniq_config_id_version` (`config_id`, `version`, `deleted`) USING BTREE
) ENGINE = InnoDB DEFAULT CHARSET = utf8mb4 COLLATE = utf8mb4_bin COMMENT = '应用配置版本';

CREATE TABLE IF NOT EXISTS `config_change_order` (
  `id` varchar(20) CHARACTER SET utf8mb4 COLLATE utf8mb4_bin NOT NULL COMMENT 'ID',
  `config_id` varchar(20) CHARACTER SET utf8mb4 COLLATE utf8mb4_bin NOT NULL COMMENT '配置ID',
  `config_center_id` varchar(255) CHARACTER SET utf8mb4 COLLATE utf8mb4_bin DEFAULT NULL COMMENT '配置中心id',
  `config_version_id` varchar(20) CHARACTER SET utf8mb4 COLLATE utf8mb4_bin NOT NULL COMMENT '配置版本ID',
  `config_version` varchar(255) CHARACTER SET utf8mb4 COLLATE utf8mb4_bin NOT NULL COMMENT '配置版本, 冗余存储以便查询',
  `type` varchar(64) CHARACTER SET utf8mb4 COLLATE utf8mb4_bin NOT NULL COMMENT '操作类型, 发布: release、回滚: rollback、灰度: beta',
  `description` varchar(255) CHARACTER SET utf8mb4 COLLATE utf8mb4_bin DEFAULT NULL COMMENT '备注',
  `tenant` varchar(255) CHARACTER SET utf8mb4 COLLATE utf8mb4_bin DEFAULT NULL COMMENT '租户信息',
  `creator` varchar(255) CHARACTER SET utf8mb4 COLLATE utf8mb4_bin DEFAULT NULL COMMENT '创建者',
  `created_at` timestamp NOT NULL DEFAULT CURRENT_TIMESTAMP COMMENT '创建时间',
  `deleted` varchar(20) CHARACTER SET utf8mb4 COLLATE utf8mb4_bin NOT NULL DEFAULT '0' COMMENT '逻辑删除标识',
  `deleted_at` timestamp NULL DEFAULT NULL COMMENT '删除时间',
  PRIMARY KEY (`id`),
  KEY `idx_config_id_config_version` (`config_id`, `config_version`) USING BTREE
) ENGINE = InnoDB DEFAULT CHARSET = utf8mb4 COLLATE = utf8mb4_bin COMMENT = '配置变更单';

CREATE TABLE IF NOT EXISTS `config_beta_policy` (
  `id` varchar(20) CHARACTER SET utf8mb4 COLLATE utf8mb4_bin NOT NULL COMMENT 'ID',
  `config_id` varchar(20) CHARACTER SET utf8mb4 COLLATE utf8mb4_bin NOT NULL COMMENT '配置ID',
  `config_change_order_id` varchar(20) CHARACTER SET utf8mb4 COLLATE utf8mb4_bin NOT NULL COMMENT '配置变更id',
  `config_version_id` varchar(20) CHARACTER SET utf8mb4 COLLATE utf8mb4_bin NOT NULL COMMENT '配置版本ID',
  `config_version` varchar(255) CHARACTER SET utf8mb4 COLLATE utf8mb4_bin NOT NULL COMMENT '配置版本, 冗余存储以便查询',
  `ips` json DEFAULT NULL COMMENT 'IP列表',
  `description` varchar(255) CHARACTER SET utf8mb4 COLLATE utf8mb4_bin DEFAULT NULL COMMENT '备注',
  `tenant` varchar(255) CHARACTER SET utf8mb4 COLLATE utf8mb4_bin DEFAULT NULL COMMENT '租户信息',
  `creator` varchar(255) CHARACTER SET utf8mb4 COLLATE utf8mb4_bin DEFAULT NULL COMMENT '创建者',
  `created_at` timestamp NOT NULL DEFAULT CURRENT_TIMESTAMP COMMENT '创建时间',
  `deleted` varchar(20) CHARACTER SET utf8mb4 COLLATE utf8mb4_bin NOT NULL DEFAULT '0' COMMENT '逻辑删除标识',
  `deleted_at` timestamp NULL DEFAULT NULL COMMENT '删除时间',
  PRIMARY KEY (`id`),
  KEY `idx_config_id_config_version` (`config_id`, `config_version`) USING BTREE
) ENGINE = InnoDB DEFAULT CHARSET = utf8mb4 COLLATE = utf8mb4_bin COMMENT = '配置灰度策略';

CREATE TABLE IF NOT EXISTS `config_temp` (
  `id` varchar(20) CHARACTER SET utf8mb4 COLLATE utf8mb4_bin NOT NULL COMMENT 'ID',
  `name` varchar(255) CHARACTER SET utf8mb4 COLLATE utf8mb4_bin NOT NULL COMMENT '配置模板名称',
  `data_type` varchar(20) CHARACTER SET utf8mb4 COLLATE utf8mb4_bin DEFAULT NULL COMMENT '数据类型:yaml、json、text、xml、html、properties、toml',
  `content` text CHARACTER SET utf8mb4 COLLATE utf8mb4_bin DEFAULT NULL COMMENT '配置信息',
  `description` varchar(255) CHARACTER SET utf8mb4 COLLATE utf8mb4_bin DEFAULT NULL COMMENT '备注',
  `tenant` varchar(255) CHARACTER SET utf8mb4 COLLATE utf8mb4_bin DEFAULT NULL COMMENT '租户信息',
  `creator` varchar(255) CHARACTER SET utf8mb4 COLLATE utf8mb4_bin DEFAULT NULL COMMENT '创建者',
  `created_at` timestamp NOT NULL DEFAULT CURRENT_TIMESTAMP COMMENT '创建时间',
  `updated_at` timestamp DEFAULT CURRENT_TIMESTAMP ON UPDATE CURRENT_TIMESTAMP COMMENT '更新时间',
  `deleted` varchar(20) CHARACTER SET utf8mb4 COLLATE utf8mb4_bin NOT NULL DEFAULT '0' COMMENT '逻辑删除标识',
  `deleted_at` timestamp NULL DEFAULT NULL COMMENT '删除时间',
  PRIMARY KEY (`id`)
) ENGINE = InnoDB DEFAULT CHARSET = utf8mb4 COLLATE = utf8mb4_bin COMMENT = '配置模板';

CREATE TABLE IF NOT EXISTS `menu` (
  `id` varchar(20) NOT NULL COMMENT 'ID',
  `code` varchar(32) DEFAULT NULL COMMENT '菜单编码',
  `name` varchar(128) DEFAULT NULL COMMENT '菜单名称',
  `description` varchar(1024) DEFAULT NULL COMMENT '描述',
  `sequence` bigint DEFAULT NULL COMMENT '序列',
  `type` varchar(20) DEFAULT NULL COMMENT '类型: page, button',
  `path` varchar(255) DEFAULT NULL COMMENT '路径',
  `properties` text COMMENT '属性',
  `status` varchar(20) DEFAULT NULL COMMENT '状态',
  `parent_id` varchar(20) DEFAULT NULL COMMENT '父ID',
  `parent_path` varchar(255) DEFAULT NULL COMMENT '父路径',
  `created_at` datetime(3) DEFAULT NULL COMMENT '创建时间',
  `updated_at` datetime(3) DEFAULT NULL COMMENT '更新时间',
  PRIMARY KEY (`id`),
  KEY `idx_menu_parent_path` (`parent_path`),
  KEY `idx_menu_name` (`name`),
  KEY `idx_menu_status` (`status`),
  KEY `idx_menu_type` (`type`),
  KEY `idx_menu_parent_id` (`parent_id`),
  KEY `idx_menu_code` (`code`),
  KEY `idx_menu_sequence` (`sequence`)
) ENGINE = InnoDB DEFAULT CHARSET = utf8mb4 COLLATE = utf8mb4_bin COMMENT = '菜单';

CREATE TABLE IF NOT EXISTS `menu_resource` (
  `id` varchar(20) NOT NULL COMMENT 'ID',
  `menu_id` varchar(20) DEFAULT NULL COMMENT '菜单ID',
  `method` varchar(20) DEFAULT NULL COMMENT '请求方法',
  `path` varchar(255) DEFAULT NULL COMMENT '请求路径',
  `created_at` datetime(3) DEFAULT NULL COMMENT '创建时间',
  `updated_at` datetime(3) DEFAULT NULL COMMENT '更新时间',
  PRIMARY KEY (`id`),
  KEY `idx_menu_resource_menu_id` (`menu_id`)
) ENGINE = InnoDB DEFAULT CHARSET = utf8mb4 COLLATE = utf8mb4_bin COMMENT = '菜单资源';

CREATE TABLE IF NOT EXISTS `menu_resource_group` (
  `id` varchar(20) NOT NULL COMMENT 'ID',
  `code` varchar(255) CHARACTER SET utf8mb4 COLLATE utf8mb4_bin DEFAULT NULL COMMENT '权限编码',
  `name` varchar(255) CHARACTER SET utf8mb4 COLLATE utf8mb4_bin DEFAULT NULL COMMENT '权限名称',
  `description` varchar(500) CHARACTER SET utf8mb4 COLLATE utf8mb4_bin DEFAULT NULL COMMENT '权限描述',
  `content` json DEFAULT NULL COMMENT '资源列表[menu_id]',
  `order_num` int DEFAULT NULL COMMENT '前端展示排序',
  `depends_on` json DEFAULT NULL COMMENT '授权依赖其他权限编码,数组结构',
  `base_auth` tinyint(1) DEFAULT '0' COMMENT '是否基础默认权限,前端默认勾选',
  `created_at` timestamp NOT NULL DEFAULT CURRENT_TIMESTAMP COMMENT '创建时间',
  `updated_at` timestamp NULL DEFAULT NULL COMMENT '更新时间',
  PRIMARY KEY (`id`),
  UNIQUE KEY `uniq_code` (`code`)
) ENGINE = InnoDB DEFAULT CHARSET = utf8mb4 COLLATE = utf8mb4_bin COMMENT = '菜单资源分组';

CREATE TABLE IF NOT EXISTS `role` (
  `id` varchar(20) NOT NULL COMMENT 'ID',
  `code` varchar(32) DEFAULT NULL COMMENT '角色编码',
  `name` varchar(128) DEFAULT NULL COMMENT '角色名称',
  `description` varchar(1024) DEFAULT NULL COMMENT '角色描述',
  `sequence` bigint DEFAULT NULL COMMENT '角色序列',
  `tenant` varchar(255) DEFAULT NULL COMMENT '租户信息',
  `status` varchar(20) DEFAULT NULL COMMENT '状态',
  `created_at` datetime(3) DEFAULT NULL COMMENT '创建时间',
  `updated_at` datetime(3) DEFAULT NULL COMMENT '更新时间',
  PRIMARY KEY (`id`),
  KEY `idx_role_code` (`code`),
  KEY `idx_role_name` (`name`),
  KEY `idx_role_sequence` (`sequence`),
  KEY `idx_role_status` (`status`)
) ENGINE = InnoDB DEFAULT CHARSET = utf8mb4 COLLATE = utf8mb4_bin COMMENT = '角色';

CREATE TABLE IF NOT EXISTS `role_menu` (
  `id` varchar(20) NOT NULL COMMENT 'ID',
  `role_id` varchar(20) DEFAULT NULL COMMENT '角色ID',
  `menu_group_id` varchar(20) DEFAULT NULL COMMENT '菜单组ID',
  `created_at` datetime(3) DEFAULT NULL COMMENT '创建时间',
  `updated_at` datetime(3) DEFAULT NULL COMMENT '更新时间',
  PRIMARY KEY (`id`),
  KEY `idx_role_menu_role_id` (`role_id`),
  KEY `idx_role_menu_menu_id` (`menu_group_id`)
) ENGINE = InnoDB DEFAULT CHARSET = utf8mb4 COLLATE = utf8mb4_bin COMMENT = '角色菜单';

CREATE TABLE IF NOT EXISTS `user` (
  `id` varchar(20) NOT NULL COMMENT 'ID',
  `username` varchar(64) DEFAULT NULL COMMENT '用户名',
  `name` varchar(64) DEFAULT NULL COMMENT '用户名称',
  `password` varchar(64) DEFAULT NULL COMMENT '密码',
  `phone` varchar(32) DEFAULT NULL COMMENT '电话',
  `email` varchar(128) DEFAULT NULL COMMENT '邮件',
  `remark` varchar(1024) DEFAULT NULL COMMENT '备注',
  `tenant` varchar(255) DEFAULT NULL COMMENT '租户信息',
  `status` varchar(20) DEFAULT NULL COMMENT '状态',
  `created_at` datetime(3) DEFAULT NULL COMMENT '创建时间',
  `updated_at` datetime(3) DEFAULT NULL COMMENT '更新时间',
  PRIMARY KEY (`id`),
  KEY `idx_user_username` (`username`),
  KEY `idx_user_name` (`name`),
  KEY `idx_user_status` (`status`)
) ENGINE = InnoDB DEFAULT CHARSET = utf8mb4 COLLATE = utf8mb4_bin COMMENT = '用户';

CREATE TABLE IF NOT EXISTS `user_role` (
  `id` varchar(20) NOT NULL COMMENT 'ID',
  `user_id` varchar(20) DEFAULT NULL COMMENT '用户ID',
  `role_id` varchar(20) DEFAULT NULL COMMENT '角色ID',
  `created_at` datetime(3) DEFAULT NULL COMMENT '创建时间',
  `updated_at` datetime(3) DEFAULT NULL COMMENT '更新时间',
  PRIMARY KEY (`id`),
  KEY `idx_user_role_user_id` (`user_id`),
  KEY `idx_user_role_role_id` (`role_id`)
) ENGINE = InnoDB DEFAULT CHARSET = utf8mb4 COLLATE = utf8mb4_bin COMMENT = '用户角色';

CREATE TABLE IF NOT EXISTS `casbin_rule` (
  `id` bigint(11) NOT NULL AUTO_INCREMENT COMMENT 'ID',
  `ptype` varchar(100) DEFAULT NULL COMMENT '策略类型（p 或 g）',
  `v0` varchar(100) DEFAULT NULL COMMENT 'v0',
  `v1` varchar(100) DEFAULT NULL COMMENT 'v1',
  `v2` varchar(100) DEFAULT NULL COMMENT 'v2',
  `v3` varchar(100) DEFAULT NULL COMMENT 'method',
  `v4` varchar(100) DEFAULT NULL COMMENT 'v4',
  `v5` varchar(100) DEFAULT NULL COMMENT 'v5',
  PRIMARY KEY (`id`)
) ENGINE = InnoDB DEFAULT CHARSET = utf8mb4 COLLATE = utf8mb4_bin COMMENT = '鉴权引擎规则';

CREATE TABLE IF NOT EXISTS `data_permission` (
  `id` varchar(20) CHARACTER SET utf8mb4 COLLATE utf8mb4_bin NOT NULL COMMENT 'ID',
  `type` varchar(50) CHARACTER SET utf8mb4 COLLATE utf8mb4_bin NOT NULL COMMENT '数据类型(表名)',
  `data_id` varchar(20) CHARACTER SET utf8mb4 COLLATE utf8mb4_bin NOT NULL COMMENT '数据ID',
  `user` varchar(50) CHARACTER SET utf8mb4 COLLATE utf8mb4_bin NOT NULL COMMENT '用户',
  `tenant` varchar(50) CHARACTER SET utf8mb4 COLLATE utf8mb4_bin NOT NULL COMMENT '租户',
  `role` varchar(20) CHARACTER SET utf8mb4 COLLATE utf8mb4_bin NOT NULL COMMENT '角色编码',
  `permission` tinyint unsigned NOT NULL DEFAULT 0 COMMENT '数据权限位 - 格式(read,write,delete)',
  `created_at` timestamp NOT NULL DEFAULT CURRENT_TIMESTAMP COMMENT '创建时间',
  `updated_at` timestamp DEFAULT CURRENT_TIMESTAMP ON UPDATE CURRENT_TIMESTAMP COMMENT '更新时间',
  `deleted` varchar(20) CHARACTER SET utf8mb4 COLLATE utf8mb4_bin NOT NULL DEFAULT '0' COMMENT '逻辑删除标识',
  `deleted_at` timestamp NULL DEFAULT NULL COMMENT '删除时间',
  PRIMARY KEY (`id`),
  UNIQUE KEY `uniq_data_permission` (
    `type`,
    `data_id`,
    `user`,
    `tenant`,
    `role`,
    `deleted`
  ) USING BTREE,
  KEY `idx_type` (`type`) USING BTREE,
  KEY `idx_data_id` (`data_id`) USING BTREE,
  KEY `idx_user` (`user`) USING BTREE
) ENGINE = InnoDB DEFAULT CHARSET = utf8mb4 COLLATE = utf8mb4_bin COMMENT = '数据权限';

CREATE TABLE IF NOT EXISTS `lane_space` (
  `id` VARCHAR(20) COLLATE utf8mb4_bin NOT NULL COMMENT 'ID',
  `created_at` TIMESTAMP NOT NULL DEFAULT CURRENT_TIMESTAMP COMMENT '创建时间',
  `updated_at` TIMESTAMP DEFAULT CURRENT_TIMESTAMP ON UPDATE CURRENT_TIMESTAMP COMMENT '更新时间',
  `deleted` VARCHAR(20) COLLATE utf8mb4_bin NOT NULL DEFAULT '0' COMMENT '逻辑删除标识',
  `deleted_at` TIMESTAMP NULL DEFAULT NULL COMMENT '删除时间',
  `tenant` VARCHAR(255) COLLATE utf8mb4_bin NOT NULL COMMENT '租户',
  `name` VARCHAR(255) COLLATE utf8mb4_bin NOT NULL COMMENT '名称',
  `version` bigint NOT NULL DEFAULT '1' COMMENT '版本号',
  `description` VARCHAR(1024) COLLATE utf8mb4_bin DEFAULT NULL COMMENT '备注',
  PRIMARY KEY (`id`),
  KEY `idx_name` (`name`) USING BTREE
) ENGINE = InnoDB DEFAULT CHARSET = utf8mb4 COLLATE = utf8mb4_bin COMMENT = '泳道组';

CREATE TABLE IF NOT EXISTS `lane` (
  `id` VARCHAR(20) COLLATE utf8mb4_bin NOT NULL COMMENT 'ID',
  `created_at` TIMESTAMP NOT NULL DEFAULT CURRENT_TIMESTAMP COMMENT '创建时间',
  `updated_at` TIMESTAMP DEFAULT CURRENT_TIMESTAMP ON UPDATE CURRENT_TIMESTAMP COMMENT '更新时间',
  `deleted` VARCHAR(20) COLLATE utf8mb4_bin NOT NULL DEFAULT '0' COMMENT '逻辑删除标识',
  `deleted_at` TIMESTAMP NULL DEFAULT NULL COMMENT '删除时间',
  `tenant` VARCHAR(255) COLLATE utf8mb4_bin NOT NULL COMMENT '租户',
  `lane_space_id` VARCHAR(32) COLLATE utf8mb4_bin NOT NULL COMMENT '泳道组id',
  `name` VARCHAR(255) COLLATE utf8mb4_bin NOT NULL COMMENT '名称',
  `code` VARCHAR(64) COLLATE utf8mb4_bin NOT NULL COMMENT '编码',
  `type` VARCHAR(64) COLLATE utf8mb4_bin NOT NULL COMMENT '泳道类型',
  `services` VARCHAR(1024) COLLATE utf8mb4_bin NULL COMMENT 'services,已改为应用名列表',
  `space_code` VARCHAR(255) CHARACTER SET utf8mb4 COLLATE utf8mb4_bin DEFAULT NULL COMMENT '微服务空间编码',
  `service_groups` json DEFAULT NULL COMMENT 'service与group的映射关系',
  `enabled` TINYINT(1) UNSIGNED NOT NULL DEFAULT '1' COMMENT '是否开启泳道',
  `fallback_type` VARCHAR(255) COLLATE utf8mb4_bin NOT NULL COMMENT '容错类型, REJECT:不容错, DEFAULT:基线容错, CUSTOM:指定泳道容错',
  `fallback_lane_id` VARCHAR(255) COLLATE utf8mb4_bin NULL COMMENT '容错泳道id',
  `fallback_lane_code` VARCHAR(255) COLLATE utf8mb4_bin NULL COMMENT '容错泳道code',
  `context_propagation_type` varchar(255) COLLATE utf8mb4_bin NULL COMMENT '上下文透传类型,TraceID: 调用链；CustomHeader: 自定义请求头',
  `context_propagation_header` varchar(255) COLLATE utf8mb4_bin NULL COMMENT '上下文透传Header名称',
  `decorator` varchar(64) COLLATE utf8mb4_bin NULL COMMENT '域名装饰',
  `description` VARCHAR(1024) COLLATE utf8mb4_bin DEFAULT NULL COMMENT '备注',
  PRIMARY KEY (`id`),
  KEY `idx_lane_space_id` (`lane_space_id`) USING BTREE,
  KEY `idx_name` (`name`) USING BTREE,
  KEY `idx_code` (`code`) USING BTREE,
  KEY `idx_fallback_lane_id` (`fallback_lane_id`) USING BTREE
) ENGINE = InnoDB DEFAULT CHARSET = utf8mb4 COLLATE = utf8mb4_bin COMMENT = '泳道';

CREATE TABLE IF NOT EXISTS `lane_rule` (
  `id` VARCHAR(20) COLLATE utf8mb4_bin NOT NULL COMMENT 'ID',
  `created_at` TIMESTAMP NOT NULL DEFAULT CURRENT_TIMESTAMP COMMENT '创建时间',
  `updated_at` TIMESTAMP DEFAULT CURRENT_TIMESTAMP ON UPDATE CURRENT_TIMESTAMP COMMENT '更新时间',
  `lane_id` VARCHAR(32) COLLATE utf8mb4_bin NOT NULL COMMENT '泳道组id',
  `condition_op_model` VARCHAR(32) COLLATE utf8mb4_bin NOT NULL COMMENT 'Condition运算模式, 与: and, 或: or',
  `conditions` TEXT COLLATE utf8mb4_bin NULL COMMENT '灰度规则Condition',
  `host` VARCHAR(255) DEFAULT NULL COMMENT '入口域名',
  `ratio` int(10) DEFAULT NULL COMMENT '流量比例',
  `enabled` TINYINT(1) UNSIGNED NOT NULL DEFAULT '1' COMMENT '是否开启泳道规则',
  PRIMARY KEY (`id`),
  KEY `idx_lane_id` (`lane_id`) USING BTREE
) ENGINE = InnoDB DEFAULT CHARSET = utf8mb4 COLLATE = utf8mb4_bin COMMENT = '泳道引流规则';

CREATE TABLE IF NOT EXISTS `k8s_cluster` (
  `id` varchar(20) CHARACTER SET utf8mb4 COLLATE utf8mb4_bin NOT NULL COMMENT 'ID',
  `name` varchar(255) CHARACTER SET utf8mb4 COLLATE utf8mb4_bin NOT NULL COMMENT '应用名称',
  `region` varchar(255) CHARACTER SET utf8mb4 COLLATE utf8mb4_bin DEFAULT NULL COMMENT '地域编码',
  `region_name` varchar(255) CHARACTER SET utf8mb4 COLLATE utf8mb4_bin DEFAULT NULL COMMENT '地域名称',
  `zones` varchar(255) CHARACTER SET utf8mb4 COLLATE utf8mb4_bin DEFAULT NULL COMMENT '多可用区信息',
  `arch` varchar(255) CHARACTER SET utf8mb4 COLLATE utf8mb4_bin DEFAULT NULL COMMENT '集群架构',
  `version` varchar(255) CHARACTER SET utf8mb4 COLLATE utf8mb4_bin DEFAULT NULL COMMENT '集群版本',
  `kube_config` text CHARACTER SET utf8mb4 COLLATE utf8mb4_bin DEFAULT NULL COMMENT '集群kubeconfig',
  `tenant` varchar(255) CHARACTER SET utf8mb4 COLLATE utf8mb4_bin DEFAULT NULL COMMENT '租户',
  `creator` varchar(255) CHARACTER SET utf8mb4 COLLATE utf8mb4_bin DEFAULT NULL COMMENT '创建者',
  `description` varchar(255) CHARACTER SET utf8mb4 COLLATE utf8mb4_bin DEFAULT NULL COMMENT '备注',
  `created_at` timestamp NOT NULL DEFAULT CURRENT_TIMESTAMP COMMENT '创建时间',
  `updated_at` timestamp DEFAULT CURRENT_TIMESTAMP ON UPDATE CURRENT_TIMESTAMP COMMENT '更新时间',
  `deleted` varchar(20) CHARACTER SET utf8mb4 COLLATE utf8mb4_bin NOT NULL DEFAULT '0' COMMENT '逻辑删除标识',
  `deleted_at` timestamp NULL DEFAULT NULL COMMENT '删除时间',
  PRIMARY KEY (`id`),
  UNIQUE KEY `uniq_name` (`name`, `deleted`) USING BTREE
) ENGINE = InnoDB DEFAULT CHARSET = utf8mb4 COLLATE = utf8mb4_bin COMMENT = '集群';

CREATE TABLE IF NOT EXISTS `setting` (
  `id` varchar(20) NOT NULL COMMENT 'ID',
  `space_code` varchar(255) DEFAULT NULL COMMENT '微服务空间编码',
  `application_id` varchar(20) DEFAULT NULL COMMENT '应用ID',
  `component_version_id` varchar(20) DEFAULT NULL COMMENT '组件版本ID',
  `setting_scope_id` varchar(20) DEFAULT NULL COMMENT '配置域ID',
  `type` varchar(20) NOT NULL COMMENT '配置项类型: label, environment, file',
  `key` varchar(255) NOT NULL COMMENT '配置项',
  `value` text COMMENT '配置内容',
  `tenant` varchar(255) DEFAULT NULL COMMENT '租户',
  `creator` varchar(255) DEFAULT NULL COMMENT '创建者',
  `modifier` varchar(255) DEFAULT NULL COMMENT '修改者',
  `description` varchar(255) DEFAULT NULL COMMENT '备注',
  `created_at` timestamp NOT NULL DEFAULT CURRENT_TIMESTAMP COMMENT '创建时间',
  `updated_at` timestamp DEFAULT CURRENT_TIMESTAMP ON UPDATE CURRENT_TIMESTAMP COMMENT '更新时间',
  `deleted` varchar(20) NOT NULL DEFAULT '0' COMMENT '逻辑删除标识',
  `deleted_at` timestamp NULL DEFAULT NULL COMMENT '删除时间',
  PRIMARY KEY (`id`),
  UNIQUE KEY `uniq_setting` (
    `key`,
    `space_code`,
    `application_id`,
    `component_version_id`,
    `setting_scope_id`,
    `type`,
    `deleted`
  ),
  KEY `idx_space_code` (`space_code`),
  KEY `idx_application_id` (`application_id`),
  KEY `idx_type` (`type`)
) ENGINE = InnoDB DEFAULT CHARSET = utf8mb4 COLLATE = utf8mb4_bin COMMENT = '基础配置';

CREATE TABLE IF NOT EXISTS `setting_scope` (
  `id` varchar(20) NOT NULL COMMENT 'ID',
  `scope_type` varchar(20) DEFAULT NULL COMMENT '''配置域类型(global,microservice,application,workload)',
  `component_type` varchar(20) DEFAULT NULL COMMENT '组件类型(istio,agent,nacos)',
  `component_id` varchar(20) DEFAULT NULL COMMENT '组件ID',
  `component_version_id` varchar(20) DEFAULT NULL COMMENT '组件版本ID',
  `cluster_id` varchar(64) DEFAULT NULL COMMENT 'k8s集群ID',
  `space_code` varchar(255) DEFAULT NULL COMMENT '空间代码',
  `service_id` varchar(20) DEFAULT NULL COMMENT '微服务ID',
  `service_name` varchar(255) DEFAULT NULL COMMENT '微服务空间编码',
  `application_id` varchar(20) DEFAULT NULL COMMENT '应用ID',
  `workload_name` varchar(255) DEFAULT NULL COMMENT '工作负载名称',
  `pod_name` varchar(255) DEFAULT NULL COMMENT 'POD名称',
  `tenant` varchar(255) DEFAULT NULL COMMENT '租户',
  `creator` varchar(255) DEFAULT NULL COMMENT '创建人',
  `description` varchar(255) DEFAULT NULL COMMENT '描述',
  `created_at` datetime(3) DEFAULT NULL COMMENT '创建时间',
  `updated_at` datetime(3) DEFAULT NULL COMMENT '更新时间',
  `deleted` varchar(20) DEFAULT NULL COMMENT '是否删除',
  `deleted_at` datetime(3) DEFAULT NULL COMMENT '删除时间',
  PRIMARY KEY (`id`)
) ENGINE = InnoDB DEFAULT CHARSET = utf8mb4 COLLATE = utf8mb4_bin COMMENT = '配置域';

-- ============================================================
-- API版本表: api_version
-- ============================================================
CREATE TABLE IF NOT EXISTS api_version (
  id VARCHAR(20) PRIMARY KEY COMMENT '主键ID',
  name VARCHAR(100) NOT NULL COMMENT '服务名称,如 user-service',
  service_id VARCHAR(20) NOT NULL COMMENT '服务ID',
  description TEXT COMMENT '服务描述',
  version VARCHAR(50) NOT NULL COMMENT '版本号,如 1.0.0',
  source_url VARCHAR(1024) DEFAULT NULL COMMENT 'Swagger JSON URL',
  base_path VARCHAR(1024) DEFAULT NULL COMMENT 'API 基础路径,如 /api/v1',
  raw_json LONGTEXT COMMENT '原始 Swagger JSON（可选）',
  common_json LONGTEXT COMMENT '原始 Swagger JSON（除paths）',
  created_at DATETIME DEFAULT CURRENT_TIMESTAMP COMMENT '创建时间',
  updated_at DATETIME DEFAULT CURRENT_TIMESTAMP ON UPDATE CURRENT_TIMESTAMP COMMENT '更新时间',
  deleted VARCHAR(20) NOT NULL DEFAULT '0' COMMENT '逻辑删除标识',
  deleted_at DATETIME NULL DEFAULT NULL COMMENT '删除时间',
  UNIQUE KEY uk_service_version (service_id, version, deleted)
) ENGINE = InnoDB DEFAULT CHARSET = utf8mb4 COLLATE = utf8mb4_bin COMMENT = 'Swagger API版本表';

-- ============================================================
-- API接口表: api_endpoint
-- ============================================================
CREATE TABLE IF NOT EXISTS api_endpoint (
  id VARCHAR(20) PRIMARY KEY COMMENT '主键ID',
  api_version_id VARCHAR(20) NOT NULL COMMENT '关联API版本ID',
  method VARCHAR(255) NOT NULL COMMENT '请求方法,如 GET/POST (缩短以适应索引)',
  path VARCHAR(512) NOT NULL COMMENT '请求路径,如 /user/{id} (缩短以适应索引)',
  summary VARCHAR(1024) DEFAULT NULL COMMENT '接口简要描述',
  description TEXT COMMENT '接口详细描述',
  deprecated TINYINT(1) DEFAULT 0 COMMENT '是否废弃',
  tag VARCHAR(1024) DEFAULT NULL COMMENT '所属标签',
  operation_id VARCHAR(255) NOT NULL COMMENT 'Swagger 的 operationId, 服务-版本下唯一',
  consumes VARCHAR(1024) DEFAULT NULL COMMENT 'Swagger 的 consumes',
  produces VARCHAR(1024) DEFAULT NULL COMMENT 'Swagger 的 produces',
  parameters JSON DEFAULT NULL COMMENT '接口参数列表',
  responses JSON DEFAULT NULL COMMENT '接口响应',
  raw_json LONGTEXT COMMENT '原始 Swagger JSON（可选）',
  created_at DATETIME DEFAULT CURRENT_TIMESTAMP COMMENT '创建时间',
  updated_at DATETIME DEFAULT CURRENT_TIMESTAMP ON UPDATE CURRENT_TIMESTAMP COMMENT '更新时间',
  deleted VARCHAR(20) NOT NULL DEFAULT '0' COMMENT '逻辑删除标识',
  deleted_at DATETIME NULL DEFAULT NULL COMMENT '删除时间',
  UNIQUE KEY uk_version_operation_id (api_version_id, operation_id, deleted),
  CONSTRAINT fk_endpoint_version FOREIGN KEY (api_version_id) REFERENCES api_version (id) ON DELETE CASCADE
) ENGINE = InnoDB DEFAULT CHARSET = utf8mb4 COLLATE = utf8mb4_bin COMMENT = 'API接口定义表';