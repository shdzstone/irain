package models

import "irain/ginServer/util/dao"

type Node struct {
	Model
	Pid uint		`gorm:"not null" json:"pid"`
	AdminId uint	`gorm:"not null" json:"admin_id"`
	NodeName string	`gorm:"not null" json:"node_name"`
	Route string	`gorm:"null" json:"route"`
	NodeAddress string`gorm:"null" json:"node_address"`
	Method string	`gorm:"not null" json:"method"`
	Icon string		`gorm:"null" json:"icon"`
	Type uint		`gorm:"not null" json:"type"`
	Status uint		`gorm:"not null" json:"status"`
}

func GetAllNode(adminId uint) (nodeList []*Node, err error){
	if err = dao.DB().Where("admin_id=?",adminId).Order("type,id").Find(&nodeList).Error; err != nil{
		return nil, err
	}
	return
}