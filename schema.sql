# Dump of table mall_ad
# ------------------------------------------------------------

CREATE TABLE `mall_ad`
(
    `id`         int(11)      NOT NULL AUTO_INCREMENT,
    `app`        varchar(16)  NOT NULL DEFAULT '',
    `name`       varchar(63)  NOT NULL DEFAULT '' COMMENT '广告标题',
    `link`       varchar(255) NOT NULL DEFAULT '' COMMENT '所广告的商品页面或者活动页面链接地址',
    `url`        varchar(255) NOT NULL COMMENT '广告宣传图片',
    `position`   tinyint(3)   NOT NULL DEFAULT '1' COMMENT '广告位置：1则是首页',
    `content`    varchar(255) NOT NULL DEFAULT '' COMMENT '活动内容',
    `start_time` datetime              DEFAULT NULL COMMENT '广告开始时间',
    `end_time`   datetime              DEFAULT NULL COMMENT '广告结束时间',
    `enabled`    tinyint(1)   NOT NULL DEFAULT '0' COMMENT '是否启动',
    `created_at` datetime              DEFAULT NULL COMMENT '创建时间',
    `updated_at` datetime              DEFAULT NULL COMMENT '更新时间',
    `deleted_at` datetime              DEFAULT NULL COMMENT '删除时间 (软删除)',
    PRIMARY KEY (`id`),
    KEY `enabled` (`enabled`)
) ENGINE = InnoDB
  DEFAULT CHARSET = utf8mb4 COMMENT ='广告表';



# Dump of table mall_address
# ------------------------------------------------------------

CREATE TABLE `mall_address`
(
    `id`             int(11)      NOT NULL AUTO_INCREMENT,
    `app`            varchar(16)  NOT NULL DEFAULT '',
    `name`           varchar(63)  NOT NULL DEFAULT '' COMMENT '收货人名称',
    `user_id`        int(11)      NOT NULL DEFAULT '0' COMMENT '用户表的用户ID',
    `province`       varchar(63)  NOT NULL COMMENT '行政区域表的省ID',
    `city`           varchar(63)  NOT NULL COMMENT '行政区域表的市ID',
    `county`         varchar(63)  NOT NULL COMMENT '行政区域表的区县ID',
    `address_detail` varchar(127) NOT NULL DEFAULT '' COMMENT '详细收货地址',
    `area_code`      char(6)      NOT NULL DEFAULT '' COMMENT '地区编码',
    `postal_code`    char(6)      NOT NULL DEFAULT '' COMMENT '邮政编码',
    `tel`            varchar(20)  NOT NULL DEFAULT '' COMMENT '手机号码',
    `is_default`     tinyint(1)   NOT NULL DEFAULT '0' COMMENT '是否默认地址',
    `created_at`     datetime              DEFAULT NULL COMMENT '创建时间',
    `updated_at`     datetime              DEFAULT NULL COMMENT '更新时间',
    `deleted_at`     datetime              DEFAULT NULL COMMENT '删除时间 (软删除)',
    PRIMARY KEY (`id`),
    KEY `user_id` (`user_id`)
) ENGINE = InnoDB
  DEFAULT CHARSET = utf8mb4 COMMENT ='收货地址表';



# Dump of table mall_brand
# ------------------------------------------------------------

CREATE TABLE `mall_brand`
(
    `id`          int(11)        NOT NULL AUTO_INCREMENT,
    `app`         varchar(16)    NOT NULL DEFAULT '',
    `name`        varchar(255)   NOT NULL DEFAULT '' COMMENT '品牌商名称',
    `desc`        varchar(255)   NOT NULL DEFAULT '' COMMENT '品牌商简介',
    `pic_url`     varchar(255)   NOT NULL DEFAULT '' COMMENT '品牌商页的品牌商图片',
    `sort_order`  tinyint(3)     NOT NULL DEFAULT '50',
    `floor_price` decimal(10, 2) NOT NULL DEFAULT '0.00' COMMENT '品牌商的商品低价，仅用于页面展示',
    `created_at`  datetime                DEFAULT NULL COMMENT '创建时间',
    `updated_at`  datetime                DEFAULT NULL COMMENT '更新时间',
    `deleted_at`  datetime                DEFAULT NULL COMMENT '删除时间 (软删除)',
    PRIMARY KEY (`id`)
) ENGINE = InnoDB
  DEFAULT CHARSET = utf8mb4 COMMENT ='品牌商表';



# Dump of table mall_cart
# ------------------------------------------------------------

