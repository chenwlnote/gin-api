package database

import (
	"chenwlnote.gin-api/app/pkg/logger"
	"chenwlnote.gin-api/app/provider/app/config"
	"database/sql"
	_ "github.com/go-sql-driver/mysql"
	"math/rand"
	"strconv"
	"time"
)

var dbReadDsn map[string][]*sql.DB
var dbWriteDsn map[string][]*sql.DB

func init() {
	initMysqlReadDsn()
	initMysqlWriteDsn()
	checkMysqlConnection()
}

func checkMysqlConnection() {
	if dbReadDsn == nil || dbWriteDsn == nil {
		panic("db 初始化失败")
	}
}

type MysqlPool struct {
}

func NewMysqlPool() *MysqlPool {
	return &MysqlPool{}
}

func (mysqlPool *MysqlPool) Read(dbAlias string) *sql.DB {
	dsnArr := getMysqlReadDsn(dbAlias)
	if dsnArr == nil || len(dsnArr) == 0 {
		return nil
	}
	if len(dsnArr) == 1 {
		return dsnArr[0]
	}
	index := rand.Intn(len(dsnArr) - 1)
	return dsnArr[index]
}

func (mysqlPool *MysqlPool) Write(dbAlias string) *sql.DB {
	dsnArr := getMysqlWriteReadDsn(dbAlias)
	if dsnArr == nil || len(dsnArr) == 0 {
		return nil
	}
	if len(dsnArr) == 1 {
		return dsnArr[0]
	}
	index := rand.Intn(len(dsnArr) - 1)
	return dsnArr[index]
}

func getMysqlReadDsn(dbAlias string) []*sql.DB {
	if len(dbReadDsn) == 0 {
		logger.Fatal("db read dsn is null", map[string]interface{}{"db_alias": dbAlias})
		return nil
	}
	if dsn, ok := dbReadDsn[dbAlias]; ok {
		return dsn
	} else {
		logger.Fatal("db read dsn miss ", map[string]interface{}{"db_alias": dbAlias})
		return nil
	}
}

func getMysqlWriteReadDsn(dbAlias string) []*sql.DB {
	if len(dbWriteDsn) == 0 {
		logger.Fatal("db write dsn is null", map[string]interface{}{"db_alias": dbAlias})
	}
	if dsn, ok := dbWriteDsn[dbAlias]; ok {
		return dsn
	} else {
		logger.Fatal("db write dsn miss ", map[string]interface{}{"db_alias": dbAlias})
		return nil
	}
}

func initMysqlReadDsn() {
	for _, mysqlConfig := range config.Get().Database.Mysql.Read {
		for _, dbConfig := range mysqlConfig.DatabaseConfig {

			db, err := sql.Open("mysql", makeMysqlDsn(mysqlConfig, dbConfig))
			if err != nil {
				logger.Fatal("db open connection error ", map[string]interface{}{"err": err})
				continue
			}
			if db == nil {
				logger.Fatal("db read connection is nil ", map[string]interface{}{"db_name": dbConfig.Name})
				continue
			}
			db.SetConnMaxIdleTime(mysqlConfig.ConnMaxIdleTime * time.Second)
			db.SetConnMaxLifetime(mysqlConfig.ConnMaxLifeTime * time.Second)
			db.SetMaxIdleConns(mysqlConfig.MaxIdleConns)
			db.SetMaxOpenConns(mysqlConfig.MaxOpenConns)
			if dbReadDsn == nil {
				dbReadDsn = make(map[string][]*sql.DB)
			}
			dbReadDsn[dbConfig.Alias] = append(dbReadDsn[dbConfig.Alias], db)
		}
	}
}

func makeMysqlDsn(mysqlConfig config.MysqlInfo, dbConfig config.DatabaseConfig) string {
	dsn := mysqlConfig.Username + ":" + mysqlConfig.Password +
		"@tcp(" + mysqlConfig.Host + ":" +
		strconv.Itoa(mysqlConfig.Port) + ")/" + dbConfig.Name +
		"?charset=" + dbConfig.Charset + "&parseTime=True"
	return dsn
}

func initMysqlWriteDsn() {
	for _, mysqlConfig := range config.Get().Database.Mysql.Write {
		for _, dbConfig := range mysqlConfig.DatabaseConfig {
			db, err := sql.Open("mysql", makeMysqlDsn(mysqlConfig, dbConfig))
			if err != nil {
				logger.Fatal("db open connection error ", map[string]interface{}{"err": err})
				continue
			}
			if db == nil {
				logger.Fatal("db write connection is nil ", map[string]interface{}{"db_name": dbConfig.Name})
				continue
			}
			db.SetConnMaxIdleTime(mysqlConfig.ConnMaxIdleTime * time.Second)
			db.SetConnMaxLifetime(mysqlConfig.ConnMaxLifeTime * time.Second)
			db.SetMaxIdleConns(mysqlConfig.MaxIdleConns)
			db.SetMaxOpenConns(mysqlConfig.MaxOpenConns)
			if dbWriteDsn == nil {
				dbWriteDsn = make(map[string][]*sql.DB)
			}
			dbWriteDsn[dbConfig.Alias] = append(dbWriteDsn[dbConfig.Alias], db)
			dbReadDsn[dbConfig.Alias] = append(dbReadDsn[dbConfig.Alias], db)
		}
	}
}
