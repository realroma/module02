package main

import (
	"os"
	"fmt"
	"time"
	"bufio"
)
var start time.Time
var m time.Time
func logTime() {
	if start == m { //Проверка на наличие точки отсчёта
		start = time.Now() //Записывает момент во времени который будет точкой отсчёта
		return //Выход(завершение) функции
	}
	time := time.Now()//Создание второй точки
	res := time.Sub(start)//Вычисление промежутка между time и start
        fmt.Println(res)//Вывод времени между time и start

}



func main() {
	logTime()//Вызов функции и начало отсчёта
	file, err := os.Create("out.txt")     // создаем файл
	r, _ := os.Open("in.txt") // Открываем файл который будем сканировать построчно
	writer := bufio.NewWriter(file)//Указываем куда ведётся запись с буфера
	reader := bufio.NewScanner(r)
	var x int = 1
	for reader.Scan() {
		fmt.Println(x, reader.Text())
		x++
	}
	writer.WriteString("1")//Запись строки в буфер
	writer.Flush()//Запись из буфера в файл
	if err != nil {
		fmt.Println("i have been stopping of error")
	} //Проверка на наличие ошибок
	logTime()//Запись времени от прошлого logTime
	file.Close()//Закрывает функцию os.Open
}
