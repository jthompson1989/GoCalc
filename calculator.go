package main

import (
	"fmt"
	"strconv"
	"strings"
)

type Calculator struct {
	displayText        string
	calculationHistory []Calculation
}

type Calculation struct {
	Equation string
	Result   int
}

var operators = map[string]int{
	"+": CalculationType.Addition,
	"-": CalculationType.Subtract,
	"/": CalculationType.Division,
	"*": CalculationType.Multiply,
	"x": CalculationType.Multiply,
}

func (calc *Calculator) UpdateDisplayText(displayValue string) {
	calc.displayText = displayValue
}

func (calc *Calculator) SimpleCalculation(calculation string) int {
	var result int
	var isContinuingEquation = strings.Contains("+x*/-", strings.Split(calculation, "")[0])
	if isContinuingEquation {
		var lastResult = calc.calculationHistory[len(calc.calculationHistory)-1].Result
		calculation = strconv.Itoa(lastResult) + calculation
	}

	//Foreach Loop in GoLang
	for operator, calcType := range operators {
		if strings.Contains(calculation, operator) {
			values := strings.Split(calculation, operator)
			value1, _ := strconv.Atoi(values[0])
			value2, _ := strconv.Atoi(values[1])
			result = Calculate(value1, value2, calcType)
			break
		}
	}

	calc.displayText = strconv.Itoa(result)
	calc.calculationHistory = append(calc.calculationHistory, Calculation{
		Equation: calculation,
		Result:   result,
	})
	return result
}

func (calc *Calculator) DisplayHistory() {
	var calcHistory Calculation
	for index := 0; index < len(calc.calculationHistory); index++ {
		calcHistory = calc.calculationHistory[index]
		fmt.Println(calcHistory.Equation + "=" + strconv.Itoa(calcHistory.Result))
	}
}

//My Attempt at Enum

func newCalcuationTypeRegistry() *calculationType {
	return &calculationType{
		Subtract: 0,
		Addition: 1,
		Multiply: 2,
		Division: 3,
	}
}

type calculationType struct {
	Subtract int
	Addition int
	Multiply int
	Division int
}

func Calculate(value1, value2 int, typeOfCalculation int) int {
	switch typeOfCalculation {
	case CalculationType.Subtract:
		return value1 - value2
	case CalculationType.Addition:
		return value1 + value2
	case CalculationType.Multiply:
		return value1 * value2
	case CalculationType.Division:
		return value1 / value2
	default:
		return value1 + value2
	}
}

var CalculationType = newCalcuationTypeRegistry()
