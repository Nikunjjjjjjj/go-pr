package main

import (
	"encoding/json"
	"fmt"
)

func main() {
	//fitnessTracker()
	//EmployeeAttendaceTracker()
	//employeeSalaryCal()
	//orderSummaryAnalyzwer()
	//libraryBookManager()
	studentPerformanceTasker()
	//concurrentFileDownload()
}

// 1.
func fitnessTracker() {
	stepsPerHour := [24]int{599, 501, 501, 501, 501, 501, 501, 0, 501, 501, 5015, 499, 150, 150, 234, 1501, 501, 501, 501, 501, 501, 501, 389, 900}
	//stepsPerHour := [24]int{499, 150, 150, 234, 1501, 501, 501, 501, 501, 501, 501, 389, 900}

	totalSteps := 0
	hoursWhereStepLess := []int{}
	for i, v := range stepsPerHour {
		if v == 0 {
			continue
		}
		if v < 500 {
			hoursWhereStepLess = append(hoursWhereStepLess, i+1) //isse print karu?
		}
		totalSteps += v
	}
	if totalSteps < 8000 {
		fmt.Printf("you walked %d steps.try harder tommmorow!", totalSteps)
		fmt.Println("time when steps are less", hoursWhereStepLess)
	} else {
		fmt.Printf("Awesome! you walked %d steps today", totalSteps)
	}
}

// 2
type Attendance struct {
	checkInTime  string
	checkOutTime string
	status       string
}
type employee struct {
	name       string
	id         int
	attendance Attendance
}

func EmployeeAttendaceTracker() {
	var emp1 employee = employee{"niku", 284, Attendance{"10:00", "null", "Present"}}
	//fmt.Println(emp1)
	emp1.markCheckOutTime("6:00")

	str := emp1.getSummary()
	fmt.Println(str)
}
func (e *employee) markCheckOutTime(time string) {
	if e.attendance.checkOutTime == "null" {
		e.attendance.checkOutTime = time
		e.attendance.status = "completed"
	}
	//if checkouttime!=null will not do anything
	fmt.Printf("employee name :%s,checkoutTime: %s", e.name, e.attendance.checkOutTime)

}

func (e *employee) getSummary() string {
	str := fmt.Sprintf("Employee %s(ID:%d)-CheckIn:%s|ChecKout:%s|status:%s", e.name, e.id, e.attendance.checkInTime, e.attendance.checkOutTime, e.attendance.status)
	return str
}

// 3
type employee3Q struct {
	name       string
	basesalary float64
	department string
}

func employeeSalaryCal() {
	// var emp1 employee3Q = employee3Q{"niku", 40000, "Engineering"}
	var emp1 employee3Q = employee3Q{"niku", 40000, "Sales"}

	emp1.calculateSalary()
}

func (e employee3Q) calculateSalary() {
	var sal float64
	if e.department == "Engineering" {
		sal = e.basesalary + e.basesalary*0.05
	} else if e.department == "Sales" {
		sal = e.basesalary + e.basesalary*0.1
	}
	fmt.Printf("Final salary for %s(%s):%f", e.name, e.department, sal)
}

// 4
func orderSummaryAnalyzwer() {
	orders := []float64{
		499.5, 1200, 799, 300, 1500,
	}
	var totalRevenue float64
	var avgOrdervalue float64
	numberOfHighvalueOrder := 0
	for _, v := range orders {
		totalRevenue += v
		if v > 1000 {
			numberOfHighvalueOrder++
		}
	}
	avgOrdervalue = totalRevenue / float64(len(orders))
	fmt.Println("Total Revenue:", totalRevenue)
	fmt.Println("Average order value:", avgOrdervalue)
	fmt.Println("High value orders", numberOfHighvalueOrder)
}

// 5
var books = map[string]int{
	"go programming":  3,
	"python101":       5,
	"data structures": 2,
}

func libraryBookManager() {
	// books := map[string]int{
	// 	"go programming":3,
	// 	"python101":5,
	// 	"data structures":2,
	// }
	addBook("hyledger", 5)
	RemoveBook("python101")
	listBooks()
}

func addBook(title string, qty int) {
	books[title] += qty
}

func RemoveBook(title string) {
	delete(books, title)
}

func listBooks() {
	fmt.Println(books)
}

// 6
type student struct {
	name  string
	Marks int
	Grade string
}

func studentPerformanceTasker() {
	studentData := []student{
		student{"Amit", 85, "A"},
		student{"Sara", 70, "B"},
		student{"Nikunj", 100, "A+"},
	}
	fmt.Println("before", studentData)
	data, _ := json.Marshal(studentData)
	fmt.Println(string(data))

}

// 7
func concurrentFileDownload() {
	// // 	files:=[]string{"file1.pdf","file2.pdf"}
	// for i,file :range files{
	// 	//we will download file one by one
	// 	then downloaded content will be passed to different channels naming them as chan+i

	// }
	// then we will add aitgroaup.add(len(files))
	// then i will make another loop for len(files) in which we will make a goroutine and in
	// each we will receive value from channel after that we will do wait group.done()

	// at the end we will place waitgroup.wait which is waiting for all goroutines to finish and we will print result
}

//8 on other file
