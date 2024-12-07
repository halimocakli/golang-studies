
# Golang Application

This project demonstrates basic mathematical operations in Golang, including rounding, flooring, and ceiling of a floating-point number inputted by the user.

## 📘 **Project Overview**
This simple command-line application takes a floating-point number as input and outputs its rounded, floored, and ceiled values using Golang's standard `math` package.

---

## 🚀 **How to Run**

To run this application, follow these steps:

1. **Ensure Golang is installed** on your system. To verify, run:
   ```bash
   go version
   ```

2. **Run the application** with the following command:
   ```bash
   go run main.go
   ```

3. **Input a floating-point number** when prompted (e.g., `3.14`).

---

## 🔧 **Requirements**

- Golang (version 1.18 or higher is recommended)

---

## 🛠️ **Code Explanation**

The application follows these steps:

1. **Input Handling**: 
   - Prompts the user to input a floating-point number.
   - Reads the input and checks for any input errors.
   
2. **Mathematical Operations**:
   - **Original**: Displays the original input number.
   - **Rounded**: Rounds the input to the nearest whole number using `math.Round()`.
   - **Floored**: Returns the largest whole number less than or equal to the input using `math.Floor()`.
   - **Ceiled**: Returns the smallest whole number greater than or equal to the input using `math.Ceil()`.

---

## 📂 **Project Structure**

```
├── main.go       # Main Golang application file
```

---

## 📘 **Example Usage**

```bash
$ go run main.go
Input a float value (e.g., 3.14): 3.14

Original: 3.14
Rounded: 3.00
Floored: 3.00
Ceiled: 4.00
```

---

## ❗ **Error Handling**

- If the user inputs an invalid number (like a string), the application will exit with an error message.

---

## 📚 **Golang Concepts Used**

- **Package `fmt`**: For input and output operations.
- **Package `math`**: For rounding, flooring, and ceiling calculations.
- **Package `log`**: For error handling and application termination on invalid input.

---

## 📜 **License**

This project is open-source and can be used freely for educational purposes.
