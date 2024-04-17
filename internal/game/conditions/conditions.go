package conditions

import (
	"fmt"
	"math/rand"
	"time"
	"investgo/internal/game/investment"
)

var (
	r = rand.New(rand.NewSource(time.Now().UnixNano()))
)

// Обработка событий
func EventProcessing(investments []investment.Investment) {
	investment := investments[r.Intn(len(investments))]
	count := r.Intn(investment.GetMaxCount()+investment.GetStartCount()) - investment.GetStartCount()

	if count == 0 {
		return
	}

	if count > 0 {
		investment.Buy(min(count, investment.GetMaxCount()))
	} else {
		if investment.GetStartCount() > -count {
			investment.Sell(-count)
		}
	}
}

// Обновление инвестиций
func UpdateInvestment(investments []investment.Investment) {
	eventNumber := r.Intn(80) + 20
	for i := 0; i < eventNumber; i++ {
		EventProcessing(investments)
	}

	fmt.Println("Возможные инвестиции:")
	for i, investment := range investments {
		fmt.Printf("%d. %s\n", i+1, investment)
	}
}

// Получение дохода
func GetIncome(capital int) int {
	return r.Intn(capital / 50)
}

// Получение расходов
func GetExpenses(capital int) int {
	return r.Intn(capital / 50)
}

func min(a, b int) int {
	if a < b {
		return a
	}
	return b
}
