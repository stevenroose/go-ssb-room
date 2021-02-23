// SPDX-License-Identifier: MIT

package sqlite

import (
	"database/sql"

	"github.com/ssb-ngi-pointer/go-ssb-room/admindb"
)

// make sure to implement interfaces correctly
var _ admindb.AliasService = (*Aliases)(nil)

type Aliases struct {
	db *sql.DB
}