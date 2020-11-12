package models

import (
	"irain/ginServer/util/dao"
)

type Admin struct {
	Model
	Account   string `gorm:"unique;not null" json:"account,omitempty"`
	Password  string `gorm:"not null" json:"password,omitempty"`
	NickName  string `gorm:"not null" json:"nick_name,omitempty"`
	Avatar    string `json:"avatar,omitempty"`
	BriefDesc string `gorm:"null" json:"brief_desc,omitempty"`
}

func CreateAnAdmin(admin *Admin) (err error) {
	err = dao.DB().Create(&admin).Error
	return
}

func GetAllAdmin() (adminList []*Admin, err error) {
	if err = dao.DB().Find(&adminList).Error; err != nil {
		return nil, err
	}
	return
}

func GetExistAdmin(acc, pwd string) (admin *Admin, err error) {
	admin = new(Admin)
	if err = dao.DB().Where("account=?", acc).Where("password=?", pwd).First(admin).Error; err != nil {
		return nil, err
	}
	
	return admin, nil
}

func GetAnAdmin(id string) (admin *Admin, err error) {
	admin = new(Admin)
	if err = dao.DB().Where("id=?", id).First(admin).Error; err != nil {
		return nil, err
	}
	return
}

func UpdateAnAdmin(admin *Admin) (err error) {
	err = dao.DB().Save(admin).Error
	return
}

func DeleteAnAdmin(id string) (err error) {
	err = dao.DB().Where("id=?", id).Delete(&Admin{}).Error
	return
}
