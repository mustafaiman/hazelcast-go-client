// Copyright (c) 2008-2017, Hazelcast, Inc. All Rights Reserved.
//
// Licensed under the Apache License, Version 2.0 (the "License")
// you may not use this file except in compliance with the License.
// You may obtain a copy of the License at
//
// http://www.apache.org/licenses/LICENSE-2.0
//
// Unless required by applicable law or agreed to in writing, software
// distributed under the License is distributed on an "AS IS" BASIS,
// WITHOUT WARRANTIES OR CONDITIONS OF ANY KIND, either express or implied.
// See the License for the specific language governing permissions and
// limitations under the License.

package protocol

import (
	. "github.com/hazelcast/hazelcast-go-client/internal/serialization"
)

type MapEntriesWithPredicateResponseParameters struct {
	Response *[]Pair
}

func MapEntriesWithPredicateCalculateSize(name *string, predicate *Data) int {
	// Calculates the request payload size
	dataSize := 0
	dataSize += StringCalculateSize(name)
	dataSize += DataCalculateSize(predicate)
	return dataSize
}

func MapEntriesWithPredicateEncodeRequest(name *string, predicate *Data) *ClientMessage {
	// Encode request into clientMessage
	clientMessage := NewClientMessage(nil, MapEntriesWithPredicateCalculateSize(name, predicate))
	clientMessage.SetMessageType(MAP_ENTRIESWITHPREDICATE)
	clientMessage.IsRetryable = true
	clientMessage.AppendString(name)
	clientMessage.AppendData(predicate)
	clientMessage.UpdateFrameLength()
	return clientMessage
}

func MapEntriesWithPredicateDecodeResponse(clientMessage *ClientMessage) *MapEntriesWithPredicateResponseParameters {
	// Decode response from client message
	parameters := new(MapEntriesWithPredicateResponseParameters)

	responseSize := clientMessage.ReadInt32()
	response := make([]Pair, responseSize)
	for responseIndex := 0; responseIndex < int(responseSize); responseIndex++ {
		var responseItem Pair
		responseItemKey := clientMessage.ReadData()
		responseItemVal := clientMessage.ReadData()
		responseItem.key = responseItemKey
		responseItem.value = responseItemVal
		response[responseIndex] = responseItem
	}
	parameters.Response = &response

	return parameters
}
