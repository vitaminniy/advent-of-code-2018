package main

import (
	"strconv"
	"strings"
)

type rectangle struct {
	id            int
	x, y          int
	width, height int
}

//#1 @ 871,327: 16x20
func newRectangle(input string) (r rectangle, err error) {
	values := strings.Split(input, " ")

	if r.id, err = strconv.Atoi(values[0][1:]); err != nil {
		return r, err
	}

	coords := strings.Split(values[2], ",")
	if r.x, err = strconv.Atoi(coords[0]); err != nil {
		return r, err
	}
	if r.y, err = strconv.Atoi(coords[1][:len(coords[1])-1]); err != nil {
		return r, err
	}

	sizes := strings.Split(values[3], "x")
	if r.width, err = strconv.Atoi(sizes[0]); err != nil {
		return r, err
	}
	if r.height, err = strconv.Atoi(sizes[1]); err != nil {
		return r, err
	}

	return r, err
}
