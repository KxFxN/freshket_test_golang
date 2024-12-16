package main

import (
	"fmt"
	"math"
)

type FoodStore struct {
	Menu           map[string]float64
	BundleDiscount []string
}

func NewFoodStore() *FoodStore {
	return &FoodStore{
		Menu: map[string]float64{
			"Red":    50,
			"Green":  40,
			"Blue":   30,
			"Yellow": 50,
			"Pink":   80,
			"Purple": 90,
			"Orange": 120,
		},
		BundleDiscount: []string{"Orange", "Pink", "Green"},
	}
}

func (c *FoodStore) CalculatePrice(order map[string]int, isMember bool) (float64, error) {
	if err := c.ValidateOrder(order); err != nil {
		return 0, err
	}

	totalPrice := c.calculateBaseTotal(order)

	totalPrice = c.applyBundleDiscounts(order, totalPrice)

	if isMember {
		totalPrice *= 0.9
	}

	return math.Round(totalPrice*100) / 100, nil
}

func (c *FoodStore) ValidateOrder(order map[string]int) error {
	for item := range order {
		if _, exisis := c.Menu[item]; !exisis {
			return fmt.Errorf("รายการอาหารไม่ถูกต้อง: %s", item)
		}
	}
	return nil
}

func (c *FoodStore) calculateBaseTotal(order map[string]int) float64 {
	var total float64
	for item, quantity := range order {
		total += c.Menu[item] * float64(quantity)
	}

	return total
}

// คำนวณส่วนลดแบบชุด
func (c *FoodStore) applyBundleDiscounts(order map[string]int, totalPrice float64) float64 {
	for _, item := range c.BundleDiscount {
		if quantity, exists := order[item]; exists {
			// คำนวณคู่ที่จะได้ส่วนลด
			bundles := quantity / 2
			if bundles > 0 {
				// ลด 5% ต่อคู่ แต่ไม่ลดจากยอดรวม
				bundleDiscount := c.Menu[item] * 2 * float64(bundles) * 0.05
				totalPrice -= bundleDiscount
			}
		}
	}
	return totalPrice
}

func main() {
	calculator := NewFoodStore()

	order1 := map[string]int{
		"Red":   1,
		"Green": 1,
	}
	Desk1(calculator, order1)

	order2 := map[string]int{
		"Orange": 5,
	}
	Desk2(calculator, order2)

}

func Desk1(c *FoodStore, order map[string]int) {
	fmt.Println("ไม่มีสมาชิก")
	fmt.Printf("ออเดอร์: %v\n", order)

	price, _ := c.CalculatePrice(order, false)
	fmt.Printf("ราคารวม: %.2f บาท\n", price)

	fmt.Println("มีสมาชิก")
	fmt.Printf("ออเดอร์: %v\n", order)

	memberPrice, _ := c.CalculatePrice(order, true)
	fmt.Printf("ราคารวม: %.2f บาท\n", memberPrice)
}

func Desk2(c *FoodStore, order map[string]int) {
	fmt.Println("Bundles")
	fmt.Printf("ออเดอร์: %v\n", order)

	price, _ := c.CalculatePrice(order, false)
	fmt.Printf("ราคารวม: %.2f บาท\n", price)

	fmt.Println("Bundles order สำหรับ สมาชิก")
	fmt.Printf("ออเดอร์: %v\n", order)

	memberPrice, _ := c.CalculatePrice(order, true)
	fmt.Printf("ราคารวม: %.2f บาท\n", memberPrice)
}
