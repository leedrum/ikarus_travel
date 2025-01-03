// Code generated by templ - DO NOT EDIT.

// templ: version: v0.2.778
package icons

//lint:file-ignore SA4006 This context is only used if a nested component is present.

import "github.com/a-h/templ"
import templruntime "github.com/a-h/templ/runtime"

func Adult() templ.Component {
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
		_, templ_7745c5c3_Err = templ_7745c5c3_Buffer.WriteString("<svg class=\"w-5 h-5\" width=\"24\" height=\"24\" viewBox=\"0 0 24 24\" fill=\"none\" xmlns=\"http://www.w3.org/2000/svg\"><path fill-rule=\"evenodd\" clip-rule=\"evenodd\" d=\"M8 4C5.79086 4 4 5.79086 4 8C4 10.2091 5.79086 12 8 12C10.2091 12 12 10.2091 12 8C12 5.79086 10.2091 4 8 4ZM6 13C3.79086 13 2 14.7909 2 17V18C2 19.1046 2.89543 20 4 20H12C13.1046 20 14 19.1046 14 18V17C14 14.7909 12.2091 13 10 13H6Z\" fill=\"#383838\"></path> <path fill-rule=\"evenodd\" clip-rule=\"evenodd\" d=\"M13.2508 10.9055C13.7282 10.0447 14 9.05408 14 8C14 6.94592 13.7282 5.95532 13.2508 5.09447C13.9676 4.41606 14.9352 4 16 4C18.2091 4 20 5.79086 20 8C20 10.2091 18.2091 12 16 12C14.9352 12 13.9676 11.5839 13.2508 10.9055ZM15.4649 20C15.8052 19.4117 16 18.7286 16 18V17C16 15.4633 15.4223 14.0615 14.4722 13H18C20.2091 13 22 14.7909 22 17V18C22 19.1046 21.1046 20 20 20H15.4649Z\" fill=\"#383838\"></path></svg>")
		if templ_7745c5c3_Err != nil {
			return templ_7745c5c3_Err
		}
		return templ_7745c5c3_Err
	})
}

var _ = templruntime.GeneratedTemplate
