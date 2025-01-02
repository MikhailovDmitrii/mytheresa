package infra

type Product struct {
	Sku      string `gorm:"column:sku;primaryKey;<-:create"`
	Name     string `gorm:"column:name;not null"`
	Category string `gorm:"column:category;not null"`
	Price    int    `gorm:"column:price;not null"`
}
