package data

import (
	"github.com/pso-dev/delivery-dashboard/backend/internal/validator"
)

type Resource struct {
	ID             int64    `json:"id"`
	Name           string   `json:"name"`
	Email          string   `json:"email"`
	JobTitle       string   `json:"jobTitle"`
	Manager        string   `json:"manager"`
	Location       string   `json:"location"`
	WorkGroup      string   `json:"workGroup"`
	Clearance      string   `json:"clearance"`
	Specialties    []string `json:"specialties"`
	Certifications []string `json:"certifications"`
	Active         bool     `json:"active"`
}

func ValidateID(v *validator.Validator, id int64) {
	v.Check(id != 0, "id", "must be provided")
	v.Check(id > 0, "id", "cannot be a negative number")
}

func ValidateName(v *validator.Validator, name string) {
	v.Check(name != "", "name", "must be provided")
	v.Check(len(name) <= 256, "name", "cannot be more than 256 bytes")
}

func ValidateJobTitle(v *validator.Validator, jobTitle string) {
	// TODO: Remove this hard-coding
	jobTitles := []string{}
	v.Check(validator.PermittedValue(jobTitle, jobTitles...), "jobTitle", "does not exist")
}

func ValidateManager(v *validator.Validator, manager string) {
	// TODO: remove this hard-coding
	managers := []string{}
	v.Check(validator.PermittedValue(manager, managers...), "manager", "is not a manager")
}

func ValidateLocation(v *validator.Validator, location string) {
	v.Check(location != "", "location", "must be provided")
	v.Check(len(location) <= 256, "location", "cannot be more than 256 bytes")
}

func ValidateWorkgroup(v *validator.Validator, workgroup string) {
	// TODO: Remove this hard-coding
	workgroups := []string{}
	v.Check(validator.PermittedValue(workgroup, workgroups...), "workgroup", "does not exist")
}

func ValidatorClearance(v *validator.Validator, clearance string) {
	clearances := []string{
		"None",
		"Baseline",
		"NV1",
		"NV2",
		"TSPV",
	}
	v.Check(validator.PermittedValue(clearance, clearances...), "clearance", "must be one of ['None', 'Baseline', 'NV1', 'NV2', 'TSPV']")
}

func ValidateResource(v *validator.Validator, r Resource) {
	ValidateID(v, r.ID)
	ValidateName(v, r.Name)
	ValidateJobTitle(v, r.JobTitle)
	ValidateManager(v, r.Manager)
	ValidateLocation(v, r.Location)
	ValidateWorkgroup(v, r.WorkGroup)
	ValidatorClearance(v, r.Clearance)
	v.Check(validator.Unique(r.Specialties), "specialties", "cannot contain duplicate values")
	v.Check(validator.Unique(r.Certifications), "certifications", "cannot contain duplicate values")
}
