package main

import "fmt"
//Домашнее задание 4
func main() {
  fmt.Println("========================")
  fmt.Println("Задача 1.")
  fmt.Println("========================")
  Task1()
  fmt.Println("========================")
  fmt.Println("Задача 2.")
  fmt.Println("========================")
  Task2()
  fmt.Println("========================")
  fmt.Println("Задача 3.")
  fmt.Println("========================")
  Task3()
  fmt.Println("========================")
  fmt.Println("Задача 4.")
  fmt.Println("========================")
  Task4()
  fmt.Println("========================")
  fmt.Println("Задача 5.")
  fmt.Println("========================")
  Task5()

}

func task1() {
/*Задача 1. Баллы ЕГЭ
  Сумма проходных баллов в институт составляет 275 баллов.
  Напишите программу, которая запрашивает у пользователя результаты ЕГЭ по трём экзаменам и проверяет, поступил ли он в институт или нет.
*/  
  
  var val, amountExams int
  messageBal := "Количество баллов за экзамен"
  fmt.Println("Баллы ЕГЭ \n")
  fmt.Println("Введите параметры:")
  
  fmt.Print(messageBal, "1:")
  fmt.Scan(&val)
  amountExams += val
  fmt.Print(messageBal, "2:")
  fmt.Scan(&val)
  amountExams += val
  fmt.Print(messageBal, "3:")
  fmt.Scan(&val)
  amountExams += val
  
  if amountExams  >= 275{
    fmt.Println("Поздравляю! Вы приняты в институт!")
  } else {
    fmt.Println("Вы не приняты в институт!")
  }
}

func task2() {
/*Задача 2. Три числа
  Напишите программу, которая запрашивает у пользователя три числа и сообщает, есть ли среди них число, большее, чем 5.
*/
  var val1, val2, val3 int
  fmt.Println("Три числа \n")
  fmt.Print("Введите первое число:")
  fmt.Scan(&val1)
  fmt.Print("Введите второе число:")
  fmt.Scan(&val2)
  fmt.Print("Введите третье число:")
  fmt.Scan(&val3)


  if val1  > 5 {
    fmt.Println("Первое число", val1, "больше 5")
  }
  if val2  > 5 {
    fmt.Println("Второе число", val2, "больше 5")
  }
  if val3  > 5 {
    fmt.Println("Теретье число", val3, "больше 5")
  }
}
func task3() {
/*Задача 3. Банкомат
  Пользователи банкомата хотят снимать деньги. Но банкомат умеет выдавать только купюры по 100 рублей, а максимальная сумма снятия — 100 000 рублей.
  Напишите программу, которая проверяет допустимость суммы средств, которую ввёл пользователь. Сумма не должна быть меньше 100 и больше 100 000 руб.
*/
  var val int
  fmt.Println("Банкомат \n")
  fmt.Print("Введите сумму для снятия:")
  fmt.Scan(&val)

  if val < 100 {
    fmt.Println("Банкомат не может выдать сумму меньше 100 руб.")
  } else if val > 100000 {
    fmt.Println("В банкомате не достаточно средств")
  } else if val % 100 == 0 {
    fmt.Println("Выдача наличных", val, " руб.")
  } else {
    fmt.Println("Банкомат может выдать только ратное 100 руб.")
  }
}

func task4() {
/*Задача 4. Три числа: еще попытка
  Напишите программу, которая запрашивает у пользователя три числа и выводит количество чисел, которые больше, либо равны 5.
*/ var val1, val2, val3, numberHigher5 int
  fmt.Println("Три числа: еще попытка \n")
  fmt.Print("Введите первое число:")
  fmt.Scan(&val1)
  fmt.Print("Введите второе число:")
  fmt.Scan(&val2)
  fmt.Print("Введите третье число:")
  fmt.Scan(&val3)


  if val1  > 5 {
    numberHigher5 ++
  }
  if val2  > 5 {
    numberHigher5 ++
  }
  if val3  > 5 {
    numberHigher5 ++
  }

  fmt.Println("Чесел больше 5:", numberHigher5)
}

func task5() {
/*Задача 5. Ресторан
В ресторане действуют следующие правила:
      по понедельникам должна применяться скидка 10% на всё меню, потому что понедельник — день тяжёлый!
      по пятницам, если сумма чека превышает 10 000 руб., включается дополнительная скидка в размере 5%;
      если число гостей в одной компании превышает 5 человек, автоматически включается надбавка на обслуживание 10%.
  Напишите программу, которая запрашивает день недели, число гостей и сумму чека и рассчитывает сумму к оплате.
*/
  var guests, check, discount, surcharge  int
  var day string
  fmt.Println("Ресторан \n")
  fmt.Print("Введите день недели:")
  fmt.Scan(&day)
  fmt.Print("Введите количество гостей:")
  fmt.Scan(&guests)
  fmt.Print("Введите сумму чека:")
  fmt.Scan(&check)

  if day == "Понедельник"{
    discount = check * 10  / 100
  } else if day == "Пятница"{
    if check > 10000{
    discount = check * 5  / 100
    }
  }
  if guests > 5 {
    surcharge = check * 10 / 100  
  } 
  fmt.Print("Сумма к оплате:", check - discount + surcharge, "\n")
}