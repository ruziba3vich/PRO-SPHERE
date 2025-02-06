/*
 * @Author: javohir-a abdusamatovjavohir@gmail.com
 * @Date: 2024-12-05 04:02:50
 * @LastEditors: javohir-a abdusamatovjavohir@gmail.com
 * @LastEditTime: 2024-12-18 22:55:15
 * @FilePath: /auth/internal/items/models/proid.go
 * @Description: 这是默认设置,请设置`customMade`, 打开koroFileHeader查看配置 进行设置: https://github.com/OBKoro1/koro1FileHeader/wiki/%E9%85%8D%E7%BD%AE
 */
package models

import (
	"encoding/json"
	"fmt"
	"time"
)

type ProIDUser struct {
	ID          int    `json:"id"`
	FirstName   string `json:"name"`
	LastName    string `json:"surname"`
	PhoneNumber string `json:"phone_number"`
	Email       string `json:"email"`
	DateOfBirth string `json:"date_of_birth"` // Updated type
	Gender      Gender `json:"gender"`
	AvatarUrl   string `json:"avatar_full_url"`
	Avatar      string `json:"avatar"`
}

type Date time.Time

const dateLayout = "2006-01-02"

// UnmarshalJSON parses a date in the "YYYY-MM-DD" format
func (d *Date) UnmarshalJSON(b []byte) error {
	// Remove quotes from the JSON string
	str := string(b)
	str = str[1 : len(str)-1] // Strip quotes

	// Parse the date
	t, err := time.Parse(dateLayout, str)
	if err != nil {
		return fmt.Errorf("invalid date format: %s", err)
	}
	*d = Date(t)
	return nil
}

// MarshalJSON converts the Date back to JSON
func (d Date) MarshalJSON() ([]byte, error) {
	return json.Marshal(time.Time(d).Format(dateLayout))
}

// ToTime converts the custom Date type back to time.Time
func (d Date) ToTime() time.Time {
	return time.Time(d)
}

func NormalizeGender(gender string) string {
	switch gender {
	case "Male":
		return "male"
	case "Female":
		return "female"
	case "":
		return "none"
	default:
		return ""
	}
}
