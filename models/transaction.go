package models

type Transaction struct {
	ID      uint              `gorm:"primary_key" json:"id"`
	UserID  uint              `json:"user_id"`
	Amount  float64           `json:"amount"`
	Items   []TransactionItem `gorm:"foreignkey:TransactionID" json:"items"`
	Pricing float64           `json:"pricing"`
}
