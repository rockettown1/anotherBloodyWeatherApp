![Weather App Demo](./demo.gif)

REST client using Go's built in http package.

The first API call (to Mapbox) takes a location string and returns the latitude and longitude for that place. The lat and lng coordinates are then sent to the Darksky API to return the weather.