CREATE TABLE `mall_cart`
(
    `id`             int(11)        NOT NULL AUTO_INCREMENT,
    `app`            varchar(16)    NOT NULL DEFAULT '',
    `user_id`        int(11)        NOT NULL COMMENT '用户表的用户ID',
    `goods_id`       int(11)        NOT NULL COMMENT '商品表的商品ID',
    `goods_sn`       varchar(63)    NOT NULL DEFAULT '' COMMENT '商品编号',
    `goods_name`     varchar(127)   NOT NULL DEFAULT '' COMMENT '商品名称',
    `product_id`     int(11)        NOT NULL COMMENT '商品货品表的货品ID',
    `price`          decimal(10, 2) NOT NULL DEFAULT '0.00' COMMENT '商品货品的价格',
    `number`         smallint(5)    NOT NULL DEFAULT '0' COMMENT '商品货品的数量',
    `specifications` varchar(1023)  NOT NULL DEFAULT '' COMMENT '商品规格值列表，采用JSON数组格式',
    `checked`        tinyint(1)     NOT NULL DEFAULT '1' COMMENT '购物车中商品是否选择状态',
    `pic_url`        varchar(255)   NOT NULL DEFAULT '' COMMENT '商品图片或者商品货品图片',
    `created_at`     datetime                DEFAULT NULL COMMENT '创建时间',
    `updated_at`     datetime                DEFAULT NULL COMMENT '更新时间',
    `deleted_at`     datetime                DEFAULT NULL COMMENT '删除时间 (软删除)',
    PRIMARY KEY (`id`)
) ENGINE = InnoDB
  DEFAULT CHARSET = utf8mb4 COMMENT ='购物车商品表';



# Dump of table mall_category
# ------------------------------------------------------------

CREATE TABLE `mall_category`
(
    `id`         int(11)       NOT NULL AUTO_INCREMENT,
    `app`        varchar(16)   NOT NULL DEFAULT '',
    `name`       varchar(63)   NOT NULL DEFAULT '' COMMENT '类目名称',
    `keywords`   varchar(1023) NOT NULL DEFAULT '' COMMENT '类目关键字，以JSON数组格式',
    `desc`       varchar(255)  NOT NULL DEFAULT '' COMMENT '类目广告语介绍',
    `pid`        int(11)       NOT NULL DEFAULT '0' COMMENT '父类目ID',
    `icon_url`   varchar(255)  NOT NULL DEFAULT '' COMMENT '类目图标',
    `pic_url`    varchar(255)  NOT NULL DEFAULT '' COMMENT '类目图片',
    `level`      varchar(255)  NOT NULL DEFAULT 'L1',
    `sort_order` tinyint(3)    NOT NULL DEFAULT '50' COMMENT '排序',
    `created_at` datetime               DEFAULT NULL COMMENT '创建时间',
    `updated_at` datetime               DEFAULT NULL COMMENT '更新时间',
    `deleted_at` datetime               DEFAULT NULL COMMENT '删除时间 (软删除)',
    PRIMARY KEY (`id`),
    KEY `parent_id` (`pid`)
) ENGINE = InnoDB
  DEFAULT CHARSET = utf8mb4 COMMENT ='类目表';



# Dump of table mall_collect
# ------------------------------------------------------------

CREATE TABLE `mall_collect`
(
    `id`         int(11)     NOT NULL AUTO_INCREMENT,
    `app`        varchar(16) NOT NULL DEFAULT '',
    `user_id`    int(11)     NOT NULL DEFAULT '0' COMMENT '用户表的用户ID',
    `value_id`   int(11)     NOT NULL DEFAULT '0' COMMENT '如果type=0，则是商品ID；如果type=1，则是专题ID',
    `type`       tinyint(3)  NOT NULL DEFAULT '0' COMMENT '收藏类型，如果type=0，则是商品ID；如果type=1，则是专题ID',
    `created_at` datetime             DEFAULT NULL COMMENT '创建时间',
    `updated_at` datetime             DEFAULT NULL COMMENT '更新时间',
    `deleted_at` datetime             DEFAULT NULL COMMENT '删除时间 (软删除)',
    PRIMARY KEY (`id`),
    KEY `user_id` (`user_id`),
    KEY `goods_id` (`value_id`)
) ENGINE = InnoDB
  DEFAULT CHARSET = utf8mb4 COMMENT ='收藏表';



# Dump of table mall_comment
# ------------------------------------------------------------

CREATE TABLE `mall_comment`
(
    `id`          int(11)       NOT NULL AUTO_INCREMENT,
    `app`         varchar(16)   NOT NULL DEFAULT '',
    `value_id`    int(11)       NOT NULL DEFAULT '0' COMMENT '如果type=0，则是商品评论；如果是type=1，则是专题评论。',
    `type`        tinyint(3)    NOT NULL DEFAULT '0' COMMENT '评论类型，如果type=0，则是商品评论；如果是type=1，则是专题评论；如果type=3，则是订单商品评论。',
    `content`     varchar(1023) NOT NULL COMMENT '评论内容',
    `user_id`     int(11)       NOT NULL DEFAULT '0' COMMENT '用户表的用户ID',
    `has_picture` tinyint(1)    NOT NULL DEFAULT '0' COMMENT '是否含有图片',
    `pic_urls`    varchar(1023) NOT NULL DEFAULT '' COMMENT '图片地址列表，采用JSON数组格式',
    `star`        smallint(6)   NOT NULL DEFAULT '1' COMMENT '评分， 1-5',
    `created_at`  datetime               DEFAULT NULL COMMENT '创建时间',
    `updated_at`  datetime               DEFAULT NULL COMMENT '更新时间',
    `deleted_at`  datetime               DEFAULT NULL COMMENT '删除时间 (软删除)',
    PRIMARY KEY (`id`),
    KEY `id_value` (`value_id`)
) ENGINE = InnoDB
  DEFAULT CHARSET = utf8mb4 COMMENT ='评论表';



