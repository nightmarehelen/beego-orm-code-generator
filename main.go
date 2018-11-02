package main

import (
	_ "codegen/db"
	"codegen/gen"
)



func main(){
	gen.GenerateModels()
}


