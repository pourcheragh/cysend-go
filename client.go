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
	"encoding/json"
	"encoding/xml"
	"errors"
	"fmt"
	"io"
	"log"
	"mime/multipart"
	"net/http"
	"net/http/httputil"
	"net/url"
	"os"
	"path/filepath"
	"reflect"
	"regexp"
	"strconv"
	"strings"
	"time"
	"unicode/utf8"

)

var (
	JsonCheck       = regexp.MustCompile(`(?i:(?:application|text)/(?:[^;]+\+)?json)`)
	XmlCheck        = regexp.MustCompile(`(?i:(?:application|text)/(?:[^;]+\+)?xml)`)
	queryParamSplit = regexp.MustCompile(`(^|&)([^&]+)`)
	queryDescape    = strings.NewReplacer( "%5B", "[", "%5D", "]" )
)

// APIClient manages communication with the CY.SEND® OpenAPI Documentation API v5.2.3
// In most cases there should be only one, shared, APIClient.
type APIClient struct {
	cfg    *Configuration
	common service // Reuse a single struct instead of allocating one for each service on the heap.

	// API Services

	AccessAPI *AccessAPIService

	CYSENDGiftCardAPI *CYSENDGiftCardAPIService

	CustomerAPI *CustomerAPIService

	NotificationsAPI *NotificationsAPIService

	OpenAPIAPI *OpenAPIAPIService

	StoreAPI *StoreAPIService

	VerificationAPI *VerificationAPIService
}

type service struct {
	client *APIClient
}

// NewAPIClient creates a new API client. Requires a userAgent string describing your application.
// optionally a custom http.Client to allow for advanced features such as caching.
func NewAPIClient(cfg *Configuration) *APIClient {
	if cfg.HTTPClient == nil {
		cfg.HTTPClient = http.DefaultClient
	}

	c := &APIClient{}
	c.cfg = cfg
	c.common.client = c

	// API Services
	c.AccessAPI = (*AccessAPIService)(&c.common)
	c.CYSENDGiftCardAPI = (*CYSENDGiftCardAPIService)(&c.common)
	c.CustomerAPI = (*CustomerAPIService)(&c.common)
	c.NotificationsAPI = (*NotificationsAPIService)(&c.common)
	c.OpenAPIAPI = (*OpenAPIAPIService)(&c.common)
	c.StoreAPI = (*StoreAPIService)(&c.common)
	c.VerificationAPI = (*VerificationAPIService)(&c.common)

	return c
}

func atoi(in string) (int, error) {
	return strconv.Atoi(in)
}

// selectHeaderContentType select a content type from the available list.
func selectHeaderContentType(contentTypes []string) string {
	if len(contentTypes) == 0 {
		return ""
	}
	if contains(contentTypes, "application/json") {
		return "application/json"
	}
	return contentTypes[0] // use the first content type specified in 'consumes'
}

// selectHeaderAccept join all accept types and return
func selectHeaderAccept(accepts []string) string {
	if len(accepts) == 0 {
		return ""
	}

	if contains(accepts, "application/json") {
		return "application/json"
	}

	return strings.Join(accepts, ",")
}

// contains is a case insensitive match, finding needle in a haystack
func contains(haystack []string, needle string) bool {
	for _, a := range haystack {
		if strings.EqualFold(a, needle) {
			return true
		}
	}
	return false
}

// Verify optional parameters are of the correct type.
func typeCheckParameter(obj interface{}, expected string, name string) error {
	// Make sure there is an object.
	if obj == nil {
		return nil
	}

	// Check the type is as expected.
	if reflect.TypeOf(obj).String() != expected {
		return fmt.Errorf("expected %s to be of type %s but received %s", name, expected, reflect.TypeOf(obj).String())
	}
	return nil
}

