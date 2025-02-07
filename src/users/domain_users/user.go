package domain

type User struct {
	ID        int32  `gorm:"primaryKey;autoIncrement" json:"id"`
	Name      string `json:"name"`
	Email     string `json:"email"`
	Password  string `json:"password"`
	UpdatedAt int64  `json:"updated_at"`
}

func NewUser(name string, email string, password string) *User {
	return &User{Name: name, Email: email, Password: password}
}
