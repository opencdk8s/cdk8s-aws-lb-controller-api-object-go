// @opencdk8s/cdk8s-aws-lb-controller-api-object
package opencdk8scdk8sawslbcontrollerapiobject

import (
	"time"

	_jsii_ "github.com/aws/jsii-runtime-go/runtime"
	_init_ "github.com/opencdk8s/cdk8s-aws-lb-controller-api-object-go/opencdk8scdk8sawslbcontrollerapiobject/jsii"

	"github.com/aws/constructs-go/constructs/v3"
	"github.com/cdk8s-team/cdk8s-go/cdk8s"
	"github.com/opencdk8s/cdk8s-aws-lb-controller-api-object-go/opencdk8scdk8sawslbcontrollerapiobject/internal"
)

// Experimental.
type AWSLoadBalancerControllerObject interface {
	cdk8s.ApiObject
	ApiGroup() *string
	ApiVersion() *string
	Chart() cdk8s.Chart
	Kind() *string
	Metadata() cdk8s.ApiObjectMetadataDefinition
	Name() *string
	AddDependency(dependencies ...constructs.IConstruct)
	AddJsonPatch(ops ...cdk8s.JsonPatch)
	OnPrepare()
	OnSynthesize(session constructs.ISynthesisSession)
	OnValidate() *[]*string
	ToJson() interface{}
	ToString() *string
}

// The jsii proxy struct for AWSLoadBalancerControllerObject
type jsiiProxy_AWSLoadBalancerControllerObject struct {
	internal.Type__cdk8sApiObject
}

