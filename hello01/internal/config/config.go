// Code scaffolded by goctl. Safe to edit.
// goctl 1.9.2

package config

// import "github.com/zeromicro/go-zero/rest"

// type Config struct {
// 	rest.RestConf
// 	Address string `json:"address"`
// 	Prop	string `json:"myProp"`
// 	NoConf  string `json:"noConf,optional"`
// 	NoConf2 string `json:"noConf2,default=默认值"`
// }

type Config struct {
	Name	string       `json:"name"`
	Host    string       `json:"host"`
	Port    int          `json:"port"`	
	DataBase DataBase `json:"database"`
}

type DataBase struct {
	Host     string `json:"host"`
	Port     int    `json:"port"`
	DBName   string `json:"dbName"`
	User     string `json:"user"`
	Password string `json:"password"`
}