func parameterValueToString( obj interface{}, key string ) string {
	if reflect.TypeOf(obj).Kind() != reflect.Ptr {
		return fmt.Sprintf("%v", obj)
	}
	var param,ok = obj.(MappedNullable)
	if !ok {
		return ""
	}
	dataMap,err := param.ToMap()
	if err != nil {
		return ""
	}
	return fmt.Sprintf("%v", dataMap[key])
}

// parameterAddToHeaderOrQuery adds the provided object to the request header or url query
// supporting deep object syntax
func parameterAddToHeaderOrQuery(headerOrQueryParams interface{}, keyPrefix string, obj interface{}, collectionType string) {
	var v = reflect.ValueOf(obj)
	var value = ""
	if v == reflect.ValueOf(nil) {
		value = "null"
	} else {
		switch v.Kind() {
			case reflect.Invalid:
				value = "invalid"

			case reflect.Struct:
				if t,ok := obj.(MappedNullable); ok {
					dataMap,err := t.ToMap()
					if err != nil {
						return
					}
					parameterAddToHeaderOrQuery(headerOrQueryParams, keyPrefix, dataMap, collectionType)
					return
				}
				if t, ok := obj.(time.Time); ok {
					parameterAddToHeaderOrQuery(headerOrQueryParams, keyPrefix, t.Format(time.RFC3339), collectionType)
					return
				}
				value = v.Type().String() + " value"
			case reflect.Slice:
				var indValue = reflect.ValueOf(obj)
				if indValue == reflect.ValueOf(nil) {
					return
				}
				var lenIndValue = indValue.Len()
				for i:=0;i<lenIndValue;i++ {
					var arrayValue = indValue.Index(i)
					parameterAddToHeaderOrQuery(headerOrQueryParams, keyPrefix, arrayValue.Interface(), collectionType)
				}
				return

			case reflect.Map:
				var indValue = reflect.ValueOf(obj)
				if indValue == reflect.ValueOf(nil) {
					return
				}
				iter := indValue.MapRange()
				for iter.Next() {
					k,v := iter.Key(), iter.Value()
					parameterAddToHeaderOrQuery(headerOrQueryParams, fmt.Sprintf("%s[%s]", keyPrefix, k.String()), v.Interface(), collectionType)
				}
				return

			case reflect.Interface:
				fallthrough
			case reflect.Ptr:
				parameterAddToHeaderOrQuery(headerOrQueryParams, keyPrefix, v.Elem().Interface(), collectionType)
				return

			case reflect.Int, reflect.Int8, reflect.Int16,
				reflect.Int32, reflect.Int64:
				value = strconv.FormatInt(v.Int(), 10)
			case reflect.Uint, reflect.Uint8, reflect.Uint16,
				reflect.Uint32, reflect.Uint64, reflect.Uintptr:
				value = strconv.FormatUint(v.Uint(), 10)
			case reflect.Float32, reflect.Float64:
				value = strconv.FormatFloat(v.Float(), 'g', -1, 32)
			case reflect.Bool:
				value = strconv.FormatBool(v.Bool())
			case reflect.String:
				value = v.String()
			default:
				value = v.Type().String() + " value"
		}
	}

	switch valuesMap := headerOrQueryParams.(type) {
		case url.Values:
			if collectionType == "csv" && valuesMap.Get(keyPrefix) != "" {
				valuesMap.Set(keyPrefix, valuesMap.Get(keyPrefix) + "," + value)
			} else {
				valuesMap.Add(keyPrefix, value)
			}
			break
		case map[string]string:
			valuesMap[keyPrefix] = value
			break
	}
}

// helper for converting interface{} parameters to json strings
func parameterToJson(obj interface{}) (string, error) {
	jsonBuf, err := json.Marshal(obj)
	if err != nil {
		return "", err
	}
	return string(jsonBuf), err
}

// callAPI do the request.
func (c *APIClient) callAPI(request *http.Request) (*http.Response, error) {
	if c.cfg.Debug {
		dump, err := httputil.DumpRequestOut(request, true)
		if err != nil {
			return nil, err
		}
		log.Printf("\n%s\n", string(dump))
	}

	resp, err := c.cfg.HTTPClient.Do(request)
	if err != nil {
		return resp, err
	}

	if c.cfg.Debug {
		dump, err := httputil.DumpResponse(resp, true)
		if err != nil {
			return resp, err
		}
		log.Printf("\n%s\n", string(dump))
	}
	return resp, err
}

