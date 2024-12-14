# Smart Cities Location Insights Platform

## Project Overview

This Smart Cities Location Insights Platform is a comprehensive tool designed to provide detailed demographic and infrastructure information for specific locations in Kenya. Users can explore and analyze local amenities, population statistics, and spatial characteristics at the county, constituency, and ward levels.

## Features

- **Granular Location Selection**: 
  - Drill down from County to Constituency to Ward levels
  - Intuitive interface for exploring administrative boundaries in Kenya

- **Comprehensive Amenities Mapping**:
  - Detailed inventory of local amenities
  - Categories include:
    * Educational Facilities (Schools, Colleges)
    * Healthcare Centers (Hospitals, Clinics)
    * Public Services (Police Stations, Post Offices)
    * Transportation Infrastructure
    * Recreational Spaces
    * Commercial Zones

- **Population Analytics**:
  - Total population for selected area
  - Demographic breakdown
  - Population density calculations
  - Amenities-to-population ratio analysis

## Prerequisites

- Python 3.8+
- Required Libraries:
  - `geopandas`
  - `pandas`
  - `streamlit`
  - `plotly`
  - `requests`

## Installation

1. Clone the repository:
   ```bash
   git clone https://github.com/VictorPaulArony/smart-planner.git
   cd smart-planner
   ```


## Usage

2. Run the website in the CLI for localhost:
   ```bash
   go run main.go
   ```
``

### User Interface Workflow
1. Select County from dropdown
2. Select Constituency within the County
3. Select specific Ward
4. View:
   - Total Population
   - Detailed Amenities List
   - Population-to-Amenities Ratio
   - Interactive Visualizations

## Data Sources

- Kenya National Bureau of Statistics
- Open Government Data
- Geospatial Databases
- Local Administrative Registers

## Data Privacy & Accuracy

- All data is sourced from official government sources
- Personal identification information is anonymized
- Regular data updates ensure accuracy

## Contributing

1. Fork the repository
2. Create your feature branch (`git checkout -b feature/AmazingFeature`)
3. Commit your changes (`git commit -m 'Add some AmazingFeature'`)
4. Push to the branch (`git push origin feature/AmazingFeature`)
5. Open a Pull Request

## License

Distributed under the MIT License. See `LICENSE` for more information.

## Contact

Project Maintainer: [Victor Paul Arony]
- Email: victorpaularony@gmail.com
- Project Link: [https://github.com/VictorPaulArony/smart-planner.git](https://github.com/VictorPaulArony/smart-planner.git)

## Acknowledgements

- Kenya National Bureau of Statistics
- Local Government Departments
- Open Data Initiatives

---

### Future Roadmap

- [ ] Machine Learning-based Predictive Analytics
- [ ] Real-time Amenities Updates
- [ ] Mobile Application Development
- [ ] Expanded Geographic Coverage
