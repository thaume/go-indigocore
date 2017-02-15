// Copyright 2017 Stratumn SAS. All rights reserved.
//
// This Source Code Form is subject to the terms of the Mozilla Public
// License, v. 2.0. If a copy of the MPL was not distributed with this
// file, You can obtain one at http://mozilla.org/MPL/2.0/.

package storehttp

import (
	"net/http"
	"strconv"
	"strings"

	"github.com/stratumn/go/store"
	"github.com/stratumn/go/types"
)

func parseFilter(r *http.Request) (*store.Filter, error) {
	pagination, err := parsePagination(r)
	if err != nil {
		return nil, err
	}

	var (
		mapID           = r.URL.Query().Get("mapId")
		prevLinkHashStr = r.URL.Query().Get("prevLinkHash")
		tagsStr         = r.URL.Query().Get("tags")
		prevLinkHash    *types.Bytes32
		tags            []string
	)

	if prevLinkHashStr != "" {
		prevLinkHash, err = types.NewBytes32FromString(prevLinkHashStr)
		if err != nil {
			return nil, newErrPrevLinkHash("")
		}
	}

	if tagsStr != "" {
		spaceTags := strings.Split(tagsStr, " ")
		for _, t := range spaceTags {
			tags = append(tags, strings.Split(t, "+")...)
		}
	}

	return &store.Filter{
		Pagination:   *pagination,
		MapID:        mapID,
		PrevLinkHash: prevLinkHash,
		Tags:         tags,
	}, nil
}

func parsePagination(r *http.Request) (*store.Pagination, error) {
	var err error

	offsetstr := r.URL.Query().Get("offset")
	offset := 0
	if offsetstr != "" {
		if offset, err = strconv.Atoi(offsetstr); err != nil || offset < 0 {
			return nil, newErrOffset("")
		}
	}

	limitstr := r.URL.Query().Get("limit")
	limit := store.DefaultLimit
	if limitstr != "" {
		if limit, err = strconv.Atoi(limitstr); err != nil || limit < 0 {
			return nil, newErrLimit("")
		}
	}

	if limit > store.MaxLimit {
		return nil, newErrLimit("")
	}

	return &store.Pagination{
		Offset: offset,
		Limit:  limit,
	}, nil
}