# Dump of table mall_coupon
# ------------------------------------------------------------

CREATE TABLE `mall_coupon`
(
    `id`          int(11)        NOT NULL AUTO_INCREMENT,
    `app`         varchar(16)    NOT NULL DEFAULT '',
    `name`        varchar(63)    NOT NULL COMMENT '优惠券名称',
    `desc`        varchar(127)   NOT NULL DEFAULT '' COMMENT '优惠券介绍，通常是显示优惠券使用限制文字',
    `tag`         varchar(63)    NOT NULL DEFAULT '' COMMENT '优惠券标签，例如新人专用',
    `total`       int(11)        NOT NULL DEFAULT '0' COMMENT '优惠券数量，如果是0，则是无限量',
    `discount`    decimal(10, 2) NOT NULL DEFAULT '0.00' COMMENT '优惠金额，',
    `min`         decimal(10, 2) NOT NULL DEFAULT '0.00' COMMENT '最少消费金额才能使用优惠券。',
    `limit`       smallint(6)    NOT NULL DEFAULT '1' COMMENT '用户领券限制数量，如果是0，则是不限制；默认是1，限领一张.',
    `type`        smallint(6)    NOT NULL DEFAULT '0' COMMENT '优惠券赠送类型，如果是0则通用券，用户领取；如果是1，则是注册赠券；如果是2，则是优惠券码兑换；',
    `status`      smallint(6)    NOT NULL DEFAULT '0' COMMENT '优惠券状态，如果是0则是正常可用；如果是1则是过期; 如果是2则是下架。',
    `goods_type`  smallint(6)    NOT NULL DEFAULT '0' COMMENT '商品限制类型，如果0则全商品，如果是1则是类目限制，如果是2则是商品限制。',
    `goods_value` varchar(1023)  NOT NULL DEFAULT '[]' COMMENT '商品限制值，goods_type如果是0则空集合，如果是1则是类目集合，如果是2则是商品集合。',
    `code`        varchar(63)    NOT NULL DEFAULT '' COMMENT '优惠券兑换码',
    `time_type`   smallint(6)    NOT NULL DEFAULT '0' COMMENT '有效时间限制，如果是0，则基于领取时间的有效天数days；如果是1，则start_time和end_time是优惠券有效期；',
    `days`        smallint(6)    NOT NULL DEFAULT '0' COMMENT '基于领取时间的有效天数days。',
    `start_time`  datetime                DEFAULT NULL COMMENT '使用券开始时间',
    `end_time`    datetime                DEFAULT NULL COMMENT '使用券截至时间',
    `created_at`  datetime                DEFAULT NULL COMMENT '创建时间',
    `updated_at`  datetime                DEFAULT NULL COMMENT '更新时间',
    `deleted_at`  datetime                DEFAULT NULL COMMENT '删除时间 (软删除)',
    PRIMARY KEY (`id`)
) ENGINE = InnoDB
  DEFAULT CHARSET = utf8mb4 COMMENT ='优惠券信息及规则表';



# Dump of table mall_coupon_user
# ------------------------------------------------------------

CREATE TABLE `mall_coupon_user`
(
    `id`         int(11)     NOT NULL AUTO_INCREMENT,
    `app`        varchar(16) NOT NULL DEFAULT '',
    `user_id`    int(11)     NOT NULL COMMENT '用户ID',
    `coupon_id`  int(11)     NOT NULL COMMENT '优惠券ID',
    `status`     smallint(6) NOT NULL DEFAULT '0' COMMENT '使用状态, 如果是0则未使用；如果是1则已使用；如果是2则已过期；如果是3则已经下架；',
    `used_time`  datetime             DEFAULT NULL COMMENT '使用时间',
    `start_time` datetime    NOT NULL COMMENT '有效期开始时间',
    `end_time`   datetime    NOT NULL COMMENT '有效期截至时间',
    `order_id`   int(11)     NOT NULL COMMENT '订单ID',
    `created_at` datetime             DEFAULT NULL COMMENT '创建时间',
    `updated_at` datetime             DEFAULT NULL COMMENT '更新时间',
    `deleted_at` datetime             DEFAULT NULL COMMENT '删除时间 (软删除)',
    PRIMARY KEY (`id`)
) ENGINE = InnoDB
  DEFAULT CHARSET = utf8mb4 COMMENT ='优惠券用户使用表';



