# CS1699-Deliverable-4-Impact-of-Blockchain

# Hyperledger Fabric Setup

CouchDB - http://localhost:5984/_utils/#/database/mychannel_provider_sync/_all_docs

Dependencies:
- node + npm
    - may need to do: NODE_TLS_REJECT_UNAUTHORIZED=0
- docker + docker_compose
- Python 2.7
- Go
    - Fabric source code should be available locally in GOPATH
			cd $GOPATH/src/github.com
			mkdir hyperledger
			cd hyperledger
			git clone http://gerrit.hyperledger.org/r/fabric
    - Setting GOPATH
        - https://github.com/golang/go/wiki/SettingGOPATH#bash

# Chaincode
https://github.com/Mavirek/CS1699-Deliverable-4-Impact-of-Blockchain/blob/master/chaincode/provider_sync/provider_sync.go

# To run
  * Go to provider_sync folder
  * Run `./startFabric.sh`
    * Make sure Docker is running!
  * Run `npm install`, `node enrollAdmin.js`, and `node registerUser.js` in order to start interacting with the ledger
  * Run `node invoke.js` or `node query.js` if you want to add/update or query a provider respectively
    * Edit the request parameters in `invoke.js` and `query.js` to try out additional methods in the chaincode
  * Visit the CouchDB localhost link (above) to view a table representation of the blockchain data
  
# Shutdown network
  * Run `./stopFabric.sh`