func (j *jsiiProxy_AWSLoadBalancerControllerObject) ApiGroup() *string {
	var returns *string
	_jsii_.Get(
		j,
		"apiGroup",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_AWSLoadBalancerControllerObject) ApiVersion() *string {
	var returns *string
	_jsii_.Get(
		j,
		"apiVersion",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_AWSLoadBalancerControllerObject) Chart() cdk8s.Chart {
	var returns cdk8s.Chart
	_jsii_.Get(
		j,
		"chart",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_AWSLoadBalancerControllerObject) Kind() *string {
	var returns *string
	_jsii_.Get(
		j,
		"kind",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_AWSLoadBalancerControllerObject) Metadata() cdk8s.ApiObjectMetadataDefinition {
	var returns cdk8s.ApiObjectMetadataDefinition
	_jsii_.Get(
		j,
		"metadata",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_AWSLoadBalancerControllerObject) Name() *string {
	var returns *string
	_jsii_.Get(
		j,
		"name",
		&returns,
	)
	return returns
}


// Defines an "extentions" API object for AWS Load Balancer Controller - https://github.com/kubernetes-sigs/aws-load-balancer-controller.
// Experimental.
func NewAWSLoadBalancerControllerObject(scope constructs.Construct, id *string, props *AWSLoadBalancerControllerProps) AWSLoadBalancerControllerObject {
	_init_.Initialize()

	j := jsiiProxy_AWSLoadBalancerControllerObject{}

	_jsii_.Create(
		"@opencdk8s/cdk8s-aws-lb-controller-api-object.AWSLoadBalancerControllerObject",
		[]interface{}{scope, id, props},
		&j,
	)

	return &j
}

// Defines an "extentions" API object for AWS Load Balancer Controller - https://github.com/kubernetes-sigs/aws-load-balancer-controller.
// Experimental.
func NewAWSLoadBalancerControllerObject_Override(a AWSLoadBalancerControllerObject, scope constructs.Construct, id *string, props *AWSLoadBalancerControllerProps) {
	_init_.Initialize()

	_jsii_.Create(
		"@opencdk8s/cdk8s-aws-lb-controller-api-object.AWSLoadBalancerControllerObject",
		[]interface{}{scope, id, props},
		a,
	)
}

// Renders a Kubernetes manifest for an ingress object. https://github.com/kubernetes-sigs/aws-load-balancer-controller.
//
// This can be used to inline resource manifests inside other objects (e.g. as templates).
// Experimental.
func AWSLoadBalancerControllerObject_Manifest(props *AWSLoadBalancerControllerProps) interface{} {
	_init_.Initialize()

	var returns interface{}

	_jsii_.StaticInvoke(
		"@opencdk8s/cdk8s-aws-lb-controller-api-object.AWSLoadBalancerControllerObject",
		"manifest",
		[]interface{}{props},
		&returns,
	)

	return returns
}

// Returns the `ApiObject` named `Resource` which is a child of the given construct.
//
// If `c` is an `ApiObject`, it is returned directly. Throws an
// exception if the construct does not have a child named `Default` _or_ if
// this child is not an `ApiObject`.
// Experimental.
func AWSLoadBalancerControllerObject_Of(c constructs.IConstruct) cdk8s.ApiObject {
	_init_.Initialize()

	var returns cdk8s.ApiObject

	_jsii_.StaticInvoke(
		"@opencdk8s/cdk8s-aws-lb-controller-api-object.AWSLoadBalancerControllerObject",
		"of",
		[]interface{}{c},
		&returns,
	)

	return returns
}

func AWSLoadBalancerControllerObject_GVK() *cdk8s.GroupVersionKind {
	_init_.Initialize()
	var returns *cdk8s.GroupVersionKind
	_jsii_.StaticGet(
		"@opencdk8s/cdk8s-aws-lb-controller-api-object.AWSLoadBalancerControllerObject",
		"GVK",
		&returns,
	)
	return returns
}

// Create a dependency between this ApiObject and other constructs.
//
// These can be other ApiObjects, Charts, or custom.
// Experimental.
func (a *jsiiProxy_AWSLoadBalancerControllerObject) AddDependency(dependencies ...constructs.IConstruct) {
	args := []interface{}{}
	for _, a := range dependencies {
		args = append(args, a)
	}

	_jsii_.InvokeVoid(
		a,
		"addDependency",
		args,
	)
}

// Applies a set of RFC-6902 JSON-Patch operations to the manifest synthesized for this API object.
//
// TODO: EXAMPLE
//
// Experimental.
func (a *jsiiProxy_AWSLoadBalancerControllerObject) AddJsonPatch(ops ...cdk8s.JsonPatch) {
	args := []interface{}{}
	for _, a := range ops {
		args = append(args, a)
	}

	_jsii_.InvokeVoid(
		a,
		"addJsonPatch",
		args,
	)
}

// Perform final modifications before synthesis.
//
// This method can be implemented by derived constructs in order to perform
// final changes before synthesis. prepare() will be called after child
// constructs have been prepared.
//
// This is an advanced framework feature. Only use this if you
// understand the implications.
// Experimental.
func (a *jsiiProxy_AWSLoadBalancerControllerObject) OnPrepare() {
	_jsii_.InvokeVoid(
		a,
		"onPrepare",
		nil, // no parameters
	)
}

// Allows this construct to emit artifacts into the cloud assembly during synthesis.
//
// This method is usually implemented by framework-level constructs such as `Stack` and `Asset`
// as they participate in synthesizing the cloud assembly.
// Experimental.
func (a *jsiiProxy_AWSLoadBalancerControllerObject) OnSynthesize(session constructs.ISynthesisSession) {
	_jsii_.InvokeVoid(
		a,
		"onSynthesize",
		[]interface{}{session},
	)
}

// Validate the current construct.
//
// This method can be implemented by derived constructs in order to perform
// validation logic. It is called on all constructs before synthesis.
//
// Returns: An array of validation error messages, or an empty array if there the construct is valid.
// Deprecated: use `Node.addValidation()` to subscribe validation functions on this construct
// instead of overriding this method.
func (a *jsiiProxy_AWSLoadBalancerControllerObject) OnValidate() *[]*string {
	var returns *[]*string

	_jsii_.Invoke(
		a,
		"onValidate",
		nil, // no parameters
		&returns,
	)

	return returns
}

// Renders the object to Kubernetes JSON.
//
// To disable sorting of dictionary keys in output object set the
// `CDK8S_DISABLE_SORT` environment variable to any non-empty value.
// Experimental.
func (a *jsiiProxy_AWSLoadBalancerControllerObject) ToJson() interface{} {
	var returns interface{}

	_jsii_.Invoke(
		a,
		"toJson",
		nil, // no parameters
		&returns,
	)

	return returns
}

// Returns a string representation of this construct.
// Experimental.
func (a *jsiiProxy_AWSLoadBalancerControllerObject) ToString() *string {
	var returns *string

	_jsii_.Invoke(
		a,
		"toString",
		nil, // no parameters
		&returns,
	)

	return returns
}

// Experimental.
type AWSLoadBalancerControllerProps struct {
	// Standard object's metadata.
	// Experimental.
	Metadata *ObjectMeta `json:"metadata"`
	// Spec defines the behavior of the ingress.
	// Experimental.
	Spec *IngressSpec `json:"spec"`
}

// HTTPIngressPath associates a path regex with a backend.
//
// Incoming urls matching the path are forwarded to the backend.
// Experimental.
type HttpIngressPath struct {
	// Backend defines the referenced service endpoint to which the traffic will be forwarded to.
	// Experimental.
	Backend *IngressBackend `json:"backend"`
	// Path is an extended POSIX regex as defined by IEEE Std 1003.1, (i.e this follows the egrep/unix syntax, not the perl syntax) matched against the path of an incoming request. Currently it can contain characters disallowed from the conventional "path" part of a URL as defined by RFC 3986. Paths must begin with a '/'. If unspecified, the path defaults to a catch all sending traffic to the backend.
	// Experimental.
	Path *string `json:"path"`
}

// HTTPIngressRuleValue is a list of http selectors pointing to backends.
//
// In the example: http://<host>/<path>?<searchpart> -> backend where where parts of the url correspond to RFC 3986, this resource will be used to match against everything after the last '/' and before the first '?' or '#'.
// Experimental.
type HttpIngressRuleValue struct {
	// A collection of paths that map requests to backends.
	// Experimental.
	Paths *[]*HttpIngressPath `json:"paths"`
}

// IngressBackend describes all endpoints for a given service and port.
// Experimental.
type IngressBackend struct {
	// Specifies the name of the referenced service.
	// Experimental.
	ServiceName *string `json:"serviceName"`
	// Specifies the port of the referenced service.
	// Experimental.
	ServicePort IntOrString `json:"servicePort"`
}

// IngressRule represents the rules mapping the paths under a specified host to the related backend services.
//
// Incoming requests are first evaluated for a host match, then routed to the backend associated with the matching IngressRuleValue.
// Experimental.
type IngressRule struct {
	// Host is the fully qualified domain name of a network host, as defined by RFC 3986.
	//
	// Note the following deviations from the "host" part of the URI as defined in the RFC: 1. IPs are not allowed. Currently an IngressRuleValue can only apply to the
	// IP in the Spec of the parent Ingress.
	// 2. The `:` delimiter is not respected because ports are not allowed.
	// Currently the port of an Ingress is implicitly :80 for http and
	// :443 for https.
	// Both these may change in the future. Incoming requests are matched against the host before the IngressRuleValue. If the host is unspecified, the Ingress routes all traffic based on the specified IngressRuleValue.
	// Experimental.
	Host *string `json:"host"`
	// Experimental.
	Http *HttpIngressRuleValue `json:"http"`
}

// IngressSpec describes the Ingress the user wishes to exist.
// Experimental.
type IngressSpec struct {
	// A default backend capable of servicing requests that don't match any rule.
	//
	// At least one of 'backend' or 'rules' must be specified. This field is optional to allow the loadbalancer controller or defaulting logic to specify a global default.
	// Experimental.
	Backend *IngressBackend `json:"backend"`
	// A list of host rules used to configure the Ingress.
	//
	// If unspecified, or no rule matches, all traffic is sent to the default backend.
	// Experimental.
	Rules *[]*IngressRule `json:"rules"`
	// TLS configuration.
	//
	// Currently the Ingress only supports a single TLS port, 443. If multiple members of this list specify different hosts, they will be multiplexed on the same port according to the hostname specified through the SNI TLS extension, if the ingress controller fulfilling the ingress supports SNI.
	// Experimental.
	Tls *[]*IngressTls `json:"tls"`
}

// IngressTLS describes the transport layer security associated with an Ingress.
// Experimental.
type IngressTls struct {
	// Hosts are a list of hosts included in the TLS certificate.
	//
	// The values in this list must match the name/s used in the tlsSecret. Defaults to the wildcard host setting for the loadbalancer controller fulfilling this Ingress, if left unspecified.
	// Experimental.
	Hosts *[]*string `json:"hosts"`
	// SecretName is the name of the secret used to terminate SSL traffic on 443.
	//
	// Field is left optional to allow SSL routing based on SNI hostname alone. If the SNI host in a listener conflicts with the "Host" header field used by an IngressRule, the SNI host is used for termination and value of the Host header is used for routing.
	// Experimental.
	SecretName *string `json:"secretName"`
}

// Initializer is information about an initializer that has not yet completed.
// Experimental.
type Initializer struct {
	// name of the process that is responsible for initializing this object.
	// Experimental.
	Name *string `json:"name"`
}

// Initializers tracks the progress of initialization.
// Experimental.
type Initializers struct {
	// Pending is a list of initializers that must execute in order before this object is visible.
	//
	// When the last pending initializer is removed, and no failing result is set, the initializers struct will be set to nil and the object is considered as initialized and visible to all clients.
	// Experimental.
	Pending *[]*Initializer `json:"pending"`
	// If result is set with the Failure field, the object will be persisted to storage and then deleted, ensuring that other clients can observe the deletion.
	// Experimental.
	Result *KubeStatusProps `json:"result"`
}

// Experimental.
type IntOrString interface {
}

// The jsii proxy struct for IntOrString
type jsiiProxy_IntOrString struct {
	_ byte // padding
}

// Experimental.
func IntOrString_FromNumber(value *float64) IntOrString {
	_init_.Initialize()

	var returns IntOrString

	_jsii_.StaticInvoke(
		"@opencdk8s/cdk8s-aws-lb-controller-api-object.IntOrString",
		"fromNumber",
		[]interface{}{value},
		&returns,
	)

	return returns
}

// Experimental.
func IntOrString_FromString(value *string) IntOrString {
	_init_.Initialize()

	var returns IntOrString

	_jsii_.StaticInvoke(
		"@opencdk8s/cdk8s-aws-lb-controller-api-object.IntOrString",
		"fromString",
		[]interface{}{value},
		&returns,
	)

	return returns
}

// Status is a return value for calls that don't return other objects.
// Experimental.
type KubeStatusProps struct {
	// Suggested HTTP return code for this status, 0 if not set.
	// Experimental.
	Code *float64 `json:"code"`
	// Extended data associated with the reason.
	//
	// Each reason may define its own extended details. This field is optional and the data returned is not guaranteed to conform to any schema except that defined by the reason type.
	// Experimental.
	Details *StatusDetails `json:"details"`
	// A human-readable description of the status of this operation.
	// Experimental.
	Message *string `json:"message"`
	// Standard list metadata.
	//
	// More info: https://git.k8s.io/community/contributors/devel/api-conventions.md#types-kinds
	// Experimental.
	Metadata *ListMeta `json:"metadata"`
	// A machine-readable description of why this operation is in the "Failure" status.
	//
	// If this value is empty there is no information available. A Reason clarifies an HTTP status code but does not override it.
	// Experimental.
	Reason *string `json:"reason"`
}

// ListMeta describes metadata that synthetic resources must have, including lists and various status objects.
//
// A resource may have only one of {ObjectMeta, ListMeta}.
// Experimental.
type ListMeta struct {
	// continue may be set if the user set a limit on the number of items returned, and indicates that the server has more data available.
	//
	// The value is opaque and may be used to issue another request to the endpoint that served this list to retrieve the next set of available objects. Continuing a consistent list may not be possible if the server configuration has changed or more than a few minutes have passed. The resourceVersion field returned when using this continue value will be identical to the value in the first response, unless you have received this token from an error message.
	// Experimental.
	Continue *string `json:"continue"`
	// remainingItemCount is the number of subsequent items in the list which are not included in this list response.
	//
	// If the list request contained label or field selectors, then the number of remaining items is unknown and the field will be left unset and omitted during serialization. If the list is complete (either because it is not chunking or because this is the last chunk), then there are no more remaining items and this field will be left unset and omitted during serialization. Servers older than v1.15 do not set this field. The intended use of the remainingItemCount is *estimating* the size of a collection. Clients should not rely on the remainingItemCount to be set or to be exact.
	//
	// This field is alpha and can be changed or removed without notice.
	// Experimental.
	RemainingItemCount *float64 `json:"remainingItemCount"`
	// String that identifies the server's internal version of this object that can be used by clients to determine when objects have changed.
	//
	// Value must be treated as opaque by clients and passed unmodified back to the server. Populated by the system. Read-only. More info: https://git.k8s.io/community/contributors/devel/api-conventions.md#concurrency-control-and-consistency
	// Experimental.
	ResourceVersion *string `json:"resourceVersion"`
	// selfLink is a URL representing this object.
	//
	// Populated by the system. Read-only.
	// Experimental.
	SelfLink *string `json:"selfLink"`
}

// ManagedFieldsEntry is a workflow-id, a FieldSet and the group version of the resource that the fieldset applies to.
// Experimental.
type ManagedFieldsEntry struct {
	// APIVersion defines the version of this resource that this field set applies to.
	//
	// The format is "group/version" just like the top-level APIVersion field. It is necessary to track the version of a field set because it cannot be automatically converted.
	// Experimental.
	ApiVersion *string `json:"apiVersion"`
	// Fields identifies a set of fields.
	// Experimental.
	Fields interface{} `json:"fields"`
	// Manager is an identifier of the workflow managing these fields.
	// Experimental.
	Manager *string `json:"manager"`
	// Operation is the type of operation which lead to this ManagedFieldsEntry being created.
	//
	// The only valid values for this field are 'Apply' and 'Update'.
	// Experimental.
	Operation *string `json:"operation"`
	// Time is timestamp of when these fields were set.
	//
	// It should always be empty if Operation is 'Apply'
	// Experimental.
	Time *time.Time `json:"time"`
}

// ObjectMeta is metadata that all persisted resources must have, which includes all objects users must create.
// Experimental.
type ObjectMeta struct {
	// Annotations is an unstructured key value map stored with a resource that may be set by external tools to store and retrieve arbitrary metadata.
	//
	// They are not queryable and should be preserved when modifying objects. More info: http://kubernetes.io/docs/user-guide/annotations
	// Experimental.
	Annotations *map[string]*string `json:"annotations"`
	// The name of the cluster which the object belongs to.
	//
	// This is used to distinguish resources with same name and namespace in different clusters. This field is not set anywhere right now and apiserver is going to ignore it if set in create or update request.
	// Experimental.
	ClusterName *string `json:"clusterName"`
	// CreationTimestamp is a timestamp representing the server time when this object was created.
	//
	// It is not guaranteed to be set in happens-before order across separate operations. Clients may not set this value. It is represented in RFC3339 form and is in UTC.
	//
	// Populated by the system. Read-only. Null for lists. More info: https://git.k8s.io/community/contributors/devel/api-conventions.md#metadata
	// Experimental.
	CreationTimestamp *time.Time `json:"creationTimestamp"`
	// Number of seconds allowed for this object to gracefully terminate before it will be removed from the system.
	//
	// Only set when deletionTimestamp is also set. May only be shortened. Read-only.
	// Experimental.
	DeletionGracePeriodSeconds *float64 `json:"deletionGracePeriodSeconds"`
	// DeletionTimestamp is RFC 3339 date and time at which this resource will be deleted.
	//
	// This field is set by the server when a graceful deletion is requested by the user, and is not directly settable by a client. The resource is expected to be deleted (no longer visible from resource lists, and not reachable by name) after the time in this field, once the finalizers list is empty. As long as the finalizers list contains items, deletion is blocked. Once the deletionTimestamp is set, this value may not be unset or be set further into the future, although it may be shortened or the resource may be deleted prior to this time. For example, a user may request that a pod is deleted in 30 seconds. The Kubelet will react by sending a graceful termination signal to the containers in the pod. After that 30 seconds, the Kubelet will send a hard termination signal (SIGKILL) to the container and after cleanup, remove the pod from the API. In the presence of network partitions, this object may still exist after this timestamp, until an administrator or automated process can determine the resource is fully terminated. If not set, graceful deletion of the object has not been requested.
	//
	// Populated by the system when a graceful deletion is requested. Read-only. More info: https://git.k8s.io/community/contributors/devel/api-conventions.md#metadata
	// Experimental.
	DeletionTimestamp *time.Time `json:"deletionTimestamp"`
	// Must be empty before the object is deleted from the registry.
	//
	// Each entry is an identifier for the responsible component that will remove the entry from the list. If the deletionTimestamp of the object is non-nil, entries in this list can only be removed.
	// Experimental.
	Finalizers *[]*string `json:"finalizers"`
	// GenerateName is an optional prefix, used by the server, to generate a unique name ONLY IF the Name field has not been provided.
	//
	// If this field is used, the name returned to the client will be different than the name passed. This value will also be combined with a unique suffix. The provided value has the same validation rules as the Name field, and may be truncated by the length of the suffix required to make the value unique on the server.
	//
	// If this field is specified and the generated name exists, the server will NOT return a 409 - instead, it will either return 201 Created or 500 with Reason ServerTimeout indicating a unique name could not be found in the time allotted, and the client should retry (optionally after the time indicated in the Retry-After header).
	//
	// Applied only if Name is not specified. More info: https://git.k8s.io/community/contributors/devel/api-conventions.md#idempotency
	// Experimental.
	GenerateName *string `json:"generateName"`
	// A sequence number representing a specific generation of the desired state.
	//
	// Populated by the system. Read-only.
	// Experimental.
	Generation *float64 `json:"generation"`
	// An initializer is a controller which enforces some system invariant at object creation time.
	//
	// This field is a list of initializers that have not yet acted on this object. If nil or empty, this object has been completely initialized. Otherwise, the object is considered uninitialized and is hidden (in list/watch and get calls) from clients that haven't explicitly asked to observe uninitialized objects.
	//
	// When an object is created, the system will populate this list with the current set of initializers. Only privileged users may set or modify this list. Once it is empty, it may not be modified further by any user.
	//
	// DEPRECATED - initializers are an alpha field and will be removed in v1.15.
	// Experimental.
	Initializers *Initializers `json:"initializers"`
	// Map of string keys and values that can be used to organize and categorize (scope and select) objects.
	//
	// May match selectors of replication controllers and services. More info: http://kubernetes.io/docs/user-guide/labels
	// Experimental.
	Labels *map[string]*string `json:"labels"`
	// ManagedFields maps workflow-id and version to the set of fields that are managed by that workflow.
	//
	// This is mostly for internal housekeeping, and users typically shouldn't need to set or understand this field. A workflow can be the user's name, a controller's name, or the name of a specific apply path like "ci-cd". The set of fields is always in the version that the workflow used when modifying the object.
	//
	// This field is alpha and can be changed or removed without notice.
	// Experimental.
	ManagedFields *[]*ManagedFieldsEntry `json:"managedFields"`
	// Name must be unique within a namespace.
	//
	// Is required when creating resources, although some resources may allow a client to request the generation of an appropriate name automatically. Name is primarily intended for creation idempotence and configuration definition. Cannot be updated. More info: http://kubernetes.io/docs/user-guide/identifiers#names
	// Experimental.
	Name *string `json:"name"`
	// Namespace defines the space within each name must be unique.
	//
	// An empty namespace is equivalent to the "default" namespace, but "default" is the canonical representation. Not all objects are required to be scoped to a namespace - the value of this field for those objects will be empty.
	//
	// Must be a DNS_LABEL. Cannot be updated. More info: http://kubernetes.io/docs/user-guide/namespaces
	// Experimental.
	Namespace *string `json:"namespace"`
	// List of objects depended by this object.
	//
	// If ALL objects in the list have been deleted, this object will be garbage collected. If this object is managed by a controller, then an entry in this list will point to this controller, with the controller field set to true. There cannot be more than one managing controller.
	// Experimental.
	OwnerReferences *[]*OwnerReference `json:"ownerReferences"`
	// An opaque value that represents the internal version of this object that can be used by clients to determine when objects have changed.
	//
	// May be used for optimistic concurrency, change detection, and the watch operation on a resource or set of resources. Clients must treat these values as opaque and passed unmodified back to the server. They may only be valid for a particular resource or set of resources.
	//
	// Populated by the system. Read-only. Value must be treated as opaque by clients and . More info: https://git.k8s.io/community/contributors/devel/api-conventions.md#concurrency-control-and-consistency
	// Experimental.
	ResourceVersion *string `json:"resourceVersion"`
	// SelfLink is a URL representing this object.
	//
	// Populated by the system. Read-only.
	// Experimental.
	SelfLink *string `json:"selfLink"`
	// UID is the unique in time and space value for this object.
	//
	// It is typically generated by the server on successful creation of a resource and is not allowed to change on PUT operations.
	//
	// Populated by the system. Read-only. More info: http://kubernetes.io/docs/user-guide/identifiers#uids
	// Experimental.
	Uid *string `json:"uid"`
}

// OwnerReference contains enough information to let you identify an owning object.
//
// An owning object must be in the same namespace as the dependent, or be cluster-scoped, so there is no namespace field.
// Experimental.
type OwnerReference struct {
	// API version of the referent.
	// Experimental.
	ApiVersion *string `json:"apiVersion"`
	// Kind of the referent.
	//
	// More info: https://git.k8s.io/community/contributors/devel/api-conventions.md#types-kinds
	// Experimental.
	Kind *string `json:"kind"`
	// Name of the referent.
	//
	// More info: http://kubernetes.io/docs/user-guide/identifiers#names
	// Experimental.
	Name *string `json:"name"`
	// UID of the referent.
	//
	// More info: http://kubernetes.io/docs/user-guide/identifiers#uids
	// Experimental.
	Uid *string `json:"uid"`
	// If true, AND if the owner has the "foregroundDeletion" finalizer, then the owner cannot be deleted from the key-value store until this reference is removed.
	//
	// Defaults to false. To set this field, a user needs "delete" permission of the owner, otherwise 422 (Unprocessable Entity) will be returned.
	// Experimental.
	BlockOwnerDeletion *bool `json:"blockOwnerDeletion"`
	// If true, this reference points to the managing controller.
	// Experimental.
	Controller *bool `json:"controller"`
}

// StatusCause provides more information about an api.Status failure, including cases when multiple errors are encountered.
// Experimental.
type StatusCause struct {
	// The field of the resource that has caused this error, as named by its JSON serialization.
	//
	// May include dot and postfix notation for nested attributes. Arrays are zero-indexed.  Fields may appear more than once in an array of causes due to fields having multiple errors. Optional.
	//
	// Examples:
	// "name" - the field "name" on the current resource
	// "items[0].name" - the field "name" on the first array entry in "items"
	// Experimental.
	Field *string `json:"field"`
	// A human-readable description of the cause of the error.
	//
	// This field may be presented as-is to a reader.
	// Experimental.
	Message *string `json:"message"`
	// A machine-readable description of the cause of the error.
	//
	// If this value is empty there is no information available.
	// Experimental.
	Reason *string `json:"reason"`
}

// StatusDetails is a set of additional properties that MAY be set by the server to provide additional information about a response.
//
// The Reason field of a Status object defines what attributes will be set. Clients must ignore fields that do not match the defined type of each attribute, and should assume that any attribute may be empty, invalid, or under defined.
// Experimental.
type StatusDetails struct {
	// The Causes array includes more details associated with the StatusReason failure.
	//
	// Not all StatusReasons may provide detailed causes.
	// Experimental.
	Causes *[]*StatusCause `json:"causes"`
	// The group attribute of the resource associated with the status StatusReason.
	// Experimental.
	Group *string `json:"group"`
	// The kind attribute of the resource associated with the status StatusReason.
	//
	// On some operations may differ from the requested resource Kind. More info: https://git.k8s.io/community/contributors/devel/api-conventions.md#types-kinds
	// Experimental.
	Kind *string `json:"kind"`
	// The name attribute of the resource associated with the status StatusReason (when there is a single name which can be described).
	// Experimental.
	Name *string `json:"name"`
	// If specified, the time in seconds before the operation should be retried.
	//
	// Some errors may indicate the client must take an alternate action - for those errors this field may indicate how long to wait before taking the alternate action.
	// Experimental.
	RetryAfterSeconds *float64 `json:"retryAfterSeconds"`
	// UID of the resource.
	//
	// (when there is a single resource which can be described). More info: http://kubernetes.io/docs/user-guide/identifiers#uids
	// Experimental.
	Uid *string `json:"uid"`
}

