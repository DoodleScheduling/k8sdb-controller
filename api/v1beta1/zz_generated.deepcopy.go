// +build !ignore_autogenerated

/*


Licensed under the Apache License, Version 2.0 (the "License");
you may not use this file except in compliance with the License.
You may obtain a copy of the License at

    http://www.apache.org/licenses/LICENSE-2.0

Unless required by applicable law or agreed to in writing, software
distributed under the License is distributed on an "AS IS" BASIS,
WITHOUT WARRANTIES OR CONDITIONS OF ANY KIND, either express or implied.
See the License for the specific language governing permissions and
limitations under the License.
*/

// Code generated by controller-gen. DO NOT EDIT.

package v1beta1

import (
	runtime "k8s.io/apimachinery/pkg/runtime"
)

// DeepCopyInto is an autogenerated deepcopy function, copying the receiver, writing into out. in must be non-nil.
func (in *MongoDB) DeepCopyInto(out *MongoDB) {
	*out = *in
	out.TypeMeta = in.TypeMeta
	in.ObjectMeta.DeepCopyInto(&out.ObjectMeta)
	in.Spec.DeepCopyInto(&out.Spec)
	in.Status.DeepCopyInto(&out.Status)
}

// DeepCopy is an autogenerated deepcopy function, copying the receiver, creating a new MongoDB.
func (in *MongoDB) DeepCopy() *MongoDB {
	if in == nil {
		return nil
	}
	out := new(MongoDB)
	in.DeepCopyInto(out)
	return out
}

// DeepCopyObject is an autogenerated deepcopy function, copying the receiver, creating a new runtime.Object.
func (in *MongoDB) DeepCopyObject() runtime.Object {
	if c := in.DeepCopy(); c != nil {
		return c
	}
	return nil
}

// DeepCopyInto is an autogenerated deepcopy function, copying the receiver, writing into out. in must be non-nil.
func (in *MongoDBCredential) DeepCopyInto(out *MongoDBCredential) {
	*out = *in
	out.Vault = in.Vault
}

// DeepCopy is an autogenerated deepcopy function, copying the receiver, creating a new MongoDBCredential.
func (in *MongoDBCredential) DeepCopy() *MongoDBCredential {
	if in == nil {
		return nil
	}
	out := new(MongoDBCredential)
	in.DeepCopyInto(out)
	return out
}

// DeepCopyInto is an autogenerated deepcopy function, copying the receiver, writing into out. in must be non-nil.
func (in *MongoDBCredentialStatus) DeepCopyInto(out *MongoDBCredentialStatus) {
	*out = *in
}

// DeepCopy is an autogenerated deepcopy function, copying the receiver, creating a new MongoDBCredentialStatus.
func (in *MongoDBCredentialStatus) DeepCopy() *MongoDBCredentialStatus {
	if in == nil {
		return nil
	}
	out := new(MongoDBCredentialStatus)
	in.DeepCopyInto(out)
	return out
}

// DeepCopyInto is an autogenerated deepcopy function, copying the receiver, writing into out. in must be non-nil.
func (in MongoDBCredentials) DeepCopyInto(out *MongoDBCredentials) {
	{
		in := &in
		*out = make(MongoDBCredentials, len(*in))
		copy(*out, *in)
	}
}

// DeepCopy is an autogenerated deepcopy function, copying the receiver, creating a new MongoDBCredentials.
func (in MongoDBCredentials) DeepCopy() MongoDBCredentials {
	if in == nil {
		return nil
	}
	out := new(MongoDBCredentials)
	in.DeepCopyInto(out)
	return *out
}

// DeepCopyInto is an autogenerated deepcopy function, copying the receiver, writing into out. in must be non-nil.
func (in MongoDBCredentialsStatus) DeepCopyInto(out *MongoDBCredentialsStatus) {
	{
		in := &in
		*out = make(MongoDBCredentialsStatus, len(*in))
		copy(*out, *in)
	}
}

// DeepCopy is an autogenerated deepcopy function, copying the receiver, creating a new MongoDBCredentialsStatus.
func (in MongoDBCredentialsStatus) DeepCopy() MongoDBCredentialsStatus {
	if in == nil {
		return nil
	}
	out := new(MongoDBCredentialsStatus)
	in.DeepCopyInto(out)
	return *out
}

// DeepCopyInto is an autogenerated deepcopy function, copying the receiver, writing into out. in must be non-nil.
func (in *MongoDBDatabaseStatus) DeepCopyInto(out *MongoDBDatabaseStatus) {
	*out = *in
}

// DeepCopy is an autogenerated deepcopy function, copying the receiver, creating a new MongoDBDatabaseStatus.
func (in *MongoDBDatabaseStatus) DeepCopy() *MongoDBDatabaseStatus {
	if in == nil {
		return nil
	}
	out := new(MongoDBDatabaseStatus)
	in.DeepCopyInto(out)
	return out
}

