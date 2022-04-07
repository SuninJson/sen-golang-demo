package model

type person struct {
	Name string
	age  int //小写开头其它包不能访问到
}

// NewPerson 为其它包提供方便创建该结构体的功能
func NewPerson(name string) *person {
	//校验名称的正确性

	return &person{
		Name: name,
	}
}

func (p *person) SetAge(age int) {
	//校验年龄是否正确
	if age > 0 && age < 150 {
		p.age = age
	} else {
		//处理年龄数据不正确的情况
	}
}
