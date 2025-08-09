package main

import "fmt"

func divide(a, b int) (int, error) {
	if b == 0 {
		return 0, fmt.Errorf("На ноль делить нельзя")
	}
	return a / b, nil
}

func main() {
	var answer string

	for {
		var a, b int
		fmt.Println("Введите целые числа: делимое и делитель")
		_, scanErr := fmt.Scan(&a, &b)
		if scanErr != nil {
			fmt.Println("Ошибка ввода:", scanErr)
			// спросим, пробовать ли ещё
			fmt.Print("Хотите попробовать снова? (y/n): ")
			fmt.Scan(&answer)
			if answer != "y" && answer != "Y" {
				fmt.Println("Выход из программы.")
				break
			}
			continue
		}

		result, err := divide(a, b)
		if err != nil {
			fmt.Println(err)
			// можно спросить, продолжать ли после ошибки
		} else {
			fmt.Println("Результат:", result)
			switch {
			case result > 10:
				fmt.Println("Результат большой")
			case result >= 1 && result <= 10:
				fmt.Println("Результат средний")
			default:
				fmt.Println("Результат маленький или отрицательный")
			}
		}

		fmt.Print("Хотите ввести новые числа? (y/n): ")
		fmt.Scan(&answer)
		if answer != "y" && answer != "Y" {
			fmt.Println("Выход из программы.")
			break
		}
	}
}
