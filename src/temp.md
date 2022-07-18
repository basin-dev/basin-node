# Methods for `list`

// returns list of all Basin URLs
// mode == producer OR consumer
getSources(mode)

// returns metadata for a specific Basin URL
// url == a local basin url
getMetadata(url)

// returns network reputation score (and metadata)
// mode == producer OR consumer
getReputation(mode)

// returns information about the wallet connected to node
getWalletInfo()

**Maybe those four are enough, but it might also be nice to have...**

// returns list of all outstanding requests from / for that node (including metadata)
// if URL is specified, then it returns request for that specific resource
// mode == producer OR consumer
getRequests(mode, url=)

// returns list of all ongoing royalties from / for that node (including metadata)
// if URL is specified, then it returns request for that specific resource
// mode == producer OR consumer
getRoyalties(mode, url=)

// returns list of all current schemas from / for that node (including metadata)
// if URL is specified, then it returns request for that specific resource
// mode == producer OR consumer
getSchemas(mode, url=)

// returns list of all current cache expectations from / for that node (including metadata)
// if URL is specified, then it returns request for that specific resource
// mode == producer OR consumer
getCacheExepectations(mode, url=)