package packageone

var privateVar = "I am private"
var PublicVar = "I am public (or exported)" // This is exported because it starts with a capital letter

// notExported is not exported because it starts with a lowercase letter
func notExported() {

}

// Exported is exported because it starts with a capital letter
func Exported() {

}
