package fund

import (
	"fmt"
	in "investgo/internal/game/investment"
)

type Pair struct {
	I in.Investment
	Q int
}

// Класс для портфеля
type Portfolio struct {
    Deposits []Pair
    Metals   []Pair
    Bonds    []Pair
    Stocks   []Pair
}

// Конструктор
func NewPortfolio() *Portfolio {
    return &Portfolio{
        Deposits: make([]Pair, 0),
        Metals:   make([]Pair, 0),
        Bonds:    make([]Pair, 0),
        Stocks:   make([]Pair, 0),
    }
}

// Расчет доходности портфеля
func (p *Portfolio) CalculatePortfolioProfitability() int {
    var totalProfitability int

    // Расчет доходности для депозитов
    totalProfitability += p.calculateInvestmentTypeProfitability(p.Deposits)

    // Расчет доходности для металлов
    totalProfitability += p.calculateInvestmentTypeProfitability(p.Metals)

    // Расчет доходности для облигаций
    totalProfitability += p.calculateInvestmentTypeProfitability(p.Bonds)

    // Расчет доходности для акций
    totalProfitability += p.calculateInvestmentTypeProfitability(p.Stocks)

    return totalProfitability
}

// Расчет доходности любого типа инвестиций
func (p *Portfolio) calculateInvestmentTypeProfitability(investments []Pair) int {
    var typeProfitability int
    for _, investment := range investments {
        fmt.Println(typeProfitability)
        typeProfitability += investment.I.CalculateProfitability()
    }
    return typeProfitability
}

// Покупка инвестиций
func (p *Portfolio) BuyInvestment(investment in.Investment, quantity int) {
    switch inv := investment.(type) {
    case *in.Deposit:
        p.buyInvestment(&p.Deposits, inv, quantity)
    case *in.Metal:
        p.buyInvestment(&p.Metals, inv, quantity)
    case *in.Bond:
        p.buyInvestment(&p.Bonds, inv, quantity)
    case *in.Stock:
        p.buyInvestment(&p.Stocks, inv, quantity)
    }
    fmt.Println(p.Deposits)
    investment.Buy(quantity)
}

// Продажа инвестиций
func (p *Portfolio) SellInvestment(investment in.Investment, quantity int) {
    switch inv := investment.(type) {
    case *in.Deposit:
        p.sellInvestment(p.Deposits, inv, quantity)
    case *in.Metal:
        p.sellInvestment(p.Metals, inv, quantity)
    case *in.Bond:
        p.sellInvestment(p.Bonds, inv, quantity)
    case *in.Stock:
        p.sellInvestment(p.Stocks, inv, quantity)
    }
    investment.Sell(quantity)
}

// Вспомогательные методы для покупки/продажи инвестиций
func (p *Portfolio) buyInvestment(investments *[]Pair, investment in.Investment, quantity int) {
    for i, tuple := range *investments {
        if tuple.I.Name() == investment.Name() {
            (*investments)[i].Q += quantity
            return
        }
    }
    *investments = append(*investments, Pair{investment, quantity})
}

func (p *Portfolio) sellInvestment(investments []Pair, investment in.Investment, quantity int) {
    for i, tuple := range investments {
        if tuple.I.Name() == investment.Name() {
            investments[i].Q -= quantity
            if investments[i].Q <= 0 {
                investments = append(investments[:i], investments[i+1:]...)
            }
            return
        }
    }
}
