package schema

var Models = []interface{}{}

type Book struct {
	//lint:ignore U1000 tableName
	tableName interface{} `pg:"book"`
	Id        int         `pg:"id,pk"`
	Name      string      `pg:"name"`
}

type OrderItem struct {
	//lint:ignore U1000 tableName
	tableName interface{} `pg:"order_item"`
	Id        int         `pg:"id,pk"`
	OrderId   int         `pg:"order_id"`
	BookId    int         `pg:"book_id"`
	Quantity  int         `pg:"quantity"`
}

type Order struct {
	//lint:ignore U1000 tableName
	tableName interface{} `pg:"order"`
	Id        int         `pg:"id,pk"`
	// Items     []OrderItem `pg:"rel:has-many"`
}

type OrderWithItem struct {
	//lint:ignore U1000 tableName
	// 组合不能访问私有变量，需要重新定义表名
	tableName interface{} `pg:"order"`
	Order
	Items []OrderItem `pg:"rel:has-many"`
}

func init() {
	Models = append(Models, (*Book)(nil))
	Models = append(Models, (*OrderItem)(nil))
	Models = append(Models, (*Order)(nil))
}
