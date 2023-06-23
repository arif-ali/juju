// Copyright 2023 Canonical Ltd.
// Licensed under the AGPLv3, see LICENCE file for details.

package schema

import "github.com/juju/juju/core/database"

// ModelDDL is used to create model databases.
func ModelDDL() []database.Delta {
	schemas := []func() database.Delta{
		changeLogSchema,
	}

	deltas := make([]database.Delta, 0, len(schemas))
	for _, fn := range schemas {
		deltas = append(deltas, fn())
	}

	return deltas
}
