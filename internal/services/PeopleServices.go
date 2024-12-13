package services

import (
	"churchflowx/internal/models"
	"churchflowx/internal/objects"
	"log"
)

// ######################## MEMBERS SERVICES ###########################

func AddMemberToDb(taks *objects.Member) (bool, string) {

	return true, "success"
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

func AddMinisterToDb(taks *objects.Minister) (bool, string) {

	return true, "success"
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

func AddVisitorToDb(taks *objects.Visitor) (bool, string) {

	return true, "success"
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

func AddPastorToDb(taks *objects.Pastor) (bool, string) {

	return true, "success"
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
	if result.Error != nil {
		log.Fatal(result.Error)
		// return false
	}
	db.Create(&new_admin)
	return result.Error == nil
}

func GetAdminFromDb(user_admin *objects.Admin) (bool, string) {

	return true, "success"
}

func GetPAdminFromDb(user_admin *objects.Admin) (bool, string) {

	return true, "success"
}

func UpdateAdminInDb(user_admin *objects.Admin) (bool, string) {

	return true, "success"
}

func DeleteAdminFromDb(user_admin *objects.Admin) (bool, string) {

	return true, "success"
}
