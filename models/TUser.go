package models

import (
	"github.com/astaxie/beego/orm"
	//_"github.com/go-sql-driver/mysql"
)

//用户表
type TUser struct {
	//用户序号
	Id int64
	//电话号码
	Tep string
	//密码
	Pwd string
	//收款人
	Payee string
	//地址
	Address string
	//收款帐号
	Amount string
	//账号类别
	AmountType string
	//是否消费者
	IsCustomer bool
	//是否商家
	IsSeller bool
	//是否配送员
	IsDiliver bool
	//是否管理员
	IsManager bool
	//微信openId
	Vid string
	//是否冻结
	IsLock bool
	//创建时间 --- 时间戳
	AddTime int64
}

//新建用户
func AddUser(user *TUser) (int64, error) {
	o := orm.NewOrm() //数据库
	userId, err := o.Insert(user) //插入数据
	return userId, err
}

//查询账号
func GetUserById(userId int64) (*TUser, error) {
	o := orm.NewOrm() //数据库
	user := new(TUser) //TUser就是第9行struct的数据库，就是mysql的表
	qs := o.QueryTable("t_user") //表名为t_user
	err := qs.Filter("id", userId).One(user) //One是指查询一条数据，One(user)是查询mysql表中一条数据
	return user, err
}

//手机号查询账号
func GetUserByTel(tel string) (*TUser, error) {
	o := orm.NewOrm()
	user := new(TUser) //TUser就是第9行struct的数据库，就是mysql的表
	qs := o.QueryTable("t_user") //表名为t_user
	err := qs.Filter("tel", tel).One(user) //One是指查询一条数据，One(user)是查询mysql表中一条数据
	return user, err
}

//微信Id查询账号
func GetUserByVid(vid int64) (*TUser, error) {
	o := orm.NewOrm()
	user := new(TUser) //TUser就是第9行struct的数据库，就是mysql的表
	qs := o.QueryTable("t_user") //表名为t_user
	err := qs.Filter("vid", vid).One(user) //One是指查询一条数据，One(user)是查询mysql表中一条数据
	return user, err
}