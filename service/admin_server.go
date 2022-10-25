package service

import (
	"fmt"
	"github.com/go-xorm/xorm"
	"github.com/kataras/iris/v12"
	"webProject/model"
)

/*
service层：标准的开发模式是将每个实体的功能以接口形式定义，供控制层进行调用
来自网络：/在我们实际的开发过程中，我们往往将数据提供服务模块设计成接口，这样设计的目的是接口定义和具体
//的功能编程实现了分离，有助于我们在不同的实现方案之间进行切换，成本非常小

*/

type AdminService interface {

	//数据库查询方法
	GetAdminNameAndPassword(name string, pwd string) (model.Admin, bool)

	//统计方法
	GetAdminCount() (int64, error)

	PostAdminInfo(offset, limit int) []model.Admin

	PostAvatar(adminId int, fileName string) error
}

// 定义一个AdminService的（内部）实现类,注意这个实现类首字母小写
type adminService struct {
	//实现类多一个成员变量
	engine *xorm.Engine
}

// 为了给内部实现类成员变量赋值还要写一个方法实现？
func NewAdminService(db *xorm.Engine) AdminService {
	//这个地方为什么取地址？在没有实现全部方法时候，还不能说adminService是AdminService的实现类,不实现会报错
	return &adminService{
		engine: db,
	}
}

// 为什么在这实现书库查询，不在datasource？
func (ads *adminService) GetAdminNameAndPassword(name string, pwd string) (model.Admin, bool) {
	var admin model.Admin
	ads.engine.Where("admin_name = ? and pwd = ?", name, pwd).Get(&admin)
	fmt.Printf("admin:%s,adminid:%d\n", admin, admin.AdminId)
	return admin, admin.AdminId != 0

}

func (ads *adminService) GetAdminCount() (int64, error) {
	admin := new(model.Admin)
	//conut函数返回的是int64类型
	total, err := ads.engine.Count(admin)
	if err != nil {
		panic(err.Error())
		return 0, err
	}

	return total, nil
}

func (ads *adminService) PostAdminInfo(offset, limit int) []model.Admin {
	adminInfo := make([]model.Admin, 0)
	//limit（参数1，参数2）参数1是每页多少内容，参数2是查询起始位置，需要计算一下offset
	err := ads.engine.Limit(offset, limit*offset).Find(&adminInfo)
	if err != nil {
		panic(err.Error())
	}
	return adminInfo
}

// 更新头像数据到数据库中
func (ads *adminService) PostAvatar(adminId int, fileName string) error {
	admin := model.Admin{
		Avatar: fileName,
	}
	engin := ads.engine
	affected, err := engin.ID(adminId).Cols("avatar").Update(&admin)
	iris.New().Logger().Info("更新记录：" + string(affected))
	return err
}
