create table if not exists attribute
(
    id         bigint auto_increment
        primary key,
    spu_id     bigint                             not null comment 'sup的id',
    name       varchar(64)                        not null comment '属性名称',
    value      varchar(255)                       not null comment '属性值',
    created_at datetime default CURRENT_TIMESTAMP not null comment '创建时间',
    updated_at datetime default CURRENT_TIMESTAMP not null on update CURRENT_TIMESTAMP,
    deleted_at datetime                           null
);

create table if not exists brand
(
    id         bigint auto_increment
        primary key,
    name       varchar(264)                       not null comment '品牌名称',
    icon_url   varchar(255)                       not null,
    created_at datetime default CURRENT_TIMESTAMP not null,
    updated_at datetime default CURRENT_TIMESTAMP not null on update CURRENT_TIMESTAMP,
    deleted_at datetime                           null
);

create table if not exists category
(
    id         bigint auto_increment
        primary key,
    name       varchar(64)                        not null,
    parent_id  bigint                             not null comment '父品牌id',
    level      int      default 0                 not null comment '层级：0,1,2',
    created_at datetime default CURRENT_TIMESTAMP not null,
    updated_at datetime default CURRENT_TIMESTAMP not null on update CURRENT_TIMESTAMP,
    deleted_at datetime                           null
);

create table if not exists goods_sku
(
    id          bigint auto_increment
        primary key,
    spu_id      bigint                                not null comment 'spu的id',
    sku_price   varchar(20) default '00.00'           not null,
    stock_num   int         default 0                 not null comment '库存',
    spec_values longtext                              not null comment '商品规格信息，json存储',
    enabled     tinyint(1)                            not null comment '是否上架',
    created_at  datetime    default CURRENT_TIMESTAMP not null,
    updated_at  datetime    default CURRENT_TIMESTAMP not null on update CURRENT_TIMESTAMP,
    deleted_at  datetime                              null
);

create table if not exists goods_spu
(
    id             bigint auto_increment
        primary key,
    spu_name       varchar(64)                           not null comment 'spu名称',
    spu_desc       varchar(255)                          not null comment 'spu描述',
    seller_id      bigint                                null,
    shop_id        bigint                                null,
    category_id    bigint                                not null comment '分类id',
    category_name  varchar(64)                           not null comment '分类名称',
    brand_id       bigint                                not null comment '品牌id',
    brand_name     varchar(64)                           not null comment '品牌名称',
    brand_icon_url text                                  null,
    spu_price      varchar(20) default '00.00'           not null comment '价格, 展示使用，不用于交易',
    enabled        tinyint(1)  default 1                 not null comment '是否上架',
    spu_img_url    varchar(255)                          not null comment 'spu图片url',
    attribute_ids  varchar(255)                          not null comment 'spu拥有的属性id列表 ,隔开 不额外用一张表存关系',
    spec_ids       varchar(255)                          not null comment 'spu拥有的规格项id ,隔开 不额外用一张表存关系',
    created_at     datetime    default CURRENT_TIMESTAMP not null,
    updated_at     datetime    default CURRENT_TIMESTAMP not null on update CURRENT_TIMESTAMP,
    deleted_at     datetime                              null
);

create table if not exists keyword
(
    id           bigint auto_increment
        primary key,
    value        varchar(64)                        not null comment '搜索关键字',
    search_times bigint   default 0                 not null comment '搜索次数',
    created_at   datetime default CURRENT_TIMESTAMP not null,
    updated_at   datetime default CURRENT_TIMESTAMP not null on update CURRENT_TIMESTAMP,
    deleted_at   datetime                           null
)
    comment '搜索关键词';

create table if not exists spec
(
    id         bigint auto_increment
        primary key,
    name       varchar(64)                        not null comment '规格名称',
    spu_id     bigint                             not null comment 'spu的id',
    created_at datetime default CURRENT_TIMESTAMP not null,
    updated_at datetime default CURRENT_TIMESTAMP not null on update CURRENT_TIMESTAMP,
    deleted_at datetime                           null
);

create table if not exists spec_value
(
    id         bigint auto_increment
        primary key,
    spec_id    bigint                             not null comment '规格项id',
    spu_id     bigint                             not null comment 'spu的id',
    sku_id     bigint                             not null comment 'sku的id',
    name       varchar(64)                        not null comment '规格项名称(spec的name，冗余存储)',
    value      varchar(64)                        not null comment '规格的值',
    created_at datetime default CURRENT_TIMESTAMP not null,
    updated_at datetime default CURRENT_TIMESTAMP not null on update CURRENT_TIMESTAMP,
    deleted_at datetime                           null
)
    comment '规格值';

create table if not exists stock_op_record
(
    id         bigint auto_increment
        primary key,
    trade_no   varchar(255)                       not null comment '交易号',
    op_type    varchar(20)                        null comment '操作类型',
    created_at datetime default CURRENT_TIMESTAMP not null,
    updated_at datetime default CURRENT_TIMESTAMP not null on update CURRENT_TIMESTAMP,
    deleted_at datetime                           null,
    constraint stock_record_pk_trade_op
        unique (trade_no, op_type)
);

create table if not exists trade_info
(
    id         bigint auto_increment
        primary key,
    trade_no   varchar(255)                       null comment '交易号',
    page_url   varchar(255)                       not null comment '支付链接',
    created_at datetime default CURRENT_TIMESTAMP not null,
    updated_at datetime default CURRENT_TIMESTAMP not null on update CURRENT_TIMESTAMP,
    deleted_at datetime                           null
);


