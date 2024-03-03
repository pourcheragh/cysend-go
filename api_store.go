/*
CY.SEND® OpenAPI Documentation

# 1 Table of content 1. Table of content 2. Overview 3. Changelog 4. Definitions 5. API Structure 6. Technical requirements 7. Read before you start coding 8. Integration is done, how do we go live? 9. Security 10. HTTP Status Codes 11. About us # 2 Overview The cysend.com API is a gateway to buy all the products we carry.<br /> The cysend.com API is built following the OpenAPI standards.<br /> [read more](https://www.openapis.org/) ## 2.1 Audience This document contains the technical interconnection specifications. It is designed to be read by the CUSTOMER’s technical team in charge of the integration. ## 2.2 Uses The cysend.com API can be interfaced with different kinds of devices & hosts, such as: 1. Mobile phones 2. Terminals 3. CUSTOMER web sites 4. Automated points of sales 5. Automated teller machines 6. White label websites 7. White label applications 8. Custom applications & hardware # 3 Changelog Date | Version | Changes | Author ----------- | ----------- | ----------- | ----------- 2023-12-11 | 5.2.3 | <ul>Added parameters to receive PIN usage instructions in catalog and the possibility to filter the returned languages.<li>Added `return_usage_instruction` parameter in GET `/store/catalogue` and GET `/store/catalogue/products` operations.</li><li>Added `languages` parameter in GET `/store/catalogue` and GET `/store/catalogue/products` operations.</li><li>Added `usage_instructions` parameter to the `Product` schema. This parameter is optional and is only returned if new parameter `return_usage_instruction` is set to true.</li></ul> | CY.SEND OpenApi team 2022-11-05 | 5.2.2 | <ul><li>New operation `/store/order/cost`</li><br />This operation returns an indicative cost of a face value. It does not place a purchase order. It is only a calculation tool to know the cost of a face value before placing the final purchase order. It returns the current cost at the exact time of the operation.</ul> | CY.SEND OpenApi team 2022-09-08 | 5.2.1 | <ul><li>Added `country_zone` parameter to the `Product` schema<br />Returns an indication of the geographical zone in which the product can be used.</li></ul> | CY.SEND OpenApi team 2020-03-06 | 5.2.0 | <ul><li>Added mobile to catalogue operations `/store/catalogue/products` and `/store/catalogue/face-value`.<br />This filter will detect the operator and return the available product/face values.<br />The service has a cost and is billed on usage</li><li>Added notification operations. See Notification operation group and definitions (Chapter 4.13).</li><li>Added `notification` parameter in `/store/order` POST request.</li></ul> | CY.SEND OpenApi team 2019-10-08 | 5.1.3 | <ul><li>Added `product_description` parameter in catalogue.</li></ul> | CY.SEND OpenApi team 2019-05-24 | 5.1.2 | <ul><li>Added product types (`instant` and `prepaid_code`) to the `Product` schema</li><li>Added new definition to chapter 4: Product types</li><li>Renamed `PIN` schema to `PrepaidCode`</li><li>Added multi languages prepaid code usage instructions to the `PrepaidCode` schema</li><li>Renamed `Coverage` schema to `Catalogue`</li><li>Renamed `gift_card_pin_code` parameter to `gift_card_code` in `CY.SEND Gift Card` operations</li></ul> | CY.SEND OpenApi team 2019-05-14 | 5.1.1 | <ul><li>Added support for promotions in `Catalogue` operations</li><li>Renamed `Catalogue` schema parameter `maintenance` to `maintenances`</li><li>Added purchase `date` parameter in `OrderStatus` schema</li><li>Renamed `pin-code` parameter to `gift_card_pin_code` in `CY.SEND Gift Card` operations</li></ul> | CY.SEND OpenApi team 2019-03-19 | 5.1.0 | <ul><li>Public early access</li></ul> | CY.SEND OpenApi team  # 4 Definitions ## 4.1 API Credentials API credentials are different from cysend.com credentials. Please ask the support team for your credentials. ## 4.2 HTTP Basic access authentication In the context of an HTTP transaction, basic access authentication is a method for an HTTP user agent to provide a user name and password when making a request. For security reason, this authentication method should only be used on HTTPS request. [Wikipedia](https://en.wikipedia.org/wiki/Basic_access_authentication) ## 4.3 HTTP request methods HTTP defines methods (sometimes referred to as \"verbs\") to indicate the desired action to be performed on the identified resource. ## 4.4 HTTP status code Status codes are issued by a server in response to a client's request made to the server. It includes codes from IETF Request for Comments (RFCs), other specifications, and some additional codes used in some common applications of the Hypertext Transfer Protocol (HTTP). The first digit of the status code specifies one of five standard classes of responses. The message phrases shown are typical, but any human-readable alternative may be provided. Unless otherwise stated, the status code is part of the HTTP/1.1 standard (RFC 7231). ## 4.5 Callback URL Call back URL is an option to automatically inform CUSTOMER when a status change occurs for a purchase order. The callback URL is a great option to automatically track any status change to a purchase order. It is advised to implement it. ## 4.6 First Level Support It is the initial support service provided by CUSTOMER to its clients. It requires knowledge of products, terms and conditions offered by CY.TALK rather than technical information. ## 4.7 Second Level Support It is more in-depth technical support service provided by CY.TALK for CUSTOMER support rather than direct CUSTOMER's client support. ## 4.8 CUSTOMER The CUSTOMER is a registered CY.SEND user. ## 4.9 Products It is the service delivered to the CUSTOMER’s client. Each product has a Face value and a currency. They can be in fixed values or in ranges. The products are divided into categories ## 4.10 Product types * Instant product: is a type of product where the purchased top up amount will be credited instantly and directly to the beneficiary’s balance. Example: for a mobile top up where the top up is instantly credited to the beneficiary’s balance, or to top up an utility account where the amount directly credits the beneficiary’s balance. * Prepaid code: is a type of product delivered to the beneficiary as a prepaid code with usage instructions. It is a series of number or letters or both that is loaded with a prepaid amount. To redeem the prepaid code, the beneficiary has to follow the sent usage instructions. Example: to visit the website of the issuing company and follow the shown instructions, or to send a USSD code in case of a mobile operator prepaid. Each prepaid code has a serial number ( digits) that is used to help identify the prepaid code. The serial number cannot be redeemed. Prepaid codes usually have an expiration date. ## 4.11 Order It's CUSTOMER's purchase order, in which CUSTOMER will add products and place the purchase. For now it is limited to 1 product per order (later will be increased). ## 4.12 Product issuer purchase reference number It is a number provide by some product issuers that can be used by the beneficiary for his customer care enquiries when contacting the product issuer. ## 4.13 Notifications Use notification to be proactively informed of event that happened based on the rules you have set to be sent using the notification channel that suits you. Example: to retrieve your balance or to know the status change of a purchase.  We recommend you set notifications to keep track at all times of the status of your business. <br /><br /> There are two ways you can use the notifications: * The ones you can subscribe for a resource catalogue. * The ones for specific resource. ### 4.13.1 Resource A resource is an object within a resource type. Such as: a promotion, a maintenance, balance, a purchase, a price lists, a news, credentials, a gift cards, security and a sale. ### 4.13.2 Resource type Each resource type is a homogeneous group that contains only one type of resource. You can subscribe to a resources type. <br />Example of resource types are: Promotions, Maintenances, Balances, Purchases, Price lists, News, Credentials, Gift Cards, Security and Sales ### 4.13.3 Event It is an occurrence that happens that leads to set a trigger off such as: * A balance was debited * A transaction was successful * Start of a promotion ### 4.13.4 Trigger A mechanism that initiates an action when an event occurs or upon receiving some type of input such as: * A low balance threshold * A change in top up status ### 4.13.7 Channel A channel is the means that will be used to convey the notifications. Optionally CUSTOMER can create channels to send notifications in white label; I.E.: using his own email, firebase, Mattermost server, or own SMS sender name. ### 4.13.6 Channel type It is the transport medium that will be used to send the notifications: * Email * SMS * HTTP * Firebase * Phone call * XMPP * Mattermost ### 4.13.8 Template It is the content of notifications. Each channel has his is own set of templates which are specific per channel type and per trigger. Templates are built using a pseudo language that allow the usage of variables describing the notification resource and trigger. ### 4.13.9 Subscription CUSTOMERS can subscribe to an event by resource type or by resource, you will be notified each time the event you subscribed to sets off a trigger. <br />I.E.: a CUSTOMER can subscribe to a purchase result notification by email for every purchase made from his account (it’s the default subscription for all new CY.SEND CUSTOMERS). <br />Or subscribe to a specific purchase result notification by HTTP for a single purchase using the  CY.SEND OpenAPI POST [/store/order](#/Store/createOrder) operation. ### 4.13.10 Destination It is to where the notification will be sent. For each notification subscription, the CUSTOMER should specify the destination and the parameters that are required by the Channel Type. <br />I.E.: Email channels required `to_emails` parameter and have 2 optional parameters that are `cc_emails` and `bcc_emails`. <br />HTTP channels only require an `url` parameter. ## 4.14 Purchase Scenarios The CY.SEND OpenAPI use scenarios so you can test your API implementations. There are 2 scenarios: LIVE and SIMULATION. ### 4.14.1 Live purchase scenario Will buy top up and DEBIT your CY.SEND account balance. ### 4.14.2 Simulate purchase scenario Will do a simulation of a top up, your CY.SEND balance will NOT be debited. ## 4.15 API credentials You will need them to use the CY.SEND OpenAPI, they are found in your profile in the Developer section. You can have as many credentials as you need for each of your use cases, passwords are shown only one time (make sure you save them). # 5 API Structure ## 5.1 Scopes The API is separated in multiple scopes that correspond to CY.SEND resources. A scope is always the first part (before the second slash) of an operation path. For example, the operation to retrieve catalogue countries has the path [/store/catalogue/countries](#/Store/getCountries). The scope of this operation is `Store`. ## 5.2 Methods Every scope has defined HTTP request methods. ## 5.3 Operation An operation is defined by a relative URL (containing the scope) and a method. ## 5.4 CUSTOMER rights CUSTOMER's credentials have access to some scope and methods. If the API detect an unauthorized access, it will return an 403 HTTP status code. ## 5.5 Authentication Every operation can have one or more authentication mechanisms. Only one mechanism must be used per request. ### 5.5.1 HTTP basic access authentication HTTP basic access authentication uses API credentials to authenticate the CUSTOMER. For security reason, this authentication mechanism must always be used over SSL request. ### 5.5.2 Access (Impersonation) token The `access` scope allows the CUSTOMER to generate and manage tokens in order to allow a third party (point of sale, application, partner, ...) to use CY.SEND API on his behalf without knowing his API credentials. The token usage can be restricted by source IP, time and accessible operations. Please see POST [/access/token](#/Access/createAccessToken) operation for more information. # 6 Technical requirements Representational State Transfer architecture, Knowledge of HTTP requests and JSON. ## 6.1 Technologies used * HTTP request * JSON for query results # 7 Read before you start coding, this is the part you really should read before starting! As we are developing our API and catalogue constantly, we suggest that you use the below logic for your integration since more products and features will be added based on this logic. <br />We optimize so you will not need to rewrite your code. 1. Create a customer account on www.cysend.com. 2. Activate as many OpenAPI credentials as you need from the developer section in your profile. The password will be shown only one time, make sure you save it. <br />Note: You must never send this password in clear text over the internet. This includes public chats such as Skype, WhatsApp, email…. We are not responsible if you share your password in the open. 3. Pull the CY.SEND catalogue to build your own catalogue by:     1. Use the /store/catalogue operation to build your database.     <br />This will return:         1. All countries where CY.SEND has products.         2. All product categories         3. All products name         4. All products face values and currencies         5. All products price     2.  Create a database with all of the above – You should update it once a day after 11:00 (GMT +1) to keep the prices and products current, else you might have error when doing purchases.     <br />NOTE: The face value ID are changed each time the value changes but they are NOT changed when the cost changes. This is valid for ranges and fixed values.     3. Create your retail sales price list     4. Set a callback to receive top up status update 4. Testing purchases <br />The OpenAPI can send simulations and real top ups. The simulation scenario will function exactly as a real top up but it will not debit your CY.SEND balance. It is to simulate a real top up. If you want to send a real top up where you buy a product do no use the simulate scenario     1. Choose a product to buy     2. Set a callback to receive the order statutes     3. Place the test purchase order using the Simulate scenarios     4. Wait for the callback top up status     <br />Note: Purchases are asynchronous, you must save all the purchase orders you send so you can update them when you receive the callback with the new order status. 5. Let our team know when you are ready to go live so we validate the implementation. 6. When all validated, credit you CY.SEND balance with money. 7. Use the  purchase order operation with scenario LIVE to buy products. # 8 Integration is done, how do we go live? Congratulations! You are close to starting your top up business. Follow the below steps: 1. Your account should have positive balance 2. Use the POST [/store/order](#/Store/createOrder) operation with scenario live and it will send real top ups 3. If you have any question before starting please contact our support. # 9 Security Various security mechanisms protect our infrastructure and the CUSTOMER. If suspicious activity is detected, cysend.com systems will automatically ban and terminate connections. # 10 HTTP Status Codes Code | Text | Description ----------- | ----------- | ----------- 200 | OK | Success! 201 | Created | The request has been fulfilled and resulted in a new resource being created. 202 | Accepted | The request has been accepted for processing, but the processing has not been completed.<br /><br />Note: All purchase orders are asynchronous. It means that when CUSTOMER executes a purchase order, the CY.SEND<br />OpenAPI replies “accepted” and in a separate background process the purchase order is sent to the<br />product issuer. A purchase order will remain in the “accepted” state until it receives a reply from the product issuer. This HTTP status code is not the final result of an order. To know the final result, wait for the callback URL request or pull the status using GET `/store/order` operation. 204 | No Content | The result of the search is empty. 400 | Bad Request | The request was invalid or cannot be otherwise served. An accompanying error message will explain further.<br />Placing an order with missing beneficiary information will result in this status. 401 | Unauthorized | Missing or incorrect authentication credentials. 402 | Payment Required | The order can't be placed because of insufficient balance. Please top up your account. 404 | Not Found | The URI requested is invalid or the resource requested, such as an order, does not exist. 500 | Internal Server Error | Something is broken. This is usually a temporary error, for example in a high load<br />situation or if an endpoint is temporarily having issues. Try again later or contact support. 502 | Bad Gateway | CY.SEND OpenAPI is down, or being upgraded. 503 | Service Unavailable | The CY.SEND OpenAPI servers are up, but overloaded with requests. Try again later or contact support. 504 | Gateway timeout | The CY.SEND OpenAPI servers are up, but the request couldn’t be serviced due to some failure<br />within the internal stack. Try again later or contact support. # 11 About us 

API version: 5.2.3
*/

