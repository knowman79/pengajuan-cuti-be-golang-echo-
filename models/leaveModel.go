package models

import "time"

type LeaveModel struct {
	FormId           int       `json:"form_id" validate:"required" sql:"AUTO_INCREMENT"`
	UserId           int       `json:"user_id" validate:"required"`
	Name             string    `json:"name" validate:"required"`
	Types            string    `json:"type" validate:"required"`
	CreatedBy        string    `json:"created_by" validate:"required"`
	ModifiedBy       string    `json:"modified_by"`
	CreatedDate      time.Time `json:"created_date" validate:"required"`
	LastModifiedDate time.Time `json:"last_modified_date"`
	StartDate        time.Time `json:"start_date" validate:"required"`
	EndDate          time.Time `json:"end_date" validate:"required"`
	Description      string    `json:"descpription" validate:"required"`
	ReplacementId    int       `json:"replacement_id" validate:"required"`
	Address          string    `json:"address" validate:"required"`
	Phone            string    `json:"phone" validate:"required"`
	Status           string    `json:"status" validate:"required"`
}
