package models

import "time"

type TotalPoints struct {
	Points int `json:"points"`
}

type StaticPoint struct {
	Company          Company   `json:"company"`
	Points           int       `json:"points"`
	CreatedYearMonth time.Time `json:"created_year_month"`
}

type QrmarkInfo struct {
	QrmarkID  int
	UserID    int
	CompanyID int
	Point     int
}

type School struct {
	ID        int       `json:"school_id"`
	Name      string    `json:"name"`
	CreatedAt time.Time `json:"created_at"`
}

type VerificationToken struct {
	UserID    int
	Token     string
	ExpiredAt time.Time
}

type User struct {
	ID        int       `json:"user_id"`
	Name      string    `json:"name"`
	Email     string    `json:"email"`
	Password  string    `json:"password"`
	Role      string    `json:"role"`
	Verified  bool      `json:"verified"`
	SchoolID  int       `json:"school_id"`
	CreatedAt time.Time `json:"created_at"`
}

type UserRes struct {
	ID        int       `json:"user_id"`
	Name      string    `json:"name"`
	Email     string    `json:"email"`
	Role      string    `json:"role"`
	Verified  bool      `json:"verified"`
	School    School    `json:"school"`
	CreatedAt time.Time `json:"created_at"`
}

type Company struct {
	ID        int       `json:"company_id"`
	Name      string    `json:"name"`
	CreatedAt time.Time `json:"created_at"`
}

type Qrmark struct {
	QrmarkID    int       `json:"qrmark_id"`
	UserID      int       `json:"user_id"`
	SchoolName  string    `json:"school_name"`
	CompanyName string    `json:"company_name"`
	Points      int       `json:"points"`
	CreatedAt   time.Time `json:"created_at"`
}

type QrmarkList struct {
	QrmarkList []Qrmark `json:"qrmarks"`
	HasNext    bool     `json:"has_next"`
	Page       int      `json:"page"`
}

type SchoolList struct {
	SchoolList []School `json:"schools"`
	HasNext    bool     `json:"has_next"`
	Page       int      `json:"page"`
}

type UserList struct {
	UserList []UserRes `json:"users"`
	HasNext  bool      `json:"has_next"`
	Page     int       `json:"page"`
}
