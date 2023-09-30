package main

import "errors"

func Addition(firstNumber, secondNumber float64) float64 {
	return firstNumber + secondNumber
}

func Subtraction(firstNumber, secondNumber float64) float64 {
	return firstNumber - secondNumber
}

func Division(firstNumber, secondNumber float64) (float64, error) {
  if firstNumber == 0 || secondNumber == 0 {
    return 0, errors.New("Cannot divide by zero")
  }
  return firstNumber / secondNumber, nil
}

func Multiplication(firstNumber, secondNumber float64) float64 {
	return firstNumber * secondNumber
}