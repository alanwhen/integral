package sysinit

import (
	//_ 导入models初始化
	_ "github.com/alanwhen/education-mini/models"
	"github.com/astaxie/beego"
	"github.com/astaxie/beego/orm"
	_ "github.com/go-sql-driver/mysql"
)

func InitDatabase() {
	dbType := beego.AppConfig.String("db_type")

	dbAlias := beego.AppConfig.String(dbType + "::db_alias")

	dbName := beego.AppConfig.String(dbType + "::db_name")

	dbUser := beego.AppConfig.String(dbType + "::db_user")

	dbPwd := beego.AppConfig.String(dbType + "::db_pwd")

	dbHost := beego.AppConfig.String(dbType + "::db_host")

	dbPort := beego.AppConfig.String(dbType + "::db_port")

	switch dbType {
	case "mysql":
		dbCharset := beego.AppConfig.String(dbType + "::db_charset")
		orm.RegisterDataBase(dbAlias, dbType, dbUser+":"+dbPwd+"@tcp("+dbHost+":"+dbPort+")/"+dbName+"?charset="+dbCharset+"&loc=Asia%2FShanghai", 30)
	}

	isDev := (beego.AppConfig.String("runMode") == "dev")

	orm.RunSyncdb("default", false, isDev)
	if isDev {
		orm.Debug = isDev
	}
}
