this is GoLang Training Path\
In Maximilian's course, parts 072 to 087\
https://www.udemy.com/course/go-the-complete-guide/

```
// You can pass pointer of struct to functions and when use struct type as argument you can use (*u).propertyName or u.propertyName for access value of propertyName.

outputUserData(&userData)
outputUserData(&userData2)

func outputUserData(u *user) {
 	fmt.Println("You call outputUserData() function", &u)
 	fmt.Println("First Name ==> ", (*u).firstName)
 	fmt.Println("Last Name ==> ", u.lastName)
	fmt.Println("Birth Date ==> ", u.birthDate)
  fmt.Println("Create Date ==> ", u.createDate)
}
```
```
type user struct {
	firstName  string
	lastName   string
	birthDate  string
	createDate time.Time
}

// You can initialization without use name of property but order of value is important.
userData := user{
  userFirstName,
  userLastName,
  userBirthDate,
  time.Now(),
}
```