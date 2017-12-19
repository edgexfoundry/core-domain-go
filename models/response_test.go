/*******************************************************************************
 * Copyright 2017 Dell Inc.
 *
 * Licensed under the Apache License, Version 2.0 (the "License"); you may not use this file except
 * in compliance with the License. You may obtain a copy of the License at
 *
 * http://www.apache.org/licenses/LICENSE-2.0
 *
 * Unless required by applicable law or agreed to in writing, software distributed under the License
 * is distributed on an "AS IS" BASIS, WITHOUT WARRANTIES OR CONDITIONS OF ANY KIND, either express
 * or implied. See the License for the specific language governing permissions and limitations under
 * the License.
 *
 * @microservice: core-domain-go library
 * @author: Ryan Comer & Spencer Bull, Dell
 * @version: 0.5.0
 *******************************************************************************/

package models

import (
	"reflect"
	"testing"
)

func TestResponse_MarshalJSON(t *testing.T) {
	tests := []struct {
		name    string
		r       Response
		want    []byte
		wantErr bool
	}{
	// TODO: Add test cases.
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			got, err := tt.r.MarshalJSON()
			if (err != nil) != tt.wantErr {
				t.Errorf("Response.MarshalJSON() error = %v, wantErr %v", err, tt.wantErr)
				return
			}
			if !reflect.DeepEqual(got, tt.want) {
				t.Errorf("Response.MarshalJSON() = %v, want %v", got, tt.want)
			}
		})
	}
}

func TestResponse_String(t *testing.T) {
	tests := []struct {
		name string
		a    Response
		want string
	}{
	// TODO: Add test cases.
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			if got := tt.a.String(); got != tt.want {
				t.Errorf("Response.String() = %v, want %v", got, tt.want)
			}
		})
	}
}

func TestResponse_Equals(t *testing.T) {
	type args struct {
		r2 Response
	}
	tests := []struct {
		name string
		r    Response
		args args
		want bool
	}{
	// TODO: Add test cases.
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			if got := tt.r.Equals(tt.args.r2); got != tt.want {
				t.Errorf("Response.Equals() = %v, want %v", got, tt.want)
			}
		})
	}
}
