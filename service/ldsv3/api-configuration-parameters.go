package ldsv3

import (
	"fmt"
)

// List calls

// ListLogConfigurationParameter generic get log configuration parameters call
func (lds *Ldsv3) ListLogConfigurationParameter(parameterType string) (*ConfigurationParameterResponse, error) {
	if parameterType == "" {
		return nil, fmt.Errorf("Please provide parameter type", parameterType)
	}

	apiURI := fmt.Sprintf("%s/log-configuration-parameters/%s", basePath, parameterType)

	// Create and execute request
	resp, err := lds.Client.Rclient.R().
		SetResult(ConfigurationParameterResponse{}).
		SetError(LsdErrorv3{}).
		Get(apiURI)

	if err != nil {
		return nil, err
	}

	if resp.IsError() {
		e := resp.Error().(*LsdErrorv3)

		return nil, e
	}

	return resp.Result().(*ConfigurationParameterResponse), nil
}

// ListDeliveryFrequencies returns all available delivery frequencies, each with an id and descriptive value.
// You need the id to create or modify a log delivery configuration.
func (lds *Ldsv3) ListDeliveryFrequencies() (*ConfigurationParameterResponse, error) {
	resp, err := lds.ListLogConfigurationParameter("delivery-frequencies")

	if err != nil {
		return nil, err
	}

	return resp, nil
}

// ListDeliveryThresholds returns all available log delivery thresholds, each with an id and descriptive value.
func (lds *Ldsv3) ListDeliveryThresholds() (*ConfigurationParameterResponse, error) {
	resp, err := lds.ListLogConfigurationParameter("delivery-thresholds")

	if err != nil {
		return nil, err
	}

	return resp, nil
}

// ListLogEncodings returns all available log encoding options.
// You can restrict the response by specifying optional values for the deliveryType and logSourceType,
// since available encoding types are based on these characteristics of a log delivery configuration.
func (lds *Ldsv3) ListLogEncodings(deliveryType, logSourceType string) (*ConfigurationParameterResponse, error) {
	apiURI := fmt.Sprintf("%s/log-configuration-parameters/encodings", basePath)

	query := map[string]string{}

	if deliveryType != "" {
		query["deliveryType"] = deliveryType
	}

	if logSourceType != "" {
		query["logSourceType"] = logSourceType
	}

	// Create and execute request
	resp, err := lds.Client.Rclient.R().
		SetResult(ConfigurationParameterResponse{}).
		SetError(LsdErrorv3{}).
		SetQueryParams(query).
		Get(apiURI)

	if err != nil {
		return nil, err
	}

	if resp.IsError() {
		e := resp.Error().(*LsdErrorv3)

		return nil, e
	}

	return resp.Result().(*ConfigurationParameterResponse), nil
}

// ListMessageSizes returns all available message sizes, each with an id and descriptive value.
// You need the id to create or modify a log delivery configuration.
func (lds *Ldsv3) ListMessageSizes() (*ConfigurationParameterResponse, error) {
	resp, err := lds.ListLogConfigurationParameter("message-sizes")

	if err != nil {
		return nil, err
	}

	return resp, nil
}

// ListContacts returns all contacts to which you have access.
func (lds *Ldsv3) ListContacts() (*ConfigurationParameterResponse, error) {
	resp, err := lds.ListLogConfigurationParameter("contacts")

	if err != nil {
		return nil, err
	}

	return resp, nil
}

// ListNetStorageGroups returns all NetStorage4 groups to which you have access.
func (lds *Ldsv3) ListNetStorageGroups() (*ConfigurationParameterResponse, error) {
	resp, err := lds.ListLogConfigurationParameter("netstorage-groups")

	if err != nil {
		return nil, err
	}

	return resp, nil
}

// GetLogConfigurationParameter generic get log configuration parameters call
func (lds *Ldsv3) GetLogConfigurationParameter(ID, parameterType string) (*GenericConfigurationParameterElement, error) {
	if ID == "" {
		return nil, fmt.Errorf("Please provide %s ID", parameterType)
	}

	if parameterType == "" {
		return nil, fmt.Errorf("Please provide parameter type", parameterType)
	}

	apiURI := fmt.Sprintf("%s/log-configuration-parameters/%s/%s", basePath, parameterType, ID)

	// Create and execute request
	resp, err := lds.Client.Rclient.R().
		SetResult(GenericConfigurationParameterElement{}).
		SetError(LsdErrorv3{}).
		Get(apiURI)

	if err != nil {
		return nil, err
	}

	if resp.IsError() {
		e := resp.Error().(*LsdErrorv3)

		return nil, e
	}

	return resp.Result().(*GenericConfigurationParameterElement), nil
}

// GetDeliveryFrequency returns a specific delivery frequency.
func (lds *Ldsv3) GetDeliveryFrequency(deliveryFrequencyID string) (*GenericConfigurationParameterElement, error) {
	resp, err := lds.GetLogConfigurationParameter(deliveryFrequencyID, "delivery-frequencies")

	if err != nil {
		return nil, err
	}

	return resp, nil
}

// GetDeliveryThreshold returns a specific delivery frequency.
func (lds *Ldsv3) GetDeliveryThreshold(deliveryThresholdID string) (*GenericConfigurationParameterElement, error) {
	resp, err := lds.GetLogConfigurationParameter(deliveryThresholdID, "delivery-thresholds")

	if err != nil {
		return nil, err
	}

	return resp, nil
}

// GetLogFormat returns a specific log format.
// You need this id to specify the log format for a log delivery configuration.
func (lds *Ldsv3) GetLogFormat(logFormatID string) (*GenericConfigurationParameterElement, error) {
	resp, err := lds.GetLogConfigurationParameter(logFormatID, "log-formats")

	if err != nil {
		return nil, err
	}

	return resp, nil
}

// GetLogEncoding returns a specific log encoding.
func (lds *Ldsv3) GetLogEncoding(encodingID string) (*GenericConfigurationParameterElement, error) {
	resp, err := lds.GetLogConfigurationParameter(encodingID, "encodings")

	if err != nil {
		return nil, err
	}

	return resp, nil
}

// GetMessageSize retrieves a specific message size.
func (lds *Ldsv3) GetMessageSize(messageSizeID string) (*GenericConfigurationParameterElement, error) {
	resp, err := lds.GetLogConfigurationParameter(messageSizeID, "message-sizes")

	if err != nil {
		return nil, err
	}

	return resp, nil
}

// GetContact returns a specific contact, assuming the identity associated with the API client has access to it.
func (lds *Ldsv3) GetContact(contactID string) (*GenericConfigurationParameterElement, error) {
	resp, err := lds.GetLogConfigurationParameter(contactID, "contacts")

	if err != nil {
		return nil, err
	}

	return resp, nil
}
