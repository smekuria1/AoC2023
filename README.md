# AoC2023
Advent of Code 2023 in GO
**There are better solutions out there this my humble approach**
## Day 1 completed
- Very intersting Puzzle learned a lot about string manipulation in GO
- Currently very inefficient a lot of strconv used
- Reversing each line helped find the solution
## Day 2 completed
- Relatively easier than Day1
- Triple nested loop for parsing the input could be improved using regex
- Used Hash maps for both p1 and p2(tracking minimum set of cubes)

## Day 3 Not completed
- Matrix Operations :(
- It is painfull trying to work with Matrixes in Go
- Hopefully comeback to this problem and maybe do it python or JS

## Day 4 completed
- Bless Even Days 
- Used a map to save the wining numbers of game line
    - Second pass to check if scratch numbers match and increment counts 
- Used the logic from P1 to make a helper function for P2 calculating how deep to iterate and add cards to our map with their counts
- iterate over map values and find sum
- Only solution that first came to mind was recursive approach 

## Day 5 Completed
- Was completely stumped on how to put the parsed data and make the mappings
- Had to look up hints for P1 to complete it
- P2 was fairly straight forward bruteforce approach took 7mins average to finish checking each seed pair
-  Can be improved by using go routines 

# Day 6 Completed
- Very Straightforward bruteforce approach that worked for both p1 and p2
- Had to adjust parsing slightly for P2 and used the same algorithm
- Quadratic math approach might improve run time but its already super fast 

# Day 7 P1 completed
- Didn't have time to attempt P2
- P1 took a lot of time although i believe the problem was not that complicated 
- Getting comfortable using maps and slices for parsinng and solving problems


# Day 8 
- First tried to use a custom built graph struct failed miserably cause of cycles
- Used a map with Node struct as values for left and right 
- Previous parsing functions made working with this input very easy and efficient 
- Had Absolutely no idea how to do P2 until I read LCM on the Reddit
- Same parsing and building the map only adjusted the traversal and used LCM function from the internet