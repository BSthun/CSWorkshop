package enroll

import (
	"backend/modules"
	ihub "backend/modules/hub"
	"backend/types/model"
	"backend/types/payload"
	"backend/types/response"
	"backend/utils/value"
)

func MapEnrollment(enrollment *model.Enrollment) *payload.EnrollmentInfo {
	return &payload.EnrollmentInfo{
		EnrollmentId: enrollment.Id,
		EnrolledAt:   enrollment.CreatedAt,
		LabName:      enrollment.Lab.Name,
		DbName:       enrollment.DbName,
		DbValid:      enrollment.DbValid,
		DbHost:       nil,
		DbPort:       nil,
		DbUsername:   nil,
		DbPassword:   nil,
		Tasks:        nil,
	}
}

func MapEnrollmentTask(enrollment *model.Enrollment, tasks []*model.Task, session *ihub.Session) *payload.EnrollmentInfo {
	mappedEnrollment := MapEnrollment(enrollment)

	if enrollment.User != nil {
		mappedEnrollment.DbHost = &modules.Conf.InfoDbHost
		mappedEnrollment.DbPort = &modules.Conf.InfoDbPort
		mappedEnrollment.DbUsername = enrollment.User.Credential.Username
		mappedEnrollment.DbPassword = enrollment.User.Credential.Password
	}

	mappedEnrollment.Tasks, _ = value.Iterate(tasks, func(task *model.Task) (*payload.TaskList, *response.ErrorInstance) {
		return &payload.TaskList{
			Id:    task.Id,
			Title: task.Title,
		}, nil
	})

	if session != nil {
		mappedEnrollment.Token = session.Token
	}

	return mappedEnrollment
}
