package types

// Money представляет собой денежную сумму в минимальных единицах (центы, копейки, дирамы и т.д).
type Money int64

// Category представляет собой категорию, в которой был совершен платёж (авто, аптеки, рестараны и т.д).
type Category string

// Status представляет собой статус платежа.
type Status string

// Предопределенные статусы платежей.
const (
	StatusOk       Status = "OK"
	StatusFail     Status = "FAIL"
	StatusProgress Status = "INPROGRESS"
)

// Payment представляет информацию о платеже.
type Payment struct {
	ID       int
	Amount   Money
	Category Category
	Status   Status
}
