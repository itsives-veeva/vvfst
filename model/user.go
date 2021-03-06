/*
This code serves as an example and is not meant for production use.

Copyright 2020 Veeva Systems Inc.

Licensed under the Apache License, Version 2.0 (the "License"); you may not use
this file except in compliance with the License. You may obtain a copy of the License at

http://www.apache.org/licenses/LICENSE-2.0

Unless required by applicable law or agreed to in writing, software distributed under
the License is distributed on an "AS IS" BASIS, WITHOUT WARRANTIES OR CONDITIONS OF ANY KIND,
either express or implied. See the License for the specific language governing permissions
and limitations under the License.
*/
package model

type VaultID struct {
	ID   int
	Name string
	URL  string
}

type BaseResult struct {
	ResponseStatus  string             `json:"responseStatus,omitempty"`
	ResponseMessage string             `json:"responseMessage,omitempty"`
	Errors          []*RestResultError `json:"errors,omitempty"`
}

type AuthResult struct {
	BaseResult
	SessionID string     `json:"sessionId,omitempty"`
	UserID    int        `json:"userId,omitempty"`
	VaultID   int        `json:"vaultId,omitempty"`
	VaultIDs  []*VaultID `json:"vaultIds,omitempty"`
}
