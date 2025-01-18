package api

//Coin balance params
type CoinBalanceParams struct {
	Username string
}

//Coin balance Response
type CoinBalanceResponse struct {
	//Success code 200
	Code int

	//Account Balance
	Balance int64
}

//Error Response
type Error struct {
	//Error code
	Code int
	//Error message
	Message string
}