// Code generated by OpenAPI Generator (https://openapi-generator.tech); DO NOT EDIT.

package cysend

import (
	"bytes"
	"context"
	"io"
	"net/http"
	"net/url"
)


// StoreAPIService StoreAPI service
type StoreAPIService service

type ApiCreateOrderRequest struct {
	ctx context.Context
	ApiService *StoreAPIService
	orderRequest *[]OrderRequest
}

// Create purchase order request parameters
func (r ApiCreateOrderRequest) OrderRequest(orderRequest []OrderRequest) ApiCreateOrderRequest {
	r.orderRequest = &orderRequest
	return r
}

func (r ApiCreateOrderRequest) Execute() ([]OrderStatus, *http.Response, error) {
	return r.ApiService.CreateOrderExecute(r)
}

/*
CreateOrder Place a purchase order for CY.SEND products

This operation creates a purchase order for products and executes it asynchronously.
<br />
To purchase a product, CUSTOMER must:
* Select a product. [/store/catalogue/products](#/Store/getProducts)
* Select a face value (fixed or range). [/store/catalogue/face-value](#/Store/getFaceValues)
* Insure that you have the required beneficiary information [/store/catalogue/products](#/Store/getProducts)
* Assign a CUSTOMER unique order ID
## API order scenarios
All purchase orders are asynchronous. It means that when CUSTOMER executes a purchase order, the CY.SEND
OpenAPI replies “accepted” and in a separate background process the purchase order is sent to the
product issuer. A purchase order will remain in the “accepted” state until it
receives a reply from the product issuer.
To test your callback script, create an order with a scenario and use the [callback operation](#/Store/forceCallback).
* simulate-success: This scenario will simulate a successful purchase order. To retrieve the status of your
order, use a callback.
* simulate-delayed-success: This scenario will simulate a delayed successful order purchase. The delay is
ranging from 60 to 120 seconds. To retrieve the status of your order, use a callback.
* simulate-failed: This scenario will simulate a failed order purchase. To retrieve the status of your
order, use a callback.
* simulate-delayed-failed: This scenario will failed a delayed successful order purchase. The delay is ranging
from 60 to 120 seconds. To retrieve the status of your order, use a callback.
* live: This is not a simulation. The request will be processed and your balance debited. To retrieve the
status of your order, use a callback.
## Callback
The HTTP callback body will contain the purchase order information in JSON format.
<br />
It is strongly recommended to set a callback (callback_url parameter) for each order to automatically be
notified on the URL of your choice when an order changes status even if it comes after a long time.
---
The `status` of an order can be:
* accepted: The purchase is still in progress. This means that we are waiting confirmation from the product
issuer or that CY.SEND delayed the transaction for security reasons.
* success: The  purchase was executed successfully.
* failed: The purchase failed. Additional failure information is returned in `response_code`, `response_title`
and `response_description` parameters.
---
For a successful top up, some product issuers provide a purchase reference number (`issuer_reference` parameter)
that can be used by the beneficiary for his customer care enquiries when contacting the product issuer.
<br />
<br />
If the purchase order contains a prepaid code product, the response will contain the prepaid code, the prepaid
code serial number, the prepaid code expiration and the prepaid code usage instructions (`prepaid_code` parameter).
<br />
<br />
Response codes are based on [ISO8583](https://en.wikipedia.org/wiki/ISO_8583#Response_code) codes.
<br />
<br />
Using a cron job to periodically check if there is an order status change is a waste of resources on both
sides. If you do not have the choice you can use the Retrieve an order operation.


 @param ctx context.Context - for authentication, logging, cancellation, deadlines, tracing, etc. Passed from http.Request or context.Background().
 @return ApiCreateOrderRequest
*/
func (a *StoreAPIService) CreateOrder(ctx context.Context) ApiCreateOrderRequest {
	return ApiCreateOrderRequest{
		ApiService: a,
		ctx: ctx,
	}
}

// Execute executes the request
//  @return []OrderStatus
func (a *StoreAPIService) CreateOrderExecute(r ApiCreateOrderRequest) ([]OrderStatus, *http.Response, error) {
	var (
		localVarHTTPMethod   = http.MethodPost
		localVarPostBody     interface{}
		formFiles            []formFile
		localVarReturnValue  []OrderStatus
	)

	localBasePath, err := a.client.cfg.ServerURLWithContext(r.ctx, "StoreAPIService.CreateOrder")
	if err != nil {
		return localVarReturnValue, nil, &GenericOpenAPIError{error: err.Error()}
	}

	localVarPath := localBasePath + "/store/order"

	localVarHeaderParams := make(map[string]string)
	localVarQueryParams := url.Values{}
	localVarFormParams := url.Values{}
	if r.orderRequest == nil {
		return localVarReturnValue, nil, reportError("orderRequest is required and must be specified")
	}

	// to determine the Content-Type header
	localVarHTTPContentTypes := []string{"application/json"}

	// set Content-Type header
	localVarHTTPContentType := selectHeaderContentType(localVarHTTPContentTypes)
	if localVarHTTPContentType != "" {
		localVarHeaderParams["Content-Type"] = localVarHTTPContentType
	}

	// to determine the Accept header
	localVarHTTPHeaderAccepts := []string{"application/json"}

	// set Accept header
	localVarHTTPHeaderAccept := selectHeaderAccept(localVarHTTPHeaderAccepts)
	if localVarHTTPHeaderAccept != "" {
		localVarHeaderParams["Accept"] = localVarHTTPHeaderAccept
	}
	// body params
	localVarPostBody = r.orderRequest
	if r.ctx != nil {
		// API Key Authentication
		if auth, ok := r.ctx.Value(ContextAPIKeys).(map[string]APIKey); ok {
			if apiKey, ok := auth["Token"]; ok {
				var key string
				if apiKey.Prefix != "" {
					key = apiKey.Prefix + " " + apiKey.Key
				} else {
					key = apiKey.Key
				}
				localVarHeaderParams["Api-token"] = key
			}
		}
	}
	req, err := a.client.prepareRequest(r.ctx, localVarPath, localVarHTTPMethod, localVarPostBody, localVarHeaderParams, localVarQueryParams, localVarFormParams, formFiles)
	if err != nil {
		return localVarReturnValue, nil, err
	}

	localVarHTTPResponse, err := a.client.callAPI(req)
	if err != nil || localVarHTTPResponse == nil {
		return localVarReturnValue, localVarHTTPResponse, err
	}

	localVarBody, err := io.ReadAll(localVarHTTPResponse.Body)
	localVarHTTPResponse.Body.Close()
	localVarHTTPResponse.Body = io.NopCloser(bytes.NewBuffer(localVarBody))
	if err != nil {
		return localVarReturnValue, localVarHTTPResponse, err
	}

	if localVarHTTPResponse.StatusCode >= 300 {
		newErr := &GenericOpenAPIError{
			body:  localVarBody,
			error: localVarHTTPResponse.Status,
		}
		if localVarHTTPResponse.StatusCode == 400 {
			var v APIError
			err = a.client.decode(&v, localVarBody, localVarHTTPResponse.Header.Get("Content-Type"))
			if err != nil {
				newErr.error = err.Error()
				return localVarReturnValue, localVarHTTPResponse, newErr
			}
					newErr.error = formatErrorMessage(localVarHTTPResponse.Status, &v)
					newErr.model = v
			return localVarReturnValue, localVarHTTPResponse, newErr
		}
		if localVarHTTPResponse.StatusCode == 401 {
			var v APIError
			err = a.client.decode(&v, localVarBody, localVarHTTPResponse.Header.Get("Content-Type"))
			if err != nil {
				newErr.error = err.Error()
				return localVarReturnValue, localVarHTTPResponse, newErr
			}
					newErr.error = formatErrorMessage(localVarHTTPResponse.Status, &v)
					newErr.model = v
			return localVarReturnValue, localVarHTTPResponse, newErr
		}
		if localVarHTTPResponse.StatusCode == 402 {
			var v APIError
			err = a.client.decode(&v, localVarBody, localVarHTTPResponse.Header.Get("Content-Type"))
			if err != nil {
				newErr.error = err.Error()
				return localVarReturnValue, localVarHTTPResponse, newErr
			}
					newErr.error = formatErrorMessage(localVarHTTPResponse.Status, &v)
					newErr.model = v
			return localVarReturnValue, localVarHTTPResponse, newErr
		}
		if localVarHTTPResponse.StatusCode == 404 {
			var v APIError
			err = a.client.decode(&v, localVarBody, localVarHTTPResponse.Header.Get("Content-Type"))
			if err != nil {
				newErr.error = err.Error()
				return localVarReturnValue, localVarHTTPResponse, newErr
			}
					newErr.error = formatErrorMessage(localVarHTTPResponse.Status, &v)
					newErr.model = v
		}
		return localVarReturnValue, localVarHTTPResponse, newErr
	}

	err = a.client.decode(&localVarReturnValue, localVarBody, localVarHTTPResponse.Header.Get("Content-Type"))
	if err != nil {
		newErr := &GenericOpenAPIError{
			body:  localVarBody,
			error: err.Error(),
		}
		return localVarReturnValue, localVarHTTPResponse, newErr
	}

	return localVarReturnValue, localVarHTTPResponse, nil
}