// Allow modification of underlying config for alternate implementations and testing
// Caution: modifying the configuration while live can cause data races and potentially unwanted behavior
func (c *APIClient) GetConfig() *Configuration {
	return c.cfg
}

type formFile struct {
		fileBytes []byte
		fileName string
		formFileName string
}

// prepareRequest build the request
func (c *APIClient) prepareRequest(
	ctx context.Context,
	path string, method string,
	postBody interface{},
	headerParams map[string]string,
	queryParams url.Values,
	formParams url.Values,
	formFiles []formFile) (localVarRequest *http.Request, err error) {

	var body *bytes.Buffer

	// Detect postBody type and post.
	if postBody != nil {
		contentType := headerParams["Content-Type"]
		if contentType == "" {
			contentType = detectContentType(postBody)
			headerParams["Content-Type"] = contentType
		}

		body, err = setBody(postBody, contentType)
		if err != nil {
			return nil, err
		}
	}

	// add form parameters and file if available.
	if strings.HasPrefix(headerParams["Content-Type"], "multipart/form-data") && len(formParams) > 0 || (len(formFiles) > 0) {
		if body != nil {
			return nil, errors.New("Cannot specify postBody and multipart form at the same time.")
		}
		body = &bytes.Buffer{}
		w := multipart.NewWriter(body)

		for k, v := range formParams {
			for _, iv := range v {
				if strings.HasPrefix(k, "@") { // file
					err = addFile(w, k[1:], iv)
					if err != nil {
						return nil, err
					}
				} else { // form value
					w.WriteField(k, iv)
				}
			}
		}
		for _, formFile := range formFiles {
			if len(formFile.fileBytes) > 0 && formFile.fileName != "" {
				w.Boundary()
				part, err := w.CreateFormFile(formFile.formFileName, filepath.Base(formFile.fileName))
				if err != nil {
						return nil, err
				}
				_, err = part.Write(formFile.fileBytes)
				if err != nil {
						return nil, err
				}
			}
		}

		// Set the Boundary in the Content-Type
		headerParams["Content-Type"] = w.FormDataContentType()

		// Set Content-Length
		headerParams["Content-Length"] = fmt.Sprintf("%d", body.Len())
		w.Close()
	}

	if strings.HasPrefix(headerParams["Content-Type"], "application/x-www-form-urlencoded") && len(formParams) > 0 {
		if body != nil {
			return nil, errors.New("Cannot specify postBody and x-www-form-urlencoded form at the same time.")
		}
		body = &bytes.Buffer{}
		body.WriteString(formParams.Encode())
		// Set Content-Length
		headerParams["Content-Length"] = fmt.Sprintf("%d", body.Len())
	}

	// Setup path and query parameters
	url, err := url.Parse(path)
	if err != nil {
		return nil, err
	}

	// Override request host, if applicable
	if c.cfg.Host != "" {
		url.Host = c.cfg.Host
	}

	// Override request scheme, if applicable
	if c.cfg.Scheme != "" {
		url.Scheme = c.cfg.Scheme
	}

	// Adding Query Param
	query := url.Query()
	for k, v := range queryParams {
		for _, iv := range v {
			query.Add(k, iv)
		}
	}

	// Encode the parameters.
	url.RawQuery = queryParamSplit.ReplaceAllStringFunc(query.Encode(), func(s string) string {
		pieces := strings.Split(s, "=")
		pieces[0] = queryDescape.Replace(pieces[0])
		return strings.Join(pieces, "=")
	})

	// Generate a new request
	if body != nil {
		localVarRequest, err = http.NewRequest(method, url.String(), body)
	} else {
		localVarRequest, err = http.NewRequest(method, url.String(), nil)
	}
	if err != nil {
		return nil, err
	}

	// add header parameters, if any
	if len(headerParams) > 0 {
		headers := http.Header{}
		for h, v := range headerParams {
			headers[h] = []string{v}
		}
		localVarRequest.Header = headers
	}

	// Add the user agent to the request.
	localVarRequest.Header.Add("User-Agent", c.cfg.UserAgent)

	if ctx != nil {
		// add context to the request
		localVarRequest = localVarRequest.WithContext(ctx)

		// Walk through any authentication.

		// Basic HTTP Authentication
		if auth, ok := ctx.Value(ContextBasicAuth).(BasicAuth); ok {
			localVarRequest.SetBasicAuth(auth.UserName, auth.Password)
		}

	}

	for header, value := range c.cfg.DefaultHeader {
		localVarRequest.Header.Add(header, value)
	}
	return localVarRequest, nil
}

