# GoLang_training

Go is Object Oriented
1. Encapsulations
 - State("fields")
 - Behaviour("methods")
 - Exported and unexported, viewable and not viewable (traditional public/private)

2. Reusability 
 - Inheritance("Embedded types")

3. Polymorphism
 - Interfaces

4. Overriding
 - Promotion

5. We dont create classes, we create TYPE
6. We don't instantiate, we create a VALUE of a TYPE
7. Logically organize your fields together. Readability & clarity trump performance as a design concern. 
 	Go will be performant. Go for readability first. However, if you are in a situation where you need to prioritize performance: 
 	lay the fields out from largest to smallest, eg, int 64, int64, float32, bool

Functions
- In Go are PASS BY VALUE
- Syntax --> func (r receiver) identifier(parameters) (return(s)) 	{ ... }


- keyword identifier type --> var x int
A value can be of more than one type

Interface allows to deal with data of multiple types, if the methods are attached to a type
If there is an empty interface, every value can go into that

- Functions are first class citizens in GO, can be used as a type


Variadic parameter and empty interface -- Can take any number of parmaeters