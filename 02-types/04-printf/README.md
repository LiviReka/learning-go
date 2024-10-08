# Formatted print

Implement the below functions that return the argument as nicely formatted text. Use the `fmt.Sprintf` function and consult the [documentation](https://pkg.go.dev/fmt#hdr-Printing) on how to format the output.

1. Write a function `printBool` that receives a boolean as input and returns the text: `variable of type boolean and value <value>`, where `<value>` stands for the current value of the input:

   ```go
   func printBool(b bool) string
   ```

2. Write a function `printInt` that receives an integer as input and returns the text: `variable of type integer and value <value>`, where `<value>` stands for the current value of the input:

   ```go
   func printInt(i int) string
   ```

3. Write a function `printhex` that receives an integer as input and returns the text: `variable of type integer in hexadecimal form and value <value>`, where `<value>` stands for the current value of the input:

   ```go
   func printHex(i int) string
   ```

4. Write a function `printFloat` that receives a floating point number as input and returns the text: `variable of type float and value <value>`, where `<value>` stands for the current value of the input truncated to two decimal places (use the format string `%.2f`):

   ```go
   func printFloat(f float64) string
   ```

5. Write a function `printString` that receives a string as input and returns the text: `variable of type string and value "<value>"`, where `"<value>"` stands for the current value of the input inside double quotes:

   ```go
   func printString(s string) string
   ```

6. Write a function `concatStrings` that receives two strings as argument and returns a single string that contains the two inputs concatenated:

   ```go
   func concatStrings(a, b string) string
   ```

7. Write a function `printConcatStrings` that receives two strings as argument, concatenates them and prints the result using `printString`:

   ```go
   func printConcatStrings(a, b string) string
   ```

Insert your code into the file `exercise.go` at the placeholder `// INSERT YOUR CODE HERE`.
