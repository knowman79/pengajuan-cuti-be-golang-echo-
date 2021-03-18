package models

import "time"

type LeaveByNameModel struct {
    CreatedDate time.Time `json:"created_date" validate:"required"`
    Description string    `json:"descpription" validate:"required"`
    StartDate   time.Time `json:"start_date" validate:"required"`
    EndDate     time.Time `json:"end_date" validate:"required"`
    Duration    int       `json:"duration" validate:"required"`
    Status      string    `json:"status" validate:"required"`
}