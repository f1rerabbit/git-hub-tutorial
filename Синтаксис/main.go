package main
// для вывода
import "fmt"

func main() {
  // fmt.Println("hello world")
// переменные
//   var age = 12
//   var num = 2.3456
//   var str = "hello world"
//   var res int
//   res = a + b
//   // res = a - b
//   // res = a * b
//   // res = a / b
//   // res = a % b
//   fmt.Println (age, num, str)
//
// // Константы
//
// // const pi = 3.14
// var web string = "itProger"
// //количество символо
// // fmt.Println(len(web))ы
// // \n переход на новую строку
// fmt.Println(web + " is cool \nwebsite")

// var isDone bool = true
// fmt.Printf("%t \n", isDone)
// }


// Условия
// var age = 8
// if age < 5 {
//   fmt.Println("Вам пора в детский сад")
// } else if age == 5 {
//   fmt.Println("Вам пора идти школу")
// } else if (age > 5) && (age <= 18) {
//   var grade = age - 5
//   fmt.Println("Пора идти", grade, " klass")
// } else {
//   fmt.Println("Вам пора идти в универ")
//         }
// }

// switch case
// var age = 10
// switch age {
// case 5: fmt.Println("Вам 5 лет")
// case 10: fmt.Println("Вам 10 лет")
// case 15: fmt.Println("Вам 15 лет")
// default: fmt.Println("Вам хз")
// }

//Цикл for
// var i = 1
// for i <= 10 {
//   fmt.Println(i)
//   i++
// }
// for i := 0; i <= 5; i++ {
//   fmt.Println(i)
// }

//Массивы
// var arr[3] int
// arr[0] = 45
// arr[1] = 97
// arr[2] = 76
// fmt.Println(arr[1])
//
// nums := [3]float64 {4.23, 5.23, 98.1}
// for i, value := range nums {
//   fmt.Println(value, i)
// }

//Карты
// websites := make(map[string]float64)
//
// websites["gogolev"] = 0.8
// websites["yandex"] = 0.99
// fmt.Println(websites["gogolev"])

// // Функции
// var a = 3
// var b = 2
// a, b = summ(a, b)
// fmt.Println(a, b)
// }
// func summ(num1 int, num2 int) (int, int)  {
// var res int
// res = num1 + num2
// return res, num1 * num2
// }
// // Вернуть целых 2 числа функции

// Замыкания типа функции
// var num = 3
// multiple := func() int {
// num *= 2
// return num
// }
// fmt.Println(multiple())

// Откладывание
// defer two()
// one()
// }
// func one()  {
// fmt.Println("1")
// }
// func two()  {
// fmt.Println("2")
// }
// var x = 0
// pointer (&x)
// fmt.Println(x)
// }
// // Благодаря * мы работаем с адресом переменной которую хотим изменить
// func pointer(x *int)  {
//   *x = 2
// }

// Структуры
// bob:= cats{"bob", 7, 0.87}
//   fmt.Println("bob function is", bob.test())
// }
// type cats struct {
//   name string
//   age int
//   happiness float64
// }
// func (cat *cats) test() float64  {
// return float64(cat.age) * cat.happiness
// }