type ApiForceCallbackRequest struct {
	ctx context.Context
	ApiService *StoreAPIService
	forceCallbackRequest *ForceCallbackRequest
}

// Force callback request parameters
func (r ApiForceCallbackRequest) ForceCallbackRequest(forceCallbackRequest ForceCallbackRequest) ApiForceCallbackRequest {
	r.forceCallbackRequest = &forceCallbackRequest
	return r
}

func (r ApiForceCallbackRequest) Execute() (*http.Response, error) {
	return r.ApiService.ForceCallbackExecute(r)
}

/*
ForceCallback Force a callback request

This operation forces the CY.SEND OpenAPI to resend a callback request for a previously created order.
You can use this operation to test your callback request script.


 @param ctx context.Context - for authentication, logging, cancellation, deadlines, tracing, etc. Passed from http.Request or context.Background().
 @return ApiForceCallbackRequest
*/
func (a *StoreAPIService) ForceCallback(ctx context.Context) ApiForceCallbackRequest {
	return ApiForceCallbackRequest{
		ApiService: a,
		ctx: ctx,
	}
}

// Execute executes the request
func (a *StoreAPIService) ForceCallbackExecute(r ApiForceCallbackRequest) (*http.Response, error) {
	var (
		localVarHTTPMethod   = http.MethodPost
		localVarPostBody     interface{}
		formFiles            []formFile
	)

	localBasePath, err := a.client.cfg.ServerURLWithContext(r.ctx, "StoreAPIService.ForceCallback")
	if err != nil {
		return nil, &GenericOpenAPIError{error: err.Error()}
	}

	localVarPath := localBasePath + "/store/order/callback"

	localVarHeaderParams := make(map[string]string)
	localVarQueryParams := url.Values{}
	localVarFormParams := url.Values{}
	if r.forceCallbackRequest == nil {
		return nil, reportError("forceCallbackRequest is required and must be specified")
	}

	// to determine the Content-Type header
	localVarHTTPContentTypes := []string{"application/json"}

	// set Content-Type header
	localVarHTTPContentType := selectHeaderContentType(localVarHTTPContentTypes)
	if localVarHTTPContentType != "" {
		localVarHeaderParams["Content-Type"] = localVarHTTPContentType
	}

	// to determine the Accept header
	localVarHTTPHeaderAccepts := []string{"application/json"}

	// set Accept header
	localVarHTTPHeaderAccept := selectHeaderAccept(localVarHTTPHeaderAccepts)
	if localVarHTTPHeaderAccept != "" {
		localVarHeaderParams["Accept"] = localVarHTTPHeaderAccept
	}
	// body params
	localVarPostBody = r.forceCallbackRequest
	if r.ctx != nil {
		// API Key Authentication
		if auth, ok := r.ctx.Value(ContextAPIKeys).(map[string]APIKey); ok {
			if apiKey, ok := auth["Token"]; ok {
				var key string
				if apiKey.Prefix != "" {
					key = apiKey.Prefix + " " + apiKey.Key
				} else {
					key = apiKey.Key
				}
				localVarHeaderParams["Api-token"] = key
			}
		}
	}
	req, err := a.client.prepareRequest(r.ctx, localVarPath, localVarHTTPMethod, localVarPostBody, localVarHeaderParams, localVarQueryParams, localVarFormParams, formFiles)
	if err != nil {
		return nil, err
	}

	localVarHTTPResponse, err := a.client.callAPI(req)
	if err != nil || localVarHTTPResponse == nil {
		return localVarHTTPResponse, err
	}

	localVarBody, err := io.ReadAll(localVarHTTPResponse.Body)
	localVarHTTPResponse.Body.Close()
	localVarHTTPResponse.Body = io.NopCloser(bytes.NewBuffer(localVarBody))
	if err != nil {
		return localVarHTTPResponse, err
	}

	if localVarHTTPResponse.StatusCode >= 300 {
		newErr := &GenericOpenAPIError{
			body:  localVarBody,
			error: localVarHTTPResponse.Status,
		}
		if localVarHTTPResponse.StatusCode == 401 {
			var v APIError
			err = a.client.decode(&v, localVarBody, localVarHTTPResponse.Header.Get("Content-Type"))
			if err != nil {
				newErr.error = err.Error()
				return localVarHTTPResponse, newErr
			}
					newErr.error = formatErrorMessage(localVarHTTPResponse.Status, &v)
					newErr.model = v
			return localVarHTTPResponse, newErr
		}
		if localVarHTTPResponse.StatusCode == 404 {
			var v APIError
			err = a.client.decode(&v, localVarBody, localVarHTTPResponse.Header.Get("Content-Type"))
			if err != nil {
				newErr.error = err.Error()
				return localVarHTTPResponse, newErr
			}
					newErr.error = formatErrorMessage(localVarHTTPResponse.Status, &v)
					newErr.model = v
		}
		return localVarHTTPResponse, newErr
	}

	return localVarHTTPResponse, nil
}

type ApiGetCatalogueRequest struct {
	ctx context.Context
	ApiService *StoreAPIService
	giftCardCode *string
	returnUsageInstruction *bool
	languages *string
}

// CY.SEND gift card PIN code
func (r ApiGetCatalogueRequest) GiftCardCode(giftCardCode string) ApiGetCatalogueRequest {
	r.giftCardCode = &giftCardCode
	return r
}

// Set to true to receive PIN products usage instructions.
func (r ApiGetCatalogueRequest) ReturnUsageInstruction(returnUsageInstruction bool) ApiGetCatalogueRequest {
	r.returnUsageInstruction = &returnUsageInstruction
	return r
}

// List of language code in ISO 639-1 format separated by a coma.
func (r ApiGetCatalogueRequest) Languages(languages string) ApiGetCatalogueRequest {
	r.languages = &languages
	return r
}

func (r ApiGetCatalogueRequest) Execute() (*Catalogue, *http.Response, error) {
	return r.ApiService.GetCatalogueExecute(r)
}

/*
GetCatalogue Retrieves the complete CY.SEND product catalogue

This operation returns the full content of the CY.SEND product catalogue. <br /> The catalogue can be filtered specifically for a single CY.SEND gift card. Only products that can be purchased by this gift card will be returned.


 @param ctx context.Context - for authentication, logging, cancellation, deadlines, tracing, etc. Passed from http.Request or context.Background().
 @return ApiGetCatalogueRequest
*/
func (a *StoreAPIService) GetCatalogue(ctx context.Context) ApiGetCatalogueRequest {
	return ApiGetCatalogueRequest{
		ApiService: a,
		ctx: ctx,
	}
}

// Execute executes the request
//  @return Catalogue
func (a *StoreAPIService) GetCatalogueExecute(r ApiGetCatalogueRequest) (*Catalogue, *http.Response, error) {
	var (
		localVarHTTPMethod   = http.MethodGet
		localVarPostBody     interface{}
		formFiles            []formFile
		localVarReturnValue  *Catalogue
	)

	localBasePath, err := a.client.cfg.ServerURLWithContext(r.ctx, "StoreAPIService.GetCatalogue")
	if err != nil {
		return localVarReturnValue, nil, &GenericOpenAPIError{error: err.Error()}
	}

	localVarPath := localBasePath + "/store/catalogue"

	localVarHeaderParams := make(map[string]string)
	localVarQueryParams := url.Values{}
	localVarFormParams := url.Values{}

	if r.giftCardCode != nil {
		parameterAddToHeaderOrQuery(localVarQueryParams, "gift_card_code", r.giftCardCode, "")
	}
	if r.returnUsageInstruction != nil {
		parameterAddToHeaderOrQuery(localVarQueryParams, "return_usage_instruction", r.returnUsageInstruction, "")
	} else {
		var defaultValue bool = false
		r.returnUsageInstruction = &defaultValue
	}
	if r.languages != nil {
		parameterAddToHeaderOrQuery(localVarQueryParams, "languages", r.languages, "")
	}
	// to determine the Content-Type header
	localVarHTTPContentTypes := []string{}

	// set Content-Type header
	localVarHTTPContentType := selectHeaderContentType(localVarHTTPContentTypes)
	if localVarHTTPContentType != "" {
		localVarHeaderParams["Content-Type"] = localVarHTTPContentType
	}

	// to determine the Accept header
	localVarHTTPHeaderAccepts := []string{"application/json"}

	// set Accept header
	localVarHTTPHeaderAccept := selectHeaderAccept(localVarHTTPHeaderAccepts)
	if localVarHTTPHeaderAccept != "" {
		localVarHeaderParams["Accept"] = localVarHTTPHeaderAccept
	}
	if r.ctx != nil {
		// API Key Authentication
		if auth, ok := r.ctx.Value(ContextAPIKeys).(map[string]APIKey); ok {
			if apiKey, ok := auth["Token"]; ok {
				var key string
				if apiKey.Prefix != "" {
					key = apiKey.Prefix + " " + apiKey.Key
				} else {
					key = apiKey.Key
				}
				localVarHeaderParams["Api-token"] = key
			}
		}
	}
	req, err := a.client.prepareRequest(r.ctx, localVarPath, localVarHTTPMethod, localVarPostBody, localVarHeaderParams, localVarQueryParams, localVarFormParams, formFiles)
	if err != nil {
		return localVarReturnValue, nil, err
	}

	localVarHTTPResponse, err := a.client.callAPI(req)
	if err != nil || localVarHTTPResponse == nil {
		return localVarReturnValue, localVarHTTPResponse, err
	}

	localVarBody, err := io.ReadAll(localVarHTTPResponse.Body)
	localVarHTTPResponse.Body.Close()
	localVarHTTPResponse.Body = io.NopCloser(bytes.NewBuffer(localVarBody))
	if err != nil {
		return localVarReturnValue, localVarHTTPResponse, err
	}

	if localVarHTTPResponse.StatusCode >= 300 {
		newErr := &GenericOpenAPIError{
			body:  localVarBody,
			error: localVarHTTPResponse.Status,
		}
		if localVarHTTPResponse.StatusCode == 401 {
			var v APIError
			err = a.client.decode(&v, localVarBody, localVarHTTPResponse.Header.Get("Content-Type"))
			if err != nil {
				newErr.error = err.Error()
				return localVarReturnValue, localVarHTTPResponse, newErr
			}
					newErr.error = formatErrorMessage(localVarHTTPResponse.Status, &v)
					newErr.model = v
			return localVarReturnValue, localVarHTTPResponse, newErr
		}
		if localVarHTTPResponse.StatusCode == 404 {
			var v APIError
			err = a.client.decode(&v, localVarBody, localVarHTTPResponse.Header.Get("Content-Type"))
			if err != nil {
				newErr.error = err.Error()
				return localVarReturnValue, localVarHTTPResponse, newErr
			}
					newErr.error = formatErrorMessage(localVarHTTPResponse.Status, &v)
					newErr.model = v
		}
		return localVarReturnValue, localVarHTTPResponse, newErr
	}

	err = a.client.decode(&localVarReturnValue, localVarBody, localVarHTTPResponse.Header.Get("Content-Type"))
	if err != nil {
		newErr := &GenericOpenAPIError{
			body:  localVarBody,
			error: err.Error(),
		}
		return localVarReturnValue, localVarHTTPResponse, newErr
	}

	return localVarReturnValue, localVarHTTPResponse, nil
}

