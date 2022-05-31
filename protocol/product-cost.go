package protocol

type ProductCost struct {
	ProductID() [16]byte           // product-status domain
	OrgID            [16]byte // Sell can be register just by producer organization

	MaterialsPercent   math.PerMyriad
	MaterialsCost      price.Amount
	LaborPercent       math.PerMyriad
	LaborCost          price.Amount
	InvestmentsPercent math.PerMyriad
	InvestmentsCost    price.Amount
	TotalCost          price.Amount

	Markup          math.PerMyriad // the % added to the cost to determine the wholesale price
	WholesaleProfit price.Amount   // The amount of wholesale markup
	Margin          math.PerMyriad // the % of product price to determine the retail price
	RetailProfit    price.Amount   // The amount of retail margin

	Cost price.Amount
}
