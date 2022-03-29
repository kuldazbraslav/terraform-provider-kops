package schemas

import (
	"testing"

	"github.com/google/go-cmp/cmp"
	"k8s.io/kops/pkg/apis/kops"
)

func TestExpandResourceAWSAuthenticationIdentityMappingSpec(t *testing.T) {
	_default := kops.AWSAuthenticationIdentityMappingSpec{}
	type args struct {
		in map[string]interface{}
	}
	tests := []struct {
		name string
		args args
		want kops.AWSAuthenticationIdentityMappingSpec
	}{
		{
			name: "default",
			args: args{
				in: map[string]interface{}{
					"arn":      "",
					"username": "",
					"groups":   func() []interface{} { return nil }(),
				},
			},
			want: _default,
		},
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			got := ExpandResourceAWSAuthenticationIdentityMappingSpec(tt.args.in)
			if diff := cmp.Diff(tt.want, got); diff != "" {
				t.Errorf("ExpandResourceAWSAuthenticationIdentityMappingSpec() mismatch (-want +got):\n%s", diff)
			}
		})
	}
}

func TestFlattenResourceAWSAuthenticationIdentityMappingSpecInto(t *testing.T) {
	_default := map[string]interface{}{
		"arn":      "",
		"username": "",
		"groups":   func() []interface{} { return nil }(),
	}
	type args struct {
		in kops.AWSAuthenticationIdentityMappingSpec
	}
	tests := []struct {
		name string
		args args
		want map[string]interface{}
	}{
		{
			name: "default",
			args: args{
				in: kops.AWSAuthenticationIdentityMappingSpec{},
			},
			want: _default,
		},
		{
			name: "ARN - default",
			args: args{
				in: func() kops.AWSAuthenticationIdentityMappingSpec {
					subject := kops.AWSAuthenticationIdentityMappingSpec{}
					subject.ARN = ""
					return subject
				}(),
			},
			want: _default,
		},
		{
			name: "Username - default",
			args: args{
				in: func() kops.AWSAuthenticationIdentityMappingSpec {
					subject := kops.AWSAuthenticationIdentityMappingSpec{}
					subject.Username = ""
					return subject
				}(),
			},
			want: _default,
		},
		{
			name: "Groups - default",
			args: args{
				in: func() kops.AWSAuthenticationIdentityMappingSpec {
					subject := kops.AWSAuthenticationIdentityMappingSpec{}
					subject.Groups = nil
					return subject
				}(),
			},
			want: _default,
		},
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			got := map[string]interface{}{}
			FlattenResourceAWSAuthenticationIdentityMappingSpecInto(tt.args.in, got)
			if diff := cmp.Diff(tt.want, got); diff != "" {
				t.Errorf("FlattenResourceAWSAuthenticationIdentityMappingSpec() mismatch (-want +got):\n%s", diff)
			}
		})
	}
}

func TestFlattenResourceAWSAuthenticationIdentityMappingSpec(t *testing.T) {
	_default := map[string]interface{}{
		"arn":      "",
		"username": "",
		"groups":   func() []interface{} { return nil }(),
	}
	type args struct {
		in kops.AWSAuthenticationIdentityMappingSpec
	}
	tests := []struct {
		name string
		args args
		want map[string]interface{}
	}{
		{
			name: "default",
			args: args{
				in: kops.AWSAuthenticationIdentityMappingSpec{},
			},
			want: _default,
		},
		{
			name: "ARN - default",
			args: args{
				in: func() kops.AWSAuthenticationIdentityMappingSpec {
					subject := kops.AWSAuthenticationIdentityMappingSpec{}
					subject.ARN = ""
					return subject
				}(),
			},
			want: _default,
		},
		{
			name: "Username - default",
			args: args{
				in: func() kops.AWSAuthenticationIdentityMappingSpec {
					subject := kops.AWSAuthenticationIdentityMappingSpec{}
					subject.Username = ""
					return subject
				}(),
			},
			want: _default,
		},
		{
			name: "Groups - default",
			args: args{
				in: func() kops.AWSAuthenticationIdentityMappingSpec {
					subject := kops.AWSAuthenticationIdentityMappingSpec{}
					subject.Groups = nil
					return subject
				}(),
			},
			want: _default,
		},
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			got := FlattenResourceAWSAuthenticationIdentityMappingSpec(tt.args.in)
			if diff := cmp.Diff(tt.want, got); diff != "" {
				t.Errorf("FlattenResourceAWSAuthenticationIdentityMappingSpec() mismatch (-want +got):\n%s", diff)
			}
		})
	}
}