type ApiGetCategoriesRequest struct {
	ctx context.Context
	ApiService *StoreAPIService
	countryId *int32
	countryIso2 *string
	categoryId *int32
	promotion *bool
	giftCardCode *string
}

// Country ID filter
func (r ApiGetCategoriesRequest) CountryId(countryId int32) ApiGetCategoriesRequest {
	r.countryId = &countryId
	return r
}

// Country ISO 3166-1 alpha-2 filter
func (r ApiGetCategoriesRequest) CountryIso2(countryIso2 string) ApiGetCategoriesRequest {
	r.countryIso2 = &countryIso2
	return r
}

// Category ID filter
func (r ApiGetCategoriesRequest) CategoryId(categoryId int32) ApiGetCategoriesRequest {
	r.categoryId = &categoryId
	return r
}

// Promotion filter
func (r ApiGetCategoriesRequest) Promotion(promotion bool) ApiGetCategoriesRequest {
	r.promotion = &promotion
	return r
}

// CY.SEND gift card PIN code
func (r ApiGetCategoriesRequest) GiftCardCode(giftCardCode string) ApiGetCategoriesRequest {
	r.giftCardCode = &giftCardCode
	return r
}

func (r ApiGetCategoriesRequest) Execute() ([]Category, *http.Response, error) {
	return r.ApiService.GetCategoriesExecute(r)
}

/*
GetCategories Retrieves the CY.SEND product categories list from the catalogue

This operation retrieves the list of categories where CY.SEND products are available.
<br />
The product category list can be filtered by:
* CY.SEND country ID
* Country ISO 3166-1 alpha-2
* CY.SEND category ID
* Promotions
* CY.SEND gift card PIN code


 @param ctx context.Context - for authentication, logging, cancellation, deadlines, tracing, etc. Passed from http.Request or context.Background().
 @return ApiGetCategoriesRequest
*/
func (a *StoreAPIService) GetCategories(ctx context.Context) ApiGetCategoriesRequest {
	return ApiGetCategoriesRequest{
		ApiService: a,
		ctx: ctx,
	}
}

// Execute executes the request
//  @return []Category
func (a *StoreAPIService) GetCategoriesExecute(r ApiGetCategoriesRequest) ([]Category, *http.Response, error) {
	var (
		localVarHTTPMethod   = http.MethodGet
		localVarPostBody     interface{}
		formFiles            []formFile
		localVarReturnValue  []Category
	)

	localBasePath, err := a.client.cfg.ServerURLWithContext(r.ctx, "StoreAPIService.GetCategories")
	if err != nil {
		return localVarReturnValue, nil, &GenericOpenAPIError{error: err.Error()}
	}

	localVarPath := localBasePath + "/store/catalogue/categories"

	localVarHeaderParams := make(map[string]string)
	localVarQueryParams := url.Values{}
	localVarFormParams := url.Values{}

	if r.countryId != nil {
		parameterAddToHeaderOrQuery(localVarQueryParams, "country_id", r.countryId, "")
	}
	if r.countryIso2 != nil {
		parameterAddToHeaderOrQuery(localVarQueryParams, "country_iso2", r.countryIso2, "")
	}
	if r.categoryId != nil {
		parameterAddToHeaderOrQuery(localVarQueryParams, "category_id", r.categoryId, "")
	}
	if r.promotion != nil {
		parameterAddToHeaderOrQuery(localVarQueryParams, "promotion", r.promotion, "")
	}
	if r.giftCardCode != nil {
		parameterAddToHeaderOrQuery(localVarQueryParams, "gift_card_code", r.giftCardCode, "")
	}
	// to determine the Content-Type header
	localVarHTTPContentTypes := []string{}

	// set Content-Type header
	localVarHTTPContentType := selectHeaderContentType(localVarHTTPContentTypes)
	if localVarHTTPContentType != "" {
		localVarHeaderParams["Content-Type"] = localVarHTTPContentType
	}

	// to determine the Accept header
	localVarHTTPHeaderAccepts := []string{"application/json"}

	// set Accept header
	localVarHTTPHeaderAccept := selectHeaderAccept(localVarHTTPHeaderAccepts)
	if localVarHTTPHeaderAccept != "" {
		localVarHeaderParams["Accept"] = localVarHTTPHeaderAccept
	}
	if r.ctx != nil {
		// API Key Authentication
		if auth, ok := r.ctx.Value(ContextAPIKeys).(map[string]APIKey); ok {
			if apiKey, ok := auth["Token"]; ok {
				var key string
				if apiKey.Prefix != "" {
					key = apiKey.Prefix + " " + apiKey.Key
				} else {
					key = apiKey.Key
				}
				localVarHeaderParams["Api-token"] = key
			}
		}
	}
	req, err := a.client.prepareRequest(r.ctx, localVarPath, localVarHTTPMethod, localVarPostBody, localVarHeaderParams, localVarQueryParams, localVarFormParams, formFiles)
	if err != nil {
		return localVarReturnValue, nil, err
	}

	localVarHTTPResponse, err := a.client.callAPI(req)
	if err != nil || localVarHTTPResponse == nil {
		return localVarReturnValue, localVarHTTPResponse, err
	}

	localVarBody, err := io.ReadAll(localVarHTTPResponse.Body)
	localVarHTTPResponse.Body.Close()
	localVarHTTPResponse.Body = io.NopCloser(bytes.NewBuffer(localVarBody))
	if err != nil {
		return localVarReturnValue, localVarHTTPResponse, err
	}

	if localVarHTTPResponse.StatusCode >= 300 {
		newErr := &GenericOpenAPIError{
			body:  localVarBody,
			error: localVarHTTPResponse.Status,
		}
		if localVarHTTPResponse.StatusCode == 401 {
			var v APIError
			err = a.client.decode(&v, localVarBody, localVarHTTPResponse.Header.Get("Content-Type"))
			if err != nil {
				newErr.error = err.Error()
				return localVarReturnValue, localVarHTTPResponse, newErr
			}
					newErr.error = formatErrorMessage(localVarHTTPResponse.Status, &v)
					newErr.model = v
			return localVarReturnValue, localVarHTTPResponse, newErr
		}
		if localVarHTTPResponse.StatusCode == 404 {
			var v APIError
			err = a.client.decode(&v, localVarBody, localVarHTTPResponse.Header.Get("Content-Type"))
			if err != nil {
				newErr.error = err.Error()
				return localVarReturnValue, localVarHTTPResponse, newErr
			}
					newErr.error = formatErrorMessage(localVarHTTPResponse.Status, &v)
					newErr.model = v
		}
		return localVarReturnValue, localVarHTTPResponse, newErr
	}

	err = a.client.decode(&localVarReturnValue, localVarBody, localVarHTTPResponse.Header.Get("Content-Type"))
	if err != nil {
		newErr := &GenericOpenAPIError{
			body:  localVarBody,
			error: err.Error(),
		}
		return localVarReturnValue, localVarHTTPResponse, newErr
	}

	return localVarReturnValue, localVarHTTPResponse, nil
}

type ApiGetCountriesRequest struct {
	ctx context.Context
	ApiService *StoreAPIService
	countryId *int32
	countryIso2 *string
	promotion *bool
	giftCardCode *string
}

// Country ID filter
func (r ApiGetCountriesRequest) CountryId(countryId int32) ApiGetCountriesRequest {
	r.countryId = &countryId
	return r
}

// Country ISO 3166-1 alpha-2 filter
func (r ApiGetCountriesRequest) CountryIso2(countryIso2 string) ApiGetCountriesRequest {
	r.countryIso2 = &countryIso2
	return r
}

// Promotion filter
func (r ApiGetCountriesRequest) Promotion(promotion bool) ApiGetCountriesRequest {
	r.promotion = &promotion
	return r
}

// CY.SEND gift card PIN code
func (r ApiGetCountriesRequest) GiftCardCode(giftCardCode string) ApiGetCountriesRequest {
	r.giftCardCode = &giftCardCode
	return r
}

func (r ApiGetCountriesRequest) Execute() ([]Country, *http.Response, error) {
	return r.ApiService.GetCountriesExecute(r)
}

/*
GetCountries Retrieves the CY.SEND product catalogue country list

This operation retrieves the list of countries where CY.SEND products are available.
<br />
The list can be filtered by:
* CY.SEND country ID
* Country ISO 3166-1 alpha-2
* Promotions
* CY.SEND gift card PIN code


 @param ctx context.Context - for authentication, logging, cancellation, deadlines, tracing, etc. Passed from http.Request or context.Background().
 @return ApiGetCountriesRequest
*/
func (a *StoreAPIService) GetCountries(ctx context.Context) ApiGetCountriesRequest {
	return ApiGetCountriesRequest{
		ApiService: a,
		ctx: ctx,
	}
}

