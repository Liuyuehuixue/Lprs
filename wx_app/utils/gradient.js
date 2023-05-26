import sobelHorizontal from './operations/sobelHorizontal'
import sobelVertical from './operations/sobelVertical'

const gradient = function (pixels,width,height) {

  const deltaX = sobelHorizontal(pixels,width,height);
  const deltaY = sobelVertical(pixels,width,height);

  var outputVectors = new Array(height);

  for (let y = 0; y < height; y++) {
	outputVectors[y] = new Array(width);
    for (let x = 0; x < width; x++) {
      let dstOff = (y * width + x) * 4;

      var r = Math.sqrt(Math.pow(deltaX[dstOff], 2) + Math.pow(deltaY[dstOff], 2));
      var g = Math.sqrt(Math.pow(deltaX[dstOff + 1], 2) + Math.pow(deltaY[dstOff + 1], 2));
      var b = Math.sqrt(Math.pow(deltaX[dstOff + 2], 2) + Math.pow(deltaY[dstOff + 2], 2));

      var dr = Math.atan2(deltaY[dstOff], deltaX[dstOff]);
      var dg = Math.atan2(deltaY[dstOff + 1], deltaX[dstOff + 1]);
      var db = Math.atan2(deltaY[dstOff + 2], deltaX[dstOff + 2]);
	  
	  outputVectors[y][x] = {
          mag: (r * 0.299 + g * 0.587+ b * 0.114),
          orient: (dr * 0.299 + dg * 0.587+ db * 0.114)
        }
    }
  }

  return outputVectors
}

export default gradient
