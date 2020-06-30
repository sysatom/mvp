package model

import (
	"time"
)

type Model struct {
	ID        int       `json:"id" db:"id" faker:"-"`
	CreatedAt time.Time `json:"createdAt" db:"createdAt"`
	UpdatedAt time.Time `json:"updatedAt" db:"updatedAt"`
	DeletedAt time.Time `json:"-" faker:"-"`
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
	UserID        int        `json:"userId" db:"userId"`
	OrderSN       string     `json:"orderSn" db:"orderSn"`
	OrderStatus   int        `json:"orderStatus" db:"orderStatus"`
	Consignee     string     `json:"consignee" db:"consignee"`
	Mobile        string     `json:"mobile" db:"mobile"`
	Address       string     `json:"address" db:"address"`
	Message       string     `json:"message" db:"message"`
	GoodsPrice    float64    `json:"goodsPrice" db:"goodsPrice"`
	FreightPrice  float64    `json:"freightPrice" db:"freightPrice"`
	CouponPrice   float64    `json:"couponPrice" db:"couponPrice"`
	IntegralPrice float64    `json:"integralPrice" db:"integralPrice"`
	GrouponPrice  float64    `json:"grouponPrice" db:"grouponPrice"`
	OrderPrice    float64    `json:"orderPrice" db:"orderPrice"`
	ActualPrice   float64    `json:"actualPrice" db:"actualPrice"`
	PayID         string     `json:"payId" db:"payId"`
	PayTime       *time.Time `json:"payTime" db:"payTime"`
	ShipSN        string     `json:"shipSn" db:"shipSn"`
	ShipChannel   string     `json:"shipChannel" db:"shipChannel"`
	ShipTime      *time.Time `json:"shipTime" db:"shipTime"`
	RefundAmount  float64    `json:"refundAmount" db:"refundAmount"`
	RefundType    string     `json:"refundType" db:"refundType"`
	RefundContent string     `json:"refundContent" db:"refundContent"`
	RefundTime    *time.Time `json:"refundTime" db:"refundTime"`
	ConfirmTime   *time.Time `json:"confirmTime" db:"confirmTime"`
	Comments      int        `json:"comments" db:"comments"`
	EndTime       *time.Time `json:"endTime" db:"endTime"`
}

type MallOrderGoods struct {
	Model

	App            string  `json:"-" faker:"-"`
	OrderID        int     `json:"orderId" db:"orderId"`
	GoodsID        int     `json:"goodsId" db:"goodsId"`
	GoodsName      string  `json:"goodsName" db:"goodsName"`
	GoodsSN        string  `json:"goodsSn" db:"goodsSn"`
	ProductID      int     `json:"productId" db:"productId"`
	Number         int     `json:"number" db:"number"`
	Price          float64 `json:"price" db:"price"`
	Specifications string  `json:"specifications" db:"specifications"`
	PicURL         string  `json:"picUrl" db:"picUrl"`
	Comment        int     `json:"comment" db:"comment"`
}

type MallUser struct {
	Model

	App           string     `json:"-" faker:"-"`
	Username      string     `json:"username" db:"username" faker:"username"`
	Password      string     `json:"password" db:"password" faker:"password"`
	Gender        int        `json:"gender" db:"gender" faker:"oneof: 0, 1, 2"`
	Birthday      string     `json:"birthday" db:"birthday" faker:"date"`
	LastLoginTime *time.Time `json:"lastLoginTime" db:"lastLoginTime"`
	LastLoginIP   string     `json:"lastLoginIp" db:"lastLoginIp" faker:"ipv4"`
	UserLevel     int        `json:"userLevel" db:"userLevel" faker:"oneof: 0, 1, 2"`
	Nickname      string     `json:"nickname" db:"nickname" faker:"name"`
	Mobile        string     `json:"mobile" db:"mobile" faker:"phone_number"`
	Avatar        string     `json:"avatar" db:"avatar" faker:"url"`
	Platform      string     `json:"platform" db:"platform" faker:"oneof: wx, tt"`
	OpenID        string     `json:"openId" db:"openId" faker:"uuid_digit"`
	SessionKey    string     `json:"sessionKey" db:"sessionKey" faker:"uuid_digit"`
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
	UserID         int     `json:"userId" db:"userId"`
	GoodsID        int     `json:"goodsId" db:"goodsId"`
	GoodsSN        string  `json:"goodsSn" db:"goodsSn"`
	GoodsName      string  `json:"goodsName" db:"goodsName"`
	ProductID      int     `json:"productId" db:"productId"`
	Price          float64 `json:"price" db:"price"`
	Number         int     `json:"number" db:"number"`
	Specifications string  `json:"specifications" db:"specifications"`
	Checked        bool    `json:"checked" db:"checked"`
	PicURL         string  `json:"picUrl" db:"picUrl"`
}

