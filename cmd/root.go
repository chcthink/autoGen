package cmd

import (
	"fmt"
	"os"
	"strings"

	"github.com/spf13/cobra"
)

var (
	port      int
	tableName string
	username  string
	password  string
	ip        string
	dbname    string
	tags      string
	admin     bool
	model     bool
	adminList bool
)

var rootCmd = &cobra.Command{
	Use: "gen",
	Run: func(cmd *cobra.Command, args []string) {
		if dbname == "" {
			fmt.Println("no such flag --password")
			return
		}
		if tableName == "" {
			fmt.Println("no such flag --password")
			return
		}
		// init db
		var err error
		mysql, err = NewMysql(username, dbname, ip, password, port)
		if err != nil {
			fmt.Println(err)
			return
		}
		// init params
		tables := strings.Split(tableName, ",")
		for _, table := range tables {
			// create file
			path := make(map[string]string)
			tag := ""
			if model {
				tag = "backend/dao/" + tags + "d/"
				path[tag] = tag + table + ".go"
			}
			if admin {
				tag = "admin/src/api/" + tags + "/"
				path[tag] = tag + table + ".ts"
				path["admin/src/api/"+tags+"/model/"] = "admin/src/api/" + tags + "/model/" + table + ".ts"
			}
			if adminList {
				tag = "gen_temp/" + tags + "/"
				path[tag] = tag + table + ".ts"
			}
			for k, v := range path {
				_ = os.MkdirAll(k, 0777)
				f, err := os.Create(v)
				if err != nil {
					return
				}

				insertStr := genDAO(table)
				if admin {
					if strings.Contains(k, "/model/") {
						insertStr = genAPI(table)
					} else {
						insertStr = genAPIMethod(table)
					}
				}
				if adminList {
					insertStr, _ = generateStruct(table, ADMINLIST)
				}
				if err != nil {
					fmt.Println(err)
					_ = f.Close()
					_ = os.Remove(v)
					return
				}
				_, err = f.WriteString(insertStr)
				if err != nil {
					fmt.Println(err)
					_ = f.Close()
					_ = os.Remove(v)
					return
				}
				err = f.Sync()
				if err != nil {
					fmt.Println(err)
					_ = f.Close()
					_ = os.Remove(insertStr)
					return
				}
			}

		}

	},
}

func init() {
	rootCmd.PersistentFlags().IntVar(&port, "port", 3306, "mysql port")
	rootCmd.PersistentFlags().StringVarP(&username, "username", "u", "root", "mysql username")
	rootCmd.PersistentFlags().StringVarP(&password, "password", "p", "123456ab", "mysql password")
	rootCmd.PersistentFlags().StringVarP(&tableName, "tableName", "t", "", "table names,use ',' to split")
	rootCmd.PersistentFlags().StringVarP(&ip, "ip", "i", "127.0.0.1", "ip")
	rootCmd.PersistentFlags().StringVarP(&dbname, "dbname", "d", "", "database name")
	rootCmd.PersistentFlags().StringVarP(&tags, "tags", "g", "", "tags")
	rootCmd.PersistentFlags().BoolVarP(&admin, "admin", "a", false, "is admin")
	rootCmd.PersistentFlags().BoolVarP(&model, "model", "m", false, "dao model")
	rootCmd.PersistentFlags().BoolVarP(&adminList, "admin_list", "A", false, "admin list")
}

func Execute() {
	if err := rootCmd.Execute(); err != nil {
		fmt.Println(err)
		os.Exit(1)
	}
}
