package concurrent

func printOne(a chan <- int) {
	for i :=0; i< 100; i++ {
		a <- i
		if i % 2 == 1 {
			fmt.Println("go1: %v", i)
		}
	}

}

func printTwo(a <- chan int) {
for i :=0 i< 100; i++ {
		<- a
		if i % 2 == 0 {
			fmt.Println("go2: %v", i)
		}	
	}
}