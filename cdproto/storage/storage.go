// Package storage provides the Chrome DevTools Protocol
// commands, types, and events for the Storage domain.
//
// Generated by the cdproto-gen command.
package storage

// Code generated by cdproto-gen. DO NOT EDIT.

import (
	"context"

	"bitbucket.org/ShipwrightTibi/chromecrawlingnew/cdproto/cdp"
)

// ClearDataForOriginParams clears storage for origin.
type ClearDataForOriginParams struct {
	Origin       string `json:"origin"`       // Security origin.
	StorageTypes string `json:"storageTypes"` // Comma separated list of StorageType to clear.
}

// ClearDataForOrigin clears storage for origin.
//
// parameters:
//   origin - Security origin.
//   storageTypes - Comma separated list of StorageType to clear.
func ClearDataForOrigin(origin string, storageTypes string) *ClearDataForOriginParams {
	return &ClearDataForOriginParams{
		Origin:       origin,
		StorageTypes: storageTypes,
	}
}

// Do executes Storage.clearDataForOrigin against the provided context.
func (p *ClearDataForOriginParams) Do(ctxt context.Context, h cdp.Executor) (err error) {
	return h.Execute(ctxt, CommandClearDataForOrigin, p, nil)
}

// GetUsageAndQuotaParams returns usage and quota in bytes.
type GetUsageAndQuotaParams struct {
	Origin string `json:"origin"` // Security origin.
}

// GetUsageAndQuota returns usage and quota in bytes.
//
// parameters:
//   origin - Security origin.
func GetUsageAndQuota(origin string) *GetUsageAndQuotaParams {
	return &GetUsageAndQuotaParams{
		Origin: origin,
	}
}

// GetUsageAndQuotaReturns return values.
type GetUsageAndQuotaReturns struct {
	Usage          float64         `json:"usage,omitempty"`          // Storage usage (bytes).
	Quota          float64         `json:"quota,omitempty"`          // Storage quota (bytes).
	UsageBreakdown []*UsageForType `json:"usageBreakdown,omitempty"` // Storage usage per type (bytes).
}

// Do executes Storage.getUsageAndQuota against the provided context.
//
// returns:
//   usage - Storage usage (bytes).
//   quota - Storage quota (bytes).
//   usageBreakdown - Storage usage per type (bytes).
func (p *GetUsageAndQuotaParams) Do(ctxt context.Context, h cdp.Executor) (usage float64, quota float64, usageBreakdown []*UsageForType, err error) {
	// execute
	var res GetUsageAndQuotaReturns
	err = h.Execute(ctxt, CommandGetUsageAndQuota, p, &res)
	if err != nil {
		return 0, 0, nil, err
	}

	return res.Usage, res.Quota, res.UsageBreakdown, nil
}

// TrackCacheStorageForOriginParams registers origin to be notified when an
// update occurs to its cache storage list.
type TrackCacheStorageForOriginParams struct {
	Origin string `json:"origin"` // Security origin.
}

// TrackCacheStorageForOrigin registers origin to be notified when an update
// occurs to its cache storage list.
//
// parameters:
//   origin - Security origin.
func TrackCacheStorageForOrigin(origin string) *TrackCacheStorageForOriginParams {
	return &TrackCacheStorageForOriginParams{
		Origin: origin,
	}
}

// Do executes Storage.trackCacheStorageForOrigin against the provided context.
func (p *TrackCacheStorageForOriginParams) Do(ctxt context.Context, h cdp.Executor) (err error) {
	return h.Execute(ctxt, CommandTrackCacheStorageForOrigin, p, nil)
}

// TrackIndexedDBForOriginParams registers origin to be notified when an
// update occurs to its IndexedDB.
type TrackIndexedDBForOriginParams struct {
	Origin string `json:"origin"` // Security origin.
}

// TrackIndexedDBForOrigin registers origin to be notified when an update
// occurs to its IndexedDB.
//
// parameters:
//   origin - Security origin.
func TrackIndexedDBForOrigin(origin string) *TrackIndexedDBForOriginParams {
	return &TrackIndexedDBForOriginParams{
		Origin: origin,
	}
}

// Do executes Storage.trackIndexedDBForOrigin against the provided context.
func (p *TrackIndexedDBForOriginParams) Do(ctxt context.Context, h cdp.Executor) (err error) {
	return h.Execute(ctxt, CommandTrackIndexedDBForOrigin, p, nil)
}

// UntrackCacheStorageForOriginParams unregisters origin from receiving
// notifications for cache storage.
type UntrackCacheStorageForOriginParams struct {
	Origin string `json:"origin"` // Security origin.
}

// UntrackCacheStorageForOrigin unregisters origin from receiving
// notifications for cache storage.
//
// parameters:
//   origin - Security origin.
func UntrackCacheStorageForOrigin(origin string) *UntrackCacheStorageForOriginParams {
	return &UntrackCacheStorageForOriginParams{
		Origin: origin,
	}
}

// Do executes Storage.untrackCacheStorageForOrigin against the provided context.
func (p *UntrackCacheStorageForOriginParams) Do(ctxt context.Context, h cdp.Executor) (err error) {
	return h.Execute(ctxt, CommandUntrackCacheStorageForOrigin, p, nil)
}

// UntrackIndexedDBForOriginParams unregisters origin from receiving
// notifications for IndexedDB.
type UntrackIndexedDBForOriginParams struct {
	Origin string `json:"origin"` // Security origin.
}

// UntrackIndexedDBForOrigin unregisters origin from receiving notifications
// for IndexedDB.
//
// parameters:
//   origin - Security origin.
func UntrackIndexedDBForOrigin(origin string) *UntrackIndexedDBForOriginParams {
	return &UntrackIndexedDBForOriginParams{
		Origin: origin,
	}
}

// Do executes Storage.untrackIndexedDBForOrigin against the provided context.
func (p *UntrackIndexedDBForOriginParams) Do(ctxt context.Context, h cdp.Executor) (err error) {
	return h.Execute(ctxt, CommandUntrackIndexedDBForOrigin, p, nil)
}

// Command names.
const (
	CommandClearDataForOrigin           = "Storage.clearDataForOrigin"
	CommandGetUsageAndQuota             = "Storage.getUsageAndQuota"
	CommandTrackCacheStorageForOrigin   = "Storage.trackCacheStorageForOrigin"
	CommandTrackIndexedDBForOrigin      = "Storage.trackIndexedDBForOrigin"
	CommandUntrackCacheStorageForOrigin = "Storage.untrackCacheStorageForOrigin"
	CommandUntrackIndexedDBForOrigin    = "Storage.untrackIndexedDBForOrigin"
)