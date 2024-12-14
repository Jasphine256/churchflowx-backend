package services

import (
	"churchflowx/internal/models"
	"churchflowx/internal/objects"
)

// ######################## MEMBERS SERVICES ###########################

func AddMemberToDb(user_id int, member *objects.Member) bool {
	new_member := models.Member{
		UserId:         user_id,
		Date:           member.Date,
		Name:           member.Name,
		Dob:            member.Dob,
		Zone:           member.Zone,
		Village:        member.Village,
		Parish:         member.Parish,
		Subcounty:      member.Subcounty,
		FormerReligion: member.FormerReligion,
		Educ:           member.Educ,
		Occupation:     member.Occupation,
		Where:          member.Where,
		MaritalStatus:  member.MaritalStatus,
		Children:       member.Children,
		Tel:            member.Tel,
		Email:          member.Email,
		Father:         member.Father,
		Mother:         member.Mother,
		HomeVillage:    member.HomeVillage,
		HomeParish:     member.HomeParish,
		HomeSubCounty:  member.HomeSubCounty,
		HomeDistrict:   member.HomeDistrict,
	}
	result := db.Create(&new_member)
	return result.Error == nil
}

func GetMemberFromDb(taks *objects.Member) (bool, string) {

	return true, "success"
}

func GetMembersFromDb(taks *objects.Member) (bool, string) {

	return true, "success"
}

func UpdateMemberInDb(taks *objects.Member) (bool, string) {

	return true, "success"
}

func DeleteMemberFromDb(taks *objects.Member) (bool, string) {

	return true, "success"
}

// ######################## MINISTERS SERVICES ###########################

func AddMinisterToDb(user_id int, minister *objects.Minister) bool {
	new_minister := models.Minister{
		UserId:         user_id,
		Date:           minister.Date,
		Name:           minister.Name,
		Ministry:       minister.Ministry,
		Role:           minister.Role,
		Dob:            minister.Dob,
		Zone:           minister.Zone,
		Village:        minister.Village,
		Parish:         minister.Parish,
		Subcounty:      minister.Subcounty,
		FormerReligion: minister.FormerReligion,
		Educ:           minister.Educ,
		Occupation:     minister.Occupation,
		Where:          minister.Where,
		MaritalStatus:  minister.MaritalStatus,
		Children:       minister.Children, // Convert Children from int to uint8
		Tel:            minister.Tel,
		Email:          minister.Email,
		Father:         minister.Father,
		Mother:         minister.Mother,
		HomeVillage:    minister.HomeVillage,
		HomeParish:     minister.HomeParish,
		HomeSubCounty:  minister.HomeSubCounty,
		HomeDistrict:   minister.HomeDistrict,
	}
	result := db.Create(&new_minister)
	return result.Error == nil
}

func GetMinisterFromDb(taks *objects.Minister) (bool, string) {

	return true, "success"
}

func GetMinistersFromDb(taks *objects.Minister) (bool, string) {

	return true, "success"
}

func UpdateMinisterInDb(taks *objects.Minister) (bool, string) {

	return true, "success"
}

func DeleteMinisterFromDb(taks *objects.Minister) (bool, string) {

	return true, "success"
}

// ######################## VISITORS SERVICES ###########################

func AddVisitorToDb(user_id int, visitor *objects.Visitor) bool {
	new_visitor := models.Visitor{
		UserId:       user_id,
		Tel:          visitor.Tel,
		Email:        visitor.Email,
		HomeDistrict: visitor.HomeDistrict,
	}
	result := db.Create(&new_visitor)
	return result.Error == nil
}

func GetVisitorFromDb(taks *objects.Visitor) (bool, string) {

	return true, "success"
}

func GetVisitorsFromDb(taks *objects.Visitor) (bool, string) {

	return true, "success"
}

func UpdateVisitorInDb(taks *objects.Visitor) (bool, string) {

	return true, "success"
}

func DeleteVisitorFromDb(taks *objects.Visitor) (bool, string) {

	return true, "success"
}

// ######################## PASTORS SERVICES ###########################

func AddPastorToDb(user_id int, pastor *objects.Pastor) bool {
	new_pastor := models.Pastor{
		UserId:         user_id,
		Date:           pastor.Date,
		Name:           pastor.Name,
		Ministry:       pastor.Ministry,
		Dob:            pastor.Dob,
		PastorSince:    pastor.PastorSince,
		Zone:           pastor.Zone,
		Village:        pastor.Village,
		Parish:         pastor.Parish,
		Subcounty:      pastor.Subcounty,
		FormerReligion: pastor.FormerReligion,
		Educ:           pastor.Educ,
		Occupation:     pastor.Occupation,
		Where:          pastor.Where,
		MaritalStatus:  pastor.MaritalStatus,
		Children:       pastor.Children,
		Tel:            pastor.Tel,
		Email:          pastor.Email,
		Father:         pastor.Father,
		Mother:         pastor.Mother,
		HomeVillage:    pastor.HomeVillage,
		HomeParish:     pastor.HomeParish,
		HomeSubCounty:  pastor.HomeSubCounty,
		HomeDistrict:   pastor.HomeDistrict,
	}
	result := db.Create(&new_pastor)
	return result.Error == nil
}

func GetPastorFromDb(taks *objects.Pastor) (bool, string) {

	return true, "success"
}

func GetPastorsFromDb(taks *objects.Pastor) (bool, string) {

	return true, "success"
}

func UpdatePastorInDb(taks *objects.Pastor) (bool, string) {

	return true, "success"
}

func DeletePastorFromDb(taks *objects.Pastor) (bool, string) {

	return true, "success"
}

// ######################## ADMINS SERVICES ###########################

func AddAdminToDb(user_admin *objects.Admin) bool {
	new_admin := models.Admin{
		Name:       user_admin.Name,
		Email:      user_admin.Email,
		ProfileImg: user_admin.ProfileImg,
	}
	result := db.Create(&new_admin)
	return result.Error == nil
}

func GetAdminFromDb(user_admin *objects.Admin) bool {

	return true
}

func GetPAdminFromDb(user_admin *objects.Admin) bool {

	return true
}

func UpdateAdminInDb(user_admin *objects.Admin) bool {

	return true
}

func DeleteAdminFromDb(user_admin *objects.Admin) bool {

	return true
}
