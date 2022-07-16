package consumer

import "time"

type ProductIndex struct {
	Operation   string    `json:"operation"`
	Id          int64     `json:"id"`          //商品id
	StoreName   string    `json:"storeName"`   //商品名称
	StoreInfo   string    `json:"storeInfo"`   //商品简介
	Keyword     string    `json:"keyword"`     //关键字
	CateId      int       `json:"cateId"`      //分类id
	Price       float64   `json:"price"`       //商品价格
	Sales       int32     `json:"sales"`       //销量
	Ficti       int32     `json:"ficti"`       //虚拟销量
	IsHot       int8      `json:"isHot"`       //是否热卖 (0: 否，1：是)
	IsBenefit   int8      `json:"isBenefit"`   //是否优惠(0: 否，1：是)
	IsBest      int8      `json:"isBest"`      //是否精品(0: 否，1：是)
	IsNew       int8      `json:"isNew"`       //是否新品 (0: 否，1：是)
	Description string    `json:"description"` //产品描述
	IsPostage   int8      `json:"isPostage"`   //是否包邮 (0: 否，1：是)
	IsGood      int8      `json:"isGood"`      //是否优品推荐 (0: 否，1：是)
	CreateTime  time.Time `json:"createTime"`
	UpdateTime  time.Time `json:"updateTime"`
}

type ProductMsg struct {
	Operation string `json:"operation"`
	IsShow    int8   `json:"isShow"`
	ProductIndex
}