// Execute executes the request
//  @return []Country
func (a *StoreAPIService) GetCountriesExecute(r ApiGetCountriesRequest) ([]Country, *http.Response, error) {
	var (
		localVarHTTPMethod   = http.MethodGet
		localVarPostBody     interface{}
		formFiles            []formFile
		localVarReturnValue  []Country
	)

	localBasePath, err := a.client.cfg.ServerURLWithContext(r.ctx, "StoreAPIService.GetCountries")
	if err != nil {
		return localVarReturnValue, nil, &GenericOpenAPIError{error: err.Error()}
	}

	localVarPath := localBasePath + "/store/catalogue/countries"

	localVarHeaderParams := make(map[string]string)
	localVarQueryParams := url.Values{}
	localVarFormParams := url.Values{}

	if r.countryId != nil {
		parameterAddToHeaderOrQuery(localVarQueryParams, "country_id", r.countryId, "")
	}
	if r.countryIso2 != nil {
		parameterAddToHeaderOrQuery(localVarQueryParams, "country_iso2", r.countryIso2, "")
	}
	if r.promotion != nil {
		parameterAddToHeaderOrQuery(localVarQueryParams, "promotion", r.promotion, "")
	}
	if r.giftCardCode != nil {
		parameterAddToHeaderOrQuery(localVarQueryParams, "gift_card_code", r.giftCardCode, "")
	}
	// to determine the Content-Type header
	localVarHTTPContentTypes := []string{}

	// set Content-Type header
	localVarHTTPContentType := selectHeaderContentType(localVarHTTPContentTypes)
	if localVarHTTPContentType != "" {
		localVarHeaderParams["Content-Type"] = localVarHTTPContentType
	}

	// to determine the Accept header
	localVarHTTPHeaderAccepts := []string{"application/json"}

	// set Accept header
	localVarHTTPHeaderAccept := selectHeaderAccept(localVarHTTPHeaderAccepts)
	if localVarHTTPHeaderAccept != "" {
		localVarHeaderParams["Accept"] = localVarHTTPHeaderAccept
	}
	if r.ctx != nil {
		// API Key Authentication
		if auth, ok := r.ctx.Value(ContextAPIKeys).(map[string]APIKey); ok {
			if apiKey, ok := auth["Token"]; ok {
				var key string
				if apiKey.Prefix != "" {
					key = apiKey.Prefix + " " + apiKey.Key
				} else {
					key = apiKey.Key
				}
				localVarHeaderParams["Api-token"] = key
			}
		}
	}
	req, err := a.client.prepareRequest(r.ctx, localVarPath, localVarHTTPMethod, localVarPostBody, localVarHeaderParams, localVarQueryParams, localVarFormParams, formFiles)
	if err != nil {
		return localVarReturnValue, nil, err
	}

	localVarHTTPResponse, err := a.client.callAPI(req)
	if err != nil || localVarHTTPResponse == nil {
		return localVarReturnValue, localVarHTTPResponse, err
	}

	localVarBody, err := io.ReadAll(localVarHTTPResponse.Body)
	localVarHTTPResponse.Body.Close()
	localVarHTTPResponse.Body = io.NopCloser(bytes.NewBuffer(localVarBody))
	if err != nil {
		return localVarReturnValue, localVarHTTPResponse, err
	}

	if localVarHTTPResponse.StatusCode >= 300 {
		newErr := &GenericOpenAPIError{
			body:  localVarBody,
			error: localVarHTTPResponse.Status,
		}
		if localVarHTTPResponse.StatusCode == 401 {
			var v APIError
			err = a.client.decode(&v, localVarBody, localVarHTTPResponse.Header.Get("Content-Type"))
			if err != nil {
				newErr.error = err.Error()
				return localVarReturnValue, localVarHTTPResponse, newErr
			}
					newErr.error = formatErrorMessage(localVarHTTPResponse.Status, &v)
					newErr.model = v
			return localVarReturnValue, localVarHTTPResponse, newErr
		}
		if localVarHTTPResponse.StatusCode == 404 {
			var v APIError
			err = a.client.decode(&v, localVarBody, localVarHTTPResponse.Header.Get("Content-Type"))
			if err != nil {
				newErr.error = err.Error()
				return localVarReturnValue, localVarHTTPResponse, newErr
			}
					newErr.error = formatErrorMessage(localVarHTTPResponse.Status, &v)
					newErr.model = v
		}
		return localVarReturnValue, localVarHTTPResponse, newErr
	}

	err = a.client.decode(&localVarReturnValue, localVarBody, localVarHTTPResponse.Header.Get("Content-Type"))
	if err != nil {
		newErr := &GenericOpenAPIError{
			body:  localVarBody,
			error: err.Error(),
		}
		return localVarReturnValue, localVarHTTPResponse, newErr
	}

	return localVarReturnValue, localVarHTTPResponse, nil
}

type ApiGetFaceValuesRequest struct {
	ctx context.Context
	ApiService *StoreAPIService
	countryId *int32
	countryIso2 *string
	categoryId *int32
	productId *int32
	faceValueId *int32
	promotion *bool
	giftCardCode *string
	mobile *string
}

// Country ID filter
func (r ApiGetFaceValuesRequest) CountryId(countryId int32) ApiGetFaceValuesRequest {
	r.countryId = &countryId
	return r
}

// Country ISO 3166-1 alpha-2 filter
func (r ApiGetFaceValuesRequest) CountryIso2(countryIso2 string) ApiGetFaceValuesRequest {
	r.countryIso2 = &countryIso2
	return r
}

// Category ID filter
func (r ApiGetFaceValuesRequest) CategoryId(categoryId int32) ApiGetFaceValuesRequest {
	r.categoryId = &categoryId
	return r
}

// Product ID filter
func (r ApiGetFaceValuesRequest) ProductId(productId int32) ApiGetFaceValuesRequest {
	r.productId = &productId
	return r
}

// Face value ID filter
func (r ApiGetFaceValuesRequest) FaceValueId(faceValueId int32) ApiGetFaceValuesRequest {
	r.faceValueId = &faceValueId
	return r
}

// Promotion filter
func (r ApiGetFaceValuesRequest) Promotion(promotion bool) ApiGetFaceValuesRequest {
	r.promotion = &promotion
	return r
}

// CY.SEND gift card PIN code
func (r ApiGetFaceValuesRequest) GiftCardCode(giftCardCode string) ApiGetFaceValuesRequest {
	r.giftCardCode = &giftCardCode
	return r
}

// Mobile number you want to recharge. CY.SEND will detect the mobile operator of the number and return it. This service have a cost and will be billed on this request.
func (r ApiGetFaceValuesRequest) Mobile(mobile string) ApiGetFaceValuesRequest {
	r.mobile = &mobile
	return r
}

func (r ApiGetFaceValuesRequest) Execute() (*FaceValue, *http.Response, error) {
	return r.ApiService.GetFaceValuesExecute(r)
}

/*
GetFaceValues Retrieves the CY.SEND product face value list

This operation retrieves the list of the product face values.
<br />
They are two types of face value:
* Fixed: Means a set amount. The beneficiary will receive that exact amount and the CUSTOMER's balance will be debited of the cost.
* Range: Means a set of face values with a minimum, a maximum and a step. The API returns the cost of the minimum and the maximum face value. For example: a range from 1 to 10 step 1 would be: 1, 2, 3, 4 … 8, 9, 10. To calculate the cost of a single face value, since the cost is linear, you must use a rule of three. For example, to calculate the cost of face value 5 you should divide the maximum cost (cost of face value 10) by the maximum face value (10) multiplied by desired face value (5).
---
The list can be filtered by:
* CY.SEND country ID
* Country ISO 3166-1 alpha-2
* CY.SEND category ID
* CY.SEND product ID
* CY.SEND face value ID
* Maintenance
* Promotions
* CY.SEND gift card PIN code
---
NOTE: The face value ID are changed each time the value changes but they are NOT changed when the cost changes. This is valid for ranges and fixed values.


 @param ctx context.Context - for authentication, logging, cancellation, deadlines, tracing, etc. Passed from http.Request or context.Background().
 @return ApiGetFaceValuesRequest
*/
func (a *StoreAPIService) GetFaceValues(ctx context.Context) ApiGetFaceValuesRequest {
	return ApiGetFaceValuesRequest{
		ApiService: a,
		ctx: ctx,
	}
}

// Execute executes the request
//  @return FaceValue
func (a *StoreAPIService) GetFaceValuesExecute(r ApiGetFaceValuesRequest) (*FaceValue, *http.Response, error) {
	var (
		localVarHTTPMethod   = http.MethodGet
		localVarPostBody     interface{}
		formFiles            []formFile
		localVarReturnValue  *FaceValue
	)

	localBasePath, err := a.client.cfg.ServerURLWithContext(r.ctx, "StoreAPIService.GetFaceValues")
	if err != nil {
		return localVarReturnValue, nil, &GenericOpenAPIError{error: err.Error()}
	}

	localVarPath := localBasePath + "/store/catalogue/face-value"

	localVarHeaderParams := make(map[string]string)
	localVarQueryParams := url.Values{}
	localVarFormParams := url.Values{}

	if r.countryId != nil {
		parameterAddToHeaderOrQuery(localVarQueryParams, "country_id", r.countryId, "")
	}
	if r.countryIso2 != nil {
		parameterAddToHeaderOrQuery(localVarQueryParams, "country_iso2", r.countryIso2, "")
	}
	if r.categoryId != nil {
		parameterAddToHeaderOrQuery(localVarQueryParams, "category_id", r.categoryId, "")
	}
	if r.productId != nil {
		parameterAddToHeaderOrQuery(localVarQueryParams, "product_id", r.productId, "")
	}
	if r.faceValueId != nil {
		parameterAddToHeaderOrQuery(localVarQueryParams, "face_value_id", r.faceValueId, "")
	}
	if r.promotion != nil {
		parameterAddToHeaderOrQuery(localVarQueryParams, "promotion", r.promotion, "")
	}
	if r.giftCardCode != nil {
		parameterAddToHeaderOrQuery(localVarQueryParams, "gift_card_code", r.giftCardCode, "")
	}
	if r.mobile != nil {
		parameterAddToHeaderOrQuery(localVarQueryParams, "mobile", r.mobile, "")
	}
	// to determine the Content-Type header
	localVarHTTPContentTypes := []string{}

	// set Content-Type header
	localVarHTTPContentType := selectHeaderContentType(localVarHTTPContentTypes)
	if localVarHTTPContentType != "" {
		localVarHeaderParams["Content-Type"] = localVarHTTPContentType
	}

	// to determine the Accept header
	localVarHTTPHeaderAccepts := []string{"application/json"}

	// set Accept header
	localVarHTTPHeaderAccept := selectHeaderAccept(localVarHTTPHeaderAccepts)
	if localVarHTTPHeaderAccept != "" {
		localVarHeaderParams["Accept"] = localVarHTTPHeaderAccept
	}
	if r.ctx != nil {
		// API Key Authentication
		if auth, ok := r.ctx.Value(ContextAPIKeys).(map[string]APIKey); ok {
			if apiKey, ok := auth["Token"]; ok {
				var key string
				if apiKey.Prefix != "" {
					key = apiKey.Prefix + " " + apiKey.Key
				} else {
					key = apiKey.Key
				}
				localVarHeaderParams["Api-token"] = key
			}
		}
	}
	req, err := a.client.prepareRequest(r.ctx, localVarPath, localVarHTTPMethod, localVarPostBody, localVarHeaderParams, localVarQueryParams, localVarFormParams, formFiles)
	if err != nil {
		return localVarReturnValue, nil, err
	}

	localVarHTTPResponse, err := a.client.callAPI(req)
	if err != nil || localVarHTTPResponse == nil {
		return localVarReturnValue, localVarHTTPResponse, err
	}

	localVarBody, err := io.ReadAll(localVarHTTPResponse.Body)
	localVarHTTPResponse.Body.Close()
	localVarHTTPResponse.Body = io.NopCloser(bytes.NewBuffer(localVarBody))
	if err != nil {
		return localVarReturnValue, localVarHTTPResponse, err
	}

	if localVarHTTPResponse.StatusCode >= 300 {
		newErr := &GenericOpenAPIError{
			body:  localVarBody,
			error: localVarHTTPResponse.Status,
		}
		if localVarHTTPResponse.StatusCode == 401 {
			var v APIError
			err = a.client.decode(&v, localVarBody, localVarHTTPResponse.Header.Get("Content-Type"))
			if err != nil {
				newErr.error = err.Error()
				return localVarReturnValue, localVarHTTPResponse, newErr
			}
					newErr.error = formatErrorMessage(localVarHTTPResponse.Status, &v)
					newErr.model = v
			return localVarReturnValue, localVarHTTPResponse, newErr
		}
		if localVarHTTPResponse.StatusCode == 404 {
			var v APIError
			err = a.client.decode(&v, localVarBody, localVarHTTPResponse.Header.Get("Content-Type"))
			if err != nil {
				newErr.error = err.Error()
				return localVarReturnValue, localVarHTTPResponse, newErr
			}
					newErr.error = formatErrorMessage(localVarHTTPResponse.Status, &v)
					newErr.model = v
		}
		return localVarReturnValue, localVarHTTPResponse, newErr
	}

	err = a.client.decode(&localVarReturnValue, localVarBody, localVarHTTPResponse.Header.Get("Content-Type"))
	if err != nil {
		newErr := &GenericOpenAPIError{
			body:  localVarBody,
			error: err.Error(),
		}
		return localVarReturnValue, localVarHTTPResponse, newErr
	}

	return localVarReturnValue, localVarHTTPResponse, nil
}

type ApiGetMaintenanceRequest struct {
	ctx context.Context
	ApiService *StoreAPIService
	productId *int32
}

// Product ID filter
func (r ApiGetMaintenanceRequest) ProductId(productId int32) ApiGetMaintenanceRequest {
	r.productId = &productId
	return r
}

