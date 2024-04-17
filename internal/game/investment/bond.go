package investment

import "fmt"

// Класс для облигаций
type Bond struct {
	InvestmentBase
    FaceValue      int
    CouponRate     float64
    StartCouponRate float64
}

// Конструктор
func NewBond(name string, maxCount, faceValue int, couponRate float64) *Bond {
    return &Bond{
		InvestmentBase: *NewInvestmentBase(name, maxCount),
        FaceValue:      faceValue,
        CouponRate:     couponRate,
        StartCouponRate: couponRate,
    }
}

// Расчет стоимости инвестиции
func (b *Bond) InvestmentCost() int {
    return b.FaceValue
}

// Покупка облигаций
func (b *Bond) Buy(count int) {
    b.Count += count
    b.CouponRate += b.StartCouponRate * float64(count) / float64(b.MaxCount) / 2
}

// Продажа облигаций
func (b *Bond) Sell(count int) {
    b.Count -= count
    b.CouponRate -= b.StartCouponRate * float64(count) / float64(b.MaxCount) / 2
}

// Расчет доходности инвестиции
func (b *Bond) CalculateProfitability() int {
    return int(float64(b.Count) * float64(b.FaceValue) * b.CouponRate)
}

// Переопределение метода String()
func (b *Bond) String() string {
    return fmt.Sprintf("%s (имеется %d штук, номинал %d, купонная ставка %.2f)", b.name, b.MaxCount-b.Count, b.FaceValue, b.CouponRate)
}
