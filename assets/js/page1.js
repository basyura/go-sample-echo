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
          fill: false,
          backgroundColor: 'green',
          borderColor: 'green',
          data: [9, 9, 9],
        },
        {
          fill: false,
          backgroundColor: 'orange',
          borderColor: 'orange',
          data: [null, null, 9, 9, 9, 9, 9],
        },
        {
          fill: false,
          backgroundColor: 'blue',
          borderColor: 'blue',
          data: [null, null, null, null, null, null, 9, 9, 9],
        },
        {
          fill: false,
          backgroundColor: 'red',
          borderColor: 'red',
          data: [null, null, null, null, null, null, null, null, 9, 9, 9],
        },

        {
          fill: false,
          backgroundColor: 'green',
          borderColor: 'green',
          data: [null, null, 8, 8],
        },
        {
          fill: false,
          backgroundColor: 'orange',
          borderColor: 'orange',
          data: [null, null, null, 8, 8],
        },
        {
          fill: false,
          backgroundColor: 'blue',
          borderColor: 'blue',
          data: [null, null, null, null, 8, 8, 8, 8],
        },
        {
          fill: false,
          backgroundColor: 'red',
          borderColor: 'red',
          data: [null, null, null, null, null, null, null, 8, 8],
        },
      ],
    },

    // Configuration options go here
    options: {
      legend: {
        display: false,
      },
      tooltips: {
        enabled: false,
      },
      scales: {
        yAxes: [
          {
            ticks: {
              display: false,
              min: 0,
              max: 10,
              stepSize: 1,
            },
          },
        ],
      },
    },
  };
  for (var i = 0; i < 100; i++) {
    params.data.labels.push((i * 50).toString());
  }
  var chart = new Chart(ctx, params);

  loadTable();
};

function loadTable() {
  var buf = '<table id="details" >';
  buf += '<tr><td>aaaaaaaaaaaaa</td><td>100</td></tr>';
  buf += '<tr><td>bbbbbbbbbbbbb</td><td>200</td></tr>';
  buf += '<tr><td>ccccccccccccc</td><td>300</td></tr>';
  buf += '<tr><td>ddddddddddddd</td><td>500</td></tr>';
  buf += '<tr><td>ddddddddddddd</td><td>500</td></tr>';
  buf += '<tr><td>ddddddddddddd</td><td>500</td></tr>';
  buf += '<tr><td>ddddddddddddd</td><td>500</td></tr>';
  buf += '<tr><td>ddddddddddddd</td><td>500</td></tr>';
  buf += '<tr><td>ddddddddddddd</td><td>500</td></tr>';
  buf += '<tr><td>ddddddddddddd</td><td>500</td></tr>';
  buf += '<tr><td>ddddddddddddd</td><td>500</td></tr>';
  buf += '</table>';

  var wrappr = document.getElementById('leftWrapper');
  wrappr.insertAdjacentHTML('afterbegin', buf);
}
