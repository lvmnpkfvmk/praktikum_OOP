package investment

import "fmt"

// Класс для вкладов
type Deposit struct {
	InvestmentBase
    InterestRate    float64
    StartInterestRate float64
}

// Конструктор
func NewDeposit(name string, maxCount int, interestRate float64) *Deposit {
    return &Deposit{
		InvestmentBase: *NewInvestmentBase(name, maxCount),
        InterestRate:     interestRate,
        StartInterestRate: interestRate,
    }
}

// Расчет стоимости инвестиции
func (d *Deposit) InvestmentCost() int {
    return d.Count
}

// Покупка вклада
func (d *Deposit) Buy(count int) {
    d.Count += count
    d.InterestRate += d.StartInterestRate * float64(count) / float64(d.MaxCount) / 2
}

// Продажа вклада
func (d *Deposit) Sell(count int) {
    d.Count -= count
    d.InterestRate -= d.StartInterestRate * float64(count) / float64(d.MaxCount) / 2
}

// Расчет доходности инвестиции
func (d *Deposit) CalculateProfitability() int {
    fmt.Println("Доход", int(float64(d.Count) * d.InterestRate))
    return int(float64(d.Count) * d.InterestRate)
}

// Переопределение метода String()
func (d *Deposit) String() string {
    return fmt.Sprintf("%s (Максимум %d рублей, процентная ставка %.2f)", d.name, d.MaxCount-d.Count, d.InterestRate)
}
