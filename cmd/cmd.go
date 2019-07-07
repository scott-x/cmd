/*
* @Author: sottxiong
* @Date:   2019-07-07 16:27:42
* @Last Modified by:   sottxiong
* @Last Modified time: 2019-07-07 22:41:26
*/
package main

import (
  "github.com/scott-x/go-commander/cmd"
  "fmt"
)
func main(){
   cmd.AddQuestion("name","What's your name ? ","Please input correct name: ","[a-z]+")
   cmd.AddQuestion("age","What's your age ? ","Please input correct age: ","[0-9]{2}")
   a :=cmd.Exec()
   fmt.Println(a)
}   
