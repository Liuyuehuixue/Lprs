import convolution from './convolution'

const sobelVertical = function (pixels, width, height) {
  const divider = 4;
  const kernel = [1 / divider, 0, -1 / divider, 2 / divider, 0, -2 / divider, 1 / divider, 0, -1 / divider];

  pixels = convolution(pixels, kernel, width, height);
  return pixels
}

export default sobelVertical
