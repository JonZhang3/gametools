package app

import (
	"github.com/gofiber/fiber/v2"
	"github.com/valyala/fasthttp"
)

func WrapContext(ctx *fiber.Ctx) *Context {
	return &Context{
		ctx:      ctx,
		Request:  &Request{ctx: ctx},
		Response: &Response{ctx: ctx},
		Route:    &Route{ctx: ctx},
	}
}

type Context struct {
	ctx      *fiber.Ctx
	Request  *Request
	Response *Response
	Route    *Route
}

func (ctx *Context) App() *fiber.App {
	return ctx.ctx.App()
}

func (ctx *Context) BaseURL() string {
	return ctx.ctx.BaseURL()
}

func (ctx *Context) Bind(vars fiber.Map) error {
	return ctx.ctx.Bind(vars)
}

func (ctx *Context) Context() *fasthttp.RequestCtx {
	return ctx.ctx.Context()
}

func (ctx *Context) GetRouteURL(routeName string, params fiber.Map) (string, error) {
	return ctx.ctx.GetRouteURL(routeName, params)
}

func (ctx *Context) Method(override ...string) string {
	return ctx.ctx.Method(override...)
}

func (ctx *Context) Range(size int) (fiber.Range, error) {
	return ctx.ctx.Range(size)
}

func (ctx *Context) Redirect(location string, status ...int) error {
	return ctx.ctx.Redirect(location, status...)
}

func (ctx *Context) CurrentUserId() uint64 {
	claims := ParseToken(ctx.ctx)
	return claims.ID
}

type Request struct {
	ctx *fiber.Ctx
}

func (req *Request) Accepts(offers ...string) string {
	return req.ctx.Accepts(offers...)
}

func (req *Request) AcceptsCharsets(offers ...string) string {
	return req.ctx.AcceptsCharsets(offers...)
}

func (req *Request) AcceptsEncodings(offers ...string) string {
	return req.ctx.AcceptsEncodings(offers...)
}

func (req *Request) AcceptsLanguages(offers ...string) string {
	return req.ctx.AcceptsLanguages(offers...)
}

func (req *Request) Header(key string, defaultValue ...string) string {
	return req.ctx.Get(key, defaultValue...)
}

func (req *Request) Headers() map[string]string {
	return req.ctx.GetReqHeaders()
}

func (req *Request) ContentTypeIs(extension string) bool {
	return req.ctx.Is(extension)
}

func (req *Request) Hostname() string {
	return req.ctx.Hostname()
}

func (req *Request) IP() string {
	return req.ctx.IP()
}

func (req *Request) IPs() []string {
	return req.ctx.IPs()
}

func (req *Request) XHR() bool {
	return req.ctx.XHR()
}

func (req *Request) Path(override ...string) string {
	return req.ctx.Path(override...)
}

func (req *Request) Protocol() string {
	return req.ctx.Protocol()
}

func (req *Request) Body() []byte {
	return req.ctx.Body()
}

func (req *Request) BodyParser(out interface{}) error {
	return req.ctx.BodyParser(out)
}

func (req *Request) Query(key string, defaultValue ...string) string {
	return req.ctx.Query(key, defaultValue...)
}

func (req *Request) QueryParser(out any) error {
	return req.ctx.QueryParser(out)
}

type Route struct {
	ctx *fiber.Ctx
}

func (r *Route) AllParams() map[string]string {
	return r.ctx.AllParams()
}

func (r *Route) Params(key string, defaultValue ...string) string {
	return r.ctx.Params(key, defaultValue...)
}

func (r *Route) ParamsInt(key string) (int, error) {
	return r.ctx.ParamsInt(key)
}

func (r *Route) ParamsParser(out interface{}) error {
	return r.ctx.ParamsParser(out)
}

type Response struct {
	ctx *fiber.Ctx
}

func (resp *Response) HeaderAppend(field string, values ...string) {
	resp.ctx.Append(field, values...)
}

func (resp *Response) HeaderAttachment(filename ...string) {
	resp.ctx.Attachment(filename...)
}

func (resp *Response) HeaderVary(fields ...string) {
	resp.ctx.Vary(fields...)
}

func (resp *Response) ContentType(ext string, charset ...string) {
	resp.ctx.Type(ext, charset...)
}

func (resp *Response) HeaderSet(key string, val string) {
	resp.ctx.Set(key, val)
}

func (resp *Response) Headers() map[string]string {
	return resp.ctx.GetRespHeaders()
}

func (resp *Response) Links(link ...string) {
	resp.ctx.Links(link...)
}

func (resp *Response) Location(path string) {
	resp.ctx.Location(path)
}

func (resp *Response) JSON(data interface{}) error {
	return resp.ctx.JSON(data)
}

func (resp *Response) JSONP(data interface{}, callback ...string) error {
	return resp.ctx.JSONP(data, callback...)
}

func (resp *Response) XML(data interface{}) error {
	return resp.ctx.XML(data)
}

func (resp *Response) WriteString(s string) (n int, err error) {
	return resp.ctx.WriteString(s)
}

func (resp *Response) Writef(f string, a ...interface{}) (n int, err error) {
	return resp.ctx.Writef(f, a...)
}

func (resp *Response) Write(p []byte) (n int, err error) {
	return resp.ctx.Write(p)
}

func (resp *Response) Status(status int) {
	resp.ctx.Status(status)
}

func (resp *Response) Ok(data interface{}, message ...string) error {
	return resp.ctx.JSON(Ok(data, message...))
}

func (resp *Response) Error(message string, data ...interface{}) error {
	return resp.ctx.JSON(Error(message, data...))
}

type Cookie struct {
	ctx *fiber.Ctx
}

func (c *Cookie) Set(cookie *fiber.Cookie) {
	c.ctx.Cookie(cookie)
}

func (c *Cookie) Get(key string, defaultValue ...string) string {
	return c.ctx.Cookies(key, defaultValue...)
}
