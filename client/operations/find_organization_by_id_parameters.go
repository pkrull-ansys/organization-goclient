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
)

// NewFindOrganizationByIDParams creates a new FindOrganizationByIDParams object
// with the default values initialized.
func NewFindOrganizationByIDParams() *FindOrganizationByIDParams {
	var ()
	return &FindOrganizationByIDParams{

		timeout: cr.DefaultTimeout,
	}
}

// NewFindOrganizationByIDParamsWithTimeout creates a new FindOrganizationByIDParams object
// with the default values initialized, and the ability to set a timeout on a request
func NewFindOrganizationByIDParamsWithTimeout(timeout time.Duration) *FindOrganizationByIDParams {
	var ()
	return &FindOrganizationByIDParams{

		timeout: timeout,
	}
}

// NewFindOrganizationByIDParamsWithContext creates a new FindOrganizationByIDParams object
// with the default values initialized, and the ability to set a context for a request
func NewFindOrganizationByIDParamsWithContext(ctx context.Context) *FindOrganizationByIDParams {
	var ()
	return &FindOrganizationByIDParams{

		Context: ctx,
	}
}

/*FindOrganizationByIDParams contains all the parameters to send to the API endpoint
for the find organization by Id operation typically these are written to a http.Request
*/
type FindOrganizationByIDParams struct {

	/*ID
	  ID of organization to fetch

	*/
	ID int64

	timeout    time.Duration
	Context    context.Context
	HTTPClient *http.Client
}

// WithTimeout adds the timeout to the find organization by Id params
func (o *FindOrganizationByIDParams) WithTimeout(timeout time.Duration) *FindOrganizationByIDParams {
	o.SetTimeout(timeout)
	return o
}

// SetTimeout adds the timeout to the find organization by Id params
func (o *FindOrganizationByIDParams) SetTimeout(timeout time.Duration) {
	o.timeout = timeout
}

// WithContext adds the context to the find organization by Id params
func (o *FindOrganizationByIDParams) WithContext(ctx context.Context) *FindOrganizationByIDParams {
	o.SetContext(ctx)
	return o
}

// SetContext adds the context to the find organization by Id params
func (o *FindOrganizationByIDParams) SetContext(ctx context.Context) {
	o.Context = ctx
}

// WithID adds the id to the find organization by Id params
func (o *FindOrganizationByIDParams) WithID(id int64) *FindOrganizationByIDParams {
	o.SetID(id)
	return o
}

// SetID adds the id to the find organization by Id params
func (o *FindOrganizationByIDParams) SetID(id int64) {
	o.ID = id
}

// WriteToRequest writes these params to a swagger request
func (o *FindOrganizationByIDParams) WriteToRequest(r runtime.ClientRequest, reg strfmt.Registry) error {

	r.SetTimeout(o.timeout)
	var res []error

	// path param id
	if err := r.SetPathParam("id", swag.FormatInt64(o.ID)); err != nil {
		return err
	}

	if len(res) > 0 {
		return errors.CompositeValidationError(res...)
	}
	return nil
}
