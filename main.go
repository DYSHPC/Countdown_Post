package main

import (
	"fmt"
	"math/rand"
	"time"
)

func cool(x string, y []string, f string) {
	for i := 0; i < 20; i++ {
		k := y[rand.Intn(len(y))]
		z := x + k
		fmt.Printf("%s\r", z)
		time.Sleep(10 * time.Millisecond)
	}
	fmt.Printf("%s\r", x+f)
}

func main() {
	rand.Seed(time.Now().UnixNano())

	a := "大家好我們是程式設計社~8/20 大園高中社團博覽會-茫茫人海遇見你 倒數4天！ 一定要來喔！"
	b := []string{"h", "e", "l", "o", "w", "r", "d"}
	g := []string{"大", "家", "好", "我", "們", "是", "程", "式", "設", "計", "社", "~", "8", "/", "2", "0", " ", "大", "園", "高", "中", "社", "團", "博", "覽", "會", "-", "茫", "茫", "人", "海", "遇", "見", "你", " ", "倒", "數", "4", "天", "！", " ", "一", "定", "要", "來", "喔", "！"}

	c := ""
	d := "hello world"
	for _, ch := range d {
		cool(c, b, string(ch))
		c += string(ch)
	}
	fmt.Println()

	gc := ""
	for _, ch := range a {
		cool(gc, g, string(ch))
		gc += string(ch)
	}
	fmt.Println()

	fmt.Println("╱╱╭╮")
	fmt.Println("╱╱┃┃")
	fmt.Println("╭━╯┣╮╱╭┳━━┳━━╮")
	fmt.Println("┃╭╮┃┃╱┃┃╭╮┃╭━╯")
	fmt.Println("┃╰╯┃╰━╯┃╰╯┃╰━╮")
	fmt.Println("╰━━┻━╮╭┫╭━┻━━╯")
	fmt.Println("╱╱╱╭━╯┃┃┃")
	fmt.Println("╱╱╱╰━━╯╰╯")
	fmt.Println()
	fmt.Println("     ***      ***")
	fmt.Println("    *****    *****")
	fmt.Println("     ***      ***")
	fmt.Println("*                    *")
	fmt.Println(" *                  *")
	fmt.Println("  *                *")
	fmt.Println("   *              *")
	fmt.Println("    **************")
}
