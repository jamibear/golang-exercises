# golang-exercises
## section 4: organizing data with structs

### Arrays vs Slices 

Arrays 
- Primitive data structure 
- Can't be resizes 
- Rarely used directly 

Slices 
- Can grow and shrink 
- Used 99% of the time for lists of elements

### Value Types vs Reference Types 

Value types use pointers to change values in a function 
- int 
- float 
- string 
- bool 
- structs 

Reference types don't need pointers to change value in a function 
- slices 
- maps 
- channels 
- pointers 
- functions 

