package services

import (
	"churchflowx/internal/models"
	"churchflowx/internal/objects"
)

// ######################## MEMBERS SERVICES ###########################

func AddMemberToDb(member *objects.Member) bool {
	new_member := models.Member{
		GID:            member.GID,
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

func GetMemberFromDb(current_user_id, member_id string) (bool, models.Member) {
	var member models.Member
	result := db.Where("g_id = ? AND id = ?",current_user_id, member_id ).Find(&member)
	return result.Error == nil, member
}

func GetMembersFromDb(google_id string) (bool, []models.Member) {
	var members []models.Member
	result := db.Where("g_id = ?", google_id).Find(&members)
	return result.Error == nil, members
}

func UpdateMemberInDb(taks *objects.Member) (bool, string) {

	return true, "success"
}

func DeleteMemberFromDb(taks *objects.Member) (bool, string) {

	return true, "success"
}

// ######################## MINISTERS SERVICES ###########################

func AddMinisterToDb(minister *objects.Minister) bool {
	new_minister := models.Minister{
		GID:            minister.GID,
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

func GetMinisterFromDb(current_user_id, minister_id string) (bool, models.Minister) {
	var minister models.Minister
	result := db.Where("g_id = ? AND id = ?",current_user_id, minister_id ).Find(&minister)
	return result.Error == nil, minister
}

func GetMinistersFromDb(google_id string) (bool, []models.Minister) {
	var ministers []models.Minister
	result := db.Where("g_id = ?", google_id).Find(&ministers)
	return result.Error == nil, ministers
}

func UpdateMinisterInDb(taks *objects.Minister) (bool, string) {

	return true, "success"
}

func DeleteMinisterFromDb(taks *objects.Minister) (bool, string) {

	return true, "success"
}

// ######################## VISITORS SERVICES ###########################

func AddVisitorToDb(visitor *objects.Visitor) bool {
	new_visitor := models.Visitor{
		GID:          visitor.GID,
		Name:         visitor.Name,
		Tel:          visitor.Tel,
		Email:        visitor.Email,
		HomeDistrict: visitor.HomeDistrict,
	}
	result := db.Create(&new_visitor)
	return result.Error == nil
}

func GetVisitorFromDb(current_user_id, visitor_id string) (bool, models.Visitor) {
	var visitor models.Visitor
	result := db.Where("g_id = ? AND id = ?",current_user_id, visitor_id ).Find(&visitor)
	return result.Error == nil, visitor
}

func GetVisitorsFromDb(google_id string) (bool, []models.Visitor) {
	var visitors []models.Visitor
	result := db.Where("g_id = ?", google_id).Find(&visitors)
	return result.Error == nil, visitors
}

func UpdateVisitorInDb(taks *objects.Visitor) (bool, string) {

	return true, "success"
}

func DeleteVisitorFromDb(taks *objects.Visitor) (bool, string) {

	return true, "success"
}

// ######################## PASTORS SERVICES ###########################

func AddPastorToDb(pastor *objects.Pastor) bool {
	new_pastor := models.Pastor{
		GID:            pastor.GID,
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

func GetPastorFromDb(current_user_id, pastor_id string) (bool, models.Pastor) {
	var pastor models.Pastor
	result := db.Where("g_id = ? AND id = ?",current_user_id, pastor_id ).Find(&pastor)
	return result.Error == nil, pastor
}

func GetPastorsFromDb(google_id string) (bool, []models.Pastor) {
	var pastors []models.Pastor
	result := db.Where("g_id = ?", google_id).Find(&pastors)
	return result.Error == nil, pastors
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
		GID:        user_admin.GID,
		Name:       user_admin.Name,
		Email:      user_admin.Email,
		ProfileImg: user_admin.ProfileImg,
	}
	result := db.Create(&new_admin)
	return result.Error == nil
}

func GetAdminFromDb(google_id string) (bool, []models.Admin) {
	var admin []models.Admin
	result := db.Where("g_id = ?", google_id).Find(&admin)
	return result.Error == nil, admin
}

func GetAdminsFromDb(user_admin *objects.Admin) bool {

	return true
}

func UpdateAdminInDb(user_admin *objects.Admin) bool {

	return true
}

func DeleteAdminFromDb(user_admin *objects.Admin) bool {

	return true
}
