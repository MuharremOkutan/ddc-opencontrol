package operations

// This file was generated by the swagger tool.
// Editing this file might prove futile when you re-run the swagger generate command

import (
	"net/http"
	"time"

	"golang.org/x/net/context"

	"github.com/go-openapi/errors"
	"github.com/go-openapi/runtime"
	cr "github.com/go-openapi/runtime/client"
	"github.com/go-openapi/swag"

	strfmt "github.com/go-openapi/strfmt"

	"github.com/docker/ddc-opencontrol/nlp/textanalytics/models"
)

// NewDetectLanguageParams creates a new DetectLanguageParams object
// with the default values initialized.
func NewDetectLanguageParams() *DetectLanguageParams {
	var ()
	return &DetectLanguageParams{

		timeout: cr.DefaultTimeout,
	}
}

// NewDetectLanguageParamsWithTimeout creates a new DetectLanguageParams object
// with the default values initialized, and the ability to set a timeout on a request
func NewDetectLanguageParamsWithTimeout(timeout time.Duration) *DetectLanguageParams {
	var ()
	return &DetectLanguageParams{

		timeout: timeout,
	}
}

// NewDetectLanguageParamsWithContext creates a new DetectLanguageParams object
// with the default values initialized, and the ability to set a context for a request
func NewDetectLanguageParamsWithContext(ctx context.Context) *DetectLanguageParams {
	var ()
	return &DetectLanguageParams{

		Context: ctx,
	}
}

/*DetectLanguageParams contains all the parameters to send to the API endpoint
for the detect language operation typically these are written to a http.Request
*/
type DetectLanguageParams struct {

	/*OcpApimSubscriptionKey
	  subscription key in header

	*/
	OcpApimSubscriptionKey *string
	/*BatchInputV2*/
	BatchInputV2 *models.BatchInputV2
	/*NumberOfLanguagesToDetect
	  Format - int32. (Optional) Number of languages to detect. Set to 1 by default.

	*/
	NumberOfLanguagesToDetect *int64
	/*SubscriptionKey
	  subscription key in url

	*/
	SubscriptionKey *string

	timeout    time.Duration
	Context    context.Context
	HTTPClient *http.Client
}

// WithTimeout adds the timeout to the detect language params
func (o *DetectLanguageParams) WithTimeout(timeout time.Duration) *DetectLanguageParams {
	o.SetTimeout(timeout)
	return o
}

// SetTimeout adds the timeout to the detect language params
func (o *DetectLanguageParams) SetTimeout(timeout time.Duration) {
	o.timeout = timeout
}

// WithContext adds the context to the detect language params
func (o *DetectLanguageParams) WithContext(ctx context.Context) *DetectLanguageParams {
	o.SetContext(ctx)
	return o
}

// SetContext adds the context to the detect language params
func (o *DetectLanguageParams) SetContext(ctx context.Context) {
	o.Context = ctx
}

// WithOcpApimSubscriptionKey adds the ocpApimSubscriptionKey to the detect language params
func (o *DetectLanguageParams) WithOcpApimSubscriptionKey(ocpApimSubscriptionKey *string) *DetectLanguageParams {
	o.SetOcpApimSubscriptionKey(ocpApimSubscriptionKey)
	return o
}

// SetOcpApimSubscriptionKey adds the ocpApimSubscriptionKey to the detect language params
func (o *DetectLanguageParams) SetOcpApimSubscriptionKey(ocpApimSubscriptionKey *string) {
	o.OcpApimSubscriptionKey = ocpApimSubscriptionKey
}

// WithBatchInputV2 adds the batchInputV2 to the detect language params
func (o *DetectLanguageParams) WithBatchInputV2(batchInputV2 *models.BatchInputV2) *DetectLanguageParams {
	o.SetBatchInputV2(batchInputV2)
	return o
}

// SetBatchInputV2 adds the batchInputV2 to the detect language params
func (o *DetectLanguageParams) SetBatchInputV2(batchInputV2 *models.BatchInputV2) {
	o.BatchInputV2 = batchInputV2
}

// WithNumberOfLanguagesToDetect adds the numberOfLanguagesToDetect to the detect language params
func (o *DetectLanguageParams) WithNumberOfLanguagesToDetect(numberOfLanguagesToDetect *int64) *DetectLanguageParams {
	o.SetNumberOfLanguagesToDetect(numberOfLanguagesToDetect)
	return o
}

// SetNumberOfLanguagesToDetect adds the numberOfLanguagesToDetect to the detect language params
func (o *DetectLanguageParams) SetNumberOfLanguagesToDetect(numberOfLanguagesToDetect *int64) {
	o.NumberOfLanguagesToDetect = numberOfLanguagesToDetect
}

// WithSubscriptionKey adds the subscriptionKey to the detect language params
func (o *DetectLanguageParams) WithSubscriptionKey(subscriptionKey *string) *DetectLanguageParams {
	o.SetSubscriptionKey(subscriptionKey)
	return o
}

// SetSubscriptionKey adds the subscriptionKey to the detect language params
func (o *DetectLanguageParams) SetSubscriptionKey(subscriptionKey *string) {
	o.SubscriptionKey = subscriptionKey
}

// WriteToRequest writes these params to a swagger request
func (o *DetectLanguageParams) WriteToRequest(r runtime.ClientRequest, reg strfmt.Registry) error {

	r.SetTimeout(o.timeout)
	var res []error

	if o.OcpApimSubscriptionKey != nil {

		// header param Ocp-Apim-Subscription-Key
		if err := r.SetHeaderParam("Ocp-Apim-Subscription-Key", *o.OcpApimSubscriptionKey); err != nil {
			return err
		}

	}

	if o.BatchInputV2 == nil {
		o.BatchInputV2 = new(models.BatchInputV2)
	}

	if err := r.SetBodyParam(o.BatchInputV2); err != nil {
		return err
	}

	if o.NumberOfLanguagesToDetect != nil {

		// query param numberOfLanguagesToDetect
		var qrNumberOfLanguagesToDetect int64
		if o.NumberOfLanguagesToDetect != nil {
			qrNumberOfLanguagesToDetect = *o.NumberOfLanguagesToDetect
		}
		qNumberOfLanguagesToDetect := swag.FormatInt64(qrNumberOfLanguagesToDetect)
		if qNumberOfLanguagesToDetect != "" {
			if err := r.SetQueryParam("numberOfLanguagesToDetect", qNumberOfLanguagesToDetect); err != nil {
				return err
			}
		}

	}

	if o.SubscriptionKey != nil {

		// query param subscription-key
		var qrSubscriptionKey string
		if o.SubscriptionKey != nil {
			qrSubscriptionKey = *o.SubscriptionKey
		}
		qSubscriptionKey := qrSubscriptionKey
		if qSubscriptionKey != "" {
			if err := r.SetQueryParam("subscription-key", qSubscriptionKey); err != nil {
				return err
			}
		}

	}

	if len(res) > 0 {
		return errors.CompositeValidationError(res...)
	}
	return nil
}
