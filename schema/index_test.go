package schema_test

import (
	"context"
	"testing"

	"github.com/go-pg/pg/v10"
	"github.com/go-pg/pg/v10/orm"
	"github.com/mole828/pgLearn/conn"
	"github.com/mole828/pgLearn/schema"
)

var db = conn.GetDB()

func TestConnect(t *testing.T) {
	if err := db.Ping(context.Background()); err != nil {
		t.Error(err)
	}
}

func TestCreateTable(t *testing.T) {
	var total = 0
	for _, model := range schema.Models {
		var err = db.Model(model).CreateTable(&orm.CreateTableOptions{
			IfNotExists: true,
		})
		if err != nil {
			t.Error(err)
			return
		} else {
			total += 1
		}
	}
	t.Logf("%v table aready", total)
}

func TestInsert(t *testing.T) {
	books := []*schema.Book{
		{Id: 1, Name: "Go Programming"},
		{Id: 2, Name: "Introduction to Algorithms"},
		{Id: 3, Name: "Database Systems"},
		{Id: 4, Name: "pg-go"},
	}
	db.RunInTransaction(context.Background(), func(tx *pg.Tx) error {
		var total = 0
		for _, book := range books {
			var inserted, err = tx.Model(book).Where("id=?", book.Id).SelectOrInsert()
			if err != nil {
				t.Error(book, inserted, err)
				return err
			}
			if inserted {
				total += 1
			}
		}
		t.Log(total)
		return nil
	})
}

func TestSelectBook(t *testing.T) {
	db.RunInTransaction(context.Background(), func(tx *pg.Tx) error {
		var books = make([]*schema.Book, 0)
		if err := tx.Model(&books).Select(); err != nil {
			t.Error(err)
		}
		t.Log(books)
		for _, book := range books {
			t.Log(book)
		}
		return nil
	})
}

func TestInsertOrder(t *testing.T) {
	var order = &schema.Order{
		Id: 1,
	}
	db.RunInTransaction(context.Background(), func(tx *pg.Tx) error {
		var inserted, err = tx.Model(order).Where("id=?", order.Id).SelectOrInsert()
		if err != nil {
			t.Error(inserted, err)
		}
		t.Log(inserted)
		return nil
	})
}

func TestInsertOrderItem(t *testing.T) {
	db.RunInTransaction(context.Background(), func(tx *pg.Tx) error {
		var order_item_list = []*schema.OrderItem{
			{OrderId: 1, BookId: 1, Quantity: 10},
			{OrderId: 1, BookId: 2, Quantity: 5},
			{OrderId: 1, BookId: 3, Quantity: 7},
		}
		for _, order_item := range order_item_list {
			var selected, err = tx.Model(order_item).Where("id=?", order_item.Id).SelectOrInsert()
			if err != nil {
				t.Error(err)
				return err
			}
			t.Log(order_item, selected)
		}
		return nil
	})
}

func TestSelectOrder(t *testing.T) {
	db.RunInTransaction(context.Background(), func(tx *pg.Tx) error {
		var order = &schema.OrderWithItem{
			Order: schema.Order{},
		}
		if err := tx.Model(order).
			Where("id=?", 1).
			// Relation "Items"，好像是根据反射选取的
			// 为 "items" 时候，会出现 error {model=Order does not have relation="items"}
			Relation("Items", func(q *orm.Query) (*orm.Query, error) {
				// Relation 中，指定 Items.order_id=? 时候
				// return q.Where("order_id=?", 1), nil
				// return q.Where("order_item.order_id=?", 1), nil
				return q.Where("order_item.order_id=?", 1), nil
			}).
			Select(); err != nil {
			t.Error(err)
			return err
		}
		t.Logf("%+v", order)
		if order.Id != 1 {
			t.Error(order)
		}
		return nil
	})
}
