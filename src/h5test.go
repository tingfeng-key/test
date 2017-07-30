package src

import (
	"strconv"
)

type requestData struct {
	mark string "支付接口标识"
	merchId int "商户ID"
	userId int "用户ID"
	appId int "平台分配的应用ID"
	orderNo string "订单号码"
	amount float64 "金额"
	body string "商品描述"
	name string "商家名称"
	detail string "商品详情"
	params string "穿透参数"
	returnUrl string "支付返回地址"
	noticeUrl string "后台通知地址"
}
func Pay(userId int, orderNo string, amount float64) (string){
	data := requestData{
		mark: "payOrder",
		merchId: 1,
		userId: userId,
		appId: 1,
		orderNo: orderNo,
		amount:amount,
		body: "充值金币",
		name: "腾讯游戏",
		detail: "商品详情",
		params: "",
		returnUrl: "http://192.168.3.5/return",
		noticeUrl: "http://192.168.3.5/noticy",
	};
	return "mark="+data.mark+"&merchId="+string(data.merchId)+"&userId="+
		string(data.userId)+"&appId="+string(data.appId)+
		"&orderNo="+data.orderNo+"&amount="+
		strconv.FormatFloat(data.amount, 'f', 2, 32)+
		"&body="+data.body+"&name="+data.name+
		"&noticeUrl=123"+"&returnUrl=123";//+url.QueryEscape(data.returnUrl)++url.QueryEscape(data.noticeUrl);

}