type MallGoods struct {
	Model

	App          string  `json:"-" faker:"-"`
	GoodsSN      string  `json:"goodsSn" db:"goodsSn"`
	Name         string  `json:"name" db:"name" binding:"required"`
	CategoryID   int     `json:"categoryId" db:"categoryId"`
	BrandID      int     `json:"brandId" db:"brandId"`
	Gallery      string  `json:"gallery" db:"gallery"`
	Keywords     string  `json:"keywords" db:"keywords"`
	Brief        string  `json:"brief" db:"brief"`
	IsOnSale     bool    `json:"isOnSale" db:"isOnSale"`
	SortOrder    int     `json:"sortOrder" db:"sortOrder"`
	PicURL       string  `json:"picUrl" db:"picUrl"`
	ShareURL     string  `json:"shareUrl" db:"shareUrl"`
	IsNew        bool    `json:"isNew" db:"isNew"`
	IsHot        bool    `json:"isHot" db:"isHot"`
	Unit         string  `json:"unit" db:"unit"`
	CounterPrice float64 `json:"counterPrice" db:"counterPrice"`
	RetailPrice  float64 `json:"retailPrice" db:"retailPrice"`
	Detail       string  `json:"detail" db:"detail"`
}

type MallGoodsAttribute struct {
	Model

	App       string `json:"-" faker:"-"`
	GoodsID   int    `json:"goods_id" db:"goods_id"`
	Attribute string `json:"attribute" db:"attribute"`
	Value     string `json:"value" db:"value"`
	SortOrder int    `json:"sort_order" db:"sort_order"`
}

type MallGoodsProduct struct {
	Model

	App            string  `json:"-" faker:"-"`
	GoodsID        int     `json:"goodsId" db:"goodsId"`
	Specifications string  `json:"specifications" db:"specifications"`
	Price          float64 `json:"price" db:"price"`
	Number         int     `json:"number" db:"number"`
	URL            string  `json:"url" db:"url"`
}

type MallGoodsSpecification struct {
	Model

	App           string `json:"-" faker:"-"`
	GoodsID       int    `json:"goodsId" db:"goodsId"`
	Specification string `json:"specification" db:"specification"`
	Value         string `json:"value" db:"value"`
	PicURL        string `json:"picUrl" db:"picUrl"`
}

type MallBrand struct {
	Model

	App        string  `json:"-" faker:"-"`
	Name       string  `json:"name" db:"name"`
	Desc       string  `json:"desc" db:"desc" faker:"sentence"`
	PicURL     string  `json:"picUrl" db:"picUrl" faker:"url"`
	SortOrder  int     `json:"sortOrder" db:"sortOrder"`
	FloorPrice float64 `json:"floorPrice" db:"floorPrice"`
}

type MallCategory struct {
	Model

	App       string `json:"-" faker:"-"`
	Name      string `json:"name" db:"name"`
	Keywords  string `json:"keywords" db:"keywords"`
	Desc      string `json:"desc" db:"desc"`
	Pid       int    `json:"pid" db:"pid"`
	IconURL   string `json:"iconUrl" db:"iconUrl"`
	PicURL    string `json:"picUrl" db:"picUrl"`
	Level     string `json:"level" db:"level"`
	SortOrder int    `json:"sortOrder" db:"sortOrder"`
}

type MallIssue struct {
	Model

	App      string `json:"-" faker:"-"`
	Question string `json:"question" db:"question" faker:"sentence"`
	Answer   string `json:"answer" db:"answer" faker:"sentence"`
}
