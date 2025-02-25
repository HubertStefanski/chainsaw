package template

import (
	"context"

	"github.com/jmespath-community/go-jmespath/pkg/binding"
	"github.com/kyverno/chainsaw/pkg/apis/v1alpha1"
	"github.com/kyverno/chainsaw/pkg/runner/mutate"
	"k8s.io/apimachinery/pkg/apis/meta/v1/unstructured"
)

func ResourceRef(ctx context.Context, obj *unstructured.Unstructured, bindings binding.Bindings) error {
	var temp unstructured.Unstructured
	temp.SetAPIVersion(obj.GetAPIVersion())
	temp.SetKind(obj.GetKind())
	temp.SetName(obj.GetName())
	temp.SetNamespace(obj.GetNamespace())
	temp.SetLabels(obj.GetLabels())
	template := v1alpha1.Any{
		Value: temp.UnstructuredContent(),
	}
	if merged, err := mutate.Merge(ctx, temp, bindings, template); err != nil {
		return err
	} else {
		temp = merged
	}
	obj.SetAPIVersion(temp.GetAPIVersion())
	obj.SetKind(temp.GetKind())
	obj.SetName(temp.GetName())
	obj.SetNamespace(temp.GetNamespace())
	obj.SetLabels(temp.GetLabels())
	return nil
}
