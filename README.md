# README.md

# Event Visualization App

## Overview

This project is a web-based application that visualizes real-time events sourced from NASA's EONET (Earth Observatory Natural Event Tracker) API. The application uses CesiumJS for 3D mapping and categorizes events such as wildfires, storms, and earthquakes, displaying them interactively on a world map.

---

## Features

- **Event Categorization**: Categorizes and visualizes events based on their type (e.g., Wildfires, Earthquakes).
- **Real-Time Updates**: Fetches data every 10 seconds to keep the event information up to date.
- **Interactive Map**: Displays events using CesiumJS with custom icons for different event types.
- **Web API**: Includes endpoints for fetching categorized events and last fetch timestamps.

---

## Tech Stack

- **Frontend**: 
  - HTML, JavaScript, CesiumJS
- **Backend**: 
  - Go (Gin Framework)
- **Deployment**: 
  - Docker

---

## Prerequisites

1. **Go**: Ensure Go is installed (version 1.21+).
2. **Docker**: Install Docker for containerized deployment.
3. **Token/Key**: Get a token from Cesium and a key from Nasa (See step 2)

---

## Setup and Installation

### 1. Clone the Repository
```bash
git clone https://github.com/Lach-dev/NaturalEvents.git
cd NaturalEvents
```

### 2. Set Up the API Key
Replace `YOUR_API_KEY` in `main.go` with your NASA EONET API key. Obtain one from [NASA EONET](https://eonet.gsfc.nasa.gov/).
Replace `YOUR_ACCESS_TOKEN` in `index.html` with your cesium token. Obtain one from [Cesium](https://ion.cesium.com/).

### 3. Build and Run Locally

#### a. Run Locally with Go
```bash
go mod download
go run main.go
```
Access the app at `http://localhost:8080`.

#### b. Build and Run with Docker
```bash
docker build -t NaturalEvents .
docker run -p 8080:8080 NaturalEvents
```
Access the app at `http://localhost:8080`.

---

## API Endpoints

### 1. **Fetch Categorized Events**
- **GET** `/api/categorize`
- **Response**:
  ```json
  {
      "events": [
          {
              "title": "Event Title",
              "categories": [{"title": "Category Title"}],
              "geometry": [{"coordinates": [longitude, latitude]}]
          }
      ]
  }
  ```

### 2. **Fetch Last Data Fetch Time**
- **GET** `/api/lastFetchTime`
- **Response**:
  ```json
  {
      "lastFetchTime": "2024-12-07T12:34:56Z"
  }
  ```

---

## File Structure

```
.
├── main.go               # Backend implementation
├── index.html            # Frontend with CesiumJS integration
├── static/               # Static assets (icons)
├── templates/            # HTML templates
├── Dockerfile            # Docker build file
└── README.md             # Project documentation
```

---

## Known Issues

1. **Commented Code**: Some functions (e.g., `categorizeEvents`) contain unused code.
2. **Fetch Logic**: The fetch interval might require adjustment for higher efficiency.

---

## Future Improvements

1. Implement detailed caching strategies for event data.
2. Add user interaction features like filtering events by category or location.
3. Enhance error handling for API responses.

---

## Acknowledgments

- **NASA EONET API** for providing real-time event data.
- **CesiumJS** for interactive 3D visualization.
- **Gin Framework** for the web API.

Feel free to contribute or report issues in the repository!