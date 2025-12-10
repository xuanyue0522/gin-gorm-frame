# 管理员表
CREATE TABLE `admin_user` (
      `id` bigint NOT NULL AUTO_INCREMENT COMMENT '主键-管理员ID',
      `name` varchar(255) CHARACTER SET utf8mb4 COLLATE utf8mb4_general_ci NOT NULL COMMENT '管理员名称',
      `username` varchar(255) CHARACTER SET utf8mb4 COLLATE utf8mb4_general_ci NOT NULL COMMENT '管理员用户名',
      `mobile` varchar(255) CHARACTER SET utf8mb4 COLLATE utf8mb4_general_ci NOT NULL COMMENT '手机号',
      `password` varchar(255) CHARACTER SET utf8mb4 COLLATE utf8mb4_general_ci NOT NULL COMMENT '管理员密码',
      `status` tinyint NOT NULL COMMENT '状态：1-正常，-1-禁用',
      `created_at` bigint NOT NULL  COMMENT '创建时间',
      `updated_at` bigint NOT NULL  COMMENT '更新时间',
      `deleted_at` bigint DEFAULT NULL COMMENT '删除时间',
      `create_by` bigint NOT NULL DEFAULT 0 COMMENT '创建人',
      `update_by` bigint NOT NULL DEFAULT 0 COMMENT '更新人',
      `sex` tinyint NOT NULL DEFAULT 3 COMMENT '性别：1-男，2-女,3-其他',

      PRIMARY KEY (`id`) USING BTREE,
      UNIQUE KEY `idx_mobile` (`mobile`) USING BTREE,
      KEY `idx_name` (`name`) USING BTREE
) ENGINE=InnoDB AUTO_INCREMENT=1 DEFAULT CHARSET=utf8mb4 COLLATE=utf8mb4_general_ci ROW_FORMAT=DYNAMIC COMMENT='管理员表';
