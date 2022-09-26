/*
Copyright 2022 Upbound Inc.
*/

// Code generated by upjet. DO NOT EDIT.

package v1beta1

import (
	metav1 "k8s.io/apimachinery/pkg/apis/meta/v1"
	"k8s.io/apimachinery/pkg/runtime/schema"

	v1 "github.com/crossplane/crossplane-runtime/apis/common/v1"
)

type AccessControlAllowHeadersObservation struct {
}

type AccessControlAllowHeadersParameters struct {

	// +kubebuilder:validation:Optional
	Items []*string `json:"items,omitempty" tf:"items,omitempty"`
}

type AccessControlAllowMethodsObservation struct {
}

type AccessControlAllowMethodsParameters struct {

	// +kubebuilder:validation:Optional
	Items []*string `json:"items,omitempty" tf:"items,omitempty"`
}

type AccessControlAllowOriginsObservation struct {
}

type AccessControlAllowOriginsParameters struct {

	// +kubebuilder:validation:Optional
	Items []*string `json:"items,omitempty" tf:"items,omitempty"`
}

type AccessControlExposeHeadersObservation struct {
}

type AccessControlExposeHeadersParameters struct {

	// +kubebuilder:validation:Optional
	Items []*string `json:"items,omitempty" tf:"items,omitempty"`
}

type ContentSecurityPolicyObservation struct {
}

type ContentSecurityPolicyParameters struct {

	// The policy directives and their values that CloudFront includes as values for the Content-Security-Policy HTTP response header. See Content Security Policy for more information.
	// +kubebuilder:validation:Required
	ContentSecurityPolicy *string `json:"contentSecurityPolicy" tf:"content_security_policy,omitempty"`

	// A Boolean value that determines whether CloudFront overrides a response header with the same name received from the origin with the header specifies here.
	// +kubebuilder:validation:Required
	Override *bool `json:"override" tf:"override,omitempty"`
}

type ContentTypeOptionsObservation struct {
}

type ContentTypeOptionsParameters struct {

	// A Boolean value that determines whether CloudFront overrides a response header with the same name received from the origin with the header specifies here.
	// +kubebuilder:validation:Required
	Override *bool `json:"override" tf:"override,omitempty"`
}

type CorsConfigObservation struct {
}

type CorsConfigParameters struct {

	// A Boolean value that CloudFront uses as the value for the Access-Control-Allow-Credentials HTTP response header.
	// +kubebuilder:validation:Required
	AccessControlAllowCredentials *bool `json:"accessControlAllowCredentials" tf:"access_control_allow_credentials,omitempty"`

	// Object that contains an attribute items that contains a list of HTTP header names that CloudFront includes as values for the Access-Control-Allow-Headers HTTP response header.
	// +kubebuilder:validation:Required
	AccessControlAllowHeaders []AccessControlAllowHeadersParameters `json:"accessControlAllowHeaders" tf:"access_control_allow_headers,omitempty"`

	// Object that contains an attribute items that contains a list of HTTP methods that CloudFront includes as values for the Access-Control-Allow-Methods HTTP response header. Valid values: GET | POST | OPTIONS | PUT | DELETE | HEAD | ALL
	// +kubebuilder:validation:Required
	AccessControlAllowMethods []AccessControlAllowMethodsParameters `json:"accessControlAllowMethods" tf:"access_control_allow_methods,omitempty"`

	// Object that contains an attribute items that contains a list of origins that CloudFront can use as the value for the Access-Control-Allow-Origin HTTP response header.
	// +kubebuilder:validation:Required
	AccessControlAllowOrigins []AccessControlAllowOriginsParameters `json:"accessControlAllowOrigins" tf:"access_control_allow_origins,omitempty"`

	// Object that contains an attribute items that contains a list of HTTP headers that CloudFront includes as values for the Access-Control-Expose-Headers HTTP response header.
	// +kubebuilder:validation:Optional
	AccessControlExposeHeaders []AccessControlExposeHeadersParameters `json:"accessControlExposeHeaders,omitempty" tf:"access_control_expose_headers,omitempty"`

	// A number that CloudFront uses as the value for the Access-Control-Max-Age HTTP response header.
	// +kubebuilder:validation:Optional
	AccessControlMaxAgeSec *float64 `json:"accessControlMaxAgeSec,omitempty" tf:"access_control_max_age_sec,omitempty"`

	// A Boolean value that determines how CloudFront behaves for the HTTP response header.
	// +kubebuilder:validation:Required
	OriginOverride *bool `json:"originOverride" tf:"origin_override,omitempty"`
}

type CustomHeadersConfigItemsObservation struct {
}

type CustomHeadersConfigItemsParameters struct {

	// The HTTP response header name.
	// +kubebuilder:validation:Required
	Header *string `json:"header" tf:"header,omitempty"`

	// A Boolean value that determines whether CloudFront overrides a response header with the same name received from the origin with the header specifies here.
	// +kubebuilder:validation:Required
	Override *bool `json:"override" tf:"override,omitempty"`

	// The value for the HTTP response header.
	// +kubebuilder:validation:Required
	Value *string `json:"value" tf:"value,omitempty"`
}

