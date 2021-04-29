package entity

type Product struct {
	Id         int     `gorm:"TYPE:SERIAL;PRIMARY_KEY;NOT NULL;COLUMN:id" json:"id"`
	Name       string  `gorm:"TYPE:varchar;NOT NULL;COLUMN:name" json:"name"`
	Price      float32 `gorm:"TYPE:numeric;NOT NULL;COLUMN:price" json:"price"`
	IdCategory int     `gorm:"TYPE:int;NOT NULL;COLUMN:id_category" json:"id_category"`
	IdMarket   string  `gorm:"TYPE:uuid;NOT NULL;COLUMN:id_market" json:"id_market"`
}

func (Product) TableName() string {
	return "Product"
}
