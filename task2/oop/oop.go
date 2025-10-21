package main

import "fmt"

/*
*定义一个 Shape 接口，包含 Area() 和 Perimeter() 两个方法。然后创建 Rectangle 和 Circle 结构体，
实现 Shape 接口。在主函数中，创建这两个结构体的实例，并调用它们的 Area() 和 Perimeter() 方法。
*/
type Shape interface {
	Area() float64
	Perimeter() float64
}

type Rectangle struct {
	width  float64
	height float64
}

func (rect *Rectangle) Area() float64 {
	return rect.width * rect.height
}
func (rect *Rectangle) Perimeter() float64 {
	return 2 * (rect.width + rect.height)
}

type Circle struct {
	radius float64
}

func (c *Circle) Perimeter() float64 {
	return c.radius * c.radius
}
func (c *Circle) Area() float64 {
	return (22 / 7.0) * (c.radius * c.radius)
}

func main() {
	var circ Shape = &Circle{3}
	var rect Shape = &Rectangle{10, 5}
	fmt.Println("Rectangle area:", rect.Area())
	fmt.Println("Rectangle perimeter:", rect.Perimeter())
	fmt.Println("Circle perimeter:", circ.Perimeter())
	fmt.Println("Circle area:", circ.Area())
	fmt.Println("=============================================================")
	emp := Employee{Person{"John", 30}, 1001}
	emp.PrintInfo()
}

/*
*
使用组合的方式创建一个 Person 结构体，包含 Name 和 Age 字段，再创建一个 Employee 结构体，组合 Person 结构体并添加 EmployeeID 字段。
为 Employee 结构体实现一个 PrintInfo() 方法，输出员工的信息。
*/
type Person struct {
	Name string
	Age  int
}

type Employee struct {
	Person
	EmployeeID int
}

func (e *Employee) PrintInfo() {
	fmt.Printf("Name: %s, Age: %d, EmployeeID: %d\n", e.Name, e.Age, e.EmployeeID)
}
