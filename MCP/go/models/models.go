package models

import (
	"context"
	"github.com/mark3labs/mcp-go/mcp"
)

type Tool struct {
	Definition mcp.Tool
	Handler    func(ctx context.Context, req mcp.CallToolRequest) (*mcp.CallToolResult, error)
}

// SocialLink represents the SocialLink schema from the OpenAPI specification
type SocialLink struct {
	Id string `json:"id,omitempty"` // Unique identifier of the social link
	TypeField string `json:"type,omitempty"` // Type of the social link, e.g. twitter
	Url string `json:"url"` // URL of the social link, e.g. https://www.twitter.com/apideck
}

// Address represents the Address schema from the OpenAPI specification
type Address struct {
	Name string `json:"name,omitempty"` // The name of the address.
	Line2 string `json:"line2,omitempty"` // Line 2 of the address
	Phone_number string `json:"phone_number,omitempty"` // Phone number of the address
	Latitude string `json:"latitude,omitempty"` // Latitude of the address
	State string `json:"state,omitempty"` // Name of state
	Fax string `json:"fax,omitempty"` // Fax number of the address
	Postal_code string `json:"postal_code,omitempty"` // Zip code or equivalent.
	Contact_name string `json:"contact_name,omitempty"` // Name of the contact person at the address
	Line3 string `json:"line3,omitempty"` // Line 3 of the address
	Notes string `json:"notes,omitempty"` // Additional notes
	Street_number string `json:"street_number,omitempty"` // Street number
	Row_version string `json:"row_version,omitempty"` // A binary value used to detect updates to a object and prevent data conflicts. It is incremented each time an update is made to the object.
	StringField string `json:"string,omitempty"` // The address string. Some APIs don't provide structured address data.
	Id string `json:"id,omitempty"` // Unique identifier for the address.
	Country string `json:"country,omitempty"` // country code according to ISO 3166-1 alpha-2.
	County string `json:"county,omitempty"` // Address field that holds a sublocality, such as a county
	TypeField string `json:"type,omitempty"` // The type of address.
	Website string `json:"website,omitempty"` // Website of the address
	Line4 string `json:"line4,omitempty"` // Line 4 of the address
	City string `json:"city,omitempty"` // Name of city.
	Email string `json:"email,omitempty"` // Email address of the address
	Line1 string `json:"line1,omitempty"` // Line 1 of the address e.g. number, street, suite, apt #, etc.
	Salutation string `json:"salutation,omitempty"` // Salutation of the contact person at the address
	Longitude string `json:"longitude,omitempty"` // Longitude of the address
}

// PhoneNumber represents the PhoneNumber schema from the OpenAPI specification
type PhoneNumber struct {
	Id string `json:"id,omitempty"` // Unique identifier of the phone number
	Number string `json:"number"` // The phone number
	TypeField string `json:"type,omitempty"` // The type of phone number
	Area_code string `json:"area_code,omitempty"` // The area code of the phone number, e.g. 323
	Country_code string `json:"country_code,omitempty"` // The country code of the phone number, e.g. +1
	Extension string `json:"extension,omitempty"` // The extension of the phone number
}

// LeadsSort represents the LeadsSort schema from the OpenAPI specification
type LeadsSort struct {
	By string `json:"by,omitempty"` // The field on which to sort the Leads
	Direction string `json:"direction,omitempty"` // The direction in which to sort the results
}

// Links represents the Links schema from the OpenAPI specification
type Links struct {
	Current string `json:"current,omitempty"` // Link to navigate to the current page through the API
	Next string `json:"next,omitempty"` // Link to navigate to the previous page through the API
	Previous string `json:"previous,omitempty"` // Link to navigate to the previous page through the API
}

// TooManyRequestsResponse represents the TooManyRequestsResponse schema from the OpenAPI specification
type TooManyRequestsResponse struct {
	Type_name string `json:"type_name,omitempty"` // The type of error returned
	Detail map[string]interface{} `json:"detail,omitempty"`
	ErrorField string `json:"error,omitempty"` // Contains an explanation of the status_code as defined in HTTP/1.1 standard (RFC 6585)
	Message string `json:"message,omitempty"` // A human-readable message providing more details about the error.
	Ref string `json:"ref,omitempty"` // Link to documentation of error type
	Status_code float64 `json:"status_code,omitempty"` // HTTP status code
}

