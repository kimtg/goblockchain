# goblockchain
Simple blockchain in Go

This repository implements a very simple Blockchain, based off of this
post and its code: [Simple blockchain in Ada](https://tomekw.com/simple-blockchain-in-ada/)

See the demo file for basic usage.

## Example Run
```
Simple blockchain demo

Mining first block...
Block mined.
Hash: 000005a7b6dad9a2a8992b168c2e0b8a427f91d35cee0d61963d0082eb76f8d8, Previous hash: 0000000000000000000000000000000000000000000000000000000000000000, Timestamp: 2021-06-13 12:46:51.7450617 +0900 KST m=+3.810524701, Nonce: 1095267, Data: First block

Mining second block...
Block mined.
Hash: 00000827f7318faaae8454cfb9af9ec3d060e70d7c3479878800d330dd644ec3, Previous hash: 000005a7b6dad9a2a8992b168c2e0b8a427f91d35cee0d61963d0082eb76f8d8, Timestamp: 2021-06-13 12:46:57.239535 +0900 KST m=+9.304998001, Nonce: 1626359, Data: Second block

Mining third block...
Block mined.
Hash: 0000044380e281b06b519126e846b37da033856cf91835d445aec63bd84cfeba, Previous hash: 00000827f7318faaae8454cfb9af9ec3d060e70d7c3479878800d330dd644ec3, Timestamp: 2021-06-13 12:46:58.7691071 +0900 KST m=+10.834570101, Nonce: 454301, Data: Third block

Is blockchain valid? true

Printing blockchain...
Blockchain - difficulty: 5, blocks: 3
Hash: 000005a7b6dad9a2a8992b168c2e0b8a427f91d35cee0d61963d0082eb76f8d8, Previous hash: 0000000000000000000000000000000000000000000000000000000000000000, Timestamp: 2021-06-13 12:46:51.7450617 +0900 KST m=+3.810524701, Nonce: 1095267, Data: First block
Hash: 00000827f7318faaae8454cfb9af9ec3d060e70d7c3479878800d330dd644ec3, Previous hash: 000005a7b6dad9a2a8992b168c2e0b8a427f91d35cee0d61963d0082eb76f8d8, Timestamp: 2021-06-13 12:46:57.239535 +0900 KST m=+9.304998001, Nonce: 1626359, Data: Second block
Hash: 0000044380e281b06b519126e846b37da033856cf91835d445aec63bd84cfeba, Previous hash: 00000827f7318faaae8454cfb9af9ec3d060e70d7c3479878800d330dd644ec3, Timestamp: 2021-06-13 12:46:58.7691071 +0900 KST m=+10.834570101, Nonce: 454301, Data: Third block
```