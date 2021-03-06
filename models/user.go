/*
   Copyright (c) 2016 VMware, Inc. All Rights Reserved.
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

package models

import (
	"time"
)

// User holds the details of a user.
type User struct {
	UserID       int    `orm:"column(user_id)" json:"UserId"`
	Username     string `orm:"column(username)" json:"username"`
	Email        string `orm:"column(email)" json:"email"`
	Password     string `orm:"column(password)" json:"password"`
	Realname     string `orm:"column(realname)" json:"realname"`
	Comment      string `orm:"column(comment)" json:"comment"`
	Deleted      int    `orm:"column(deleted)"`
	Rolename     string
	RoleID       int `json:"RoleId"`
	RoleList     []Role
	HasAdminRole int    `orm:"column(sysadmin_flag)"`
	ResetUUID    string `orm:"column(reset_uuid)" json:"ResetUuid"`
	Salt         string `orm:"column(salt)"`

	CreationTime time.Time `orm:"creation_time" json:"creation_time"`
	UpdateTime   time.Time `orm:"update_time" json:"update_time"`
}
