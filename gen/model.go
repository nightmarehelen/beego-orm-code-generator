package gen

import "codegen/db"

type Model struct{
	PackageName string
	ModelPackage string
	ModelName string
	Table *db.Table
}
