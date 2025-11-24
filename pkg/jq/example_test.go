package jq

import (
	"fmt"
	"log"
)

func ExampleNewStringQuery() {
	// 从JSON字符串创建查询器
	jsonStr := `{
		"name": "sam",
		"gender": "m",
		"pet": null,
		"skills": [
			"Eating",
			"Sleeping",
			"Crawling"
		],
		"hello.world": true
	}`

	parser, err := NewStringQuery(jsonStr)
	if err != nil {
		log.Fatal(err)
	}

	// 查询简单字段
	name, _ := parser.QueryToString("name")
	fmt.Printf("Name: %s\n", name)

	// 查询嵌套字段（带引号的特殊键名）
	helloWorld, _ := parser.QueryToBool("'hello.world'")
	fmt.Printf("Hello World: %t\n", helloWorld)

	// 查询数组元素
	skill, _ := parser.QueryToString("skills.[1]")
	fmt.Printf("Skill: %s\n", skill)

	// 输出:
	// Name: sam
	// Hello World: true
	// Skill: Sleeping
}

func ExampleNewQuery() {
	// 从Go数据结构创建查询器
	data := map[string]interface{}{
		"users": []interface{}{
			map[string]interface{}{
				"name":   "elgs",
				"gender": "m",
				"skills": []string{"Golang", "Java", "C"},
			},
			map[string]interface{}{
				"name":   "enny",
				"gender": "f",
				"skills": []string{"IC", "Electric design", "Verification"},
			},
		},
	}

	parser := NewQuery(data)

	// 查询数组中的对象
	name, _ := parser.QueryToString("users.[0].name")
	fmt.Printf("First user: %s\n", name)

	gender, _ := parser.QueryToString("users.[1].gender")
	fmt.Printf("Second user gender: %s\n", gender)

	// 查询嵌套数组
	skill, _ := parser.QueryToString("users.[1].skills.[2]")
	fmt.Printf("Second user's third skill: %s\n", skill)

	// 输出:
	// First user: elgs
	// Second user gender: f
	// Second user's third skill: Verification
}

func ExampleQuery() {
	// 使用默认客户端查询
	data := map[string]interface{}{
		"company": "ACME",
		"employees": []interface{}{
			map[string]interface{}{
				"id":         1,
				"name":       "John",
				"department": "Engineering",
				"salary":     50000,
			},
			map[string]interface{}{
				"id":         2,
				"name":       "Jane",
				"department": "Marketing",
				"salary":     55000,
			},
		},
	}

	// 查询公司名称
	company, _ := QueryString(data, "company")
	fmt.Printf("Company: %s\n", company)

	// 查询员工数量
	employees, _ := QueryArray(data, "employees")
	fmt.Printf("Employee count: %d\n", len(employees))

	// 查询特定员工信息
	firstEmpName, _ := QueryString(data, "employees.[0].name")
	firstEmpSalary, _ := QueryInt64(data, "employees.[0].salary")
	fmt.Printf("First employee: %s ($%d)\n", firstEmpName, firstEmpSalary)

	// 输出:
	// Company: ACME
	// Employee count: 2
	// First employee: John ($50000)
}

func ExampleJQ_QueryToInt64() {
	jsonStr := `{
		"product": "Laptop",
		"price": 1299.99,
		"inStock": true
	}`

	parser, _ := NewStringQuery(jsonStr)

	// 查询价格（浮点数转整数）
	price, _ := parser.QueryToInt64("price")
	fmt.Printf("Price: $%d\n", price)

	// 输出:
	// Price: $1299
}

func ExampleJQ_QueryToMap() {
	jsonStr := `{
		"user": {
			"id": 123,
			"profile": {
				"username": "john_doe",
				"email": "john@example.com",
				"preferences": {
					"theme": "dark",
					"notifications": true
				}
			}
		}
	}`

	parser, _ := NewStringQuery(jsonStr)

	// 查询嵌套对象
	profile, _ := parser.QueryToMap("user.profile")
	fmt.Printf("Username: %s\n", profile["username"])

	// 查询更深层的嵌套对象
	preferences, _ := parser.QueryToMap("user.profile.preferences")
	fmt.Printf("Theme: %s\n", preferences["theme"])
	fmt.Printf("Notifications: %t\n", preferences["notifications"])

	// 输出:
	// Username: john_doe
	// Theme: dark
	// Notifications: true
}