type CustomHeadersConfigObservation struct {
}

type CustomHeadersConfigParameters struct {

	// +kubebuilder:validation:Optional
	Items []CustomHeadersConfigItemsParameters `json:"items,omitempty" tf:"items,omitempty"`
}

type FrameOptionsObservation struct {
}

type FrameOptionsParameters struct {

	// The value of the X-Frame-Options HTTP response header. Valid values: DENY | SAMEORIGIN
	// +kubebuilder:validation:Required
	FrameOption *string `json:"frameOption" tf:"frame_option,omitempty"`

	// A Boolean value that determines whether CloudFront overrides a response header with the same name received from the origin with the header specifies here.
	// +kubebuilder:validation:Required
	Override *bool `json:"override" tf:"override,omitempty"`
}

type ReferrerPolicyObservation struct {
}

type ReferrerPolicyParameters struct {

	// A Boolean value that determines whether CloudFront overrides a response header with the same name received from the origin with the header specifies here.
	// +kubebuilder:validation:Required
	Override *bool `json:"override" tf:"override,omitempty"`

	// Determines whether CloudFront includes the Referrer-Policy HTTP response header and the header’s value. See Referrer Policy for more information.
	// +kubebuilder:validation:Required
	ReferrerPolicy *string `json:"referrerPolicy" tf:"referrer_policy,omitempty"`
}

type ResponseHeadersPolicyObservation struct {

	// The identifier for the response headers policy.
	ID *string `json:"id,omitempty" tf:"id,omitempty"`
}

type ResponseHeadersPolicyParameters struct {

	// A comment to describe the response headers policy. The comment cannot be longer than 128 characters.
	// +kubebuilder:validation:Optional
	Comment *string `json:"comment,omitempty" tf:"comment,omitempty"`

	// A configuration for a set of HTTP response headers that are used for Cross-Origin Resource Sharing (CORS). See Cors Config for more information.
	// +kubebuilder:validation:Optional
	CorsConfig []CorsConfigParameters `json:"corsConfig,omitempty" tf:"cors_config,omitempty"`

	// Object that contains an attribute items that contains a list of custom headers. See Custom Header for more information.
	// +kubebuilder:validation:Optional
	CustomHeadersConfig []CustomHeadersConfigParameters `json:"customHeadersConfig,omitempty" tf:"custom_headers_config,omitempty"`

	// The current version of the response headers policy.
	// +kubebuilder:validation:Optional
	Etag *string `json:"etag,omitempty" tf:"etag,omitempty"`

	// A unique name to identify the response headers policy.
	// +kubebuilder:validation:Required
	Name *string `json:"name" tf:"name,omitempty"`

	// Region is the region you'd like your resource to be created in.
	// +terrajet:crd:field:TFTag=-
	// +kubebuilder:validation:Required
	Region *string `json:"region" tf:"-"`

	// A configuration for a set of security-related HTTP response headers. See Security Headers Config for more information.
	// +kubebuilder:validation:Optional
	SecurityHeadersConfig []SecurityHeadersConfigParameters `json:"securityHeadersConfig,omitempty" tf:"security_headers_config,omitempty"`
}

type SecurityHeadersConfigObservation struct {
}

type SecurityHeadersConfigParameters struct {

	// The policy directives and their values that CloudFront includes as values for the Content-Security-Policy HTTP response header. See Content Security Policy for more information.
	// +kubebuilder:validation:Optional
	ContentSecurityPolicy []ContentSecurityPolicyParameters `json:"contentSecurityPolicy,omitempty" tf:"content_security_policy,omitempty"`

	// Determines whether CloudFront includes the X-Content-Type-Options HTTP response header with its value set to nosniff. See Content Type Options for more information.
	// +kubebuilder:validation:Optional
	ContentTypeOptions []ContentTypeOptionsParameters `json:"contentTypeOptions,omitempty" tf:"content_type_options,omitempty"`

	// Determines whether CloudFront includes the X-Frame-Options HTTP response header and the header’s value. See Frame Options for more information.
	// +kubebuilder:validation:Optional
	FrameOptions []FrameOptionsParameters `json:"frameOptions,omitempty" tf:"frame_options,omitempty"`

	// Determines whether CloudFront includes the Referrer-Policy HTTP response header and the header’s value. See Referrer Policy for more information.
	// +kubebuilder:validation:Optional
	ReferrerPolicy []ReferrerPolicyParameters `json:"referrerPolicy,omitempty" tf:"referrer_policy,omitempty"`

	// Determines whether CloudFront includes the Strict-Transport-Security HTTP response header and the header’s value. See Strict Transport Security for more information.
	// +kubebuilder:validation:Optional
	StrictTransportSecurity []StrictTransportSecurityParameters `json:"strictTransportSecurity,omitempty" tf:"strict_transport_security,omitempty"`

	// Determine whether CloudFront includes the X-XSS-Protection HTTP response header and the header’s value. See XSS Protection for more information.
	// +kubebuilder:validation:Optional
	XSSProtection []XSSProtectionParameters `json:"xssProtection,omitempty" tf:"xss_protection,omitempty"`
}

