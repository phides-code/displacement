/*Write a program which first prompts the user to enter values for acceleration, initial 
velocity, and initial displacement. Then the program should prompt the user to enter a value 
for time and the program should compute the displacement after the entered time.

You will need to define and use a function called GenDisplaceFn() which takes three float64 
arguments, acceleration a, initial velocity vo, and initial displacement so. GenDisplaceFn() 
should return a function which computes displacement as a function of time, assuming the 
given values acceleration, initial velocity, and initial displacement. The function returned 
by GenDisplaceFn() should take one float64 argument t, representing time, and return one 
float64 argument which is the displacement travelled after time t.
*/

package main 

import (
    "fmt"
)

func main() {
    var a, v0, s0, t float64
    
    // prompt the user to enter the values
    fmt.Printf("Please enter the value for acceleration: ")
    fmt.Scan(&a)
    
    fmt.Printf("Please enter the value for initial velocity: ")
    fmt.Scan(&v0)
    
    fmt.Printf("Please enter the value for initial displacement: ")
    fmt.Scan(&s0)
    
    fmt.Printf("Please enter the value for time: ")
    fmt.Scan(&t)
    
    fmt.Println("The displacement is: ")
    
    MyDisplaceFn := GenDisplaceFn(a, v0, s0)
    fmt.Println(MyDisplaceFn(t))
}

func GenDisplaceFn(a, v0, s0 float64 ) func (float64) float64 {
    
    fn := func (t float64) float64 { 
        // compute displacement as a function of t, a, v0, and s0
        // return one float64 argument which is the displacement travelled after time t.
        return 0.5 * a * t * t + v0 * t + s0
    }
    
    return fn
}
