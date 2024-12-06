
# Go Program: Calculate Distance Between Two Points

## Description
This program calculates the Euclidean distance between two points in a 2D plane using Go.

## Features
- Accepts two points as input from the user.
- Calculates the distance between the points using the formula:
  \[
  	ext{distance} = \sqrt{(x_2 - x_1)^2 + (y_2 - y_1)^2}
  \]
- Provides error handling for invalid inputs.

## Usage

### Prerequisites
- Go programming language installed. [Install Go](https://go.dev/doc/install)

### How to Run
1. Clone or download this repository.
2. Open a terminal in the directory containing the `main.go` file.
3. Run the following commands:
   ```bash
   go run main.go
   ```

### Input Format
- The program will prompt you to enter the coordinates of two points in the format:
  ```
  x1 y1
  x2 y2
  ```

### Example
#### Input:
```
Enter coordinates for 1st point [x1 y1]: 3 4
Enter coordinates for 2nd point [x2 y2]: 6 8
```

#### Output:
```
The distance between the points is: 5.00
```

## Code Structure
- **Point Struct**: Represents a point in 2D space with `x` and `y` coordinates.
- **main Function**: Handles input, validates it, and computes the distance.

## Error Handling
- Invalid inputs are handled with appropriate error messages.
- The program terminates if inputs are not in the expected format.

## License
This project is licensed under the MIT License.
