package gen

import "codegen/db"

type Model struct{
	PackageName string
	ModelName string
	Table *db.Table
}
