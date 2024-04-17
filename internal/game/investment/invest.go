package investment

// Базовый интерфейс для инвестиций
type Investment interface {
	Name() string
    InvestmentCost() int
    GetMaxCount() int
    GetStartCount() int
    Buy(int)
    Sell(int)
    CalculateProfitability() int
    String() string
}

// Базовая реализация инвестиций
type InvestmentBase struct {
    name     string
    Count    int
    MaxCount int
    StartCount int
}

// Конструктор
func NewInvestmentBase(name string, maxCount int) *InvestmentBase {
    return &InvestmentBase{
        name:     name,
        MaxCount: maxCount,
        StartCount: maxCount,
    }
}
func (i *InvestmentBase) Name() string {
    return i.name
}
// Реализация методов интерфейса Investment
func (i *InvestmentBase) GetMaxCount() int {
    return i.MaxCount - i.Count
}

func (i *InvestmentBase) GetStartCount() int {
    return i.Count
}
