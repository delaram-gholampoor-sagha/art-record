package protocol

// OrgProfit
type OrgBudgetActualUTC interface {
	Week() utc.WeekElapsed           //
	Income() protocol.AmountOfMoney  //
	Outcome() protocol.AmountOfMoney //
	Time() protocol.Time             // save time
	RequestID() [16]byte             // user-request domain
}

// Usage:
// Country Governments: need to steal organizations profits by the name of law as Tax, ...! so force them to implement these services.
// Stock Markets: need to show related data to traders to know the quality of organizations in making profit.
const (
	ProfitDomain = "profit.protocol"

	ProfitServiceBaseCurrency     = "urn:giti:profit.protocol:service:base-currency"
	ProfitServiceAllCurrency      = "urn:giti:profit.protocol:service:all-currency"
	ProfitServiceSpecificCurrency = "urn:giti:profit.protocol:service:specific-currency" // for each day
	// ProfitService              = "urn:giti:profit.protocol:service:"
)

type ProfitServiceBaseCurrencyResponse interface {
	BaseCurrency() uint32 // same as SocietyID in GP protocol
}

type ProfitServiceAllCurrencyRequest interface {
	Time() protocol.Time // it will round to a day and return profit for the day.
}
type ProfitServiceAllCurrencyResponse interface {
	Amount() int64 // in the base currency
}