# Dump of table mall_feedback
# ------------------------------------------------------------

CREATE TABLE `mall_feedback`
(
    `id`          int(11)       NOT NULL AUTO_INCREMENT,
    `app`         varchar(16)   NOT NULL DEFAULT '',
    `user_id`     int(11)       NOT NULL DEFAULT '0' COMMENT '用户表的用户ID',
    `username`    varchar(63)   NOT NULL DEFAULT '' COMMENT '用户名称',
    `mobile`      varchar(20)   NOT NULL DEFAULT '' COMMENT '手机号',
    `feed_type`   varchar(63)   NOT NULL DEFAULT '' COMMENT '反馈类型',
    `content`     varchar(1023) NOT NULL COMMENT '反馈内容',
    `status`      int(3)        NOT NULL DEFAULT '0' COMMENT '状态',
    `has_picture` tinyint(1)    NOT NULL DEFAULT '0' COMMENT '是否含有图片',
    `pic_urls`    varchar(1023) NOT NULL DEFAULT '' COMMENT '图片地址列表，采用JSON数组格式',
    `created_at`  datetime               DEFAULT NULL COMMENT '创建时间',
    `updated_at`  datetime               DEFAULT NULL COMMENT '更新时间',
    `deleted_at`  datetime               DEFAULT NULL COMMENT '删除时间 (软删除)',
    PRIMARY KEY (`id`),
    KEY `id_value` (`status`)
) ENGINE = InnoDB
  DEFAULT CHARSET = utf8mb4 COMMENT ='意见反馈表';



# Dump of table mall_footprint
# ------------------------------------------------------------

CREATE TABLE `mall_footprint`
(
    `id`         int(11)     NOT NULL AUTO_INCREMENT,
    `app`        varchar(16) NOT NULL DEFAULT '',
    `user_id`    int(11)     NOT NULL DEFAULT '0' COMMENT '用户表的用户ID',
    `goods_id`   int(11)     NOT NULL DEFAULT '0' COMMENT '浏览商品ID',
    `created_at` datetime             DEFAULT NULL COMMENT '创建时间',
    `updated_at` datetime             DEFAULT NULL COMMENT '更新时间',
    `deleted_at` datetime             DEFAULT NULL COMMENT '删除时间 (软删除)',
    PRIMARY KEY (`id`)
) ENGINE = InnoDB
  DEFAULT CHARSET = utf8mb4 COMMENT ='用户浏览足迹表';



# Dump of table mall_goods
# ------------------------------------------------------------

CREATE TABLE `mall_goods`
(
    `id`            int(11)        NOT NULL AUTO_INCREMENT,
    `app`           varchar(16)    NOT NULL DEFAULT '',
    `goods_sn`      varchar(63)    NOT NULL DEFAULT '' COMMENT '商品编号',
    `name`          varchar(127)   NOT NULL DEFAULT '' COMMENT '商品名称',
    `category_id`   int(11)        NOT NULL DEFAULT '0' COMMENT '商品所属类目ID',
    `brand_id`      int(11)        NOT NULL DEFAULT '0',
    `gallery`       varchar(1023)  NOT NULL DEFAULT '' COMMENT '商品宣传图片列表，采用JSON数组格式',
    `keywords`      varchar(255)   NOT NULL DEFAULT '' COMMENT '商品关键字，采用逗号间隔',
    `brief`         varchar(255)   NOT NULL DEFAULT '' COMMENT '商品简介',
    `is_on_sale`    tinyint(1)     NOT NULL DEFAULT '1' COMMENT '是否上架',
    `sort_order`    smallint(4)    NOT NULL DEFAULT '100',
    `pic_url`       varchar(255)   NOT NULL DEFAULT '' COMMENT '商品页面商品图片',
    `share_url`     varchar(255)   NOT NULL DEFAULT '' COMMENT '商品分享朋友圈图片',
    `is_new`        tinyint(1)     NOT NULL DEFAULT '0' COMMENT '是否新品首发，如果设置则可以在新品首发页面展示',
    `is_hot`        tinyint(1)     NOT NULL DEFAULT '0' COMMENT '是否人气推荐，如果设置则可以在人气推荐页面展示',
    `unit`          varchar(31)    NOT NULL DEFAULT '’件‘' COMMENT '商品单位，例如件、盒',
    `counter_price` decimal(10, 2) NOT NULL DEFAULT '0.00' COMMENT '专柜价格',
    `retail_price`  decimal(10, 2) NOT NULL DEFAULT '100000.00' COMMENT '零售价格',
    `detail`        text           NOT NULL COMMENT '商品详细介绍，是富文本格式',
    `created_at`    datetime                DEFAULT NULL COMMENT '创建时间',
    `updated_at`    datetime                DEFAULT NULL COMMENT '更新时间',
    `deleted_at`    datetime                DEFAULT NULL COMMENT '删除时间 (软删除)',
    PRIMARY KEY (`id`),
    KEY `goods_sn` (`goods_sn`),
    KEY `cat_id` (`category_id`),
    KEY `brand_id` (`brand_id`),
    KEY `sort_order` (`sort_order`)
) ENGINE = InnoDB
  DEFAULT CHARSET = utf8mb4 COMMENT ='商品基本信息表';



