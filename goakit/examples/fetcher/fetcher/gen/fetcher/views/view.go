// Code generated by goa v2.0.0-wip, DO NOT EDIT.
//
// fetcher views
//
// Command:
// $ goa gen goa.design/plugins/goakit/examples/fetcher/fetcher/design

package views

import (
	goa "goa.design/goa"
)

// FetchMedia is the viewed result type that is projected based on a view.
type FetchMedia struct {
	// Type to project
	Projected *FetchMediaView
	// View to render
	View string
}

// FetchMediaView is a type that runs validations on a projected type.
type FetchMediaView struct {
	// HTTP status code returned by fetched service
	Status *int
	// The href to the corresponding archive in the archiver service
	ArchiveHref *string
}

// Validate runs the validations defined on the viewed result type FetchMedia.
func (result *FetchMedia) Validate() (err error) {
	switch result.View {
	case "default", "":
		err = result.Projected.Validate()
	default:
		err = goa.InvalidEnumValueError("view", result.View, []interface{}{"default"})
	}
	return
}

// Validate runs the validations defined on FetchMediaView using the "default"
// view.
func (result *FetchMediaView) Validate() (err error) {
	if result.Status == nil {
		err = goa.MergeErrors(err, goa.MissingFieldError("status", "result"))
	}
	if result.ArchiveHref == nil {
		err = goa.MergeErrors(err, goa.MissingFieldError("archive_href", "result"))
	}
	if result.Status != nil {
		if *result.Status < 0 {
			err = goa.MergeErrors(err, goa.InvalidRangeError("result.status", *result.Status, 0, true))
		}
	}
	if result.ArchiveHref != nil {
		err = goa.MergeErrors(err, goa.ValidatePattern("result.archive_href", *result.ArchiveHref, "^/archive/[0-9]+$"))
	}
	return
}
