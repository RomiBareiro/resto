# Delivery API
We're developing a simplified version of a restaurant API for real-time queries. The API takes user location (latitude and longitude) as input and responds with a list of restaurant IDs that can deliver orders to the user at the time of the query, considering each restaurant's location, delivery radius, and operating hours. There's no defined order for the returned IDs in case of multiple matches.

## Problem Description
We need to build an API that can handle real-time queries based on user location. Given the latitude and longitude of the user, the API should return a list of restaurant IDs that can fulfill orders for the user at that moment, considering factors such as the restaurant's location, delivery radius, and operating hours. The list should be dynamically generated based on the current time of the query.

## Data Format
Restaurant information is available in a CSV file with the following columns: (In this example csv is in template folder)

| Field  | Meaning  |
| ------------ | ------------ |
|  id |  Restaurant ID |
|  latitude |Latitude of the restaurant's location   |
|longitude   | Longitude of the restaurant's location  |
|  availability_radius |   Delivery radius in kilometers|
|  close_hour |  End time of delivery hours in ISO format |
|  open_hour | Start time of delivery hours in ISO format  |
|  rating | Restaurant rating, a number between 1 and 5 |

Every 6 hs, we download this csv file and update a table with merchant information to query our requests

## API Development
The API should provide the following functionality:

- Receive a request with the user's location.
- Determine which restaurants can deliver orders to the user at the time of the query.
- Return a list of restaurant IDs that meet the delivery criteria:
  * Restaurant is open and its closing time is > to 10 min from current time
  * Restaurant availability radius < 5km

## Usage
Retrieves available merchant/restaurant IDs from the Jampp platform based on geographic coordinates. This is useful for discovering nearby merchants or restaurants at a given location.

**Base URL:** Uses the `{{local_url}}` environment variable — make sure it is set in your active environment before sending the request.

**Request Body (raw JSON)**

| Parameter | Type | Description |
|-----------|------|-------------|
| `latitude` | number | Latitude of the target location (e.g., `40.7128` for New York City) |
| `longitude` | number | Longitude of the target location (e.g., `-74.0060` for New York City) |

**Notes**
- If no merchants are available at the specified coordinates, the API returns a `404` response with `{ "error": "not_found", "cause": "no available merchants" }`.
- Ensure the coordinates are valid and within a supported region to get results.

### Start server

```code
sudo docker-compose up --build
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

### Stop server
```code
sudo docker-compose down
```
