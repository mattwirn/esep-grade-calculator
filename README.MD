# ESEP Grade Calculator

This library is a simple utility to calculate your final letter grade from a set of assignments, exams, and essays.

### Features

- Add **assignments**, **exams**, and **essays** with individual scores.  
- Compute a **weighted average**:  
  - Assignments: **50%**  
  - Exams: **35%**  
  - Essays: **15%**  
- Convert the numerical score into a **letter grade** using the following scale:
  - **A**: 90–100  
  - **B**: 80–89  
  - **C**: 70–79  
  - **D**: 60–69  
  - **F**: < 60  

### Library Functions

##### `NewGradeCalculator()`
Creates a new grade calculator object of type `GradeCalculator` with empty grade lists.

##### `.AddGrade(name string, grade int, gradeType GradeType)`
Adds a grade of the given type (`Assignment`, `Exam`, or `Essay`) to the calculator.

##### `.GetFinalGrade() string`
Computes the weighted average of all grades and returns the corresponding **letter grade**.


### Runtime Requirements:
Golang 1.24+

#### Installing Go
If you don't already have Go installed, follow these steps:

1. Visit the official Go download page: https://go.dev/dl
2. Download the installer for your operating system
3. Verify intallation in a terminal:

    ```sh
    go version
    ```
    Confirm that the command prints the installed version of Go.


### Running Unit Tests
Run unit tests:
```
cd grade_calculator/
go test -v ./...
```

Run unit tests with code coverage:
```
cd grade_calculator/
go test -v -coverprofile='coverage.txt' ./...
```