// BankAccount represents the BankAccount schema from the OpenAPI specification
type BankAccount struct {
	Bank_name string `json:"bank_name,omitempty"` // The name of the bank
	Account_name string `json:"account_name,omitempty"` // The name which you used in opening your bank account.
	Account_number string `json:"account_number,omitempty"` // A bank account number is a number that is tied to your bank account. If you have several bank accounts, such as personal, joint, business (and so on), each account will have a different account number.
	Account_type string `json:"account_type,omitempty"` // The type of bank account.
	Branch_identifier string `json:"branch_identifier,omitempty"` // A branch identifier is a unique identifier for a branch of a bank or financial institution.
	Routing_number string `json:"routing_number,omitempty"` // A routing number is a nine-digit code used to identify a financial institution in the United States.
	Bic string `json:"bic,omitempty"` // The Bank Identifier Code (BIC).
	Bsb_number string `json:"bsb_number,omitempty"` // A BSB is a 6 digit numeric code used for identifying the branch of an Australian or New Zealand bank or financial institution.
	Currency string `json:"currency,omitempty"` // Indicates the associated currency for an amount of money. Values correspond to [ISO 4217](https://en.wikipedia.org/wiki/ISO_4217).
	Iban string `json:"iban,omitempty"` // The International Bank Account Number (IBAN).
	Bank_code string `json:"bank_code,omitempty"` // A bank code is a code assigned by a central bank, a bank supervisory body or a Bankers Association in a country to all its licensed member banks or financial institutions.
}

// Email represents the Email schema from the OpenAPI specification
type Email struct {
	Email string `json:"email"` // Email address
	Id string `json:"id,omitempty"` // Unique identifier for the email address
	TypeField string `json:"type,omitempty"` // Email type
}

// LeadWebhookEvent represents the LeadWebhookEvent schema from the OpenAPI specification
type LeadWebhookEvent struct {
	Entity_id string `json:"entity_id,omitempty"` // The service provider's ID of the entity that triggered this event
	Entity_url string `json:"entity_url,omitempty"` // The url to retrieve entity detail.
	Occurred_at string `json:"occurred_at,omitempty"` // ISO Datetime for when the original event occurred
	Entity_type string `json:"entity_type,omitempty"` // The type entity that triggered this event
	Event_id string `json:"event_id,omitempty"` // Unique reference to this request event
	Execution_attempt float64 `json:"execution_attempt,omitempty"` // The current count this request event has been attempted
	Service_id string `json:"service_id,omitempty"` // Service provider identifier
	Unified_api string `json:"unified_api,omitempty"` // Name of Apideck Unified API
	Consumer_id string `json:"consumer_id,omitempty"` // Unique consumer identifier. You can freely choose a consumer ID yourself. Most of the time, this is an ID of your internal data model that represents a user or account in your system (for example account:12345). If the consumer doesn't exist yet, Vault will upsert a consumer based on your ID.
	Event_type string `json:"event_type,omitempty"`
}

// DeleteLeadResponse represents the DeleteLeadResponse schema from the OpenAPI specification
type DeleteLeadResponse struct {
	Operation string `json:"operation"` // Operation performed
	Resource string `json:"resource"` // Unified API resource name
	Service string `json:"service"` // Apideck ID of service provider
	Status string `json:"status"` // HTTP Response Status
	Status_code int `json:"status_code"` // HTTP Response Status Code
	Data UnifiedId `json:"data"`
}

// UpdateLeadResponse represents the UpdateLeadResponse schema from the OpenAPI specification
type UpdateLeadResponse struct {
	Service string `json:"service"` // Apideck ID of service provider
	Status string `json:"status"` // HTTP Response Status
	Status_code int `json:"status_code"` // HTTP Response Status Code
	Data UnifiedId `json:"data"`
	Operation string `json:"operation"` // Operation performed
	Resource string `json:"resource"` // Unified API resource name
}

// UnauthorizedResponse represents the UnauthorizedResponse schema from the OpenAPI specification
type UnauthorizedResponse struct {
	Detail string `json:"detail,omitempty"` // Contains parameter or domain specific information related to the error and why it occurred.
	ErrorField string `json:"error,omitempty"` // Contains an explanation of the status_code as defined in HTTP/1.1 standard (RFC 7231)
	Message string `json:"message,omitempty"` // A human-readable message providing more details about the error.
	Ref string `json:"ref,omitempty"` // Link to documentation of error type
	Status_code float64 `json:"status_code,omitempty"` // HTTP status code
	Type_name string `json:"type_name,omitempty"` // The type of error returned
}

