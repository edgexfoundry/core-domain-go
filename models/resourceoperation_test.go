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

func TestResourceOperation_MarshalJSON(t *testing.T) {
	tests := []struct {
		name    string
		ro      ResourceOperation
		want    []byte
		wantErr bool
	}{
	// TODO: Add test cases.
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			got, err := tt.ro.MarshalJSON()
			if (err != nil) != tt.wantErr {
				t.Errorf("ResourceOperation.MarshalJSON() error = %v, wantErr %v", err, tt.wantErr)
				return
			}
			if !reflect.DeepEqual(got, tt.want) {
				t.Errorf("ResourceOperation.MarshalJSON() = %v, want %v", got, tt.want)
			}
		})
	}
}

func TestResourceOperation_String(t *testing.T) {
	tests := []struct {
		name string
		ro   ResourceOperation
		want string
	}{
	// TODO: Add test cases.
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			if got := tt.ro.String(); got != tt.want {
				t.Errorf("ResourceOperation.String() = %v, want %v", got, tt.want)
			}
		})
	}
}
