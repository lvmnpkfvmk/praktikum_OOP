package fund

import (
    "fmt"
	cond "investgo/internal/game/conditions"
	in "investgo/internal/game/investment"
	u "investgo/internal/utils"
)

// Класс для инвестиционного фонда
type InvestmentFund struct {
	Capital     int
	MyPortfolio *Portfolio
	Taxes       float64
}

// Конструктор
func NewInvestmentFund(initialCapital int) *InvestmentFund {
	return &InvestmentFund{
		Capital:     initialCapital,
		MyPortfolio: NewPortfolio(),
		Taxes:       0.13,
	}
}

// Расчет доходности инвестиций
func (f *InvestmentFund) CalculateInvestmentReturns() {
	portfolioProfitability := f.MyPortfolio.CalculatePortfolioProfitability()
	f.Capital += portfolioProfitability

	fmt.Printf("Доходность портфеля: %d\n", portfolioProfitability)
	fmt.Printf("Капитал фонда: %d\n", f.Capital)
}

// Оплата налогов
func (f *InvestmentFund) PayTaxes() {
	taxes := int(float64(cond.GetIncome(f.Capital)) * f.Taxes)
	f.Capital -= taxes
	fmt.Printf("Налоги: %d\n", taxes)
	fmt.Printf("Капитал фонда: %d\n", f.Capital)
}

// Учет новых поступлений и расходов
func (f *InvestmentFund) AccountForNewIncomeAndExpenses() {
	newIncome := cond.GetIncome(f.Capital)
	f.Capital += newIncome
	fmt.Printf("Новые поступления: %d\n", newIncome)

	newExpenses := cond.GetExpenses(f.Capital)
	f.Capital -= newExpenses
	fmt.Printf("Новые расходы: %d\n", newExpenses)

	fmt.Printf("Капитал фонда: %d\n", f.Capital)
}

// Покупка инвестиции
func (f *InvestmentFund) BuyInvestment(investment in.Investment) {
	maxBuyCount := f.Capital / investment.InvestmentCost()
	if maxBuyCount > investment.GetMaxCount() {
		maxBuyCount = investment.GetMaxCount()
	}

	fmt.Println(f.Capital, investment.InvestmentCost(), maxBuyCount)
	if maxBuyCount <= 0 {
		fmt.Printf("Недостаточно средств для покупки %s\n", f.getInvestmentType(investment))
		return
	}

	investmentType := f.getInvestmentType(investment)
	fmt.Printf("Доступно %d. Введите количество покупаемого %s:\n", maxBuyCount, investmentType)
	count := u.ChoiceNumber(maxBuyCount, true)
	f.MyPortfolio.BuyInvestment(investment, count)
	f.Capital -= investment.InvestmentCost() * count
	fmt.Printf("Капитал фонда: %d\n", f.Capital)
}

// Продажа инвестиции
func (f *InvestmentFund) SellInvestment(investment in.Investment) {
	maxSellCount := investment.GetMaxCount()
	if maxSellCount == 0 {
		fmt.Println("Нет инвестиций для продажи")
		return
	}

	fmt.Printf("Введите количество продаваемых инвестиций (максимум %d):\n", maxSellCount)
	count := u.ChoiceNumber(maxSellCount, true)
	f.MyPortfolio.SellInvestment(investment, count)
	f.Capital += investment.InvestmentCost() * count
}

func (f *InvestmentFund) getInvestmentType(investment in.Investment) string {
	switch investment.(type) {
	case *in.Stock:
		return "акции"
	case *in.Bond:
		return "облигации"
	case *in.Metal:
		return "металла"
	case *in.Deposit:
		return "вклада"
	default:
		return ""
	}
}
