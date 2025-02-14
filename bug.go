func main() {


  x := make([]int, 0)

  for i := 0; i < 10; i++ {

    x = append(x, i)

  }

  fmt.Println(x)


  y := x[:5]

  fmt.Println(y)


  y[0] = 1000

  fmt.Println(y)

  fmt.Println(x)

}