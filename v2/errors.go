package deeplapi

import "errors"

// About Error:
//
// https://www.deepl.com/docs-api/accessing-the-api/authentication/#error-handling
//
//   Code	Description
//   ----	-----------
//   400	Bad request. Please check error message and your parameters.
//   403	Authorization failed. Please supply a valid auth_key parameter.
//   404	The requested resource could not be found.
//   413	The request size exceeds the limit.
//   414	The request URL is too long. You can avoid this error by using a POST request instead of a GET request.
//   429	Too many requests. Please wait and resend your request.
//   456	Quota exceeded. The character limit has been reached.
//   503	Resource currently unavailable. Try again later.
//   529	Too many requests. Please wait and resend your request.
//   5**	Internal error
//
// Additional information may be provided by a JSON response that contains more details about the error.
// In this case, this additional information will be contained in the message key.
//
//   GET /v2/translate?text=Translate it HTTP/1.0
//   Host: api-free.deepl.com
//   User-Agent: YourApp
//   Authorization: DeepL-Auth-Key f9999999-c999-b999-1999-799999999999
//   Accept: application/json
//
//   { "message": "Parameter 'target_lang' not specified." }

var (
	ErrBadRequest           = errors.New("bad request")
	ErrAuthorizationFailed  = errors.New("authorization failed")
	ErrNotFound             = errors.New("resource not found")
	ErrRequstSizeExceeds    = errors.New("request size exceeds the limit")
	ErrRequestURLTooLong    = errors.New("request URL is too long")
	ErrTooManyRequests      = errors.New("too many requests")
	ErrQuotaExceeded        = errors.New("quota exceeded")
	ErrCurrentlyUnavailable = errors.New("resource currently unavailable")
	ErrTooManyRequests5xx   = errors.New("internal error cause too many requests")
	ErrUnexpected           = errors.New("internal server error")
)
