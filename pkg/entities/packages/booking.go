package packages

// import (
// 	"gorm.io/gorm"
// )

// type UserData struct {
// 	gorm.Model
// 	Email    string
// 	Phone    string
// 	Password string
// 	Name     string
// }

// type Traveller struct {
// 	gorm.Model
// 	Name       string
// 	Age        string
// 	Gender     string
// 	UserId     uint
// 	PackageId  uint
// 	Package    Package `gorm:"foreignKey:PackageId"`
// 	ActivityId JSONB
// }

// type Booking struct {
// 	gorm.Model
// 	PaymentId        string `gorm:"unique"`
// 	BookingReference string `gorm:"unique;not null"`
// 	BookingStatus    string `gorm:"default:PENDING"`
// 	TotalFare        string
// 	CancelledStatus  string      `gorm:"default:false"`
// 	UserId           uint        `gorm:"foreignKey:"`
// 	User             UserData    `gorm:"foreignKey:UserId"`
// 	Bookings         []Traveller `gorm:"many2many:traveller_booking;"`
// }
