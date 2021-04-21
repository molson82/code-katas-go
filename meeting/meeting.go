package meeting

import (
	"fmt"
	"sort"
	"strings"
)

type friend struct {
  FirstName string
  LastName string
}

func Meeting(s string) string {
  names := strings.Split(s, ";")
  var friendList []friend
  for _,v := range names {
    firstName := strings.Split(v, ":")[0]
    lastName := strings.Split(v, ":")[1]
    friendList = append(friendList, friend{strings.ToLower(firstName), strings.ToLower(lastName)})
  }
  sort.Slice(friendList, func(i, j int) bool {
    if friendList[i].LastName == friendList[j].LastName {
      return friendList[i].FirstName < friendList[j].FirstName
    }
    return friendList[i].LastName < friendList[j].LastName
  })
  var finalStr string
  for _,v := range friendList {
    finalStr = finalStr + fmt.Sprintf("(%s, %s)", v.LastName, v.FirstName) 
  }
  return strings.ToUpper(finalStr)
}
