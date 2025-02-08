/*
 * @Author: javohir-a abdusamatovjavohir@gmail.com
 * @Date: 2024-12-05 00:56:53
 * @LastEditors: javohir-a abdusamatovjavohir@gmail.com
 * @LastEditTime: 2024-12-18 22:52:54
 * @FilePath: /auth/internal/items/models/user.go
 * @Description: 这是默认设置,请设置`customMade`, 打开koroFileHeader查看配置 进行设置: https://github.com/OBKoro1/koro1FileHeader/wiki/%E9%85%8D%E7%BD%AE
 */
package models

import "time"

type Gender string

const (
	Male   Gender = "male"
	Female Gender = "female"
	Other  Gender = "other"
)

type User struct {
	ID          int       `json:"id"`
	ProID       int       `json:"pro_id"`
	FirstName   string    `json:"first_name"`
	LastName    string    `json:"last_name"`
	Email       string    `json:"email"`
	DateOfBirth string    `json:"date_of_birth"`
	Gender      Gender    `json:"gender"`
	Avatar      string    `json:"avatar"`
	Role        string    `json:"role"`
	Phone       string    `json:"phone"`
	CreatedAt   time.Time `json:"created_at"`
	UpdatedAt   time.Time `json:"updated_at"`
}

type CreateUser struct {
	ProID       int    `json:"pro_id" validate:"required"`
	FirstName   string `json:"first_name" validate:"required"`
	LastName    string `json:"last_name" validate:"required"`
	Email       string `json:"email"`
	Phone       string `json:"phone"`
	DateOfBirth string `json:"date_of_birth"`
	Gender      Gender `json:"gender" validate:"required,oneof=male female other"`
	Avatar      string `json:"avatar"`
	Role        string `json:"role"`
}

type UpdateUser struct {
	ID          int     `json:"id" validate:"required"`
	ProID       *int    `json:"pro_id"`
	FirstName   *string `json:"first_name"`
	LastName    *string `json:"last_name"`
	Email       *string `json:"email"`
	DateOfBirth *string `json:"date_of_birth"`
	Gender      *Gender `json:"gender"`
	Avatar      *string `json:"avatar"`
	Role        *string `json:"role"`
}

type BirthRange struct {
	StartDate time.Time `json:"start_date"`
	EndDate   time.Time `json:"end_date"`
}

func (br BirthRange) IsValid() bool {
	return br.StartDate.Before(br.EndDate)
}

type GetAllUsers struct {
	FirstName  *string     `json:"first_name"`
	LastName   *string     `json:"last_name"`
	BirthRange *BirthRange `json:"birth_range"`
	Gender     *Gender     `json:"gender"`
	Role       *string     `json:"role"`
	Limit      int         `json:"limit"`
	Offset     int         `json:"offset"`
}

type GetAllUsersRes struct {
	TotalItems int     `json:"total_items"`
	Users      []*User `json:"users"`
}