# Dump of table mall_goods_attribute
# ------------------------------------------------------------

CREATE TABLE `mall_goods_attribute`
(
    `id`         int(11)      NOT NULL AUTO_INCREMENT,
    `app`        varchar(16)  NOT NULL DEFAULT '',
    `goods_id`   int(11)      NOT NULL DEFAULT '0' COMMENT '商品表的商品ID',
    `attribute`  varchar(255) NOT NULL COMMENT '商品参数名称',
    `value`      varchar(255) NOT NULL COMMENT '商品参数值',
    `created_at` datetime              DEFAULT NULL COMMENT '创建时间',
    `updated_at` datetime              DEFAULT NULL COMMENT '更新时间',
    `deleted_at` datetime              DEFAULT NULL COMMENT '删除时间 (软删除)',
    PRIMARY KEY (`id`),
    KEY `goods_id` (`goods_id`)
) ENGINE = InnoDB
  DEFAULT CHARSET = utf8mb4 COMMENT ='商品参数表';



# Dump of table mall_goods_product
# ------------------------------------------------------------

CREATE TABLE `mall_goods_product`
(
    `id`             int(11)        NOT NULL AUTO_INCREMENT,
    `app`            varchar(16)    NOT NULL DEFAULT '',
    `goods_id`       int(11)        NOT NULL DEFAULT '0' COMMENT '商品表的商品ID',
    `specifications` varchar(1023)  NOT NULL COMMENT '商品规格值列表，采用JSON数组格式',
    `price`          decimal(10, 2) NOT NULL DEFAULT '0.00' COMMENT '商品货品价格',
    `number`         int(11)        NOT NULL DEFAULT '0' COMMENT '商品货品数量',
    `url`            varchar(125)   NOT NULL DEFAULT '' COMMENT '商品货品图片',
    `created_at`     datetime                DEFAULT NULL COMMENT '创建时间',
    `updated_at`     datetime                DEFAULT NULL COMMENT '更新时间',
    `deleted_at`     datetime                DEFAULT NULL COMMENT '删除时间 (软删除)',
    PRIMARY KEY (`id`)
) ENGINE = InnoDB
  DEFAULT CHARSET = utf8mb4 COMMENT ='商品货品表';



# Dump of table mall_goods_specification
# ------------------------------------------------------------

CREATE TABLE `mall_goods_specification`
(
    `id`            int(11)      NOT NULL AUTO_INCREMENT,
    `app`           varchar(16)  NOT NULL DEFAULT '',
    `goods_id`      int(11)      NOT NULL DEFAULT '0' COMMENT '商品表的商品ID',
    `specification` varchar(255) NOT NULL DEFAULT '' COMMENT '商品规格名称',
    `value`         varchar(255) NOT NULL DEFAULT '' COMMENT '商品规格值',
    `pic_url`       varchar(255) NOT NULL DEFAULT '' COMMENT '商品规格图片',
    `created_at`    datetime              DEFAULT NULL COMMENT '创建时间',
    `updated_at`    datetime              DEFAULT NULL COMMENT '更新时间',
    `deleted_at`    datetime              DEFAULT NULL COMMENT '删除时间 (软删除)',
    PRIMARY KEY (`id`),
    KEY `goods_id` (`goods_id`)
) ENGINE = InnoDB
  DEFAULT CHARSET = utf8mb4 COMMENT ='商品规格表';



# Dump of table mall_groupon
# ------------------------------------------------------------

CREATE TABLE `mall_groupon`
(
    `id`                int(11)      NOT NULL AUTO_INCREMENT,
    `app`               varchar(16)  NOT NULL DEFAULT '',
    `order_id`          int(11)      NOT NULL COMMENT '关联的订单ID',
    `groupon_id`        int(11)      NOT NULL DEFAULT '0' COMMENT '如果是开团用户，则groupon_id是0；如果是参团用户，则groupon_id是团购活动ID',
    `rules_id`          int(11)      NOT NULL COMMENT '团购规则ID，关联mall_groupon_rules表ID字段',
    `user_id`           int(11)      NOT NULL COMMENT '用户ID',
    `share_url`         varchar(255) NOT NULL DEFAULT '' COMMENT '团购分享图片地址',
    `creator_user_id`   int(11)      NOT NULL COMMENT '开团用户ID',
    `creator_user_time` datetime              DEFAULT NULL COMMENT '开团时间',
    `status`            smallint(6)  NOT NULL DEFAULT '0' COMMENT '团购活动状态，开团未支付则0，开团中则1，开团失败则2',
    `created_at`        datetime     NOT NULL COMMENT '创建时间',
    `updated_at`        datetime              DEFAULT NULL COMMENT '更新时间',
    `deleted_at`        datetime              DEFAULT NULL COMMENT '删除时间 (软删除)',
    PRIMARY KEY (`id`) USING BTREE
) ENGINE = InnoDB
  DEFAULT CHARSET = utf8mb4
  ROW_FORMAT = DYNAMIC COMMENT ='团购活动表';



