package main
import "fmt"

func main() {
        var i, sum int
        score :=  []int{ 8, 7, 8, 9, 9}
        for i = 0; i < len(score); i++ {
                sum += score[i]
        }
      fmt.Println("sum is ", sum, "   i is ", i)
}