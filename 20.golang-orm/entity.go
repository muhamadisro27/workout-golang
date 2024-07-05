package golangorm

import (
	"database/sql"
	"time"

	"gorm.io/gorm"
)

type Sample struct {
	ID   string `json:"id"`
	Name string `json:"name"`
}

type UserLog struct {
	ID        uint32 `json:"id" gorm:"primaryKey;column:id;autoIncrement;<-:create"`
	UserID    uint32 `json:"user_id" gorm:"column:user_id"`
	Action    string `json:"action" gorm:"column:action"`
	CreatedAt int64  `json:"created_at" gorm:"column:created_at;autoCreateTime:milli;<-:create"`
	UpdatedAt int64  `json:"updated_at" gorm:"column:updated_at;autoCreateTime:milli;autoUpdateTime:milli"`
}

type Todo struct {
	gorm.Model
	UserID      uint32 `json:"user_id" gorm:"column:user_id"`
	Title       string `json:"title" gorm:"column:title"`
	Description string `json:"description" gorm:"column:description"`
}

type User struct {
	ID           string    `json:"id" gorm:"primaryKey;column:id;<-:create"`
	Name         Name      `gorm:"embedded" json:"name"`
	Password     string    `json:"password" gorm:"column:password"`
	CreatedAt    time.Time `json:"created_at" gorm:"column:created_at;<-:create"`
	UpdatedAt    time.Time `json:"updated_at" gorm:"column:updated_at"`
	Information  string    `gorm:"-"`
	Wallet       Wallet    `gorm:"foreignKey:user_id;references:id" json:"wallet"`
	Addresses    []Address `gorm:"foreignKey:user_id;references:id" json:"addresses"`
	LikeProducts []Product `gorm:"many2many:user_likes_product;foreignKey:id;joinForeignKey:user_id;references:id;joinReferences:product_id" json:"like_products"`
}

type Name struct {
	FirstName  string `json:"first_name"`
	MiddleName string `json:"middle_name"`
	LastName   string `json:"last_name"`
}

type Address struct {
	ID        string    `gorm:"primary_key;column:id;autoIncrement" json:"id"`
	UserID    string    `gorm:"column:user_id" json:"user_id"`
	Address   string    `gorm:"column:address" json:"address"`
	CreatedAt time.Time `gorm:"column:created_at;autoCreateTime" json:"created_at"`
	UpdatedAt time.Time `gorm:"column:updated_at;autoCreateTime;autoUpdateTime" json:"updated_at"`
	User      User      `gorm:"foreignKey:user_id;references:id" json:"user"`
}

type Wallet struct {
	ID        string    `gorm:"primary_key;column:id" json:"id"`
	UserID    string    `gorm:"column:user_id" json:"user_id"`
	Balance   int64     `gorm:"column:balance" json:"balance"`
	CreatedAt time.Time `gorm:"column:created_at;autoCreateTime" json:"created_at"`
	UpdatedAt time.Time `gorm:"column:updated_at;autoCreateTime;autoUpdateTime" json:"updated_at"`
	User      *User     `gorm:"foreignKey:user_id;references:id" json:"user"`
}

type Product struct {
	ID    string `gorm:"primary_key;column:id" json:"id"`
	Name  string `gorm:"column:name" json:"name"`
	Price int64  `gorm:"column:price" json:"price"`
	CreatedUpdatedAt
	LikedByUsers []User `gorm:"many2many:user_likes_product;foreignKey:id;joinForeignKey:product_id;references:id;joinReferences:user_id" json:"liked_by_users"`
}

type CreatedUpdatedAt struct {
	CreatedAt time.Time `gorm:"column:created_at;autoCreateTime" json:"created_at"`
	UpdatedAt time.Time `gorm:"column:updated_at;autoCreateTime;autoUpdateTime" json:"updated_at"`
}

type GuestBook struct {
	ID      int64  `gorm:"primary_key;column:id;autoIncrement"`
	Name    string `gorm:"column:name"`
	Email   string `gorm:"column:email"`
	Message string `gorm:"column:message"`
	CreatedUpdatedAt
}

func (g *GuestBook) TableName() string {
	return "guest_books"
}

func (p *Product) TableName() string {
	return "products"
}

func (a *Address) TableName() string {
	return "addresses"
}

func (w *Wallet) TableName() string {
	return "wallets"
}

func (ul *UserLog) TableName() string {
	return "user_logs"
}

func (u *User) TableName() string {
	return "users"
}

func (u *User) BeforeCreate(db *gorm.DB) error {
	if u.ID == "" {
		u.ID = "user-before-create"
	}
	return nil
}
func (t *Todo) TableName() string {
	return "todos"
}

func SetValueInNullColumn(value string) *sql.NullString {

	return &sql.NullString{value, true}

}
