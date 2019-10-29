/*
Copyright © 2019 NAME HERE <EMAIL ADDRESS>

Licensed under the Apache License, Version 2.0 (the "License");
you may not use this file except in compliance with the License.
You may obtain a copy of the License at

    http://www.apache.org/licenses/LICENSE-2.0

Unless required by applicable law or agreed to in writing, software
distributed under the License is distributed on an "AS IS" BASIS,
WITHOUT WARRANTIES OR CONDITIONS OF ANY KIND, either express or implied.
See the License for the specific language governing permissions and
limitations under the License.
*/
package cmd

import (
	"fmt"
	"Agenda/Entity"
	"github.com/spf13/cobra"
)

// loginCmd represents the login command
var loginCmd = &cobra.Command{
	Use:   "login",
	Short: "A brief description of your command",
	Long: `agenda login -n lwh -p 123`,
	Run: func(cmd *cobra.Command, args []string) {
		//fmt.Println("login called")
		username,_  := cmd.Flags().GetString("username")
		password,_  := cmd.Flags().GetString("password")
		userlist := Entity.GetUsers()
		for i := 0;i < len(userlist);i ++{
			if userlist[i].Username == username{
				if userlist[i].Password == password{
					fmt.Println("Login Successfully!")
					return 
				}
				fmt.Println("Password error!\n Please try again\n")
				return
			}
		}
		fmt.Println("Username don't exist!Please register first\n.")
	},
}

func init() {
	rootCmd.AddCommand(loginCmd)

	// Here you will define your flags and configuration settings.
	loginCmd.Flags().StringP("username","u","","username")
	loginCmd.Flags().StringP("password","p","","password")
	// Cobra supports Persistent Flags which will work for this command
	// and all subcommands, e.g.:
	// loginCmd.PersistentFlags().String("foo", "", "A help for foo")

	// Cobra supports local flags which will only run when this command
	// is called directly, e.g.:
	// loginCmd.Flags().BoolP("toggle", "t", false, "Help message for toggle")
}
