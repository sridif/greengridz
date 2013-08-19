function button1(){
 val = document.getElementById("user_ida").value;
 console.log(document.getElementById("user_ida").value);
 
 $.ajax({
    url: 'http://api.greengridz.com/ice/re',
    type: 'GET',
    datatype: 'json',
    data: 'user_id=' + val + '&op=usage' , 
    success: function(json){ handle2(json.d.Output); },
    error : handle2('0,2,0,0,0,0,0,0')
});


}


function button2(){

  val = document.getElementById("user_ida").value;
  console.log(document.getElementById("user_ida").value);
 
  $.ajax({
    url: 'http://api.greengridz.com/ice/re',
    type: 'GET',
    datatype: 'json',
    data: 'user_id=' + val + '&op=usage' , 
    success: function(json){ handle3(json.d.Output);},
    error : handle3('0,0,0,0')
  });

}

function findLineByLeastSquares(values_x, values_y) {
    var sum_x = 0;
    var sum_y = 0;
    var sum_xy = 0;
    var sum_xx = 0;
    var count = 0;
    
    /*
     * We'll use those variables for faster read/write access.
     */

    var x = 0;
    var y = 0;

    /*var values_length = values_x.length;
    if (values_length != values_y.length) {
        throw new Error('The parameters values_x and values_y need to have same size!');
    }*/

    values_length=10
    
    /*
     * Nothing to do.
     */
    if (values_length === 0) {
        return [ [], [] ];
    }
    
    /*
     * Calculate the sum for each of the parts necessary.
     */
    for (var v = 0; v < values_length; v++) {
        x = values_x[v];
        y = values_y[v];
        sum_x += x;
        sum_y += y;
        sum_xx += x*x;
        sum_xy += x*y;
        count++;
    }
    
    /*
     * Calculate m and b for the formular:
     * y = x * m + b
     */
    var m = (count*sum_xy - sum_x*sum_y) / (count*sum_xx - sum_x*sum_x);
    var b = (sum_y/count) - (m*sum_x)/count;
    
    /*
     * We will make the x and y result line now
     */
    var result_values_x = [];
    var result_values_y = [];
    
    for (var v = 0; v < values_length; v++) {
        x = values_x[v];
        y = x * m + b;
        result_values_x.push(x);
        result_values_y.push(y);
    }
    
    return[ m, b];
}

function handle3(user_data) {


  var arr = user_data[0].split(",");

  Data = [];
    console.log('user_data');
    console.log(user_data);
  for (i=0; i< 20; i ++ ) {
	       console.log(arr[i]);
	       Data.push(parseFloat(arr[i]));
  }
  
    cost =[10,102,43,3,2,3,45,6,7,8,3,4,5,6];
    co2 = [1,2,3,4,5,6,20,8,9,20,23,3,5,2];

    SData = findLineByLeastSquares(Data, co2 )   ;
   console.log(SData);
  Scatter('#container1', SData);

}

function handle2(user_data) {

  var Data = [0,0,0,0,0,0];
  var arr = user_data[0].split(",");

  Data = [];
  for (i=0; i< 20; i ++ ) {
	       console.log(arr[i]);
	       Data.push(parseFloat(arr[i]));
  }

  console.log(Data);
  Line('#container0', Data);

}

function Line(name,Data){

        $(name).highcharts({
            title: {
                text: 'Daily power consumption data',
                x: -20 //center
            },
            subtitle: {
                text: 'Source: greengridz.com',
                x: -20
            },
            xAxis: {
               // categories: ['Jan', 'Feb', 'Mar', 'Apr', 'May', 'Jun', 'Jul', 'Aug', 'Sep', 'Oct', 'Nov', 'Dec']
            },
            yAxis: {
                title: {
                    text: 'Temperature (°C)'
                },
                plotLines: [{
                    value: 0,
                    width: 1,
                    color: '#808080'
                }]
            },
            tooltip: {
                valueSuffix: '°C'
            },
            legend: {
                layout: 'vertical',
                align: 'right',
                verticalAlign: 'middle',
                borderWidth: 0
            },
            series: [{
                name: 'user data',
                data: Data 
            }
            
            ]
        });
   
    

}
function Scatter(name,SData){

  $(name).highcharts({
            chart: {
            },
            xAxis: {
                min: -0.5,
                max: 5.5
            },
            yAxis: {
                min: 0
            },
            title: {
                text: 'Scatter plot with regression line'
            },
            series: [{
                type: 'line',
                name: 'Regression Line',
                data: [ SData[0]*0 + SData[1] , SData[0]*1 + SData[1] ] ,
                marker: {
                    enabled: false
                },
                states: {
                    hover: {
                        lineWidth: 0
                    }
                },
                enableMouseTracking: false
            }, {
                type: 'scatter',
                name: 'Observations',
                data: [ ],
                marker: {
                    radius: 4
                }
            }]
        });
}
    



function Line1(name){

        $(name).highcharts({
            title: {
                text: 'Monthly Average Temperature',
                x: -20 //center
            },
            subtitle: {
                text: 'Source: WorldClimate.com',
                x: -20
            },
            xAxis: {
                categories: ['Jan', 'Feb', 'Mar', 'Apr', 'May', 'Jun',
                    'Jul', 'Aug', 'Sep', 'Oct', 'Nov', 'Dec']
            },
            yAxis: {
                title: {
                    text: 'Temperature (°C)'
                },
                plotLines: [{
                    value: 0,
                    width: 1,
                    color: '#808080'
                }]
            },
            tooltip: {
                valueSuffix: '°C'
            },
            legend: {
                layout: 'vertical',
                align: 'right',
                verticalAlign: 'middle',
                borderWidth: 0
            },
            series: [{
                name: 'Tokyo',
                data: [7.0, 6.9, 9.5, 14.5, 18.2, 21.5, 25.2, 26.5, 23.3, 18.3, 13.9, 9.6]
            }, {
                name: 'New York',
                data: [-0.2, 0.8, 5.7, 11.3, 17.0, 22.0, 24.8, 24.1, 20.1, 14.1, 8.6, 2.5]
            }, {
                name: 'Berlin',
                data: [-0.9, 0.6, 3.5, 8.4, 13.5, 17.0, 18.6, 17.9, 14.3, 9.0, 3.9, 1.0]
            }, {
                name: 'London',
                data: [3.9, 4.2, 5.7, 8.5, 11.9, 15.2, 17.0, 16.6, 14.2, 10.3, 6.6, 4.8]
            }]
        });
   
    

}
function Scatter1(name){

  $(name).highcharts({
            chart: {
            },
            xAxis: {
                min: -0.5,
                max: 5.5
            },
            yAxis: {
                min: 0
            },
            title: {
                text: 'Scatter plot with regression line'
            },
            series: [{
                type: 'line',
                name: 'Regression Line',
                data: [[0, 1.11], [5, 4.51]],
                marker: {
                    enabled: false
                },
                states: {
                    hover: {
                        lineWidth: 0
                    }
                },
                enableMouseTracking: false
            }, {
                type: 'scatter',
                name: 'Observations',
                data: [1, 1.5, 2.8, 3.5, 3.9, 4.2],
                marker: {
                    radius: 4
                }
            }]
        });
}
    


