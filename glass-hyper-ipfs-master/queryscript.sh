#!/bin/bash

# set -v


# cat <<EOF

# ========================================================================================================
# Query Resource 1 IPFS Resource (as org1.org and org2.org - should result in permitted and permitted respectively):
# ========================================================================================================

# EOF

#!!!WARNING!! These commands are no longer functional. Retained for reference only. Use glass-portal.py to interact with Hyperledger instead.

minifab="$(pwd)/minifab"

$minifab query -p '"readIPFSResource","resource1"' -t '' -o org1.org
$minifab query -p '"readIPFSResource","resource1"' -t '' -o org2.org

# cat <<EOF

# ========================================================================================================
# Query Resource 1 IPFS Resource Key (as org1.org and org2.org - should result in permitted and denied respectively):
# ========================================================================================================

$minifab query -p '"readIPFSResourceKey","resource1"' -t '' -o org1.org
$minifab query -p '"readIPFSResourceKey","resource1"' -t '' -o org2.org

# EOF


