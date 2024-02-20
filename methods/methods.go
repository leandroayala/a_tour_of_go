package main

/*
Go does not have classes. However, you can define methods on types.

A method is a function with a special receiver argument.

The receiver appears in its own argument list between the func keyword and the method name.

*/
type Names struct {
	name     string
	lastname string
}

//Define the method ListName that receives Names struct  TYPE
func (name Names) ListName() {

	println(name.name + " " + name.lastname)

}

func main() {

	//Here is like insancing a class
	name := Names{
		name:     "Leandro",
		lastname: "Ayala",
	}

	//call the method
	name.ListName()
}
