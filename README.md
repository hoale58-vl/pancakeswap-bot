# Pancake Add Liquidity contract

```
0x10ED43C718714eb63d5aA57B78B54704E256024E
```

# Event

```
addLiquidity(
    address tokenA,
    address tokenB,
    uint amountADesired,
    uint amountBDesired,
    uint amountAMin,
    uint amountBMin,
    address to,
    uint deadline
)


----Websocket------
PancakeFactory
=> createPair(address tokenA, address tokenB)
=> emit PairCreated(token0, token1, pair, allPairs.length);

token0 = address
token1 = address

----RPC------
pair =  address (PancakePair)
=> event Mint(address indexed sender, uint256 amount0, uint256 amount1);
=> emit Mint(msg.sender, amount0, amount1);

```

# How to run

```
cd src
cp empty.db database.db
go get
go run main.go
```
