package model

import (
	"time"
)

type Model struct {
	ID        int         `json:"id" db:"id" faker:"-"`
	CreatedAt time.Time   `json:"created_at" db:"created_at"`
	UpdatedAt time.Time   `json:"updated_at" db:"updated_at"`
	DeletedAt interface{} `json:"-" db:"deleted_at" faker:"-"`
}

//
// 订单流程：下单成功－》支付订单－》发货－》收货
// 订单状态：
// 101 订单生成，未支付；102，下单未支付用户取消；103，下单未支付超期系统自动取消
// 201 支付完成，商家未发货；202，订单生产，已付款未发货，用户申请退款；203，管理员执行退款操作，确认退款成功；
// 301 商家发货，用户未确认；
// 401 用户确认收货，订单结束； 402 用户没有确认收货，但是快递反馈已收货后，超过一定时间，系统自动确认收货，订单结束。
//
// 当101用户未付款时，此时用户可以进行的操作是取消或者付款
// 当201支付完成而商家未发货时，此时用户可以退款
// 当301商家已发货时，此时用户可以有确认收货
// 当401用户确认收货以后，此时用户可以进行的操作是退货、删除、去评价或者再次购买
// 当402系统自动确认收货以后，此时用户可以删除、去评价、或者再次购买
const (
	ORDER_STATUS_CREATE          = 101
	ORDER_STATUS_PAY             = 201
	ORDER_STATUS_SHIP            = 301
	ORDER_STATUS_CONFIRM         = 401
	ORDER_STATUS_CANCEL          = 102
	ORDER_STATUS_AUTO_CANCEL     = 103
	ORDER_STATUS_ADMIN_CANCEL    = 104
	ORDER_STATUS_REFUND          = 202
	ORDER_STATUS_REFUND_CONFIRM  = 203
	ORDER_STATUS_AUTO_CONFIRM    = 402
	ORDER_STATUS_PAY_GROUPON     = 200
	ORDER_STATUS_TIMEOUT_GROUPON = 204
)

const (
	MALL_USER_STATUS_ENABLE  = 0
	MALL_USER_STATUS_DISABLE = 1
	MALL_USER_STATUS_DESTROY = 2
)

type MallOrder struct {
	Model

	App           string     `json:"-" faker:"-"`
	UserID        int        `json:"user_id" db:"user_id"`
	OrderSN       string     `json:"order_sn" db:"order_sn"`
	OrderStatus   int        `json:"order_status" db:"order_status"`
	Consignee     string     `json:"consignee" db:"consignee"`
	Mobile        string     `json:"mobile" db:"mobile"`
	Address       string     `json:"address" db:"address"`
	Message       string     `json:"message" db:"message"`
	GoodsPrice    float64    `json:"goods_price" db:"goods_price"`
	FreightPrice  float64    `json:"freight_price" db:"freight_price"`
	CouponPrice   float64    `json:"coupon_price" db:"coupon_price"`
	IntegralPrice float64    `json:"integral_price" db:"integral_price"`
	GrouponPrice  float64    `json:"groupon_price" db:"groupon_price"`
	OrderPrice    float64    `json:"order_price" db:"order_price"`
	ActualPrice   float64    `json:"actual_price" db:"actual_price"`
	PayID         string     `json:"pay_id" db:"pay_id"`
	PayTime       *time.Time `json:"pay_time" db:"pay_time"`
	ShipSN        string     `json:"ship_sn" db:"ship_sn"`
	ShipChannel   string     `json:"ship_channel" db:"ship_channel"`
	ShipTime      *time.Time `json:"ship_time" db:"ship_time"`
	RefundAmount  float64    `json:"refund_amount" db:"refund_amount"`
	RefundType    string     `json:"refund_type" db:"refund_type"`
	RefundContent string     `json:"refund_content" db:"refund_content"`
	RefundTime    *time.Time `json:"refund_time" db:"refund_time"`
	ConfirmTime   *time.Time `json:"confirm_time" db:"confirm_time"`
	Comments      int        `json:"comments" db:"comments"`
	EndTime       *time.Time `json:"end_time" db:"end_time"`
}

type MallOrderGoods struct {
	Model

	App            string  `json:"-" faker:"-"`
	OrderID        int     `json:"order_id" db:"order_id"`
	GoodsID        int     `json:"goods_id" db:"goods_id"`
	GoodsName      string  `json:"goods_name" db:"goods_name"`
	GoodsSN        string  `json:"goods_sn" db:"goods_sn"`
	ProductID      int     `json:"product_id" db:"product_id"`
	Number         int     `json:"number" db:"number"`
	Price          float64 `json:"price" db:"price"`
	Specifications string  `json:"specifications" db:"specifications"`
	PicURL         string  `json:"pic_url" db:"pic_url"`
	Comment        int     `json:"comment" db:"comment"`
}

type MallUser struct {
	Model

	App           string     `json:"-" faker:"-"`
	Username      string     `json:"username" db:"username" faker:"username"`
	Password      string     `json:"password" db:"password" faker:"password"`
	Gender        int        `json:"gender" db:"gender" faker:"oneof: 0, 1, 2"`
	Birthday      string     `json:"birthday" db:"birthday" faker:"date"`
	LastLoginTime *time.Time `json:"last_login_time" db:"last_login_time"`
	LastLoginIP   string     `json:"last_login_ip" db:"last_login_ip" faker:"ipv4"`
	UserLevel     int        `json:"user_level" db:"user_level" faker:"oneof: 0, 1, 2"`
	Nickname      string     `json:"nickname" db:"nickname" faker:"name"`
	Mobile        string     `json:"mobile" db:"mobile" faker:"phone_number"`
	Avatar        string     `json:"avatar" db:"avatar" faker:"url"`
	Platform      string     `json:"platform" db:"platform" faker:"oneof: wx, tt"`
	OpenID        string     `json:"openid" db:"openid" faker:"uuid_digit"`
	SessionKey    string     `json:"session_key" db:"session_key" faker:"uuid_digit"`
	Status        int        `json:"status" db:"status" faker:"oneof: 0, 1, 2"`
}

