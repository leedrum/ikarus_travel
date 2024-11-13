// Code generated by templ - DO NOT EDIT.

// templ: version: v0.2.778
package icons

//lint:file-ignore SA4006 This context is only used if a nested component is present.

import "github.com/a-h/templ"
import templruntime "github.com/a-h/templ/runtime"

func Edit() templ.Component {
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
		_, templ_7745c5c3_Err = templ_7745c5c3_Buffer.WriteString("<svg width=\"24\" height=\"24\" viewBox=\"0 0 24 24\" fill=\"none\" xmlns=\"http://www.w3.org/2000/svg\"><path fill-rule=\"evenodd\" clip-rule=\"evenodd\" d=\"M11.3196 6.17647H5.00001C3.89543 6.17647 3 7.12457 3 8.29412V18.8824C3 20.0519 3.89543 21 5.00001 21H16.0001C17.1046 21 18.0001 20.0519 18.0001 18.8824V11.1322L14.0863 15.2762C13.7373 15.6457 13.2928 15.8976 12.8088 16.0001L10.1277 16.5678C8.37839 16.9383 6.83608 15.3053 7.18594 13.4531L7.72217 10.6142C7.81896 10.1018 8.05685 9.63114 8.40585 9.26162L11.3196 6.17647Z\" fill=\"#383838\"></path> <path fill-rule=\"evenodd\" clip-rule=\"evenodd\" d=\"M19.8464 4.31839C19.745 4.05932 19.5964 3.82394 19.4091 3.62571C19.2218 3.42734 18.9995 3.26998 18.7549 3.16262C18.5102 3.05526 18.2479 3 17.983 3C17.7182 3 17.4559 3.05526 17.2112 3.16262C16.9666 3.26998 16.7443 3.42734 16.557 3.62571L16.0112 4.2037L18.8632 7.22346L19.4091 6.64547C19.5964 6.44724 19.745 6.21186 19.8464 5.95279C19.9478 5.69372 20 5.41603 20 5.13559C20 4.85515 19.9478 4.57746 19.8464 4.31839ZM17.449 8.72087L14.5969 5.70111L9.82006 10.759C9.75026 10.8329 9.70268 10.927 9.68332 11.0295L9.1471 13.8684C9.07713 14.2388 9.38559 14.5654 9.73545 14.4913L12.4166 13.9236C12.5134 13.9031 12.6023 13.8527 12.6721 13.7788L17.449 8.72087Z\" fill=\"#383838\"></path></svg>")
		if templ_7745c5c3_Err != nil {
			return templ_7745c5c3_Err
		}
		return templ_7745c5c3_Err
	})
}

func EditWithLink(link string) templ.Component {
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
		templ_7745c5c3_Var2 := templ.GetChildren(ctx)
		if templ_7745c5c3_Var2 == nil {
			templ_7745c5c3_Var2 = templ.NopComponent
		}
		ctx = templ.ClearChildren(ctx)
		_, templ_7745c5c3_Err = templ_7745c5c3_Buffer.WriteString("<a href=\"")
		if templ_7745c5c3_Err != nil {
			return templ_7745c5c3_Err
		}
		var templ_7745c5c3_Var3 templ.SafeURL = templ.URL(link)
		_, templ_7745c5c3_Err = templ_7745c5c3_Buffer.WriteString(templ.EscapeString(string(templ_7745c5c3_Var3)))
		if templ_7745c5c3_Err != nil {
			return templ_7745c5c3_Err
		}
		_, templ_7745c5c3_Err = templ_7745c5c3_Buffer.WriteString("\">")
		if templ_7745c5c3_Err != nil {
			return templ_7745c5c3_Err
		}
		templ_7745c5c3_Err = Edit().Render(ctx, templ_7745c5c3_Buffer)
		if templ_7745c5c3_Err != nil {
			return templ_7745c5c3_Err
		}
		_, templ_7745c5c3_Err = templ_7745c5c3_Buffer.WriteString("</a>")
		if templ_7745c5c3_Err != nil {
			return templ_7745c5c3_Err
		}
		return templ_7745c5c3_Err
	})
}

var _ = templruntime.GeneratedTemplate