# Dump of table mall_groupon_rules
# ------------------------------------------------------------

CREATE TABLE `mall_groupon_rules`
(
    `id`              int(11)        NOT NULL AUTO_INCREMENT,
    `app`             varchar(16)    NOT NULL DEFAULT '',
    `goods_id`        int(11)        NOT NULL COMMENT '商品表的商品ID',
    `goods_name`      varchar(127)   NOT NULL COMMENT '商品名称',
    `pic_url`         varchar(255)   NOT NULL DEFAULT '' COMMENT '商品图片或者商品货品图片',
    `discount`        decimal(63, 0) NOT NULL COMMENT '优惠金额',
    `discount_member` int(11)        NOT NULL COMMENT '达到优惠条件的人数',
    `expire_time`     datetime                DEFAULT NULL COMMENT '团购过期时间',
    `status`          smallint(6)    NOT NULL DEFAULT '0' COMMENT '团购规则状态，正常上线则0，到期自动下线则1，管理手动下线则2',
    `created_at`      datetime                DEFAULT NULL COMMENT '创建时间',
    `updated_at`      datetime                DEFAULT NULL COMMENT '更新时间',
    `deleted_at`      datetime                DEFAULT NULL COMMENT '删除时间 (软删除)',
    PRIMARY KEY (`id`) USING BTREE
) ENGINE = InnoDB
  DEFAULT CHARSET = utf8mb4
  ROW_FORMAT = DYNAMIC COMMENT ='团购规则表';



# Dump of table mall_issue
# ------------------------------------------------------------

CREATE TABLE `mall_issue`
(
    `id`         int(11)      NOT NULL AUTO_INCREMENT,
    `app`        varchar(16)  NOT NULL DEFAULT '',
    `question`   varchar(255) NOT NULL DEFAULT '' COMMENT '问题标题',
    `answer`     varchar(255) NOT NULL DEFAULT '' COMMENT '问题答案',
    `created_at` datetime              DEFAULT NULL COMMENT '创建时间',
    `updated_at` datetime              DEFAULT NULL COMMENT '更新时间',
    `deleted_at` datetime              DEFAULT NULL COMMENT '删除时间 (软删除)',
    PRIMARY KEY (`id`)
) ENGINE = InnoDB
  DEFAULT CHARSET = utf8mb4 COMMENT ='常见问题表';



# Dump of table mall_keyword
# ------------------------------------------------------------

CREATE TABLE `mall_keyword`
(
    `id`         int(11)      NOT NULL AUTO_INCREMENT,
    `app`        varchar(16)  NOT NULL DEFAULT '',
    `keyword`    varchar(127) NOT NULL DEFAULT '' COMMENT '关键字',
    `url`        varchar(255) NOT NULL DEFAULT '' COMMENT '关键字的跳转链接',
    `is_hot`     tinyint(1)   NOT NULL DEFAULT '0' COMMENT '是否是热门关键字',
    `is_default` tinyint(1)   NOT NULL DEFAULT '0' COMMENT '是否是默认关键字',
    `sort_order` int(11)      NOT NULL DEFAULT '100' COMMENT '排序',
    `created_at` datetime              DEFAULT NULL COMMENT '创建时间',
    `updated_at` datetime              DEFAULT NULL COMMENT '更新时间',
    `deleted_at` datetime              DEFAULT NULL COMMENT '删除时间 (软删除)',
    PRIMARY KEY (`id`)
) ENGINE = InnoDB
  DEFAULT CHARSET = utf8mb4 COMMENT ='关键字表';



# Dump of table mall_order
# ------------------------------------------------------------

