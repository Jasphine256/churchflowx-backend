package services

import (
	"churchflowx/internal/database"
	"churchflowx/internal/models"
	"churchflowx/internal/objects"
)

var db = database.InitialiseDB()

// ######################## FUNDS SERVICES ###########################

func AddFundToDb(fund objects.Fund) bool {
	new_fund := models.Fund{
		GID:    fund.GID,
		Name:   fund.Name,
		Reason: fund.Reason,
		Amount: fund.Amount,
		Date:   fund.Date,
	}
	result := db.Create(&new_fund)
	return result.Error == nil
}

func GetFundFromDb(fund *objects.Task) (bool, string) {

	return true, "success"
}

func GetFundsFromDb(google_id string) (bool, []models.Fund) {
	var funds []models.Fund
	result := db.Where("g_id = ?", google_id).Find(&funds)
	return result.Error == nil, funds
}

func UpdateFundInDb(fund *objects.Fund) (bool, string) {

	return true, "success"
}

func DeleteFundFromDb(fund *objects.Fund) (bool, string) {

	return true, "success"
}

// ######################## PAYMENTS SERVICES ###########################

func AddPaymentToDb(payment objects.Payment) bool {
	new_payment := models.Payment{
		GID:    payment.GID,
		Name:   payment.Name,
		Reason: payment.Reason,
		Amount: payment.Amount,
		Date:   payment.Date,
	}
	result := db.Create(&new_payment)
	return result.Error == nil
}

func GetPaymentFromDb(fund *objects.Payment) (bool, string) {

	return true, "success"
}

func GetPaymentsFromDb(google_id string) (bool, []models.Payment) {
	var payments []models.Payment
	result := db.Where("g_id = ?", google_id).Find(&payments)
	return result.Error == nil, payments
}

func UpdatePaymentInDb(fund *objects.Payment) (bool, string) {

	return true, "success"
}

func DeletePaymentFromDb(fund *objects.Payment) (bool, string) {

	return true, "success"
}
