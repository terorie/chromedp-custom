// Package cachestorage provides the Chrome DevTools Protocol
// commands, types, and events for the CacheStorage domain.
//
// Generated by the cdproto-gen command.
package cachestorage

// Code generated by cdproto-gen. DO NOT EDIT.

import (
	"context"

	"bitbucket.org/ShipwrightTibi/chromecrawlingnew/cdproto/cdp"
)

// DeleteCacheParams deletes a cache.
type DeleteCacheParams struct {
	CacheID CacheID `json:"cacheId"` // Id of cache for deletion.
}

// DeleteCache deletes a cache.
//
// parameters:
//   cacheID - Id of cache for deletion.
func DeleteCache(cacheID CacheID) *DeleteCacheParams {
	return &DeleteCacheParams{
		CacheID: cacheID,
	}
}

// Do executes CacheStorage.deleteCache against the provided context.
func (p *DeleteCacheParams) Do(ctxt context.Context, h cdp.Executor) (err error) {
	return h.Execute(ctxt, CommandDeleteCache, p, nil)
}

// DeleteEntryParams deletes a cache entry.
type DeleteEntryParams struct {
	CacheID CacheID `json:"cacheId"` // Id of cache where the entry will be deleted.
	Request string  `json:"request"` // URL spec of the request.
}

// DeleteEntry deletes a cache entry.
//
// parameters:
//   cacheID - Id of cache where the entry will be deleted.
//   request - URL spec of the request.
func DeleteEntry(cacheID CacheID, request string) *DeleteEntryParams {
	return &DeleteEntryParams{
		CacheID: cacheID,
		Request: request,
	}
}

// Do executes CacheStorage.deleteEntry against the provided context.
func (p *DeleteEntryParams) Do(ctxt context.Context, h cdp.Executor) (err error) {
	return h.Execute(ctxt, CommandDeleteEntry, p, nil)
}

// RequestCacheNamesParams requests cache names.
type RequestCacheNamesParams struct {
	SecurityOrigin string `json:"securityOrigin"` // Security origin.
}

// RequestCacheNames requests cache names.
//
// parameters:
//   securityOrigin - Security origin.
func RequestCacheNames(securityOrigin string) *RequestCacheNamesParams {
	return &RequestCacheNamesParams{
		SecurityOrigin: securityOrigin,
	}
}

// RequestCacheNamesReturns return values.
type RequestCacheNamesReturns struct {
	Caches []*Cache `json:"caches,omitempty"` // Caches for the security origin.
}

// Do executes CacheStorage.requestCacheNames against the provided context.
//
// returns:
//   caches - Caches for the security origin.
func (p *RequestCacheNamesParams) Do(ctxt context.Context, h cdp.Executor) (caches []*Cache, err error) {
	// execute
	var res RequestCacheNamesReturns
	err = h.Execute(ctxt, CommandRequestCacheNames, p, &res)
	if err != nil {
		return nil, err
	}

	return res.Caches, nil
}

// RequestCachedResponseParams fetches cache entry.
type RequestCachedResponseParams struct {
	CacheID        CacheID   `json:"cacheId"`        // Id of cache that contains the entry.
	RequestURL     string    `json:"requestURL"`     // URL spec of the request.
	RequestHeaders []*Header `json:"requestHeaders"` // headers of the request.
}

// RequestCachedResponse fetches cache entry.
//
// parameters:
//   cacheID - Id of cache that contains the entry.
//   requestURL - URL spec of the request.
//   requestHeaders - headers of the request.
func RequestCachedResponse(cacheID CacheID, requestURL string, requestHeaders []*Header) *RequestCachedResponseParams {
	return &RequestCachedResponseParams{
		CacheID:        cacheID,
		RequestURL:     requestURL,
		RequestHeaders: requestHeaders,
	}
}

// RequestCachedResponseReturns return values.
type RequestCachedResponseReturns struct {
	Response *CachedResponse `json:"response,omitempty"` // Response read from the cache.
}

// Do executes CacheStorage.requestCachedResponse against the provided context.
//
// returns:
//   response - Response read from the cache.
func (p *RequestCachedResponseParams) Do(ctxt context.Context, h cdp.Executor) (response *CachedResponse, err error) {
	// execute
	var res RequestCachedResponseReturns
	err = h.Execute(ctxt, CommandRequestCachedResponse, p, &res)
	if err != nil {
		return nil, err
	}

	return res.Response, nil
}

// RequestEntriesParams requests data from cache.
type RequestEntriesParams struct {
	CacheID    CacheID `json:"cacheId"`              // ID of cache to get entries from.
	SkipCount  int64   `json:"skipCount"`            // Number of records to skip.
	PageSize   int64   `json:"pageSize"`             // Number of records to fetch.
	PathFilter string  `json:"pathFilter,omitempty"` // If present, only return the entries containing this substring in the path
}

// RequestEntries requests data from cache.
//
// parameters:
//   cacheID - ID of cache to get entries from.
//   skipCount - Number of records to skip.
//   pageSize - Number of records to fetch.
func RequestEntries(cacheID CacheID, skipCount int64, pageSize int64) *RequestEntriesParams {
	return &RequestEntriesParams{
		CacheID:   cacheID,
		SkipCount: skipCount,
		PageSize:  pageSize,
	}
}

// WithPathFilter if present, only return the entries containing this
// substring in the path.
func (p RequestEntriesParams) WithPathFilter(pathFilter string) *RequestEntriesParams {
	p.PathFilter = pathFilter
	return &p
}

// RequestEntriesReturns return values.
type RequestEntriesReturns struct {
	CacheDataEntries []*DataEntry `json:"cacheDataEntries,omitempty"` // Array of object store data entries.
	ReturnCount      float64      `json:"returnCount,omitempty"`      // Count of returned entries from this storage. If pathFilter is empty, it is the count of all entries from this storage.
}

// Do executes CacheStorage.requestEntries against the provided context.
//
// returns:
//   cacheDataEntries - Array of object store data entries.
//   returnCount - Count of returned entries from this storage. If pathFilter is empty, it is the count of all entries from this storage.
func (p *RequestEntriesParams) Do(ctxt context.Context, h cdp.Executor) (cacheDataEntries []*DataEntry, returnCount float64, err error) {
	// execute
	var res RequestEntriesReturns
	err = h.Execute(ctxt, CommandRequestEntries, p, &res)
	if err != nil {
		return nil, 0, err
	}

	return res.CacheDataEntries, res.ReturnCount, nil
}

// Command names.
const (
	CommandDeleteCache           = "CacheStorage.deleteCache"
	CommandDeleteEntry           = "CacheStorage.deleteEntry"
	CommandRequestCacheNames     = "CacheStorage.requestCacheNames"
	CommandRequestCachedResponse = "CacheStorage.requestCachedResponse"
	CommandRequestEntries        = "CacheStorage.requestEntries"
)