CREATE TABLE `mall_order`
(
    `id`             int(11)        NOT NULL AUTO_INCREMENT,
    `app`            varchar(16)    NOT NULL DEFAULT '',
    `user_id`        int(11)        NOT NULL COMMENT '用户表的用户ID',
    `order_sn`       varchar(63)    NOT NULL COMMENT '订单编号',
    `order_status`   smallint(6)    NOT NULL COMMENT '订单状态',
    `consignee`      varchar(63)    NOT NULL COMMENT '收货人名称',
    `mobile`         varchar(63)    NOT NULL COMMENT '收货人手机号',
    `address`        varchar(127)   NOT NULL COMMENT '收货具体地址',
    `message`        varchar(512)   NOT NULL DEFAULT '' COMMENT '用户订单留言',
    `goods_price`    decimal(10, 2) NOT NULL COMMENT '商品总费用',
    `freight_price`  decimal(10, 2) NOT NULL COMMENT '配送费用',
    `coupon_price`   decimal(10, 2) NOT NULL COMMENT '优惠券减免',
    `integral_price` decimal(10, 2) NOT NULL COMMENT '用户积分减免',
    `groupon_price`  decimal(10, 2) NOT NULL COMMENT '团购优惠价减免',
    `order_price`    decimal(10, 2) NOT NULL COMMENT '订单费用， = goods_price + freight_price - coupon_price',
    `actual_price`   decimal(10, 2) NOT NULL COMMENT '实付费用， = order_price - integral_price',
    `pay_id`         varchar(63)    NOT NULL DEFAULT '' COMMENT '微信付款编号',
    `pay_time`       datetime                DEFAULT NULL COMMENT '微信付款时间',
    `ship_sn`        varchar(63)    NOT NULL DEFAULT '' COMMENT '发货编号',
    `ship_channel`   varchar(63)    NOT NULL DEFAULT '' COMMENT '发货快递公司',
    `ship_time`      datetime                DEFAULT NULL COMMENT '发货开始时间',
    `refund_amount`  decimal(10, 2) NOT NULL COMMENT '实际退款金额，（有可能退款金额小于实际支付金额）',
    `refund_type`    varchar(63)    NOT NULL DEFAULT '' COMMENT '退款方式',
    `refund_content` varchar(127)   NOT NULL DEFAULT '' COMMENT '退款备注',
    `refund_time`    datetime                DEFAULT NULL COMMENT '退款时间',
    `confirm_time`   datetime                DEFAULT NULL COMMENT '用户确认收货时间',
    `comments`       smallint(6)    NOT NULL DEFAULT '0' COMMENT '待评价订单商品数量',
    `end_time`       datetime                DEFAULT NULL COMMENT '订单关闭时间',
    `created_at`     datetime                DEFAULT NULL COMMENT '创建时间',
    `updated_at`     datetime                DEFAULT NULL COMMENT '更新时间',
    `deleted_at`     datetime                DEFAULT NULL COMMENT '删除时间 (软删除)',
    PRIMARY KEY (`id`),
    UNIQUE KEY `order_sn` (`order_sn`)
) ENGINE = InnoDB
  DEFAULT CHARSET = utf8mb4 COMMENT ='订单表';



# Dump of table mall_order_goods
# ------------------------------------------------------------

CREATE TABLE `mall_order_goods`
(
    `id`             int(11)        NOT NULL AUTO_INCREMENT,
    `app`            varchar(16)    NOT NULL DEFAULT '',
    `order_id`       int(11)        NOT NULL DEFAULT '0' COMMENT '订单表的订单ID',
    `goods_id`       int(11)        NOT NULL DEFAULT '0' COMMENT '商品表的商品ID',
    `goods_name`     varchar(127)   NOT NULL DEFAULT '' COMMENT '商品名称',
    `goods_sn`       varchar(63)    NOT NULL DEFAULT '' COMMENT '商品编号',
    `product_id`     int(11)        NOT NULL DEFAULT '0' COMMENT '商品货品表的货品ID',
    `number`         smallint(5)    NOT NULL DEFAULT '0' COMMENT '商品货品的购买数量',
    `price`          decimal(10, 2) NOT NULL DEFAULT '0.00' COMMENT '商品货品的售价',
    `specifications` varchar(1023)  NOT NULL COMMENT '商品货品的规格列表',
    `pic_url`        varchar(255)   NOT NULL DEFAULT '' COMMENT '商品货品图片或者商品图片',
    `comment`        int(11)        NOT NULL DEFAULT '0' COMMENT '订单商品评论，如果是-1，则超期不能评价；如果是0，则可以评价；如果其他值，则是comment表里面的评论ID。',
    `created_at`     datetime                DEFAULT NULL COMMENT '创建时间',
    `updated_at`     datetime                DEFAULT NULL COMMENT '更新时间',
    `deleted_at`     datetime                DEFAULT NULL COMMENT '删除时间 (软删除)',
    PRIMARY KEY (`id`),
    KEY `order_id` (`order_id`),
    KEY `goods_id` (`goods_id`)
) ENGINE = InnoDB
  DEFAULT CHARSET = utf8mb4 COMMENT ='订单商品表';



# Dump of table mall_search_history
# ------------------------------------------------------------

CREATE TABLE `mall_search_history`
(
    `id`         int(11)     NOT NULL AUTO_INCREMENT,
    `app`        varchar(16) NOT NULL DEFAULT '',
    `user_id`    int(11)     NOT NULL COMMENT '用户表的用户ID',
    `keyword`    varchar(63) NOT NULL COMMENT '搜索关键字',
    `from`       varchar(63) NOT NULL DEFAULT '' COMMENT '搜索来源，如pc、wx、app',
    `created_at` datetime             DEFAULT NULL COMMENT '创建时间',
    `updated_at` datetime             DEFAULT NULL COMMENT '更新时间',
    `deleted_at` datetime             DEFAULT NULL COMMENT '删除时间 (软删除)',
    PRIMARY KEY (`id`)
) ENGINE = InnoDB
  DEFAULT CHARSET = utf8mb4 COMMENT ='搜索历史表';