type StrictTransportSecurityObservation struct {
}

type StrictTransportSecurityParameters struct {

	// A number that CloudFront uses as the value for the Access-Control-Max-Age HTTP response header.
	// +kubebuilder:validation:Required
	AccessControlMaxAgeSec *float64 `json:"accessControlMaxAgeSec" tf:"access_control_max_age_sec,omitempty"`

	// A Boolean value that determines whether CloudFront includes the includeSubDomains directive in the Strict-Transport-Security HTTP response header.
	// +kubebuilder:validation:Optional
	IncludeSubdomains *bool `json:"includeSubdomains,omitempty" tf:"include_subdomains,omitempty"`

	// A Boolean value that determines whether CloudFront overrides a response header with the same name received from the origin with the header specifies here.
	// +kubebuilder:validation:Required
	Override *bool `json:"override" tf:"override,omitempty"`

	// A Boolean value that determines whether CloudFront includes the preload directive in the Strict-Transport-Security HTTP response header.
	// +kubebuilder:validation:Optional
	Preload *bool `json:"preload,omitempty" tf:"preload,omitempty"`
}

type XSSProtectionObservation struct {
}

type XSSProtectionParameters struct {

	// A Boolean value that determines whether CloudFront includes the mode=block directive in the X-XSS-Protection header.
	// +kubebuilder:validation:Optional
	ModeBlock *bool `json:"modeBlock,omitempty" tf:"mode_block,omitempty"`

	// A Boolean value that determines whether CloudFront overrides a response header with the same name received from the origin with the header specifies here.
	// +kubebuilder:validation:Required
	Override *bool `json:"override" tf:"override,omitempty"`

	// A Boolean value that determines the value of the X-XSS-Protection HTTP response header. When this setting is true, the value of the X-XSS-Protection header is 1. When this setting is false, the value of the X-XSS-Protection header is 0.
	// +kubebuilder:validation:Required
	Protection *bool `json:"protection" tf:"protection,omitempty"`

	// A reporting URI, which CloudFront uses as the value of the report directive in the X-XSS-Protection header. You cannot specify a report_uri when mode_block is true.
	// +kubebuilder:validation:Optional
	ReportURI *string `json:"reportUri,omitempty" tf:"report_uri,omitempty"`
}

// ResponseHeadersPolicySpec defines the desired state of ResponseHeadersPolicy
type ResponseHeadersPolicySpec struct {
	v1.ResourceSpec `json:",inline"`
	ForProvider     ResponseHeadersPolicyParameters `json:"forProvider"`
}

// ResponseHeadersPolicyStatus defines the observed state of ResponseHeadersPolicy.
type ResponseHeadersPolicyStatus struct {
	v1.ResourceStatus `json:",inline"`
	AtProvider        ResponseHeadersPolicyObservation `json:"atProvider,omitempty"`
}

// +kubebuilder:object:root=true

// ResponseHeadersPolicy is the Schema for the ResponseHeadersPolicys API. Provides a CloudFront response headers policy resource.
// +kubebuilder:printcolumn:name="READY",type="string",JSONPath=".status.conditions[?(@.type=='Ready')].status"
// +kubebuilder:printcolumn:name="SYNCED",type="string",JSONPath=".status.conditions[?(@.type=='Synced')].status"
// +kubebuilder:printcolumn:name="EXTERNAL-NAME",type="string",JSONPath=".metadata.annotations.crossplane\\.io/external-name"
// +kubebuilder:printcolumn:name="AGE",type="date",JSONPath=".metadata.creationTimestamp"
// +kubebuilder:subresource:status
// +kubebuilder:resource:scope=Cluster,categories={crossplane,managed,aws}
type ResponseHeadersPolicy struct {
	metav1.TypeMeta   `json:",inline"`
	metav1.ObjectMeta `json:"metadata,omitempty"`
	Spec              ResponseHeadersPolicySpec   `json:"spec"`
	Status            ResponseHeadersPolicyStatus `json:"status,omitempty"`
}

// +kubebuilder:object:root=true

// ResponseHeadersPolicyList contains a list of ResponseHeadersPolicys
type ResponseHeadersPolicyList struct {
	metav1.TypeMeta `json:",inline"`
	metav1.ListMeta `json:"metadata,omitempty"`
	Items           []ResponseHeadersPolicy `json:"items"`
}

// Repository type metadata.
var (
	ResponseHeadersPolicy_Kind             = "ResponseHeadersPolicy"
	ResponseHeadersPolicy_GroupKind        = schema.GroupKind{Group: CRDGroup, Kind: ResponseHeadersPolicy_Kind}.String()
	ResponseHeadersPolicy_KindAPIVersion   = ResponseHeadersPolicy_Kind + "." + CRDGroupVersion.String()
	ResponseHeadersPolicy_GroupVersionKind = CRDGroupVersion.WithKind(ResponseHeadersPolicy_Kind)
)

func init() {
	SchemeBuilder.Register(&ResponseHeadersPolicy{}, &ResponseHeadersPolicyList{})
}
