package data

import (
	"testing"

	"github.com/kylelemons/godebug/pretty"
	appsv1 "k8s.io/api/apps/v1"
	corev1 "k8s.io/api/core/v1"
	networkingv1 "k8s.io/api/networking/v1"
	rbacv1 "k8s.io/api/rbac/v1"
	metav1 "k8s.io/apimachinery/pkg/apis/meta/v1"
	"k8s.io/apimachinery/pkg/runtime"
)

func TestNewEntry(t *testing.T) {
	t.Parallel()

	meta := metav1.ObjectMeta{
		UID: "123",
	}

	tests := []struct {
		name    string
		obj     runtime.Object
		st      SourceType
		ct      ChangeType
		want    Entry
		wantErr bool
	}{
		{
			name:    "Error: Invalid type",
			obj:     &corev1.PodList{},
			st:      STWatchList,
			ct:      CTAdd,
			wantErr: true,
		},
		{
			name: "Success: Namespace type",
			obj:  &corev1.Namespace{ObjectMeta: meta},
			st:   STInformer,
			ct:   CTDelete,
			want: Entry{
				data:       &corev1.Namespace{ObjectMeta: meta},
				sourceType: STInformer,
				changeType: CTDelete,
				objectType: OTNamespace,
				uid:        meta.UID,
			},
		},
		{
			name: "Success: Node type",
			obj:  &corev1.Node{ObjectMeta: meta},
			st:   STInformer,
			ct:   CTUpdate,
			want: Entry{
				data:       &corev1.Node{ObjectMeta: meta},
				sourceType: STInformer,
				changeType: CTUpdate,
				objectType: OTNode,
				uid:        meta.UID,
			},
		},
		{
			name: "Success: PersistentVolume type",
			obj:  &corev1.PersistentVolume{ObjectMeta: meta},
			st:   STWatchList,
			ct:   CTAdd,
			want: Entry{
				data:       &corev1.PersistentVolume{ObjectMeta: meta},
				sourceType: STWatchList,
				changeType: CTAdd,
				objectType: OTPersistentVolume,
				uid:        meta.UID,
			},
		},
		{
			name: "Success: Pod type",
			obj:  &corev1.Pod{ObjectMeta: meta},
			st:   STWatchList,
			ct:   CTAdd,
			want: Entry{
				data:       &corev1.Pod{ObjectMeta: meta},
				sourceType: STWatchList,
				changeType: CTAdd,
				objectType: OTPod,
				uid:        meta.UID,
			},
		},
		{
			name: "Success: ClusterRole type",
			obj:  &rbacv1.ClusterRole{ObjectMeta: meta},
			st:   STWatchList,
			ct:   CTAdd,
			want: Entry{
				data:       &rbacv1.ClusterRole{ObjectMeta: meta},
				sourceType: STWatchList,
				changeType: CTAdd,
				objectType: OTClusterRole,
				uid:        meta.UID,
			},
		},
		{
			name: "Success: ClusterRoleBinding type",
			obj:  &rbacv1.ClusterRoleBinding{ObjectMeta: meta},
			st:   STWatchList,
			ct:   CTAdd,
			want: Entry{
				data:       &rbacv1.ClusterRoleBinding{ObjectMeta: meta},
				sourceType: STWatchList,
				changeType: CTAdd,
				objectType: OTClusterRoleBinding,
				uid:        meta.UID,
			},
		},
		{
			name: "Success: Role type",
			obj:  &rbacv1.Role{ObjectMeta: meta},
			st:   STWatchList,
			ct:   CTAdd,
			want: Entry{
				data:       &rbacv1.Role{ObjectMeta: meta},
				sourceType: STWatchList,
				changeType: CTAdd,
				objectType: OTRole,
				uid:        meta.UID,
			},
		},
		{
			name: "Success: RoleBinding type",
			obj:  &rbacv1.RoleBinding{ObjectMeta: meta},
			st:   STWatchList,
			ct:   CTAdd,
			want: Entry{
				data:       &rbacv1.RoleBinding{ObjectMeta: meta},
				sourceType: STWatchList,
				changeType: CTAdd,
				objectType: OTRoleBinding,
				uid:        meta.UID,
			},
		},
		{
			name: "Success: Service type",
			obj:  &corev1.Service{ObjectMeta: meta},
			st:   STWatchList,
			ct:   CTAdd,
			want: Entry{
				data:       &corev1.Service{ObjectMeta: meta},
				sourceType: STWatchList,
				changeType: CTAdd,
				objectType: OTService,
				uid:        meta.UID,
			},
		},
		{
			name: "Success: Deployment type",
			obj:  &appsv1.Deployment{ObjectMeta: meta},
			st:   STWatchList,
			ct:   CTAdd,
			want: Entry{
				data:       &appsv1.Deployment{ObjectMeta: meta},
				sourceType: STWatchList,
				changeType: CTAdd,
				objectType: OTDeployment,
				uid:        meta.UID,
			},
		},
		{
			name: "Success: Ingress type",
			obj:  &networkingv1.Ingress{ObjectMeta: meta},
			st:   STWatchList,
			ct:   CTAdd,
			want: Entry{
				data:       &networkingv1.Ingress{ObjectMeta: meta},
				sourceType: STWatchList,
				changeType: CTAdd,
				objectType: OTIngressController,
				uid:        meta.UID,
			},
		},
		{
			name: "Success: Endpoints type",
			obj:  &corev1.Endpoints{ObjectMeta: meta},
			st:   STWatchList,
			ct:   CTAdd,
			want: Entry{
				data:       &corev1.Endpoints{ObjectMeta: meta},
				sourceType: STWatchList,
				changeType: CTAdd,
				objectType: OTEndpoint,
				uid:        meta.UID,
			},
		},
	}

	for _, test := range tests {
		got, err := NewEntry(test.obj, test.st, test.ct)
		switch {
		case err == nil && test.wantErr:
			t.Errorf("TestNewEntry(%s): got err == nil, want err != nil", test.name)
			continue
		case err != nil && !test.wantErr:
			t.Errorf("TestNewEntry(%s): got err == %v, want err == nil", test.name, err)
			continue
		case err != nil:
			continue
		}

		if diff := pretty.Compare(test.want, got); diff != "" {
			t.Errorf("TestNewEntry(%s): -want/+got:\n%s", test.name, diff)
		}
	}
}

func TestAssertTo(t *testing.T) {
	t.Parallel()

	tests := []struct {
		name    string
		obj     runtime.Object
		want    any
		wantErr bool
	}{
		{
			name:    "Error: object is nil",
			obj:     nil,
			wantErr: true,
		},
		{
			name:    "Error: object is not a assertion type",
			obj:     &corev1.Node{},
			wantErr: true,
		},
		{
			name: "Success: object is assertion type",
			obj:  &corev1.Pod{},
			want: &corev1.Pod{},
		},
	}

	for _, test := range tests {
		got, err := assert[*corev1.Pod](test.obj)
		switch {
		case err == nil && test.wantErr:
			t.Errorf("TestAssertTo(%s): got err == nil, want err != nil", test.name)
			continue
		case err != nil && !test.wantErr:
			t.Errorf("TestAssertTo(%s): got err == %v, want err == nil", test.name, err)
			continue
		case err != nil:
			continue
		}

		if diff := pretty.Compare(test.want, got); diff != "" {
			t.Errorf("TestAssertTo(%s): -want/+got:\n%s", test.name, diff)
		}
	}
}
