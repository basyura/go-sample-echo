/*
 *
 */
window.onload = function() {
  var ctx = document.getElementById('myChart').getContext('2d');

  params = {
    // The type of chart we want to create
    type: 'line',

    // The data for our dataset
    data: {
      labels: [],
      datasets: [
        {
          label: 'line1',
          fill: false,
          backgroundColor: 'rgb(255, 99, 132)',
          borderColor: 'rgb(255, 99, 132)',
          data: [0, 10, 5, 2, 20, 30, 45],
        },
        {
          label: 'line2',
          fill: false,
          backgroundColor: 'blue',
          borderColor: 'blue',
          data: [5, 15, 15, 25, 25, 30, 35],
        },
        {
          label: 'line3',
          fill: false,
          backgroundColor: 'green',
          borderColor: 'green',
          data: [5, 10, 10, 35, 15, 20, 35],
        },
        {
          label: 'line4',
          fill: false,
          backgroundColor: 'green',
          borderColor: 'green',
          data: [40, 40, 40, null, 40, 40, 40],
        },
      ],
    },

    // Configuration options go here
    options: {},
  };

  for (var i = 0; i < 100; i++) {
    params.data.labels.push(i.toString());
  }
  var chart = new Chart(ctx, params);
};
