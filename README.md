# ✈️ Estimated Zero Fuel Weight (EZFW) Calculator

A client-side web application designed for fast, accurate calculation of an aircraft's Estimated Zero Fuel Weight (EZFW), a critical figure for flight planning and weight & balance compliance.

This project is structured as a **Go (Golang) web server** serving static HTML, CSS, and JavaScript assets.

## ✨ Features

* **Go Server Backend:** Uses the `net/http` package to securely serve the frontend assets.
* **Dynamic Aircraft Selection:** Loads aircraft-specific data (Registration, DOW, and Type) from the `aircraft.csv` file.
* **Automated Weight Calculations:** Instantly computes Total Passenger Weight, Total Bag Weight, and Total ULD Tare based on user inputs.
* **Wide Body Logic:** Automatically switches to calculate Unit Load Device (ULD) requirements (Total ULDs, Total ULD Tare) when a Wide Body aircraft is selected.
* **Zero Fuel Weight (ZFW) Rounding:** Applies the industry standard rounding rule: EZFW is rounded up to the nearest 100 units.
* **Data Export:** Allows the user to save all input and calculated results as a structured JSON file.

---

## ⚙️ Technical Deep Dive (Aviation & Calculation Logic)

### Core Concept: Zero Fuel Weight (ZFW)

The Zero Fuel Weight (ZFW) is the total weight of the aircraft and all its contents **excluding** usable fuel. It is a critical structural limitation for most aircraft.

This calculator determines the **Estimated** Zero Fuel Weight (EZFW) by aggregating the following calculated components:

$$\text{EZFW} = \text{DOW} + \text{Total Pax Wt} + \text{Total Bag Wt} + \text{Cargo Wt} + \text{Service Wt} + \text{Total ULD Tare Wt (if applicable)}$$

The final result is rounded up to the nearest 100 units: $$\text{ZFW (Final)} = \text{Ceiling}(\text{EZFW} / 100) \times 100$$

---

## 🚀 Project Setup and Execution

This project requires a Go environment to run the server.

### Prerequisites

* **Go (1.18 or later)**
* `aircraft.csv` file present (as shown in your example).

### Project Structure

For the `main.go` server to function correctly, your files must be organized as follows:
