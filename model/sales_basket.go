package model

// PaymentMethod defines the payment method type
type PaymentMethod string

const (
	PaymentMethodCash   PaymentMethod = "CASH"
	PaymentMethodCredit PaymentMethod = "CREDIT"
	PaymentMethodDebit  PaymentMethod = "DEBIT"
)

// SalesBasket represents the sales_basket table in the database
type SalesBasket struct {
	ID            int           `json:"id_sales" db:"id_sales"`
	SalesDate     int           `json:"sales_date" db:"sales_date"` // This might need to be a time.Time depending on actual usage
	UserID        int           `json:"id_user" db:"id_user"`
	MemberID      int           `json:"id_member" db:"id_member"`
	PaymentMethod PaymentMethod `json:"payment_method" db:"payment_method"`
	Total         int           `json:"total" db:"total"`
	
	// Optional relation fields (not in database)
	User          *User         `json:"user,omitempty" db:"-"`
	Member        *Member       `json:"member,omitempty" db:"-"`
	Items         []SalesItem   `json:"items,omitempty" db:"-"`
}
