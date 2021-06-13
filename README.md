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
Hash: 00000a51a11c4e1d5ee2383ae5b02a77216775ed1dab4c92f102803b64bae33f, Previous hash: 000000000000000000000000000000000
0000000000000000000000000000000, Timestamp: 2021-06-13 12:40:15.2232105 +0900 KST m=+4.842418201, Nonce: 935519, Data: F
irst block

Mining second block...
Block mined.
Hash: 0000085fe3f049157462645dfd04384606bbd3844240afef11e12f1b489e0ab3, Previous hash: 00000a51a11c4e1d5ee2383ae5b02a772
16775ed1dab4c92f102803b64bae33f, Timestamp: 2021-06-13 12:40:24.9754905 +0900 KST m=+14.594698201, Nonce: 1929334, Data:
 Second block

Mining third block...
Block mined.
Hash: 00000573ec89408ab59724a0eaa2c0df86e0d67b8b38bbdbf431ace8751b65d5, Previous hash: 0000085fe3f049157462645dfd0438460
6bbd3844240afef11e12f1b489e0ab3, Timestamp: 2021-06-13 12:40:25.2692706 +0900 KST m=+14.888478301, Nonce: 57782, Data: T
hird block

Is blockchain valid? true

Printing blockchain...
Blockchain - difficulty: 5, blocks: 3
Hash: 00000a51a11c4e1d5ee2383ae5b02a77216775ed1dab4c92f102803b64bae33f, Previous hash: 000000000000000000000000000000000
0000000000000000000000000000000, Timestamp: 2021-06-13 12:40:15.2232105 +0900 KST m=+4.842418201, Nonce: 935519, Data: F
irst block
Hash: 0000085fe3f049157462645dfd04384606bbd3844240afef11e12f1b489e0ab3, Previous hash: 00000a51a11c4e1d5ee2383ae5b02a772
16775ed1dab4c92f102803b64bae33f, Timestamp: 2021-06-13 12:40:24.9754905 +0900 KST m=+14.594698201, Nonce: 1929334, Data:
 Second block
Hash: 00000573ec89408ab59724a0eaa2c0df86e0d67b8b38bbdbf431ace8751b65d5, Previous hash: 0000085fe3f049157462645dfd0438460
6bbd3844240afef11e12f1b489e0ab3, Timestamp: 2021-06-13 12:40:25.2692706 +0900 KST m=+14.888478301, Nonce: 57782, Data: T
hird block
```