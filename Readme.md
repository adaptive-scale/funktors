# Funktors

A high order function library for apply various array methods like map, reduce, filters etc, using generics.

## Currently supports:

- Foreach
- Map
- Filter
- Concat
- Fill
- GroupBy
- Reduce
- Find
- Sum
- Every

## Example

For eg. given an array for strings, filter if the string contain the word "test". With funktors, you can solve it as follows:

```go
    input := []string{"test123", "123all", "testvalue"}
	values := funktors.Filter(input, func(index int, a string) bool {
		return strings.Contains(a, "test")
	})
	log.Println(values)
}
```

Without funktors, it would look something like this:

```go
    input := []string{"test123", "123all", "testvalue"}
    var output []string{}
        
   for _, v := range input {
        if strings.Contains(v, "test") {
            output = append(output, v)
        }
   }
   log.Println(output)
```

Where it gets interesting is, if the filtering logic starts becoming very complex, that code block looks more entangled.
