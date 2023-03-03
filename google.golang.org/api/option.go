package api

import (
	"crypto/tls"
	"net/http"
)

google.golang.org/api/option

Overview ¶
包 option 包含 Google API 客户端的选项。


Types ¶
type ClientCertSource = func(*tls.CertificateRequestInfo) (*tls.Certificate, error)				// ClientCertSource 是一个函数，它返回打开 TLS 连接时要使用的 TLS 客户端证书。
																								// 它遵循与 crypto/tls.Config.GetClientCertificate 相同的语义。
																								// 这是一个实验性 API，将来可能会更改或删除。


type ClientOption interface {																	// ClientOption 是 Google API 客户端的一个选项。
	Apply(*internal.DialSettings)
}
1.func WithAPIKey(apiKey string) ClientOption								// WithAPIKey 返回一个 ClientOption，它指定要用作身份验证基础的 API 密钥。
																			// API 密钥只能用于 JSON-over-HTTP API，包括导入路径 google.golang.org/api/.... 下的 API。
2.func WithAudiences(audience ...string) ClientOption						// WithAudiences 返回一个 ClientOption，指定要用作 JWT 令牌身份验证的受众字段（“aud”）的受众。
3.func WithClientCertSource(s ClientCertSource) ClientOption				// WithClientCertSource 返回一个 ClientOption，指定用于获取 TLS 客户端证书的回调函数。
																			// 此选项用于支持 mTLS 身份验证，其中服务器在建立连接时验证客户端证书。
																			// 只要服务器向客户端请求证书，就会调用回调函数。回调函数的实现应尽量确保在传输客户端的整个生命周期内可以按需重复返回有效证书。如果返回 nil Certificate（即无法获得 Certificate），则应返回错误。
																			// 这是一个实验性 API，将来可能会更改或删除。
4.func WithCredentials(creds *google.Credentials) ClientOption				// WithCredentials 返回一个 ClientOption 来验证 API 调用。
5.func WithCredentialsFile(filename string) ClientOption					// WithCredentialsFile 返回一个 ClientOption，它使用给定的服务帐户或刷新令牌 JSON 凭据文件对 API 调用进行身份验证。
6.func WithCredentialsJSON(p []byte) ClientOption							// WithCredentialsJSON 返回一个 ClientOption，它使用给定的服务帐户或刷新令牌 JSON 凭据对 API 调用进行身份验证。
7.func WithEndpoint(url string) ClientOption								// WithEndpoint 返回一个 ClientOption，它覆盖用于服务的默认端点。
8.func WithGRPCConn(conn *grpc.ClientConn) ClientOption						// WithGRPCConn 返回一个 ClientOption，它指定要用作通信基础的 gRPC 客户端连接。此选项只能用于支持 gRPC 作为其通信传输的服务。使用时，WithGRPCConn 选项优先于所有其他提供的选项。
9.func WithGRPCConnectionPool(size int) ClientOption						// WithGRPCConnectionPool 返回一个 ClientOption，它创建一个 gRPC 连接池，请求将在这些连接之间进行平衡。
10.func WithGRPCDialOption(opt grpc.DialOption) ClientOption				// WithGRPCDialOption 返回一个 ClientOption，它将新的 grpc.DialOption 附加到底层 gRPC 拨号。它不适用于 WithGRPCConn。
11.func WithHTTPClient(client *http.Client) ClientOption					// WithHTTPClient 返回一个 ClientOption，指定要用作通信基础的 HTTP 客户端。此选项只能用于支持 HTTP 作为其通信传输的服务。使用时，WithHTTPClient 选项优先于所有其他提供的选项。
12.func WithQuotaProject(quotaProject string) ClientOption					// WithQuotaProject 返回一个 ClientOption，指定用于配额和计费目的的项目。
																			// 有关更多信息，请阅读：https://cloud.google.com/apis/docs/system-parameters
13.func WithRequestReason(requestReason string) ClientOption				// WithRequestReason 返回一个 ClientOption，它指定发出请求的原因，旨在记录在审计日志中。一个示例原因是支持案例票号。
																			// 有关更多信息，请阅读：https://cloud.google.com/apis/docs/system-parameters
14.func WithScopes(scope ...string) ClientOption							// WithScopes 返回一个 ClientOption，它覆盖用于服务的默认 OAuth2 范围。
																			// 如果同时使用 WithScopes 和 WithTokenSource，则将改用令牌源中的范围设置。
15.func WithTelemetryDisabled() ClientOption								// WithTelemetryDisabled 返回一个 ClientOption，它禁用 gRPC 和 HTTP 客户端上的默认遥测 (OpenCensus) 设置。一个示例原因是绑定覆盖默认值的自定义遥测。
16.func WithTokenSource(s oauth2.TokenSource) ClientOption					// WithTokenSource 返回一个 ClientOption，它指定要用作身份验证基础的 OAuth2 令牌源。
17.func WithUserAgent(ua string) ClientOption								// WithUserAgent 返回一个设置 User-Agent 的 ClientOption。此选项与 WithHTTPClient 选项不兼容。如果您希望提供自定义客户端，则需要通过 RoundTripper 中间件添加此标头。
18.func WithoutAuthentication() ClientOption								// WithoutAuthentication 返回一个 ClientOption，指定不应使用身份验证。它仅适用于测试和访问公共资源，例如公共 Google Cloud Storage 存储桶。
																			// 同时提供 WithoutAuthentication 和 WithAPIKey、WithTokenSource、WithCredentialsFile 或 WithServiceAccountFile 中的任何一个都是错误的。
																			