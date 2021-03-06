// Copyright 2012-2015 Oliver Eilhard. All rights reserved.
// Use of this source code is governed by a MIT-license.
// See http://olivere.mit-license.org/license.txt for details.

package elastic

import (
	"context"

	"gopkg.in/olivere/elastic.v2/uritemplates"
)

type DeleteIndexService struct {
	client *Client
	index  string
}

func NewDeleteIndexService(client *Client) *DeleteIndexService {
	builder := &DeleteIndexService{
		client: client,
	}
	return builder
}

func (b *DeleteIndexService) Index(index string) *DeleteIndexService {
	b.index = index
	return b
}

// Do runs DoC() with default context.
func (b *DeleteIndexService) Do() (*DeleteIndexResult, error) {
	return b.DoC(nil)
}

func (b *DeleteIndexService) DoC(ctx context.Context) (*DeleteIndexResult, error) {
	// Build url
	path, err := uritemplates.Expand("/{index}/", map[string]string{
		"index": b.index,
	})
	if err != nil {
		return nil, err
	}

	// Get response
	res, err := b.client.PerformRequestC(ctx, "DELETE", path, nil, nil)
	if err != nil {
		return nil, err
	}

	// Return result
	ret := new(DeleteIndexResult)
	if err := b.client.decoder.Decode(res.Body, ret); err != nil {
		return nil, err
	}
	return ret, nil
}

// -- Result of a delete index request.

type DeleteIndexResult struct {
	Acknowledged bool `json:"acknowledged"`
}
