// Code generated by templ - DO NOT EDIT.

// templ: version: v0.2.778
package components

//lint:file-ignore SA4006 This context is only used if a nested component is present.

import "github.com/a-h/templ"
import templruntime "github.com/a-h/templ/runtime"

// ToggleSize represents the size of the toggle
type ToggleSize string

// ToggleLabelPlacement represents where the label should be placed
type ToggleLabelPlacement string

// ToggleProps defines the properties for the Toggle component
type ToggleProps struct {
	// ID is the unique identifier for the toggle input
	ID string

	// Name is the name attribute for the toggle input
	Name string

	// Label is the text label for the toggle
	LabelLeft string

	// Label is the text label for the toggle
	LabelRight string

	// It's treated as an Alpine.js expression for dynamic checking.
	// Example bool: Checked: "true"
	// Example string: Checked: "darkMode"
	Checked string

	// It's treated as an Alpine.js expression for dynamic checking.
	// Example bool: Disabled: "true"
	// Example string: Disabled: "isLoading"
	Disabled string

	// Class specifies additional CSS classes
	Class string

	// Attributes allows passing additional HTML attributes
	Attributes templ.Attributes
}

// Toggle renders a toggle switch component based on the provided props.
// It can be customized with different label placements, and supports
// both static and dynamic states through Alpine.js integration.
//
// Usage:
//
//	// Basic toggle with label
//	@components.Toggle(components.ToggleProps{
//	    ID: "dark-mode",
//	    Name: "darkMode",
//	    LabelLeft: "Dark Mode",
//	})
//
//	// Toggle with Alpine.js binding
//	@components.Toggle(components.ToggleProps{
//	    ID: "notifications",
//	    Name: "notifications",
//	    LabelLeft: "Enable Notifications",
//	    Checked: "notificationsEnabled",
//	    Disabled: "isLoading",
//	})
//
//	// Toggle with custom size and label placement
//	@components.Toggle(components.ToggleProps{
//	    ID: "alerts",
//	    Name: "alerts",
//	    LabelLeft: "Alerts",
//	})
//
// Props:
//   - ID: The unique identifier for the toggle input. Required.
//   - Name: The name attribute for the toggle input. Required.
//   - LabelLeft: The text label for the toggle. Optional.
//   - LabelRight: The text label for the toggle. Optional.
//   - Checked: Controls the checked state. Can be bool or string for Alpine.js binding. Optional.
//   - Disabled: Controls the disabled state. Can be bool or string for Alpine.js binding. Optional.
//   - Class: Additional CSS classes. Optional.
//   - Attributes: Additional HTML attributes. Optional.
func Toggle(props ToggleProps) templ.Component {
	return templruntime.GeneratedTemplate(func(templ_7745c5c3_Input templruntime.GeneratedComponentInput) (templ_7745c5c3_Err error) {
		templ_7745c5c3_W, ctx := templ_7745c5c3_Input.Writer, templ_7745c5c3_Input.Context
		if templ_7745c5c3_CtxErr := ctx.Err(); templ_7745c5c3_CtxErr != nil {
			return templ_7745c5c3_CtxErr
		}
		templ_7745c5c3_Buffer, templ_7745c5c3_IsBuffer := templruntime.GetBuffer(templ_7745c5c3_W)
		if !templ_7745c5c3_IsBuffer {
			defer func() {
				templ_7745c5c3_BufErr := templruntime.ReleaseBuffer(templ_7745c5c3_Buffer)
				if templ_7745c5c3_Err == nil {
					templ_7745c5c3_Err = templ_7745c5c3_BufErr
				}
			}()
		}
		ctx = templ.InitializeContext(ctx)
		templ_7745c5c3_Var1 := templ.GetChildren(ctx)
		if templ_7745c5c3_Var1 == nil {
			templ_7745c5c3_Var1 = templ.NopComponent
		}
		ctx = templ.ClearChildren(ctx)
		_, templ_7745c5c3_Err = templ_7745c5c3_Buffer.WriteString("<div x-data=\"{ checked: false }\" class=\"flex items-center gap-2\"><input id=\"")
		if templ_7745c5c3_Err != nil {
			return templ_7745c5c3_Err
		}
		var templ_7745c5c3_Var2 string
		templ_7745c5c3_Var2, templ_7745c5c3_Err = templ.JoinStringErrs(props.ID)
		if templ_7745c5c3_Err != nil {
			return templ.Error{Err: templ_7745c5c3_Err, FileName: `pkg/components/toggle.templ`, Line: 81, Col: 16}
		}
		_, templ_7745c5c3_Err = templ_7745c5c3_Buffer.WriteString(templ.EscapeString(templ_7745c5c3_Var2))
		if templ_7745c5c3_Err != nil {
			return templ_7745c5c3_Err
		}
		_, templ_7745c5c3_Err = templ_7745c5c3_Buffer.WriteString("\" type=\"checkbox\" name=\"")
		if templ_7745c5c3_Err != nil {
			return templ_7745c5c3_Err
		}
		var templ_7745c5c3_Var3 string
		templ_7745c5c3_Var3, templ_7745c5c3_Err = templ.JoinStringErrs(props.Name)
		if templ_7745c5c3_Err != nil {
			return templ.Error{Err: templ_7745c5c3_Err, FileName: `pkg/components/toggle.templ`, Line: 83, Col: 20}
		}
		_, templ_7745c5c3_Err = templ_7745c5c3_Buffer.WriteString(templ.EscapeString(templ_7745c5c3_Var3))
		if templ_7745c5c3_Err != nil {
			return templ_7745c5c3_Err
		}
		_, templ_7745c5c3_Err = templ_7745c5c3_Buffer.WriteString("\" class=\"hidden\"")
		if templ_7745c5c3_Err != nil {
			return templ_7745c5c3_Err
		}
		if props.Checked != "" {
			_, templ_7745c5c3_Err = templ_7745c5c3_Buffer.WriteString(" x-init=\"")
			if templ_7745c5c3_Err != nil {
				return templ_7745c5c3_Err
			}
			var templ_7745c5c3_Var4 string
			templ_7745c5c3_Var4, templ_7745c5c3_Err = templ.JoinStringErrs("checked = " + props.Checked)
			if templ_7745c5c3_Err != nil {
				return templ.Error{Err: templ_7745c5c3_Err, FileName: `pkg/components/toggle.templ`, Line: 86, Col: 41}
			}
			_, templ_7745c5c3_Err = templ_7745c5c3_Buffer.WriteString(templ.EscapeString(templ_7745c5c3_Var4))
			if templ_7745c5c3_Err != nil {
				return templ_7745c5c3_Err
			}
			_, templ_7745c5c3_Err = templ_7745c5c3_Buffer.WriteString("\"")
			if templ_7745c5c3_Err != nil {
				return templ_7745c5c3_Err
			}
		}
		if props.Disabled != "" {
			_, templ_7745c5c3_Err = templ_7745c5c3_Buffer.WriteString(" :disabled=\"")
			if templ_7745c5c3_Err != nil {
				return templ_7745c5c3_Err
			}
			var templ_7745c5c3_Var5 string
			templ_7745c5c3_Var5, templ_7745c5c3_Err = templ.JoinStringErrs(props.Disabled)
			if templ_7745c5c3_Err != nil {
				return templ.Error{Err: templ_7745c5c3_Err, FileName: `pkg/components/toggle.templ`, Line: 89, Col: 30}
			}
			_, templ_7745c5c3_Err = templ_7745c5c3_Buffer.WriteString(templ.EscapeString(templ_7745c5c3_Var5))
			if templ_7745c5c3_Err != nil {
				return templ_7745c5c3_Err
			}
			_, templ_7745c5c3_Err = templ_7745c5c3_Buffer.WriteString("\"")
			if templ_7745c5c3_Err != nil {
				return templ_7745c5c3_Err
			}
		}
		_, templ_7745c5c3_Err = templ_7745c5c3_Buffer.WriteString(" x-model=\"checked\"> ")
		if templ_7745c5c3_Err != nil {
			return templ_7745c5c3_Err
		}
		if props.LabelLeft != "" {
			_, templ_7745c5c3_Err = templ_7745c5c3_Buffer.WriteString("<label @click=\"$refs.switchButton.click(); $refs.switchButton.focus()\" :id=\"$id(&#39;switch&#39;)\" :class=\"{ &#39;text-foreground&#39;: checked, &#39;text-muted-foreground&#39;: !checked, &#39;opacity-50 cursor-not-allowed&#39;: disabled }\" class=\"text-sm select-none\" x-cloak>")
			if templ_7745c5c3_Err != nil {
				return templ_7745c5c3_Err
			}
			var templ_7745c5c3_Var6 string
			templ_7745c5c3_Var6, templ_7745c5c3_Err = templ.JoinStringErrs(props.LabelLeft)
			if templ_7745c5c3_Err != nil {
				return templ.Error{Err: templ_7745c5c3_Err, FileName: `pkg/components/toggle.templ`, Line: 101, Col: 21}
			}
			_, templ_7745c5c3_Err = templ_7745c5c3_Buffer.WriteString(templ.EscapeString(templ_7745c5c3_Var6))
			if templ_7745c5c3_Err != nil {
				return templ_7745c5c3_Err
			}
			_, templ_7745c5c3_Err = templ_7745c5c3_Buffer.WriteString("</label> ")
			if templ_7745c5c3_Err != nil {
				return templ_7745c5c3_Err
			}
		}
		var templ_7745c5c3_Var7 = []any{
			"relative inline-flex h-6 py-0.5 focus:outline-none rounded-full w-10",
			"disabled:opacity-50 disabled:cursor-not-allowed",
		}
		templ_7745c5c3_Err = templ.RenderCSSItems(ctx, templ_7745c5c3_Buffer, templ_7745c5c3_Var7...)
		if templ_7745c5c3_Err != nil {
			return templ_7745c5c3_Err
		}
		_, templ_7745c5c3_Err = templ_7745c5c3_Buffer.WriteString("<button x-ref=\"switchButton\" type=\"button\" @click=\"checked = !checked\" class=\"")
		if templ_7745c5c3_Err != nil {
			return templ_7745c5c3_Err
		}
		var templ_7745c5c3_Var8 string
		templ_7745c5c3_Var8, templ_7745c5c3_Err = templ.JoinStringErrs(templ.CSSClasses(templ_7745c5c3_Var7).String())
		if templ_7745c5c3_Err != nil {
			return templ.Error{Err: templ_7745c5c3_Err, FileName: `pkg/components/toggle.templ`, Line: 1, Col: 0}
		}
		_, templ_7745c5c3_Err = templ_7745c5c3_Buffer.WriteString(templ.EscapeString(templ_7745c5c3_Var8))
		if templ_7745c5c3_Err != nil {
			return templ_7745c5c3_Err
		}
		_, templ_7745c5c3_Err = templ_7745c5c3_Buffer.WriteString("\" :class=\"checked ? &#39;bg-primary&#39; : &#39;bg-neutral-200&#39;\" :disabled=\"")
		if templ_7745c5c3_Err != nil {
			return templ_7745c5c3_Err
		}
		var templ_7745c5c3_Var9 string
		templ_7745c5c3_Var9, templ_7745c5c3_Err = templ.JoinStringErrs(props.Disabled)
		if templ_7745c5c3_Err != nil {
			return templ.Error{Err: templ_7745c5c3_Err, FileName: `pkg/components/toggle.templ`, Line: 113, Col: 29}
		}
		_, templ_7745c5c3_Err = templ_7745c5c3_Buffer.WriteString(templ.EscapeString(templ_7745c5c3_Var9))
		if templ_7745c5c3_Err != nil {
			return templ_7745c5c3_Err
		}
		_, templ_7745c5c3_Err = templ_7745c5c3_Buffer.WriteString("\" x-cloak><span :class=\"checked ? &#39;translate-x-[18px] bg-secondary&#39; : &#39;translate-x-0.5 bg-muted-foreground&#39;\" class=\"w-5 h-5 duration-200 ease-in-out rounded-full shadow-md\"></span></button> ")
		if templ_7745c5c3_Err != nil {
			return templ_7745c5c3_Err
		}
		if props.LabelRight != "" {
			_, templ_7745c5c3_Err = templ_7745c5c3_Buffer.WriteString("<label @click=\"$refs.switchButton.click(); $refs.switchButton.focus()\" :id=\"$id(&#39;switch&#39;)\" :class=\"{ &#39;text-foreground&#39;: checked, &#39;text-muted-foreground&#39;: !checked, &#39;opacity-50 cursor-not-allowed&#39;: disabled }\" class=\"text-sm select-none\" x-cloak>")
			if templ_7745c5c3_Err != nil {
				return templ_7745c5c3_Err
			}
			var templ_7745c5c3_Var10 string
			templ_7745c5c3_Var10, templ_7745c5c3_Err = templ.JoinStringErrs(props.LabelRight)
			if templ_7745c5c3_Err != nil {
				return templ.Error{Err: templ_7745c5c3_Err, FileName: `pkg/components/toggle.templ`, Line: 129, Col: 22}
			}
			_, templ_7745c5c3_Err = templ_7745c5c3_Buffer.WriteString(templ.EscapeString(templ_7745c5c3_Var10))
			if templ_7745c5c3_Err != nil {
				return templ_7745c5c3_Err
			}
			_, templ_7745c5c3_Err = templ_7745c5c3_Buffer.WriteString("</label>")
			if templ_7745c5c3_Err != nil {
				return templ_7745c5c3_Err
			}
		}
		_, templ_7745c5c3_Err = templ_7745c5c3_Buffer.WriteString("</div>")
		if templ_7745c5c3_Err != nil {
			return templ_7745c5c3_Err
		}
		return templ_7745c5c3_Err
	})
}

var _ = templruntime.GeneratedTemplate
