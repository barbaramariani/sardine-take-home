package main

import (
	"encoding/json"
	"fmt"
	"io/ioutil"
	"net/http"
	"slices"
)

func risk(w http.ResponseWriter, req *http.Request) {
	body, err := ioutil.ReadAll(req.Body)
	if err != nil {
		fmt.Printf("could not read body: %s\n", err)
	}

	rules := initializeRules()
	result := evaluateRules(body, rules)

	responseJSON, err := json.Marshal(map[string]interface{}{"risk_ratings": result})
	if err != nil {
		fmt.Printf("could not marshal response: %s\n", err)
		http.Error(w, "Internal Server Error", http.StatusInternalServerError)
		return
	}

	w.Header().Set("Content-Type", "application/json")

	w.Write(responseJSON)
}

func evaluateRules(transactionsBytes []byte, rules []Rule) RiskRatings {
	var transactions Transactions
	err := json.Unmarshal(transactionsBytes, &transactions)
	if err != nil {
		fmt.Printf("could not unmarshal transactions: %s\n", err)
		return nil
	}

	var riskRatings RiskRatings

	for idx, transaction := range transactions.Transactions {
		user := User{Id: transaction.UserId, TotalSpent: 0, CardIds: []int{}}
		for i := 0; i <= idx; i++ {
			if transactions.Transactions[i].UserId == user.Id {
				user.TotalSpent += transactions.Transactions[i].AmountUsCents
				if !slices.Contains(user.CardIds, transactions.Transactions[i].CardId) {
					user.CardIds = append(user.CardIds, transactions.Transactions[i].CardId)
				}
			}
		}
		highestRisk := low
		for _, rule := range rules {
			if rule.Condition(transaction, user) && rule.Risk > highestRisk {
				highestRisk = rule.Risk
			}
		}
		riskRatings = append(riskRatings, highestRisk.String())
	}
	return riskRatings
}

func initializeRules() []Rule {
	return []Rule{
		{
			Id: 1,
			Condition: func(transaction Transaction, _ User) bool {
				return transaction.AmountUsCents > 5000_00
			},
			Risk: medium,
		},
		{
			Id: 2,
			Condition: func(transaction Transaction, _ User) bool {
				return transaction.AmountUsCents > 10000_00
			},
			Risk: high,
		},
		{
			Id: 3,
			Condition: func(_ Transaction, user User) bool {
				return user.TotalSpent > 10000_00
			},
			Risk: medium,
		},
		{
			Id: 4,
			Condition: func(_ Transaction, user User) bool {
				return user.TotalSpent > 20000_00
			},
			Risk: high,
		},
		{
			Id: 5,
			Condition: func(_ Transaction, user User) bool {
				return len(user.CardIds) > 1
			},
			Risk: medium,
		},
		{
			Id: 6,
			Condition: func(_ Transaction, user User) bool {
				return len(user.CardIds) > 2
			},
			Risk: high,
		},
	}
}
