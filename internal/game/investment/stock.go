package investment

import "fmt"

type IStock interface {
}

// Класс для акций
type Stock struct {
    InvestmentBase
    PricePerShare  int
    StartPricePerShare int
}

// Конструктор
func NewStock(name string, maxCount, pricePerShare int) *Stock {
    return &Stock{
        InvestmentBase: *NewInvestmentBase(name, maxCount),
        PricePerShare:  pricePerShare,
        StartPricePerShare: pricePerShare,
    }
}

// Расчет стоимости инвестиции
func (s *Stock) InvestmentCost() int {
    return s.PricePerShare
}

// Покупка акций
func (s *Stock) Buy(count int) {
    s.Count += count
    s.PricePerShare += s.StartPricePerShare * count / s.MaxCount / 2
}

// Продажа акций
func (s *Stock) Sell(count int) {
    s.Count -= count
    s.PricePerShare -= s.StartPricePerShare * count / s.MaxCount / 2
}

// Расчет доходности инвестиции
func (s *Stock) CalculateProfitability() int {
    return 0
}

// Переопределение метода String()
func (s *Stock) String() string {
    return fmt.Sprintf("%s (имеется %d штук, цена за штуку %d)", s.name, s.MaxCount-s.Count, s.PricePerShare)
}
