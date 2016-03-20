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
	conf, err := revel.LoadConfig("app.conf")
	CheckErr(err, "Load app.conf failed")
	mysqluser, _ := conf.String("mysql.user")
	mysqlpass, _ := conf.String("mysql.pass")
	mysqlurls, _ := conf.String("mysql.urls")
	mysqldb, _ := conf.String("mysql.db")

	db, err := sql.Open("mymysql",
		fmt.Sprintf("tcp:%s:3306*%s/%s/%s",
			mysqlurls,
			mysqldb,
			mysqluser,
			mysqlpass,
		),
	)
	CheckErr(err, "mysql.Open.failed")

	// construct a gorp DbMap
	dbmap = &gorp.DbMap{Db: db, Dialect: gorp.MySQLDialect{"InnoDB", "UTF8"}}

	setColumnSizes := func(t *gorp.TableMap, colSizes map[string]int) {
		for col, size := range colSizes {
			t.ColMap(col).MaxSize = size
		}
	}
	userTable := dbmap.AddTableWithName(User{}, "s_user")
	userTable.SetKeys(false, "UserId")
	userTable.ColMap("Username").SetUnique(true).SetNotNull(true)
	userTable.ColMap("Password").SetNotNull(true)
	userTable.ColMap("Role").SetNotNull(true)
	setColumnSizes(userTable, map[string]int{
		"UserId":   50,
		"Username": 50,
		"Password": 50,
		"Role":     1,
	})

	agentAable := dbmap.AddTableWithName(Agent{}, "o_agent")
	agentAable.SetKeys(false, "AgentId")
	agentAable.ColMap("Host").SetNotNull(true)
	agentAable.ColMap("Port").SetNotNull(true)
	agentAable.ColMap("LoginName").SetNotNull(true)
	agentAable.ColMap("LoginPass").SetNotNull(true)
	setColumnSizes(agentAable, map[string]int{
		"Host":         100,
		"Port":         4,
		"LoginName":    30,
		"LoginPass":    100,
		"InitShellCmd": 2000,
		"Memo":         500,
		"UserId":       45,
	})

	cmdAable := dbmap.AddTableWithName(Cmd{}, "o_cmd")
	cmdAable.SetKeys(false, "CmdId")
	cmdAable.ColMap("Shell").SetNotNull(true)
	setColumnSizes(cmdAable, map[string]int{
		"Shell":   100,
		"CmdId":   45,
		"AgentId": 45,
	})

	serverAable := dbmap.AddTableWithName(Server{}, "p_server")
	serverAable.SetKeys(false, "ServerId")
	serverAable.ColMap("Theme").SetNotNull(true)
	serverAable.ColMap("Content").SetNotNull(true)
	setColumnSizes(serverAable, map[string]int{
		"Theme":   100,
		"Content": 30000,
	})

	dbmap.TraceOn("[gorp]", revel.INFO)
	err = dbmap.CreateTablesIfNotExists()
	CheckErr(err, "Create tables failed")

	return dbmap
}

// 执行sql
func Exec(sql string, v interface{}) bool {
	r, err := dbmap.Exec(sql, v)
	CheckErr(err, "Exec ")
	if row, _ := r.RowsAffected(); row > 0 {
		return true
	} else {
		return false
	}
}

// base model common method
const alphanum = "abcdefghijklmnopqrstuvwxyz"

func RandomString(pixff string, strlen int) string {
	rand.Seed(time.Now().UTC().UnixNano())
	result := make([]byte, strlen)
	for i := 0; i < strlen; i++ {
		result[i] = alphanum[rand.Intn(len(alphanum))]
	}
	return time.Now().Format("20151212010203") + string(result) + pixff
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

func CheckErrSQLSelectList(vs []interface{}, err error) (r bool) {
	return CheckErr(err, "SelectList VS")
}

func CheckErrSQLSelectOne(err error) (r bool) {
	return CheckErr(err, "SelectOne")
}

func CheckPanic(err error, msg string) {
	if err != nil {
		log.Fatalln(msg, err)
	}
}

func ValidationAddErr(v *revel.Validation, message string) {
	v.Errors = append(v.Errors,
		&revel.ValidationError{
			Message: message,
			Key:     "Err",
		},
	)
}
