package fastpurgev3

import (
	"fmt"
)

// PurgeCacheByURL Invalidates content on the selected URL for the selected network.
// Akamai API docs: https://developer.akamai.com/api/core_features/fast_purge/v3.html
func (fp *Fastpurgev3) PurgeCacheByURL(opts FastPurgeRequest, tier AkamaiEnvironment, purgeStrategy AkamaiPurgeStrategy) (*FastPurgeResult, error) {

	resp, err := fp.executePurgeRequest(opts, purgeStrategy, tier, URL)

	return resp, err

}

// PurgeCacheByCPCode Invalidates content on the selected CPCODE for the selected network.
// Akamai API docs: https://developer.akamai.com/api/core_features/fast_purge/v3.html
func (fp *Fastpurgev3) PurgeCacheByCPCode(opts FastPurgeRequest, tier AkamaiEnvironment, purgeStrategy AkamaiPurgeStrategy) (*FastPurgeResult, error) {

	resp, err := fp.executePurgeRequest(opts, purgeStrategy, tier, URL)

	return resp, err

}

// PurgeCacheByCacheTag Invalidates content on the selected CPCODE for the selected network.
// Akamai API docs: https://developer.akamai.com/api/core_features/fast_purge/v3.html
func (fp *Fastpurgev3) PurgeCacheByCacheTag(opts FastPurgeRequest, tier AkamaiEnvironment, purgeStrategy AkamaiPurgeStrategy) (*FastPurgeResult, error) {

	resp, err := fp.executePurgeRequest(opts, purgeStrategy, tier, URL)

	return resp, err

}

//executePurgeRequest executes request to invalidate cache based on given conditions.
//AkamaiPurgeStrategy: delete | invalidate
//AkamaiEnvironment: production | staging
//AkamaiPurgeType: URL|cpcode|cache tag
func (fp *Fastpurgev3) executePurgeRequest(opts FastPurgeRequest, purgeStrategy AkamaiPurgeStrategy, tier AkamaiEnvironment, purgeType AkamaiPurgeType) (*FastPurgeResult, error) {

	// Create and execute request
	resp, err := fp.Client.Rclient.R().
		SetBody(opts).
		SetResult(FastPurgeResult{}).
		SetError(FastpurgeError{}).
		Post(fmt.Sprintf("%s/%s/%s/%s", basePath, purgeStrategy, purgeType, tier))

	if err != nil {
		return nil, err
	}

	if resp.IsError() {
		e := resp.Error().(*FastpurgeError)
		return nil, e
	}

	return resp.Result().(*FastPurgeResult), nil

}
