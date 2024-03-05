# Go API client for cysend

# 1 Table of content
1. Table of content
2. Overview
3. Changelog
4. Definitions
5. API Structure
6. Technical requirements
7. Read before you start coding
8. Integration is done, how do we go live?
9. Security
10. HTTP Status Codes
11. About us
# 2 Overview
The cysend.com API is a gateway to buy all the products we carry.<br />
The cysend.com API is built following the OpenAPI standards.<br />
[read more](https://www.openapis.org/)
## 2.1 Audience
This document contains the technical interconnection specifications. It is designed to be read by the CUSTOMER’s technical team in charge of the integration.
## 2.2 Uses
The cysend.com API can be interfaced with different kinds of devices &
hosts, such as:
1. Mobile phones
2. Terminals
3. CUSTOMER web sites
4. Automated points of sales
5. Automated teller machines
6. White label websites
7. White label applications
8. Custom applications & hardware
# 3 Changelog
Date | Version | Changes | Author
----------- | ----------- | ----------- | -----------
2023-12-11 | 5.2.3 | <ul>Added parameters to receive PIN usage instructions in catalog and the possibility to filter the returned languages.<li>Added `return_usage_instruction` parameter in GET `/store/catalogue` and GET `/store/catalogue/products` operations.</li><li>Added `languages` parameter in GET `/store/catalogue` and GET `/store/catalogue/products` operations.</li><li>Added `usage_instructions` parameter to the `Product` schema. This parameter is optional and is only returned if new parameter `return_usage_instruction` is set to true.</li></ul> | CY.SEND OpenApi team
2022-11-05 | 5.2.2 | <ul><li>New operation `/store/order/cost`</li><br />This operation returns an indicative cost of a face value. It does not place a purchase order. It is only a calculation tool to know the cost of a face value before placing the final purchase order. It returns the current cost at the exact time of the operation.</ul> | CY.SEND OpenApi team
2022-09-08 | 5.2.1 | <ul><li>Added `country_zone` parameter to the `Product` schema<br />Returns an indication of the geographical zone in which the product can be used.</li></ul> | CY.SEND OpenApi team
2020-03-06 | 5.2.0 | <ul><li>Added mobile to catalogue operations `/store/catalogue/products` and `/store/catalogue/face-value`.<br />This filter will detect the operator and return the available product/face values.<br />The service has a cost and is billed on usage</li><li>Added notification operations. See Notification operation group and definitions (Chapter 4.13).</li><li>Added `notification` parameter in `/store/order` POST request.</li></ul> | CY.SEND OpenApi team
2019-10-08 | 5.1.3 | <ul><li>Added `product_description` parameter in catalogue.</li></ul> | CY.SEND OpenApi team
2019-05-24 | 5.1.2 | <ul><li>Added product types (`instant` and `prepaid_code`) to the `Product` schema</li><li>Added new definition to chapter 4: Product types</li><li>Renamed `PIN` schema to `PrepaidCode`</li><li>Added multi languages prepaid code usage instructions to the `PrepaidCode` schema</li><li>Renamed `Coverage` schema to `Catalogue`</li><li>Renamed `gift_card_pin_code` parameter to `gift_card_code` in `CY.SEND Gift Card` operations</li></ul> | CY.SEND OpenApi team
2019-05-14 | 5.1.1 | <ul><li>Added support for promotions in `Catalogue` operations</li><li>Renamed `Catalogue` schema parameter `maintenance` to `maintenances`</li><li>Added purchase `date` parameter in `OrderStatus` schema</li><li>Renamed `pin-code` parameter to `gift_card_pin_code` in `CY.SEND Gift Card` operations</li></ul> | CY.SEND OpenApi team
2019-03-19 | 5.1.0 | <ul><li>Public early access</li></ul> | CY.SEND OpenApi team

# 4 Definitions
## 4.1 API Credentials
API credentials are different from cysend.com credentials. Please ask the support team for your credentials.
## 4.2 HTTP Basic access authentication
In the context of an HTTP transaction, basic access authentication is a
method for an HTTP user agent to provide a user name and password when
making a request. For security reason, this authentication method should
only be used on HTTPS request.
[Wikipedia](https://en.wikipedia.org/wiki/Basic_access_authentication)
## 4.3 HTTP request methods
HTTP defines methods (sometimes referred to as \"verbs\") to indicate the
desired action to be performed on the identified resource.
## 4.4 HTTP status code
Status codes are issued by a server in response to a client's request made
to the server. It includes codes from IETF Request for Comments (RFCs),
other specifications, and some additional codes used in some common
applications of the Hypertext Transfer Protocol (HTTP). The first digit of
the status code specifies one of five standard classes of responses. The
message phrases shown are typical, but any human-readable alternative may be
provided. Unless otherwise stated, the status code is part of the HTTP/1.1
standard (RFC 7231).
## 4.5 Callback URL
Call back URL is an option to automatically inform CUSTOMER when a status
change occurs for a purchase order. The callback URL is a great option to
automatically track any status change to a purchase order. It is advised to implement it.
## 4.6 First Level Support
It is the initial support service provided by CUSTOMER to its clients. It
requires knowledge of products, terms and conditions offered by CY.TALK
rather than technical information.
## 4.7 Second Level Support
It is more in-depth technical support service provided by CY.TALK for
CUSTOMER support rather than direct CUSTOMER's client support.
## 4.8 CUSTOMER
The CUSTOMER is a registered CY.SEND user.
## 4.9 Products
It is the service delivered to the CUSTOMER’s client. Each product has a Face value and a currency. They can be in fixed values or in ranges. The products are divided into categories
## 4.10 Product types
* Instant product: is a type of product where the purchased top up amount will be credited instantly and directly to the beneficiary’s balance. Example: for a mobile top up where the top up is instantly credited to the beneficiary’s balance, or to top up an utility account where the amount directly credits the beneficiary’s balance.
* Prepaid code: is a type of product delivered to the beneficiary as a prepaid code with usage instructions. It is a series of number or letters or both that is loaded with a prepaid amount. To redeem the prepaid code, the beneficiary has to follow the sent usage instructions. Example: to visit the website of the issuing company and follow the shown instructions, or to send a USSD code in case of a mobile operator prepaid. Each prepaid code has a serial number ( digits) that is used to help identify the prepaid code. The serial number cannot be redeemed. Prepaid codes usually have an expiration date.
## 4.11 Order
It's CUSTOMER's purchase order, in which CUSTOMER will add products and place the purchase. For now it is limited to 1 product per order (later will be increased).
## 4.12 Product issuer purchase reference number
It is a number provide by some product issuers that can be used by the beneficiary for his customer care enquiries when contacting the product issuer.
## 4.13 Notifications
Use notification to be proactively informed of event that happened based on the rules you have set to be sent using the notification channel that suits you. Example: to retrieve your balance or to know the status change of a purchase.  We recommend you set notifications to keep track at all times of the status of your business.
<br /><br />
There are two ways you can use the notifications:
* The ones you can subscribe for a resource catalogue.
* The ones for specific resource.
### 4.13.1 Resource
A resource is an object within a resource type. Such as: a promotion, a maintenance, balance, a purchase, a price lists, a news, credentials, a gift cards, security and a sale.
### 4.13.2 Resource type
Each resource type is a homogeneous group that contains only one type of resource. You can subscribe to a resources type.
<br />Example of resource types are: Promotions, Maintenances, Balances, Purchases, Price lists, News, Credentials, Gift Cards, Security and Sales
### 4.13.3 Event
It is an occurrence that happens that leads to set a trigger off such as:
* A balance was debited
* A transaction was successful
* Start of a promotion
### 4.13.4 Trigger
A mechanism that initiates an action when an event occurs or upon receiving some type of input such as:
* A low balance threshold
* A change in top up status
### 4.13.7 Channel
A channel is the means that will be used to convey the notifications. Optionally CUSTOMER can create channels to send notifications in white label; I.E.: using his own email, firebase, Mattermost server, or own SMS sender name.
### 4.13.6 Channel type
It is the transport medium that will be used to send the notifications:
* Email
* SMS
* HTTP
* Firebase
* Phone call
* XMPP
* Mattermost
### 4.13.8 Template
It is the content of notifications. Each channel has his is own set of templates which are specific per channel type and per trigger. Templates are built using a pseudo language that allow the usage of variables describing the notification resource and trigger.
### 4.13.9 Subscription
CUSTOMERS can subscribe to an event by resource type or by resource, you will be notified each time the event you subscribed to sets off a trigger.
<br />I.E.: a CUSTOMER can subscribe to a purchase result notification by email for every purchase made from his account (it’s the default subscription for all new CY.SEND CUSTOMERS).
<br />Or subscribe to a specific purchase result notification by HTTP for a single purchase using the  CY.SEND OpenAPI POST [/store/order](#/Store/createOrder) operation.
### 4.13.10 Destination
It is to where the notification will be sent. For each notification subscription, the CUSTOMER should specify the destination and the parameters that are required by the Channel Type.
<br />I.E.: Email channels required `to_emails` parameter and have 2 optional parameters that are `cc_emails` and `bcc_emails`.
<br />HTTP channels only require an `url` parameter.
## 4.14 Purchase Scenarios
The CY.SEND OpenAPI use scenarios so you can test your API implementations. There are 2 scenarios: LIVE and SIMULATION.
### 4.14.1 Live purchase scenario
Will buy top up and DEBIT your CY.SEND account balance.
### 4.14.2 Simulate purchase scenario
Will do a simulation of a top up, your CY.SEND balance will NOT be debited.
## 4.15 API credentials
You will need them to use the CY.SEND OpenAPI, they are found in your profile in the Developer section. You can have as many credentials as you need for each of your use cases, passwords are shown only one time (make sure you save them).
# 5 API Structure
## 5.1 Scopes
The API is separated in multiple scopes that correspond to CY.SEND
resources. A scope is always the first part (before the second slash) of an
operation path. For example, the operation to retrieve catalogue countries
has the path [/store/catalogue/countries](#/Store/getCountries). The scope of this operation is
`Store`.
## 5.2 Methods
Every scope has defined HTTP request methods.
## 5.3 Operation
An operation is defined by a relative URL (containing the scope) and a
method.
## 5.4 CUSTOMER rights
CUSTOMER's credentials have access to some scope and methods. If the API
detect an unauthorized access, it will return an 403 HTTP status code.
## 5.5 Authentication
Every operation can have one or more authentication mechanisms. Only one
mechanism must be used per request.
### 5.5.1 HTTP basic access authentication
HTTP basic access authentication uses API credentials to authenticate the
CUSTOMER. For security reason, this authentication mechanism must always be
used over SSL request.
### 5.5.2 Access (Impersonation) token
The `access` scope allows the CUSTOMER to generate and manage tokens in order
to allow a third party (point of sale, application, partner, ...) to use
CY.SEND API on his behalf without knowing his API credentials. The token
usage can be restricted by source IP, time and accessible operations. Please
see POST [/access/token](#/Access/createAccessToken) operation for more information.
# 6 Technical requirements
Representational State Transfer architecture, Knowledge of HTTP requests and
JSON.
## 6.1 Technologies used
* HTTP request
* JSON for query results
# 7 Read before you start coding, this is the part you really should read before starting!
As we are developing our API and catalogue constantly, we suggest that you use the below logic for your integration since more products and features will be added based on this logic.
<br />We optimize so you will not need to rewrite your code.
1. Create a customer account on www.cysend.com.
2. Activate as many OpenAPI credentials as you need from the developer section in your profile. The password will be shown only one time, make sure you save it.
<br />Note: You must never send this password in clear text over the internet. This includes public chats such as Skype, WhatsApp, email…. We are not responsible if you share your password in the open.
3. Pull the CY.SEND catalogue to build your own catalogue by:
    1. Use the /store/catalogue operation to build your database.
    <br />This will return:
        1. All countries where CY.SEND has products.
        2. All product categories
        3. All products name
        4. All products face values and currencies
        5. All products price
    2.  Create a database with all of the above – You should update it once a day after 11:00 (GMT +1) to keep the prices and products current, else you might have error when doing purchases.
    <br />NOTE: The face value ID are changed each time the value changes but they are NOT changed when the cost changes. This is valid for ranges and fixed values.
    3. Create your retail sales price list
    4. Set a callback to receive top up status update
4. Testing purchases
<br />The OpenAPI can send simulations and real top ups. The simulation scenario will function exactly as a real top up but it will not debit your CY.SEND balance. It is to simulate a real top up. If you want to send a real top up where you buy a product do no use the simulate scenario
    1. Choose a product to buy
    2. Set a callback to receive the order statutes
    3. Place the test purchase order using the Simulate scenarios
    4. Wait for the callback top up status
    <br />Note: Purchases are asynchronous, you must save all the purchase orders you send so you can update them when you receive the callback with the new order status.
5. Let our team know when you are ready to go live so we validate the implementation.
6. When all validated, credit you CY.SEND balance with money.
7. Use the  purchase order operation with scenario LIVE to buy products.
# 8 Integration is done, how do we go live?
Congratulations! You are close to starting your top up business. Follow the below steps:
1. Your account should have positive balance
2. Use the POST [/store/order](#/Store/createOrder) operation with scenario live and it will send real top ups
3. If you have any question before starting please contact our support.
# 9 Security
Various security mechanisms protect our infrastructure and the CUSTOMER. If
suspicious activity is detected, cysend.com systems will automatically ban
and terminate connections.
# 10 HTTP Status Codes
Code | Text | Description
----------- | ----------- | -----------
200 | OK | Success!
201 | Created | The request has been fulfilled and resulted in a new resource being created.
202 | Accepted | The request has been accepted for processing, but the processing has not been completed.<br /><br />Note: All purchase orders are asynchronous. It means that when CUSTOMER executes a purchase order, the CY.SEND<br />OpenAPI replies “accepted” and in a separate background process the purchase order is sent to the<br />product issuer. A purchase order will remain in the “accepted” state until it receives a reply from the product issuer. This HTTP status code is not the final result of an order. To know the final result, wait for the callback URL request or pull the status using GET `/store/order` operation.
204 | No Content | The result of the search is empty.
400 | Bad Request | The request was invalid or cannot be otherwise served. An accompanying error message will explain further.<br />Placing an order with missing beneficiary information will result in this status.
401 | Unauthorized | Missing or incorrect authentication credentials.
402 | Payment Required | The order can't be placed because of insufficient balance. Please top up your account.
404 | Not Found | The URI requested is invalid or the resource requested, such as an order, does not exist.
500 | Internal Server Error | Something is broken. This is usually a temporary error, for example in a high load<br />situation or if an endpoint is temporarily having issues. Try again later or contact support.
502 | Bad Gateway | CY.SEND OpenAPI is down, or being upgraded.
503 | Service Unavailable | The CY.SEND OpenAPI servers are up, but overloaded with requests. Try again later or contact support.
504 | Gateway timeout | The CY.SEND OpenAPI servers are up, but the request couldn’t be serviced due to some failure<br />within the internal stack. Try again later or contact support.
# 11 About us


## Overview
This API client was generated by the [OpenAPI Generator](https://openapi-generator.tech) project.  By using the [OpenAPI-spec](https://www.openapis.org/) from a remote server, you can easily generate an API client.

- API version: 5.2.3
- Package version: 1.0.0
- Build package: org.openapitools.codegen.languages.GoClientCodegen
For more information, please visit [https://www.cysend.com](https://www.cysend.com)

## Installation

Install the following dependencies:

```sh
go get github.com/stretchr/testify/assert
go get golang.org/x/net/context
```

Put the package under your project folder and add the following in import:

```go
import cysend "github.com/pourcheragh/cysend-go"
```

To use a proxy, set the environment variable `HTTP_PROXY`:

```go
os.Setenv("HTTP_PROXY", "http://proxy_name:proxy_port")
```

## Configuration of Server URL

Default configuration comes with `Servers` field that contains server objects as defined in the OpenAPI specification.

### Select Server Configuration

For using other server than the one defined on index 0 set context value `cysend.ContextServerIndex` of type `int`.

```go
ctx := context.WithValue(context.Background(), cysend.ContextServerIndex, 1)
```

### Templated Server URL

Templated server URL is formatted using default variables from configuration or from context value `cysend.ContextServerVariables` of type `map[string]string`.

```go
ctx := context.WithValue(context.Background(), cysend.ContextServerVariables, map[string]string{
	"basePath": "v2",
})
```

Note, enum values are always validated and all unused variables are silently ignored.

### URLs Configuration per Operation

Each operation can use different server URL defined using `OperationServers` map in the `Configuration`.
An operation is uniquely identified by `"{classname}Service.{nickname}"` string.
Similar rules for overriding default operation server index and variables applies by using `cysend.ContextOperationServerIndices` and `cysend.ContextOperationServerVariables` context maps.

```go
ctx := context.WithValue(context.Background(), cysend.ContextOperationServerIndices, map[string]int{
	"{classname}Service.{nickname}": 2,
})
ctx = context.WithValue(context.Background(), cysend.ContextOperationServerVariables, map[string]map[string]string{
	"{classname}Service.{nickname}": {
		"port": "8443",
	},
})
```

## Documentation for API Endpoints

All URIs are relative to *https://www.cysend.com/openapi/prod/v5.2.3*

Class | Method | HTTP request | Description
------------ | ------------- | ------------- | -------------
*AccessAPI* | [**CreateAccessToken**](docs/AccessAPI.md#createaccesstoken) | **Post** /access/token | Creates a new access token
*AccessAPI* | [**GetAccessToken**](docs/AccessAPI.md#getaccesstoken) | **Get** /access/token | Retrieves an existing access token
*CYSENDGiftCardAPI* | [**GetActive**](docs/CYSENDGiftCardAPI.md#getactive) | **Get** /gift-card/{gift_card_code}/active | Checks if a CY.SEND gift card is active
*CYSENDGiftCardAPI* | [**GetGiftCard**](docs/CYSENDGiftCardAPI.md#getgiftcard) | **Get** /gift-card/{gift_card_code} | Retrieves a CY.SEND gift card information by PIN code
*CYSENDGiftCardAPI* | [**PutActive**](docs/CYSENDGiftCardAPI.md#putactive) | **Put** /gift-card/{gift_card_code}/active | Activates or deactivates a CY.SEND gift card
*CustomerAPI* | [**Balance**](docs/CustomerAPI.md#balance) | **Get** /customer/balance | Retrieves customer account default balance
*CustomerAPI* | [**GetSubscriptions**](docs/CustomerAPI.md#getsubscriptions) | **Get** /customer/subscriptions | Retrieves the list of subscriptions
*NotificationsAPI* | [**GetNotificationChannels**](docs/NotificationsAPI.md#getnotificationchannels) | **Get** /notification/channels | Retrieve the list of notification channels
*NotificationsAPI* | [**GetNotificationEvents**](docs/NotificationsAPI.md#getnotificationevents) | **Get** /notification/events | Retrieve the list of notification events
*NotificationsAPI* | [**GetNotificationResourceTypes**](docs/NotificationsAPI.md#getnotificationresourcetypes) | **Get** /notification/resource_types | Retrieve the list of notification resource types
*NotificationsAPI* | [**GetSubscriptions**](docs/NotificationsAPI.md#getsubscriptions) | **Get** /customer/subscriptions | Retrieves the list of subscriptions
*OpenAPIAPI* | [**GetErrors**](docs/OpenAPIAPI.md#geterrors) | **Get** /openapi/errors | Retrieves the list of CY.SEND OpenAPI errors
*OpenAPIAPI* | [**Ping**](docs/OpenAPIAPI.md#ping) | **Get** /openapi/ping | CY.SEND OpenAPI connection status monitor
*StoreAPI* | [**CreateOrder**](docs/StoreAPI.md#createorder) | **Post** /store/order | Place a purchase order for CY.SEND products
*StoreAPI* | [**ForceCallback**](docs/StoreAPI.md#forcecallback) | **Post** /store/order/callback | Force a callback request
*StoreAPI* | [**GetCatalogue**](docs/StoreAPI.md#getcatalogue) | **Get** /store/catalogue | Retrieves the complete CY.SEND product catalogue
*StoreAPI* | [**GetCategories**](docs/StoreAPI.md#getcategories) | **Get** /store/catalogue/categories | Retrieves the CY.SEND product categories list from the catalogue
*StoreAPI* | [**GetCountries**](docs/StoreAPI.md#getcountries) | **Get** /store/catalogue/countries | Retrieves the CY.SEND product catalogue country list
*StoreAPI* | [**GetFaceValues**](docs/StoreAPI.md#getfacevalues) | **Get** /store/catalogue/face-value | Retrieves the CY.SEND product face value list
*StoreAPI* | [**GetMaintenance**](docs/StoreAPI.md#getmaintenance) | **Get** /store/catalogue/maintenance | Retrieves current and future product maintenances
*StoreAPI* | [**GetOrder**](docs/StoreAPI.md#getorder) | **Get** /store/order | (Optional) Retrieves order information by CY.SEND order unique ID or by CUSTOMER order unique ID
*StoreAPI* | [**GetProductCost**](docs/StoreAPI.md#getproductcost) | **Put** /store/order/cost | Retrieve the cost of a face value
*StoreAPI* | [**GetProducts**](docs/StoreAPI.md#getproducts) | **Get** /store/catalogue/products | Retrieves the CY.SEND product catalogue list
*StoreAPI* | [**GetPromotions**](docs/StoreAPI.md#getpromotions) | **Get** /store/catalogue/promotion | Retrieves current and future product promotions
*StoreAPI* | [**GetResponses**](docs/StoreAPI.md#getresponses) | **Get** /store/order/responses | Retrieve the list of possible order response codes
*VerificationAPI* | [**PhoneNumberVerificationRequest**](docs/VerificationAPI.md#phonenumberverificationrequest) | **Post** /verification/phone-number | Creates a phone number verification request
*VerificationAPI* | [**PhoneNumberVerify**](docs/VerificationAPI.md#phonenumberverify) | **Put** /verification/phone-number | Confirm a verification request


## Documentation For Models

 - [APIError](docs/APIError.md)
 - [APIErrorDescription](docs/APIErrorDescription.md)
 - [BeneficiaryInformation](docs/BeneficiaryInformation.md)
 - [Catalogue](docs/Catalogue.md)
 - [Category](docs/Category.md)
 - [CostRequest](docs/CostRequest.md)
 - [CostResponse](docs/CostResponse.md)
 - [Country](docs/Country.md)
 - [CustomerBalance](docs/CustomerBalance.md)
 - [FaceValue](docs/FaceValue.md)
 - [FixedFaceValue](docs/FixedFaceValue.md)
 - [ForceCallbackRequest](docs/ForceCallbackRequest.md)
 - [GiftCard](docs/GiftCard.md)
 - [GiftCardActive](docs/GiftCardActive.md)
 - [Maintenance](docs/Maintenance.md)
 - [MaintenanceDescriptionsInner](docs/MaintenanceDescriptionsInner.md)
 - [NotificationChannel](docs/NotificationChannel.md)
 - [NotificationDestinationParameter](docs/NotificationDestinationParameter.md)
 - [NotificationDestinationParameterDescription](docs/NotificationDestinationParameterDescription.md)
 - [NotificationEvent](docs/NotificationEvent.md)
 - [NotificationResourceType](docs/NotificationResourceType.md)
 - [OrderRequest](docs/OrderRequest.md)
 - [OrderResponseCode](docs/OrderResponseCode.md)
 - [OrderScenario](docs/OrderScenario.md)
 - [OrderStatus](docs/OrderStatus.md)
 - [Permission](docs/Permission.md)
 - [PrepaidCode](docs/PrepaidCode.md)
 - [Product](docs/Product.md)
 - [ProductParameter](docs/ProductParameter.md)
 - [ProductParameterElementsInner](docs/ProductParameterElementsInner.md)
 - [ProductProductDescriptionInner](docs/ProductProductDescriptionInner.md)
 - [ProductType](docs/ProductType.md)
 - [ProductUsageInstructionsInner](docs/ProductUsageInstructionsInner.md)
 - [Promotion](docs/Promotion.md)
 - [PromotionDescriptionsInner](docs/PromotionDescriptionsInner.md)
 - [PromotionTimeZone](docs/PromotionTimeZone.md)
 - [RangeFaceValue](docs/RangeFaceValue.md)
 - [Subscription](docs/Subscription.md)
 - [Token](docs/Token.md)
 - [TokenOptions](docs/TokenOptions.md)
 - [Verification](docs/Verification.md)
 - [VerificationMethod](docs/VerificationMethod.md)
 - [VerificationRequest](docs/VerificationRequest.md)
 - [VerificationStatus](docs/VerificationStatus.md)


## Documentation For Authorization


Authentication schemes defined for the API:
### BasicAuthentication

- **Type**: HTTP basic authentication

Example

```go
auth := context.WithValue(context.Background(), cysend.ContextBasicAuth, cysend.BasicAuth{
	UserName: "username",
	Password: "password",
})
r, err := client.Service.Operation(auth, args)
```

### Token

- **Type**: API key
- **API key parameter name**: Api-token
- **Location**: HTTP header

Note, each API key must be added to a map of `map[string]APIKey` where the key is: Api-token and passed in as the auth context for each request.

Example

```go
auth := context.WithValue(
		context.Background(),
		cysend.ContextAPIKeys,
		map[string]cysend.APIKey{
			"Api-token": {Key: "API_KEY_STRING"},
		},
	)
r, err := client.Service.Operation(auth, args)
```


## Documentation for Utility Methods

Due to the fact that model structure members are all pointers, this package contains
a number of utility functions to easily obtain pointers to values of basic types.
Each of these functions takes a value of the given basic type and returns a pointer to it:

* `PtrBool`
* `PtrInt`
* `PtrInt32`
* `PtrInt64`
* `PtrFloat`
* `Ptrdecimal.Decimal`
* `PtrFloat64`
* `PtrString`
* `PtrTime`

## Author