// GetLeadsResponse represents the GetLeadsResponse schema from the OpenAPI specification
type GetLeadsResponse struct {
	Meta Meta `json:"meta,omitempty"` // Response metadata
	Operation string `json:"operation"` // Operation performed
	Resource string `json:"resource"` // Unified API resource name
	Service string `json:"service"` // Apideck ID of service provider
	Status string `json:"status"` // HTTP Response Status
	Status_code int `json:"status_code"` // HTTP Response Status Code
	Data []Lead `json:"data"`
	Links Links `json:"links,omitempty"` // Links to navigate to previous or next pages through the API
}

// UnexpectedErrorResponse represents the UnexpectedErrorResponse schema from the OpenAPI specification
type UnexpectedErrorResponse struct {
	ErrorField string `json:"error,omitempty"` // Contains an explanation of the status_code as defined in HTTP/1.1 standard (RFC 7231)
	Message string `json:"message,omitempty"` // A human-readable message providing more details about the error.
	Ref string `json:"ref,omitempty"` // Link to documentation of error type
	Status_code float64 `json:"status_code,omitempty"` // HTTP status code
	Type_name string `json:"type_name,omitempty"` // The type of error returned
	Detail interface{} `json:"detail,omitempty"` // Contains parameter or domain specific information related to the error and why it occurred.
}

// Website represents the Website schema from the OpenAPI specification
type Website struct {
	Id string `json:"id,omitempty"` // Unique identifier for the website
	TypeField string `json:"type,omitempty"` // The type of website
	Url string `json:"url"` // The website URL
}

// NotImplementedResponse represents the NotImplementedResponse schema from the OpenAPI specification
type NotImplementedResponse struct {
	Status_code float64 `json:"status_code,omitempty"` // HTTP status code
	Type_name string `json:"type_name,omitempty"` // The type of error returned
	Detail interface{} `json:"detail,omitempty"` // Contains parameter or domain specific information related to the error and why it occurred.
	ErrorField string `json:"error,omitempty"` // Contains an explanation of the status_code as defined in HTTP/1.1 standard (RFC 7231)
	Message string `json:"message,omitempty"` // A human-readable message providing more details about the error.
	Ref string `json:"ref,omitempty"` // Link to documentation of error type
}

// Meta represents the Meta schema from the OpenAPI specification
type Meta struct {
	Cursors map[string]interface{} `json:"cursors,omitempty"` // Cursors to navigate to previous or next pages through the API
	Items_on_page int `json:"items_on_page,omitempty"` // Number of items returned in the data property of the response
}

// BadRequestResponse represents the BadRequestResponse schema from the OpenAPI specification
type BadRequestResponse struct {
	Detail interface{} `json:"detail,omitempty"` // Contains parameter or domain specific information related to the error and why it occurred.
	ErrorField string `json:"error,omitempty"` // Contains an explanation of the status_code as defined in HTTP/1.1 standard (RFC 7231)
	Message string `json:"message,omitempty"` // A human-readable message providing more details about the error.
	Ref string `json:"ref,omitempty"` // Link to documentation of error type
	Status_code float64 `json:"status_code,omitempty"` // HTTP status code
	Type_name string `json:"type_name,omitempty"` // The type of error returned
}

// CustomMappings represents the CustomMappings schema from the OpenAPI specification
type CustomMappings struct {
}

// LeadsFilter represents the LeadsFilter schema from the OpenAPI specification
type LeadsFilter struct {
	Name string `json:"name,omitempty"` // Name of the lead to filter on
	Email string `json:"email,omitempty"` // E-mail of the lead to filter on
	First_name string `json:"first_name,omitempty"` // First name of the lead to filter on
	Last_name string `json:"last_name,omitempty"` // Last name of the lead to filter on
}

// CustomField represents the CustomField schema from the OpenAPI specification
type CustomField struct {
	Id string `json:"id"` // Unique identifier for the custom field.
	Name string `json:"name,omitempty"` // Name of the custom field.
	Value interface{} `json:"value,omitempty"`
	Description string `json:"description,omitempty"` // More information about the custom field
}

// UnifiedId represents the UnifiedId schema from the OpenAPI specification
type UnifiedId struct {
	Id string `json:"id"` // The unique identifier of the resource
}

// PaymentRequiredResponse represents the PaymentRequiredResponse schema from the OpenAPI specification
type PaymentRequiredResponse struct {
	Type_name string `json:"type_name,omitempty"` // The type of error returned
	Detail string `json:"detail,omitempty"` // Contains parameter or domain specific information related to the error and why it occurred.
	ErrorField string `json:"error,omitempty"` // Contains an explanation of the status_code as defined in HTTP/1.1 standard (RFC 7231)
	Message string `json:"message,omitempty"` // A human-readable message providing more details about the error.
	Ref string `json:"ref,omitempty"` // Link to documentation of error type
	Status_code float64 `json:"status_code,omitempty"` // HTTP status code
}

