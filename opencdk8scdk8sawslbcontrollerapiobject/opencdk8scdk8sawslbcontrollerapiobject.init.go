package opencdk8scdk8sawslbcontrollerapiobject

import (
	"reflect"

	_jsii_ "github.com/aws/jsii-runtime-go/runtime"
)

func init() {
	_jsii_.RegisterClass(
		"@opencdk8s/cdk8s-aws-lb-controller-api-object.AWSLoadBalancerControllerObject",
		reflect.TypeOf((*AWSLoadBalancerControllerObject)(nil)).Elem(),
		[]_jsii_.Member{
			_jsii_.MemberMethod{JsiiMethod: "addDependency", GoMethod: "AddDependency"},
			_jsii_.MemberMethod{JsiiMethod: "addJsonPatch", GoMethod: "AddJsonPatch"},
			_jsii_.MemberProperty{JsiiProperty: "apiGroup", GoGetter: "ApiGroup"},
			_jsii_.MemberProperty{JsiiProperty: "apiVersion", GoGetter: "ApiVersion"},
			_jsii_.MemberProperty{JsiiProperty: "chart", GoGetter: "Chart"},
			_jsii_.MemberProperty{JsiiProperty: "kind", GoGetter: "Kind"},
			_jsii_.MemberProperty{JsiiProperty: "metadata", GoGetter: "Metadata"},
			_jsii_.MemberProperty{JsiiProperty: "name", GoGetter: "Name"},
			_jsii_.MemberProperty{JsiiProperty: "node", GoGetter: "Node"},
			_jsii_.MemberMethod{JsiiMethod: "toJson", GoMethod: "ToJson"},
			_jsii_.MemberMethod{JsiiMethod: "toString", GoMethod: "ToString"},
		},
		func() interface{} {
			j := jsiiProxy_AWSLoadBalancerControllerObject{}
			_jsii_.InitJsiiProxy(&j.Type__cdk8sApiObject)
			return &j
		},
	)
	_jsii_.RegisterStruct(
		"@opencdk8s/cdk8s-aws-lb-controller-api-object.AWSLoadBalancerControllerProps",
		reflect.TypeOf((*AWSLoadBalancerControllerProps)(nil)).Elem(),
	)
	_jsii_.RegisterStruct(
		"@opencdk8s/cdk8s-aws-lb-controller-api-object.HttpIngressPath",
		reflect.TypeOf((*HttpIngressPath)(nil)).Elem(),
	)
	_jsii_.RegisterStruct(
		"@opencdk8s/cdk8s-aws-lb-controller-api-object.HttpIngressRuleValue",
		reflect.TypeOf((*HttpIngressRuleValue)(nil)).Elem(),
	)
	_jsii_.RegisterStruct(
		"@opencdk8s/cdk8s-aws-lb-controller-api-object.IngressBackend",
		reflect.TypeOf((*IngressBackend)(nil)).Elem(),
	)
	_jsii_.RegisterStruct(
		"@opencdk8s/cdk8s-aws-lb-controller-api-object.IngressRule",
		reflect.TypeOf((*IngressRule)(nil)).Elem(),
	)
	_jsii_.RegisterStruct(
		"@opencdk8s/cdk8s-aws-lb-controller-api-object.IngressSpec",
		reflect.TypeOf((*IngressSpec)(nil)).Elem(),
	)
	_jsii_.RegisterStruct(
		"@opencdk8s/cdk8s-aws-lb-controller-api-object.IngressTls",
		reflect.TypeOf((*IngressTls)(nil)).Elem(),
	)
	_jsii_.RegisterStruct(
		"@opencdk8s/cdk8s-aws-lb-controller-api-object.Initializer",
		reflect.TypeOf((*Initializer)(nil)).Elem(),
	)
	_jsii_.RegisterStruct(
		"@opencdk8s/cdk8s-aws-lb-controller-api-object.Initializers",
		reflect.TypeOf((*Initializers)(nil)).Elem(),
	)
	_jsii_.RegisterClass(
		"@opencdk8s/cdk8s-aws-lb-controller-api-object.IntOrString",
		reflect.TypeOf((*IntOrString)(nil)).Elem(),
		nil, // no members
		func() interface{} {
			return &jsiiProxy_IntOrString{}
		},
	)
	_jsii_.RegisterStruct(
		"@opencdk8s/cdk8s-aws-lb-controller-api-object.KubeStatusProps",
		reflect.TypeOf((*KubeStatusProps)(nil)).Elem(),
	)
	_jsii_.RegisterStruct(
		"@opencdk8s/cdk8s-aws-lb-controller-api-object.ListMeta",
		reflect.TypeOf((*ListMeta)(nil)).Elem(),
	)
	_jsii_.RegisterStruct(
		"@opencdk8s/cdk8s-aws-lb-controller-api-object.ManagedFieldsEntry",
		reflect.TypeOf((*ManagedFieldsEntry)(nil)).Elem(),
	)
	_jsii_.RegisterStruct(
		"@opencdk8s/cdk8s-aws-lb-controller-api-object.ObjectMeta",
		reflect.TypeOf((*ObjectMeta)(nil)).Elem(),
	)
	_jsii_.RegisterStruct(
		"@opencdk8s/cdk8s-aws-lb-controller-api-object.OwnerReference",
		reflect.TypeOf((*OwnerReference)(nil)).Elem(),
	)
	_jsii_.RegisterStruct(
		"@opencdk8s/cdk8s-aws-lb-controller-api-object.StatusCause",
		reflect.TypeOf((*StatusCause)(nil)).Elem(),
	)
	_jsii_.RegisterStruct(
		"@opencdk8s/cdk8s-aws-lb-controller-api-object.StatusDetails",
		reflect.TypeOf((*StatusDetails)(nil)).Elem(),
	)
}