create table if not exists `order`
(
    order_no     varchar(255)                       not null
        primary key,
    trade_no     varchar(255)                       not null,
    order_amount varchar(20)                        not null comment '订单金额',
    order_status varchar(64)                        not null comment '订单状态',
    seller_id    bigint                             not null comment '卖家id',
    buyer_id     bigint                             not null comment '买家id',
    pay_type     varchar(20)                        not null comment '支付方式',
    created_at   datetime default CURRENT_TIMESTAMP not null,
    updated_at   datetime default CURRENT_TIMESTAMP not null on update CURRENT_TIMESTAMP,
    deleted_at   datetime                           null,
    shop_id      bigint   default 0                 not null
);

create table if not exists order_item
(
    id            bigint auto_increment
        primary key,
    trade_no      varchar(255)                         not null comment '交易号',
    order_no      varchar(255)                         not null comment '订单号',
    spu_id        bigint                               not null comment 'spu的id',
    sku_id        bigint                               not null comment 'sku的id',
    shop_id       bigint                               not null comment '店铺id',
    seller_id     bigint                               not null comment '卖家id',
    buyer_id      bigint                               not null,
    spu_name      varchar(255)                         not null comment 'spu名称',
    category_id   bigint                               not null comment '商品分类id',
    category_name varchar(255)                         not null comment '商品分类',
    brand_id      bigint                               not null comment '品牌id',
    brand_name    varchar(255)                         null comment '品牌名称',
    item_img_url  varchar(255)                         not null comment '订单项图片url',
    spec_values   varchar(255)                         not null comment '规格值',
    sku_amount    varchar(20)                          not null comment 'sku的金额',
    is_delivered  tinyint(1) default 0                 not null comment '是否发货',
    is_confirmed  tinyint(1) default 0                 not null comment '是否签收',
    created_at    datetime   default CURRENT_TIMESTAMP not null,
    updated_at    datetime   default CURRENT_TIMESTAMP not null on update CURRENT_TIMESTAMP,
    deleted_at    datetime                             null
);

create table if not exists trade
(
    trade_no     varchar(255)                       not null comment '交易号'
        primary key,
    trade_amount varchar(255)                       not null comment '交易金额',
    buyer_id     bigint   default 0                 not null comment '买家id',
    pay_type     varchar(20)                        not null comment '支付方式',
    trade_status varchar(20)                        not null comment '交易状态',
    created_at   datetime default CURRENT_TIMESTAMP not null,
    updated_at   datetime default CURRENT_TIMESTAMP not null on update CURRENT_TIMESTAMP,
    deleted_at   datetime                           null
);


create table if not exists pay_flow
(
    pay_no         varchar(255)                       not null
        primary key,
    trade_no       varchar(255)                       not null comment '电商交易号',
    order_no       varchar(255)                       not null comment '订单号',
    third_trade_no varchar(255)                       null comment '第三方支付机构交易号',
    pay_state      int                                not null comment '支付状态: 1 - 支付中,  2 - 已支付，3 - 完成结算，	4-支付取消， 5-退款，  6-部分退款， 7-退款失败',
    total_amount   varchar(30)                        not null comment '支付金额',
    created_at     datetime default CURRENT_TIMESTAMP not null,
    updated_at     datetime default CURRENT_TIMESTAMP not null on update CURRENT_TIMESTAMP,
    deleted_at     datetime                           null,
    constraint t_pay_flow__index_out_trade_no
        unique (trade_no)
);

create table if not exists pay_info
(
    id           bigint       null,
    trade_no     varchar(255) null,
    pay_page_url text         null,
    created_at   datetime     null,
    updated_at   datetime     null,
    deleted_at   datetime     null
);

create table if not exists shop
(
    id           bigint unsigned auto_increment comment '店铺ID'
        primary key,
    user_id      bigint                             not null,
    shop_name    varchar(128)                       not null comment '店铺名称',
    shop_desc    text                               null comment '店铺描述',
    shop_img_url varchar(255)                       null comment '店铺图片URL',
    created_at   datetime default CURRENT_TIMESTAMP null comment '创建时间',
    updated_at   datetime default CURRENT_TIMESTAMP null on update CURRENT_TIMESTAMP comment '更新时间',
    deleted_at   datetime                           null comment '删除时间',
    constraint idx_shop_name
        unique (shop_name),
    constraint shop_pk_user_id
        unique (user_id)
)
    comment '店铺表' collate = utf8mb4_unicode_ci;

create table if not exists user
(
    id            bigint unsigned auto_increment comment '用户ID'
        primary key,
    username      varchar(64)                        not null comment '用户名',
    avatar_url    varchar(255)                       null comment '头像URL',
    salt          varchar(255)                       not null comment '密码盐值',
    password      varchar(255)                       not null comment '密码',
    register_type varchar(32)                        not null comment '注册类型',
    shop_id       bigint unsigned                    null comment '关联店铺ID',
    created_at    datetime default CURRENT_TIMESTAMP null comment '创建时间',
    updated_at    datetime default CURRENT_TIMESTAMP null on update CURRENT_TIMESTAMP comment '更新时间',
    deleted_at    datetime                           null comment '删除时间',
    constraint idx_username
        unique (username)
)
    comment '用户表' collate = utf8mb4_unicode_ci;

create index idx_shop_id
    on user (shop_id);

