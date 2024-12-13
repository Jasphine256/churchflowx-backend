package services

import (
	"churchflowx/internal/database"
	"churchflowx/internal/models"
	"churchflowx/internal/objects"
)

var db = database.InitialiseDB()

// ######################## FUNDS SERVICES ###########################

func AddFundToDb(user_id int, fund objects.Fund) bool {
	new_fund := models.Fund{
		UserId: user_id,
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

func GetFundsFromDb(fund *objects.Fund) (bool, string) {

	return true, "success"
}

func UpdateFundInDb(fund *objects.Fund) (bool, string) {

	return true, "success"
}

func DeleteFundFromDb(fund *objects.Fund) (bool, string) {

	return true, "success"
}

// ######################## PAYMENTS SERVICES ###########################

func AddPaymentToDb(user_id int, payment objects.Payment) bool {
	new_payment := models.Payment{
		UserId: user_id,
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

func GetPaymentsFromDb(fund *objects.Payment) (bool, string) {

	return true, "success"
}

func UpdatePaymentInDb(fund *objects.Payment) (bool, string) {

	return true, "success"
}

func DeletePaymentFromDb(fund *objects.Payment) (bool, string) {

	return true, "success"
}
