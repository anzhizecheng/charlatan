// Copyright (C) 2019 Guillaume CAMUS
//
// This file is part of Charlatan.
//
// Charlatan is free software: you can redistribute it and/or modify
// it under the terms of the GNU General Public License as published by
// the Free Software Foundation, either version 3 of the License, or
// (at your option) any later version.
//
// Charlatan is distributed in the hope that it will be useful,
// but WITHOUT ANY WARRANTY; without even the implied warranty of
// MERCHANTABILITY or FITNESS FOR A PARTICULAR PURPOSE.  See the
// GNU General Public License for more details.
//
// You should have received a copy of the GNU General Public License
// along with Charlatan.  If not, see <http://www.gnu.org/licenses/>.

package normalizer

import (
	"fmt"
	"regexp"
	"strconv"

	"github.com/guiyomh/charlatan/contract"
)

var RANGE_REGEX = regexp.MustCompile(`(?i)(?P<refbase>[\p{L}\._\/\d]+)\{(?P<from>[0-9]+)(?:\.{2})(?P<to>[0-9]+)((,\s?(?P<step>[0-9]+))?)\}.*`)

type Range struct {
	chainabler
}

func (r *Range) CanDenormalize(ref string) bool {
	return RANGE_REGEX.MatchString(ref)
}

// BuildIds parse the referense to build a range of ids
// For example user_{1..6,3} => []string{
// 	"user_1",
// 	"user_3",
// 	"user_6",
// }
func (r *Range) BuildIds(matches map[string]string) ([]string, error) {
	var ids = make([]string, 0)
	refbase, ok := matches["refbase"]
	if !ok {
		return ids, fmt.Errorf("Could not retrieve 'refbase' in '%v'", matches)
	}
	from, err := getInt("from", matches)
	if err != nil {
		return ids, err
	}
	to, err := getInt("to", matches)
	if err != nil {
		return ids, err
	}
	step, err := getInt("step", matches)
	if err != nil {
		step = 1
	}

	for current := from; current <= to; current += step {
		ids = append(ids, fmt.Sprintf("%s%d", refbase, current))
	}

	return ids, nil
}

func getInt(key string, matches map[string]string) (int, error) {
	var value int
	v, ok := matches[key]
	if !ok {
		return value, fmt.Errorf("Could not retrieve '%s' in '%v'", key, matches)
	}
	return strconv.Atoi(v)
}

func (r *Range) Denormalize(bag contract.Bager, ref string, data map[string]interface{}) (contract.Bager, error) {
	matches, err := r.buildMatches(ref, RANGE_REGEX)
	if err != nil {
		return bag, err
	}
	ids, err := r.BuildIds(matches)
	if err != nil {
		return bag, err
	}
	fixtures := r.buildFixture(ids, data)
	for _, f := range fixtures {
		bag.Add(f)
	}
	return bag, nil
}
