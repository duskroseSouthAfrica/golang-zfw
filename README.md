# ✈️ Estimated Zero Fuel Weight (EZFW) Calculator

A client-side web application designed for fast, accurate calculation of an aircraft's Estimated Zero Fuel Weight (EZFW), a critical figure for flight planning and weight & balance compliance.

This calculator simplifies the load sheet process by automating complex weight accumulation and rounding rules, particularly those concerning Wide Body (containerized) operations.

## ✨ Features

* **Dynamic Aircraft Selection:** Loads aircraft-specific data (Registration, Dry Operating Weight (DOW), and Type) from a local `aircraft.csv` file.
* **Automated Weight Calculations:** Instantly computes Total Passenger Weight, Total Bag Weight, and Total ULD Tare based on user inputs.
* **Wide Body Logic:** Automatically switches to calculate Unit Load Device (ULD) requirements (Total ULDs, Total ULD Tare) when a Wide Body aircraft is selected.
* **Zero Fuel Weight (ZFW) Rounding:** Applies the industry standard rounding rule: EZFW is rounded up to the nearest 100 units (e.g., 95,350 kg rounds to 95,400 kg).
* **Data Export:** Allows the user to save all input and calculated results as a structured JSON file for record-keeping and regulatory compliance.

## ⚙️ Technical Deep Dive

### Core Concept: Zero Fuel Weight (ZFW)

The Zero Fuel Weight (ZFW) is the total weight of the aircraft and all its contents **excluding** usable fuel. It is a critical structural limitation for most aircraft, ensuring that the bending stress on the wing root caused by the weight of the fuselage and its contents does not exceed certified limits.

$$\text{ZFW} = \text{Operating Empty Weight (OEW)} + \text{Payload (PL)}$$

This calculator determines the **Estimated** Zero Fuel Weight (EZFW) by aggregating the following calculated components:

$$\text{EZFW} = \text{DOW} + \text{Total Pax Wt} + \text{Total Bag Wt} + \text{Cargo Wt} + \text{Service Wt} + \text{Total ULD Tare Wt (if applicable)}$$

Where:

| Abbreviation | Description |
| :--- | :--- |
| **DOW** | Dry Operating Weight (Aircraft-specific baseline weight, including crew and fluids). |
| **Total Pax Wt** | Total number of passengers multiplied by their average weight. |
| **Total Bag Wt** | Total number of bags multiplied by their average weight. |
| **Service Wt** | Weight of additional services (e.g., catering, potable water). |
| **Total ULD Tare**| Total container (ULD) weight required for baggage storage on wide-body aircraft. |

### Wide Body Calculation Logic

For wide-body aircraft, baggage is assumed to be loaded into Unit Load Devices (containers). The calculation determines the minimum number of ULDs required and accounts for the weight of the empty containers (tare weight):

1.  **Total Bags:** $\text{Total Pax} \times \text{Avg Bags Per Pax}$
2.  **Total ULDs (Rounded Up):** $\text{Ceiling}(\text{Total Bags} / \text{Avg Bags Per ULD})$
3.  **Total ULD Tare:** $\text{Total ULDs} \times \text{Avg ULD Tare Weight}$

### Final ZFW Determination

The final result, which is the required figure for the load sheet, applies a conservative rounding rule to the EZFW:

$$\text{ZFW (Final)} = \text{Ceiling}(\text{EZFW} / 100) \times 100$$

*(Example: EZFW of 95,350.45 rounds up to 95,400)*

## 🚀 Usage (How to Run Locally)

This is a static web application and requires no complex build process or external server (other than serving the files).

### Prerequisites

* A web browser (Chrome, Firefox, etc.)
* (Optional but recommended) A local web server to avoid browser security restrictions when loading local files (e.g., Python's `http.server`).

### Installation

1.  **Clone the repository:**
    ```bash
    git clone [YOUR_REPOSITORY_LINK]
    cd EZFW-Calculator
    ```
2.  **Ensure `aircraft.csv` is present:**
    The file must be in the root directory and formatted as: `Registration,DOW (kg/lbs),Type (Narrow/Wide)`
    *(Example Row: A6-ABC,120500,Wide)*
3.  **Run the application:**
    Open the `index.html` file directly in your browser, **OR** (recommended) start a simple local server:
    ```bash
    # For Python 3
    python3 -m http.server
    ```
    Then, navigate to `http://localhost:8000` (or the address provided by your server).

## 🗂️ Project Structure

| File | Purpose |
| :--- | :--- |
| `index.html` | The main structure, containing all input fields and the results table. |
| `script` tags | Contains the core JavaScript logic for data loading, calculating, and dynamic display updates. |
| `aircraft.csv`| External data source for aircraft registration, Dry Operating Weight (DOW), and Type. |

---

## ✍️ Built With

* **HTML5:** For structure and form elements.
* **CSS:** For presentation (implied by classes like `.container`).
* **Vanilla JavaScript:** For all dynamic calculations and DOM manipulation (no external frameworks or libraries).
