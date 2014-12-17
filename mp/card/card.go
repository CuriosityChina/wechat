// @description wechat 是腾讯微信公众平台 api 的 golang 语言封装
// @link        https://github.com/chanxuehong/wechat for the canonical source repository
// @license     https://github.com/chanxuehong/wechat/blob/master/LICENSE
// @authors     chanxuehong(chanxuehong@gmail.com)

package card

type Card struct {
	CardType      string `json:"card_type"`
	GeneralCoupon struct {
		BaseInfo      BaseInfo `json:"base_info"`
		DefaultDetail string   `json:"default_detail"`
	} `json:"general_coupon,omitempty"`
	Groupon struct {
		BaseInfo   BaseInfo `json:"base_info"`
		DealDetail string   `json:"deal_detail"`
	} `json:"groupon,omitempty"`
	Gift struct {
		BaseInfo BaseInfo `json:"base_info"`
		Gift     string   `json:"gift"`
	} `json:"gift,omitempty"`
	Cash struct {
		BaseInfo   BaseInfo `json:"base_info"`
		LeastCost  int      `json:"least_cost"`
		ReduceCost int      `json:"reduce_cost"`
	} `json:"cash,omitempty"`
	Discount struct {
		BaseInfo BaseInfo `json:"base_info"`
		Discount int      `json:"discount"`
	} `json:"discount,omitempty"`
	MemberCard struct {
		BaseInfo       BaseInfo `json:"base_info"`
		SupplyBonus    bool     `json:"supply_bonus"`
		SupplyBalance  bool     `json:"supply_balance"`
		BonusCleared   string   `json:"bonus_cleared,omitempty"`
		BonusRules     string   `json:"bonus_rules"`
		BalanceRules   string   `json:"balance_rules,omitempty"`
		Prerogative    string   `json:"prerogative"`
		BindOldCardUrl string   `json:"bind_old_card_url,omitempty"`
		ActivateUrl    string   `json:"activate_url,omitempty"`
	} `json:"member_card,omitempty"`
	ScenicTicket struct {
		BaseInfo    BaseInfo `json:"base_info"`
		TicketClass string   `json:"ticket_class,omitempty"`
		GuideURl    string   `json:"guide_url,omitempty"`
	} `json:"scenic_ticket,omitempty"`
	MovieTicket struct {
		BaseInfo BaseInfo `json:"base_info"`
		Detail   string   `json:"detail,omitempty"`
	} `json:"movie_ticket,omitempty"`
	BoardingPass struct {
		BaseInfo      BaseInfo `json:"base_info"`
		From          string   `json:"from"`
		To            string   `json:"to"`
		Flight        string   `json:"flight"`
		DepartureTime string   `json:"departure_time,omitempty"`
		LandingTime   string   `json:"landing_time,omitempty"`
		CheckInUrl    string   `json:"check_in_url,omitempty"`
		AirModel      string   `json:"air_model,omitempty"`
	} `json:"boarding_pass,omitempty"`
	LuckMoney struct {
		BaseInfo BaseInfo `json:"base_info,omitempty"`
	} `json:"lucky_money,omitempty"`
}

// 优惠券的基本类型
type BaseInfo struct {
	// 卡券的商户尺寸，300*300
	LogoUrl        string  `json:"logo_url"`
	CodeType       string  `json:"code_type"`
	BrandName      string  `json:"brand_name"`
	Title          string  `json:"title"` // 9个汉字
	SubTitle       string  `json:"sub_title,omitempty"`
	Color          string  `json:"color"`
	Notice         string  `json:"notice"`
	ServicePhone   string  `json:"service_phone,omitempty"`
	Source         string  `json:"source,omitempty"`
	Description    string  `json:"description"`
	UseLimit       int     `json:"user_limit,omitempty"`
	GetLimit       int     `json:"get_limit,omitempty"`
	UseCustomCode  bool    `json:"use_custom_code,omitempty"`
	BindOpenid     bool    `json:"bind_openid,omitempty"`
	CanShare       bool    `json:"can_share,omitempty"`
	CanGiveFriend  bool    `json:"can_give_friend,omitempty"`
	LocationIdList []int64 `json:"location_id_list,omitempty"`
	DateInfo       `json:"date_info"`
	Sku            struct {
		Quantity int64 `json:"quantity"`
	} `json:"sku"`
	UrlNameType string `json:"url_name_type,omitempty"`
	CustomUrl   string `json:"custom_url,omitempty"`
}

func (baesinfo *BaseInfo) Init(logo_url, code_type, brand_name, title, color, notice, description string, date_info DateInfo, sku_quantity int64) {
	baesinfo.LogoUrl = logo_url
	baesinfo.CodeType = code_type
	baesinfo.BrandName = brand_name
	baesinfo.Title = title
	baesinfo.Color = color
	baesinfo.Notice = notice
	baesinfo.Description = description
	baesinfo.DateInfo = date_info
	baesinfo.Sku.Quantity = sku_quantity
}

// 使用日期
type DateInfo struct {
	Type           int   `json:"type"`
	BeginTimestamp int64 `json:"begin_timestamp"`
	EndTimestamp   int64 `json:"end_timestamp"`
	FixedTerm      int   `json:"fixed_term"`
	FixedBeginTerm int   `json:"fixed_begin_term"`
}

func (dateinfo *DateInfo) NewTimestamp(begin int64, end int64) {
	dateinfo.Type = Date_Info_Type_Begin_End
	dateinfo.BeginTimestamp = begin
	dateinfo.EndTimestamp = end
}

// fixed term 是 天
func (dateinfo *DateInfo) NewFixed(fixed_term int, fixed_begin_term int) {
	dateinfo.Type = Date_Info_Type_Fixed_Term
	dateinfo.FixedTerm = fixed_term
	dateinfo.FixedBeginTerm = fixed_begin_term
}

func (card *Card) InitGeneralCoupon(base_info BaseInfo, default_detail string) {
	card.CardType = CARD_TYPE_GENERAL_COUPON
	card.GeneralCoupon.BaseInfo = base_info
	card.GeneralCoupon.DefaultDetail = default_detail
}
func (card *Card) InitGroupon(base_info BaseInfo, deal_detail string) {
	card.CardType = CARD_TYPE_GROUPON
	card.Groupon.BaseInfo = base_info
	card.Groupon.DealDetail = deal_detail
}
func (card *Card) InitGift(base_info BaseInfo, gift string)                           {}
func (card *Card) InitCash(base_info BaseInfo, least_cost string, reduce_cost string) {}
func (card *Card) InitDiscount(base_info BaseInfo, discount int)                      {}
func (card *Card) InitMemberCard(
	base_info BaseInfo, supply_bonus bool,
	supply_balance, bonus_cleared, bonus_rules, balance_rules, prerogative, bind_old_card_url, activate_url string) {
}
func (card *Card) InitScenicTicket(base_info BaseInfo, ticket_class, guide_url string) {}
func (card *Card) InitMovieTicket(base_info BaseInfo, detail string)                   {}
func (card *Card) InitBoardingPass(base_info BaseInfo, from, to, flight, departure_time, landing_time, check_in_url, air_model string) {
}
func (card *Card) InitLuckMoney(base_info BaseInfo) {}

func (card *Card) CheckValidate() {

}
