package protocol


// ProductCost indicate the domain record data fields.
type ProductCost interface {
	ProductID() [16]byte // product domain
	OrgID() [16]byte     // Sell can be register just by producer organization

	Materials() math.PerMyriad
	MaterialsCost() protocol.AmountOfMoney
	Labor() math.PerMyriad
	LaborCost() protocol.AmountOfMoney
	Investments() math.PerMyriad
	InvestmentsCost() protocol.AmountOfMoney
	TotalCost() protocol.AmountOfMoney

	Markup() math.PerMyriad                  // the % added to the cost to determine the wholesale price
	WholesaleProfit() protocol.AmountOfMoney // The amount of wholesale markup
	Margin() math.PerMyriad                  // the % of product price to determine the retail price
	RetailProfit() protocol.AmountOfMoney    // The amount of retail margin

	Cost() protocol.AmountOfMoney
}
