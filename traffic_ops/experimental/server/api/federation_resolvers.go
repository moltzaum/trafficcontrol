
// Licensed under the Apache License, Version 2.0 (the "License");
// you may not use this file except in compliance with the License.
// You may obtain a copy of the License at

// http://www.apache.org/licenses/LICENSE-2.0

// Unless required by applicable law or agreed to in writing, software
// distributed under the License is distributed on an "AS IS" BASIS,
// WITHOUT WARRANTIES OR CONDITIONS OF ANY KIND, either express or implied.
// See the License for the specific language governing permissions and
// limitations under the License.

// This file was initially generated by gen_to_start.go (add link), as a start
// of the Traffic Ops golang data model

package api

import (
	"encoding/json"
	_ "github.com/apache/trafficcontrol/traffic_ops/experimental/server/output_format" // needed for swagger
	"github.com/jmoiron/sqlx"
	"log"
	"time"
)

type FederationResolvers struct {
	Id        int64                    `db:"id" json:"id"`
	IpAddress string                   `db:"ip_address" json:"ipAddress"`
	Type      string                   `db:"type" json:"type"`
	CreatedAt time.Time                `db:"created_at" json:"createdAt"`
	Links     FederationResolversLinks `json:"_links" db:-`
}

type FederationResolversLinks struct {
	Self string `db:"self" json:"_self"`
}

// @Title getFederationResolversById
// @Description retrieves the federation_resolvers information for a certain id
// @Accept  application/json
// @Param   id              path    int     false        "The row id"
// @Success 200 {array}    FederationResolvers
// @Resource /api/2.0
// @Router /api/2.0/federation_resolvers/{id} [get]
func getFederationResolver(id int64, db *sqlx.DB) (interface{}, error) {
	ret := []FederationResolvers{}
	arg := FederationResolvers{}
	arg.Id = id
	queryStr := "select *, concat('" + API_PATH + "federation_resolvers/', id) as self"
	queryStr += " from federation_resolvers WHERE id=:id"
	nstmt, err := db.PrepareNamed(queryStr)
	err = nstmt.Select(&ret, arg)
	if err != nil {
		log.Println(err)
		return nil, err
	}
	nstmt.Close()
	return ret, nil
}

// @Title getFederationResolverss
// @Description retrieves the federation_resolvers
// @Accept  application/json
// @Success 200 {array}    FederationResolvers
// @Resource /api/2.0
// @Router /api/2.0/federation_resolvers [get]
func getFederationResolvers(db *sqlx.DB) (interface{}, error) {
	ret := []FederationResolvers{}
	queryStr := "select *, concat('" + API_PATH + "federation_resolvers/', id) as self"
	queryStr += " from federation_resolvers"
	err := db.Select(&ret, queryStr)
	if err != nil {
		log.Println(err)
		return nil, err
	}
	return ret, nil
}

// @Title postFederationResolvers
// @Description enter a new federation_resolvers
// @Accept  application/json
// @Param                 Body body     FederationResolvers   true "FederationResolvers object that should be added to the table"
// @Success 200 {object}    output_format.ApiWrapper
// @Resource /api/2.0
// @Router /api/2.0/federation_resolvers [post]
func postFederationResolver(payload []byte, db *sqlx.DB) (interface{}, error) {
	var v FederationResolvers
	err := json.Unmarshal(payload, &v)
	if err != nil {
		log.Println(err)
		return nil, err
	}
	sqlString := "INSERT INTO federation_resolvers("
	sqlString += "ip_address"
	sqlString += ",type"
	sqlString += ",created_at"
	sqlString += ") VALUES ("
	sqlString += ":ip_address"
	sqlString += ",:type"
	sqlString += ",:created_at"
	sqlString += ")"
	result, err := db.NamedExec(sqlString, v)
	if err != nil {
		log.Println(err)
		return nil, err
	}
	return result, err
}

// @Title putFederationResolvers
// @Description modify an existing federation_resolversentry
// @Accept  application/json
// @Param   id              path    int     true        "The row id"
// @Param                 Body body     FederationResolvers   true "FederationResolvers object that should be added to the table"
// @Success 200 {object}    output_format.ApiWrapper
// @Resource /api/2.0
// @Router /api/2.0/federation_resolvers/{id}  [put]
func putFederationResolver(id int64, payload []byte, db *sqlx.DB) (interface{}, error) {
	var arg FederationResolvers
	err := json.Unmarshal(payload, &arg)
	arg.Id = id
	if err != nil {
		log.Println(err)
		return nil, err
	}
	sqlString := "UPDATE federation_resolvers SET "
	sqlString += "ip_address = :ip_address"
	sqlString += ",type = :type"
	sqlString += ",created_at = :created_at"
	sqlString += " WHERE id=:id"
	result, err := db.NamedExec(sqlString, arg)
	if err != nil {
		log.Println(err)
		return nil, err
	}
	return result, err
}

// @Title delFederationResolversById
// @Description deletes federation_resolvers information for a certain id
// @Accept  application/json
// @Param   id              path    int     false        "The row id"
// @Success 200 {array}    FederationResolvers
// @Resource /api/2.0
// @Router /api/2.0/federation_resolvers/{id} [delete]
func delFederationResolver(id int64, db *sqlx.DB) (interface{}, error) {
	arg := FederationResolvers{}
	arg.Id = id
	result, err := db.NamedExec("DELETE FROM federation_resolvers WHERE id=:id", arg)
	if err != nil {
		log.Println(err)
		return nil, err
	}
	return result, err
}
