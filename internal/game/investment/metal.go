package investment

import "fmt"

// Класс для металлов
type Metal struct {
    InvestmentBase
    Weight         int
    PricePerGram   float64
    StartPricePerGram float64
}

// Конструктор
func NewMetal(name string, maxCount, weight int, pricePerGram float64) *Metal {
    return &Metal{
        InvestmentBase: *NewInvestmentBase(name, maxCount),
        Weight:         weight,
        PricePerGram:   pricePerGram,
        StartPricePerGram: pricePerGram,
    }
}

// Расчет стоимости инвестиции
func (m *Metal) InvestmentCost() int {
    return int(float64(m.Weight) * m.PricePerGram)
}

// Покупка металла
func (m *Metal) Buy(count int) {
    m.Count += count
    m.PricePerGram += m.StartPricePerGram * float64(count) / float64(m.MaxCount) / 2
}

// Продажа металла
func (m *Metal) Sell(count int) {
    m.Count -= count
    m.PricePerGram -= m.StartPricePerGram * float64(count) / float64(m.MaxCount) / 2
}

// Расчет доходности инвестиции
func (m *Metal) CalculateProfitability() int {
    return 0
}

// Переопределение метода String()
func (m *Metal) String() string {
    return fmt.Sprintf("%s (имеется %d штук, вес одной штуки %d, цена за грамм %.2f)", m.name, m.MaxCount-m.Count, m.Weight, m.PricePerGram)
}
