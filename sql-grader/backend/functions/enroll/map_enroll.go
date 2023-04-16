package enroll

import (
	"backend/modules"
	"backend/types/model"
	"backend/types/payload"
	"backend/types/response"
	"backend/utils/value"
)

func MapEnrollment(enrollment *model.Enrollment) *payload.EnrollInfo {
	return &payload.EnrollInfo{
		EnrollmentId: enrollment.Id,
		DbName:       enrollment.DbName,
		DbValid:      enrollment.DbValid,
	}
}

func MapEnrollmentTask(enrollment *model.Enrollment, tasks []*model.Task) *payload.EnrollInfo {
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

	return mappedEnrollment
}
