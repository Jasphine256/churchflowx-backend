package services

import (
	"churchflowx/internal/models"
	"churchflowx/internal/objects"
)

// ######################## TASKS SERVICES ###########################

func AddTaskToDb(google_id int, user_task objects.Task) bool {

	new_task := models.Task{
		GID:         google_id,
		Title:       user_task.Title,
		Description: user_task.Description,
		StartDate:   user_task.StartDate,
		DateDue:     user_task.DateDue,
		Status:      user_task.Status,
	}
	result := db.Create(&new_task)
	return result.Error == nil
}

func GetTaskFromDb(task *objects.Task) (bool, string) {
	return true, "success"
}

func GetTasksFromDb(google_id int) (bool, []models.Task) {
	var tasks []models.Task
	result := db.Where("g_id = ?", google_id).Find(&tasks)
	return result.Error == nil, tasks
}

func UpdateTaskInDb(task *objects.Task) (bool, string) {
	return true, "success"
}

func DeleteTaskFromDb(task *objects.Task) (bool, string) {
	return true, "success"
}

// ######################## PLANS SERVICES ###########################

func AddPlanToDb(google_id int, plan objects.Plan) bool {
	new_plan := models.Plan{
		GID:         google_id,
		Title:       plan.Title,
		Description: plan.Description,
		Handler:     plan.Handler,
		Team:        plan.Team,
		Budget:      plan.Budget,
	}
	result := db.Create(&new_plan)
	return result.Error == nil
}

func GetPlanFromDb(plan *objects.Plan) (bool, string) {
	return true, "success"
}

func GetPlansFromDb(google_id int) (bool, []models.Plan) {
	var plans []models.Plan
	result := db.Where("g_id = ?", google_id).Find(&plans)
	return result.Error == nil, plans
}

func UpdatePlanInDb(plan *objects.Plan) (bool, string) {
	return true, "success"
}

func DeletePlanFromDb(plan *objects.Plan) (bool, string) {
	return true, "success"
}

// ######################## PROJECTS SERVICES ###########################

func AddProjectToDb(google_id int, project *objects.Project) bool {
	new_project := models.Project{
		GID:         google_id,
		Title:       project.Title,
		Description: project.Description,
		Handler:     project.Description,
		Budget:      project.Budget,
		StartDate:   project.StartDate,
		EndDate:     project.EndDate,
	}
	result := db.Create(&new_project)
	return result.Error == nil
}

func GetProjectFromDb(project *objects.Project) (bool, string) {
	return true, "success"
}

func GetProjectsFromDb(google_id int) (bool, []models.Project) {
	var projects []models.Project
	result := db.Where("g_id = ?", google_id).Find(&projects)
	return result.Error == nil, projects
}

func UpdateProjectInDb(project *objects.Project) (bool, string) {
	return true, "success"
}

func DeleteProjectFromDb(project *objects.Project) (bool, string) {
	return true, "success"
}
