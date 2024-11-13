// Code generated by templ - DO NOT EDIT.

// templ: version: v0.2.778
package icons

//lint:file-ignore SA4006 This context is only used if a nested component is present.

import "github.com/a-h/templ"
import templruntime "github.com/a-h/templ/runtime"

func EndTime() templ.Component {
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
		_, templ_7745c5c3_Err = templ_7745c5c3_Buffer.WriteString("<svg width=\"24\" height=\"24\" viewBox=\"0 0 24 24\" fill=\"none\" xmlns=\"http://www.w3.org/2000/svg\"><path fill-rule=\"evenodd\" clip-rule=\"evenodd\" d=\"M13 11.1492L13 4C13 3.44771 12.5523 3 12 3C11.4477 3 11 3.44772 11 4L11 11.1492L8.78087 8.37531C8.43586 7.94404 7.80657 7.87412 7.37531 8.21913C6.94404 8.56414 6.87412 9.19343 7.21913 9.6247L11.2191 14.6247C11.4089 14.8619 11.6962 15 12 15C12.3038 15 12.5911 14.8619 12.7809 14.6247L16.7809 9.62469C17.1259 9.19343 17.056 8.56414 16.6247 8.21913C16.1934 7.87412 15.5641 7.94404 15.2191 8.37531L13 11.1492Z\" fill=\"#383838\"></path> <path fill-rule=\"evenodd\" clip-rule=\"evenodd\" d=\"M9.65739 15.8741L7.35812 13H5C3.89543 13 3 13.8954 3 15V19C3 20.1046 3.89543 21 5 21H19C20.1046 21 21 20.1046 21 19V15C21 13.8954 20.1046 13 19 13H16.6419L14.3426 15.8741C13.7733 16.5857 12.9114 17 12 17C11.0886 17 10.2267 16.5857 9.65739 15.8741ZM17 16C16.4477 16 16 16.4477 16 17C16 17.5523 16.4477 18 17 18H17.01C17.5623 18 18.01 17.5523 18.01 17C18.01 16.4477 17.5623 16 17.01 16H17Z\" fill=\"#383838\"></path></svg>")
		if templ_7745c5c3_Err != nil {
			return templ_7745c5c3_Err
		}
		return templ_7745c5c3_Err
	})
}

var _ = templruntime.GeneratedTemplate