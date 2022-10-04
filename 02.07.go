package main

import (
	"os"
	"fmt"
	"bufio"
)

func write (in string, out string) {
	InFile, _ := os.Open(in)// Открываем файл для импорта данных
        OutFile, _ := os.Create(out)// Создаём файл для экспорта данных
	defer InFile.Close()// Закрываем файл импорта после исполнения кода
	defer OutFile.Close()// Закрываем файл экспорта после исполнения кода
	reader := bufio.NewScanner(InFile)// Создаём сканер для считывания
        writer := bufio.NewWriter(OutFile)// Создаём сканер для записи
	defer writer.Flush()// Записываем данные из буфера в файл
	var text string //Инициализируем переменную
        for reader.Scan() {
		if reader.Text() == "" {
			panic ("Plase is empty")//Если количество байт в строке 0 вызвать панику
		}
               	writer.WriteString(reader.Text() + "\n")//Загружаем строку в буфер
		fmt.Println(string(reader.Text))
		text = reader.Text()//Присваиваем возвращаемую функцию в переменную
		fmt.Println(text)//Выводим результат функции
        }
}


func main() {
	write("InFile.txt", "OutFile")
}