// UnprocessableResponse represents the UnprocessableResponse schema from the OpenAPI specification
type UnprocessableResponse struct {
	Type_name string `json:"type_name,omitempty"` // The type of error returned
	Detail string `json:"detail,omitempty"` // Contains parameter or domain specific information related to the error and why it occurred.
	ErrorField string `json:"error,omitempty"` // Contains an explanation of the status_code as defined in HTTP/1.1 standard (RFC 7231)
	Message string `json:"message,omitempty"` // A human-readable message providing more details about the error.
	Ref string `json:"ref,omitempty"` // Link to documentation of error type
	Status_code float64 `json:"status_code,omitempty"` // HTTP status code
}

// NotFoundResponse represents the NotFoundResponse schema from the OpenAPI specification
type NotFoundResponse struct {
	ErrorField string `json:"error,omitempty"` // Contains an explanation of the status_code as defined in HTTP/1.1 standard (RFC 7231)
	Message string `json:"message,omitempty"` // A human-readable message providing more details about the error.
	Ref string `json:"ref,omitempty"` // Link to documentation of error type
	Status_code float64 `json:"status_code,omitempty"` // HTTP status code
	Type_name string `json:"type_name,omitempty"` // The type of error returned
	Detail interface{} `json:"detail,omitempty"` // Contains parameter or domain specific information related to the error and why it occurred.
}

// Lead represents the Lead schema from the OpenAPI specification
type Lead struct {
	Lead_id string `json:"lead_id,omitempty"` // The identifier of the lead.
	Title string `json:"title,omitempty"` // The job title of the lead.
	First_name string `json:"first_name,omitempty"` // The first name of the lead.
	Tags []string `json:"tags,omitempty"`
	Created_at string `json:"created_at,omitempty"` // Date created in ISO 8601 format
	Status string `json:"status,omitempty"`
	Custom_fields []CustomField `json:"custom_fields,omitempty"`
	Description string `json:"description,omitempty"` // The description of the lead.
	Prefix string `json:"prefix,omitempty"` // The prefix of the lead.
	Addresses []Address `json:"addresses,omitempty"`
	Emails []Email `json:"emails,omitempty"`
	Currency string `json:"currency,omitempty"` // Indicates the associated currency for an amount of money. Values correspond to [ISO 4217](https://en.wikipedia.org/wiki/ISO_4217).
	Company_id string `json:"company_id,omitempty"` // The company the lead is associated with.
	Monetary_amount float64 `json:"monetary_amount,omitempty"` // The monetary amount of the lead.
	Websites []Website `json:"websites,omitempty"`
	Owner_id string `json:"owner_id,omitempty"` // The owner of the lead.
	Phone_numbers []PhoneNumber `json:"phone_numbers,omitempty"`
	Custom_mappings CustomMappings `json:"custom_mappings,omitempty"` // When custom mappings are configured on the resource, the result is included here.
	Lead_source string `json:"lead_source,omitempty"` // The source of the lead.
	Social_links []SocialLink `json:"social_links,omitempty"`
	Name string `json:"name"` // Full name of the lead.
	Last_name string `json:"last_name,omitempty"` // The last name of the lead.
	Updated_at string `json:"updated_at,omitempty"` // Date updated in ISO 8601 format
	Company_name string `json:"company_name"` // The name of the company the lead is associated with.
	Fax string `json:"fax,omitempty"` // The fax number of the lead.
	Language string `json:"language,omitempty"` // language code according to ISO 639-1. For the United States - EN
	Id string `json:"id,omitempty"` // Unique identifier for the contact.
}

// CreateLeadResponse represents the CreateLeadResponse schema from the OpenAPI specification
type CreateLeadResponse struct {
	Status_code int `json:"status_code"` // HTTP Response Status Code
	Data UnifiedId `json:"data"`
	Operation string `json:"operation"` // Operation performed
	Resource string `json:"resource"` // Unified API resource name
	Service string `json:"service"` // Apideck ID of service provider
	Status string `json:"status"` // HTTP Response Status
}

// GetLeadResponse represents the GetLeadResponse schema from the OpenAPI specification
type GetLeadResponse struct {
	Data Lead `json:"data"`
	Operation string `json:"operation"` // Operation performed
	Resource string `json:"resource"` // Unified API resource name
	Service string `json:"service"` // Apideck ID of service provider
	Status string `json:"status"` // HTTP Response Status
	Status_code int `json:"status_code"` // HTTP Response Status Code
}
