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

// registerCmd represents the register command
var registerCmd = &cobra.Command{
	Use:   "register",
	Short: "A brief description of your command",
	Long: `to register a new account`,
	Run: func(cmd *cobra.Command, args []string) {
		fmt.Println("register called")
		username, _ := cmd.Flags().GetString("user")
		password, _ := cmd.Flags().GetString("password")
		email, _ := cmd.Flags().GetString("email")
		phone, _ := cmd.Flags().GetString("phone")
		userList := Entity.GetUsers()
		for i := 0; i < len(userList); i++ {
			if userList[i].Username == username {
				fmt.Println("Username exist!!!!")
				return
			}
		}
		fmt.Println("Register successfully!!!\nNow you could login in with username ",username)
		newUser := Entity.User{Username: username, Password: password, Email: email, Phone: phone}
		userList = append(userList, newUser)
		Entity.UpdateUsers(userList)
		
	},
}

func init() {
	rootCmd.AddCommand(registerCmd)

	// Here you will define your flags and configuration settings.

	// Cobra supports Persistent Flags which will work for this command
	// and all subcommands, e.g.:
	// registerCmd.PersistentFlags().String("foo", "", "A help for foo")
	registerCmd.Flags().StringP("user", "u", "", "username")
	registerCmd.Flags().StringP("password", "p", "", "password")
	registerCmd.Flags().StringP("email", "e", "", "email")
	registerCmd.Flags().StringP("phone", "o", "", "phone")
	// Cobra supports local flags which will only run when this command
	// is called directly, e.g.:
	// registerCmd.Flags().BoolP("toggle", "t", false, "Help message for toggle")
}
