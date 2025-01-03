package animals

import "gorm.io/gorm"

type Animal struct {
	gorm.Model
	Name string
	Type string
	Size string
}
