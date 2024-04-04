package main

type Person struct {
	Name    string
	Country *Country
}

type Country struct {
	Name     string
	Citizens []*Person
}

func main() {
	usa := &Country{Name: "USA"}
	john := &Person{Name: "John", Country: usa}
	mary := &Person{Name: "Mary", Country: usa}
	usa.Citizens = []*Person{john, mary}
	// 手动解除 john 和 usa 的相互引用
	john.Country = nil
	usa.Citizens = []*Person{mary}
}
