# Ga e Of Life

Implementation done for takehome test.

Build the application:
```
$ cd {PROJECT_DIR}/
$ go build main.go -o out
```

Create a test file:
```
$ cat > in <<EOF
> 0 0 0 0 0
> 0 0 1 0 0
> 0 0 0 1 0
> 0 1 1 1 0
> 0 0 0 0 0
> ^D
```

Exectue:
```
./out -ng 3 < in
0 0 0 0 0 
0 0 1 0 0 
0 0 0 1 0 
0 1 1 1 0 
0 0 0 0 0 

0 0 0 0 0 
0 0 0 0 0 
0 1 0 1 0 
0 0 1 1 0 
0 0 1 0 0 

0 0 0 0 0 
0 0 0 0 0 
0 0 0 1 0 
0 1 0 1 0 
0 0 1 1 0 

0 0 0 0 0 
0 0 0 0 0 
0 0 1 0 0 
0 0 0 1 1 
0 0 1 1 0 
```

* There are no test cases since they were not required by the test. Please let me know if you need them.
* Tests would contain table based method that match multiple cases of input grid with corresponding final state after `n` generations.
* Input only performs very basic validations.
