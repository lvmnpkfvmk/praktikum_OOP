package program

import (
	"fmt"
	"investgo/internal/game/fund"
	in "investgo/internal/game/investment"
	cond "investgo/internal/game/conditions"
	u "investgo/internal/utils"
)

// Класс для игры
type Game struct {
	CyclesCount      int
	CurrentCycle     int
	MyInvestmentFund *fund.InvestmentFund
	Investments      []in.Investment
}

// Конструктор
func NewGame() *Game {
	return &Game{
		CyclesCount:      12,
		MyInvestmentFund: fund.NewInvestmentFund(560000),
		Investments: []in.Investment{
			in.NewStock("Компания №1", 1000, 1000),
			in.NewStock("Компания №2", 100, 100),
			in.NewStock("Компания №3", 10, 10),
			in.NewBond("Облигация №1", 1000, 1000, 0.05),
			in.NewBond("Облигация №2", 100, 100, 0.1),
			in.NewBond("Облигация №3", 10, 10, 0.15),
			in.NewMetal("Золото №1", 1000, 1000, 0.1),
			in.NewMetal("Серебро №2", 100, 100, 1),
			in.NewMetal("Бронза №3", 10, 10, 10),
			in.NewDeposit("Вклад №1", 1000, 1000),
			in.NewDeposit("Вклад №2", 100, 100),
			in.NewDeposit("Вклад №3", 10, 10),
		},
	}
}

// Метод для расчета доходности инвестиций
func (g *Game) CalculateInvestmentReturns() {
	fmt.Println("Расчитываем доходность инвестиций...")
	g.MyInvestmentFund.CalculateInvestmentReturns()
}

// Метод для оплаты налогов
func (g *Game) PayTaxes() {
	fmt.Println("Платим налоги...")
	g.MyInvestmentFund.PayTaxes()
}

// Метод для учета новых поступлений и расходов
func (g *Game) AccountForNewIncomeAndExpenses() {
	fmt.Println("Учитываем новые поступления и расходы...")
	g.MyInvestmentFund.AccountForNewIncomeAndExpenses()
}
// Метод для покупки инвестиций
func (g *Game) BuyInvestment() {
    fmt.Println("Выберите тип инвестиций, которые хотите купить:")
    fmt.Println("1. Акции")
    fmt.Println("2. Облигации")
    fmt.Println("3. Металлы")
    fmt.Println("4. Вклады")
    fmt.Println("0. Отмена")

    option := u.ChoiceNumber(4, true)
    if option == 0 {
        return
    }

    var investments []in.Investment
    var investmentType string
    switch option {
    case 1:
    	for _, inv := range g.Investments {
			if _, ok := inv.(*in.Stock); ok {
				investments = append(investments, inv)
			}
    	}
        investmentType = "акцию, которую хотите купить"
    case 2:
    	for _, inv := range g.Investments {
			if _, ok := inv.(*in.Bond); ok {
				investments = append(investments, inv)
			}
    	}
        investmentType = "облигацию, которую хотите купить"
    case 3:
    	for _, inv := range g.Investments {
			if _, ok := inv.(*in.Metal); ok {
				investments = append(investments, inv)
			}
    	}
        investmentType = "металл, который хотите купить"
    case 4:
    	for _, inv := range g.Investments {
			if _, ok := inv.(*in.Deposit); ok {
				investments = append(investments, inv)
			}
    	}
        investmentType = "вклад, который хотите сделать"
    }

    fmt.Printf("Выберите %s:\n", investmentType)
    for i, inv := range investments {
        fmt.Printf("%d. %s\n", i+1, inv)
    }
    fmt.Println("0. Отмена")

    option = u.ChoiceNumber(len(investments), true)
    if option == 0 {
        return
    }

    g.MyInvestmentFund.BuyInvestment(investments[option-1])
}

// Метод для продажи инвестиций
func (g *Game) SellInvestment() {
    fmt.Println("Выберите тип инвестиций, которые хотите продать:")
    fmt.Println("1. Акции")
    fmt.Println("2. Облигации")
    fmt.Println("3. Металлы")
    fmt.Println("4. Вклады")
    fmt.Println("0. Отмена")

    option := u.ChoiceNumber(4, true)
    if option == 0 {
        return
    }

    var investments []in.Investment
    switch option {
    case 1:
    	for _, inv := range g.Investments {
			if _, ok := inv.(*in.Stock); ok {
				investments = append(investments, inv)
			}
    	}
    case 2:
    	for _, inv := range g.Investments {
			if _, ok := inv.(*in.Bond); ok {
				investments = append(investments, inv)
			}
    	}
    case 3:
    	for _, inv := range g.Investments {
			if _, ok := inv.(*in.Metal); ok {
				investments = append(investments, inv)
			}
    	}
    case 4:
    	for _, inv := range g.Investments {
			if _, ok := inv.(*in.Deposit); ok {
				investments = append(investments, inv)
			}
    	}
    }

    fmt.Println("Выберите инвестицию, которую хотите продать:")
    for i, inv := range investments {
        fmt.Printf("%d. %s\n", i+1, inv)
    }
    fmt.Println("0. Отмена")

    option = u.ChoiceNumber(len(investments), true)
    if option == 0 {
        return
    }

    g.MyInvestmentFund.SellInvestment(investments[option-1])
}

// Метод для реструктуризации портфеля
func (g *Game) RestructurePortfolio() {
	fmt.Println("Реструктуризация портфеля...")

	for {
		fmt.Println("Выберите действие:")
		fmt.Println("1. Купить инвестиции")
		fmt.Println("2. Продать инвестиции")
		fmt.Println("0. Конец реструктуризации")

		option := u.ChoiceNumber(2, true)
		if option == 0 {
			return
		}

		switch option {
		case 1:
			g.BuyInvestment()
		case 2:
			g.SellInvestment()
		}
	}
}

// Метод для запуска цикла
func (g *Game) Cycle() {
	fmt.Printf("Начинаем цикл %d из %d...\n", g.CurrentCycle+1, g.CyclesCount)

	cond.UpdateInvestment(g.Investments)

	g.CalculateInvestmentReturns()

	g.PayTaxes()

	g.AccountForNewIncomeAndExpenses()

	g.RestructurePortfolio()

	g.CurrentCycle++
}

// Метод для запуска игры
func (g *Game) Run() {
	fmt.Println("Добро пожаловать в систему управления инвестиционным портфелем!")
	fmt.Printf("Всего будет %d циклов.\n", g.CyclesCount)
	for i := 0; i < g.CyclesCount; i++ {
		g.Cycle()
	}
}
