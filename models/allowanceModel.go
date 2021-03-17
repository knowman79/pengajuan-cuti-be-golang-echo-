package models

type AllowanceModel struct {
	LeaveId       int `json:"leave_id" validate:"required" sql:"AUTO_INCREMENT"`
	UserId        int `json:"user_id" validate:"required"`
	CurrentLeave  int `json:"current_leave" validate:"required"`
	LastYearLeave int `json:"last_year_leave" validate:"required"`
}
