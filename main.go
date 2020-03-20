// Package main App API.
//
// token: 用户登录后获取的访问凭证，有使用期限
//
// Terms Of Service:
//
//     Schemes: http
//     Host: localhost
//     BasePath: /v1
//     Version: 0.0.1
//
//     Consumes:
//     - application/json
//
//     Produces:
//     - application/json
//
//     Security:
//     - access-authorized
//
//     SecurityDefinitions:
//     access-authorized:
//          type: apiKey
//          name: Access-Token
//          in: header
//          description: value is token that  get from user login
//
//     Extensions:
//     x-meta-value: value
//     x-meta-array:
//       - value1
//       - value2
//     x-meta-array-obj:
//       - name: obj
//         value: field
//
// swagger:meta
package main

import (
	"best.me/cmd"
)

func main() {
	cmd.Execute()
}
