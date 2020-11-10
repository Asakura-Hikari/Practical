package util

type Customer struct {
	Id int
	Name string
	Gander string
	Age int
	Phone string
	Email string
}

func (c *Customer)NewCustomer(id int, name string, gander string, age int, phone string, email string){
	c.Id = id
	c.Name = name
	c.Gander = gander
	c.Age = age
	c.Phone = phone
	c.Email = email
}