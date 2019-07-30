package Trim

import (

	"fmt"
	"strings"
	"github.com/TIBCOSoftware/flogo-lib/core/activity"
	"github.com/TIBCOSoftware/flogo-lib/logger"
)

// ActivityLog is the default logger for the Log Activity
var activityLog = logger.GetLogger("activity-flogo-Trim")

// MyActivity is a stub for your Activity implementation
type Trim struct {
	metadata *activity.Metadata
}

// NewActivity creates a new activity
func NewActivity(metadata *activity.Metadata) activity.Activity {
	return &Trim{metadata: metadata}
}

// Metadata implements activity.Activity.Metadata
func (a *Trim) Metadata() *activity.Metadata {
	return a.metadata
}

// Eval implements activity.Activity.Eval
func (a *Trim) Eval(ctx activity.Context) (done bool, err error) {
	
		input := ctx.GetInput("InputString").(string)
		
			//a = strings.TrimSpace(input)
			
		ctx.SetOutput("TrimString", strings.TrimSpace(input))

	activityLog.Debugf("Activity has trimmed the leading and trailing spaces Successfully")
	fmt.Println("Activity has trimmed the leading and trailing spaces Successfully")
	
	return true, nil
}

