/*                          _       _
 *__      _____  __ ___   ___  __ _| |_ ___
 *\ \ /\ / / _ \/ _` \ \ / / |/ _` | __/ _ \
 * \ V  V /  __/ (_| |\ V /| | (_| | ||  __/
 *  \_/\_/ \___|\__,_| \_/ |_|\__,_|\__\___|
 *
 * Copyright © 2016 Weaviate. All rights reserved.
 * LICENSE: https://github.com/weaviate/weaviate/blob/master/LICENSE
 * AUTHOR: Bob van Luijt (bob@weaviate.com)
 * See www.weaviate.com for details
 * Contact: @weaviate_iot / yourfriends@weaviate.com
 */

package dgraph

import ()

// TotalResultsResult is the root of a single node with the result count
type TotalResultsResult struct {
	Root Count `dgraph:"totalResults"`
}

// Count is the struct for the counter
type Count struct {
	Count int64 `dgraph:"count"`
}

// NodeIDResult is the struct for getting the thing with an UUID
type NodeIDResult struct {
	Root NodeID `dgraph:"node"`
}

// NodeID reversed thing-id
type NodeID struct {
	UUID string   `dgraph:"uuid"`
	ID   uint64   `dgraph:"_uid_"`
	Node UUIDNode `dgraph:"~id"`
}

// UUIDNode by Thing ID
type UUIDNode struct {
	ID uint64 `dgraph:"_uid_"`
}