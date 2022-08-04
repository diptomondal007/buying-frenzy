// Licensed to Dipto Mondal under one or more contributor
// license agreements. See the NOTICE file distributed with
// this work for additional information regarding copyright
// ownership. Dipto Mondal licenses this file to you under
// the Apache License, Version 2.0 (the "License"); you may
// not use this file except in compliance with the License.
// You may obtain a copy of the License at
//
//     http://www.apache.org/licenses/LICENSE-2.0
//
// Unless required by applicable law or agreed to in writing,
// software distributed under the License is distributed on an
// "AS IS" BASIS, WITHOUT WARRANTIES OR CONDITIONS OF ANY
// KIND, either express or implied.  See the License for the
// specific language governing permissions and limitations
// under the License.

package cmd

import (
	"log"
	"os"

	"github.com/spf13/cobra"

	"github.com/diptomondal007/buying-frenzy/app/etl"
	"github.com/diptomondal007/buying-frenzy/infrastructure/conn"
)

var (
	// data dir
	dataDir        string
	userData       string
	restaurantData string
)

// etlCmd represents the etl command
var etlCmd = &cobra.Command{
	Use:   "etl",
	Short: "etl sub command extract data from json files and load them to database.",
	Long:  `etl sub command extract data from json files and load them to database.`,
	Run: func(cmd *cobra.Command, args []string) {
		err := conn.ConnectDB()
		if err != nil {
			log.Println("db connection unsuccessful! response: ", err)
			os.Exit(1)
		}

		loader := etl.NewLoader(conn.GetDB().DB)
		t := etl.NewTransformer(loader)

		e := etl.NewETL(etl.Args{
			Directory:      dataDir,
			UserData:       userData,
			RestaurantData: restaurantData,
		}, t)

		err = e.Run()
		if err != nil {
			log.Println(err)
			// we can exit from the app.
			os.Exit(1)
		}
	},
}

func init() {
	// add flags to etl cmd
	etlCmd.Flags().StringVarP(&dataDir, "data-dir", "d", "/data", "The directory where json file data should be read from")
	_ = etlCmd.MarkFlagRequired("data-dir")
	_ = etlCmd.MarkFlagDirname("data-dir")

	etlCmd.Flags().StringVarP(&userData, "user-data", "u", "users_with_purchase_history.json", "user data file name which contains user order data")
	etlCmd.Flags().StringVarP(&restaurantData, "restaurant-data", "r", "restaurant_with_menu.json", "restaurant data file name which contains restaurant data")

	// add etl sub command to root cmd
	rootCmd.AddCommand(etlCmd)
}
