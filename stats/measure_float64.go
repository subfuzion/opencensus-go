// Copyright 2017, OpenCensus Authors
//
// Licensed under the Apache License, Version 2.0 (the "License");
// you may not use this file except in compliance with the License.
// You may obtain a copy of the License at
//
//     http://www.apache.org/licenses/LICENSE-2.0
//
// Unless required by applicable law or agreed to in writing, software
// distributed under the License is distributed on an "AS IS" BASIS,
// WITHOUT WARRANTIES OR CONDITIONS OF ANY KIND, either express or implied.
// See the License for the specific language governing permissions and
// limitations under the License.
//

package stats

// MeasureFloat64 is a measure of type float64.
type MeasureFloat64 struct {
	name        string
	unit        string
	description string
}

// Name returns the name of the measure.
func (m *MeasureFloat64) Name() string {
	return m.name
}

// Description returns the description of the measure.
func (m *MeasureFloat64) Description() string {
	return m.description
}

// Unit returns the unit of the measure.
func (m *MeasureFloat64) Unit() string {
	return m.unit
}

// M creates a new float64 measurement.
// Use Record to record measurements.
func (m *MeasureFloat64) M(v float64) Measurement {
	return Measurement{m: m, v: v}
}
