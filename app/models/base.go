//****************************************************/
// Author      : ningzhong.zeng
// Revision    : 2016-01-26 21:49:47
//****************************************************/
package models

import (
	"database/sql"
	"fmt"
	"log"
	"math/rand"
	"reflect"
	"time"

	"github.com/go-gorp/gorp"
	"github.com/revel/revel"
	_ "github.com/ziutek/mymysql/godrv"
)

type Pre interface {
	PreInsert(_ gorp.SqlExecutor) error
	PreUpdate(_ gorp.SqlExecutor) error
}

type Base struct {
	Txn       *gorp.Transaction `db:"-"`
	GmtMotify time.Time
	GmtCreate time.Time
}

func (b *Base) PerInsert(_ gorp.SqlExecutor) error {
	return b.PreI()
}

func (b *Base) PerUpdate(_ gorp.SqlExecutor) error {
	return b.PreU()
}

func (b *Base) PreI() error {
	b.GmtCreate = time.Now()
	b.GmtMotify = time.Now()
	return nil
}

func (b *Base) PreU() error {
	b.GmtMotify = time.Now()
	return nil
}

var (
	dbmap *gorp.DbMap
)

func DBMap() *gorp.DbMap {
	return dbmap
}

func InitDB() *gorp.DbMap {
	// connect to db using standard Go database/sql API
	// use whatever database/sql driver you wish
	db, err := sql.Open("mymysql", "tcp:localhost:3306*gorp/admin/zhong")
	CheckErr(err, "mysql.Open.failed")

	// construct a gorp DbMap
	dbmap = &gorp.DbMap{Db: db, Dialect: gorp.MySQLDialect{"InnoDB", "UTF8"}}

	// add a table, setting the table name to 'posts' and
	// specifying that the Id property is an auto incrementing PK
	// dbmap.AddTableWithName(Post{}, "posts").SetKeys(true, "Id")
	fmt.Sprintf("random %d", rand.Int63())
	setColumnSizes := func(t *gorp.TableMap, colSizes map[string]int) {
		for col, size := range colSizes {
			t.ColMap(col).MaxSize = size
		}
	}
	userTable := dbmap.AddTableWithName(User{}, "s_user")
	userTable.SetKeys(false, "UserId")
	userTable.ColMap("Username").SetUnique(true).SetNotNull(true)
	userTable.ColMap("Password").SetNotNull(true)
	setColumnSizes(userTable, map[string]int{
		"UserId":   50,
		"Username": 50,
		"Password": 50,
	})

	dbmap.TraceOn("[gorp]", revel.INFO)
	// create the table. in a production system you'd generally
	// use a migration tool, or create the tables via scripts
	err = dbmap.CreateTablesIfNotExists()
	CheckErr(err, "Create tables failed")

	return dbmap
}

// base model common method
func Select(v interface{}) {
}

func Insert(v interface{}) {
	// revel.INFO.Println("Insert Num:", len(v))
	// user := make([]User, len(v))
	// for _, item := range v {
	// if model, ok := item.(User); ok {
	// model.UserId = RandomString("u", 4)
	// user[0] = model
	// }
	// }
	revel.INFO.Println(dbmap)

	ptrv := reflect.ValueOf(v)
	revel.INFO.Println("kind: ", ptrv.Kind(), ",ptr: ", reflect.Ptr)
	// if err := dbmap.Insert(&list); err != nil {
	// revel.ERROR.Println("Insert Err:", err)
	// }
}

func Update(v ...interface{}) {
	dbmap.Update(v)
}

func Delete(v ...interface{}) {
	dbmap.Delete(v)
}

const alphanum = "abcdefghijklmnopqrstuvwxyz"

func RandomString(pixff string, strlen int) string {
	rand.Seed(time.Now().UTC().UnixNano())
	result := make([]byte, strlen)
	for i := 0; i < strlen; i++ {
		result[i] = alphanum[rand.Intn(len(alphanum))]
	}
	return time.Now().Format("20151212010203") + "-" + pixff + "-" + string(result)
}

// base model transferion
func (b *Base) Begin() {
	txn, err := dbmap.Begin()
	if err != nil {
		panic(err)
	}
	b.Txn = txn
}

func (b *Base) Commit() {
	if b.Txn == nil {
		return
	}
	if err := b.Txn.Commit(); err != nil && err != sql.ErrTxDone {
		panic(err)
	}
	b.Txn = nil
}

func (b *Base) Rollback() {
	if b.Txn == nil {
		return
	}
	if err := b.Txn.Rollback(); err != nil && err != sql.ErrTxDone {
		panic(err)
	}
	b.Txn = nil
}

func CheckErr(err error, msg string) (r bool) {
	r = true
	if err != nil {
		// log.Fatalln(msg, err)
		revel.ERROR.Println(msg, err)
		r = false
	}
	return r
}

func CheckPanic(err error, msg string) {
	if err != nil {
		log.Fatalln(msg, err)
	}
}
