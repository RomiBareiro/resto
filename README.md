# Delivery API
We're developing a simplified version of a restaurant API for real-time queries. The API takes user location (latitude and longitude) as input and responds with a list of restaurant IDs that can deliver orders to the user at the time of the query, considering each restaurant's location, delivery radius, and operating hours. There's no defined order for the returned IDs in case of multiple matches.

## Problem Description
We need to build an API that can handle real-time queries based on user location. Given the latitude and longitude of the user, the API should return a list of restaurant IDs that can fulfill orders for the user at that moment, considering factors such as the restaurant's location, delivery radius, and operating hours. The list should be dynamically generated based on the current time of the query.

## Data Format
Restaurant information is available in a CSV file with the following columns: (In this example csv is in template folder)

```code
id: Restaurant ID
latitude: Latitude of the restaurant's location
longitude: Longitude of the restaurant's location
availability_radius: Delivery radius in kilometers
open_hour: Start time of delivery hours in ISO format
close_hour: End time of delivery hours in ISO format
rating: Restaurant rating, a number between 1 and 5
```

## API Development
The API should provide the following functionality:

- Receive a request with the user's location.
- Determine which restaurants can deliver orders to the user at the time of the query.
- Return a list of restaurant IDs that meet the delivery criteria.

## Usage

### Start server

```code
sudo docker-compose build 
sudo docker-compose up
```

### Send request: 
```code
{
    "latitude": 40.7128, //user location
    "longitude" : -74.0060
}
```

### Example response:
```code
{
    "ids": [
        "1", 
        "312"
    ]
}
```