package main

type Transactions struct {
	Transactions []Transaction `json:"transactions"`
}

type Transaction struct {
	Id            int `json:"id"`
	UserId        int `json:"user_id"`
	AmountUsCents int `json:"amount_us_cents"`
	CardId        int `json:"card_id"`
}

type User struct {
	Id         int
	TotalSpent int
	CardIds    []int
}

type Condition func(Transaction, User) bool

type Rule struct {
	Id        int
	Condition Condition
	Risk      Risk
}

type Risk int

const (
	low Risk = iota
	medium
	high
)

func (r Risk) String() string {
	return [...]string{"low", "medium", "high"}[r]
}

type RiskRatings []string
