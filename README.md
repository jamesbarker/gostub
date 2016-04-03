# gostub
Demonstrates building a simple stub for the web service http://api.openweathermap.org

# Test Data
* Lat = -27.468 Brisbane (1 <= Cnt <= 3) 
* Lat = -33.867 Sydney (Delay 10s) 
* Lat = -10.500 Bad Request
* Lat = -10.404 Not Found
* All others Lats = Melbourne (Cnt is always 10)

# URLs:
* Real URL:
 http://api.openweathermap.org/data/2.5/find?appid=0a12b8f2f0dd011ed6085cb995ff61b4&lat=-37.814&lon=144.963&cnt=2
* Count with if's:
 http://localhost:5001/data/2.5/find?appdid=0a12b8f2f0dd011ed6085cb995ff61b4&lat=-27.468&lon=153.028&cnt=3
* Delay 5s
 http://localhost:5001/data/2.5/find?appid=0a12b8f2f0dd011ed6085cb995ff61b4&lat=-33.867&lon=151.207&cnt=10 
* Bad Request:
 http://localhost:5001/data/2.5/find?appdid=0a12b8f2f0dd011ed6085cb995ff61b4&lat=-10.500&lon=100.100&cnt=1
* Not Found:
 http://localhost:5001/data/2.5/find?appdid=0a12b8f2f0dd011ed6085cb995ff61b4&lat=-10.404&lon=100.100&cnt=1
* All Others (Always 10):
 http://localhost:5001/data/2.5/find?appid=0a12b8f2f0dd011ed6085cb995ff61b4&lat=-37.814&lon=144.963&cnt=10
 