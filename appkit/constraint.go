package appkit

// #cgo CFLAGS: -x objective-c
// #cgo LDFLAGS: -framework Cocoa
// #include "constraint.h"
import "C"
import "unsafe"

// LayoutAttribute constants for constraint attributes
type LayoutAttribute int

const (
	LayoutAttributeNotAnAttribute       LayoutAttribute = 0
	LayoutAttributeLeft                 LayoutAttribute = 1
	LayoutAttributeRight                LayoutAttribute = 2
	LayoutAttributeTop                  LayoutAttribute = 3
	LayoutAttributeBottom               LayoutAttribute = 4
	LayoutAttributeLeading              LayoutAttribute = 5
	LayoutAttributeTrailing             LayoutAttribute = 6
	LayoutAttributeWidth                LayoutAttribute = 7
	LayoutAttributeHeight               LayoutAttribute = 8
	LayoutAttributeCenterX              LayoutAttribute = 9
	LayoutAttributeCenterY              LayoutAttribute = 10
	LayoutAttributeLastBaseline         LayoutAttribute = 11
	LayoutAttributeFirstBaseline        LayoutAttribute = 12
	LayoutAttributeLeftMargin           LayoutAttribute = 13
	LayoutAttributeRightMargin          LayoutAttribute = 14
	LayoutAttributeTopMargin            LayoutAttribute = 15
	LayoutAttributeBottomMargin         LayoutAttribute = 16
	LayoutAttributeLeadingMargin        LayoutAttribute = 17
	LayoutAttributeTrailingMargin       LayoutAttribute = 18
	LayoutAttributeCenterXWithinMargins LayoutAttribute = 19
	LayoutAttributeCenterYWithinMargins LayoutAttribute = 20
)

// LayoutRelation constants for constraint relations
type LayoutRelation int

const (
	LayoutRelationLessThanOrEqual    LayoutRelation = -1
	LayoutRelationEqual              LayoutRelation = 0
	LayoutRelationGreaterThanOrEqual LayoutRelation = 1
)

// LayoutConstraint represents an NSLayoutConstraint
type LayoutConstraint struct {
	constraintPtr unsafe.Pointer
}

// LayoutAnchor represents a layout anchor (NSLayoutAnchor)
type LayoutAnchor struct {
	anchorPtr unsafe.Pointer
}

// CreateConstraint creates a new layout constraint
func CreateConstraint(item1 unsafe.Pointer, attr1 LayoutAttribute, relation LayoutRelation, item2 unsafe.Pointer, attr2 LayoutAttribute, multiplier, constant float64) *LayoutConstraint {
	constraintPtr := C.LayoutConstraint_Create(
		item1,
		C.int(attr1),
		C.int(relation),
		item2,
		C.int(attr2),
		C.double(multiplier),
		C.double(constant),
	)
	return &LayoutConstraint{constraintPtr: constraintPtr}
}

// Activate activates this constraint
func (c *LayoutConstraint) Activate() {
	C.LayoutConstraint_ActivateConstraint(c.constraintPtr)
}

// Deactivate deactivates this constraint
func (c *LayoutConstraint) Deactivate() {
	C.LayoutConstraint_DeactivateConstraint(c.constraintPtr)
}

// ActivateConstraints activates multiple constraints at once
func ActivateConstraints(constraints []*LayoutConstraint) {
	if len(constraints) == 0 {
		return
	}

	// Convert Go slice to C array
	cConstraints := make([]unsafe.Pointer, len(constraints))
	for i, constraint := range constraints {
		cConstraints[i] = constraint.constraintPtr
	}

	C.Constraints_Activate(&cConstraints[0], C.int(len(constraints)))
}

// Anchor methods for View

// CenterXAnchor returns the center X anchor for the view
func (v *View) CenterXAnchor() *LayoutAnchor {
	anchorPtr := C.View_CenterXAnchor(v.viewPtr)
	return &LayoutAnchor{anchorPtr: anchorPtr}
}

// CenterYAnchor returns the center Y anchor for the view
func (v *View) CenterYAnchor() *LayoutAnchor {
	anchorPtr := C.View_CenterYAnchor(v.viewPtr)
	return &LayoutAnchor{anchorPtr: anchorPtr}
}

// LeftAnchor returns the left anchor for the view
func (v *View) LeftAnchor() *LayoutAnchor {
	anchorPtr := C.View_LeftAnchor(v.viewPtr)
	return &LayoutAnchor{anchorPtr: anchorPtr}
}

// RightAnchor returns the right anchor for the view
func (v *View) RightAnchor() *LayoutAnchor {
	anchorPtr := C.View_RightAnchor(v.viewPtr)
	return &LayoutAnchor{anchorPtr: anchorPtr}
}

// TopAnchor returns the top anchor for the view
func (v *View) TopAnchor() *LayoutAnchor {
	anchorPtr := C.View_TopAnchor(v.viewPtr)
	return &LayoutAnchor{anchorPtr: anchorPtr}
}

// BottomAnchor returns the bottom anchor for the view
func (v *View) BottomAnchor() *LayoutAnchor {
	anchorPtr := C.View_BottomAnchor(v.viewPtr)
	return &LayoutAnchor{anchorPtr: anchorPtr}
}