// DeepCopyInto is an autogenerated deepcopy function, copying the receiver, writing into out. in must be non-nil.
func (in *MongoDBList) DeepCopyInto(out *MongoDBList) {
	*out = *in
	out.TypeMeta = in.TypeMeta
	in.ListMeta.DeepCopyInto(&out.ListMeta)
	if in.Items != nil {
		in, out := &in.Items, &out.Items
		*out = make([]MongoDB, len(*in))
		for i := range *in {
			(*in)[i].DeepCopyInto(&(*out)[i])
		}
	}
}

// DeepCopy is an autogenerated deepcopy function, copying the receiver, creating a new MongoDBList.
func (in *MongoDBList) DeepCopy() *MongoDBList {
	if in == nil {
		return nil
	}
	out := new(MongoDBList)
	in.DeepCopyInto(out)
	return out
}

// DeepCopyObject is an autogenerated deepcopy function, copying the receiver, creating a new runtime.Object.
func (in *MongoDBList) DeepCopyObject() runtime.Object {
	if c := in.DeepCopy(); c != nil {
		return c
	}
	return nil
}

// DeepCopyInto is an autogenerated deepcopy function, copying the receiver, writing into out. in must be non-nil.
func (in *MongoDBSpec) DeepCopyInto(out *MongoDBSpec) {
	*out = *in
	if in.Credentials != nil {
		in, out := &in.Credentials, &out.Credentials
		*out = make(MongoDBCredentials, len(*in))
		copy(*out, *in)
	}
}

// DeepCopy is an autogenerated deepcopy function, copying the receiver, creating a new MongoDBSpec.
func (in *MongoDBSpec) DeepCopy() *MongoDBSpec {
	if in == nil {
		return nil
	}
	out := new(MongoDBSpec)
	in.DeepCopyInto(out)
	return out
}

// DeepCopyInto is an autogenerated deepcopy function, copying the receiver, writing into out. in must be non-nil.
func (in *MongoDBStatus) DeepCopyInto(out *MongoDBStatus) {
	*out = *in
	out.MongoDBAvailabilityStatus = in.MongoDBAvailabilityStatus
	if in.CredentialsStatus != nil {
		in, out := &in.CredentialsStatus, &out.CredentialsStatus
		*out = make(MongoDBCredentialsStatus, len(*in))
		copy(*out, *in)
	}
}

// DeepCopy is an autogenerated deepcopy function, copying the receiver, creating a new MongoDBStatus.
func (in *MongoDBStatus) DeepCopy() *MongoDBStatus {
	if in == nil {
		return nil
	}
	out := new(MongoDBStatus)
	in.DeepCopyInto(out)
	return out
}

// DeepCopyInto is an autogenerated deepcopy function, copying the receiver, writing into out. in must be non-nil.
func (in *PostgreSQL) DeepCopyInto(out *PostgreSQL) {
	*out = *in
	out.TypeMeta = in.TypeMeta
	in.ObjectMeta.DeepCopyInto(&out.ObjectMeta)
	in.Spec.DeepCopyInto(&out.Spec)
	out.Status = in.Status
}

// DeepCopy is an autogenerated deepcopy function, copying the receiver, creating a new PostgreSQL.
func (in *PostgreSQL) DeepCopy() *PostgreSQL {
	if in == nil {
		return nil
	}
	out := new(PostgreSQL)
	in.DeepCopyInto(out)
	return out
}

// DeepCopyObject is an autogenerated deepcopy function, copying the receiver, creating a new runtime.Object.
func (in *PostgreSQL) DeepCopyObject() runtime.Object {
	if c := in.DeepCopy(); c != nil {
		return c
	}
	return nil
}

// DeepCopyInto is an autogenerated deepcopy function, copying the receiver, writing into out. in must be non-nil.
func (in *PostgreSQLCredential) DeepCopyInto(out *PostgreSQLCredential) {
	*out = *in
	out.Vault = in.Vault
}

// DeepCopy is an autogenerated deepcopy function, copying the receiver, creating a new PostgreSQLCredential.
func (in *PostgreSQLCredential) DeepCopy() *PostgreSQLCredential {
	if in == nil {
		return nil
	}
	out := new(PostgreSQLCredential)
	in.DeepCopyInto(out)
	return out
}

// DeepCopyInto is an autogenerated deepcopy function, copying the receiver, writing into out. in must be non-nil.
func (in *PostgreSQLCredentialStatus) DeepCopyInto(out *PostgreSQLCredentialStatus) {
	*out = *in
}

// DeepCopy is an autogenerated deepcopy function, copying the receiver, creating a new PostgreSQLCredentialStatus.
func (in *PostgreSQLCredentialStatus) DeepCopy() *PostgreSQLCredentialStatus {
	if in == nil {
		return nil
	}
	out := new(PostgreSQLCredentialStatus)
	in.DeepCopyInto(out)
	return out
}

