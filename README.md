# Biscepter Test Repo

This repository is used to test and demonstrate [biscepter](https://github.com/DominicWuest/biscepter), the efficient git bisecter.

There is a single file `main.go` which is a simple Go webserver.
It contains several endpoints, accessible under `/i`, where `i` is an integer that ranges from 0 to 5.
The endpoint `/i` is considered functional if it returns `i`, otherwise it is considered buggy.

This repository's git history looks as follows:

```
*   80afecd - Merge branch 'make-endpoints-buggy'
|\  
| *   07f20d6 - Merge branch 'make-endpoint-5-buggy' into make-endpoints-buggy
| |\  
| | * 50c1c00 - Empty Commit 1
| | *   a398ec4 - Merge branch 'make-endpoint-5-buggy-nested-branch' into make-endpoint-5-buggy
| | |\  
| | | * cfad207 - Add bug to /5
| | |/  
| | * 586f00a - Empty Commit 0
| |/  
| *   7639476 - Merge pull request #1 from DominicWuest/make-endpoint-4-buggy
| |\  
| | * 100d1a2 - Empty Commit 1
| | * 72cad4a - Add bug to /4
| | * 7b540da - Empty Commit 0
| |/  
| *   e4d49ab - Merge branch 'make-endpoint-3-buggy' into make-endpoints-buggy
| |\  
| | * db9cf6a - Add bug to /3
| |/  
| * bfeaae4 - Empty Commit 0
|/  
* 76b5c32 - Add /3, /4 and /5 testing endpoints
* d3245c0 - Empty Commit 2
* 0497a90 - Empty Commit 1
* 45b648f - Empty Commit 0
* 9b70eda - Add bug to /2
* 22a405d - Add bug to /1
* 03cdf84 - Add bug to /0
* 8ee0e2a - Make simple project
```
