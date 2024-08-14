package handler

import (
	"fmt"
	"time"
)

func errParamsisRequired(name, typ string) error {
	return fmt.Errorf("param: %s (type: %s) is required", name, typ)
}

//create scheduling

type CreateSchedulingRequest struct {
	Description string    `json:"description"`
	StartTime   time.Time `json:"starttime"`
	BarberID    uint      `json:"barber-id"`
	ClientID    uint      `json:"client-id"`
	EndTime     time.Time `json:"endtime"`
	Status      string    `json:"status"`
}

// show schedulings
type ShowSchedulingsRequest struct {
	Day time.Time `json:"day"`
}

// create barber
type CreateBarberRequest struct {
	Name  string `json:"name"`
	Email string `json:"email"`
	Phone string `json:"phone"`
}

func (r *ShowSchedulingsRequest) Validate() error {
	if r.Day.IsZero() == true {
		return errParamsisRequired("day", "time")
	}
	return nil
}

func (r *CreateSchedulingRequest) Validate() error {
	if r.EndTime.IsZero() == true && r.Status == "" && r.StartTime.IsZero() == true && r.Description == "" {
		return fmt.Errorf("request body is empty or malformed")
	}

	if r.Description == "" {
		return errParamsisRequired("description", "string")
	}
	if r.Status == "" {
		return errParamsisRequired("status", "string")
	}
	if r.StartTime.IsZero() == true {
		return errParamsisRequired("starttime", "time")
	}
	if r.EndTime.IsZero() == true {
		return errParamsisRequired("endtime", "time")
	}
	return nil
}
