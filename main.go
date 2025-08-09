package main

import "fmt"

func divide(a, b int) (int, error) {
	if b == 0 {
		return 0, fmt.Errorf("На ноль делить нельзя")
	}
	return a / b, nil
}

func main() {
	for {
		var a, b int
		fmt.Println("Введите целые числа: делимое и делитель")
		fmt.Scan(&a, &b)
		result, err := divide(a, b)
		if err != nil {
			fmt.Println(err)
			return
		}
		fmt.Println(result)

		switch {
		case result > 10:
			fmt.Println("Результат большой")
		case result >= 1 && result <= 10:
			fmt.Println("Результат средний")
		case result < 1:
			fmt.Println("Результат маленький или ноль")
		}
	}
	var answer string
	fmt.Print("Хотите ввести новые числа? (y/n): ")
	fmt.Scan(&answer)
	if answer != "y" && answer != "Y" {
		fmt.Println("Выход из программы.")
		break
	}
}
