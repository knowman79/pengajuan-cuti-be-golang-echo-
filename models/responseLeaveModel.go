package models

import "time"

type ResponseLeaveModel struct {
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
    LeaveDuration    int       `json:"leave_duration" validate:"required"`
    Description      string    `json:"descpription" validate:"required"`
    ReplacementId    int       `json:"replacement_id" validate:"required"`
    ReplacementName  string    `json:"replacement_name" validate:"required"`
    Address          string    `json:"address" validate:"required"`
    Phone            string    `json:"phone" validate:"required"`
    Status           string    `json:"status" validate:"required"`
    LeaveId          int       `json:"leave_id" validate:"required"`
    Current_leave    int       `json:"current_leave" validate:"required"`
}
