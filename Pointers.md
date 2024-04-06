## Problems

1. How do you get the memory address of a variable?

   - By using `&`

2. How do you assign a value to a pointer?

   - By using `*` (asterisk)

3. How do you create a new pointer?

   - By using new() function like `xPtr := new(int)`

4. What is the value of x after running this program:

   - x=2.25

```sh
func square(x *float64) {
    *x = *x * *x
}

func main() {
    x := 1.5
    square(&x)
}
```

5. Write a program that can swap two integers (x := 1; y := 2; swap(&x, &y) should give you x=2 and y=1).

   ```sh
    func swap(x *int, y *int){
        temp:=*x
        *x=*y
        *y=temp
   }

   func main(){
    swap(&x,&y)
   }
   ```