// DeepCopyInto is an autogenerated deepcopy function, copying the receiver, writing into out. in must be non-nil.
func (in PostgreSQLCredentials) DeepCopyInto(out *PostgreSQLCredentials) {
	{
		in := &in
		*out = make(PostgreSQLCredentials, len(*in))
		copy(*out, *in)
	}
}

// DeepCopy is an autogenerated deepcopy function, copying the receiver, creating a new PostgreSQLCredentials.
func (in PostgreSQLCredentials) DeepCopy() PostgreSQLCredentials {
	if in == nil {
		return nil
	}
	out := new(PostgreSQLCredentials)
	in.DeepCopyInto(out)
	return *out
}

// DeepCopyInto is an autogenerated deepcopy function, copying the receiver, writing into out. in must be non-nil.
func (in PostgreSQLCredentialsStatus) DeepCopyInto(out *PostgreSQLCredentialsStatus) {
	{
		in := &in
		*out = make(PostgreSQLCredentialsStatus, len(*in))
		copy(*out, *in)
	}
}

// DeepCopy is an autogenerated deepcopy function, copying the receiver, creating a new PostgreSQLCredentialsStatus.
func (in PostgreSQLCredentialsStatus) DeepCopy() PostgreSQLCredentialsStatus {
	if in == nil {
		return nil
	}
	out := new(PostgreSQLCredentialsStatus)
	in.DeepCopyInto(out)
	return *out
}

// DeepCopyInto is an autogenerated deepcopy function, copying the receiver, writing into out. in must be non-nil.
func (in *PostgreSQLDatabaseStatus) DeepCopyInto(out *PostgreSQLDatabaseStatus) {
	*out = *in
}

// DeepCopy is an autogenerated deepcopy function, copying the receiver, creating a new PostgreSQLDatabaseStatus.
func (in *PostgreSQLDatabaseStatus) DeepCopy() *PostgreSQLDatabaseStatus {
	if in == nil {
		return nil
	}
	out := new(PostgreSQLDatabaseStatus)
	in.DeepCopyInto(out)
	return out
}

// DeepCopyInto is an autogenerated deepcopy function, copying the receiver, writing into out. in must be non-nil.
func (in *PostgreSQLList) DeepCopyInto(out *PostgreSQLList) {
	*out = *in
	out.TypeMeta = in.TypeMeta
	in.ListMeta.DeepCopyInto(&out.ListMeta)
	if in.Items != nil {
		in, out := &in.Items, &out.Items
		*out = make([]PostgreSQL, len(*in))
		for i := range *in {
			(*in)[i].DeepCopyInto(&(*out)[i])
		}
	}
}

// DeepCopy is an autogenerated deepcopy function, copying the receiver, creating a new PostgreSQLList.
func (in *PostgreSQLList) DeepCopy() *PostgreSQLList {
	if in == nil {
		return nil
	}
	out := new(PostgreSQLList)
	in.DeepCopyInto(out)
	return out
}

// DeepCopyObject is an autogenerated deepcopy function, copying the receiver, creating a new runtime.Object.
func (in *PostgreSQLList) DeepCopyObject() runtime.Object {
	if c := in.DeepCopy(); c != nil {
		return c
	}
	return nil
}

// DeepCopyInto is an autogenerated deepcopy function, copying the receiver, writing into out. in must be non-nil.
func (in *PostgreSQLSpec) DeepCopyInto(out *PostgreSQLSpec) {
	*out = *in
	if in.Credentials != nil {
		in, out := &in.Credentials, &out.Credentials
		*out = make(PostgreSQLCredentials, len(*in))
		copy(*out, *in)
	}
}

// DeepCopy is an autogenerated deepcopy function, copying the receiver, creating a new PostgreSQLSpec.
func (in *PostgreSQLSpec) DeepCopy() *PostgreSQLSpec {
	if in == nil {
		return nil
	}
	out := new(PostgreSQLSpec)
	in.DeepCopyInto(out)
	return out
}

// DeepCopyInto is an autogenerated deepcopy function, copying the receiver, writing into out. in must be non-nil.
func (in *PostgreSQLStatus) DeepCopyInto(out *PostgreSQLStatus) {
	*out = *in
	out.DatabaseStatus = in.DatabaseStatus
	out.CredentialsStatus = in.CredentialsStatus
}

// DeepCopy is an autogenerated deepcopy function, copying the receiver, creating a new PostgreSQLStatus.
func (in *PostgreSQLStatus) DeepCopy() *PostgreSQLStatus {
	if in == nil {
		return nil
	}
	out := new(PostgreSQLStatus)
	in.DeepCopyInto(out)
	return out
}

// DeepCopyInto is an autogenerated deepcopy function, copying the receiver, writing into out. in must be non-nil.
func (in *Vault) DeepCopyInto(out *Vault) {
	*out = *in
}

// DeepCopy is an autogenerated deepcopy function, copying the receiver, creating a new Vault.
func (in *Vault) DeepCopy() *Vault {
	if in == nil {
		return nil
	}
	out := new(Vault)
	in.DeepCopyInto(out)
	return out
}
