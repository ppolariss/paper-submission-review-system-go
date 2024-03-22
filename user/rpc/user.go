package rpc

import (
	"github.com/ppolariss/paper-submission-review-system-go/user/model"
	"github.com/ppolariss/paper-submission-review-system-go/user/schema"
	"gorm.io/gorm/clause"
)

type UserRPC struct{}

func (u *UserRPC) LoadUserRolesInConference(args schema.UserConferenceRole, reply *[]schema.Role) error {
	return model.DB.
		Model(&model.UserConferenceRole{}).
		Where("user_id = ? and conference_id = ?", args.UserID, args.ConferenceID).
		Pluck("role", reply).Error
}

func (u *UserRPC) LoadUserRolesInConferenceByUsername(userID int, reply *[]schema.UserConferenceRole) error {
	return model.DB.
		Model(&model.UserConferenceRole{}).
		Where("user_id = ?", userID).
		Find(reply).Error
}

func (u *UserRPC) AddRoleToUserInConference(args schema.UserConferenceRole, reply *string) error {
	role := args.ToModel()
	result := model.DB.Clauses(clause.OnConflict{DoNothing: true}).Create(&role)
	if result.Error != nil {
		return result.Error
	}
	if result.RowsAffected == 0 {
		*reply = "已存在"
	} else {
		*reply = "添加成功"
	}
	return nil
}

func (u *UserRPC) RemoveRoleFromUserInConference(args schema.UserConferenceRole, reply *string) error {
	role := args.ToModel()
	result := model.DB.Delete(&role)
	if result.Error != nil {
		return result.Error
	}
	if result.RowsAffected == 0 {
		*reply = "不存在"
	} else {
		*reply = "删除成功"
	}
	return nil
}

func (u *UserRPC) CheckRoleOfUserInConference(args schema.UserConferenceRole, reply *bool) error {
	var count int64
	result := model.DB.Model(&model.UserConferenceRole{}).
		Where("user_id = ? and conference_id = ? and role = ?", args.UserID, args.ConferenceID, args.Role).
		Count(&count)
	if result.Error != nil {
		return result.Error
	}
	*reply = count > 0
	return nil
}
