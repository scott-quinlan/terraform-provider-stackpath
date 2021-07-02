// Code generated by go-swagger; DO NOT EDIT.

package workload_models

// This file was generated by the swagger tool.
// Editing this file might prove futile when you re-run the swagger generate command

import (
	"context"
	"encoding/json"

	"github.com/go-openapi/errors"
	"github.com/go-openapi/strfmt"
	"github.com/go-openapi/validate"
)

// VolumeClaimVolumeClaimPhase Which phase a volume claim is currently in
//
// - VOLUME_CLAIM_PHASE_UNSPECIFIED: StackPath is unable to determine the volume claim's state
//  - PENDING: The volume claim is pending
//  - UNBOUND: The volume claim is unbound
//  - BOUND: The volume claim is bound to an instance
//
// swagger:model VolumeClaimVolumeClaimPhase
type VolumeClaimVolumeClaimPhase string

func NewVolumeClaimVolumeClaimPhase(value VolumeClaimVolumeClaimPhase) *VolumeClaimVolumeClaimPhase {
	v := value
	return &v
}

const (

	// VolumeClaimVolumeClaimPhaseVOLUMECLAIMPHASEUNSPECIFIED captures enum value "VOLUME_CLAIM_PHASE_UNSPECIFIED"
	VolumeClaimVolumeClaimPhaseVOLUMECLAIMPHASEUNSPECIFIED VolumeClaimVolumeClaimPhase = "VOLUME_CLAIM_PHASE_UNSPECIFIED"

	// VolumeClaimVolumeClaimPhasePENDING captures enum value "PENDING"
	VolumeClaimVolumeClaimPhasePENDING VolumeClaimVolumeClaimPhase = "PENDING"

	// VolumeClaimVolumeClaimPhaseUNBOUND captures enum value "UNBOUND"
	VolumeClaimVolumeClaimPhaseUNBOUND VolumeClaimVolumeClaimPhase = "UNBOUND"

	// VolumeClaimVolumeClaimPhaseBOUND captures enum value "BOUND"
	VolumeClaimVolumeClaimPhaseBOUND VolumeClaimVolumeClaimPhase = "BOUND"
)

// for schema
var volumeClaimVolumeClaimPhaseEnum []interface{}

func init() {
	var res []VolumeClaimVolumeClaimPhase
	if err := json.Unmarshal([]byte(`["VOLUME_CLAIM_PHASE_UNSPECIFIED","PENDING","UNBOUND","BOUND"]`), &res); err != nil {
		panic(err)
	}
	for _, v := range res {
		volumeClaimVolumeClaimPhaseEnum = append(volumeClaimVolumeClaimPhaseEnum, v)
	}
}

func (m VolumeClaimVolumeClaimPhase) validateVolumeClaimVolumeClaimPhaseEnum(path, location string, value VolumeClaimVolumeClaimPhase) error {
	if err := validate.EnumCase(path, location, value, volumeClaimVolumeClaimPhaseEnum, true); err != nil {
		return err
	}
	return nil
}

// Validate validates this volume claim volume claim phase
func (m VolumeClaimVolumeClaimPhase) Validate(formats strfmt.Registry) error {
	var res []error

	// value enum
	if err := m.validateVolumeClaimVolumeClaimPhaseEnum("", "body", m); err != nil {
		return err
	}

	if len(res) > 0 {
		return errors.CompositeValidationError(res...)
	}
	return nil
}

// ContextValidate validates this volume claim volume claim phase based on context it is used
func (m VolumeClaimVolumeClaimPhase) ContextValidate(ctx context.Context, formats strfmt.Registry) error {
	return nil
}
