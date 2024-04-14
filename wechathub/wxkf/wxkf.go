package wxkf

// 微信客服API文档地址： https://kf.weixin.qq.com/api/doc/path/93304

// 客服账号
const (
	addAccountURL    = "https://qyapi.weixin.qq.com/cgi-bin/kf/account/add"
	delAccountURL    = "https://qyapi.weixin.qq.com/cgi-bin/kf/account/del"
	updateAccountURL = "https://qyapi.weixin.qq.com/cgi-bin/kf/account/update"
	listAccountURL   = "https://qyapi.weixin.qq.com/cgi-bin/kf/account/list"
	linkAccountURL   = "https://qyapi.weixin.qq.com/cgi-bin/kf/add_contact_way"
)

// 客服消息收发
const (
	sendMsgURL         = "https://qyapi.weixin.qq.com/cgi-bin/kf/send_msg"
	sendLuckWelcomeURL = "https://qyapi.weixin.qq.com/cgi-bin/kf/send_msg_on_event"
	reCallMsgURL       = "https://qyapi.weixin.qq.com/cgi-bin/kf/recall_msg"
)

// 基础信息获取
const (
	customerBasicInfoURL = "https://qyapi.weixin.qq.com/cgi-bin/kf/customer/batchget"
)
