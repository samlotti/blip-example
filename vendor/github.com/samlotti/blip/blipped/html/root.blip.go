package html
// Do Not Edit
// Generated by Blip
// source blip: ./html/root.blip.html

import (
	"context"
	"fmt"
	"github.com/samlotti/blip/blipUtil"
	"io"
	"time"
)



func RootRender( pageName string, c context.Context, w io.Writer ) (terror error) {
    start := time.Now()

	var si = blipUtil.Instance()
	var escaper = si.GetEscaperFor( "html") 
	defer func() {
		if err := recover(); err != nil {
			fmt.Printf("Catch panic %s: %s\n", "RootRender", err)
			terror = fmt.Errorf("%v", err)
		}
	    si.RenderComplete(escaper, "root", "html", time.Since(start), terror)
	}()
	title := c.Value("title").(string)
	// Line: 2
	si.Write(w, []byte("<!DOCTYPE html>\n"))
	// Line: 4
	si.Write(w, []byte("\n"))
	// Line: 7
	si.Write(w, []byte("\n"))
	// Line: 13
	si.Write(w, []byte("\n<html>\n<head>\n    <meta charset=\"utf-8\">\n    <meta name=\"viewport\" content=\"width=device-width, initial-scale=1\">\n    <title>"))
	// Line: 13
	si.WriteStrSafe(w, title, escaper)
	// Line: 14
	si.Write(w, []byte("</title>\n"))
	// Line: 15
	// Text block follows
		// Line: 17
		si.Write(w, []byte("\n    <link rel=\"stylesheet\" href=\"https://cdn.jsdelivr.net/npm/bootstrap@3.3.7/dist/css/bootstrap.min.css\" integrity=\"sha384-BVYiiSIFeK1dGmJRAkycuHAHRg32OmUcww7on3RYdg4Va+PmSTsz/K68vbdEjh4u\" crossorigin=\"anonymous\">\n    <link rel=\"stylesheet\" href=\"https://cdn.jsdelivr.net/npm/bootstrap@3.3.7/dist/css/bootstrap-theme.min.css\" integrity=\"sha384-rHyoN1iRsVXV4nD0JutlnGaslCJuC7uwjduW9SVrLvRYooPp2bWYgmgJQIXwl/Sp\" crossorigin=\"anonymous\">\n"))
	// Line: 20
	si.Write(w, []byte("\n</head>\n<body>\n"))
	// Line: 21
	terror = si.CallCtxFunc(c, "styles")
	if terror != nil { return }
	// Line: 22
	si.Write(w, []byte("\n"))
	// Line: 23

	terror = NavRender(pageName, c, w)
	if terror != nil { return }
	// Line: 27
	si.Write(w, []byte("\n\n\n<div class=\"container\">\n    Page: "))
	// Line: 27
	si.WriteStrSafe(w, pageName , escaper)
	// Line: 31
	si.Write(w, []byte("\n</div>\n\n<div class=\"container\">\n    "))
	// Line: 32
	terror = si.CallCtxFunc(c, "mainContent")
	if terror != nil { return }
	// Line: 35
	si.Write(w, []byte("</div>\n\n\n"))
	// Line: 36
	// Text block follows
		// Line: 40
		si.Write(w, []byte("\n\n<!-- jQuery (necessary for Bootstrap's JavaScript plugins) -->\n<script src=\"https://code.jquery.com/jquery-1.12.4.min.js\" integrity=\"sha384-nvAa0+6Qg9clwYCGGPpDQLVpLNn0fRaROjHqs13t4Ggj3Ez50XnGQqc/r8MhnRDZ\" crossorigin=\"anonymous\"></script>\n<script src=\"https://cdn.jsdelivr.net/npm/bootstrap@3.3.7/dist/js/bootstrap.min.js\" integrity=\"sha384-Tc5IQib027qvyjSMfHjOMaLkfuWVxZxUPnCJA7l2mCWNIpG9mGCD8wGNIcPD7Txa\" crossorigin=\"anonymous\"></script>\n"))
	// Line: 48
	si.Write(w, []byte("\n</body>\n\n\n</html>\n\n\n\n"))
	return
}