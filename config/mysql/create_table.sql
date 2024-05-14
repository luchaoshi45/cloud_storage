CREATE TABLE File (
    sha1 VARCHAR(255) PRIMARY KEY,
    name VARCHAR(255) NOT NULL,
    dir VARCHAR(255) NOT NULL,
    size BIGINT NOT NULL,
    upload_at DATETIME NOT NULL,
    created_at TIMESTAMP DEFAULT CURRENT_TIMESTAMP,
    updated_at TIMESTAMP DEFAULT CURRENT_TIMESTAMP ON UPDATE CURRENT_TIMESTAMP
);

CREATE TABLE User (
    id INT(11) NOT NULL AUTO_INCREMENT COMMENT '用户名ID',
    user_name VARCHAR(64) NOT NULL DEFAULT '' COMMENT '用户名',
    user_pwd VARCHAR(256) NOT NULL DEFAULT '' COMMENT '用户encoded密码',
    email VARCHAR(64) DEFAULT NULL COMMENT '邮箱',
    phone VARCHAR(128) DEFAULT NULL COMMENT '手机号',
    email_validated TINYINT(1) DEFAULT 0 COMMENT '邮箱是否已验证',
    phone_validated TINYINT(1) DEFAULT 0 COMMENT '手机号是否已验证',
    signup_at DATETIME DEFAULT CURRENT_TIMESTAMP COMMENT '注册日期',
    last_active DATETIME DEFAULT CURRENT_TIMESTAMP ON UPDATE CURRENT_TIMESTAMP COMMENT '最后活跃时间',
    profile TEXT COMMENT'用户属性',
    status INT(11) NOT NULL DEFAULT 0 COMMENT '账户状态(启用/禁用/锁定/标记删除',
    PRIMARY KEY (id),
    UNIQUE KEY idx_user_name (user_name),
    UNIQUE KEY idx_phone (phone),
    KEY idx_status (status)
)ENGINE=InnoDB AUTO_INCREMENT=5 DEFAULT CHARSET=utf8mb4;

CREATE TABLE UserToken (
    id INT(11) NOT NULL AUTO_INCREMENT,
    user_name VARCHAR(64) NOT NULL DEFAULT '' COMMENT '用户名',
    token CHAR(40) NOT NULL DEFAULT '' COMMENT '用户登录token',
    PRIMARY KEY (id),
    UNIQUE KEY idx_user_name (user_name)
) ENGINE=InnoDB DEFAULT CHARSET=utf8mb4 COLLATE=utf8mb4_general_ci;

CREATE TABLE UserFile (
    id INT(11) NOT NULL PRIMARY KEY AUTO_INCREMENT,
    user_name VARCHAR(64) NOT NULL COMMENT '用户名',
    file_sha1 VARCHAR(64) NOT NULL DEFAULT '' COMMENT '文件哈希',
    file_size INT(11) DEFAULT 0 COMMENT '文件大小',
    file_name VARCHAR(256) NOT NULL DEFAULT '' COMMENT '文件名',
    upload_at DATETIME DEFAULT CURRENT_TIMESTAMP COMMENT '上传时间',
    last_update DATETIME DEFAULT CURRENT_TIMESTAMP ON UPDATE CURRENT_TIMESTAMP COMMENT '最后修改时间',
    status INT(11) NOT NULL DEFAULT 0 COMMENT '文件状态(0正常1已删除2禁用)',
    UNIQUE KEY idx_user_file (user_name, file_sha1),
    KEY idx_status (status),
    KEY idx_user_id (user_name)
) ENGINE=InnoDB DEFAULT CHARSET=utf8mb4;
