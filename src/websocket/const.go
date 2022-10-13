package websocket

var (
	PANCAKE_FACTORY_ADDRESS string = "0xca143ce32fe78f1f7019d7d551a6402fc5350c73"
	WS_URL                  string = "wss://bsc-ws-node.nariox.org:443"
	// PairCreated(address,address,address,uint256)
	CREATE_PAIR_TOPIC string = "0x0d3648bd0f6ba80134a33ba9275ac585d9d315f0ad8355cddefde31afa28d0e9"
	// Mint(address,uint256,uint256)
	PAIR_MINT_TOPIC string = "0x4c209b5fc8ad50758f13e2e1088ba56a560dff690a1c6fef26394f4c03821c4f"
	BNB_ADDRESS     string = ""
	BUSD_ADDRESS    string = "0xe9e7cea3dedca5984780bafc599bd69add087d56"
	// 5 BNB
	BNB_MINIMUM_REQUIRED_AMOUNT string = "0x000000000000000000000000000000000000000000004563918244f4000000"
	// 2000 BUSD
	BUSD_MINIMUM_REQUIRED_AMOUNT string = "0x000000000000000000000000000000000000000000006c6b935b8bbd400000"
	DEBUG                        bool   = true
)