func (r ApiGetMaintenanceRequest) Execute() ([]Maintenance, *http.Response, error) {
	return r.ApiService.GetMaintenanceExecute(r)
}

/*
GetMaintenance Retrieves current and future product maintenances

This operation retrieves the list of current and future product maintenances.
<br />
When a product is under maintenance, it is not available for purchase.
<br />
The list can be filtered by:
* CY.SEND product ID


 @param ctx context.Context - for authentication, logging, cancellation, deadlines, tracing, etc. Passed from http.Request or context.Background().
 @return ApiGetMaintenanceRequest
*/
func (a *StoreAPIService) GetMaintenance(ctx context.Context) ApiGetMaintenanceRequest {
	return ApiGetMaintenanceRequest{
		ApiService: a,
		ctx: ctx,
	}
}

// Execute executes the request
//  @return []Maintenance
func (a *StoreAPIService) GetMaintenanceExecute(r ApiGetMaintenanceRequest) ([]Maintenance, *http.Response, error) {
	var (
		localVarHTTPMethod   = http.MethodGet
		localVarPostBody     interface{}
		formFiles            []formFile
		localVarReturnValue  []Maintenance
	)

	localBasePath, err := a.client.cfg.ServerURLWithContext(r.ctx, "StoreAPIService.GetMaintenance")
	if err != nil {
		return localVarReturnValue, nil, &GenericOpenAPIError{error: err.Error()}
	}

	localVarPath := localBasePath + "/store/catalogue/maintenance"

	localVarHeaderParams := make(map[string]string)
	localVarQueryParams := url.Values{}
	localVarFormParams := url.Values{}

	if r.productId != nil {
		parameterAddToHeaderOrQuery(localVarQueryParams, "product_id", r.productId, "")
	}
	// to determine the Content-Type header
	localVarHTTPContentTypes := []string{}

	// set Content-Type header
	localVarHTTPContentType := selectHeaderContentType(localVarHTTPContentTypes)
	if localVarHTTPContentType != "" {
		localVarHeaderParams["Content-Type"] = localVarHTTPContentType
	}

	// to determine the Accept header
	localVarHTTPHeaderAccepts := []string{"application/json"}

	// set Accept header
	localVarHTTPHeaderAccept := selectHeaderAccept(localVarHTTPHeaderAccepts)
	if localVarHTTPHeaderAccept != "" {
		localVarHeaderParams["Accept"] = localVarHTTPHeaderAccept
	}
	if r.ctx != nil {
		// API Key Authentication
		if auth, ok := r.ctx.Value(ContextAPIKeys).(map[string]APIKey); ok {
			if apiKey, ok := auth["Token"]; ok {
				var key string
				if apiKey.Prefix != "" {
					key = apiKey.Prefix + " " + apiKey.Key
				} else {
					key = apiKey.Key
				}
				localVarHeaderParams["Api-token"] = key
			}
		}
	}
	req, err := a.client.prepareRequest(r.ctx, localVarPath, localVarHTTPMethod, localVarPostBody, localVarHeaderParams, localVarQueryParams, localVarFormParams, formFiles)
	if err != nil {
		return localVarReturnValue, nil, err
	}

	localVarHTTPResponse, err := a.client.callAPI(req)
	if err != nil || localVarHTTPResponse == nil {
		return localVarReturnValue, localVarHTTPResponse, err
	}

	localVarBody, err := io.ReadAll(localVarHTTPResponse.Body)
	localVarHTTPResponse.Body.Close()
	localVarHTTPResponse.Body = io.NopCloser(bytes.NewBuffer(localVarBody))
	if err != nil {
		return localVarReturnValue, localVarHTTPResponse, err
	}

	if localVarHTTPResponse.StatusCode >= 300 {
		newErr := &GenericOpenAPIError{
			body:  localVarBody,
			error: localVarHTTPResponse.Status,
		}
		if localVarHTTPResponse.StatusCode == 401 {
			var v APIError
			err = a.client.decode(&v, localVarBody, localVarHTTPResponse.Header.Get("Content-Type"))
			if err != nil {
				newErr.error = err.Error()
				return localVarReturnValue, localVarHTTPResponse, newErr
			}
					newErr.error = formatErrorMessage(localVarHTTPResponse.Status, &v)
					newErr.model = v
			return localVarReturnValue, localVarHTTPResponse, newErr
		}
		if localVarHTTPResponse.StatusCode == 404 {
			var v APIError
			err = a.client.decode(&v, localVarBody, localVarHTTPResponse.Header.Get("Content-Type"))
			if err != nil {
				newErr.error = err.Error()
				return localVarReturnValue, localVarHTTPResponse, newErr
			}
					newErr.error = formatErrorMessage(localVarHTTPResponse.Status, &v)
					newErr.model = v
		}
		return localVarReturnValue, localVarHTTPResponse, newErr
	}

	err = a.client.decode(&localVarReturnValue, localVarBody, localVarHTTPResponse.Header.Get("Content-Type"))
	if err != nil {
		newErr := &GenericOpenAPIError{
			body:  localVarBody,
			error: err.Error(),
		}
		return localVarReturnValue, localVarHTTPResponse, newErr
	}

	return localVarReturnValue, localVarHTTPResponse, nil
}

type ApiGetOrderRequest struct {
	ctx context.Context
	ApiService *StoreAPIService
	uid *string
	userUid *string
}

// CY.SEND order unique ID filter
func (r ApiGetOrderRequest) Uid(uid string) ApiGetOrderRequest {
	r.uid = &uid
	return r
}

// CUSTOMER unique order ID filter
func (r ApiGetOrderRequest) UserUid(userUid string) ApiGetOrderRequest {
	r.userUid = &userUid
	return r
}

func (r ApiGetOrderRequest) Execute() ([]OrderStatus, *http.Response, error) {
	return r.ApiService.GetOrderExecute(r)
}

/*
GetOrder (Optional) Retrieves order information by CY.SEND order unique ID or by CUSTOMER order unique ID

This optional operation retrieves past purchase order information if you have not received it using the callback URL.
<br />
It is strongly recommended to set a callback (callback_url parameter) for each order to automatically receive a
notification on the URL of your choice when an order changes status even if it comes after a long wait.
<br />
The `status` of an order can be:
* accepted: The purchase is still in progress. This means that we are waiting for a confirmation from the product
issuer or that CY.SEND delayed the transaction for security reasons.
* success: The purchase was executed successfully.
* failed: The purchase failed. Additional information is returned in `response_code`, `response_title` and
`response_description` parameters.
---
For a successful top up, some product issuers provide a purchase reference number (`issuer_reference` parameter)
that can be used by the beneficiary for his customer care enquiries when contacting the product issuer.
<br />
<br />
If the purchase order contains a prepaid code product, the response will contain the prepaid code, the prepaid
code serial number, the prepaid code expiration and the prepaid code usage instructions (`prepaid_code` parameter).
<br />
<br />
Response codes (`response_code` parameter) are based on [ISO8583](https://en.wikipedia.org/wiki/ISO_8583#Response_code) codes.
<br />
<br />
Using a cron job to periodically check if there is an order status change is a waste of resources on both
sides. If you do not have a choice you can use the Retrieve a order operation.


 @param ctx context.Context - for authentication, logging, cancellation, deadlines, tracing, etc. Passed from http.Request or context.Background().
 @return ApiGetOrderRequest
*/
func (a *StoreAPIService) GetOrder(ctx context.Context) ApiGetOrderRequest {
	return ApiGetOrderRequest{
		ApiService: a,
		ctx: ctx,
	}
}

// Execute executes the request
//  @return []OrderStatus
func (a *StoreAPIService) GetOrderExecute(r ApiGetOrderRequest) ([]OrderStatus, *http.Response, error) {
	var (
		localVarHTTPMethod   = http.MethodGet
		localVarPostBody     interface{}
		formFiles            []formFile
		localVarReturnValue  []OrderStatus
	)

	localBasePath, err := a.client.cfg.ServerURLWithContext(r.ctx, "StoreAPIService.GetOrder")
	if err != nil {
		return localVarReturnValue, nil, &GenericOpenAPIError{error: err.Error()}
	}

	localVarPath := localBasePath + "/store/order"

	localVarHeaderParams := make(map[string]string)
	localVarQueryParams := url.Values{}
	localVarFormParams := url.Values{}

	if r.uid != nil {
		parameterAddToHeaderOrQuery(localVarQueryParams, "uid", r.uid, "")
	}
	if r.userUid != nil {
		parameterAddToHeaderOrQuery(localVarQueryParams, "user_uid", r.userUid, "")
	}
	// to determine the Content-Type header
	localVarHTTPContentTypes := []string{}

	// set Content-Type header
	localVarHTTPContentType := selectHeaderContentType(localVarHTTPContentTypes)
	if localVarHTTPContentType != "" {
		localVarHeaderParams["Content-Type"] = localVarHTTPContentType
	}

	// to determine the Accept header
	localVarHTTPHeaderAccepts := []string{"application/json"}

	// set Accept header
	localVarHTTPHeaderAccept := selectHeaderAccept(localVarHTTPHeaderAccepts)
	if localVarHTTPHeaderAccept != "" {
		localVarHeaderParams["Accept"] = localVarHTTPHeaderAccept
	}
	if r.ctx != nil {
		// API Key Authentication
		if auth, ok := r.ctx.Value(ContextAPIKeys).(map[string]APIKey); ok {
			if apiKey, ok := auth["Token"]; ok {
				var key string
				if apiKey.Prefix != "" {
					key = apiKey.Prefix + " " + apiKey.Key
				} else {
					key = apiKey.Key
				}
				localVarHeaderParams["Api-token"] = key
			}
		}
	}
	req, err := a.client.prepareRequest(r.ctx, localVarPath, localVarHTTPMethod, localVarPostBody, localVarHeaderParams, localVarQueryParams, localVarFormParams, formFiles)
	if err != nil {
		return localVarReturnValue, nil, err
	}

	localVarHTTPResponse, err := a.client.callAPI(req)
	if err != nil || localVarHTTPResponse == nil {
		return localVarReturnValue, localVarHTTPResponse, err
	}

	localVarBody, err := io.ReadAll(localVarHTTPResponse.Body)
	localVarHTTPResponse.Body.Close()
	localVarHTTPResponse.Body = io.NopCloser(bytes.NewBuffer(localVarBody))
	if err != nil {
		return localVarReturnValue, localVarHTTPResponse, err
	}

	if localVarHTTPResponse.StatusCode >= 300 {
		newErr := &GenericOpenAPIError{
			body:  localVarBody,
			error: localVarHTTPResponse.Status,
		}
		if localVarHTTPResponse.StatusCode == 401 {
			var v APIError
			err = a.client.decode(&v, localVarBody, localVarHTTPResponse.Header.Get("Content-Type"))
			if err != nil {
				newErr.error = err.Error()
				return localVarReturnValue, localVarHTTPResponse, newErr
			}
					newErr.error = formatErrorMessage(localVarHTTPResponse.Status, &v)
					newErr.model = v
			return localVarReturnValue, localVarHTTPResponse, newErr
		}
		if localVarHTTPResponse.StatusCode == 404 {
			var v APIError
			err = a.client.decode(&v, localVarBody, localVarHTTPResponse.Header.Get("Content-Type"))
			if err != nil {
				newErr.error = err.Error()
				return localVarReturnValue, localVarHTTPResponse, newErr
			}
					newErr.error = formatErrorMessage(localVarHTTPResponse.Status, &v)
					newErr.model = v
		}
		return localVarReturnValue, localVarHTTPResponse, newErr
	}

	err = a.client.decode(&localVarReturnValue, localVarBody, localVarHTTPResponse.Header.Get("Content-Type"))
	if err != nil {
		newErr := &GenericOpenAPIError{
			body:  localVarBody,
			error: err.Error(),
		}
		return localVarReturnValue, localVarHTTPResponse, newErr
	}

	return localVarReturnValue, localVarHTTPResponse, nil
}

