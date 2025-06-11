#import "constraint.h"
#include <stdlib.h>

// Layout constraint creation
void* LayoutConstraint_Create(void *item1, int attr1, int relation, void *item2, int attr2, double multiplier, double constant)
{
    NSView* firstItem = (NSView*)item1;
    NSView* secondItem = (NSView*)item2;
    
    NSLayoutConstraint* constraint = [NSLayoutConstraint 
        constraintWithItem:firstItem
        attribute:(NSLayoutAttribute)attr1
        relatedBy:(NSLayoutRelation)relation
        toItem:secondItem
        attribute:(NSLayoutAttribute)attr2
        multiplier:multiplier
        constant:constant];
    
    return constraint;
}

// Layout constraint activation
void LayoutConstraint_Activate(void **constraints, int count)
{
    NSMutableArray* constraintArray = [NSMutableArray arrayWithCapacity:count];
    
    for (int i = 0; i < count; i++) {
        NSLayoutConstraint* constraint = (NSLayoutConstraint*)constraints[i];
        [constraintArray addObject:constraint];
    }
    
    [NSLayoutConstraint activateConstraints:constraintArray];
}

void LayoutConstraint_ActivateConstraint(void *constraint)
{
    NSLayoutConstraint* layoutConstraint = (NSLayoutConstraint*)constraint;
    layoutConstraint.active = YES;
}

void LayoutConstraint_DeactivateConstraint(void *constraint)
{
    NSLayoutConstraint* layoutConstraint = (NSLayoutConstraint*)constraint;
    layoutConstraint.active = NO;
}

// Anchor-based constraint creation (modern approach)
void* View_CenterXAnchor(void *viewPtr)
{
    NSView* view = (NSView*)viewPtr;
    return view.centerXAnchor;
}

void* View_CenterYAnchor(void *viewPtr)
{
    NSView* view = (NSView*)viewPtr;
    return view.centerYAnchor;
}

void* View_LeftAnchor(void *viewPtr)
{
    NSView* view = (NSView*)viewPtr;
    return view.leftAnchor;
}

void* View_RightAnchor(void *viewPtr)
{
    NSView* view = (NSView*)viewPtr;
    return view.rightAnchor;
}

void* View_TopAnchor(void *viewPtr)
{
    NSView* view = (NSView*)viewPtr;
    return view.topAnchor;
}

void* View_BottomAnchor(void *viewPtr)
{
    NSView* view = (NSView*)viewPtr;
    return view.bottomAnchor;
}

void* View_WidthAnchor(void *viewPtr)
{
    NSView* view = (NSView*)viewPtr;
    return view.widthAnchor;
}

void* View_HeightAnchor(void *viewPtr)
{
    NSView* view = (NSView*)viewPtr;
    return view.heightAnchor;
}

void* View_LeadingAnchor(void *viewPtr)
{
    NSView* view = (NSView*)viewPtr;
    return view.leadingAnchor;
}

void* View_TrailingAnchor(void *viewPtr)
{
    NSView* view = (NSView*)viewPtr;
    return view.trailingAnchor;
}

// Anchor constraint creation
void* Anchor_ConstraintEqualToAnchor(void *anchor1, void *anchor2)
{
    NSLayoutXAxisAnchor* xAnchor1 = (NSLayoutXAxisAnchor*)anchor1;
    NSLayoutXAxisAnchor* xAnchor2 = (NSLayoutXAxisAnchor*)anchor2;
    
    // Try as X axis anchors first
    if ([xAnchor1 isKindOfClass:[NSLayoutXAxisAnchor class]] && 
        [xAnchor2 isKindOfClass:[NSLayoutXAxisAnchor class]]) {
        return [xAnchor1 constraintEqualToAnchor:xAnchor2];
    }
    
    // Try as Y axis anchors
    NSLayoutYAxisAnchor* yAnchor1 = (NSLayoutYAxisAnchor*)anchor1;
    NSLayoutYAxisAnchor* yAnchor2 = (NSLayoutYAxisAnchor*)anchor2;
    
    if ([yAnchor1 isKindOfClass:[NSLayoutYAxisAnchor class]] && 
        [yAnchor2 isKindOfClass:[NSLayoutYAxisAnchor class]]) {
        return [yAnchor1 constraintEqualToAnchor:yAnchor2];
    }
    
    // Try as dimension anchors
    NSLayoutDimension* dimAnchor1 = (NSLayoutDimension*)anchor1;
    NSLayoutDimension* dimAnchor2 = (NSLayoutDimension*)anchor2;
    
    if ([dimAnchor1 isKindOfClass:[NSLayoutDimension class]] && 
        [dimAnchor2 isKindOfClass:[NSLayoutDimension class]]) {
        return [dimAnchor1 constraintEqualToAnchor:dimAnchor2];
    }
    
    return nil;
}

void* Anchor_ConstraintEqualToAnchorConstant(void *anchor1, void *anchor2, double constant)
{
    NSLayoutXAxisAnchor* xAnchor1 = (NSLayoutXAxisAnchor*)anchor1;
    NSLayoutXAxisAnchor* xAnchor2 = (NSLayoutXAxisAnchor*)anchor2;
    
    // Try as X axis anchors first
    if ([xAnchor1 isKindOfClass:[NSLayoutXAxisAnchor class]] && 
        [xAnchor2 isKindOfClass:[NSLayoutXAxisAnchor class]]) {
        return [xAnchor1 constraintEqualToAnchor:xAnchor2 constant:constant];
    }
    
    // Try as Y axis anchors
    NSLayoutYAxisAnchor* yAnchor1 = (NSLayoutYAxisAnchor*)anchor1;
    NSLayoutYAxisAnchor* yAnchor2 = (NSLayoutYAxisAnchor*)anchor2;
    
    if ([yAnchor1 isKindOfClass:[NSLayoutYAxisAnchor class]] && 
        [yAnchor2 isKindOfClass:[NSLayoutYAxisAnchor class]]) {
        return [yAnchor1 constraintEqualToAnchor:yAnchor2 constant:constant];
    }
    
    // Try as dimension anchors
    NSLayoutDimension* dimAnchor1 = (NSLayoutDimension*)anchor1;
    NSLayoutDimension* dimAnchor2 = (NSLayoutDimension*)anchor2;
    
    if ([dimAnchor1 isKindOfClass:[NSLayoutDimension class]] && 
        [dimAnchor2 isKindOfClass:[NSLayoutDimension class]]) {
        return [dimAnchor1 constraintEqualToAnchor:dimAnchor2 constant:constant];
    }
    
    return nil;
}

// Constraint activation (batch)
void Constraints_Activate(void **constraints, int count)
{
    LayoutConstraint_Activate(constraints, count);
} 