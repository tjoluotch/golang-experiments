package main

type MyReader struct {

}

// Implement a Reader type that emits an infinite stream of the ASCII character 'A'.

// Type my reader implements Read() method in reader interface
func(r MyReader) Read(bytes []byte)(int, error){
	for i := range bytes{
		bytes[i] = 65
	}
	return len(bytes), nil
}

func main() {
}