// WidthAnchor returns the width anchor for the view
func (v *View) WidthAnchor() *LayoutAnchor {
	anchorPtr := C.View_WidthAnchor(v.viewPtr)
	return &LayoutAnchor{anchorPtr: anchorPtr}
}

// HeightAnchor returns the height anchor for the view
func (v *View) HeightAnchor() *LayoutAnchor {
	anchorPtr := C.View_HeightAnchor(v.viewPtr)
	return &LayoutAnchor{anchorPtr: anchorPtr}
}

// LeadingAnchor returns the leading anchor for the view
func (v *View) LeadingAnchor() *LayoutAnchor {
	anchorPtr := C.View_LeadingAnchor(v.viewPtr)
	return &LayoutAnchor{anchorPtr: anchorPtr}
}

// TrailingAnchor returns the trailing anchor for the view
func (v *View) TrailingAnchor() *LayoutAnchor {
	anchorPtr := C.View_TrailingAnchor(v.viewPtr)
	return &LayoutAnchor{anchorPtr: anchorPtr}
}

// Anchor constraint creation methods

// ConstraintEqualToAnchor creates a constraint that makes this anchor equal to another anchor
func (a *LayoutAnchor) ConstraintEqualToAnchor(anchor *LayoutAnchor) *LayoutConstraint {
	constraintPtr := C.Anchor_ConstraintEqualToAnchor(a.anchorPtr, anchor.anchorPtr)
	return &LayoutConstraint{constraintPtr: constraintPtr}
}

// ConstraintEqualToAnchorConstant creates a constraint that makes this anchor equal to another anchor plus a constant
func (a *LayoutAnchor) ConstraintEqualToAnchorConstant(anchor *LayoutAnchor, constant float64) *LayoutConstraint {
	constraintPtr := C.Anchor_ConstraintEqualToAnchorConstant(a.anchorPtr, anchor.anchorPtr, C.double(constant))
	return &LayoutConstraint{constraintPtr: constraintPtr}
}

// Button anchor methods (since Button is a subclass of NSView)

// CenterXAnchor returns the center X anchor for the button
func (btn *Button) CenterXAnchor() *LayoutAnchor {
	anchorPtr := C.View_CenterXAnchor(unsafe.Pointer(btn.buttonPtr))
	return &LayoutAnchor{anchorPtr: anchorPtr}
}

// CenterYAnchor returns the center Y anchor for the button
func (btn *Button) CenterYAnchor() *LayoutAnchor {
	anchorPtr := C.View_CenterYAnchor(unsafe.Pointer(btn.buttonPtr))
	return &LayoutAnchor{anchorPtr: anchorPtr}
}

// LeftAnchor returns the left anchor for the button
func (btn *Button) LeftAnchor() *LayoutAnchor {
	anchorPtr := C.View_LeftAnchor(unsafe.Pointer(btn.buttonPtr))
	return &LayoutAnchor{anchorPtr: anchorPtr}
}

// RightAnchor returns the right anchor for the button
func (btn *Button) RightAnchor() *LayoutAnchor {
	anchorPtr := C.View_RightAnchor(unsafe.Pointer(btn.buttonPtr))
	return &LayoutAnchor{anchorPtr: anchorPtr}
}

// TopAnchor returns the top anchor for the button
func (btn *Button) TopAnchor() *LayoutAnchor {
	anchorPtr := C.View_TopAnchor(unsafe.Pointer(btn.buttonPtr))
	return &LayoutAnchor{anchorPtr: anchorPtr}
}

// BottomAnchor returns the bottom anchor for the button
func (btn *Button) BottomAnchor() *LayoutAnchor {
	anchorPtr := C.View_BottomAnchor(unsafe.Pointer(btn.buttonPtr))
	return &LayoutAnchor{anchorPtr: anchorPtr}
}

// WidthAnchor returns the width anchor for the button
func (btn *Button) WidthAnchor() *LayoutAnchor {
	anchorPtr := C.View_WidthAnchor(unsafe.Pointer(btn.buttonPtr))
	return &LayoutAnchor{anchorPtr: anchorPtr}
}

// HeightAnchor returns the height anchor for the button
func (btn *Button) HeightAnchor() *LayoutAnchor {
	anchorPtr := C.View_HeightAnchor(unsafe.Pointer(btn.buttonPtr))
	return &LayoutAnchor{anchorPtr: anchorPtr}
}

// LeadingAnchor returns the leading anchor for the button
func (btn *Button) LeadingAnchor() *LayoutAnchor {
	anchorPtr := C.View_LeadingAnchor(unsafe.Pointer(btn.buttonPtr))
	return &LayoutAnchor{anchorPtr: anchorPtr}
}

// TrailingAnchor returns the trailing anchor for the button
func (btn *Button) TrailingAnchor() *LayoutAnchor {
	anchorPtr := C.View_TrailingAnchor(unsafe.Pointer(btn.buttonPtr))
	return &LayoutAnchor{anchorPtr: anchorPtr}
}
