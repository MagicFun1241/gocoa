#import <Cocoa/Cocoa.h>
#import <Foundation/Foundation.h>

// Layout constraint creation
void* LayoutConstraint_Create(void *item1, int attr1, int relation, void *item2, int attr2, double multiplier, double constant);

// Layout constraint activation
void LayoutConstraint_Activate(void **constraints, int count);
void LayoutConstraint_ActivateConstraint(void *constraint);
void LayoutConstraint_DeactivateConstraint(void *constraint);

// Constraint attributes (NSLayoutAttribute equivalents)
#define LayoutAttributeLeft 1
#define LayoutAttributeRight 2
#define LayoutAttributeTop 3
#define LayoutAttributeBottom 4
#define LayoutAttributeLeading 5
#define LayoutAttributeTrailing 6
#define LayoutAttributeWidth 7
#define LayoutAttributeHeight 8
#define LayoutAttributeCenterX 9
#define LayoutAttributeCenterY 10
#define LayoutAttributeLastBaseline 11
#define LayoutAttributeFirstBaseline 12
#define LayoutAttributeLeftMargin 13
#define LayoutAttributeRightMargin 14
#define LayoutAttributeTopMargin 15
#define LayoutAttributeBottomMargin 16
#define LayoutAttributeLeadingMargin 17
#define LayoutAttributeTrailingMargin 18
#define LayoutAttributeCenterXWithinMargins 19
#define LayoutAttributeCenterYWithinMargins 20
#define LayoutAttributeNotAnAttribute 0

// Constraint relations (NSLayoutRelation equivalents)
#define LayoutRelationLessThanOrEqual -1
#define LayoutRelationEqual 0
#define LayoutRelationGreaterThanOrEqual 1

// Anchor-based constraint creation (modern approach)
void* View_CenterXAnchor(void *viewPtr);
void* View_CenterYAnchor(void *viewPtr);
void* View_LeftAnchor(void *viewPtr);
void* View_RightAnchor(void *viewPtr);
void* View_TopAnchor(void *viewPtr);
void* View_BottomAnchor(void *viewPtr);
void* View_WidthAnchor(void *viewPtr);
void* View_HeightAnchor(void *viewPtr);
void* View_LeadingAnchor(void *viewPtr);
void* View_TrailingAnchor(void *viewPtr);

// Anchor constraint creation
void* Anchor_ConstraintEqualToAnchor(void *anchor1, void *anchor2);
void* Anchor_ConstraintEqualToAnchorConstant(void *anchor1, void *anchor2, double constant);

// Constraint activation (batch)
void Constraints_Activate(void **constraints, int count); 