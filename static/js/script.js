const TAU = 2 * Math.PI;

document.addEventListener('DOMContentLoaded', (event) => {
main();
})
const main = () => {
  const ctx = document.getElementById("drawing").getContext('2d');
  drawGrid(ctx, 1, 1, 15, 13, {
    radius: 20,
    inset: 2,
    randomColors: generateColors(30, 1.0, 0.667)
  });
};

const defaultGridOptions = {
  radius: 10,
  sides: 6,
  inset: 0,
  // Context
  lineWidth: 1,
  fillStyle: '',
  strokeStyle: 'black',
  // Other
  randomColors: null
};

const drawGrid = (ctx, x, y, w, h, options = {}) => {
  const opts = { ...defaultGridOptions, ...options };
  const points = createPoly(opts);
  opts.diameter = opts.radius * 2;
  for (let gy = y; gy < y + h; gy++) {
    for (let gx = x; gx < x + w; gx++) {
      ctx.fillStyle = opts.randomColors ? pickRandom(opts.randomColors) : opts.fillStyle;
      drawPoly(ctx, gridToPixel(gx, gy, opts), points, opts);
    }
  }
};

const gridToPixel = (gridX, gridY, opts) => {
  const m = gridMeasurements(opts);
  return toPoint(
    Math.floor(gridX * m.gridSpaceX),
    Math.floor(gridY * m.gridSpaceY + (gridX % 2 ? m.gridOffsetY : 0))
  );
};

const drawPoly = (ctx, origin, points, opts) => {
  ctx.strokeStyle = opts.strokeStyle;
  ctx.save();
  ctx.translate(origin.x, origin.y);
  polyPath3(ctx, points);
  ctx.restore();
  if (opts.lineWidth) ctx.lineWidth = opts.lineWidth;
  if (opts.fillStyle || opts.randomColors) ctx.fill();
  if (opts.strokeStyle) ctx.stroke();
};

const createPoly = (opts, points = []) => {
  const
    { inset, radius, sides } = opts,
    size = radius - inset,
    step = TAU / sides;
  for (let i = 0; i < sides; i++) {
    points.push(toPolarCoordinate(0, 0, size, step * i));
  }
  return points;
};

const gridMeasurements = (opts) => {
  const
    { diameter, inset, radius, sides } = opts,
    edgeLength = Math.sin(Math.PI / sides) * diameter,
    gridSpaceX = diameter - edgeLength / 2,
    gridSpaceY = Math.cos(Math.PI / sides) * diameter,
    gridOffsetY = gridSpaceY / 2;
  return {
    diameter,
    edgeLength,
    gridSpaceX,
    gridSpaceY,
    gridOffsetY
  };
};

/** @unused */
const polyPath = (ctx, x, y, radius, sides = 3) => {
  ctx.beginPath();
  ctx.moveTo(...fromPoint(toPolarCoordinate2(x, y, radius)));
  for (let i = 1; i <= sides; i += 1) {
    ctx.lineTo(...fromPoint(toPolarCoordinate2(x, y, radius, sides, i)));
  }
  ctx.closePath();
};

/** @unused */
const polyPath2 = (ctx, points = []) => {
  ctx.beginPath();
  ctx.moveTo(points[0], points[1]);
  for (let i = 2; i < points.length - 1; i += 2) {
    ctx.lineTo(points[i], points[i + 1]);
  }
  ctx.closePath();
};

const polyPath3 = (ctx, points = []) => {
  const [{ x: startX, y: startY }] = points;
  ctx.beginPath();
  ctx.moveTo(startX, startY);
  points.forEach(({ x, y }) => { ctx.lineTo(x, y); });
  ctx.closePath();
};

const pickRandom = (arr) => arr[Math.floor(Math.random() * arr.length)];

const toPoint = (x, y) => ({ x, y });

const fromPoint = ({ x, y }) => [ x, y ];

const toPolarCoordinate = (centerX, centerY, radius, angle) => ({
  x: centerX + radius * Math.cos(angle),
  y: centerY + radius * Math.sin(angle)
});

const toPolarCoordinate2 = (centerX, centerY, radius, sides, i) =>
  toPolarCoordinate(centerX, centerY, radius, i === 0 ? 0 : (i * TAU / sides));

const generateColors = (count, saturation = 1.0, lightness = 0.5, alpha = 1.0) =>
  Array.from({ length: count }, (_, i) =>
    `hsla(${[
      Math.floor(i / count * 360),
      `${Math.floor(saturation * 100)}%`,
      `${Math.floor(lightness * 100)}%`,
      alpha
    ].join(', ')})`);


    function getItems() {
      // Specify the API endpoint
      const apiUrl = 'http://localhost:3000/items';
    
      // Make a GET request to the API
      fetch(apiUrl)
        .then(response => {
          // Check if the request was successful (status code 200)
          if (!response.ok) {
            throw new Error(`HTTP error! Status: ${response.status}`);
          }
    
          // Parse the JSON response
          return response.json();
        })
        .then(data => {
          // Handle the retrieved items
          console.log('Items:', data);
          // You can perform further processing with the items here
        })
        .catch(error => {
          // Handle any errors that occurred during the fetch
          console.error('Fetch error:', error);
        });
    }
    
    // Call the function to get items when needed
    console.log(getItems());
