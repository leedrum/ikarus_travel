// Code generated by templ - DO NOT EDIT.

// templ: version: v0.2.778
package icons

//lint:file-ignore SA4006 This context is only used if a nested component is present.

import "github.com/a-h/templ"
import templruntime "github.com/a-h/templ/runtime"

func Hotel() templ.Component {
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
		_, templ_7745c5c3_Err = templ_7745c5c3_Buffer.WriteString("<svg width=\"24\" height=\"24\" viewBox=\"0 0 24 24\" fill=\"none\" xmlns=\"http://www.w3.org/2000/svg\"><path fill-rule=\"evenodd\" clip-rule=\"evenodd\" d=\"M4 4.00012C4 3.44784 4.44772 3.00012 5 3.00012H19C19.5523 3.00012 20 3.44784 20 4.00012C20 4.55241 19.5523 5.00012 19 5.00012V19.0001C19.5523 19.0001 20 19.4478 20 20.0001C20 20.5524 19.5523 21.0001 19 21.0001H5C4.44772 21.0001 4 20.5524 4 20.0001C4 19.4478 4.44772 19.0001 5 19.0001V5.00012C4.44772 5.00012 4 4.55241 4 4.00012ZM9 6.00012C8.44772 6.00012 8 6.44784 8 7.00012V8.00012C8 8.55241 8.44772 9.00012 9 9.00012H10C10.5523 9.00012 11 8.55241 11 8.00012V7.00012C11 6.44784 10.5523 6.00012 10 6.00012H9ZM14 6.00012C13.4477 6.00012 13 6.44784 13 7.00012V8.00012C13 8.55241 13.4477 9.00012 14 9.00012H15C15.5523 9.00012 16 8.55241 16 8.00012V7.00012C16 6.44784 15.5523 6.00012 15 6.00012H14ZM9 10.0001C8.44772 10.0001 8 10.4478 8 11.0001V12.0001C8 12.5524 8.44772 13.0001 9 13.0001H10C10.5523 13.0001 11 12.5524 11 12.0001V11.0001C11 10.4478 10.5523 10.0001 10 10.0001H9ZM14 10.0001C13.4477 10.0001 13 10.4478 13 11.0001V12.0001C13 12.5524 13.4477 13.0001 14 13.0001H15C15.5523 13.0001 16 12.5524 16 12.0001V11.0001C16 10.4478 15.5523 10.0001 15 10.0001H14ZM11 14.0001C10.4696 14.0001 9.96086 14.2108 9.58579 14.5859C9.21071 14.961 9 15.4697 9 16.0001V19H11V16.0001H13V19H15V16.0001C15 15.4697 14.7893 14.961 14.4142 14.5859C14.0391 14.2108 13.5304 14.0001 13 14.0001H11Z\" fill=\"#383838\"></path></svg>")
		if templ_7745c5c3_Err != nil {
			return templ_7745c5c3_Err
		}
		return templ_7745c5c3_Err
	})
}

var _ = templruntime.GeneratedTemplate
