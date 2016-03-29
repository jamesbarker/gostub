# gostub
Demonstrates building a simple stub for the web service http://api.openweathermap.org

# Test Data
* Lat = -27.468 Brisbane (1 <= Cnt <= 3) 
* Lat = -10.500 Bad Request
* Lat = -10.404 Not Found
* All others Lats = Melbourne (Cnt is always 10)

# URLs:
* Real URL:
 http://api.openweathermap.org/data/2.5/find?appid=0a12b8f2f0dd011ed6085cb995ff61b4&lat=-37.814&lon=144.963&cnt=2
* Stubbed URL: 
 http://localhost:5001/data/2.5/find?appid=0a12b8f2f0dd011ed6085cb995ff61b4&lat=-37.814&lon=144.963&cnt=2		
* Count with if's:
 http://localhost:5001/data/2.5/find?appdid=0a12b8f2f0dd011ed6085cb995ff61b4&lat=-27.468&lon=144.963&cnt=3
* Bad Request:
 http://localhost:5001/data/2.5/find?appdid=0a12b8f2f0dd011ed6085cb995ff61b4&lat=-10.500&lon=144.963&cnt=1
* Not Found:
 http://localhost:5001/data/2.5/find?appdid=0a12b8f2f0dd011ed6085cb995ff61b4&lat=-10.404&lon=144.96&cnt=1
* All Others (Always 10):
 http://localhost:5001/data/2.5/find?appid=0a12b8f2f0dd011ed6085cb995ff61b4&lat=-37.814&lon=144.963&cnt=10
 