func (c *APIClient) decode(v interface{}, b []byte, contentType string) (err error) {
	if len(b) == 0 {
		return nil
	}
	if s, ok := v.(*string); ok {
		*s = string(b)
		return nil
	}
	if f, ok := v.(*os.File); ok {
		f, err = os.CreateTemp("", "HttpClientFile")
		if err != nil {
			return
		}
		_, err = f.Write(b)
		if err != nil {
			return
		}
		_, err = f.Seek(0, io.SeekStart)
		return
	}
	if f, ok := v.(**os.File); ok {
		*f, err = os.CreateTemp("", "HttpClientFile")
		if err != nil {
			return
		}
		_, err = (*f).Write(b)
		if err != nil {
			return
		}
		_, err = (*f).Seek(0, io.SeekStart)
		return
	}
	if XmlCheck.MatchString(contentType) {
		if err = xml.Unmarshal(b, v); err != nil {
			return err
		}
		return nil
	}
	if JsonCheck.MatchString(contentType) {
		if actualObj, ok := v.(interface{ GetActualInstance() interface{} }); ok { // oneOf, anyOf schemas
			if unmarshalObj, ok := actualObj.(interface{ UnmarshalJSON([]byte) error }); ok { // make sure it has UnmarshalJSON defined
				if err = unmarshalObj.UnmarshalJSON(b); err != nil {
					return err
				}
			} else {
				return errors.New("Unknown type with GetActualInstance but no unmarshalObj.UnmarshalJSON defined")
			}
		} else if err = json.Unmarshal(b, v); err != nil { // simple model
			return err
		}
		return nil
	}
	return errors.New("undefined response type")
}

// Add a file to the multipart request
func addFile(w *multipart.Writer, fieldName, path string) error {
	file, err := os.Open(filepath.Clean(path))
	if err != nil {
		return err
	}
	err = file.Close()
	if err != nil {
		return err
	}

	part, err := w.CreateFormFile(fieldName, filepath.Base(path))
	if err != nil {
		return err
	}
	_, err = io.Copy(part, file)

	return err
}

// Prevent trying to import "fmt"
func reportError(format string, a ...interface{}) error {
	return fmt.Errorf(format, a...)
}

// A wrapper for strict JSON decoding
func newStrictDecoder(data []byte) *json.Decoder {
	dec := json.NewDecoder(bytes.NewBuffer(data))
	dec.DisallowUnknownFields()
	return dec
}

// Set request body from an interface{}
func setBody(body interface{}, contentType string) (bodyBuf *bytes.Buffer, err error) {
	if bodyBuf == nil {
		bodyBuf = &bytes.Buffer{}
	}

	if reader, ok := body.(io.Reader); ok {
		_, err = bodyBuf.ReadFrom(reader)
	} else if fp, ok := body.(*os.File); ok {
		_, err = bodyBuf.ReadFrom(fp)
	} else if b, ok := body.([]byte); ok {
		_, err = bodyBuf.Write(b)
	} else if s, ok := body.(string); ok {
		_, err = bodyBuf.WriteString(s)
	} else if s, ok := body.(*string); ok {
		_, err = bodyBuf.WriteString(*s)
	} else if JsonCheck.MatchString(contentType) {
		err = json.NewEncoder(bodyBuf).Encode(body)
	} else if XmlCheck.MatchString(contentType) {
		var bs []byte
		bs, err = xml.Marshal(body)
		if err == nil {
			bodyBuf.Write(bs)
		}
	}

	if err != nil {
		return nil, err
	}

	if bodyBuf.Len() == 0 {
		err = fmt.Errorf("invalid body type %s\n", contentType)
		return nil, err
	}
	return bodyBuf, nil
}