# Dump of table mall_system
# ------------------------------------------------------------

CREATE TABLE `mall_system`
(
    `id`         int(11)      NOT NULL AUTO_INCREMENT,
    `app`        varchar(16)  NOT NULL DEFAULT '',
    `key_name`   varchar(255) NOT NULL COMMENT '系统配置名',
    `key_value`  varchar(255) NOT NULL COMMENT '系统配置值',
    `created_at` datetime              DEFAULT NULL COMMENT '创建时间',
    `updated_at` datetime              DEFAULT NULL COMMENT '更新时间',
    `deleted_at` datetime              DEFAULT NULL COMMENT '删除时间 (软删除)',
    PRIMARY KEY (`id`) USING BTREE
) ENGINE = InnoDB
  DEFAULT CHARSET = utf8mb4
  ROW_FORMAT = DYNAMIC COMMENT ='系统配置表';



# Dump of table mall_topic
# ------------------------------------------------------------

CREATE TABLE `mall_topic`
(
    `id`         int(11)        NOT NULL AUTO_INCREMENT,
    `app`        varchar(16)    NOT NULL DEFAULT '',
    `title`      varchar(255)   NOT NULL DEFAULT '''' COMMENT '专题标题',
    `subtitle`   varchar(255)   NOT NULL DEFAULT '''' COMMENT '专题子标题',
    `content`    text           NOT NULL COMMENT '专题内容，富文本格式',
    `price`      decimal(10, 2) NOT NULL DEFAULT '0.00' COMMENT '专题相关商品最低价',
    `read_count` varchar(255)   NOT NULL DEFAULT '1k' COMMENT '专题阅读量',
    `pic_url`    varchar(255)   NOT NULL DEFAULT '' COMMENT '专题图片',
    `sort_order` int(11)        NOT NULL DEFAULT '100' COMMENT '排序',
    `goods`      varchar(1023)  NOT NULL DEFAULT '' COMMENT '专题相关商品，采用JSON数组格式',
    `created_at` datetime                DEFAULT NULL COMMENT '创建时间',
    `updated_at` datetime                DEFAULT NULL COMMENT '更新时间',
    `deleted_at` datetime                DEFAULT NULL COMMENT '删除时间 (软删除)',
    PRIMARY KEY (`id`),
    KEY `topic_id` (`id`)
) ENGINE = InnoDB
  DEFAULT CHARSET = utf8mb4 COMMENT ='专题表';


# Dump of table mall_user
# ------------------------------------------------------------

CREATE TABLE `mall_user`
(
    `id`              int(11)      NOT NULL AUTO_INCREMENT,
    `app`             varchar(16)  NOT NULL DEFAULT '',
    `username`        varchar(63)  NOT NULL COMMENT '用户名称',
    `password`        varchar(63)  NOT NULL DEFAULT '' COMMENT '用户密码',
    `gender`          tinyint(3)   NOT NULL DEFAULT '0' COMMENT '性别：0 未知， 1男， 1 女',
    `birthday`        date                  DEFAULT NULL COMMENT '生日',
    `last_login_time` datetime              DEFAULT NULL COMMENT '最近一次登录时间',
    `last_login_ip`   varchar(63)  NOT NULL DEFAULT '' COMMENT '最近一次登录IP地址',
    `user_level`      tinyint(3)   NOT NULL DEFAULT '0' COMMENT '0 普通用户，1 VIP用户，2 高级VIP用户',
    `nickname`        varchar(63)  NOT NULL DEFAULT '' COMMENT '用户昵称或网络名称',
    `mobile`          varchar(20)  NOT NULL DEFAULT '' COMMENT '用户手机号码',
    `avatar`          varchar(255) NOT NULL DEFAULT '' COMMENT '用户头像图片',
    `platform`        varchar(20)  NOT NULL DEFAULT '' COMMENT '平台',
    `openid`          varchar(63)  NOT NULL DEFAULT '' COMMENT '登录openid',
    `session_key`     varchar(100) NOT NULL DEFAULT '' COMMENT '登录会话KEY',
    `status`          tinyint(3)   NOT NULL DEFAULT '0' COMMENT '0 可用, 1 禁用, 2 注销',
    `created_at`      datetime              DEFAULT NULL COMMENT '创建时间',
    `updated_at`      datetime              DEFAULT NULL COMMENT '更新时间',
    `deleted_at`      datetime              DEFAULT NULL COMMENT '删除时间 (软删除)',
    PRIMARY KEY (`id`),
    UNIQUE KEY `user_name` (`username`)
) ENGINE = InnoDB
  DEFAULT CHARSET = utf8mb4 COMMENT ='用户表';
