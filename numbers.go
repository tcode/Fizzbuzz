package main;

import fmt "fmt";

func main() {
for num := 0; num < 100; num++ {
if ((num%3) == 0) {
if ((num%5) == 0) {
fmt.Printf("FizzBuzz\n");
} else {
fmt.Printf("Fizz\n");
                }
} else if ((num%5) == 0) {
fmt.Printf("Buzz\n");
} else {
fmt.Printf("%d\n", num);
}
        }
}