// detectContentType method is used to figure out `Request.Body` content type for request header
func detectContentType(body interface{}) string {
	contentType := "text/plain; charset=utf-8"
	kind := reflect.TypeOf(body).Kind()

	switch kind {
	case reflect.Struct, reflect.Map, reflect.Ptr:
		contentType = "application/json; charset=utf-8"
	case reflect.String:
		contentType = "text/plain; charset=utf-8"
	default:
		if b, ok := body.([]byte); ok {
			contentType = http.DetectContentType(b)
		} else if kind == reflect.Slice {
			contentType = "application/json; charset=utf-8"
		}
	}

	return contentType
}

// Ripped from https://github.com/gregjones/httpcache/blob/master/httpcache.go
type cacheControl map[string]string

func parseCacheControl(headers http.Header) cacheControl {
	cc := cacheControl{}
	ccHeader := headers.Get("Cache-Control")
	for _, part := range strings.Split(ccHeader, ",") {
		part = strings.Trim(part, " ")
		if part == "" {
			continue
		}
		if strings.ContainsRune(part, '=') {
			keyval := strings.Split(part, "=")
			cc[strings.Trim(keyval[0], " ")] = strings.Trim(keyval[1], ",")
		} else {
			cc[part] = ""
		}
	}
	return cc
}

// CacheExpires helper function to determine remaining time before repeating a request.
func CacheExpires(r *http.Response) time.Time {
	// Figure out when the cache expires.
	var expires time.Time
	now, err := time.Parse(time.RFC1123, r.Header.Get("date"))
	if err != nil {
		return time.Now()
	}
	respCacheControl := parseCacheControl(r.Header)

	if maxAge, ok := respCacheControl["max-age"]; ok {
		lifetime, err := time.ParseDuration(maxAge + "s")
		if err != nil {
			expires = now
		} else {
			expires = now.Add(lifetime)
		}
	} else {
		expiresHeader := r.Header.Get("Expires")
		if expiresHeader != "" {
			expires, err = time.Parse(time.RFC1123, expiresHeader)
			if err != nil {
				expires = now
			}
		}
	}
	return expires
}

func strlen(s string) int {
	return utf8.RuneCountInString(s)
}

// GenericOpenAPIError Provides access to the body, error and model on returned errors.
type GenericOpenAPIError struct {
	body  []byte
	error string
	model interface{}
}

// Error returns non-empty string if there was an error.
func (e GenericOpenAPIError) Error() string {
	return e.error
}

// Body returns the raw bytes of the response
func (e GenericOpenAPIError) Body() []byte {
	return e.body
}

// Model returns the unpacked model of the error
func (e GenericOpenAPIError) Model() interface{} {
	return e.model
}

// format error message using title and detail when model implements rfc7807
func formatErrorMessage(status string, v interface{}) string {
	str := ""
	metaValue := reflect.ValueOf(v).Elem()

	if metaValue.Kind() == reflect.Struct {
		field := metaValue.FieldByName("Title")
		if field != (reflect.Value{}) {
			str = fmt.Sprintf("%s", field.Interface())
		}

		field = metaValue.FieldByName("Detail")
		if field != (reflect.Value{}) {
			str = fmt.Sprintf("%s (%s)", str, field.Interface())
		}
	}

	return strings.TrimSpace(fmt.Sprintf("%s %s", status, str))
}