type MallAddress struct {
	Model

	App           string `json:"-" faker:"-"`
	Name          string `json:"name" db:"name"`
	UserID        int    `json:"userId" db:"userId"`
	Province      string `json:"province" db:"province"`
	City          string `json:"city" db:"city"`
	County        string `json:"county" db:"county"`
	AddressDetail string `json:"addressDetail" db:"addressDetail"`
	AreaCode      string `json:"areaCode" db:"areaCode"`
	PostalCode    string `json:"postalCode" db:"postalCode"`
	Tel           string `json:"tel" db:"tel"`
	IsDefault     bool   `json:"isDefault" db:"isDefault"`
}

type MallSearchHistory struct {
	Model

	App     string `json:"-" faker:"-"`
	UserID  int    `json:"userId" db:"userId"`
	Keyword string `json:"keyword" db:"keyword"`
	From    string `json:"from" db:"from"`
}

type MallCart struct {
	Model

	App            string  `json:"-" faker:"-"`
	UserID         int     `json:"user_id" db:"user_id"`
	GoodsID        int     `json:"goods_id" db:"goods_id"`
	GoodsSN        string  `json:"goods_sn" db:"goods_sn"`
	GoodsName      string  `json:"goods_name" db:"goods_name"`
	ProductID      int     `json:"product_id" db:"product_id"`
	Price          float64 `json:"price" db:"price"`
	Number         int     `json:"number" db:"number"`
	Specifications string  `json:"specifications" db:"specifications"`
	Checked        bool    `json:"checked" db:"checked"`
	PicURL         string  `json:"pic_url" db:"pic_url"`
}

type MallGoods struct {
	Model

	App          string  `json:"-" faker:"-"`
	GoodsSN      string  `json:"goods_sn" db:"goods_sn"`
	Name         string  `json:"name" db:"name"`
	CategoryID   int     `json:"category_id" db:"category_id"`
	BrandID      int     `json:"brand_id" db:"brand_id"`
	Gallery      string  `json:"gallery" db:"gallery" faker:"-"`
	Keywords     string  `json:"keywords" db:"keywords" faker:"word"`
	Brief        string  `json:"brief" db:"brief" faker:"sentence"`
	IsOnSale     bool    `json:"is_on_sale" db:"is_on_sale"`
	SortOrder    int     `json:"sort_order" db:"sort_order"`
	PicURL       string  `json:"pic_url" db:"pic_url" faker:"url"`
	ShareURL     string  `json:"share_url" db:"share_url" faker:"url"`
	IsNew        bool    `json:"is_new" db:"is_new"`
	IsHot        bool    `json:"is_hot" db:"is_hot"`
	Unit         string  `json:"unit" db:"unit" faker:"word"`
	CounterPrice float64 `json:"counter_price" db:"counter_price"`
	RetailPrice  float64 `json:"retail_price" db:"retail_price"`
	Detail       string  `json:"detail" db:"detail" faker:"sentence"`
}

type MallGoodsAttribute struct {
	Model

	App       string `json:"-" faker:"-"`
	GoodsID   int    `json:"goods_id" db:"goods_id" faker:"-"`
	Attribute string `json:"attribute" db:"attribute" faker:"word"`
	Value     string `json:"value" db:"value" faker:"sentence"`
	SortOrder int    `json:"sortOrder" db:"sortOrder"`
}

type MallGoodsProduct struct {
	Model

	App            string  `json:"-" faker:"-"`
	GoodsID        int     `json:"goods_id" db:"goods_id" faker:"-"`
	Specifications string  `json:"specifications" db:"specifications" faker:"-"`
	Price          float64 `json:"price" db:"price"`
	Number         int     `json:"number" db:"number"`
	URL            string  `json:"url" db:"url" faker:"url"`
}

type MallGoodsSpecification struct {
	Model

	App           string `json:"-" faker:"-"`
	GoodsID       int    `json:"goods_id" db:"goods_id" faker:"-"`
	Specification string `json:"specification" db:"specification" faker:"-"`
	Value         string `json:"value" db:"value"`
	PicURL        string `json:"pic_url" db:"pic_url" faker:"url"`
}

type MallBrand struct {
	Model

	App        string  `json:"-" faker:"-"`
	Name       string  `json:"name" db:"name"`
	Desc       string  `json:"desc" db:"desc" faker:"sentence"`
	PicURL     string  `json:"pic_url" db:"pic_url" faker:"url"`
	SortOrder  int     `json:"sort_order" db:"sort_order"`
	FloorPrice float64 `json:"floor_price" db:"floor_price"`
}

type MallCategory struct {
	Model

	App       string `json:"-" faker:"-"`
	Name      string `json:"name" db:"name"`
	Keywords  string `json:"keywords" db:"keywords" faker:"word"`
	Desc      string `json:"desc" db:"desc" faker:"sentence"`
	Pid       int    `json:"pid" db:"pid" faker:"-"`
	IconURL   string `json:"icon_url" db:"icon_url" faker:"url"`
	PicURL    string `json:"pic_url" db:"pic_url" faker:"url"`
	Level     string `json:"level" db:"level" faker:"-"`
	SortOrder int    `json:"sort_order" db:"sort_order"`
}

type MallIssue struct {
	Model

	App      string `json:"-" faker:"-"`
	Question string `json:"question" db:"question" faker:"sentence"`
	Answer   string `json:"answer" db:"answer" faker:"sentence"`
}