type ApiGetProductCostRequest struct {
	ctx context.Context
	ApiService *StoreAPIService
	costRequest *[]CostRequest
}

// Order you want to calculate the cost
func (r ApiGetProductCostRequest) CostRequest(costRequest []CostRequest) ApiGetProductCostRequest {
	r.costRequest = &costRequest
	return r
}

func (r ApiGetProductCostRequest) Execute() ([]CostResponse, *http.Response, error) {
	return r.ApiService.GetProductCostExecute(r)
}

/*
GetProductCost Retrieve the cost of a face value

This operation returns an indicative cost of a face value. It does not place a purchase order. It is only a calculation tool to know the cost of a face value before placing the final purchase order. It returns the current cost at the exact time of the operation.


 @param ctx context.Context - for authentication, logging, cancellation, deadlines, tracing, etc. Passed from http.Request or context.Background().
 @return ApiGetProductCostRequest
*/
func (a *StoreAPIService) GetProductCost(ctx context.Context) ApiGetProductCostRequest {
	return ApiGetProductCostRequest{
		ApiService: a,
		ctx: ctx,
	}
}

// Execute executes the request
//  @return []CostResponse
func (a *StoreAPIService) GetProductCostExecute(r ApiGetProductCostRequest) ([]CostResponse, *http.Response, error) {
	var (
		localVarHTTPMethod   = http.MethodPut
		localVarPostBody     interface{}
		formFiles            []formFile
		localVarReturnValue  []CostResponse
	)

	localBasePath, err := a.client.cfg.ServerURLWithContext(r.ctx, "StoreAPIService.GetProductCost")
	if err != nil {
		return localVarReturnValue, nil, &GenericOpenAPIError{error: err.Error()}
	}

	localVarPath := localBasePath + "/store/order/cost"

	localVarHeaderParams := make(map[string]string)
	localVarQueryParams := url.Values{}
	localVarFormParams := url.Values{}
	if r.costRequest == nil {
		return localVarReturnValue, nil, reportError("costRequest is required and must be specified")
	}

	// to determine the Content-Type header
	localVarHTTPContentTypes := []string{"application/json"}

	// set Content-Type header
	localVarHTTPContentType := selectHeaderContentType(localVarHTTPContentTypes)
	if localVarHTTPContentType != "" {
		localVarHeaderParams["Content-Type"] = localVarHTTPContentType
	}

	// to determine the Accept header
	localVarHTTPHeaderAccepts := []string{"application/json"}

	// set Accept header
	localVarHTTPHeaderAccept := selectHeaderAccept(localVarHTTPHeaderAccepts)
	if localVarHTTPHeaderAccept != "" {
		localVarHeaderParams["Accept"] = localVarHTTPHeaderAccept
	}
	// body params
	localVarPostBody = r.costRequest
	if r.ctx != nil {
		// API Key Authentication
		if auth, ok := r.ctx.Value(ContextAPIKeys).(map[string]APIKey); ok {
			if apiKey, ok := auth["Token"]; ok {
				var key string
				if apiKey.Prefix != "" {
					key = apiKey.Prefix + " " + apiKey.Key
				} else {
					key = apiKey.Key
				}
				localVarHeaderParams["Api-token"] = key
			}
		}
	}
	req, err := a.client.prepareRequest(r.ctx, localVarPath, localVarHTTPMethod, localVarPostBody, localVarHeaderParams, localVarQueryParams, localVarFormParams, formFiles)
	if err != nil {
		return localVarReturnValue, nil, err
	}

	localVarHTTPResponse, err := a.client.callAPI(req)
	if err != nil || localVarHTTPResponse == nil {
		return localVarReturnValue, localVarHTTPResponse, err
	}

	localVarBody, err := io.ReadAll(localVarHTTPResponse.Body)
	localVarHTTPResponse.Body.Close()
	localVarHTTPResponse.Body = io.NopCloser(bytes.NewBuffer(localVarBody))
	if err != nil {
		return localVarReturnValue, localVarHTTPResponse, err
	}

	if localVarHTTPResponse.StatusCode >= 300 {
		newErr := &GenericOpenAPIError{
			body:  localVarBody,
			error: localVarHTTPResponse.Status,
		}
		return localVarReturnValue, localVarHTTPResponse, newErr
	}

	err = a.client.decode(&localVarReturnValue, localVarBody, localVarHTTPResponse.Header.Get("Content-Type"))
	if err != nil {
		newErr := &GenericOpenAPIError{
			body:  localVarBody,
			error: err.Error(),
		}
		return localVarReturnValue, localVarHTTPResponse, newErr
	}

	return localVarReturnValue, localVarHTTPResponse, nil
}

type ApiGetProductsRequest struct {
	ctx context.Context
	ApiService *StoreAPIService
	countryId *int32
	countryIso2 *string
	categoryId *int32
	productId *int32
	maintenance *bool
	promotion *bool
	giftCardCode *string
	mobile *string
	returnUsageInstruction *bool
	languages *string
}

// Country ID filter
func (r ApiGetProductsRequest) CountryId(countryId int32) ApiGetProductsRequest {
	r.countryId = &countryId
	return r
}

// Country ISO 3166-1 alpha-2 filter
func (r ApiGetProductsRequest) CountryIso2(countryIso2 string) ApiGetProductsRequest {
	r.countryIso2 = &countryIso2
	return r
}

// Category ID filter
func (r ApiGetProductsRequest) CategoryId(categoryId int32) ApiGetProductsRequest {
	r.categoryId = &categoryId
	return r
}

// Product ID filter
func (r ApiGetProductsRequest) ProductId(productId int32) ApiGetProductsRequest {
	r.productId = &productId
	return r
}

// Maintenance filter
func (r ApiGetProductsRequest) Maintenance(maintenance bool) ApiGetProductsRequest {
	r.maintenance = &maintenance
	return r
}

// Promotion filter
func (r ApiGetProductsRequest) Promotion(promotion bool) ApiGetProductsRequest {
	r.promotion = &promotion
	return r
}

// CY.SEND gift card PIN code
func (r ApiGetProductsRequest) GiftCardCode(giftCardCode string) ApiGetProductsRequest {
	r.giftCardCode = &giftCardCode
	return r
}

// This mobile lookup filter will attempt to detect the mobile operator (check local regulations to be sure it is fully supported) and return the corresponding product. This operation has a cost of 0.01 EUR per query. Fund your account before you use this operation. [Load your balance.](https://www.cysend.com/en/add-funds)
func (r ApiGetProductsRequest) Mobile(mobile string) ApiGetProductsRequest {
	r.mobile = &mobile
	return r
}

// Set to true to receive PIN products usage instructions.
func (r ApiGetProductsRequest) ReturnUsageInstruction(returnUsageInstruction bool) ApiGetProductsRequest {
	r.returnUsageInstruction = &returnUsageInstruction
	return r
}

// List of language code in ISO 639-1 format separated by a coma.
func (r ApiGetProductsRequest) Languages(languages string) ApiGetProductsRequest {
	r.languages = &languages
	return r
}

func (r ApiGetProductsRequest) Execute() ([]Product, *http.Response, error) {
	return r.ApiService.GetProductsExecute(r)
}

/*
GetProducts Retrieves the CY.SEND product catalogue list

This operation retrieves the list of the CY.SEND products.
<br />
A product can belong to multiple countries and categories.
<br />
Each product requires specific beneficiary information to place a purchase order. Some will require only a mobile
number (the Mobile filter will attempt to detect the mobile operator, some countries do not support this feature,
check the local regulations for more details), some an account number, and others won’t require anything.
<br />

This list can be filtered by:
* CY.SEND country ID
* Country ISO 3166-1 alpha-2
* CY.SEND category ID
* CY.SEND product ID
* Maintenance
* Promotions
* CY.SEND gift card PIN code
* Mobile


 @param ctx context.Context - for authentication, logging, cancellation, deadlines, tracing, etc. Passed from http.Request or context.Background().
 @return ApiGetProductsRequest
*/
func (a *StoreAPIService) GetProducts(ctx context.Context) ApiGetProductsRequest {
	return ApiGetProductsRequest{
		ApiService: a,
		ctx: ctx,
	}
}

// Execute executes the request
//  @return []Product
func (a *StoreAPIService) GetProductsExecute(r ApiGetProductsRequest) ([]Product, *http.Response, error) {
	var (
		localVarHTTPMethod   = http.MethodGet
		localVarPostBody     interface{}
		formFiles            []formFile
		localVarReturnValue  []Product
	)

	localBasePath, err := a.client.cfg.ServerURLWithContext(r.ctx, "StoreAPIService.GetProducts")
	if err != nil {
		return localVarReturnValue, nil, &GenericOpenAPIError{error: err.Error()}
	}

	localVarPath := localBasePath + "/store/catalogue/products"

	localVarHeaderParams := make(map[string]string)
	localVarQueryParams := url.Values{}
	localVarFormParams := url.Values{}

	if r.countryId != nil {
		parameterAddToHeaderOrQuery(localVarQueryParams, "country_id", r.countryId, "")
	}
	if r.countryIso2 != nil {
		parameterAddToHeaderOrQuery(localVarQueryParams, "country_iso2", r.countryIso2, "")
	}
	if r.categoryId != nil {
		parameterAddToHeaderOrQuery(localVarQueryParams, "category_id", r.categoryId, "")
	}
	if r.productId != nil {
		parameterAddToHeaderOrQuery(localVarQueryParams, "product_id", r.productId, "")
	}
	if r.maintenance != nil {
		parameterAddToHeaderOrQuery(localVarQueryParams, "maintenance", r.maintenance, "")
	}
	if r.promotion != nil {
		parameterAddToHeaderOrQuery(localVarQueryParams, "promotion", r.promotion, "")
	}
	if r.giftCardCode != nil {
		parameterAddToHeaderOrQuery(localVarQueryParams, "gift_card_code", r.giftCardCode, "")
	}
	if r.mobile != nil {
		parameterAddToHeaderOrQuery(localVarQueryParams, "mobile", r.mobile, "")
	}
	if r.returnUsageInstruction != nil {
		parameterAddToHeaderOrQuery(localVarQueryParams, "return_usage_instruction", r.returnUsageInstruction, "")
	} else {
		var defaultValue bool = false
		r.returnUsageInstruction = &defaultValue
	}
	if r.languages != nil {
		parameterAddToHeaderOrQuery(localVarQueryParams, "languages", r.languages, "")
	}
	// to determine the Content-Type header
	localVarHTTPContentTypes := []string{}

	// set Content-Type header
	localVarHTTPContentType := selectHeaderContentType(localVarHTTPContentTypes)
	if localVarHTTPContentType != "" {
		localVarHeaderParams["Content-Type"] = localVarHTTPContentType
	}

	// to determine the Accept header
	localVarHTTPHeaderAccepts := []string{"application/json"}

	// set Accept header
	localVarHTTPHeaderAccept := selectHeaderAccept(localVarHTTPHeaderAccepts)
	if localVarHTTPHeaderAccept != "" {
		localVarHeaderParams["Accept"] = localVarHTTPHeaderAccept
	}
	if r.ctx != nil {
		// API Key Authentication
		if auth, ok := r.ctx.Value(ContextAPIKeys).(map[string]APIKey); ok {
			if apiKey, ok := auth["Token"]; ok {
				var key string
				if apiKey.Prefix != "" {
					key = apiKey.Prefix + " " + apiKey.Key
				} else {
					key = apiKey.Key
				}
				localVarHeaderParams["Api-token"] = key
			}
		}
	}
	req, err := a.client.prepareRequest(r.ctx, localVarPath, localVarHTTPMethod, localVarPostBody, localVarHeaderParams, localVarQueryParams, localVarFormParams, formFiles)
	if err != nil {
		return localVarReturnValue, nil, err
	}

	localVarHTTPResponse, err := a.client.callAPI(req)
	if err != nil || localVarHTTPResponse == nil {
		return localVarReturnValue, localVarHTTPResponse, err
	}

	localVarBody, err := io.ReadAll(localVarHTTPResponse.Body)
	localVarHTTPResponse.Body.Close()
	localVarHTTPResponse.Body = io.NopCloser(bytes.NewBuffer(localVarBody))
	if err != nil {
		return localVarReturnValue, localVarHTTPResponse, err
	}

	if localVarHTTPResponse.StatusCode >= 300 {
		newErr := &GenericOpenAPIError{
			body:  localVarBody,
			error: localVarHTTPResponse.Status,
		}
		if localVarHTTPResponse.StatusCode == 401 {
			var v APIError
			err = a.client.decode(&v, localVarBody, localVarHTTPResponse.Header.Get("Content-Type"))
			if err != nil {
				newErr.error = err.Error()
				return localVarReturnValue, localVarHTTPResponse, newErr
			}
					newErr.error = formatErrorMessage(localVarHTTPResponse.Status, &v)
					newErr.model = v
			return localVarReturnValue, localVarHTTPResponse, newErr
		}
		if localVarHTTPResponse.StatusCode == 404 {
			var v APIError
			err = a.client.decode(&v, localVarBody, localVarHTTPResponse.Header.Get("Content-Type"))
			if err != nil {
				newErr.error = err.Error()
				return localVarReturnValue, localVarHTTPResponse, newErr
			}
					newErr.error = formatErrorMessage(localVarHTTPResponse.Status, &v)
					newErr.model = v
		}
		return localVarReturnValue, localVarHTTPResponse, newErr
	}

	err = a.client.decode(&localVarReturnValue, localVarBody, localVarHTTPResponse.Header.Get("Content-Type"))
	if err != nil {
		newErr := &GenericOpenAPIError{
			body:  localVarBody,
			error: err.Error(),
		}
		return localVarReturnValue, localVarHTTPResponse, newErr
	}

	return localVarReturnValue, localVarHTTPResponse, nil
}

