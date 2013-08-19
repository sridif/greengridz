import ystockquote as ystock
import matplotlib.pyplot as plt
from numpy import arange,array,ones,linalg
from pylab import plot,show
import numpy as np

def getdata(stock,start,end):
  data = []
  pre_data = ystock.get_historical_prices(stock,start,end)

  for i in pre_data[1:]:
    data.append(i[-1])

  return data


def plotdata1(data):
  plt.plot(range(len(data)), data)

def predictdata(data):
  print len(data)/7
  avg = [0,0,0, 0,0,0,0]
  for i in range(len(data)/7 ):
    for j in range(7):
      avg[j] = avg[j] + float(data[i*7 + j])
  
  for j in range(7):
    avg[j] = avg[j] *7 / len(data) 
  return data[:-7] + avg 

def main():

  # two type -- regression and time series.
  # importance not avraging out instances.
  # implies scale to millions
  # second type
  # represent on maps .. global maps.

  regression()  
  forecast()

def regression():
  '''
  data = getdata('GOOG', '2013-04-01', '2013-05-01')
  data2 =getdata('AAPL', '2013-04-01', '2013-05-01')

  A = array([ arange(0,len(data)) , ones(len(data))])
  y = data2

  w = linalg.lstsq(A.T, y)[0]
  http://glowingpython.blogspot.se/2012/03/linear-regression-with-numpy.html  
  '''
  #show()  
  data = getdata('GOOG', '2013-04-01', '2013-05-01')
  data2 = getdata('AAPL', '2013-04-01', '2013-05-01')  
  data3 = getdata('AMZN', '2013-04-01', '2013-05-01')  
  
  
  floatdata = []
  for d in range(len(data2)):
    floatdata.append(str( float(data2[d]) + float(data3[d]) ) )
  
  
  w1 = getweight(data,data2)
  # plotting the line
  w2 = getweight(data,data3)

  w3 = getweight(data,floatdata)
  show()

def getweight(data,data2):

  floatdata = []
  for d in data2:
    floatdata.append(float(d))
  

  xi = np.array(floatdata)
  A = array([ xi, ones(len(data))])
  # linearly generated sequence

  w = linalg.lstsq(A.T,data)[0] # obtaining the parameters
  line = w[0]*xi+w[1] # regression line
  plot(xi,line,'r-',xi,data,'o')

  return w


  '''
  data = getdata('GOOG', '2013-04-01' , '2013-05-01')
  plotdata1(data)
  plt.show()
  '''
  
def forecast():
  data = getdata('GOOG', '2013-04-01' , '2013-05-01')
  plotdata1(data)
  data2 = predictdata(data)
  plotdata1(data2)

  data = getdata('AAPL', '2013-04-01' , '2013-05-01')
  plotdata1(data)
  data2 = predictdata(data)
  plotdata1(data2)

  data = getdata('AMZN', '2013-04-01' , '2013-05-01')
  plotdata1(data)
  data2 = predictdata(data)
  plotdata1(data2)

  plt.show()

if __name__ == "__main__" :
  main()

