#!/bin/bash
#
# Copyright IBM Corp All Rights Reserved
#
# SPDX-License-Identifier: Apache-2.0
#
# Exit on first error
set -ex

# Bring the test network down
pushd ../network
./network.sh down
popd

# clean out any old identites in the wallets
rm -rf server/wallet/*
rm -rf java/wallet/*
rm -rf typescript/wallet/*
