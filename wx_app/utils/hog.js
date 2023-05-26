import gradient from './gradient';
var norms = require("./norms");

module.exports = {
  extractHOG: extractHOG,
  extractHistograms: extractHistograms,
  extractHOGFromHistograms: extractHOGFromHistograms
}

function extractHOG(pixels,width,height,options) {
  var histograms = extractHistograms(pixels,width,height,options);
  return extractHOGFromHistograms(histograms, options);
}

function extractHistograms(pixels,width,height, options) {
  var vectors = gradient(pixels,width,height);


  var cellSize = options.cellSize || 4;
  var bins = options.bins || 6;

  var cellsWide = Math.floor(vectors[0].length / cellSize);
  var cellsHigh = Math.floor(vectors.length / cellSize);

  var histograms = new Array(cellsHigh);

  for (var i = 0; i < cellsHigh; i++) {
    histograms[i] = new Array(cellsWide);

    for (var j = 0; j < cellsWide; j++) {
      histograms[i][j] = getHistogram(vectors, j * cellSize, i * cellSize,
                                      cellSize, bins);
    }
  }
  return histograms;
}

/* This function is decoupled so you can extract histograms for an entire image
 * to save recomputation, and use these to get HOGs from individual windows
 */
function extractHOGFromHistograms(histograms, options) {
  var blockSize = options.blockSize || 2;
  var blockStride = options.blockStride || (blockSize / 2);
  var normalize = norms[options.norm || "L2"];

  var blocks = [];
  var blocksHigh = histograms.length - blockSize + 1;
  var blocksWide = histograms[0].length - blockSize + 1;

  for (var y = 0; y < blocksHigh; y += blockStride) {
    for (var x = 0; x < blocksWide; x += blockStride) {
      var block = getBlock(histograms, x, y, blockSize);
      normalize(block);
      blocks.push(block);
    }
  }
  return Array.prototype.concat.apply([], blocks);
}

function getBlock(matrix, x, y, blockSize) {
  var square = [];
  for (var i = y; i < y + blockSize; i++) {
    for (var j = x; j < x + blockSize; j++) {
      square.push(matrix[i][j]);
    }
  }
  return Array.prototype.concat.apply([], square);
}

function getHistogram(elements, x, y, size, bins) {
  var histogram = zeros(bins);
  var bin,softbin;

  for (var i = 0; i < size; i++) {
    for (var j = 0; j < size; j++) {
      var vector = elements[y + i][x + j];
      bin,softbin = binFor(vector.orient, bins);
      histogram[bin] += vector.mag;
	  histogram[softbin] += vector.mag*0.5;
    }
  }
  return histogram;
}

function binFor(radians, bins) {
  var angle = radians * (180 / Math.PI);
  if (angle < 0) {
    angle += 180;
  }

  // center the first bin around 0
  //angle += 90 / bins;
  angle %= 180;

  var bin = Math.floor(angle / 180 * bins);
  if(angle-180/bins*bin > 90/bins)
	var softBin = (bin+1)%bins
  else
	var softBin = (bin-1)%bins
  return bin,softBin;
}

function zeros(size) {
  var array = new Array(size);
  for (var i = 0; i < size; i++) {
    array[i] = 0;
  }
  return array;
}
