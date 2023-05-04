package main

import (
	"fmt"
	"focus/common"
	"github.com/gin-gonic/gin"
	"gorm.io/driver/mysql"
	"gorm.io/gorm"
)

var db *gorm.DB

func main() {
	dsn := "root:123456@tcp(127.0.0.1:3306)/focus?charset=utf8mb4&parseTime=True&loc=Local"
	dbT, err := gorm.Open(mysql.Open(dsn), &gorm.Config{})
	if err != nil {
		fmt.Println("连接数据库错误", err)
		return
	}
	db = dbT
	r := gin.Default()
	//api := r.Group("api")
	r.GET("/api/v1/order/list", handler)
	r.Run() // 监听并在 0.0.0.0:8080 上启动服务
}

func handler(c *gin.Context) {
	//获取前端传递参数
	page, err := common.GetQueryInt(c, "page")
	if err != nil {
		common.ErrorResp(c, err)
		return
	}
	//获取前端传递参数
	pers, err := common.GetQueryInt(c, "pers")
	if err != nil {
		common.ErrorResp(c, err)
		return
	}
	// 用户分页及最新创建的订单数据
	var userOrder []*UserOrder
	// 包装用户及分页信息
	var respData = &RespData[UserOrder]{}
	// 获取总量
	var total Total
	// 总量sql
	err = db.Raw("select count(id) as total from p_user").Scan(&total).Error
	if err != nil {
		common.ErrorResp(c, err)
		return
	}
	// 用户分页及最新创建的订单数据sql
	err = db.Raw("select id,created_at,updated_at,nickname,mobile,"+
		"(select p_order.product_name from p_order where p_user.id=p_order.user_id order by created_at desc limit 1) as product_name,"+
		"(select p_order.total_price from p_order where p_user.id=p_order.user_id order by created_at desc limit 1) as total_price,"+
		"(select p_order.count from p_order where p_user.id=p_order.user_id order by created_at desc limit 1) as count,"+
		"(select p_order.unit_price from p_order where p_user.id=p_order.user_id order by created_at desc limit 1) as unit_price,"+
		"(select p_order.pay_type from p_order where p_user.id=p_order.user_id order by created_at desc limit 1) as pay_type,"+
		"(select p_order.status from p_order where p_user.id=p_order.user_id order by created_at desc limit 1) as status,"+
		"(select p_order.created_at from p_order where p_user.id=p_order.user_id order by p_order.created_at desc limit 1) as create_at "+
		"from p_user limit ?, ?", page*pers, pers).Scan(&userOrder).Error
	if err != nil {
		common.ErrorResp(c, err)
		return
	}
	// 赋值
	respData.Total = total.Total
	respData.HasNext = page*pers+pers < respData.Total
	respData.Page = page
	respData.Pers = pers
	respData.Data = userOrder
	// 返回内容
	common.SuccessResp(c, respData)
}
