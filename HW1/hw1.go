/*
 * +===============================================
 * | Author:        Parham Alvani <parham.alvani@gmail.com>
 * |
 * | Creation Date: 23-10-2017
 * |
 * | File Name:     linesearch.go
 * +===============================================
 */

package main

import (
	"fmt"
	"math"
)

type variable struct {
	x1, x2, x3 float64
}

func main() {
	var x []variable
	x = make([]variable, 100)

	// starting point
	fmt.Printf("x0 (x1, x2, x3) = ")
	fmt.Scanf("%f %f %f", &x[0].x1, &x[0].x2, &x[0].x3)

	// backtracking
	var alpha, beta, c float64
	fmt.Printf("alpha beta c = ")
	fmt.Scanf("%f %f %f", &alpha, &beta, &c)

	// error
	var epsilon float64
	fmt.Printf("epsilon = ")
	fmt.Scanf("%f", &epsilon)

	fmt.Printf("=====\n")
	fmt.Printf("x0 = (%f, %f, %f)\n", x[0].x1, x[0].x2, x[0].x3)
	fmt.Printf("alpha = %f\n", alpha)
	fmt.Printf("beta = %f\n", beta)
	fmt.Printf("c = %f\n", c)
	fmt.Printf("epsilon = %f\n", epsilon)
	fmt.Printf("=====\n")

	var k int

	for gf(x[k]).norm() >= epsilon {

		// steepest descent
		p := gf(x[k]).neg()

		// backtracking
		for f(x[k].add(p.scale(alpha))) > f(x[k])+c*alpha*p.dot(gf(x[k])) {
			alpha = alpha * beta
		}

		// x[k + 1] = x[k] + p * alpha
		x = append(x, x[k].add(p))

		// results
		fmt.Printf("===== k = %d\n", k)
		fmt.Printf("x[%d] = (%f, %f, %f)\n", k, x[k].x1, x[k].x2, x[k].x3)
		fmt.Printf("=====\n")

		k++
	}
}

func f(v variable) float64 {
	return -v.x1 - v.x2 - v.x3 - 1/(v.x1-20) - 1/(v.x1+v.x2-30) - 1/(v.x2-20) - 1/(v.x2+v.x3-30) - 1/(v.x3-25) - 1/(v.x2-math.Log(v.x1))
}

func gf(v variable) variable {
	return variable{
		x1: -1 + 1/math.Pow(v.x1-20, 2) + 1/math.Pow(v.x1+v.x2-30, 2) - 1/(v.x1*math.Pow(v.x2-math.Log(v.x1), 2)),
		x2: -1 + 1/math.Pow(v.x2-20, 2) + 1/math.Pow(v.x1+v.x2-30, 2) + 1/math.Pow(v.x2+v.x3-30, 2) + 1/math.Pow(v.x2-math.Log(v.x1), 2),
		x3: -1 + 1/math.Pow(v.x3-25, 2) + 1/math.Pow(v.x2+v.x3-30, 2),
	}
}

func (v variable) norm() float64 {
	return math.Sqrt(v.x1*v.x1 + v.x2*v.x2 + v.x3*v.x3)
}

func (v variable) neg() variable {
	return variable{
		-v.x1,
		-v.x2,
		-v.x3,
	}
}

func (v variable) add(o variable) variable {
	return variable{
		v.x1 + o.x1,
		v.x2 + o.x2,
		v.x3 + o.x3,
	}
}

func (v variable) dot(o variable) float64 {
	return v.x1*o.x1 + v.x2*o.x2 + v.x3*o.x3
}

func (v variable) scale(s float64) variable {
	return variable{
		v.x1 * s,
		v.x2 * s,
		v.x3 * s,
	}
}
