# Search

Write a function that reads the below text and returns `true` or `false` if "mind" appears in the text *more than once*, without respect to letter case (i.e., the each of words `Hello`, `heLLo`, and `HELlo` match for the search strings `HELLO` and `hello`).

> "One Mind there is; but under it two principles contend. The Mind lets in the light, then the dark, in interaction; so time is generated. At the end Mind awards victory to the light; time ceases and the Mind is complete. He causes things to look different so it would appear time has passed. Matter is plastic in the face of Mind."

Insert your code into the file `exercise.go` at the placeholder `// INSERT YOUR CODE HERE`.

HINT:

- Make sure to remove punctuation from the input (like `.`, `;`, or `?`) and trim the text to alphanumeric characters. For instance, you can use [`regexp.ReplaceAllString`](https://pkg.go.dev/regexp#Regexp.ReplaceAllString) with the regular expression `[^a-zA-Z0-9 ]+` for this purpose.
- You can use a `map[string]int` to remember the number of times a particular word is found while iterating through the text word by word.