type ApiGetPromotionsRequest struct {
	ctx context.Context
	ApiService *StoreAPIService
	productId *int32
}

// Product ID filter
func (r ApiGetPromotionsRequest) ProductId(productId int32) ApiGetPromotionsRequest {
	r.productId = &productId
	return r
}

func (r ApiGetPromotionsRequest) Execute() ([]Promotion, *http.Response, error) {
	return r.ApiService.GetPromotionsExecute(r)
}

/*
GetPromotions Retrieves current and future product promotions

This operation retrieves the list of current and future product promotions.
<br />
The list can be filtered by:
* CY.SEND product ID


 @param ctx context.Context - for authentication, logging, cancellation, deadlines, tracing, etc. Passed from http.Request or context.Background().
 @return ApiGetPromotionsRequest
*/
func (a *StoreAPIService) GetPromotions(ctx context.Context) ApiGetPromotionsRequest {
	return ApiGetPromotionsRequest{
		ApiService: a,
		ctx: ctx,
	}
}

// Execute executes the request
//  @return []Promotion
func (a *StoreAPIService) GetPromotionsExecute(r ApiGetPromotionsRequest) ([]Promotion, *http.Response, error) {
	var (
		localVarHTTPMethod   = http.MethodGet
		localVarPostBody     interface{}
		formFiles            []formFile
		localVarReturnValue  []Promotion
	)

	localBasePath, err := a.client.cfg.ServerURLWithContext(r.ctx, "StoreAPIService.GetPromotions")
	if err != nil {
		return localVarReturnValue, nil, &GenericOpenAPIError{error: err.Error()}
	}

	localVarPath := localBasePath + "/store/catalogue/promotion"

	localVarHeaderParams := make(map[string]string)
	localVarQueryParams := url.Values{}
	localVarFormParams := url.Values{}

	if r.productId != nil {
		parameterAddToHeaderOrQuery(localVarQueryParams, "product_id", r.productId, "")
	}
	// to determine the Content-Type header
	localVarHTTPContentTypes := []string{}

	// set Content-Type header
	localVarHTTPContentType := selectHeaderContentType(localVarHTTPContentTypes)
	if localVarHTTPContentType != "" {
		localVarHeaderParams["Content-Type"] = localVarHTTPContentType
	}

	// to determine the Accept header
	localVarHTTPHeaderAccepts := []string{"application/json"}

	// set Accept header
	localVarHTTPHeaderAccept := selectHeaderAccept(localVarHTTPHeaderAccepts)
	if localVarHTTPHeaderAccept != "" {
		localVarHeaderParams["Accept"] = localVarHTTPHeaderAccept
	}
	if r.ctx != nil {
		// API Key Authentication
		if auth, ok := r.ctx.Value(ContextAPIKeys).(map[string]APIKey); ok {
			if apiKey, ok := auth["Token"]; ok {
				var key string
				if apiKey.Prefix != "" {
					key = apiKey.Prefix + " " + apiKey.Key
				} else {
					key = apiKey.Key
				}
				localVarHeaderParams["Api-token"] = key
			}
		}
	}
	req, err := a.client.prepareRequest(r.ctx, localVarPath, localVarHTTPMethod, localVarPostBody, localVarHeaderParams, localVarQueryParams, localVarFormParams, formFiles)
	if err != nil {
		return localVarReturnValue, nil, err
	}

	localVarHTTPResponse, err := a.client.callAPI(req)
	if err != nil || localVarHTTPResponse == nil {
		return localVarReturnValue, localVarHTTPResponse, err
	}

	localVarBody, err := io.ReadAll(localVarHTTPResponse.Body)
	localVarHTTPResponse.Body.Close()
	localVarHTTPResponse.Body = io.NopCloser(bytes.NewBuffer(localVarBody))
	if err != nil {
		return localVarReturnValue, localVarHTTPResponse, err
	}

	if localVarHTTPResponse.StatusCode >= 300 {
		newErr := &GenericOpenAPIError{
			body:  localVarBody,
			error: localVarHTTPResponse.Status,
		}
		if localVarHTTPResponse.StatusCode == 401 {
			var v APIError
			err = a.client.decode(&v, localVarBody, localVarHTTPResponse.Header.Get("Content-Type"))
			if err != nil {
				newErr.error = err.Error()
				return localVarReturnValue, localVarHTTPResponse, newErr
			}
					newErr.error = formatErrorMessage(localVarHTTPResponse.Status, &v)
					newErr.model = v
			return localVarReturnValue, localVarHTTPResponse, newErr
		}
		if localVarHTTPResponse.StatusCode == 404 {
			var v APIError
			err = a.client.decode(&v, localVarBody, localVarHTTPResponse.Header.Get("Content-Type"))
			if err != nil {
				newErr.error = err.Error()
				return localVarReturnValue, localVarHTTPResponse, newErr
			}
					newErr.error = formatErrorMessage(localVarHTTPResponse.Status, &v)
					newErr.model = v
		}
		return localVarReturnValue, localVarHTTPResponse, newErr
	}

	err = a.client.decode(&localVarReturnValue, localVarBody, localVarHTTPResponse.Header.Get("Content-Type"))
	if err != nil {
		newErr := &GenericOpenAPIError{
			body:  localVarBody,
			error: err.Error(),
		}
		return localVarReturnValue, localVarHTTPResponse, newErr
	}

	return localVarReturnValue, localVarHTTPResponse, nil
}

type ApiGetResponsesRequest struct {
	ctx context.Context
	ApiService *StoreAPIService
}

func (r ApiGetResponsesRequest) Execute() ([]OrderResponseCode, *http.Response, error) {
	return r.ApiService.GetResponsesExecute(r)
}

/*
GetResponses Retrieve the list of possible order response codes

This operation retrieves the list of all possible purchase order response codes.
<br />
Response codes are based on [ISO8583](https://en.wikipedia.org/wiki/ISO_8583#Response_code) codes.


 @param ctx context.Context - for authentication, logging, cancellation, deadlines, tracing, etc. Passed from http.Request or context.Background().
 @return ApiGetResponsesRequest
*/
func (a *StoreAPIService) GetResponses(ctx context.Context) ApiGetResponsesRequest {
	return ApiGetResponsesRequest{
		ApiService: a,
		ctx: ctx,
	}
}

// Execute executes the request
//  @return []OrderResponseCode
func (a *StoreAPIService) GetResponsesExecute(r ApiGetResponsesRequest) ([]OrderResponseCode, *http.Response, error) {
	var (
		localVarHTTPMethod   = http.MethodGet
		localVarPostBody     interface{}
		formFiles            []formFile
		localVarReturnValue  []OrderResponseCode
	)

	localBasePath, err := a.client.cfg.ServerURLWithContext(r.ctx, "StoreAPIService.GetResponses")
	if err != nil {
		return localVarReturnValue, nil, &GenericOpenAPIError{error: err.Error()}
	}

	localVarPath := localBasePath + "/store/order/responses"

	localVarHeaderParams := make(map[string]string)
	localVarQueryParams := url.Values{}
	localVarFormParams := url.Values{}

	// to determine the Content-Type header
	localVarHTTPContentTypes := []string{}

	// set Content-Type header
	localVarHTTPContentType := selectHeaderContentType(localVarHTTPContentTypes)
	if localVarHTTPContentType != "" {
		localVarHeaderParams["Content-Type"] = localVarHTTPContentType
	}

	// to determine the Accept header
	localVarHTTPHeaderAccepts := []string{"application/json"}

	// set Accept header
	localVarHTTPHeaderAccept := selectHeaderAccept(localVarHTTPHeaderAccepts)
	if localVarHTTPHeaderAccept != "" {
		localVarHeaderParams["Accept"] = localVarHTTPHeaderAccept
	}
	req, err := a.client.prepareRequest(r.ctx, localVarPath, localVarHTTPMethod, localVarPostBody, localVarHeaderParams, localVarQueryParams, localVarFormParams, formFiles)
	if err != nil {
		return localVarReturnValue, nil, err
	}

	localVarHTTPResponse, err := a.client.callAPI(req)
	if err != nil || localVarHTTPResponse == nil {
		return localVarReturnValue, localVarHTTPResponse, err
	}

	localVarBody, err := io.ReadAll(localVarHTTPResponse.Body)
	localVarHTTPResponse.Body.Close()
	localVarHTTPResponse.Body = io.NopCloser(bytes.NewBuffer(localVarBody))
	if err != nil {
		return localVarReturnValue, localVarHTTPResponse, err
	}

	if localVarHTTPResponse.StatusCode >= 300 {
		newErr := &GenericOpenAPIError{
			body:  localVarBody,
			error: localVarHTTPResponse.Status,
		}
		return localVarReturnValue, localVarHTTPResponse, newErr
	}

	err = a.client.decode(&localVarReturnValue, localVarBody, localVarHTTPResponse.Header.Get("Content-Type"))
	if err != nil {
		newErr := &GenericOpenAPIError{
			body:  localVarBody,
			error: err.Error(),
		}
		return localVarReturnValue, localVarHTTPResponse, newErr
	}

	return localVarReturnValue, localVarHTTPResponse, nil
}
