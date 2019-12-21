/*
 *
 */
window.onload = function() {
  var ctx = document.getElementById('myChart').getContext('2d');
  var chart = new Chart(ctx, {
    // The type of chart we want to create
    type: 'line',

    // The data for our dataset
    data: {
      labels: ['January', 'February', 'March', 'April', 'May', 'June', 'July'],
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
      ],
    },

    // Configuration options go here
    options: {},
  });
};
