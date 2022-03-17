package protocol

import "time"

type Customer struct {
	ID uint64 `json:"id"`
	// Customer's name. Alphanumeric, with period (.), apostrophe (') and parentheses allowed.
	// Name must be between 3-50 characters in length.
	Name string `json:"name"`
	// Maximum length of 15 characters, inclusive of country code. For example, +919876543210.
	Contact string `json:"contact"`
	// Maximum length of 64 characters.
	Email string `json:"email"`
	// GST number linked to the customer. For example, 29XAbbA4369J1PA.
	// GSTIN, short for Goods and Services Tax Identification Number is a unique 15 digit identification number assigned to every taxpayer ( primarily dealer or supplier or any business entity) registered under the GST regime.
	// Obtaining GSTIN and registering for GST is absolutely free of cos
	Gstin string `json:"gstin"`
	// This is a key-value pair that can be used to store additional information about the entity.
	// It can hold a maximum of 15 key-value pairs, 256 characters (maximum) each. For example, "note_key": "Beam me up Scotty”.
	Notes string `json:"notes"`
	Created_At time.Time `json:"created_at"`
}

// If the GSTIN is correct, the details that can be verified here are-

//     The legal name of the business
//     State
//     Date of registration
//     Constitution of business – company, sole-proprietor or partnership
//     Taxpayer type – regular taxpayer or composition dealer
//     GSTIN status / UIN status

// ============================
// fail_existing
// optional
//     string Possible values:

//         0: If a customer with the same details already exists, fetches details of the existing customer.
//         1 (default): If a customer with the same details already exists, throws an error.
